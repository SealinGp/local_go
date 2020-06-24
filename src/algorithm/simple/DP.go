package simple

import (
	"fmt"
)

//本文件全为动态规划(dynamic-programming)相关的题目


//https://leetcode-cn.com/problems/range-sum-query-immutable/
func (*Ref)CS()  {
	arr := NewArr([]int{-2,0,3,-5,2,-1})
	fmt.Println(arr.SumRange3(0,2))
	fmt.Println(arr.SumRange3(2,5))
	fmt.Println(arr.SumRange3(0,5))
}
type NumArray struct {
	Nums  []int
}
func NewArr(nums []int) NumArray {
	Arr :=  NumArray{
		Nums:make([]int,len(nums) + 1),
	}
	for i, v := range nums {
		if i < 1 {
			Arr.Nums[i+1] = v
		} else {
			Arr.Nums[i+1] = Arr.Nums[i] + v
		}
	}
	return Arr
}

func (this *NumArray) SumRange3(i int, j int) int {
	return this.Nums[j+1] - this.Nums[i]
}

func (*Ref)MSA()  {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	dp   := make([]int,len(nums))
	for i, v := range nums {
		if i-1 < 0 {
			dp[i] = v
			continue
		}

		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + v
		} else {
			dp[i] = v
		}
	}

	max := dp[0]
	for i, v := range dp {
		if i == 0 {
			continue
		}

		if max < v {
			max = v
		}
	}
	fmt.Println(max)
}