package simple

import "log"

func numWays(n int) int {
	if n <= 2 {
		return n
	}

	f1 := 1
	f2 := 2
	for i := 2; i <= n; i++ {
		current := (f1 + f2) % mod
		f1 = f2
		f2 = current
	}

	return f2
}

func (*Ref) Fib() {
	log.Printf("%v", fib(45))
	log.Printf("134903163")
}

const mod int = 1e9 + 7

func fib(n int) int {
	if n <= 1 {
		return n
	}

	j1 := 1
	j2 := 0
	for i := 2; i <= n; i++ {
		current := (j1 + j2) % mod
		j2 = j1
		j1 = current
	}

	return j1
}

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
