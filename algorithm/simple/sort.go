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
