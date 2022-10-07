package simple

import "log"

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
