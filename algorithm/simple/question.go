package simple

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。不能重复利用这个数组中同样的元素。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
*/

/*
方法1 暴力解法

时间复杂度:O(n)
5个元素的nums,查找次数为
len = 5;
O(5) = (0+1)*(5-1)+(1+1)*(5-2)+(2+1)*(5-3)+(3+1)*(5-4)

n个元素的nums,查找次数为
len = n;
O(n) = (index+1)*(n-1);
当n趋近于无穷大时
O(n) = n + n +.... n;
O(n) = n;
*/
func (*Ref) TwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9

	numIndex := make([]int, 0)
	valueLeft := 0

	for index, value := range nums {
		valueLeft = target - value
		sli := nums[index+1:]

		index2, find := inSlice(valueLeft, sli)
		if find {
			numIndex = append(numIndex, index, index2+(index+1))
		}
	}

	fmt.Println(numIndex)
}

/**
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转
https://leetcode-cn.com/problems/reverse-integer/
*/
func (*Ref) Reverse() {
	//转string
	x := rand.New(rand.NewSource(time.Now().UnixNano())).Int31()
	if x == 0 {
		log.Fatal(x)
	}

	negative := false
	xStr := strconv.Itoa(int(x))
	if x < 0 {
		negative = true
		xStr = xStr[1:]
	}
	xStrLen := len(xStr)
	var xByte []byte

	//string反转
	for i := xStrLen - 1; i >= 0; i-- {
		if i == xStrLen-1 && xStr[i] == 48 {
			continue
		}
		xByte = append(xByte, xStr[i])
	}
	xStr = string(xByte)

	//转int
	result, Err := strconv.Atoi(xStr)
	if Err != nil {
		log.Fatal(Err.Error())
	}
	if negative {
		result = -result
	}
	if result < -2147483648 || result > 2147483647 {
		result = 0
	}
	fmt.Println(x)
	fmt.Println(result)
}
func (*Ref) Reverse2() {
	//转string
	//x  := rand.New(rand.NewSource(time.Now().UnixNano())).Int31()
	x := -120

	negative := false
	if x < 0 {
		negative = true
		x = -x
	}
	result := 0
	xStr := strconv.Itoa(x)
	xLen := len(xStr)
	intMax := int(math.Pow(2, 31))
	xMax := intMax - 1
	xMin := -intMax
	for i := xLen; i > 0; i-- {
		result += (x % 10) * int(math.Pow10(i-1))
		x /= 10
	}

	if negative {
		result = -result
	}
	if result > xMax || result < xMin {
		result = 0
	}
	fmt.Println(result)
}

/**
https://leetcode-cn.com/problems/palindrome-number/
*/
func (*Ref) IsPalindrome() {
	x := 12121

	isPalindrome := false
	xReverse := 0
	xReverseS := []int{}
	xLeft := x
	for {
		remainder := xLeft % 10
		xReverseS = append(xReverseS, remainder)
		if xLeft = xLeft / 10; xLeft == 0 {
			break
		}
	}

	xReverseSLen := len(xReverseS)
	for i, v := range xReverseS {
		xReverse += v * int(math.Pow10(xReverseSLen-i-1))
	}

	isPalindrome = x == xReverse
	fmt.Println(isPalindrome)
}
func (*Ref) IsPalindrome1() {
	x := 1212
	if x < 0 {
		log.Fatal(false)
	}

	tmp := x
	new := 0
	for tmp != 0 {
		new = new*10 + tmp%10
		tmp /= 10
	}
	fmt.Println(new == x)
}

func (*Ref) RomanToInt() {
	s := "III"
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	spe := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	I := 0
	for k, v := range spe {
		if num := strings.Count(s, k); num != 0 {
			s = strings.Replace(s, k, "", -1)
			I += v * num
		}
	}

	fmt.Println(s)
	for k, v := range roman {
		if num := strings.Count(s, k); num != 0 {
			I += v * num
		}
	}
	fmt.Println(I)
}
func (*Ref) RomanToInt2() {
	s := "MCMXCIV"
	roman := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sLen := len(s)
	I := 0
	for i, v := range s {
		if roman[v] == 0 {
			continue
		}
		if i+1 < sLen && roman[v] < roman[rune(s[i+1])] {
			I -= roman[v]
		} else {
			I += roman[v]
		}
	}
	fmt.Println(I)
}
func (*Ref) LongestCommonPrefix() {
	strs := []string{"aa", "a"}
	strsL := len(strs)
	res := ""
	x := 0
	hasCommon := false

	if strsL == 0 {
		fmt.Println(res)
		return
	}
	if strsL == 1 {
		res = strs[0]
		fmt.Println(res)
		return
	}

f:
	for {
		for i := 0; i < strsL; i++ {
			//["","abc"]
			if strs[i] == "" {
				hasCommon = false
				break f
			}

			//["a","a"] ["aa","a"]
			if x == len(strs[i]) || (i+1 != strsL && x == len(strs[i+1])) {
				break f
			}

			if i+1 != strsL && strs[i][x] != strs[i+1][x] {
				break f
			}
			hasCommon = true
		}
		x++
	}

	if hasCommon {
		res = strs[0][:x]
	}
	fmt.Println(res)
}

func (*Ref) LongestCommonPrefix1() {
	strs := []string{"a", "ab", "ac"}
	res := ""
	defer func() {
		fmt.Println(res)
	}()

	if len(strs) == 0 {
		return
	}
	res = strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], res) {
			res = res[:len(res)-1]

			if len(res) == 0 {
				res = ""
				return
			}
		}
	}
}

func (*Ref) IsValid() {
	//1 2 5
	s := "[(({})}]"
	res := false
	defer func() {
		fmt.Println(res)
	}()

	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	validNum := 0
	for i, v := range s {
		if vRight, ok := m[v]; ok {
			for j := i + 1; j < len(s); j++ {
				if rune(s[j]) == vRight && (j-i)%2 != 0 {
					validNum++
					//不重复寻找匹配的
					break
				}
			}
		}
	}

	res = validNum*2 == len(s)
}

func (*Ref) Check() {

}

// https://leetcode-cn.com/problems/add-two-numbers
type ListNode struct {
	Val  int
	Next *ListNode
}

func (*Ref) AddTwoNumbers() {
	// l1,l2 : 输入节点
	// l3    : l1+l2 的节点
	// l3P   : l3节点指针
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l3 := &ListNode{}
	l3P := l3
	nextPlus := 0

	for l1 != nil || l2 != nil || nextPlus > 0 {
		val := 0

		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		val += nextPlus
		nextPlus = val / 10
		val %= 10

		l3P.Next = &ListNode{Val: val}
		l3P = l3P.Next
	}

	fmt.Println(l3.Next)
}
func (*Ref) AddTwoNumbers1() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l3P := &ListNode{}
	l3 := l3P
	nextPlus := 0

	for l1 != nil || l2 != nil || nextPlus > 0 {
		var val = 0
		if l1 != nil && l2 != nil {
			val += l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		} else if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		val += nextPlus

		//进多少位
		nextPlus = val / 10
		//取余
		val %= 10

		l3P.Next = &ListNode{Val: val, Next: nil}
		l3P = l3P.Next
	}

	fmt.Println(l3.Next)
}

func (*Ref) MaxSubString() {
	s := "pwwkew"
	resm := make(map[rune]int)
	maxL := 0
	for i, sr := range s {
		resm[sr]++

		//存在重复的
		if resm[sr] > 1 {

			if resl := len(resm); maxL < resl {
				maxL = resl
			}
			resm = make(map[rune]int)

			resm[sr]++
			sr1 := s[:i]
			sr1Index := len(sr1) - 1
			for sr1Index >= 0 {
				if int32(sr1[sr1Index]) == sr {
					break
				}
				resm[int32(sr1[sr1Index])]++
				sr1Index--
			}
		}
	}

	if resl := len(resm); maxL < resl {
		maxL = resl
	}

	fmt.Println(maxL)
}
func (*Ref) MaxSubString2() {
	s := "bbtablud"
	maxLen, left := 0, 0
	for right := range s {
		sCur := s[left:right]
		index := strings.IndexByte(sCur, s[right])
		if index != -1 {
			left += index + 1
		} else {
			le := right - left + 1
			if maxLen < le {
				maxLen = le
			}
		}
	}
	fmt.Println(maxLen)
}

//https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
func (*Ref) MidSortedArrays() {

}

// https://leetcode-cn.com/problems/the-masseuse-lcci/
func (*Ref) Message() {
	nums := []int{2, 1, 4, 5, 3, 1, 1, 3}
	numsL := len(nums)
	if numsL == 0 {
		log.Println(0)
		return
	}
	/**
	dp[i][0] : 前i个预约, 第i个预约 不接的最长预约时间
	dp[i][1] : 前i个预约, 第i个预约 接的最长预约时间
	nums : 第i个预约的预约时长

	从前往后计算dp值, 假设已经计算出前i-1个dp值 考虑计算dp[i][0] dp[i][1]
	dp[i][0] = max(dp[i-1][0],dp[i-1][1])
	dp[i][1] = dp[i-1][0] + nums
	*/
	//{2,1,4,5,3,1,1,3}
	dp0 := 0       //不接的
	dp1 := nums[0] //接的
	for i := 1; i < numsL; i++ {
		tdp0 := math.Max(float64(dp0), float64(dp1))
		tdp1 := dp0 + nums[i]
		dp0 = int(tdp0)
		dp1 = tdp1
	}
	log.Println(math.Max(float64(dp0), float64(dp1)))
}

//https://leetcode-cn.com/problems/house-robber/
func (*Ref) Rob() {
	nums := []int{2, 1, 1, 2, 3}

	n := len(nums)
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	if n == 0 {
		fmt.Println(0)
		return
	}
	if n == 1 {
		fmt.Println(nums[0])
		return
	}
	if n == 2 {
		fmt.Println(max(nums[0], nums[1]))
		return
	}

	//i第i间房屋
	//dp[i] 前i间房屋能偷到的最大金额
	dp := map[int]int{}
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	tp1 := nums[0]
	tp2 := max(nums[0], nums[1])
	for i := range nums {
		if i < 2 {
			continue
		}
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])

		//tp1,tp2 = tp2,max(tp1 + nums[i],tp2)
		tmp := tp1
		tp1 = tp2
		tp2 = max(tmp+nums[i], tp2)
	}

	fmt.Println(dp[n-1], tp2)
}
