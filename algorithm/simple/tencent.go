package simple

import (
	"fmt"
	"math/rand"
	"time"
)

//https://leetcode-cn.com/problems/reverse-linked-list/
func (*Ref) RLL() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next

		curr.Next = prev
		prev = curr

		curr = next
	}
}

//https://leetcode-cn.com/problems/merge-two-sorted-lists/
func (*Ref) MTSL() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	n := mtsl(l1, l2)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
func mtsl(l1, l2 *ListNode) *ListNode {
	if l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l1.Next = mtsl(l1.Next, l2)
			return l1
		} else {
			l2.Next = mtsl(l1, l2.Next)
			return l2
		}
	}
	if l1 == nil {
		return l2
	}
	return l1
}

//https://leetcode-cn.com/problems/two-sum/
func (*Ref) TS() {
	nums := []int{3, 3}
	target := 6

	m := make(map[int]int, len(nums))
	for i := range nums {
		m[nums[i]] = i
	}

	a := []int{}
	for i, v := range nums {
		left := target - v
		a = []int{i}
		if leftI, ok := m[left]; ok {
			if leftI != i {
				a = append(a, leftI)
				break
			}
		}
	}
	fmt.Println(a)
}

//https://leetcode-cn.com/problems/3sum/
func (*Ref) TreeSum() {
	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}

	numsMaxI := len(nums) - 1
	//先排序
	dobuleFastSort(nums, 0, numsMaxI)
	tmp := make(map[int]int, len(nums))
	for i, v := range nums {
		tmp[v] = i
	}

	arr := [][]int{}
	for i, v := range nums {
		//跟上一个元素不同,避免重复
		if i == 0 || v != nums[i-1] {
			third := numsMaxI
			for second := i + 1; second <= numsMaxI; second++ {
				if second == i+1 || nums[second-1] != nums[second] {
					for second < third && nums[i]+nums[second]+nums[third] > 0 {
						third--
					}
					if second == third {
						break
					}
					if nums[i]+nums[second]+nums[third] == 0 {
						arr = append(arr, []int{nums[i], nums[second], nums[third]})
					}
				}
			}
		}
	}
	fmt.Println(arr)
}
func dobuleFastSort(arr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}

	privotIndex := dfs(arr, startIndex, endIndex)
	dobuleFastSort(arr, startIndex, privotIndex-1)
	dobuleFastSort(arr, privotIndex+1, endIndex)
}
func dfs(arr []int, startIndex, endIndex int) int {
	privot := arr[startIndex]
	left := startIndex
	right := endIndex
	for right != left {
		//右指针左移
		for left < right && arr[right] > privot {
			right--
		}
		//左指针右移
		for left < right && arr[left] <= privot {
			left++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[startIndex], arr[left] = arr[left], privot
	return left
}

//https://leetcode-cn.com/problems/linked-list-cycle/
func (*Ref) HC() {
	head := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
	}

	isCycle := false
	if head == nil {
		return
	}
	slow := head
	fast := head
	for {
		slow = slow.Next
		if slow == nil {
			break
		}
		if fast == nil || fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		if slow == nil || fast == nil {
			break
		}
		if slow.Val == fast.Val {
			isCycle = true
			break
		}
	}

	fmt.Println(isCycle)
}

//https://leetcode-cn.com/problems/implement-rand10-using-rand7/
func (*Ref) Rand10() {
	fmt.Println(rand10())
}
func rand10() int {
	row, col, idx := 0, 0, 0
	for {
		row = rand7()
		col = rand7()
		idx = col + (row-1)*7
		if idx <= 40 {
			break
		}
	}
	return 1 + (idx-1)%10

}
func rand7() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(7)
}

//https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
func (*Ref) FMSA() {
	nums1 := []int{3}
	nums2 := []int{-2, -1}

	nums1L := len(nums1)
	nums2L := len(nums2)
	//偶数
	isEven := (nums1L+nums2L)%2 == 0
	if nums1L > nums2L {
		nums1, nums2 = nums2, nums1
		nums1L, nums2L = nums2L, nums1L
	}

	findIndex := (nums1L + nums2L) / 2
	mid := 0
	i1 := 0
	i2 := 0
	for {
		if isEven && findIndex == 1 {
			if i1 < nums1L && i2 < nums2L {
				if nums1[i1] < nums2[i2] {
					mid += nums1[i1]
				} else {
					mid += nums2[i2]
				}
			} else if i1 < nums1L {
				mid += nums1[i1]
			} else if i2 < nums2L {
				mid += nums2[i2]
			}
		}
		if findIndex == 0 {
			if i1 < nums1L && i2 < nums2L {
				if nums1[i1] < nums2[i2] {
					mid += nums1[i1]
				} else {
					mid += nums2[i2]
				}
			} else if i1 < nums1L {
				mid += nums1[i1]
			} else if i2 < nums2L {
				mid += nums2[i2]
			}
			break
		}

		if i1 < nums1L && i2 < nums2L {
			if nums1[i1] < nums2[i2] {
				i1++
			} else {
				i2++
			}
		} else if i1 < nums1L {
			i1++
		} else if i2 < nums2L {
			i2++
		} else {
			break
		}
		findIndex--
	}

	if isEven {
		fmt.Println(float64(mid) / 2)
		return
	}
	fmt.Println(mid)
}
