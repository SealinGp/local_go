package simple

import (
	"strings"
)

func (*Ref) ReverseWords() {
	reverseWords("the sky is blue")
}

func reverseWords(s string) string {
	rw := make([]string, 0)
	sl := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if len(sl) > 0 {
				rw = append(rw, string(sl))
				sl = sl[:0]
			}
			continue
		}

		sl = append(sl, s[i])
	}

	if len(sl) > 0 {
		rw = append(rw, string(sl))
	}

	for i := 0; i < len(rw)/2; i++ {
		j := len(rw) - i - 1
		rw[i], rw[j] = rw[j], rw[i]
	}

	return strings.Join(rw, " ")
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	isPa := func(substr string, i, j int) bool {
		for i < j {
			if substr[i] != substr[j] {
				return false
			}
			i++
			j--
		}

		return true
	}

	begin := 0
	maxLen := 1
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if j-i+1 > maxLen && isPa(s, i, j) {
				maxLen = j - i + 1
				begin = i
			}
		}
	}

	return s[begin : begin+maxLen]
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
