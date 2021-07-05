package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

//简单时间轮实现
//http://russellluo.com/2018/10/golang-implementation-of-hierarchical-timing-wheels.html

func main() {
	log.Println("start")

	tw := NewTimeWheel()

	tw.AddNode(time.Second*2, func() {
		log.Println("now")
	})

	go func() {
		tw.Start()
	}()

	<-time.After(time.Second * 5)
	tw.Stop()
}

type timeWheel struct {
	stopCh       chan bool
	slots        []slot        //所有的槽
	slotLen      int           //多少个槽
	slotInterval time.Duration //转动一个槽需要的时间  默认 1 s/个槽
	curSlot      int           //当前指向第几个槽
}
type slot struct {
	head *node
	sLen int
}
type node struct {
	round int //转多少轮
	task  func()
	prev  *node
	next  *node
}

func NewTimeWheel() *timeWheel {
	return &timeWheel{
		stopCh:       make(chan bool),
		slotInterval: 1 * time.Second,
		slots:        make([]slot, 60),
		slotLen:      60,
		curSlot:      0,
	}
}

func (tw *timeWheel) AddNode(dur time.Duration, task func()) {
	//不足1s按1s算
	if dur < tw.slotInterval {
		dur = tw.slotInterval
	}

	//1/60
	slotIndex := int(math.Ceil(dur.Seconds())) % (int(tw.slotInterval.Seconds()) * tw.slotLen)
	newNode := &node{
		round: int(math.Ceil(dur.Seconds())) / len(tw.slots),
		task:  task,
		prev:  nil,
		next:  tw.slots[slotIndex].head,
	}
	tw.slots[slotIndex].head = newNode
}

func (tw *timeWheel) Start() {
	for {
		select {
		case <-tw.stopCh:
			fmt.Println("stopped")
			return
		default:
			tw.checkSlot()
			time.Sleep(tw.slotInterval)
		}
	}
}
func (tw *timeWheel) Stop() {
	tw.stopCh <- true
}

func (tw *timeWheel) checkSlot() {
	tw.curSlot = tw.curSlot % 60
	fmt.Println("current slot:", tw.curSlot)
	curSlot := tw.slots[tw.curSlot]
	curNode := curSlot.head
	for curNode != nil {
		curNode.round--
		//任务到期
		if curNode.round < 0 {

			//1.执行任务
			curNode.task()

			//2.删除该任务节点
			//当前槽的链表节点不是首节点的时候
			if curNode.prev != nil {
				curNode.prev.next = curNode.next
				//当前槽的链表节点是首节点的时候
			} else {
				tw.slots[tw.curSlot].head = nil
			}
			tw.slots[tw.curSlot].sLen--
		}
		curNode = curNode.next
	}
	tw.curSlot++
}
