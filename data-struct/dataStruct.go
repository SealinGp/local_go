package main

import (
	"container/list"
	"fmt"
	"log"
)

var structFuncs = map[string]func(){
	"linkTable": linkTable,
	"stack":     stack,
	"queue":     queue,
	"hashTable": hashTable,
	"dsu":       dsu,
	"Heap":      Heap,
	"ds1":       ds1,
	"ds2":       ds2,
}

// 栈-数组实现-----------------------------------
type stack1 struct {
	arr []interface{}
}

func NewStack() *stack1 {
	return &stack1{arr: make([]interface{}, 0)}
}
func (this *stack1) Push(val interface{}) {
	this.arr = append(this.arr, val)
}
func (this *stack1) Pop() interface{} {
	if len(this.arr) > 0 {
		val := this.arr[len(this.arr)-1]
		this.arr = this.arr[:len(this.arr)-1]
		return val
	}
	return nil
}
func (this *stack1) Len() int {
	return len(this.arr)
}

// 单向链表
type Node struct {
	Data int   //数据域
	Next *Node //指针域
}

// 双向链表
type DNode struct {
	Data  int
	Left  *DNode
	Right *DNode
}

// 链表 https://oi-wiki.org/ds/linked-list/
func linkTable() {
	//单向链表插入数据
	SNode := &Node{
		Data: 1,
		Next: &Node{
			Data: 2,
			Next: nil,
		},
	}

	//插入到索引为1 的位置,则找到索引为1-1=0的位置做插入操作
	i1 := 1
	iNode := &Node{
		Data: 3,
		Next: nil,
	}

	//插入链表
	i := 0
	for preN := SNode; preN.Next != nil; preN = preN.Next {
		if i == i1-1 {
			iNode.Next, preN.Next = preN.Next, iNode
			break
		}
		i++
	}

	for {
		log.Println(SNode.Data)
		if SNode.Next == nil {
			break
		}
		SNode = SNode.Next
	}

}

// 栈: https://oi-wiki.org/ds/stack/ 规律: 坐电梯
func stack() {
	//sta := []int{}
	//sta  = append(sta,1,2,3)
	//for i := len(sta)-1;i >= 0; i-- {
	//	log.Println(sta[i])
	//}
	sta1 := []int{}
	sta1 = append(sta1, 1, 2)
	fmt.Println(sta1[len(sta1)-1], sta1)

	sta1 = sta1[:len(sta1)-1]
	fmt.Println(sta1[len(sta1)-1], sta1)

	sta1 = sta1[:len(sta1)-1]
	fmt.Println(sta1)
}

// 队列: https://oi-wiki.org/ds/queue/ 规律: 排队
func queue() {
	que := []int{}
	que = append(que, 1, 2, 3)
	for _, v := range que {
		log.Println(v)
	}

	//双端队列:可以在队首|队尾 插入|删除 的队列

	//队首插入
	que = append([]int{0}, que...)
	log.Println(que)

	//队首删除
	que = que[1:]
	log.Println(que)

	//队尾插入
	que = append(que, 4)
	log.Println(que)

	//队尾删除
	que = que[:len(que)-1]
	log.Println(que)
}

// https://oi-wiki.org/ds/hash/
func hashTable() {
	m := make(map[string]interface{})
	m["a"] = "123"
	log.Println(m)
}

// https://oi-wiki.org/ds/dsu/ 并合集,用来查找环
// https://www.bilibili.com/video/av38498175/
var set = make([]int, 5)

func dsu() {
	makeSet(2)
	makeSet(3)
	Union(2, 3)
	log.Println(set)

	//makeSet(4)
	//Union(4,3)
	//log.Println(set)
	//
	//log.Println(findSet(3),findSet(2),"?")
}

// 新增一个集合,根节点的值为-1
func makeSet(i int) {
	set[i] = i
}

// 查找集合的根节点
func findSet(i int) int {
	if set[i] == i {
		return i
	}
	return findSet(set[i])
}

// 合并集合
func Union(i, j int) {
	i = findSet(i)
	j = findSet(j)
	if i > j {
		set[j] = i
	} else {
		set[i] = j
	}
}

// 深度优先遍历
// 前序遍历:  根节点 ——> 左子树 ——> 右子树
// 中序遍历:  左子树 ——> 根节点 ——> 右子树
// 后序遍历:  左子树 ——> 右子树 ——> 根节点

// 大根二叉树(从上->下,从右->左,从大->小) 排序二叉树
type Leaf struct {
	Key   int   //节点的大小 ?
	Value int   //节点值
	Left  *Leaf //左节点
	Right *Leaf //右节点
}
type tree interface {
	Insert(key int, value int)           //插入节点
	Min() int                            //值最小的节点
	Max() int                            //值最大的节点
	Search(key int) bool                 //查询是否存在
	InOrderTraverse(f func(value int))   //中序优先遍历
	PreOrderTraverse(f func(value int))  //前序优先遍历
	PostOrderTraverse(f func(value int)) //后序优先遍历
	String()                             //打印树
}

type T1 struct {
	Root *Leaf
}

func (t *T1) Insert(key, value int) {
	n := &Leaf{key, value, nil, nil}

	if t.Root == nil {
		t.Root = n
	} else {
		insertNode(t.Root, n)
	}
}
func insertNode(currentLeaf, newLeaf *Leaf) {
	if newLeaf.Key < currentLeaf.Key {
		if currentLeaf.Left == nil {
			currentLeaf.Left = newLeaf
		} else {
			insertNode(currentLeaf.Left, newLeaf)
		}
	} else {
		if currentLeaf.Right == nil {
			currentLeaf.Right = newLeaf
		} else {
			insertNode(currentLeaf.Right, newLeaf)
		}
	}
}

// 最小的节点在二叉树的最左边
func (t *T1) Min() int {
	n := t.Root
	if n == nil {
		return 0
	}
	for {
		if n.Left == nil {
			return n.Value
		}
		n = n.Left
	}
	return 0
}

func (t *T1) Max() int {
	n := t.Root
	if n == nil {
		return 0
	}
	for {
		if n.Right == nil {
			return n.Value
		}
		n = n.Right
	}
	return 0
}

func (t *T1) Search(key int) bool {
	return search(t.Root, key)
}
func search(currentLeaf *Leaf, key int) bool {
	if currentLeaf == nil {
		return false
	}
	if key < currentLeaf.Key {
		return search(currentLeaf.Left, key)
	}
	if key > currentLeaf.Key {
		return search(currentLeaf.Right, key)
	}
	return true
}

// 深度优先遍历-中序-递归实现
func (t *T1) InOrderTraverse(f func(value int)) {
	inOrderTraver(t.Root, f)
}
func inOrderTraver(currentLeaf *Leaf, f func(value int)) {
	if currentLeaf != nil {
		inOrderTraver(currentLeaf.Left, f)
		f(currentLeaf.Value)
		inOrderTraver(currentLeaf.Right, f)
	}
}

// 深度优先遍历-前序-递归实现
func (t *T1) PreOrderTraverse(f func(value int)) {
	preOrderTraver(t.Root, f)
}
func preOrderTraver(currentLeaf *Leaf, f func(value int)) {
	if currentLeaf != nil {
		f(currentLeaf.Value)
		preOrderTraver(currentLeaf.Left, f)
		preOrderTraver(currentLeaf.Right, f)
	}
}

// 深度优先遍历-前序-栈实现
func (t *T1) PreOrderTraverseStack(f func(value int)) {
	stack := NewStack()
	node := t.Root
	for node != nil || stack.Len() > 0 {
		for node != nil {
			f(node.Value)
			stack.Push(node)
			node = node.Left
		}

		if stack.Len() > 0 {
			node = stack.Pop().(*Leaf)
			node = node.Right
		}
	}
}

// 深度优先遍历-后序-递归实现
func (t *T1) PostOrderTraverse(f func(value int)) {
	postOrderTraver(t.Root, f)
}
func postOrderTraver(currentLeaf *Leaf, f func(value int)) {
	if currentLeaf != nil {
		postOrderTraver(currentLeaf.Left, f)
		postOrderTraver(currentLeaf.Right, f)
		f(currentLeaf.Value)
	}
}

// 广度优先遍历-队列实现
func (t *T1) LevelOrderTraversal(f func(value int)) {
	l := list.New()
	l.PushBack(t.Root)
	for l.Front() != nil {
		leaf := l.Front().Value.(*Leaf)
		f(leaf.Value)
		l.Remove(l.Front())
		if leaf.Left != nil {
			l.PushBack(leaf.Left)
		}
		if leaf.Right != nil {
			l.PushBack(leaf.Right)
		}
	}
}

// 广度优先遍历-递归实现 ?
func (t *T1) LevelOrderTraversal1(f func(value int)) {
	levelOrderTraversal1(t.Root, f)
}
func levelOrderTraversal1(current *Leaf, f func(value int)) {

}

func (t *T1) String() {
	fmt.Println(" ---------------------------------------------------- ")
	stringify(t.Root, 0)
	fmt.Println(" ---------------------------------------------------- ")
}
func stringify(currentLeaf *Leaf, level int) {

	if currentLeaf != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "              "
		}

		format += "----["
		level++
		fmt.Println(level)
		stringify(currentLeaf.Left, level)
		fmt.Printf(format+"%d\n", currentLeaf.Key)
		stringify(currentLeaf.Right, level)
	}
}

func ds1() {
	t1 := T1{
		Root: &Leaf{
			Key:   0,
			Value: 1,
			Left:  nil,
			Right: nil,
		},
	}
	t1.Insert(1, 2)
	t1.Insert(2, 3)
	t1.Insert(3, 4)
	t1.Insert(4, 5)
	t1.Insert(6, 6)
	//stringify(t1.Root,1)
	t1.PreOrderTraverseStack(func(value int) {
		fmt.Println(value)
	})
	fmt.Println("-------------")
	t1.PreOrderTraverse(func(value int) {
		fmt.Println(value)
	})
}

func ds2() {
	t1 := T1{}
	t1.Insert(8, 8)
	t1.Insert(3, 3)
	t1.Insert(1, 1)
	t1.Insert(10, 10)
	t1.Insert(6, 6)
	t1.Insert(4, 4)
	t1.Insert(7, 7)
	t1.Insert(13, 13)
	t1.Insert(14, 14)

	t1.LevelOrderTraversal(func(value int) {
		fmt.Println(value)
	})

	fmt.Println("-------------")

	t1.LevelOrderTraversal1(func(value int) {
		fmt.Println(value)
	})
}

// 二叉堆 = 最大堆(顶点>左>右) + 最小堆(顶点<左<右) = 基于数组结构实现
// 下面以小顶堆为例 讲述堆的弹出,新增节点
type heapType struct {
	arr []int
}

func Heap() {
	arr := []int{1, 3, 2, 6, 5, 7, 8, 9, 10}
	h := NewHeap(arr)
	h.Insert(0)
	fmt.Println(h.arr)

	arr1 := []int{1, 3, 2, 6, 5, 7, 8, 9, 10}
	h1 := NewHeap(arr1)
	fmt.Println(h1.Pop(), h1.arr) //1,[2,3,7,6,5,10,8,9]
}
func NewHeap(ar []int) heapType {
	ht := heapType{
		arr: ar,
	}
	for i := (len(ar) - 2) / 2; i >= 0; i-- {
		ht.downAdjust(i)
	}
	return ht
}

// 3
//2 1

//[3,2,1]
//

// 插入底部,向上调整
func (h *heapType) upAdjust() {
	childIndex := len(h.arr) - 1        //最后一个节点索引
	parentIndex := (childIndex - 1) / 2 //左节点 = 2*parent+1 ,右节点 = 2*parent+2 => parent = (子节点-1)/2
	tmp := h.arr[childIndex]

	for childIndex > 0 && tmp < h.arr[parentIndex] {
		//交换
		h.arr[childIndex] = h.arr[parentIndex]
		childIndex = parentIndex
		parentIndex = (childIndex - 1) / 2
	}
	h.arr[childIndex] = tmp
}

// 移除节点,把需要删除的节点跟顶部交换删除,然后把最后一个节点移动到顶部,然后做向下调整
func (h *heapType) downAdjust(parentIndex int) {
	childIndex := 2*parentIndex + 1
	len1 := len(h.arr)
	tmp := h.arr[parentIndex]

	for childIndex < len1 {
		//假设有右子节点并且右子节点比左子节点更小,则子节点为右子节点
		if childIndex+1 < len1 && h.arr[childIndex+1] < h.arr[childIndex] {
			childIndex = childIndex + 1
		}

		//假设父节点的值 < 子节点,则跳出
		if tmp <= h.arr[childIndex] {
			break
		}

		h.arr[parentIndex] = h.arr[childIndex]
		parentIndex = childIndex
		childIndex = 2*childIndex + 1
	}

	h.arr[parentIndex] = tmp
}
func (h *heapType) Insert(ele int) {
	h.arr = append(h.arr, ele)
	h.upAdjust()
}
func (h *heapType) Pop() int {
	//拿出堆顶节点
	last := h.arr[0]

	//把最后一个节点移动到堆顶节点
	h.arr[0] = h.arr[len(h.arr)-1]

	//堆长度调整
	h.arr = h.arr[:len(h.arr)-1]

	//堆顶节点向下调整
	h.downAdjust(0)
	return last
}

// 最小堆
type miniHeap struct {
	arr []int
}

func NewMiniHeap(a []int) *miniHeap {
	mh := &miniHeap{
		arr: a,
	}

	for i := (len(a) - 2) / 2; i >= 0; i-- {
		mh.downAdjust(i)
	}
	return mh
}
func (mh *miniHeap) Pop() int {
	if len(mh.arr) < 1 {
		return -1
	}
	tmp := mh.arr[0]
	mh.arr[0] = mh.arr[len(mh.arr)-1]
	mh.arr = mh.arr[:len(mh.arr)-1]
	mh.downAdjust(0)
	return tmp
}
func (mh *miniHeap) Insert(val int) {
	mh.arr = append(mh.arr, val)
	mh.upAdjust(len(mh.arr) - 1)
}
func (mh *miniHeap) upAdjust(index int) {
	tmp := mh.arr[index]
	parentIndex := (index - 1) / 2

	for index > 0 && tmp < mh.arr[parentIndex] {
		mh.arr[index] = mh.arr[parentIndex]
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
	mh.arr[index] = tmp
}
func (mh *miniHeap) downAdjust(index int) {
	arrLen := len(mh.arr)
	if index > arrLen-1 {
		return
	}
	tmp := mh.arr[index]
	leftChild := 2*index + 1
	rightChild := leftChild + 1
	for leftChild < arrLen {
		if rightChild < arrLen-1 && mh.arr[rightChild] < mh.arr[leftChild] {
			leftChild = rightChild
		}

		if tmp <= mh.arr[leftChild] {
			break
		}

		mh.arr[index] = mh.arr[leftChild]
		index = leftChild
		leftChild = 2*index + 1
		rightChild = leftChild + 1
	}

	mh.arr[index] = tmp
}
