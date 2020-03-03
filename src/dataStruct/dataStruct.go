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
		"linkTable2" : linkTable2,
	}
	fun[os.Args[1]]()
}


type Object interface {}
//单个节点
type Node struct {
	Data Object
	Next *Node
}
//单向链表
type list struct {
	head *Node  //首节点
	tail *Node  //尾节点
	size uint64 //链表总长度
}
func NewList() *list {
	return &list{}
}
//向尾部添加节点
func (l *list)Append(node *Node) bool {
	if node == nil {
		return false
	}
	node.Next = nil

	//新节点放入单向链表中
	if l.size == 0 {
		l.head = node
	} else {
		oldTail     := l.tail
		oldTail.Next = node
	}

	l.tail = node
	l.size++
	return true
}

//插入节点
func (l *list)Insert(i uint64,node *Node) bool {
	//空链表|插入的索引超出范围|空节点 均无法做插入操作
	if l.size == 0 || i > l.size || node == nil {
		return false
	}

	//首个插入
	if i == 0 {
		node.Next = l.head
		l.head    = node
	} else {
		//找到对应索引i的前面一个节点
		iPreNode := l.head
		for j := 0; uint64(j) < i; j++{
			iPreNode = iPreNode.Next
		}
		node.Next     = iPreNode.Next
		iPreNode.Next = node
	}

	l.size++
	return true
}

//删除某节点
func (l *list)Remove(i uint64) bool {
	var node *Node

	//要删除的节点的索引不得 >= 链表长度
	if i >= l.size {
		return false
	}

	//删除首节点
	if i == 0 {
		node   = l.head
		l.head = node.Next
		if l.size == 1 {
			l.tail = nil
		}

	//删除其他节点
	} else {
		//找到索引为i的前一个节点j
		preNode := l.head
		for j := 1; uint64(j) < i; j++ {
			preNode = preNode.Next
		}
		node         = preNode.Next //i节点
		preNode.Next = node.Next    //j节点的next = i节点的next

		//若删除尾部,则链表尾部需要调整
		if i == l.size-1 {
			l.tail = preNode
		}
	}

	l.size--
	return true
}
func (l *list)Get(i uint64) *Node {
	if i >= l.size {
		return nil
	}

	node := l.head
	for j := 0; uint64(j) < i;j++ {
		node = node.Next
	}
	return node
}

func linkTable()  {
	l := NewList()

	n := &Node{
		Data:"123",
	}
	l.Append(n)

	n1 := &Node{
		Data:"1234",
	}
	l.Append(n1)

	for i := 0; uint64(i) < l.size; i++ {
		fmt.Println(l.Get(uint64(i)).Data)
	}

	fmt.Println(11%10)
}

type ListNode struct {
	 Val int
	 Next *ListNode
}
func linkTable2()  {
	l1 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:8,
		},
	}

	l2 := &ListNode{
		Val:0,
	}

	if l1.Val == 0 && l1.Next != nil {
		log.Fatal("l1 begin with zero")
	}
	if l2.Val == 0 && l2.Next != nil {
		log.Fatal("l2 begin with zero")
	}

	l3P      := &ListNode{}
	l3       := l3P
	i        := 0
	nextPlus := 0
	for {
		l3P.Val  = l1.Val + l2.Val
		if l3P.Val >= 10 {
			l3P.Val = l3P.Val%10
			nextPlus++
		}

		if l1.Next == nil && l2.Next == nil {
			if nextPlus != 0 {
				l3P.Next = &ListNode{}
				l3P      = l3P.Next
				l3P.Val++
			}
			break
		}

		l3P.Next = &ListNode{}
		l3P      = l3P.Next
		l1       = l1.Next
		l2       = l2.Next
		if l1 == nil {
			l1 = &ListNode{
				Val:0,
			}
		}
		if l2 == nil {
			l2 = &ListNode{
				Val:0,
			}
		}

		i++
		l1.Val += nextPlus
		nextPlus = 0
	}

	for  {
		fmt.Println(l3.Val)
		l3 = l3.Next
		if l3 == nil {
			break
		}
	}
}