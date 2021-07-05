package simple

import "fmt"

//https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
func (*Ref) KEIAR() {
	nums := []int{2, 1}
	k := 2

	h := NewBigHeap(nums)
	kValue := 0
	for i := 0; i < k; i++ {
		kValue = h.Pop()
	}

	fmt.Println(kValue)
}

type bigHeap struct {
	arr []int
}

func NewBigHeap(nums []int) *bigHeap {
	h := &bigHeap{arr: nums}
	for i := (len(h.arr) - 2) / 2; i >= 0; i-- {
		h.downAdjust(i)
	}
	return h
}
func (this *bigHeap) downAdjust(parentIndex int) {
	if parentIndex > len(this.arr)-1 {
		return
	}
	childIndex := 2*parentIndex + 1
	heapL := len(this.arr)
	tmp := this.arr[parentIndex]

	for childIndex < heapL {
		//选出子节点中最小的
		if childIndex+1 < heapL && this.arr[childIndex+1] > this.arr[childIndex] {
			childIndex = childIndex + 1
		}

		if tmp >= this.arr[childIndex] {
			break
		}

		this.arr[parentIndex] = this.arr[childIndex]
		parentIndex = childIndex
		childIndex = 2*parentIndex + 1
	}

	this.arr[parentIndex] = tmp
}
func (this *bigHeap) Pop() int {
	if len(this.arr) < 1 {
		return -1
	}

	last := this.arr[0]
	this.arr[0] = this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	this.downAdjust(0)

	return last
}
