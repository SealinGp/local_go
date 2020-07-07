package simple

import (
	"container/list"
	"fmt"
)

//二叉树专题

//队列-链表实现--------------------------------
type queue struct {
	l *list.List
}
func NewQueue() *queue {
	return &queue{
		list.New(),
	}
}
func (this *queue)InQueue(val interface{}){
	this.l.PushBack(val)
}
func (this *queue)OutQueue() interface{} {
	ele := this.l.Front()
	if ele == nil {
		return nil
	}

	this.l.Remove(ele)
	return ele.Value
}
func (this *queue)Len() int {
	return this.l.Len()
}

//栈-数组实现-----------------------------------
type stack struct {
	arr []interface{}
}
func NewStack() *stack {
	return &stack{arr:make([]interface{},0)}
}
func (this *stack)Push(val interface{})  {
	this.arr = append(this.arr,val)
}
func (this *stack)Pop() interface{} {
	if len(this.arr) > 0 {
		val     := this.arr[len(this.arr)-1]
		this.arr = this.arr[:len(this.arr)-1]
		return val
	}
	return nil
}
func (this *stack)Len() int {
	return len(this.arr)
}


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//前序----递归
func (this *TreeNode)PreOrder(f func(curNode *TreeNode))  {
	preOrder(this,f)
}
func preOrder(current *TreeNode,f func(curNode *TreeNode))  {
	if current != nil {
		f(current)
		preOrder(current.Left,f)
		preOrder(current.Right,f)
	}
}
//前序----栈
func (this *TreeNode)PreOrderStack(f func(node *TreeNode))  {
	preOrderStack(this,f)
}
func preOrderStack(root *TreeNode,f func(node *TreeNode))  {
	sta  := NewStack()
	curr := root
	for curr != nil || sta.Len() > 0  {
		for curr != nil  {
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

//中序----
func (this *TreeNode)InOrder(f func(curNode *TreeNode))  {
	inOrder(this,f)
}
func inOrder(current *TreeNode,f func(curNode *TreeNode))  {
	if current != nil {
		inOrder(current.Left,f)
		f(current)
		inOrder(current.Right,f)
	}
}

//后序
func (this *TreeNode)PostOrder(f func(curNode *TreeNode))  {
	postOrder(this,f)
}
func postOrder(current *TreeNode,f func(curNode *TreeNode))  {
	if current != nil {
		postOrder(current.Left,f)
		postOrder(current.Right,f)
		f(current)
	}
}

//广度优先遍历-队列实现
func (this *TreeNode)LevelOrder(f func(node *TreeNode))  {
	//根节点入队
	qu := NewQueue()
	qu.InQueue(this)

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
func (this *TreeNode)DepthLevel() int {
	return depth1(this)
}
func (this *TreeNode)DepthPost() int {
	return depth2(this)
}
//广度优先遍历-实现获取树的深度
func depth1(root *TreeNode) int {
	//根节点入队
	depth := 0
	qu := NewQueue()
	qu.InQueue(root)
	depth++
	tmp := NewQueue()
	for {
		if qu.Len() == 0 && tmp.Len() == 0 {
			break
		}
		if qu.Len() == 0 && tmp.Len() > 0 {
			depth++
			qu  = tmp
			tmp = NewQueue()
		}
		first := qu.OutQueue()


		current := first.(*TreeNode)
		//f(current)
		if current.Left != nil {
			tmp.InQueue(current.Left)
		}
		if current.Right != nil {
			tmp.InQueue(current.Right)
		}
	}

	return depth
}
//深度优先遍历-后序-递归-实现获取树的深度
func depth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth2(root.Left),depth2(root.Right)) + 1
}

//https://leetcode-cn.com/problems/minimum-height-tree-lcci/
func (*Ref)STBST()  {
	nums     := []int{-10,-3,0,5,9}

	if len(nums) == 0 {
		return
	}
	node := &TreeNode{}
	c(node,nums)

	node.PreOrder(func(curNode *TreeNode) {
		fmt.Println(curNode.Val)
	})
}
func c(current *TreeNode,arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	rootNode := len(arr)/2

	current.Val   = arr[rootNode]
	if current.Left == nil {
		current.Left = &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		}
	}
	if current.Right == nil {
		current.Right = &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		}
	}
	current.Left  = c(current.Left,arr[:rootNode])
	current.Right = c(current.Right,arr[rootNode+1:])

	return current
}

//https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
func (*Ref)Mt()  {
	rootA     := []int{2,3,-1,1}
	root      := ArrToNode(rootA)
	toMirror(root)
	root.LevelOrder(func(node *TreeNode) {
		fmt.Println(node.Val)
	})
}
func toMirror(current *TreeNode) *TreeNode {
	if current == nil {
		return nil
	}
	tmp          := current.Left
	current.Left  = toMirror(current.Right)
	current.Right = toMirror(tmp)
	return current
}

//层级遍历输出的数组 转为 树
func ArrToNode(arr []int) *TreeNode {
	rootNode := &TreeNode{
		Val:   arr[0],
		Left:  nil,
		Right: nil,
	}
	currentNode := rootNode

	for i := 0; i < len(arr); i++ {
		left  := 2*i+1  //1,2|3,4
		right := left+1


		if left < len(arr) && arr[left] != -1 {
			currentNode.Left = &TreeNode{
				Val:   arr[left],
				Left:  nil,
				Right: nil,
			}
		}
		if right < len(arr) && arr[right] != -1 {
			currentNode.Right = &TreeNode{
				Val:   arr[right],
				Left:  nil,
				Right: nil,
			}
		}
		if left == len(arr) {
			break
		}
		if right == len(arr) {
			break
		}
		if i+1 > len(arr) {
			break
		}
		if  arr[i+1] != -1{
			preOrder(rootNode,func(curNode *TreeNode) {
				if curNode.Val == arr[i+1] {
					currentNode = curNode
				}
			})
		}
	}
	return rootNode
}

//https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
func (*Ref)Md()  {
	rootA := []int{3,9,20,-1,-1,15,7}
	root  := ArrToNode(rootA)
	fmt.Println(root.DepthLevel())
	fmt.Println(root.DepthPost())
}

func (*Ref)RsBst()  {
	rootA := []int{15,9,21,7,13,19,23,5,-1,11,-1,17}
	root  := ArrToNode(rootA)

	fmt.Println(sum(root,19,21,0))
}

func sum(root *TreeNode,L,R,ans int) int {
	if root != nil {
		if L <= root.Val && root.Val <= R {
			ans += root.Val
		}
		if L < root.Val {
			ans = sum(root.Left,L,R,ans)
		}
		if root.Val < R {
			ans = sum(root.Right,L,R,ans)
		}
	}
	return ans
}