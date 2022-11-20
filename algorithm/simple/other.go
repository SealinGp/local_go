package simple

import (
	"log"
	"math"
	"math/bits"
)

//i份  n> i >  0
//n长度 n > 1

//n=2 i=1 f=1x1=1
//n=3 i=1,2 f=max(1x2,1x1x1)

func hammingWeight1(num uint32) int {
	count := 0
	var flag uint32 = 1
	for flag > 0 {
		if num&flag > 1 {
			count++
		}

		flag = flag << 1
	}

	return count
}

//贪婪算法
func cuttingRope1(n int) int {
	if n < 2 {
		return 0
	}

	if n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	//尽可能剪长度为3的绳子段
	timesOf3 := n / 3

	if n-timesOf3*3 == 1 {
		timesOf3 -= 1
	}

	timesOf2 := (n - timesOf3*3) / 2

	return int(math.Pow(3, float64(timesOf3))) * int(math.Pow(2, float64(timesOf2)))
}

//动态规划
func cuttingRope(n int) int {
	if n < 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	fn := make([]int, n+12)
	fn[1] = 1
	fn[2] = 2
	fn[3] = 3

	maxF := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	for i := 4; i <= n; i++ {
		max := 0

		for j := 1; j <= i/2; j++ {
			p := fn[j] * fn[i-j]
			max = maxF(max, p)
			fn[i] = max
		}
	}

	return fn[n]
}

//00000010100101000001111010011100
//?
//00111001011110000010100101000000
func reverseBits(num uint32) uint32 {
	return bits.Reverse32(num)
}

//numRows = 3

//	1
// 1 1
//1 2 1
// (2,1) = (1,0) + (1,1)

func (*Ref) Generate() {
	log.Printf("%v", generate(5))
}

func generate(numRows int) [][]int {
	m := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, 0, i)

		for j := 0; j <= i; j++ {
			value := 1
			if i-1 >= 0 && j-1 >= 0 && j <= i-1 {
				value = m[i-1][j-1] + m[i-1][j]
			}
			row = append(row, value)
		}

		m[i] = row
	}

	return m
}
func (*Ref) IsValid() {
	log.Printf("%v", isValid("([)]"))
}

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, c := range s {
		//left
		if r, ok := m[c]; ok {
			stack = append(stack, r)
		} else {
			//right
			if len(stack) == 0 {
				return false
			}

			if stack[len(stack)-1] != c {
				return false
			}

			stack = stack[:len(stack)-1]
		}

	}

	if len(stack) == 0 {
		return true
	}

	return false
}

func (*Ref) MissingNumber() {
	log.Printf("%v", missingNumber([]int{0, 1}))

}

func missingNumber(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for i := 0; i <= len(nums); i++ {
		m[i] = struct{}{}
	}

	log.Printf("%v", m)

	for _, num := range nums {
		if _, ok := m[num]; ok {
			delete(m, num)
		}
	}

	for missingNum := range m {
		return missingNum
	}

	return 0
}
