package simple

import (
	"fmt"
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