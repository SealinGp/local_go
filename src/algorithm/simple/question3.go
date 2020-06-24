package simple

import (
	"fmt"
)

//二分查找
func (*Ref)TwoSearch()  {
	nums  := []int{
		1,2,3,4,5,
	}
	target := 5
	fmt.Println(search(nums,target))
}
func search(nums []int,target int) int {
	start := 0
	end   :=  len(nums)-1
	for start + 1 < end  {
		mid := start + (end - start)/2
		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			start = mid
		} else if target < nums[mid] {
			end   = mid
		}
	}

	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}

//https://leetcode-cn.com/problems/smallest-k-lcci/submissions/
//采用小顶堆实现
func (*Ref)SK(arr []int, k int) []int {
	mh   := NewMinHeap(arr)
	arr1 := []int{}

	for k > 0  {
		arr1 = append(arr1,mh.Pop())

		k--
	}
	return arr1
}
//小顶堆
type minHeap struct {
	h []int
}
func NewMinHeap(arr []int) *minHeap {
	mh := &minHeap{h:arr}
	for i := (len(arr)-2)/2; i >=0 ; i-- {
		mh.downAdjust(i)
	}
	return mh
}
func (this *minHeap)downAdjust(parentIndex int)  {
	tmp        := this.h[parentIndex]
	childIndex := 2*parentIndex + 1
	Hlen       := len(this.h)

	for childIndex < Hlen {
		if childIndex + 1 < Hlen  && this.h[childIndex + 1] < this.h[childIndex] {
			childIndex++
		}
		if tmp < this.h[childIndex] {
			break
		}
		this.h[parentIndex] = this.h[childIndex]
		parentIndex = childIndex
		childIndex  = 2 * parentIndex + 1
	}
	this.h[parentIndex] = tmp
}
func (this *minHeap)upAdjust()  {
	childIndex  := len(this.h)-1
	parentIndex := (childIndex - 1) / 2
	tmp         := this.h[childIndex]

	for childIndex > 0 && tmp < this.h[parentIndex] {
		this.h[childIndex] = this.h[parentIndex]
		childIndex = parentIndex
		parentIndex = (childIndex - 1)  / 2
	}
	this.h[childIndex] = tmp
}
func (this *minHeap)Insert(a int)  {
	this.h = append(this.h,a)
	this.upAdjust()
}
func (this *minHeap)Pop() int {
	first := this.h[0]
	this.h[0] = this.h[len(this.h) - 1]
	this.h = this.h[:len(this.h)-1]
	this.downAdjust(0)
	return first
}