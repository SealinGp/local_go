package simple

import (
	"fmt"
)

//二分查找
func (*Ref)TwoSearch()  {
	nums  := []int{
		1,2,3,4,5,
	}
	target := 5
	fmt.Println(search(nums,target))
}
func search(nums []int,target int) int {
	start := 0
	end   :=  len(nums)-1
	for start + 1 < end  {
		mid := start + (end - start)/2
		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			start = mid
		} else if target < nums[mid] {
			end   = mid
		}
	}

	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}