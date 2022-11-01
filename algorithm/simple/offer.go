package simple

import "log"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	arr := make([]int, 0)
	for current := head; current != nil; current = current.Next {
		arr = append(arr, current.Val)
	}

	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func (*Ref) FindRepeatNumber() {
	log.Printf("%v", findRepeatNumber([]int{3, 4, 2, 1, 1, 0}))
}

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}

			v := nums[i]
			nums[i], nums[v] = nums[v], nums[i]
		}
	}

	return 0
}
