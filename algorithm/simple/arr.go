package simple

import (
	"log"
	"math"
)

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

func (r *Ref) MaxProfit() {
	log.Printf("%v", maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

//对于每组i,j 求max(prices[i], prices[j]) j > i
//假设历史最低价格minPrice
//假设我们在第i天卖出股票
//profit = prices[i] - minPrice
func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}

	}

	return maxProfit
}
