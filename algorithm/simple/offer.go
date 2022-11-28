package simple

import (
	"log"
)

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var pHead *ListNode

	if l1.Val < l2.Val {
		pHead = l1
		l1.Next = mergeTwoLists1(l1.Next, l2)
	} else {
		pHead = l2
		l2.Next = mergeTwoLists1(l1, l2.Next)
	}

	return pHead
}

// func m1(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	}

// 	if l2 == nil {
// 		return l1
// 	}

// 	var pHead *ListNode

// 	if l1.Val < l2.Val {
// 		pHead = l1
// 		l1.Next =
// 	}

// 	return pHead
// }

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pReverseHead *ListNode
	var pPreNode *ListNode
	pNode := head
	for pNode != nil {
		pNext := pNode.Next
		if pNext == nil {
			pReverseHead = pNode
		}

		pNode.Next = pPreNode
		pPreNode = pNode
		pNode = pNext
	}

	return pReverseHead
}

func detectCycle(head *ListNode) *ListNode {
	meetingNode := MeetingNode(head)
	if meetingNode == nil {
		return nil
	}

	ringLen := 1
	pNode1 := meetingNode
	for pNode1.Next != meetingNode {
		pNode1 = pNode1.Next
		ringLen++
	}

	pNode1 = head
	for i := 0; i < ringLen; i++ {
		pNode1 = pNode1.Next
	}

	pNode2 := head
	for pNode1 != pNode2 {
		pNode1 = pNode1.Next
		pNode2 = pNode2.Next
	}

	return pNode1
}

func MeetingNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head.Next
	if slow == nil {
		return nil
	}

	fast := slow.Next
	for slow != nil && fast != nil {
		if slow == fast {
			return fast
		}

		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	return nil
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return nil
	}

	fast := head
	for x := 0; x < k-1; x++ {
		if fast.Next == nil {
			return nil
		}

		fast = fast.Next
	}

	slow := head
	for {
		if fast.Next == nil {
			break
		}

		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	isOdd := func(a int) bool {
		return a&0x1 != 0
	}

	i := 0
	j := len(nums) - 1
	for i < j {

		if isOdd(nums[i]) {
			i++
		}

		if !isOdd(nums[j]) {
			j--
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return nums
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	return nil
}

func printNumbers(n int) []int {
	count := 1
	for i := 0; i < n; i++ {
		count = count * 10
	}

	arr := make([]int, count-1, count)
	for i := 1; i < count; i++ {
		arr[i-1] = i
	}

	return arr
}

//1.不可重复访问
//2.返回访问
//3.可访问条件

func (*Ref) MovingCount() {
	log.Printf("%v", movingCount(2, 3, 1))
}

//回溯法
func movingCount(m int, n int, k int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	isMoved := make([][]bool, m)
	for i := range isMoved {
		isMoved[i] = make([]bool, n)
	}

	return movingCountCore(0, m, 0, n, k, isMoved)
}

func movingCountCore(row, rows, col, cols, k int, isMoved [][]bool) int {
	count := 0
	if row >= 0 && row < rows && col >= 0 && col < cols && !isMoved[row][col] && shuweihe1(row, col) <= k {
		isMoved[row][col] = true
		count += 1

		count += movingCountCore(row, rows, col-1, cols, k, isMoved)
		count += movingCountCore(row-1, rows, col, cols, k, isMoved)
		count += movingCountCore(row, rows, col+1, cols, k, isMoved)
		count += movingCountCore(row+1, rows, col, cols, k, isMoved)
	}

	return count
}

func shuweihe1(row, col int) int {
	return shuweihe(row) + shuweihe(col)
}

func shuweihe(num int) int {
	sum := 0

	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return sum
}

func (*Ref) Exist() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	log.Printf("%v", exist(board, word))
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	rows := len(board)
	cols := len(board[0])
	visited := make([][]bool, rows)
	for row := range visited {
		visited[row] = make([]bool, cols)
	}

	pathLen := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if hasPathCore(board, row, rows, col, cols, word, pathLen, visited) {
				return true
			}
		}
	}

	return false
}

func hasPathCore(matrix [][]byte, row, rows, col, cols int, str string, pathLen int, visited [][]bool) bool {
	if pathLen == len(str) {
		return true
	}

	hasPath := false

	if row >= 0 && row < rows && col >= 0 && col < cols && matrix[row][col] == str[pathLen] && !visited[row][col] {
		pathLen++
		visited[row][col] = true

		hasPath = hasPathCore(matrix, row, rows, col-1, cols, str, pathLen, visited) || hasPathCore(matrix, row-1, rows, col, cols, str, pathLen, visited) ||
			hasPathCore(matrix, row, rows, col+1, cols, str, pathLen, visited) || hasPathCore(matrix, row+1, rows, col, cols, str, pathLen, visited)

		if !hasPath {
			pathLen--
			visited[row][col] = false
		}
	}

	return hasPath
}

func numWays(n int) int {
	if n <= 2 {
		return n
	}

	f1 := 1
	f2 := 2
	for i := 2; i <= n; i++ {
		current := (f1 + f2) % mod
		f1 = f2
		f2 = current
	}

	return f2
}

func (*Ref) Fib() {
	log.Printf("%v", fib(45))
	log.Printf("134903163")
}

const mod int = 1e9 + 7

func fib(n int) int {
	if n <= 1 {
		return n
	}

	j1 := 1
	j2 := 0
	for i := 2; i <= n; i++ {
		current := (j1 + j2) % mod
		j2 = j1
		j1 = current
	}

	return j1
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	arr := make([]int, 0)
	for current := head; current != nil; current = current.Next {
		arr = append(arr, current.Val)
	}

	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func (*Ref) FindRepeatNumber() {
	log.Printf("%v", findRepeatNumber([]int{3, 4, 2, 1, 1, 0}))
}

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}

			v := nums[i]
			nums[i], nums[v] = nums[v], nums[i]
		}
	}

	return 0
}
