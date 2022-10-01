package simple

import (
	"log"
	"math/bits"
)

// 二进制位不同的位置的数目
// ^异或: 不同为1 相同为0
// & 11为1
//bits.OnesCount 计算二进制位数为1的个数,通常用来检测错误
//https://leetcode.cn/problems/hamming-distance/solution/yi-ming-ju-chi-by-leetcode-solution-u1w7/
func (*Ref) HD() {
	x, y := 1, 4
	xor := x ^ y

	log.Printf("%v", bits.OnesCount(uint(xor)))
}

//颠倒给定的 32 位无符号整数的二进制位

//x = 1101
//y = 1011 ?
func (*Ref) Testxxx() {
	x := 43261596
	s := 0
	for i := x; i >= 0; i-- {
		//1
		if 1<<i&x > 0 {
			s += 1 << i
		}
	}

	log.Printf("%v", s)
}
