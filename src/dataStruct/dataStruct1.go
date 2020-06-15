package main

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"sort"
)
func main() {
	if len(os.Args) <= 1 {
		log.Fatal("func required")
	}

	fun := map[string]func(){
		"ds3" : ds3,
		"ds4" : ds4,
		"ds5" : ds5,
		"ds6" : ds6,
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

//计数排序(稳定排序=基于计数排序的优化版本)-适用场景:1.已知数组范围 2.最大值和最小值间隔不大 3.整数数组
//n = 数组长度, m = 数组最大值和最小值的间隔,时间复杂度O(n) = 3n+m = n + m 空间复杂度O(n) = m
func CountSort(arr []int) []int {
	//找出数组最大值
	max := arr[0]
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	//确定数组长度
	countArr := make([]int,max - min + 1)
	for i := range arr {
		countArr[arr[i] - min]++
	}

	//变形累加前面的值
	for i := range countArr  {
		if i > 0 {
			countArr[i] = countArr[i] + countArr[i-1]
		}
	}

	//倒序遍历原数组,
	sortArr := make([]int,len(arr))
	for i := len(arr)-1; i >= 0 ; i-- {
		//找到该值的索引(索引-1=表示当前值在排序后的数组中的索引)
		index := countArr[arr[i] - min]
		index--
		countArr[arr[i] - min] = index

		sortArr[index]         = arr[i]
	}

	return sortArr
}
func ds5()  {
	arr := []int{
		9,3,5,4,9,1,2,7,8,1,3,6,5,3,4,0,10,9,7,9,
	}
	arr = CountSort(arr)
	fmt.Println(arr)

	arr1 := []int{
		95,94,91,98,99,90,99,93,91,92,
	}
	fmt.Println(CountSort(arr1))

	arr2 := []float64{
		4.12,6.421,0.0023,3.0,2.123,8.122,4.12,10.09,
	}
	fmt.Println(BucketSort(arr2))
}

//桶排序
func BucketSort(arr []float64) []float64 {
	if len(arr) < 1 {
		return arr
	}

	//1.找出最大和最小值
	max := arr[0]
	min := arr[0]
	for i := range arr  {
		if i > 0 {
			if arr[i] > max {
				max = arr[i]
			}
			if arr[i] < min {
				min = arr[i]
			}
		}
	}


	//2.创建桶
	bucketLen := len(arr)
	buckets := make([][]float64,bucketLen)           //桶的数量 = 数组的长度
	//jg      := (max - min)/(float64(len(arr)) - 1)   //每个桶的区间间隔
	d       := max - min

	//3.将每个元素放到对应的范围的桶中
	for i := range arr  {
		index := ( (arr[i] - min) * (float64(bucketLen) - 1) ) / d
		buckets[int(index)] = append(buckets[int(index)],arr[i])
	}

	//4.对每个桶进行排序
	for i := range buckets  {
		sort.Float64s(buckets[i])
	}

	//5.输出全部元素
	sortedArr := make([]float64,len(arr))
	j := 0
	for i := range buckets {
		for i1 := range buckets[i] {
			sortedArr[j] = buckets[i][i1]
			j++
		}
	}
	return sortedArr
}
func ds6()  {

	fmt.Println(1^1)
	fmt.Println(0^1)
}