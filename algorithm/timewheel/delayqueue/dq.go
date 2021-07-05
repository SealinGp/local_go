package delayqueue

import (
	"container/heap"
	"sync"
	"sync/atomic"
	"time"
)

//http://russellluo.com/2018/10/golang-implementation-of-hierarchical-timing-wheels.html
//层级时间轮实现

func New(wheelSize int) *DelayQueue {
	return &DelayQueue{
		C:       make(chan interface{}),
		pq:      newPriorityQueue(wheelSize),
		wakeupC: make(chan struct{}),
	}
}

//DelayQueue 是一个包含延迟任务元素的非阻塞队列
//该任务元素只能在过期时被使用
//队列头部元素是过期时间最长的任务元素
type DelayQueue struct {
	C chan interface{}

	mu sync.Mutex
	pq priorityQueue

	//跟 runtime.timers的 sleeping state 相似
	sleeping int32
	wakeupC  chan struct{}
}

//吧元素插入当前延迟队列中
func (dq *DelayQueue) Offer(elem interface{}, expiration int64) {
	item := &item{
		Value:    elem,
		Priority: expiration,
	}
	dq.mu.Lock()
	heap.Push(&dq.pq, item)
	index := item.Index
	dq.mu.Unlock()

	if index == 0 {
		//如果在sleeping == 1 ,则将其修改为0,并返回true
		if atomic.CompareAndSwapInt32(&dq.sleeping, 1, 0) {
			dq.wakeupC <- struct{}{}
		}
	}
}

//Poll函数初始化延迟队列的循环,并等待该队列中的一个元素过期后,将其元素发送到channel C里面去
func (dq *DelayQueue) Poll(exitC chan struct{}, nowF func() int64) {
	for {
		now := nowF()

		dq.mu.Lock()
		item, delta := dq.pq.PeekAndShift(now)
		if item == nil {
			//没有元素 或 元素处理中
			//我们必须确保整个操作的原子性
			atomic.StoreInt32(&dq.sleeping, 1)
		}
		dq.mu.Unlock()

		if item == nil {
			//该延迟优先级队列中没有元素插入
			if delta == 0 {

				select {
				//等待元素添加
				case <-dq.wakeupC:
					continue
				case <-exitC:
					goto exit
				}

				//至少有一个元素正在处理
			} else if delta > 0 {

				select {
				//新元素比当前队列中最早过期的元素更早
				case <-dq.wakeupC:
					continue
				//在delta时间(milliseconds)后
				case <-time.After(time.Duration(delta) * time.Millisecond):
					if atomic.SwapInt32(&dq.sleeping, 0) == 0 {
						<-dq.wakeupC
					}
					continue
				case <-exitC:
					goto exit
				}
			}
		}

		select {
		case dq.C <- item.Value:
		case <-exitC:
			goto exit
		}
	}

exit:
	//重置睡眠状态为0
	atomic.StoreInt32(&dq.sleeping, 0)
}

type priorityQueue []*item
type item struct {
	Value    interface{}
	Priority int64 //根据过期时间排列的优先级
	Index    int
}

func newPriorityQueue(wheelSize int) priorityQueue {
	return make(priorityQueue, 0, wheelSize)
}

func (pq priorityQueue) Len() int {
	return len(pq)
}
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

//把延迟任务推入优先级队列
func (pq *priorityQueue) Push(x interface{}) {

	n := len(*pq)
	c := cap(*pq)
	if n+1 > c {
		npq := make(priorityQueue, n, c*2)
		copy(npq, *pq)
		*pq = npq
	}

	*pq = (*pq)[0 : n+1]
	item := x.(*item)
	item.Index = n
	(*pq)[n] = item
}

//移除并返回移除的元素
func (pq *priorityQueue) Pop() interface{} {
	n := len(*pq)
	c := cap(*pq)
	if n < (c/2) && c > 25 {
		npq := make(priorityQueue, n, c/2)
		copy(npq, *pq)
		*pq = npq
	}
	item := (*pq)[n-1]
	item.Index = -1
	*pq = (*pq)[0 : n-1]
	return item
}

func (pq *priorityQueue) PeekAndShift(max int64) (*item, int64) {
	if pq.Len() == 0 {
		return nil, 0
	}

	//如果第一个元素 > max 说明没有过期,因此不移除
	item := (*pq)[0]
	if item.Priority > max {
		return nil, item.Priority - max
	}

	heap.Remove(pq, 0)
	return item, 0
}
