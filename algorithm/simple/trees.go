package simple

import (
	"container/list"
	"fmt"
	"log"
	"math"
	"strconv"
)

func (*Ref) RebuildTree() {
	buildTree([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
}

/*
1.根据preorder找到root value
2.根据inorder找到leftArr,rightArr
3.
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{
		Val: rootVal,
	}

	rootIndex := 0
	for i, v := range inorder {
		if v == rootVal {
			rootIndex = i
			break
		}
	}

	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}

func (this *TreeNode) DepthLevel() int {
	return depth1(this)
}
func (this *TreeNode) DepthPost() int {
	return depth2(this)
}

func (*Ref) FindMinHeightTrees() {

}

func findMinHeightTrees() {

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
			qu = tmp
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
	return max(depth2(root.Left), depth2(root.Right)) + 1
}

//https://leetcode-cn.com/problems/minimum-height-tree-lcci/
func (*Ref) STBST() {
	nums := []int{-10, -3, 0, 5, 9}

	if len(nums) == 0 {
		return
	}
	node := &TreeNode{}
	c(node, nums)

	preOrder(node, func(curNode *TreeNode) {
		fmt.Println(curNode.Val)
	})
}
func c(current *TreeNode, arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	rootNode := len(arr) / 2

	current.Val = arr[rootNode]
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
	current.Left = c(current.Left, arr[:rootNode])
	current.Right = c(current.Right, arr[rootNode+1:])

	return current
}

//https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
func (*Ref) Mt() {
	rootA := []int{2, 3, -1, 1}
	root := ArrToNode(rootA)
	toMirror(root)
	levelOrder(root, func(node *TreeNode) {
		fmt.Println(node.Val)
	})
}
func toMirror(current *TreeNode) *TreeNode {
	if current == nil {
		return nil
	}
	tmp := current.Left
	current.Left = toMirror(current.Right)
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
		left := 2*i + 1 //1,2|3,4
		right := left + 1

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
		if arr[i+1] != -1 {
			preOrder(rootNode, func(curNode *TreeNode) {
				if curNode.Val == arr[i+1] {
					currentNode = curNode
				}
			})
		}
	}
	return rootNode
}
func PreArrToNode(arr []int) *TreeNode {
	rootList := list.New()
	for _, v := range arr {
		rootList.PushBack(v)
	}
	return preArrToNode(rootList)
}
func preArrToNode(list2 *list.List) *TreeNode {
	var node *TreeNode
	if list2.Len() <= 0 {
		return nil
	}
	ele := list2.Front()
	list2.Remove(ele)

	data := ele.Value.(int)
	if data >= 0 {
		node = &TreeNode{
			Val: data,
		}
		node.Left = preArrToNode(list2)
		node.Right = preArrToNode(list2)
	}
	return node
}

func (*Ref) Test() {
	arr := []int{3, 2, 9, -1, -1, 10, -1, -1, 8, -1, 4}
	node := PreArrToNode(arr)
	preOrder(node, func(curNode *TreeNode) {
		fmt.Println(curNode.Val)
	})
}

//https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
func (*Ref) Md() {
	rootA := []int{3, 9, 20, -1, -1, 15, 7}
	root := ArrToNode(rootA)
	fmt.Println(root.DepthLevel())
	fmt.Println(root.DepthPost())
}

func (*Ref) RsBst() {
	rootA := []int{15, 9, 21, 7, 13, 19, 23, 5, -1, 11, -1, 17}
	root := ArrToNode(rootA)

	fmt.Println(sum(root, 19, 21, 0))
}

func sum(root *TreeNode, L, R, ans int) int {
	if root != nil {
		if L <= root.Val && root.Val <= R {
			ans += root.Val
		}
		if L < root.Val {
			ans = sum(root.Left, L, R, ans)
		}
		if root.Val < R {
			ans = sum(root.Right, L, R, ans)
		}
	}
	return ans
}

// https://leetcode-cn.com/problems/merge-two-binary-trees/
func (*Ref) MergeTrees() {
	t1 := ArrToNode([]int{1, 3, 2, 5})
	t2 := ArrToNode([]int{2, 1, 3, -1, 4, -1, 7})

	t1 = preOrder1(t1, t2)
	levelOrder(t1, func(node *TreeNode) {
		fmt.Println(node.Val)
	})
}
func preOrder1(n1 *TreeNode, n2 *TreeNode) *TreeNode {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	n1.Val += n2.Val
	n1.Left = preOrder1(n1.Left, n2.Left)
	n1.Right = preOrder1(n1.Right, n2.Right)
	return n1
}

//https://leetcode-cn.com/problems/invert-binary-tree/
func (*Ref) InvertTree() {
	root := ArrToNode([]int{4, 2, 7, 1, 3, 6, 9})
	mirror(root)
	levelOrder(root, func(node *TreeNode) {
		fmt.Println(node.Val)
	})
}
func mirror(node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		return
	}

	tmp := node.Left
	node.Left = node.Right
	node.Right = tmp
	mirror(node.Left)
	mirror(node.Right)
}

//https://leetcode-cn.com/problems/search-in-a-binary-search-tree/
func (*Ref) SearchBST() {
	root := ArrToNode([]int{4, 2, 7, 1, 3})
	val := 2

	var valNode *TreeNode
	valNode = bstFindTree(root, val)

	if valNode != nil {
		levelOrder(valNode, func(node *TreeNode) {
			fmt.Println(node.Val)
		})
	}
}
func bstFindTree(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val == val {
		return node
	}

	if val > node.Val {
		return bstFindTree(node.Right, val)
	}
	return bstFindTree(node.Left, val)
}

type Node struct {
	Val      int
	Children []*Node
}

//https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
func (*Ref) PO() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{3, []*Node{
				{5, nil},
				{6, nil},
			}},
			{2, nil},
			{4, nil},
		},
	}

	arr := []int{}
	Npo1(root, func(cn *Node) {
		arr = append(arr, cn.Val)
	})
	sta2 := NewStack()
	for _, v := range arr {
		sta2.Push(v)
	}
	i := 0
	for sta2.Len() > 0 {
		arr[i] = sta2.Pop().(int)
		i++
	}
	fmt.Println(arr)
}

//递归实现
func Npo(n *Node, f func(cn *Node)) {
	if n == nil {
		return
	}

	for len(n.Children) > 0 {
		Npo(n.Children[0], f)
		n.Children = n.Children[1:]
	}
	f(n)
}

//栈实现
func Npo1(root *Node, f func(cn *Node)) {
	sta := NewStack()
	sta.Push(root)
	for sta.Len() > 0 {
		n := sta.Pop().(*Node)
		f(n)

		if n.Children != nil {
			for _, v := range n.Children {
				sta.Push(v)
			}
		}
	}
}

func (*Ref) KthLargest() {
	root := ArrToNode([]int{3, 1, 4, -1, 2})
	k := 4

	val := 0
	inOrder(root, func(curNode *TreeNode) {
		if k < 1 {
			return
		}
		val = curNode.Val
		k--
	})
	fmt.Println(val)
}

//https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
func (*Ref) Pre() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{3, []*Node{
				{5, nil},
				{6, nil},
			}},
			{2, nil},
			{4, nil},
		},
	}

	//递归
	arr := []int{}
	pre1(root, func(node *Node) {
		arr = append(arr, node.Val)
	})

	arr2 := []int{}
	pre2(root, func(node *Node) {
		arr2 = append(arr2, node.Val)
	})

	fmt.Println(arr)
	fmt.Println(arr2)
}

//递归实现
func pre1(n *Node, f func(node *Node)) {
	if n == nil {
		return
	}

	f(n)
	if n.Children != nil {
		for i := 0; i < len(n.Children); i++ {
			pre1(n.Children[i], f)
		}
	}
}

//迭代实现
func pre2(root *Node, f func(node *Node)) {
	sta := NewStack()
	sta.Push(root)
	for sta.Len() > 0 {
		cur := sta.Pop().(*Node)
		f(cur)
		if cur.Children != nil {
			for i := len(cur.Children) - 1; i >= 0; i-- {
				sta.Push(cur.Children[i])
			}
		}
	}
}

//https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
func (*Ref) SAToBST() {
	nums := []int{-10, -3, 0, 5, 9}
	tree := SortArrToBst(nums, 0, len(nums)-1)
	levelOrder(tree, func(node *TreeNode) {
		fmt.Println(node.Val)
	})
}

//把小->大的数组转换成BST(二叉搜索树)
func SortArrToBst(arr []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	rootIndex := (left + right + 1) / 2
	root := &TreeNode{
		Val:   arr[rootIndex],
		Left:  nil,
		Right: nil,
	}
	root.Left = SortArrToBst(arr, left, rootIndex-1)
	root.Right = SortArrToBst(arr, rootIndex+1, right)
	return root
}

//https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/
func (*Ref) MaxNDepth() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{3, []*Node{
				{5, nil},
				{6, nil},
			}},
			{2, nil},
			{4, nil},
		},
	}
	fmt.Println(Ndepth(root))
}
func Ndepth(root *Node) int {
	if root == nil {
		return 0
	}

	max1 := 0
	if root.Children != nil {
		for _, c := range root.Children {
			cDepth := Ndepth(c)
			if max1 < cDepth {
				max1 = cDepth
			}
		}
	}
	return max1 + 1
}

//https://leetcode-cn.com/problems/increasing-order-search-tree/
func (*Ref) ICBST() {
	root := ArrToNode([]int{5, 3, 6, 2, 4, -1, 8, 1, -1, -1, -1, 7, 9})
	inOrder(ToRight(root), func(curNode *TreeNode) {
		fmt.Println(curNode.Val)
	})
}
func ToRight(root *TreeNode) *TreeNode {
	var rootNew *TreeNode
	var tmp *TreeNode

	inOrder(root, func(curNode *TreeNode) {
		if rootNew == nil {
			rootNew = curNode
			tmp = rootNew
		} else {
			tmp.Right = &TreeNode{
				Val:   curNode.Val,
				Left:  nil,
				Right: nil,
			}
			tmp = tmp.Right
		}
	})
	return rootNew
}

//https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
func (*Ref) LO() {
	root := ArrToNode([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println(lo(root))
}
func lo(root *TreeNode) [][]int {
	arr := make([][]int, 0)

	que := NewQueue()
	que.InQueue(root)
	tmp := NewQueue()
	arr1 := []int{}
	for que.Len() > 0 || tmp.Len() > 0 {
		if que.Len() == 0 {
			que = tmp
			tmp = NewQueue()
			arr = append(arr, arr1)
			arr1 = []int{}
		}

		node := que.OutQueue().(*TreeNode)
		arr1 = append(arr1, node.Val)
		if node.Left != nil {
			tmp.InQueue(node.Left)
		}
		if node.Right != nil {
			tmp.InQueue(node.Right)
		}
	}

	arr = append(arr, arr1)
	return arr
}

func (*Ref) TrimTree() {
	root := ArrToNode([]int{1, 0, 2})
	L := 1
	R := 2

	root = tt(root, L, R)

	levelOrder(root, func(node *TreeNode) {
		fmt.Println(node.Val)
	})
}
func tt(treeNode *TreeNode, L, R int) *TreeNode {
	if treeNode == nil {
		return nil
	}

	if treeNode.Val > R {
		return tt(treeNode.Left, L, R)
	} else if treeNode.Val < L {
		return tt(treeNode.Right, L, R)
	}

	treeNode.Left = tt(treeNode.Left, L, R)
	treeNode.Right = tt(treeNode.Right, L, R)
	return treeNode
}

//https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
func (*Ref) LOB() {
	root := ArrToNode([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println(levelOrder1(root))
}
func levelOrder1(root *TreeNode) [][]int {
	que1 := NewQueue()
	next := NewQueue()
	que1.InQueue(root)
	sta := NewStack()
	arr := []int{}
	for que1.Len() > 0 || next.Len() > 0 {
		if que1.Len() == 0 {
			que1 = next
			next = NewQueue()
			sta.Push(arr)
			arr = []int{}
		}

		no := que1.OutQueue().(*TreeNode)
		arr = append(arr, no.Val)
		if no.Left != nil {
			next.InQueue(no.Left)
		}
		if no.Right != nil {
			next.InQueue(no.Right)
		}
	}
	sta.Push(arr)

	arr1 := make([][]int, 0)
	for sta.Len() > 0 {
		arr1 = append(arr1, sta.Pop().([]int))
	}
	return arr1
}

//https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/
func (*Ref) Aof() {
	root := ArrToNode([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println(levelOrder2(root))
}
func levelOrder2(root *TreeNode) []float64 {
	que1 := NewQueue()
	next := NewQueue()
	que1.InQueue(root)
	arr1 := []float64{}
	arr := []int{}
	for que1.Len() > 0 || next.Len() > 0 {
		if que1.Len() == 0 {
			que1 = next
			next = NewQueue()
			var av float64
			for _, v := range arr {
				av += float64(v)
			}
			arr1 = append(arr1, av/float64(len(arr)))
			arr = []int{}
		}

		no := que1.OutQueue().(*TreeNode)
		arr = append(arr, no.Val)
		if no.Left != nil {
			next.InQueue(no.Left)
		}
		if no.Right != nil {
			next.InQueue(no.Right)
		}
	}
	var av float64
	for _, v := range arr {
		av += float64(v)
	}
	arr1 = append(arr1, av/float64(len(arr)))

	return arr1
}

//https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
func (*Ref) Lca() {
	root := ArrToNode([]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5})
	p := &TreeNode{Val: 2}
	q := &TreeNode{Val: 3}

	fmt.Println(lca(root, p, q).Val)
	fmt.Println(lca1(root, p, q).Val)
}

//数组
func lca(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	//a > b
	if p.Val > q.Val {
		p, q = q, p
	}
	arrA, arrB := []*TreeNode{}, []*TreeNode{}
	tmp := root
	for {
		if tmp.Val == p.Val {
			arrA = append(arrA, tmp)
			break
		}

		arrA = append(arrA, tmp)
		if p.Val < tmp.Val {
			tmp = tmp.Left
		} else {
			tmp = tmp.Right
		}
	}

	tmp = root
	for {
		if tmp.Val == q.Val {
			arrB = append(arrB, tmp)
			break
		}
		arrB = append(arrB, tmp)
		if q.Val < tmp.Val {
			tmp = tmp.Left
		} else {
			tmp = tmp.Right
		}
	}

	if len(arrA) > len(arrB) {
		arrA, arrB = arrB, arrA
	}

	tmp = nil
	for i := range arrA {
		if arrA[i].Val == arrB[i].Val {
			tmp = arrA[i]
		}
	}

	return tmp
}

//递归 N N
func lca1(root, p, q *TreeNode) *TreeNode {
	parentV := root.Val
	pVal := p.Val
	qVal := q.Val
	if pVal > parentV && qVal > parentV {
		return lca1(root.Right, p, q)
	} else if pVal < parentV && qVal < parentV {
		return lca1(root.Left, p, q)
	} else {
		return root
	}
}

//迭代
func lca2(root, p, q *TreeNode) *TreeNode {
	pVal := p.Val
	qVal := q.Val
	parent := root

	for parent != nil {
		if pVal > parent.Val && qVal > parent.Val {
			parent = parent.Right
		} else if pVal < parent.Val && qVal < parent.Val {
			parent = parent.Left
		} else {
			return parent
		}
	}
	return nil
}

//https://leetcode-cn.com/problems/binary-tree-paths/
func (*Ref) Btp() {
	root := ArrToNode([]int{1, 2, 3, -1, 5, 6})
	if root == nil {
		return
	}

	fmt.Println(btp(root, "", []string{}))
}
func btp(root *TreeNode, path string, paths []string) []string {
	if root == nil {
		return paths
	}

	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		paths = append(paths, path)
	} else {
		path += "->"
		paths = btp(root.Left, path, paths)
		paths = btp(root.Right, path, paths)
	}
	return paths
}

//https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers/
func (*Ref) SRTL() {
	root := ArrToNode([]int{1, 0, 1, 0, 1, 0, 1})

	fmt.Println(srtl(root, 0, []int{}))
}
func srtl(root *TreeNode, sum int, qu []int) int {
	if root == nil {
		return sum
	}

	qu = append(qu, root.Val)
	if root.Left == nil && root.Right == nil {
		Len := len(qu) - 1
		for _, v := range qu {
			if v > 0 {
				sum += 1 << uint(Len)
			}
			Len--
		}
	} else {
		sum = srtl(root.Left, sum, qu)
		sum = srtl(root.Right, sum, qu)
	}
	return sum
}

//https://leetcode-cn.com/problems/leaf-similar-trees/
func (*Ref) LS() {
	root1 := ArrToNode([]int{})
	root2 := ArrToNode([]int{})

	l1 := []int{}
	preOrder(root1, func(curNode *TreeNode) {
		if curNode.Left == nil && curNode.Right == nil {
			l1 = append(l1, curNode.Val)
		}
	})
	l2 := []int{}
	preOrder(root2, func(curNode *TreeNode) {
		if curNode.Left == nil && curNode.Right == nil {
			l2 = append(l2, curNode.Val)
		}
	})
	if len(l1) != len(l2) {
		return
	}
	for i := range l1 {
		if l1[i] != l2[i] {
			return
		}
	}
	return
}

func (*Ref) CBST() {
	root := ArrToNode([]int{5, 2, 13})
	val := 0
	InOrderReverse(root, func(curNode *TreeNode) {
		curNode.Val += val
		val = curNode.Val
	})
}
func InOrderReverse(node *TreeNode, f func(curNode *TreeNode)) {
	if node == nil {
		return
	}
	InOrderReverse(node.Right, f)
	f(node)
	InOrderReverse(node.Left, f)
}

//https://leetcode-cn.com/problems/same-tree/submissions/
func (*Ref) SameTree() {
	p := ArrToNode([]int{1, 2, 1})
	q := ArrToNode([]int{1, 1, 2})
	a := true
	preOrder2(p, q, func(node1 *TreeNode, node2 *TreeNode) {
		if node1 == nil || node2 == nil {
			if node1 != nil || node2 != nil {
				a = false
			}
		} else {
			if node1.Val != node2.Val {
				a = false
			}
		}
	})
	fmt.Println(a)
}
func preOrder2(n1, n2 *TreeNode, f func(node1 *TreeNode, node2 *TreeNode)) {
	f(n1, n2)
	if n1 != nil && n2 != nil {
		preOrder2(n1.Left, n2.Left, f)
		preOrder2(n1.Right, n2.Right, f)
	}
}

//https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/
func (*Ref) IsBan() {
	root := ArrToNode([]int{1, 2, 2, 3, -1, -1, 3, 4, -1, -1, 4})
	fmt.Println(levelOrderHigh(root) != -1)
}
func levelOrderHigh(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := levelOrderHigh(root.Left)
	if left == -1 {
		return -1
	}
	right := levelOrderHigh(root.Right)
	if right == -1 {
		return -1
	}

	if abs(left-right) < 2 {
		return max(left, right) + 1
	}
	return -1
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
func (*Ref) IsSymm() {
	root := ArrToNode([]int{1, 2, 2, 3, 4, 4, 3})
	sym := true
	defer func() {
		fmt.Println(sym)
	}()

	if root == nil {
		return
	}
	symm(root.Left, root.Right, func(cn1, cn2 *TreeNode) {
		if cn1 == nil || cn2 == nil {
			if cn1 != nil || cn2 != nil {
				sym = false
			}
		}
		if cn1 != nil && cn2 != nil {
			if cn1.Val != cn2.Val {
				sym = false
			}
		}
	})
}

//https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
func (*Ref) Gmd() {
	root := ArrToNode([]int{})

	prev := -1
	gap := math.MaxInt64
	inOrder(root, func(curNode *TreeNode) {
		if prev >= 0 {
			gap = min(gap, abs(curNode.Val-prev))
		}
		prev = curNode.Val
	})
}

//https://leetcode-cn.com/problems/binary-tree-tilt/
func (*Ref) FindTilt() {
	root := ArrToNode([]int{1, 2, 3, 4})

	tilt := 0
	after(root, func(ls, rs int) {
		tilt += abs(ls - rs)
	})
	fmt.Println(tilt)
}
func after(node *TreeNode, f func(ls, rs int)) int {
	if node == nil {
		return 0
	}

	leftSum := after(node.Left, f)
	rightSum := after(node.Right, f)
	f(leftSum, rightSum)
	return leftSum + rightSum + node.Val
}

//https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/
func (*Ref) FindTarget() {
	root := ArrToNode([]int{2, 1, 3})
	k := 1

	exists := false
	inOrder(root, func(curNode *TreeNode) {
		if inBst(root, k-curNode.Val, curNode) {
			exists = true
		}
	})
	fmt.Println(exists)
}
func inBst(root *TreeNode, target int, avoid *TreeNode) bool {
	if root == nil {
		return false
	}

	if target > root.Val {
		return inBst(root.Right, target, avoid)
	} else if target < root.Val {
		return inBst(root.Left, target, avoid)
	}
	return target == root.Val && root != avoid
}

//https://leetcode-cn.com/problems/path-sum-iii/
func (*Ref) PathSum() {
	root := ArrToNode([]int{10, 5, -3, 3, 2, -1, 11, 3, -2, -1, 1})
	sum := 8
	fmt.Println(pathSum1(root, sum, []int{}))
	//[5,4,8,11,null,13,4,7,2,null,null,5,1]
	//22
}

func pathSum1(root *TreeNode, sum int, pathSum []int) int {
	if root == nil {
		return 0
	}
	tmp := root.Val
	n := 0
	if root.Val == sum {
		n++
	}
	for i := len(pathSum) - 1; i >= 0; i-- {
		tmp += pathSum[i]
		if tmp == sum {
			n++
		}
	}
	pathSum = append(pathSum, root.Val)

	return n + pathSum1(root.Left, sum, pathSum) + pathSum1(root.Right, sum, pathSum)
}

//https://leetcode-cn.com/problems/sum-of-left-leaves/
func (*Ref) SLL() {
	root := ArrToNode([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println(sll(root, 0, ""))
}

func sll(n *TreeNode, sum int, ty string) int {
	if n == nil {
		return sum
	}
	if ty == "left" && n.Left == nil && n.Right == nil {
		return sum + n.Val
	}
	sum = sll(n.Left, sum, "left")
	sum = sll(n.Right, sum, "right")
	return sum
}

//https://leetcode-cn.com/problems/symmetric-tree/
func (*Ref) IsSym() {
	root := ArrToNode([]int{1, 2, 2, 3, 4, 4, 3})

	sym := true
	if root == nil {
		return
	}
	symm(root.Left, root.Right, func(cn1, cn2 *TreeNode) {
		if cn1 == nil || cn2 == nil {
			if cn1 != nil || cn2 != nil {
				sym = false
			}
		}
		if cn1 != nil && cn2 != nil {
			if cn1.Val != cn2.Val {
				sym = false
			}
		}
	})

	log.Printf("%v", sym)
}

//https://leetcode-cn.com/problems/balanced-binary-tree/
func (*Ref) IsBalanced() {
	root := ArrToNode([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println(isBalanced1(root))
}
func isBalanced1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if isBalanced1(root.Left) && isBalanced1(root.Right) &&
		abs(height(root.Left)-height(root.Right)) <= 1 {
		return true
	}
	return false
}
func height(root *TreeNode) int {
	if root == nil {
		return -1
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	return max(leftHeight, rightHeight) + 1
}

//https://leetcode-cn.com/problems/maximum-binary-tree/
func (*Ref) CMB() {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := cmb(nums)
	if root != nil {
		levelOrder(root, func(node *TreeNode) {
			fmt.Println(node.Val)
		})
	}
}
func cmb(a []int) *TreeNode {
	if len(a) < 1 {
		return nil
	}

	left := 0
	right := len(a) - 1
	maxI := left
	for left <= right {
		if a[left] > a[maxI] {
			maxI = left
		}
		if a[right] > a[maxI] {
			maxI = right
		}
		left++
		right--
	}

	n := &TreeNode{
		Val:   a[maxI],
		Left:  nil,
		Right: nil,
	}
	n.Left = cmb(a[:maxI])
	n.Right = cmb(a[maxI+1:])
	return n
}

//https://leetcode-cn.com/problems/sum-of-nodes-with-even-valued-grandparent/submissions/
func (*Ref) SEG() {
	root := ArrToNode([]int{})
	fmt.Println(seg(root, 1, 1))
}
func seg(node *TreeNode, parentVal, grandparentVal int) int {
	if node == nil {
		return 0
	}
	ans := 0
	if grandparentVal%2 == 0 {
		ans += node.Val
	}
	ans += seg(node.Left, node.Val, parentVal)
	ans += seg(node.Right, node.Val, parentVal)
	return ans
}

//https://leetcode-cn.com/problems/deepest-leaves-sum/
var maxDep = -1
var total = 0

func (*Ref) DLS() {
	root := ArrToNode([]int{1, 2, 3, 4, 5, -1, 6, 7, -1, -1, -1, -1, 8})
	preOrderDLS(root, 0, -1)
	fmt.Println(total)
}
func preOrderDLS(current *TreeNode, dep, total int) {
	if current != nil {
		if dep > maxDep {
			maxDep = dep
			total = current.Val
		} else if dep == maxDep {
			total += current.Val
		}
		preOrderDLS(current.Left, dep+1, total)
		preOrderDLS(current.Right, dep+1, total)
	}
}

//https://leetcode-cn.com/problems/binary-search-tree-iterator/
func (*Ref) BSTIF() {

}

//https://leetcode-cn.com/problems/all-elements-in-two-binary-search-trees/
func (*Ref) GAE() {
	root1 := ArrToNode([]int{2, 1, 4})
	root2 := ArrToNode([]int{1, 0, 3})

	arr1 := []int{}
	inOrder(root1, func(curNode *TreeNode) {
		arr1 = append(arr1, curNode.Val)
	})

	arr2 := []int{}
	inOrder(root2, func(curNode *TreeNode) {
		arr2 = append(arr2, curNode.Val)
	})

	arr3 := make([]int, len(arr1)+len(arr2))
	i, j, x := 0, 0, 0
	for i < len(arr1) || j < len(arr2) {
		if i >= len(arr1) {
			arr3[x] = arr2[j]
			j++
		} else if j >= len(arr2) {
			arr3[x] = arr1[i]
			i++
		} else {
			if arr1[i] <= arr2[j] {
				arr3[x] = arr1[i]
				i++
			} else {
				arr3[x] = arr2[j]
				j++
			}
		}
		x++
	}
	fmt.Println(arr3)
}

//https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
func (*Ref) IBST() {
	root := ArrToNode([]int{4, 2, 7, 1, 3})
	val := 5
	levelOrder(ibst(root, val), func(node *TreeNode) {
		fmt.Println(node.Val)
	})

}
func ibst(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = ibst(root.Right, val)
	} else {
		root.Left = ibst(root.Left, val)
	}

	return root
}

func (*Ref) IsValidBST() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	log.Printf("%v", isValidBST(root))
}

func isValidBST(root *TreeNode) bool {
	var preNode *TreeNode
	isValid := true

	inOrder(root, func(curNode *TreeNode) {
		if preNode == nil {
			preNode = curNode
			return
		}

		if preNode.Val >= curNode.Val {
			isValid = false
			return
		}

		preNode = curNode
	})

	return isValid
}

func isSymmetric(root *TreeNode) bool {
	sym := true
	if root == nil {
		return sym
	}
	symm(root.Left, root.Right, func(cn1, cn2 *TreeNode) {
		if cn1 == nil && cn2 != nil {
			sym = false
		}
		if cn1 != nil && cn2 == nil {
			sym = false
		}

		if cn1 != nil && cn2 != nil {
			if cn1.Val != cn2.Val {
				sym = false
			}
		}
	})
	return sym
}

func (r *Ref) LevelOrder12() {
	var root *TreeNode

	v := levelOrder12(root)
	log.Printf("%v", v)
}

func levelOrder12(root *TreeNode) [][]int {
	data := make([][]int, 0)

	q1 := NewQueue()
	q1.InQueue(root)
	q2 := NewQueue()

	data1 := make([]int, 0)
	for {
		if q1.Len() == 0 && q2.Len() == 0 {
			break
		}

		if q1.Len() == 0 {
			q1 = q2
			q2 = NewQueue()
			data = append(data, data1)
			data1 = make([]int, 0)
		}

		e := q1.OutQueue()
		if e == nil {
			break
		}

		node := e.(*TreeNode)
		data1 = append(data1, node.Val)
		if node.Left != nil {
			q2.InQueue(node.Left)
		}
		if node.Right != nil {
			q2.InQueue(node.Right)
		}
	}

	if len(data1) > 0 {
		data = append(data, data1)
	}

	return data
}
