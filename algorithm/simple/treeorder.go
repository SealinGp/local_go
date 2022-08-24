package simple

import "container/list"

//二叉树专题
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//队列-链表实现--------------------------------
type queue struct {
	l *list.List
}

func NewQueue() *queue {
	return &queue{
		list.New(),
	}
}

func (this *queue) InQueue(val interface{}) {
	this.l.PushBack(val)
}

func (this *queue) OutQueue() interface{} {	
	ele := this.l.Front()
	if ele == nil {
		return nil
	}	

	this.l.Remove(ele)
	return ele.Value
}

func (this *queue) Len() int {
	return this.l.Len()
}

//栈-数组实现-----------------------------------
type stack struct {
	arr []interface{}
}

func NewStack() *stack {
	return &stack{arr: make([]interface{}, 0)}
}

func (this *stack) Push(val interface{}) {
	this.arr = append(this.arr, val)
}

func (this *stack) Pop() interface{} {
	if len(this.arr) > 0 {
		val := this.arr[len(this.arr)-1]
		this.arr = this.arr[:len(this.arr)-1]
		return val
	}
	return nil
}

func (this *stack) Len() int {
	return len(this.arr)
}

//深度优先遍历-前序-栈实现
func preOrderStack(root *TreeNode, f func(node *TreeNode)) {
	sta := NewStack()
	curr := root
	for curr != nil || sta.Len() > 0 {
		for curr != nil {
			f(curr)
			sta.Push(curr)
			curr = curr.Left
		}

		for sta.Len() > 0 {
			curr = sta.Pop().(*TreeNode)
			curr = curr.Right
		}
	}
}

//深度优先遍历-前序-递归
func preOrder(current *TreeNode, f func(curNode *TreeNode)) {
	if current != nil {
		f(current)
		preOrder(current.Left, f)
		preOrder(current.Right, f)
	}
}

func preOrderReverse(root *TreeNode, f func(curNode *TreeNode)) {
	if root == nil {
		return
	}

	f(root)
	preOrderReverse(root.Right, f)
	preOrderReverse(root.Left, f)
}

//深度优先遍历-中序-递归
func inOrder(current *TreeNode, f func(curNode *TreeNode)) {
	if current != nil {
		inOrder(current.Left, f)
		f(current)
		inOrder(current.Right, f)
	}
}

//深度优先遍历-后序-递归
func postOrder(current *TreeNode, f func(curNode *TreeNode)) {
	if current != nil {
		postOrder(current.Left, f)
		postOrder(current.Right, f)
		f(current)
	}
}

//广度优先遍历-队列实现
func levelOrder(root *TreeNode, f func(node *TreeNode)) {
	//根节点入队
	qu := NewQueue()
	qu.InQueue(root)

	for {
		first := qu.OutQueue()
		if first == nil {
			break
		}
		current := first.(*TreeNode)
		f(current)
		if current.Left != nil {
			qu.InQueue(current.Left)
		}
		if current.Right != nil {
			qu.InQueue(current.Right)
		}
	}
}

func symm(n1, n2 *TreeNode, f func(cn1, cn2 *TreeNode)) {
	f(n1, n2)
	if n1 != nil && n2 != nil {
		symm(n1.Left, n2.Right, f)
		symm(n1.Right, n2.Left, f)
	}
}
