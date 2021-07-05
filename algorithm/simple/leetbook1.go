package simple

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
func (*Ref) Nums() {
	nums := []int{0, 1, 1, 2, 2}

	len1 := removeDuplicates(nums)

	// 在函数里修改输入数组对于调用者是可见的。
	// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
	for i := 0; i < len1; i++ {
		fmt.Println(nums[i])
	}
}
func removeDuplicates(nums []int) int {
	numsL := len(nums)
	if numsL < 1 {
		return numsL
	}
	left := 0
	right := left + 1
	for right < numsL && left < numsL {
		if nums[left] == nums[right] {
			right++
		} else {
			if right != left+1 {
				nums[left+1], nums[right] = nums[right], nums[left+1]
			}
			left++
			right++
		}
	}
	return left + 1
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2zsx1/
func (*Ref) Mp2() {

}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x21ib6/
func (*Ref) SN() {
	nums := []int{2, 2, 1}

	s := 0
	for i := range nums {
		s = s ^ nums[i]
	}
	fmt.Println(s)
}

//https://leetcode-cn.com/problems/combination-sum/
func (*Ref) S1() {
	candidates := []int{2, 3, 5}
	target := 8

	fmt.Println(combinationSum(candidates, target))
}
func combinationSum(candidates []int, target int) [][]int {
	current := list.New()
	results := [][]int{}
	dfs1(candidates, 0, current, target, func(a []int) {
		results = append(results, a)
	})
	return results
}

//深度优先搜索
func dfs1(candidates []int, index int, current *list.List, remainTarget int, fun func(a []int)) {
	if remainTarget < 0 {
		return
	}
	if remainTarget == 0 {
		f := current.Front()
		arr := []int{}
		for f != nil {
			arr = append(arr, f.Value.(int))
			f = f.Next()
		}
		fun(arr)
		return
	}

	for i := index; i < len(candidates); i++ {
		current.PushBack(candidates[i])
		dfs1(candidates, i, current, remainTarget-candidates[i], fun)
		current.Remove(current.Back())
	}
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2y0c2/
func (*Ref) Intersect() {
	fmt.Println(intersect([]int{7, 2, 2, 4, 7, 0, 3, 4, 5}, []int{3, 9, 8, 6, 1, 9}))
	fmt.Println(intersect1([]int{7, 2, 2, 4, 7, 0, 3, 4, 5}, []int{3, 9, 8, 6, 1, 9}))
}
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	nums1L := len(nums1)
	nums2L := len(nums2)

	arr := []int{}
	left := 0
	right := 0
	for left < nums1L && right < nums2L {
		for left < nums1L && right < nums2L && nums1[left] < nums2[right] {
			left++
		}
		for left < nums1L && right < nums2L && nums2[right] < nums1[left] {
			right++
		}

		if left < nums1L && right < nums2L && nums1[left] == nums2[right] {
			arr = append(arr, nums1[left])
			left++
			right++
		}
	}
	return arr
}
func intersect1(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	tmp := make(map[int]int)
	for _, v := range nums1 {
		tmp[v]++
	}
	arr := []int{}
	for _, v := range nums2 {
		if _, ok := tmp[v]; ok {
			arr = append(arr, v)
			tmp[v]--
			if tmp[v] <= 0 {
				delete(tmp, v)
			}
		}
	}
	return arr
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2cv1c/
func (*Ref) PlusOne() {
	fmt.Println(plusOne([]int{9}))
}
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			break
		}
		digits[i] = digits[i] % 10
	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
func (*Ref) MoveZeros() {
	num := []int{0, 1, 0, 3, 12}
	MoveZeros(num)
	fmt.Println(num)
}
func MoveZeros(nums []int) {
	tmp := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[tmp] = nums[i]
			tmp++
		}
	}
	numsL := len(nums)
	for i := tmp; i < numsL; i++ {
		nums[i] = 0
	}
}

func (*Ref) TwoSum1() {
	fmt.Println(twoSum1([]int{2, 7, 11, 15}, 9))
}
func twoSum1(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i := range nums {
		numsMap[nums[i]] = i
	}
	for i := range nums {
		tmp := target - nums[i]
		if i1, ok := numsMap[tmp]; ok && i1 != i {
			return []int{i, i1}
		}
	}
	return nil
}

func (*Ref) IsValidSudo() {

}
func isValidSudo(board [][]byte) bool {
	rowsMap := make([][]bool, 9)
	columnMap := make([][]bool, 9)
	boxes := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowsMap[i] = make([]bool, 9)
		columnMap[i] = make([]bool, 9)
		boxes[i] = make([]bool, 9)
	}

	for row := range board {
		for column := range board[row] {
			if board[row][column] == '.' {
				continue
			}
			val := board[row][column] - '0' - 1

			boxIndex := (row/3)*3 + column/3
			if rowsMap[row][val] || columnMap[column][val] || boxes[boxIndex][val] {
				return false
			}
			rowsMap[row][val] = true
			columnMap[column][val] = true
			boxes[boxIndex][val] = true
		}
	}
	return true
}

func (*Ref) Rotate1() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
func rotate(matrix [][]int) {
	n := len(matrix)

	for i := range matrix {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			j1 := n - j - 1
			matrix[i][j], matrix[i][j1] = matrix[i][j1], matrix[i][j]
		}
	}
}

func (*Ref) ReverseStr() {
	s := []byte{'h', 'e', 'l', 'l', 'o', 'a'}
	reverseString(s)
	fmt.Println(string(s))
}
func reverseString(s []byte) {
	reverStr(s, 0, len(s)-1)
}

func reverStr(s []byte, left, right int) {
	if left >= right {
		return
	}
	s[left], s[right] = s[right], s[left]
	reverStr(s, left+1, right-1)
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnx13t/
func (*Ref) Reverse1() {
	reverse(-123)
}
func reverse(x int) int {
	var x1 int
	for x != 0 {
		pop := x % 10
		x /= 10

		if x1 > math.MaxInt32/10 || (x1 == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if x1 < math.MinInt32/10 || (x1 == math.MinInt32/10 && pop < -8) {
			return 0
		}
		x1 = x1*10 + pop
	}

	return x1
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn5z8r/
func (*Ref) FirstUniqChar() {
	fmt.Println(firstUniqChar("leetcode"))
}
func firstUniqChar(s string) int {
	c := make([]uint, 26, 26)
	for _, c1 := range s {
		c[c1-'a']++
	}

	for i, c1 := range s {
		if c[c1-'a'] == 1 {
			return i
		}
	}
	return -1
}

//https://leetcode-cn.com/problems/valid-anagram/
func (*Ref) IsAnagram() {
	fmt.Println(isAnagram("rat", "car"))
}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	tmp := make([]int, 26, 26)
	for _, v := range s {
		tmp[v-'a']++
	}
	for _, v := range t {
		tmp[v-'a']--
		if tmp[v-'a'] < 0 {
			return false
		}
	}
	return true
}

func (*Ref) IPD() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func (*Ref) MyAtoi() {
	fmt.Println(myAtoi(" +0 123"))
}
func myAtoi(str string) int {
	str = strings.TrimSpace(str)

	var strI int
	var negative bool

	if len(str) < 1 {
		return strI
	}

	//确认正负
	switch str[0] {
	case '+':
		str = str[1:]
	case '-':
		negative = true
		str = str[1:]
	default:
		if str[0] < '0' || str[0] > '9' {
			return 0
		}
	}

	for i := range str {
		if v, ok := isNum(str[i]); ok {
			if negative {
				v = -v
			}
			if strI > math.MaxInt32/10 || (strI == math.MaxInt32/10 && v > 7) {
				return math.MaxInt32
			}
			if strI < math.MinInt32/10 || (strI == math.MinInt32/10 && v < -8) {
				return math.MinInt32
			}

			strI = strI*10 + v
		} else {
			return strI
		}
	}

	return strI
}

func isNum(rune2 uint8) (int, bool) {
	if rune2 < '0' || rune2 > '9' {
		return 0, false
	}
	return int(rune2 - '0'), true
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnr003/
func (*Ref) Strstr() {
	fmt.Println(strStr("a", "a"))
}
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	findI := -1
	nL := len(needle)
	hL := len(haystack)
	for i := range haystack {
		if i+nL <= hL && haystack[i:i+nL] == needle {
			return i
		}
	}
	return findI
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnpvdm/
func (*Ref) CountAndSay() {
	fmt.Println(countAndSay(6))
}

func countAndSay(n int) string {
	if n <= 1 {
		return "1"
	}
	pre := countAndSay(n - 1)
	i, j := 0, 0
	newStr := ""
	preL := len(pre)
	for i < preL && j < preL {
		for j < preL && pre[i] == pre[j] {
			j++
		}

		if i < preL {
			L := j - i
			v := pre[i]
			newStr += strconv.Itoa(L) + string(v)
		}

		i = j
	}
	return newStr
}

func (*Ref) LCP() {
	strs := []string{"caa", "", "a", "acb"}
	fmt.Println(longestCommonPrefix(strs))
}
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	comm := strs[0]
	for i := range strs {
		comm = findCommon(comm, strs[i])
	}
	return comm
}
func findCommon(str1, str2 string) string {
	minL := min(len(str1), len(str2))
	for i := 0; i < minL; i++ {
		if str1[i] != str2[i] {
			return str1[:i]
		}
	}
	return str1[:minL]
}

func (*Ref) DN() {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}
	root := deleteNode(head, 5)
	for root != nil {
		fmt.Println(root.Val)
		root = root.Next
	}
}
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}

	curr := head
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}
	return head
}

func (*Ref) RNFE() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	root := removeNthFromEnd(head, 1)
	for root != nil {
		fmt.Println(root.Val)
		root = root.Next
	}
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first := head
	cur := head
	for first != nil && n+1 > 0 {
		first = first.Next
		n--
	}

	for first != nil {
		cur = cur.Next
		first = first.Next
	}

	if cur == head && n >= 0 {
		head = head.Next
	} else {
		cur.Next = cur.Next.Next
	}
	return head
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnnhm6/
func (*Ref) RL() {
}
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	left := head
	right := head.Next

	for right != nil {
		if left == head {
			left.Next = nil
		}
		tmp := right.Next
		right.Next = left
		left = right
		right = tmp
	}
	return left
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnnbp2/
//双指针
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newN := &ListNode{
		Val:  -1,
		Next: nil,
	}
	head := newN
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				newN.Next = l1
				l1 = l1.Next
			} else {
				newN.Next = l2
				l2 = l2.Next
			}
		} else if l1 != nil {
			newN.Next = l1
			l1 = l1.Next
		} else if l2 != nil {
			newN.Next = l2
			l2 = l2.Next
		} else {
			break
		}
		newN = newN.Next
	}
	return head.Next
}

//递归
func mtl(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l1.Next = mtl(l1.Next, l2)
			return l1
		} else {
			l2.Next = mtl(l1, l2.Next)
			return l2
		}
	}

	if l1 != nil {
		return l1
	}
	return l2
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnv1oc/
func isPalindromeL(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	cur := head

	newHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	newHeadCur := newHead
	for cur != nil {
		newHeadCur.Next = &ListNode{
			Val:  cur.Val,
			Next: nil,
		}
		cur = cur.Next
		newHeadCur = newHeadCur.Next
	}

	newHead = reverseList(newHead.Next)
	for newHead != nil && head != nil {
		if newHead.Val != head.Val {
			return false
		}
		newHead = newHead.Next
		head = head.Next
	}
	return true
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnwzei/
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && slow != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next

		if fast == slow {
			return true
		}
	}
	return false
}
