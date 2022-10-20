package simple

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return ""
	}

	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		}
		return ""
	}

	//abc 3/2=1
	//ab 3/2=1
	
	//P(i,j)  s[i:j+1]
	//p(i,j) = p(i+1,j-1) && s[i] == s[j]

	isPa := func(substr string) bool {
		j := 

	}

}

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
