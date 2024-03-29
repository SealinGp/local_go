package simple

import (
	"container/list"
	"log"
	"math"
	"math/rand"
	"sort"
	"time"
)

func getCommon(nums1 []int, nums2 []int) int {
	minCommon := -1
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	if nums1Len == 0 || nums2Len == 0 {
		return minCommon
	}

	minLen := nums1Len
	if nums2Len < minLen {
		minLen = nums2Len
	}

	i := 0
	j := 0
	for i < nums1Len && j < nums2Len {
		iVal := nums1[i]
		jVal := nums2[j]

		if iVal == jVal {
			return iVal
		}

		if iVal < jVal {
			i++
			continue
		}

		j++
	}

	return minCommon
}

func (*Ref) M2() {
	nums1 := []int{0}
	nums2 := []int{1}
	m := 0
	n := 1

	merge(nums1, m, nums2, n)
	log.Printf("%v", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1 //0
	j := n - 1 //0

	for x := m + n - 1; x >= 0; x-- {
		if x == i {
			break
		}

		if i < 0 {
			nums1[x] = nums2[j]
			j--
			continue
		}

		if j < 0 {
			nums1[x] = nums1[i]
			i--
			continue
		}

		l, r := nums1[i], nums2[j]
		if l > r {
			nums1[x] = l
			i--
		} else {
			nums1[x] = r
			j--
		}
	}
}

func (*Ref) HW() {
	log.Printf("%v", 1<<10)
	log.Printf("%v", hammingWeight(00000000000000000000000000001011))
}

// &: 1 1 = 1
// ^: 异或 不同为1
// |: 有1则1
// 1<<n: 1*2^n
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

func (r *Ref) MaxProfit() {
	log.Printf("%v", maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

// 对于每组i,j 求max(prices[i], prices[j]) j > i
// 假设历史最低价格minPrice
// 假设我们在第i天卖出股票
// profit = prices[i] - minPrice
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

//[1,2,3]
//[-,2,3,4,5,6]
//如何取交集?

//[[1,3],[2,6]]
//[1,6]

func (*Ref) Merge() {

	// [[5,5],[1,2],[2,4],[2,3],[4,4],[5,5],[2,3],[5,6],[0,0],[5,6]]

	log.Printf("%v", merge1([][]int{
		{5, 5},
		{1, 2},
		{2, 4},
		{2, 3},
		{4, 4},
		{5, 5},
		{2, 3},
		{5, 6},
		{0, 0},
		{5, 6},
	}))

}

type I struct {
	intervals [][]int
}

func (i1 *I) Len() int {
	return len(i1.intervals)
}

func (i1 *I) Less(i, j int) bool {
	return i1.intervals[i][0]-i1.intervals[j][1] < 0
}

func (i1 *I) Swap(i, j int) {
	i1.intervals[i], i1.intervals[j] = i1.intervals[j], i1.intervals[i]
}

// [[1,4],[0,2],[3,5]]
func merge1(intervals [][]int) [][]int {
	i := &I{
		intervals: intervals,
	}
	sort.Sort(i)

	newIntervals := make([][]int, 0)

	hasJ := false
	for i := 0; i < len(intervals); {
		interval := intervals[i]
		newInterval := make([]int, 0)
		newInterval = append(newInterval, interval...)

		if i+1 < len(intervals) {
			nextInterval := intervals[i+1]
			if J(interval, nextInterval) {
				newInterval[0] = interval[0]
				newInterval[1] = interval[1]

				if nextInterval[0] < newInterval[0] {
					newInterval[0] = nextInterval[0]
				}

				if nextInterval[1] > newInterval[1] {
					newInterval[1] = nextInterval[1]
				}

				i += 2
				newIntervals = append(newIntervals, newInterval)
				hasJ = true
				continue
			}

		}

		newIntervals = append(newIntervals, newInterval)
		i++
	}

	if hasJ {
		return merge1(newIntervals)
	}

	return newIntervals
}

func J(a, b []int) bool {
	m1 := make(map[int]struct{})
	for i := a[0]; i <= a[1]; i++ {
		m1[i] = struct{}{}
	}

	for i := b[0]; i <= b[1]; i++ {
		if _, ok := m1[i]; ok {
			return ok
		}
	}

	return false
}
