package main

import (
	"container/heap"
	"context"
	"fmt"
	"log"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

func main()  {
	tw := NewTimeWheel(time.Second*6)
	tw.AddNode(time.Second*3, func() {
		log.Println("now")
	})
	tw.AddNode(time.Second*3, func() {
		log.Println("now1")
	})
	tw.AddNode(time.Second*4, func() {
		log.Println("now1")
	})

	tw.Start()
}

//http://russellluo.com/2018/10/golang-implementation-of-hierarchical-timing-wheels.html
//层级时间轮实现

//1.实现一个优先级队列
type item struct {
	Value    interface{}
	Priority int64
	Index    int
}
type priorityQueue []*item
func (this priorityQueue)Len() int {
	return len(this)
}
func (this priorityQueue)Less(i, j int) bool {
	return this[i].Priority < this[j].Priority
}
func (this priorityQueue)Swap(i, j int)  {
	this[i], this[j] = this[j], this[i]
	this[i].Index = i
	this[j].Index = j
}
func (this *priorityQueue)Push(x interface{})  {
	n := len(*this)
	xItem := x.(*item)
	xItem.Index = n
	*this = append(*this,xItem)

	/*
	n := len(*this)
	c := cap(*this)

	if n + 1 > c {
		npq := make(priorityQueue,n,c*2)
		copy(npq,*this)
		*this = npq
	}

	//插入最后一个元素,索引为n
	*this = (*this)[0:n+1]
	xItem := x.(*item)
	xItem.Index = n

	(*this)[n] = xItem*/
}
func (this *priorityQueue)Pop() interface{} {
	n := len(*this)
	c := cap(*this)

	if n < (c/2) && c > 25 {
		npq := make(priorityQueue,n,c/2)
		copy(npq,*this)
		*this = npq
	}

	//弹出最后一个元素
	item := (*this)[n-1]
	item.Index = -1
	*this = (*this)[0:n-1]
	return item
}
func (this *priorityQueue)PeekAndShift(max int64) (*item,int64) {
	if this.Len() == 0 {
		return nil,0
	}

	item := (*this)[0]
	if item.Priority > max {
		return nil,item.Priority - max
	}
	heap.Remove(this,0)

	return item, 0
}
func newPriorityQueue(capacity int) priorityQueue {
	return make(priorityQueue, 0, capacity)
}

//2.实现一个延迟队列
type DelayQueue struct {
	C chan interface{}
	mu sync.Mutex
	pq priorityQueue

	sleeping int32
	wakeupC chan struct{}
}

func NewDelayQueue(size int) *DelayQueue {
	return &DelayQueue{
		C:        make(chan interface{}),
		pq:       newPriorityQueue(size),
		wakeupC:  make(chan struct{}),
	}
}

//根据时间插入优先级队列,插入延迟队列
func (this *DelayQueue)Offer(elem interface{},expiration int64)  {
	newItem := &item{
		Value:    elem,
		Priority: expiration,
	}

	this.mu.Lock()
	heap.Push(&(this.pq),newItem)
	index := newItem.Index
	this.mu.Unlock()

	if index == 0 {
		//unlock
		if atomic.CompareAndSwapInt32(&this.sleeping,1,0) {
			this.wakeupC <- struct{}{}
		}
	}
}

//阻塞等待过期
// Poll starts an infinite loop, in which it continually waits for an element
// to expire and then send the expired element to the channel C.
func (this *DelayQueue)Poll(exitC chan struct{},nowF func() int64)  {
	for {
		now := nowF()

		this.mu.Lock()

		//pop一个任务
		item, delta := this.pq.PeekAndShift(now)

		//如果里面没有任务了
		if item == nil {
			// No items left or at least one item is pending.
			// We must ensure the atomicity of the whole operation, which is
			// composed of the above PeekAndShift and the following StoreInt32,
			// to avoid possible race conditions between Offer and Poll.
			atomic.StoreInt32(&this.sleeping,1)
		}
		this.mu.Unlock()

		if item == nil {
			if delta == 0 {
				select {
				case <-this.wakeupC:
					continue
				case <-exitC:
					goto exit
				}
			} else if delta > 0 {
				select {
				case <-this.wakeupC:
					continue
				case <-time.After(time.Duration(delta) * time.Millisecond):
					if atomic.SwapInt32(&this.sleeping,0) == 0 {
						<-this.wakeupC
					}
					continue
				case <-exitC:
					goto exit
				}
			}
		}

		select {
		case this.C <- item.Value:
			//成功弹出了过期的任务
		case <-exitC:
			goto exit
		}
	}

exit:
	atomic.StoreInt32(&this.sleeping,0)
}



//简单时间轮实现
type timeWheel struct {
	ctx context.Context
	cel context.CancelFunc
	stopCh  chan bool
	slots   []slot
	curSlot int
	tickDur time.Duration
}
type slot struct {
	head *node
	sLen  int
}
type node struct {
	round int     //转多少轮
	task func()
	prev *node
	next *node
}

func NewTimeWheel(timeout time.Duration) *timeWheel {
	ctx, cancelFunc := context.WithTimeout(context.Background(),timeout)
	return &timeWheel{
		ctx     : ctx,
		cel     : cancelFunc,
		stopCh  : make(chan bool),
		slots   : make([]slot,60),  //一共有60个槽
		curSlot : 0,                //初始为第0号槽
		tickDur : time.Second,      //1s 移动 1个槽
	}
}

func (tw *timeWheel)AddNode(dur time.Duration,task func())  {
	//1/60
	slotIndex := int(math.Ceil(dur.Seconds()))%len(tw.slots) //0
	newNode := &node{
		round: int(math.Ceil(dur.Seconds()))/len(tw.slots),
		task:  task,
		prev:  nil,
		next:  tw.slots[slotIndex].head,
	}
	tw.slots[slotIndex].head = newNode
}

func (tw *timeWheel)Start()  {
	for {
		select {
		case <-tw.stopCh:
			return
		case <-tw.ctx.Done():
			tw.cel()
			log.Println(tw.ctx.Err())
			return
		default:
			fmt.Println("current slot:",tw.curSlot)
			tw.checkSlot()
			time.Sleep(tw.tickDur)
		}
	}
}
func (tw *timeWheel)Stop()  {
	tw.stopCh <- true
}

func (tw *timeWheel)checkSlot()  {
	tw.curSlot = tw.curSlot%60
	curSlot   := tw.slots[tw.curSlot]
	curNode   := curSlot.head
	for curNode != nil  {
		curNode.round--
		if curNode.round < 0 {
			curNode.task()
			curNode.prev = curNode.next
		}
		curNode = curNode.next
	}
	tw.curSlot++
}
