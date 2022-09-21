package simple

import (
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

//[1]     第1行1个元素 a[0][0]=1
//[1,1]   第2行2个元素 a[1][0]=1 a[1][1]=1
//[1,2,1] 第3行3个元素 a[2][0]=1 a[2][1] = a[1][0]+a[1][1]= 2 a[2][2]=1
//[...]   第n行n个元素 a[n-1][0]=1 a[n-1][1] = a[n-1][1-1] + ... a[n-1][1]

func generate(numRows int) [][]int {
	m := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, 0, i)

		for j := 0; j < i; j++ {

			value := 0

			row = append(row, value)
		}
		m[i] = row
	}

}
