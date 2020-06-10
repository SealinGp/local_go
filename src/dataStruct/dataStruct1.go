package main

import (
	"fmt"
	"log"
	"os"
)
func main() {
	if len(os.Args) <= 1 {
		log.Fatal("func required")
	}

	fun := map[string]func(){
		"ds3" : ds3,
	}
	fun[os.Args[1]]()
}

//优先级队列的实现 = 二叉堆 + 队列 (逻辑结构) = 数组 (物理结构)
//这里大顶堆为例
func ds3()  {
	ds3 := priorityQueue{}
	ds3.EnQueue(3)
	ds3.EnQueue(5)
	ds3.EnQueue(10)
	ds3.EnQueue(2)
	ds3.EnQueue(7)

	fmt.Println(ds3.DeQueue())
	fmt.Println(ds3.DeQueue())
}
type priorityQueue []int
func (pq *priorityQueue)EnQueue(ele int)  {
	*pq = append(*pq,ele)
	pq.upAdjust()
}
func (pq *priorityQueue)upAdjust()  {
	childIndex  := len(*pq)-1
	parentIndex := (childIndex - 1)/2
	tmp         := (*pq)[childIndex]

	for childIndex > 0  && (*pq)[parentIndex] < tmp {
		(*pq)[childIndex] = (*pq)[parentIndex]
		childIndex        = parentIndex
		parentIndex       = parentIndex/2
	}
	(*pq)[childIndex] = tmp
}
func (pq *priorityQueue)DeQueue() int {
	first   := (*pq)[0]
	(*pq)[0] = (*pq)[len(*pq)-1]
	*pq      =  (*pq)[:len(*pq)-1]
	pq.downAdjust()
	return first
}
func (pq *priorityQueue)downAdjust()  {
	parentIndex := 0
	tmp         := (*pq)[parentIndex]
	childIndex  :=  2*parentIndex + 1
	Len         := len(*pq)

	for childIndex < Len {
		if childIndex + 1 < Len && (*pq)[childIndex+1] > (*pq)[childIndex] {
			childIndex = childIndex+1
		}

		if tmp > (*pq)[childIndex] {
			break
		}

		(*pq)[parentIndex] = (*pq)[childIndex]
		parentIndex        = childIndex
		childIndex         = 2*childIndex + 1
	}

	(*pq)[parentIndex] = tmp
}