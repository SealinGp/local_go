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
//向尾部添加数据
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
func (l *list)Remove(i uint64) bool {
	var node *Node
	if i >= l.size {
		return false
	}

	if i == 0 {
		node   = l.head
		l.head = node.Next
		if l.size == 1 {
			l.tail = nil
		}
	} else {
		//找到前一个节点
		preNode := l.head
		for j := 1; uint64(j) < i; j++ {
			preNode = preNode.Next
		}
		node         = preNode.Next
		preNode.Next = node.Next

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
	i  := 0
	j1 := 0
	for j := 0 ; j < i; j++ {
		fmt.Println(j)
		j1 = j
	}
	fmt.Println(j1)
}