package simple

// [1, 7, 3, 6, 5, 6]

// 1.sum[left] -sum[right] = 0
// sum[left] = total - numsi - sum[left] => 2*sumLeft = total - numsi
func PivotIndex(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}

	sum := 0
	for i, n := range nums {
		if 2*sum == total-n {
			return i
		}
		sum += n
	}

	return -1
}
