package main

func main() {}

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dst := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < dst {
			return nums[i]
		}

	}

	return dst
}

func getRow(rowIndex int) []int {
	rows := generate(rowIndex + 1)
	return rows[rowIndex]
}

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

func reverseWords(s string) string {
	arr := make([]byte, 0)
	newStr := make([]byte, 0, len(s))
	for _, v := range s {
		switch v {
		case ' ':
			newStr = append(newStr, ' ')
			newStr = append(newStr, reverseArr(arr)...)
			arr = arr[:0]
		default:
			arr = append(arr, byte(v))
		}
	}

	if len(arr) > 0 {
		newStr = append(newStr, ' ')
		newStr = append(newStr, reverseArr(arr)...)
	}

	return string(newStr[1:])
}

func reverseArr(arr []byte) []byte {
	if len(arr) == 0 {
		return arr
	}

	i := 0
	end := len(arr) - 1
	for i < end {
		arr[i], arr[end] = arr[end], arr[i]
		i++
		end--
	}

	return arr
}
