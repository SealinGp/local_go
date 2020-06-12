package main

import (
	"container/list"
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
		"ds4" : ds4,
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

func ds4()  {
	a := []int{4,7,6,5,3,2,8,1}
	FastSortDouble(a,0,len(a)-1)
	fmt.Println(a)

	b := []int{4,7,3,5,6,2,8,1}
	FastSortSingle(b,0,len(b)-1)
	fmt.Println(b)

	c := []int{4,7,3,5,6,2,8,1}
	FastSortStack(c,0,len(c)-1)
	fmt.Println(c)
}
//快排(小->大)的双边循环交换法-递归实现-左右两个指针,左找比基准元素大的,右找币基准元素小的,然后交换直到左右指针重合,然后吧基准元素跟左指针对应的元素交换
func FastSortDouble(arr []int,startIndex,endIndex int)  {
	if startIndex >= endIndex {
		return
	}

	privotIndex := partition(arr,startIndex,endIndex)
	FastSortDouble(arr,startIndex,privotIndex - 1)
	FastSortDouble(arr,privotIndex+1,endIndex)
}
func partition(arr []int,startIndex,endIndex int) int {
	privot := arr[startIndex]
	left   := startIndex
	right  := endIndex
	for right != left {
		//右指针左移
		for left < right && arr[right] > privot  {
			right--
		}
		//左指针右移
		for left < right && arr[left] <= privot  {
			left++
		}
		if left < right {
			arr[left],arr[right] = arr[right],arr[left]
		}
	}

	//privot和指针重合点交换
	arr[startIndex],arr[left] = arr[left],privot
	return left
}

//快排(小->大)的单边循环法-递归实现
func FastSortSingle(arr []int,startIndex,endIndex int)  {
	if startIndex >= endIndex {
		return
	}
	//mark是边界
	privotIndex := partition1(arr,startIndex,endIndex)
	FastSortSingle(arr,startIndex,privotIndex-1)
	FastSortSingle(arr,privotIndex+1,endIndex)
}
func partition1(arr []int,startIndex,endIndex int) int {
	privot := arr[startIndex]
	mark   := startIndex

	for i := startIndex + 1; i <= endIndex; i++ {
		if arr[i] < privot {
			mark++
			arr[mark],arr[i] = arr[i],arr[mark]
		}
	}

	arr[startIndex],arr[mark] = arr[mark],privot
	return mark
}

//快排(小->大)的栈实现
func FastSortStack(arr []int, startIndex,endIndex int)  {
	stack           := list.New()
	ma              := make(map[string]int)
	ma["startIndex"] = startIndex
	ma["endIndex"]   = endIndex

	stack.PushBack(ma)
	for stack.Len() != 0  {
		ma1         := stack.Remove(stack.Back())
		param       := ma1.(map[string]int)
		privotIndex := partition1(arr,param["startIndex"],param["endIndex"])
		if param["startIndex"] < privotIndex - 1 {
			leftParam := map[string]int{
				"startIndex" : param["startIndex"],
				"endIndex"   : privotIndex - 1,
			}
			stack.PushBack(leftParam)
		}
		if privotIndex + 1 < param["endIndex"] {
			rightParam := map[string]int{
				"startIndex" : privotIndex + 1,
				"endIndex"   : endIndex,
			}
			stack.PushBack(rightParam)
		}
	}
}