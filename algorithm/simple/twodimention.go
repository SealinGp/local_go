package simple

import "log"

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
