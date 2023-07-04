package main

func main() {}

func generate(numRows int) [][]int {
	trangle := make([][]int, 0, numRows)
	for row := 0; row < numRows; row++ {
		maxCol := row + 1

		cols := make([]int, 0, maxCol)
		for col := 0; col < maxCol; col++ {
			colX := 1
			if row >= 1 && col >= 1 {
				if row-1 < len(trangle) && col < len(trangle[row-1]) && col-1 < len(trangle[row-1]) {
					colX = trangle[row-1][col] + trangle[row-1][col-1]
				}
			}

			cols = append(cols, colX)
		}

		trangle = append(trangle, cols)
	}

	return trangle
}
