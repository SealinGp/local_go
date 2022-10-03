package simple

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
