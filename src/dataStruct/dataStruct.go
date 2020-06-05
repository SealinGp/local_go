package main

import (
	"fmt"
	"log"
	"os"
)


func main() {
	if len(os.Args) <= 1 {
		log.Fatal("func required")
	}

	fun := map[string]func(){
		"linkTable" : linkTable,
		"stack" : stack,
		"queue" : queue,
		"hashTable" : hashTable,
		"dsu" : dsu,
		"Heap" : Heap,
		"ds1" : ds1,
	}
	fun[os.Args[1]]()
}


//单向链表
type Node struct {
	Data int //数据域
	Next *Node       //指针域
}
//双向链表
type DNode struct {
	Data  int
	Left  *DNode
	Right *DNode
}

//链表 https://oi-wiki.org/ds/linked-list/
func linkTable()  {
	//单向链表插入数据
	SNode := &Node{
		Data: 1,
		Next: &Node{
			Data: 2,
			Next: nil,
		},
	}

	//插入到索引为1 的位置,则找到索引为1-1=0的位置做插入操作
	i1     := 1
	iNode := &Node{
		Data: 3,
		Next: nil,
	}

	//插入链表
	i := 0
	for preN := SNode;preN.Next != nil;preN = preN.Next {
		if i == i1-1 {
			iNode.Next,preN.Next = preN.Next,iNode
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

//栈: https://oi-wiki.org/ds/stack/ 规律: 坐电梯
func stack()  {
	//sta := []int{}
	//sta  = append(sta,1,2,3)
	//for i := len(sta)-1;i >= 0; i-- {
	//	log.Println(sta[i])
	//}
	sta1 := []int{}
	sta1 = append(sta1,1,2)
	fmt.Println(sta1[len(sta1)-1],sta1)

	sta1 = sta1[:len(sta1)-1]
	fmt.Println(sta1[len(sta1)-1],sta1)

	sta1 = sta1[:len(sta1)-1]
	fmt.Println(sta1)
}

//队列: https://oi-wiki.org/ds/queue/ 规律: 排队
func queue()  {
	que := []int{}
	que = append(que,1,2,3)
	for _,v := range que {
		log.Println(v)
	}

	//双端队列:可以在队首|队尾 插入|删除 的队列

	//队首插入
	que = append([]int{0},que...)
	log.Println(que)

	//队首删除
	que = que[1:]
	log.Println(que)

	//队尾插入
	que = append(que,4)
	log.Println(que)

	//队尾删除
	que = que[:len(que)-1]
	log.Println(que)
}

// https://oi-wiki.org/ds/hash/
func hashTable()  {
	m := make(map[string]interface{})
	m["a"] = "123"
	log.Println(m)
}

// https://oi-wiki.org/ds/dsu/ 并合集,用来查找环
// https://www.bilibili.com/video/av38498175/
var set  = make([]int,5)
func dsu()  {
	makeSet(2)
	makeSet(3)
	Union(2,3)
	log.Println(set)

	//makeSet(4)
	//Union(4,3)
	//log.Println(set)
	//
	//log.Println(findSet(3),findSet(2),"?")
}
//新增一个集合,根节点的值为-1
func makeSet(i int)  {
	set[i]  = i
}
//查找集合的根节点
func findSet(i int) int {
	if set[i] == i {
		return i
	}
	return findSet(set[i])
}
//合并集合
func Union(i,j int)  {
	i = findSet(i)
	j = findSet(j)
	if i > j {
		set[j] = i
	} else {
		set[i] = j
	}
}


//https://oi-wiki.org/ds/heap/
//https://zhuanlan.zhihu.com/p/40294173
func Heap()  {

}

// 前序遍历:  根节点 ——> 左子树 ——> 右子树
// 中序遍历:  左子树 ——> 根节点 ——> 右子树
// 后序遍历:  左子树 ——> 右子树 ——> 根节点

//大根二叉树(从上->下,从右->左,从大->小) 排序二叉树
type Leaf struct {
	Key   int     //节点的大小 ?
	Value int     //节点值
	Left  *Leaf   //左节点
	Right *Leaf   //右节点
}
type tree interface {
	Insert(key int,value int)            //插入节点
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
func (t *T1)Insert(key,value int)  {
	n := &Leaf{key,value,nil,nil}

	if t.Root == nil {
		t.Root = n
	} else {
		insertNode(t.Root,n)
	}
}
func insertNode(currentLeaf,newLeaf *Leaf)  {
	if newLeaf.Key < currentLeaf.Key {
		if currentLeaf.Left == nil {
			currentLeaf.Left = newLeaf
		} else {
			insertNode(currentLeaf.Left,newLeaf)
		}
	} else {
		if currentLeaf.Right == nil {
			currentLeaf.Right = newLeaf
		} else {
			insertNode(currentLeaf.Right,newLeaf)
		}
	}
}

//最小的节点在二叉树的最左边
func (t *T1)Min() int {
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

func (t *T1)Max() int {
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

func (t *T1)Search(key int) bool {
	return search(t.Root,key)
}
func search(currentLeaf *Leaf,key int) bool {
	if currentLeaf == nil {
		return false
	}
	if key < currentLeaf.Key {
		return search(currentLeaf.Left,key)
	}
	if key > currentLeaf.Key {
		return search(currentLeaf.Right,key)
	}
	return true
}

//中序
func (t *T1)InOrderTraverse(f func(value int))  {
	inOrderTraver(t.Root,f)
}
func inOrderTraver(currentLeaf *Leaf,f func(value int))  {
	if currentLeaf != nil {
		inOrderTraver(currentLeaf.Left,f)
		f(currentLeaf.Value)
		inOrderTraver(currentLeaf.Right,f)
	}
}

//前序 递归实现
func (t *T1)PreOrderTraverse(f func(value int))  {
	preOrderTraver(t.Root,f)
}
func preOrderTraver(currentLeaf *Leaf,f func(value int))  {
	if currentLeaf != nil {
		f(currentLeaf.Value)
		preOrderTraver(currentLeaf.Left,f)
		preOrderTraver(currentLeaf.Right,f)
	}
}
//前序栈实现
func (t *T1)PreOrderTraverseStack(f func(value int))  {
	stack := make([]*Leaf,0)
	node  := t.Root
	for node != nil || len(stack) != 0 {
		for node != nil {
			f(node.Value)
			stack = append(stack,node)
			node  = node.Left
		}

		if len(stack) != 0 {
			node  = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node  = node.Right
		}
	}
}

//后序
func (t *T1)PostOrderTraverse(f func(value int))  {
	postOrderTraver(t.Root,f)
}
func postOrderTraver(currentLeaf *Leaf,f func(value int))  {
	if currentLeaf != nil {
		postOrderTraver(currentLeaf.Left,f)
		postOrderTraver(currentLeaf.Right,f)
		f(currentLeaf.Value)
	}
}
func (t *T1)String()  {
	fmt.Println(" ---------------------------------------------------- ")
	stringify(t.Root,0)
	fmt.Println(" ---------------------------------------------------- ")
}
func stringify(currentLeaf *Leaf,level int)  {

	if currentLeaf != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "----["
			level++
			stringify(currentLeaf.Left, level)
			fmt.Printf(format+"%d\n", currentLeaf.Key)
			stringify(currentLeaf.Right, level)
		}
	}
}

func ds1()  {
	t1 := T1{
		Root:&Leaf{
			Key:   0,
			Value: 1,
			Left:  nil,
			Right: nil,
		},
	}
	t1.Insert(1,2)
	t1.Insert(2,3)
	t1.Insert(3,4)
	t1.Insert(4,5)
	t1.Insert(6,6)
	//stringify(t1.Root,1)
	t1.PreOrderTraverseStack(func(value int) {
		fmt.Println(value)
	})
	//t1.PreOrderTraverse(func(value int) {
	//	fmt.Println(value)
	//})
}