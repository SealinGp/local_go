package simple

import (
	"log"
)

func (*Ref) HW() {
	log.Printf("%v", 1<<10)
	log.Printf("%v", hammingWeight(00000000000000000000000000001011))
}


//&: 1 1 = 1
//^: 异或 不同为1
//|: 有1则1
//1<<n: 1*2^n
func hammingWeight(num uint32) int {
	s := 0
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			s++
		}
	}
	return s
}

func (r *Ref) BalanceIndex() {
	// sumi =  total - num[i] - sumi
	// 2*sumi = total - num[i]
	// i => ?
	log.Printf("%v", BalanceIndex([]int{1, 7, 3, 6, 5, 6}))
}

func BalanceIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	sumi := 0
	for i := 0; i < len(nums); i++ {
		if 2*sumi == total-nums[i] {
			return i
		}

		sumi += nums[i]
	}

	return -1
}
