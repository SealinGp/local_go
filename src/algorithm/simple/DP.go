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

//https://leetcode-cn.com/problems/contiguous-sequence-lcci/
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

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
func (*Ref)MP()  {
	//1 - 4 +
	prices   := []int{2,4,2}

	maxI := arrMax(prices,0)

	//动态规划-递归
	fmt.Println(mp(prices,0,maxI))

	//动态规划-数组
	fmt.Println(mp2(prices,maxI))

	//动态规划-滚动数组
	fmt.Println(mp3(prices))

	fmt.Println(mp4(prices))
}

func arrMax(arr []int,i int) int {
	maxI := i
	for j := i+1; j < len(arr); j++  {
		if arr[j] > arr[maxI] {
			maxI = j
		}
	}
	return maxI
}
func max(i,j int) int {
	if j > i {
		i = j
	}
	return i
}
func mp(arr []int,i,ma int) int {
	if len(arr) <= 1 || i > len(arr)-1 {
		return 0
	}

	if ma < i {
		ma = arrMax(arr,i)
	}
	return max(arr[ma] - arr[i],mp(arr,i+1,ma))
}
func mp2(arr []int,maxI int) int {
	maxPriceArr := make([]int,len(arr))
	for i := range arr {
		if i+1 > len(arr) {
			maxPriceArr[i] = 0
		} else {
			if maxI < i {
				maxI = arrMax(arr,i)
			}
			maxPriceArr[i] = arr[maxI] - arr[i]
		}
	}
	maxPriceGapIndex := 0
	for i := 1; i < len(maxPriceArr) ; i++ {
		if maxPriceArr[i] > maxPriceArr[maxPriceGapIndex] {
			maxPriceGapIndex = i
		}
	}
	return maxPriceArr[maxPriceGapIndex]
}
func mp3(arr []int) int {
	maxI       := arrMax(arr,0)
	maxPriceGap := 0
	for i := range arr {
		if i+1 < len(arr) {
			if maxI < i {
				maxI = arrMax(arr,i)
			}
			if arr[maxI] - arr[i] > maxPriceGap {
				maxPriceGap = arr[maxI] - arr[i]
			}
		}
	}
	return maxPriceGap
}
func mp4(arr []int) int {
	minPrice  := 2147483647
	maxProfit := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < minPrice {
			minPrice = arr[i]
		} else if arr[i] - minPrice > maxProfit {
			maxProfit = arr[i] - minPrice
		}
	}
	return maxProfit
}

//https://leetcode-cn.com/problems/maximum-subarray/
func (*Ref)Msa()  {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	tmp  := make([]int,len(nums))
	for i := range nums {
		if i < 1 {
			tmp[i] = nums[i]
			continue
		}

		tmp[i] = max(tmp[i-1] + nums[i],nums[i])
	}

	maxA := tmp[0]
	for i := 1; i < len(tmp);i++ {
		if tmp[i] > maxA {
			maxA = tmp[i]
		}
	}
	fmt.Println(maxA)
}

//https://leetcode-cn.com/problems/climbing-stairs/
func (*Ref)Cs()  {
	n   := 44

	tmp := make([]int,n)
	//动态规划状态转移方程 F(i) = F(i-1) + F[i-2] , F(i) 表示阶梯为i的时候,有多少种方法爬到楼顶
	for i := 1; i <= n ; i++ {
		//边界 i < 3 (i = 2,F(2) = 2  i = 1,F(1) = 1, F(3) = F(2) + F(1) = 3)
		if i-1 == 0 {
			tmp[i-1] = 1
			continue
		}
		if i-3 < 0 {
			tmp[i-1] = 2
			continue
		}

		//i >= 3
		tmp[i-1] = tmp[i-2] + tmp[i-3]
	}

	//长度为n,但是索引为n-1
	fmt.Println(tmp[n-1])
}