package simple

import (
	"log"
)

//-.二分查找
//-.归并排序
//-.快速排序

//https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
// [1,3,5]
//  s m e

// start=0 end=2 mid=1
//  3 >= 1 ? ok

// s=1 e=2 m=2

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	if len(numbers) == 1 {
		return numbers[0]
	}

	minFunc := func(arr []int, s, e int) int {
		min := arr[s]
		for i := s + 1; i <= e; i++ {
			if arr[i] < min {
				min = arr[i]
			}
		}

		return min
	}

	start := 0
	end := len(numbers) - 1
	mid := start
	for numbers[start] >= numbers[end] {
		if start+1 == end {
			mid = end
			break
		}

		if numbers[start] == numbers[end] && numbers[start] == numbers[mid] {
			return minFunc(numbers, start, end)
		}

		mid = (start + end) / 2
		if numbers[mid] >= numbers[start] {
			start = mid
			mid++
		} else if numbers[mid] <= numbers[end] {
			end = mid
			mid--
		}

	}

	return numbers[mid]
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnto1s/

//合并两个递增数组,合并后的数组也为递增
func (*Ref) MergeTwoSortedArr() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3

	merge(nums1, m, nums2, n)
	log.Printf("%v", nums1)
}

//二分查找
func (*Ref) FirstBadVersion() {
	v := firstBadVersion(1926205968)
	log.Printf("%v", v)
}

func firstBadVersion(n int) int {
	i := 0
	j := n
	if j-i <= 1 {
		return j
	}

	mid := (i + j) / 2

	for {
		if j-i == 1 {
			return j
		}

		if isBadVersion(mid) {
			j = mid
		} else {
			i = mid
		}

		mid = (j + i) / 2
	}
}

func isBadVersion(version int) bool {
	return version >= 1167880583
}
