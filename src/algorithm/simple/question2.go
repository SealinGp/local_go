package simple

import (
	"fmt"
	"strconv"
	"strings"
)

//https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/
func (*Ref)SubArrDivByK()  {
	fmt.Println(-6%4)
	A      := []int{5}
	K      := 9
	record := map[int]int{0:1}
	sum, ans := 0, 0
	for _, v := range A {
		sum += v
		modules := (sum % K + K)%K
		ans += record[modules]
		record[modules]++
	}
	fmt.Println(ans)
}

func (*Ref)DecodeString()  {
	s        := "2[abc]3[cd]ef"

	lastLeft := 0
	F : for i, v := range s {
		if v == '[' {
			lastLeft = i
		}
		if v == ']' {
			//找出被重复字符串的次数
			numIndex := lastLeft-1
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
			newV,_   := strconv.Atoi(s[numIndex:lastLeft])
			res      := strings.Repeat(
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