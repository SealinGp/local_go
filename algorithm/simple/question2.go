package simple

import (
	"fmt"
	"strconv"
	"strings"
)

//https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/
func (*Ref) SubArrDivByK() {
	fmt.Println(-6 % 4)
	A := []int{5}
	K := 9
	record := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, v := range A {
		sum += v
		modules := (sum%K + K) % K
		ans += record[modules]
		record[modules]++
	}
	fmt.Println(ans)
}

func (*Ref) DecodeString() {
	s := "2[abc]3[cd]ef"

	lastLeft := 0
F:
	for i, v := range s {
		if v == '[' {
			lastLeft = i
		}
		if v == ']' {
			//找出被重复字符串的次数
			numIndex := lastLeft - 1
			for {
				if numIndex < 0 {
					numIndex = 0
					break
				}
				if s[numIndex] < '0' || s[numIndex] > '9' {
					numIndex++
					break
				}
				numIndex--
			}

			//找出被重复字符串
			newV, _ := strconv.Atoi(s[numIndex:lastLeft])
			res := strings.Repeat(
				s[lastLeft+1:i],
				newV,
			)

			//字符串重新拼接
			s = s[:numIndex] + res + s[i+1:]
			lastLeft = 0
			goto F
		}
	}

	fmt.Println(s == "abcabccdcdcdef")
}

func (*Ref) ProductExceptSelf() {
	//n > 1
	nums := []int{1, 2, 3, 4}
	L := len(nums)
	a := make([]int, L)
	larr := make([]int, L)
	rarr := make([]int, L)
	larr[0] = 1
	rarr[L-1] = 1

	for i := 1; i < L; i++ {
		larr[i] = nums[i-1] * larr[i-1]
	}
	for i := L - 2; i >= 0; i-- {
		rarr[i] = nums[i+1] * rarr[i+1]
	}
	for i := range nums {
		a[i] = larr[i] * rarr[i]
	}
	fmt.Println(a)
}

func (*Ref) EquationsPossible() {
	equations := []string{
		"a==b", "b!=c", "c==a",
	}

	ma := make([]byte, 'z'-'a'+1)
	for i := range ma {
		ma[i] = byte(i)
	}
	for _, s := range equations {
		if s[1] == '=' {
			Union(s[0], s[3], ma)
		}
	}

	is := true
	for _, s := range equations {
		if s[1] == '!' {
			if findSet(s[0]-'a', ma) == findSet(s[3]-'a', ma) {
				is = false
				break
			}
		}
	}

	fmt.Println(is)
}

func makeSet(i byte, mab []byte) {
	index := i - 'a'
	mab[index] = index
}

func findSet(i byte, mab []byte) byte {
	if mab[i] == i {
		return i
	}
	return findSet(mab[i], mab)
}
func Union(i, j byte, mab []byte) {
	i = findSet(i-'a', mab)
	j = findSet(j-'a', mab)
	if i > j {
		mab[j] = i
	} else {
		mab[i] = j
	}
}

//https://leetcode-cn.com/problems/daily-temperatures/
func (*Ref) DT() {
	T := []int{
		73, 74, 75, 71, 69, 72, 76, 73,
	}
	Len := len(T)
	stack := make([]int, 0)
	T1 := make([]int, Len)

	for i := 0; i < Len; i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			T1[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}

}

//https://leetcode-cn.com/problems/1-bit-and-2-bit-characters/
func (*Ref) IBC() {
	bits := []int{1, 1, 1, 0}

	bitsLen := len(bits)
	is2Bit := true
	if len(bits) == 1 {
		is2Bit = bits[0] == 0
	}
	startIndex := 0

F:
	for i := startIndex; i < bitsLen-1; i++ {
		if i < bitsLen-2 {
			if bits[i] == 1 {
				startIndex = i + 2
				goto F
			}
		}

		if i == bitsLen-2 && bits[i] != 0 {
			is2Bit = false
		}
	}

	fmt.Println(is2Bit)
}
