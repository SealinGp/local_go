package simple

import (
	"log"
	"strconv"
)

func fizzBuzz(n int) []string {
	fb := make([]string, n)

	for i := 1; i <= n; i++ {

		v := ""

		if i%3 == 0 {
			v += "Fizz"
		}

		if i%5 == 0 {
			v += "Buzz"
		}

		if v == "" {
			v = strconv.FormatInt(int64(i), 10)
		}

		fb[i-1] = v
	}

	return fb
}

func (r *Ref) IsPowerOfThree() {
	log.Printf("%v", 2<<30) //2<<1 2^2
}

func isPowerOfThree(n int) bool {
	return (n > 0 && 1162261467%n == 0)
}

// answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数。
// answer[i] == "Fizz" 如果 i 是 3 的倍数。
// answer[i] == "Buzz" 如果 i 是 5 的倍数。
// answer[i] == i （以字符串形式）如果上述条件全不满足。

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xngt85/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// linear-gradient(135deg,#f54ea2,#ff7676)

func countPrimes(n int) int {
	primes := make([]bool, n)
	for i := 0; i < n; i++ {
		primes[i] = true
	}

	primeLen := 0
	for i := 2; i < n; i++ {
		if primes[i] {
			primeLen++

			for j := 2 * i; j < n; j += i {
				primes[j] = false
			}
		}

	}

	return primeLen
}

// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

//"MCMXCIV" 1000
//2216
//1994
func romanToInt(s string) int {
	if len(s) <= 0 {
		return 0
	}

	r2Int := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	for i := 0; i < len(s); i++ {
		cur := r2Int[s[i]]
		if i < len(s)-1 && cur < r2Int[s[i+1]] {
			sum -= cur
		} else {
			sum += cur
		}
	}
	return sum
}

func (*Ref) PivotIndex() {
	//[-1,-1,-1,-1,-1,0]
	log.Printf("%v", pivotIndex([]int{-1, -1, -1, -1, -1, 0}))

}

//左侧元素和 sum
//右侧元素和 total - numi - sum
//sum == total - numi - sum
func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	left := 0
	for i, v := range nums {
		if 2*left+v == total {
			return i
		}

		left += v
	}

	return -1
}

// target = 2
// [1,3,5,6]
//  0 1 2 3

func (*Ref) SearchInsert() {
	searchInsert([]int{1, 3, 5, 6}, 2)
}

func searchInsert(nums []int, target int) int {
	if len(nums) <= 0 {
		return 0
	}

	return searchInsert1(nums, target, 0, len(nums)-1)
}

func searchInsert1(nums []int, target int, start, end int) int {
	if start == end {
		if target == nums[start] {
			return start
		}

		if target < nums[start] {
			return start
		}

		start++
		return start
	}

	mid := (start + end) / 2
	if target == nums[mid] {
		return mid
	}

	if target < nums[mid] {
		return searchInsert1(nums, target, start, mid)
	}

	return searchInsert1(nums, target, mid+1, end)
}
