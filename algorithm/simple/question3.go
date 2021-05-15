package simple

import (
	"fmt"
)

//二分查找
func (*Ref) TwoSearch() {
	nums := []int{
		1, 2, 3, 4, 5,
	}
	target := 5
	fmt.Println(search(nums, target))
}
func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			start = mid
		} else if target < nums[mid] {
			end = mid
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
func (*Ref) SK(arr []int, k int) []int {
	mh := NewMinHeap(arr)
	arr1 := []int{}

	for k > 0 {
		arr1 = append(arr1, mh.Pop())

		k--
	}
	return arr1
}

//小顶堆
type minHeap struct {
	h []int
}

func NewMinHeap(arr []int) *minHeap {
	mh := &minHeap{h: arr}
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		mh.downAdjust(i)
	}
	return mh
}
func (this *minHeap) downAdjust(parentIndex int) {
	tmp := this.h[parentIndex]
	childIndex := 2*parentIndex + 1
	Hlen := len(this.h)

	for childIndex < Hlen {
		if childIndex+1 < Hlen && this.h[childIndex+1] < this.h[childIndex] {
			childIndex++
		}
		if tmp < this.h[childIndex] {
			break
		}
		this.h[parentIndex] = this.h[childIndex]
		parentIndex = childIndex
		childIndex = 2*parentIndex + 1
	}
	this.h[parentIndex] = tmp
}
func (this *minHeap) upAdjust() {
	childIndex := len(this.h) - 1
	parentIndex := (childIndex - 1) / 2
	tmp := this.h[childIndex]

	for childIndex > 0 && tmp < this.h[parentIndex] {
		this.h[childIndex] = this.h[parentIndex]
		childIndex = parentIndex
		parentIndex = (childIndex - 1) / 2
	}
	this.h[childIndex] = tmp
}
func (this *minHeap) Insert(a int) {
	this.h = append(this.h, a)
	this.upAdjust()
}
func (this *minHeap) Pop() int {
	first := this.h[0]
	this.h[0] = this.h[len(this.h)-1]
	this.h = this.h[:len(this.h)-1]
	this.downAdjust(0)
	return first
}

//https://leetcode-cn.com/problems/lru-cache/submissions/
type LRUCache struct {
	ma  map[int]*DNode
	dt  *DLinkTable
	cap int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		ma:  make(map[int]*DNode, capacity),
		dt:  &DLinkTable{},
	}
}
func (this *LRUCache) Get(key int) int {
	v, ok := this.ma[key]
	if !ok {
		return -1
	}

	this.dt.MoveToHead(v)
	return v.Val
}
func (this *LRUCache) Put(key int, value int) {
	_, ok := this.ma[key]
	if ok {
		this.ma[key].Val = value
		this.dt.MoveToHead(this.ma[key])
		return
	}

	if len(this.ma)+1 > this.cap {
		if this.dt.tail != nil {
			delete(this.ma, this.dt.tail.Key)
		}
		this.dt.DelLast()
	}
	newNode := &DNode{
		Key: key,
		Val: value,
	}
	this.ma[key] = newNode
	this.dt.AddToHead(newNode)
}

type DLinkTable struct {
	head *DNode
	tail *DNode
}
type DNode struct {
	Key  int
	Val  int
	Prev *DNode
	Next *DNode
}

func (this *DLinkTable) DelNode(node *DNode) {
	if node == nil {
		return
	}

	if node == this.head {
		this.head = node.Next
		if this.head != nil {
			this.head.Prev = nil
		}
		return
	}
	if node == this.tail {
		this.tail = node.Prev
		this.tail.Next = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
func (this *DLinkTable) DelLast() {
	this.DelNode(this.tail)
}
func (this *DLinkTable) MoveToHead(node *DNode) {
	if node == nil || this.head == node {
		return
	}

	this.DelNode(node)
	node.Next = this.head
	node.Prev = nil
	this.head.Prev = node
	this.head = node
}
func (this *DLinkTable) AddToHead(node *DNode) {
	if node == nil {
		return
	}
	if this.head == nil {
		this.head = node
	} else {
		node.Next = this.head
		node.Prev = nil
		this.head.Prev = node
		this.head = node
	}

	curN := this.head
	for curN != nil {
		if curN.Next == nil {
			break
		}
		curN = curN.Next
	}
	this.tail = curN
}

func (*Ref) Test1() {
	t := Constructor(1)
	t.Put(2, 1)
	t.Get(2)
	t.Put(3, 2)
	fmt.Println(t.Get(2))
}
