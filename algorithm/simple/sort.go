package simple

import (
	"log"
)

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnto1s/

//合并两个递增数组,合并后的数组也为递增
func (*Ref) MergeTwoSortedArr() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3

	//nums1
	// 0, 1, 2, 3, 4, 5
	// 2, 3, 9, 0, 0, 0

	//nums2
	// 0, 1, 3
	// 6, 6, 8
	merge(nums1, m, nums2, n)
	log.Printf("%v", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)

	i := 0
	j := 0
	for {
		if i == m {
			sorted = append(sorted, nums2[j:]...)
			break
		}

		if j == n {
			sorted = append(sorted, nums1[i:]...)
			break
		}

		if nums1[i] < nums2[j] {
			sorted = append(sorted, nums1[i])
			i++
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
	}
	copy(nums1, sorted)
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