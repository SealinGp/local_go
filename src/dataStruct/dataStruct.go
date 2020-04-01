package main

import (
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
	sta := []int{}
	sta  = append(sta,1,2,3)
	for i := len(sta)-1;i >= 0; i-- {
		log.Println(sta[i])
	}
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