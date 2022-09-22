package simple

import (
	"log"
	"math/bits"
)

func (*Ref) ReverseBits() {

	// log.Printf("%v", reverseBits())
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
