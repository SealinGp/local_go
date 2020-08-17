package simple

import (
	"fmt"
)

//https://leetcode-cn.com/problems/reverse-linked-list/
func (*Ref)RLL()  {
	head    := &ListNode{Val:1,Next:&ListNode{Val:2,Next:&ListNode{Val:3}}}
	var prev *ListNode
	curr    := head
	for curr != nil  {
		next := curr.Next

		curr.Next = prev
		prev      = curr

		curr = next
	}
}

//https://leetcode-cn.com/problems/merge-two-sorted-lists/
func (*Ref)MTSL()  {
	l1 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	n := mtsl(l1,l2)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
func mtsl(l1,l2 *ListNode) *ListNode {
	if l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l1.Next = mtsl(l1.Next,l2)
			return l1
		} else {
			l2.Next = mtsl(l1,l2.Next)
			return l2
		}
	}
	if l1 == nil {
		return l2
	}
	return l1
}

//https://leetcode-cn.com/problems/two-sum/
func (*Ref)TS()  {
	nums   := []int{3,3}
	target := 6

	m      := make(map[int]int,len(nums))
	for i := range nums {
		m[nums[i]] = i
	}

	a := []int{}
	for i,v := range nums {
		left := target - v
		a = []int{i}
		if leftI,ok := m[left];ok {
			if leftI != i {
				a = append(a,leftI)
				break
			}
		}
	}
	fmt.Println(a)
}
