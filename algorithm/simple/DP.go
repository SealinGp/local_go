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
func min(i,j int) int {
	if j < i {
		i = j
	}
	return i
}
func (*Ref)Cs()  {
	n   := 44

	a,b,c := 1,2,0
	//tmp := 0
	//tmp := make([]int,n)
	//动态规划状态转移方程 F(i) = F(i-1) + F[i-2] , F(i) 表示阶梯为i的时候,有多少种方法爬到楼顶
	for i := 3; i <= n ; i++ {
		c = a + b
		a = b
		b = c
	}

	//长度为n,但是索引为n-1
	fmt.Println(c)
}
func (*Ref)Mccs()  {
	cost := []int{10,15,20}

	tmp  := make([]int,len(cost))
	tmp[0] = 0
	tmp[1] = min(cost[0],cost[1])
	for i := range cost {
		if i < 2 {
			continue
		}
		tmp[i] = min(tmp[i-1] + cost[i],tmp[i-2] + cost[i-1])
	}
	fmt.Println(tmp[len(tmp)-1])
}

func (*Ref)Wts()  {
	n     := 76
	res   := 1000000007
	a,b,c := 1,2,4
	tmp   := 0

	//F(i) = F(i-1) + F(i-2) + F(i-3)
	for i := 4; i <= n; i++ {
		tmp = a + b + c
		a = b
		b = c
		c = tmp % res
	}

	fmt.Println(c)
}

//https://leetcode-cn.com/problems/counting-bits/
func (*Ref)Cb()  {
	num := 5
	arr := make([]int,num+1)
	arr[0] = 0
	for i := 1; i <= num ; i++  {
		if i == 2 {
			arr[i] = 1
			continue
		}
		//偶数: 2 跟 i/2 的值的最大值 F[i] = max(F[i/2],F[2])
		if i % 2 == 0 {
			arr[i] = max(arr[2],arr[i / 2])
		//奇数: 前一个数+1 F[i] = F[i-1] + 1
		} else {
			arr[i] = arr[i-1] + 1
		}
	}
	fmt.Println(arr)
}

//https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones/
func (*Ref)CountS()  {
	matrix := [][]int{
		{0,1,1,1},
		{1,1,1,1},
		{0,1,1,1},
	}

	x    := len(matrix)    //长
	y    := len(matrix[0]) //宽
	dp := make([][]int,x)
	for i := range dp {
		dp[i] = make([]int,y)
	}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				continue
			}

			//i-1 >= 0
			//j-1 >= 0
			//状态转移方程
			if i -1 >= 0 && j-1 >= 0 {
				dp[i][j] = min(
					dp[i-1][j-1],
					min(dp[i-1][j],dp[i][j-1]),
				) + 1
			} else {
				dp[i][j] = 1
			}
		}
	}
	answer := 0
	for i := range dp {
		for j := range dp[i] {
			answer += dp[i][j]
		}
	}

	fmt.Println(answer)
}

//https://leetcode-cn.com/problems/stone-game/
type pair struct {
	first  int
	second int
}
func (*Ref)Sg()  {
	piles := []int{4,5,7,1,10,6,3,5}
	//dp[i][j] 表示从数组索引 i~j 里面 先手|后手 选择拿到的最优数字
	dp    := make([][]pair,len(piles))
	for i := range dp  {
		dp[i] = make([]pair,len(piles))
		dp[i][i].first  = piles[i]
		dp[i][i].second = 0
	}

	for j := 1; j < len(piles) ; j++ {
		i := 0
		for jx := j; jx < len(piles) ;jx++  {
			left  := piles[i]  + dp[i+1][jx].second
			right := piles[jx] + dp[i][jx-1].second

			//先手选左边
			if left > right {
				dp[i][jx].first  = left
				dp[i][jx].second = dp[i+1][jx].first
			} else {
				dp[i][jx].first  = right
				dp[i][jx].second = dp[i][jx-1].first
			}
			i++
		}
	}

	first := dp[0][len(piles)-1].first
	right := dp[0][len(piles)-1].second
	fmt.Println(first > right)
}