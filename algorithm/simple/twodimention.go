package simple

import "log"

//https://leetcode.cn/leetbook/read/array-and-string/cuxq3/

// m x n
// 3 x 3

// 1 2 3
// 4 5 6
// 7 8 9

// 3x3

//0 m[0][0]
//1 m[0][1],m[1][0]
//2 m[2][0],m[0][2]

//3 m[1][2],m[2][1]
//4 m[2][2]


//0,1,2
//0,1,1
//a[2]=2
//a[2]=1
//a[1]=1

//奇数=up
//偶数=down
func (*Ref) FindDiagonalOrder() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	log.Printf("%v", findDiagonalOrder(mat))
}

func findDiagonalOrder(mat [][]int) []int {
	m := make([]int, 0, len(mat))
	if len(mat) <= 0 {
		return m
	}

	x := len(mat) - 1
	y := len(mat[0]) - 1
	line := len(mat) + len(mat[0]) - 1
	r, c := 0, 0

	for i := 0; i < line; i++ {

		down := true
		if i > 0 {
			down = i%2 == 0
		}

		if down {
			for {
				if i == 2 {
					log.Printf("r:%v, c:%v", r, c)
				}

				if r < 0 || c > y {
					r++
					c--
					break
				}

				m = append(m, mat[r][c])
				r--
				c++
			}

			if c+1 <= y {
				c++
			} else {
				r++
			}
			continue
		}

		for {
			if r > x || c < 0 {
				r--
				c++
				break
			}

			m = append(m, mat[r][c])
			r++
			c--
		}

		if r+1 <= x {
			r++
		} else {
			c++
		}
	}

	return m
}

//编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。

// {0, 1, 2, 0},
// {3, 4, 5, 2},
// {1, 3, 1, 5},

//3x4
//a[0][0] = 0

// col
// a[0][0]...a[2][0]

// row
// a[0][0]...a[0][3]
func (*Ref) SetZeroes() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)
	log.Printf("%v", matrix)
}

func setZeroes(matrix [][]int) {
	x := len(matrix)
	if x == 0 {
		return
	}

	tmp := make([][]int, len(matrix))
	for i := range matrix {
		tmp[i] = make([]int, len(matrix[i]))
		copy(tmp[i], matrix[i])
	}

	y := 0

	for i := 0; i < x; i++ {
		if i == 0 {
			y = len(tmp[i])
		}

		for j := 0; j < y; j++ {
			if tmp[i][j] == 0 {
				//col
				for i1 := 0; i1 < x; i1++ {
					matrix[i1][j] = 0
				}

				//row
				for j1 := 0; j1 < y; j1++ {
					matrix[i][j1] = 0
				}
			}

		}
	}

	return
}

func (*Ref) Rotate90() {
	rotate1(
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	)
}

/**

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]

NxN = 3x3

对于矩阵中第 i 行的第 j 个元素，在旋转后，它出现在倒数第 i 列的第 j 个位置。
m[i][j] = m[j][y-i-1]
*/
func rotate1(matrix [][]int) {
	if len(matrix) <= 0 {
		return
	}

	x := len(matrix)
	y := 0
	matrixTmp := make([][]int, len(matrix))
	for i := 0; i < x; i++ {
		y = len(matrix[i])
		if x != y {
			return
		}
		matrixTmp[i] = make([]int, y)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			matrixTmp[j][y-i-1] = matrix[i][j]
		}
	}

	copy(matrix, matrixTmp)
}
