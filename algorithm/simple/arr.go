package simple

import (
	"container/list"
	"log"
	"math"
	"math/rand"
	"time"
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

type Solution struct {
	ori []int
}

func Constructor1(nums []int) Solution {
	rand.Seed(time.Now().UnixNano())

	s := Solution{
		ori: nums,
	}

	return s
}

func (this *Solution) Reset() []int {
	return this.ori
}

func (this *Solution) Shuffle() []int {
	cp := make([]int, len(this.ori))
	copy(cp, this.ori)

	cpLen := len(cp)
	for i := range cp {
		j := i + rand.Intn(cpLen-i)
		cp[i], cp[j] = cp[j], cp[i]
	}

	return cp
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func (r *Ref) MinStackTest() {
	ms := Constructor2()

	ms.Push(2)
	ms.Push(0)
	ms.Push(3)
	ms.Push(0)

	log.Printf("%v", ms.GetMin())
	ms.Pop()

	log.Printf("%v", ms.GetMin())
	ms.Pop()

	log.Printf("%v", ms.GetMin())
	ms.Pop()

	log.Printf("%v", ms.GetMin())
}

type MinStack struct {
	l    *list.List
	minL *list.List
}

func Constructor2() MinStack {
	return MinStack{
		l:    list.New(),
		minL: list.New(),
	}
}

// [2,0,3,0] -> 0
// [2,0,3]
func (this *MinStack) Push(val int) {
	this.l.PushBack(val)

	if this.minL.Len() == 0 || val <= this.minL.Back().Value.(int) {
		this.minL.PushBack(val)
	}
}

func (this *MinStack) Pop() {
	if this.l.Back() != nil && this.minL.Back() != nil {
		last := this.l.Back()
		this.l.Remove(last)

		minEle := this.minL.Back()
		if minEle.Value.(int) == last.Value.(int) {
			this.minL.Remove(minEle)
		}
	}
}

func (this *MinStack) Top() int {
	if ele := this.l.Back(); ele != nil {
		return ele.Value.(int)
	}

	return 0
}

func (this *MinStack) GetMin() int {
	if minEle := this.minL.Back(); minEle != nil {
		return minEle.Value.(int)
	}

	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
