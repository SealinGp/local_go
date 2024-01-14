package main

type MyCircularQueue struct {
	font int
	rear int
	arr  []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		arr: make([]int, k+1),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.arr[this.rear] = value
	this.rear = (this.rear + 1) % len(this.arr)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.font = (this.font + 1) % len(this.arr)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.arr[this.font]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	idx := (this.rear - 1 + len(this.arr)) % len(this.arr)
	return this.arr[idx]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.rear == this.font
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.rear+1)%len(this.arr) == this.font
}

//k=5
//1%5 = 1
//2%5 = 2
//3%5 = 3
//4%5 = 4

//5%5 = 0

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
