package main

import (
	"container/list"
	"fmt"
	"sort"
	"strconv"
)

var struct1Funcs = map[string]func(){
	"ds3":  ds3,
	"ds4":  ds4,
	"ds5":  ds5,
	"ds6":  ds6,
	"ds7":  ds7,
	"ds8":  ds8,
	"ds9":  ds9,
	"ds10": ds10,
	"ds11": ds11,
	"ds12": ds12,
	"ds13": ds13,
	"ds14": ds14,
	"ds15": ds15,
}

// 优先级队列的实现 = 二叉堆 + 队列 (逻辑结构) = 数组 (物理结构)
// 这里大顶堆为例
func ds3() {
	ds3 := priorityQueue{}
	ds3.EnQueue(3)
	ds3.EnQueue(5)
	ds3.EnQueue(10)
	ds3.EnQueue(2)
	ds3.EnQueue(7)

	fmt.Println(ds3.DeQueue())
	fmt.Println(ds3.DeQueue())
}

type priorityQueue []int

func (pq *priorityQueue) EnQueue(ele int) {
	*pq = append(*pq, ele)
	pq.upAdjust()
}
func (pq *priorityQueue) upAdjust() {
	childIndex := len(*pq) - 1
	parentIndex := (childIndex - 1) / 2
	tmp := (*pq)[childIndex]

	for childIndex > 0 && (*pq)[parentIndex] < tmp {
		(*pq)[childIndex] = (*pq)[parentIndex]
		childIndex = parentIndex
		parentIndex = parentIndex / 2
	}
	(*pq)[childIndex] = tmp
}
func (pq *priorityQueue) DeQueue() int {
	first := (*pq)[0]
	(*pq)[0] = (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	pq.downAdjust()
	return first
}
func (pq *priorityQueue) downAdjust() {
	parentIndex := 0
	tmp := (*pq)[parentIndex]
	childIndex := 2*parentIndex + 1
	Len := len(*pq)

	for childIndex < Len {
		if childIndex+1 < Len && (*pq)[childIndex+1] > (*pq)[childIndex] {
			childIndex = childIndex + 1
		}

		if tmp > (*pq)[childIndex] {
			break
		}

		(*pq)[parentIndex] = (*pq)[childIndex]
		parentIndex = childIndex
		childIndex = 2*childIndex + 1
	}

	(*pq)[parentIndex] = tmp
}

func ds4() {
	// a := []int{4, 7, 6, 5, 3, 2, 8, 1}
	// FastSortDouble(a, 0, len(a)-1)
	// fmt.Println(a)

	b := []int{2, 1, 0}
	FastSortSingle(b, 0, len(b)-1)
	fmt.Println(b)

	// c := []int{4, 7, 3, 5, 6, 2, 8, 1}
	// FastSortStack(c, 0, len(c)-1)
	// fmt.Println(c)
}

// 快排(小->大)的双边循环交换法-递归实现-左右两个指针,左找比基准元素大的,右找币基准元素小的,然后交换直到左右指针重合,然后吧基准元素跟左指针对应的元素交换
func FastSortDouble(arr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}

	privotIndex := partition(arr, startIndex, endIndex)
	FastSortDouble(arr, startIndex, privotIndex-1)
	FastSortDouble(arr, privotIndex+1, endIndex)
}
func partition(arr []int, startIndex, endIndex int) int {
	privot := arr[startIndex]
	left := startIndex
	right := endIndex
	for right != left {
		//右指针左移
		for left < right && arr[right] > privot {
			right--
		}
		//左指针右移
		for left < right && arr[left] <= privot {
			left++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	//privot和指针重合点交换
	arr[startIndex], arr[left] = arr[left], privot
	return left
}

// 快排(小->大)的单边循环法-递归实现
func FastSortSingle(arr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	//mark是边界
	privotIndex := partition1(arr, startIndex, endIndex)
	FastSortSingle(arr, startIndex, privotIndex-1)
	FastSortSingle(arr, privotIndex+1, endIndex)
}

// [4, 7, 6, 5, 3, 2, 8, 1]
// [4, 3, 6, 5, 7, 2, 8, 1]
// [4, 3, 2, 5, 7, 6, 8, 1]
// [4, 3, 2, 1, 7, 6, 8, 5]
// [1, 3, 2, 4] [7, 6, 8, 5]
// mark是边界
func partition1(arr []int, startIndex, endIndex int) int {
	privot := arr[startIndex]
	mark := startIndex

	for i := startIndex + 1; i <= endIndex; i++ {
		if arr[i] < privot {
			mark++

			arr[mark], arr[i] = arr[i], arr[mark]
		}
	}

	arr[startIndex], arr[mark] = arr[mark], privot
	return mark
}

// 快排(小->大)的栈实现
func FastSortStack(arr []int, startIndex, endIndex int) {
	stack := list.New()
	ma := make(map[string]int)
	ma["startIndex"] = startIndex
	ma["endIndex"] = endIndex

	stack.PushBack(ma)
	for stack.Len() != 0 {
		ma1 := stack.Remove(stack.Back())
		param := ma1.(map[string]int)
		privotIndex := partition1(arr, param["startIndex"], param["endIndex"])
		if param["startIndex"] < privotIndex-1 {
			leftParam := map[string]int{
				"startIndex": param["startIndex"],
				"endIndex":   privotIndex - 1,
			}
			stack.PushBack(leftParam)
		}
		if privotIndex+1 < param["endIndex"] {
			rightParam := map[string]int{
				"startIndex": privotIndex + 1,
				"endIndex":   endIndex,
			}
			stack.PushBack(rightParam)
		}
	}
}

// 计数排序(稳定排序=基于计数排序的优化版本)-适用场景:1.已知数组范围 2.最大值和最小值间隔不大 3.整数数组
// n = 数组长度, m = 数组最大值和最小值的间隔,时间复杂度O(n) = 3n+m = n + m 空间复杂度O(n) = m
func CountSort(arr []int) []int {
	//找出数组最大值
	max := arr[0]
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	//确定数组长度
	countArr := make([]int, max-min+1)
	for i := range arr {
		countArr[arr[i]-min]++
	}

	//变形累加前面的值
	for i := range countArr {
		if i > 0 {
			countArr[i] = countArr[i] + countArr[i-1]
		}
	}

	//倒序遍历原数组,
	sortArr := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		//找到该值的索引(索引-1=表示当前值在排序后的数组中的索引)
		index := countArr[arr[i]-min]
		index--
		countArr[arr[i]-min] = index

		sortArr[index] = arr[i]
	}

	return sortArr
}
func ds5() {
	arr := []int{
		9, 3, 5, 4, 9, 1, 2, 7, 8, 1, 3, 6, 5, 3, 4, 0, 10, 9, 7, 9,
	}
	arr = CountSort(arr)
	fmt.Println(arr)

	arr1 := []int{
		95, 94, 91, 98, 99, 90, 99, 93, 91, 92,
	}
	fmt.Println(CountSort(arr1))

	arr2 := []float64{
		4.12, 6.421, 0.0023, 3.0, 2.123, 8.122, 4.12, 10.09,
	}
	fmt.Println(BucketSort(arr2))
}

// 桶排序
func BucketSort(arr []float64) []float64 {
	if len(arr) < 1 {
		return arr
	}

	//1.找出最大和最小值
	max := arr[0]
	min := arr[0]
	for i := range arr {
		if i > 0 {
			if arr[i] > max {
				max = arr[i]
			}
			if arr[i] < min {
				min = arr[i]
			}
		}
	}

	//2.创建桶
	bucketLen := len(arr)
	buckets := make([][]float64, bucketLen) //桶的数量 = 数组的长度
	//jg      := (max - min)/(float64(len(arr)) - 1)   //每个桶的区间间隔
	d := max - min

	//3.将每个元素放到对应的范围的桶中
	for i := range arr {
		index := ((arr[i] - min) * (float64(bucketLen) - 1)) / d
		buckets[int(index)] = append(buckets[int(index)], arr[i])
	}

	//4.对每个桶进行排序
	for i := range buckets {
		sort.Float64s(buckets[i])
	}

	//5.输出全部元素
	sortedArr := make([]float64, len(arr))
	j := 0
	for i := range buckets {
		for i1 := range buckets[i] {
			sortedArr[j] = buckets[i][i1]
			j++
		}
	}
	return sortedArr
}

// 链表环的问题
type SingleList struct {
	Root *Node1
	Last *Node1
}
type Node1 struct {
	Val  int
	Next *Node1
}

func (s *SingleList) Push(n *Node1) {
	s.Last.Next = n
	s.Last = n
}
func NewSL(root *Node1) *SingleList {
	s := &SingleList{
		Root: root,
		Last: root,
	}
	return s
}

// 单向链表关于环的问题
func ds6() {
	two := &Node1{Val: 2}

	sl := NewSL(&Node1{Val: 5})
	sl.Push(&Node1{Val: 3})
	sl.Push(&Node1{Val: 7})
	sl.Push(two)
	sl.Push(&Node1{Val: 6})
	sl.Push(&Node1{Val: 8})
	sl.Push(&Node1{Val: 1})
	sl.Push(two)

	fmt.Println(isCycle(sl))
	fmt.Println(CycleLen(sl))
	fmt.Println(CycleEntry(sl), two)
}

// 是否有环
func isCycle(l *SingleList) bool {
	first := l.Root
	next := first
	for next != nil && next.Next != nil {
		first = first.Next
		next = next.Next.Next

		if first == nil || next == nil {
			break
		}
		if first.Val == next.Val {
			return true
		}
	}
	return false
}

// 延伸问题-环长度是多少 => 首次相遇后开始计数,第二次相遇时计数的长度 = 环长度
func CycleLen(l *SingleList) int {
	first := l.Root
	next := first
	lLen := 0
	isMet := false
	for first != nil && next.Next != nil {
		first = first.Next
		next = next.Next.Next
		if first == nil || next == nil {
			break
		}

		if first.Val == next.Val {
			if lLen != 0 {
				break
			}
			isMet = true
		}
		if isMet {
			lLen++
		}
	}
	return lLen
}

// 延伸问题-入环点在哪里 => p1(速度v1)和p2(速度v2 = 2*v1) => 在首次相遇后,将两者速度均调整为1,将p1移动到首节点,p2继续在环内行走,后续相遇的位置则为入环点
func CycleEntry(l *SingleList) *Node1 {
	p1 := l.Root
	p2 := p1
	var (
		entry *Node1
		speed = 2
	)
	for p1 != nil && p2.Next != nil {
		p1 = p1.Next
		if speed == 1 {
			p2 = p2.Next
		} else {
			p2 = p2.Next.Next
		}
		if p1 == nil || p2 == nil {
			break
		}

		if p1.Val == p2.Val {
			//第二次相遇后的位置,则为入环点
			if speed == 1 {
				entry = p1
				break

				//首次相遇,将p2的速度改为1,调整p1的位置到首节点继续行走
			} else if speed == 2 {
				speed = 1
				p1 = l.Root
			}
		}
	}
	return entry
}

// 最小栈的实现
type minStack struct {
	l    *list.List
	minL *list.List
}

func NewStack1() *minStack {
	return &minStack{
		l:    list.New(),
		minL: list.New(),
	}
}
func (this *minStack) Push(val int) {
	this.l.PushBack(val)
	//最小栈为空 或 最小栈的栈顶元素 大于 插入的元素 则插入最小栈中
	if this.minL.Len() == 0 || val < this.minL.Back().Value.(int) {
		this.minL.PushBack(val)
	}
}
func (this *minStack) Pop() int {
	if this.l.Back() != nil && this.minL.Back() != nil {
		last := this.l.Back()
		lastVal := last.Value.(int)
		this.l.Remove(last)

		minEle := this.minL.Back()
		minEleVal := minEle.Value.(int)
		if minEleVal == lastVal {
			this.minL.Remove(minEle)
		}
		return lastVal
	}
	return 0
}
func (this *minStack) GetMin() int {
	if this.minL.Back() != nil {
		return this.minL.Back().Value.(int)
	}
	return 0
}

// 最小栈面试题
func ds7() {
	minS := NewStack1()
	minS.Push(4)
	minS.Push(9)
	minS.Push(7)
	minS.Push(3)
	fmt.Println(minS.GetMin())
	minS.Pop()

	fmt.Println(minS.GetMin())
	minS.Pop()
}

// 求最大公约数解法
func ds8() {
	fmt.Println(MaxYue(12, 16))
	fmt.Println(MaxYue1(12, 16))
	fmt.Println(MaxYue2(12, 16))
}

// 辗转相除法
func MaxYue(a, b int) int {
	if a < b {
		a, b = b, a
	}
	y := a % b
	if y == 0 || b <= 1 {
		return b
	}
	return MaxYue(y, b)
}

// 辗转相减法
func MaxYue1(a, b int) int {
	if a == b {
		return a
	}
	if a < b {
		a, b = b, a
	}
	return MaxYue1(a-b, b)
}

func MaxYue2(a, b int) int {
	if a == b {
		return a
	}
	if a < b {
		a, b = b, a
	}
	a1 := a % 2
	b1 := b % 2
	if a1 == 0 && b1 == 0 {
		return MaxYue2(a>>1, b>>1) << 1
	} else if a1 == 0 && b1 != 0 {
		return MaxYue2(a>>1, b)
	} else if a1 != 0 && b1 == 0 {
		return MaxYue2(a, b>>1)
	} else if a1 != 0 && b1 != 0 {
		return MaxYue2(b, a-b)
	}
	return 1
}

// 求给定一个正整数n,判断是否是2的整数次幂
func ds9() {
	fmt.Println(Mi(1))
	fmt.Println(Mi(2))
	fmt.Println(Mi(3))
	fmt.Println(Mi(4))
}
func Mi(n int) bool {
	return n&(n-1) == 0
}

// 无序数组排序后的最大相邻差
func ds10() {
	arr := []int{2, 6, 3, 4, 5, 10, 9}
	fmt.Println(BucketMaxGap(arr))
}

// 利用桶排序的思想
type bucket struct {
	min int
	max int
}

func BucketMaxGap(arr []int) int {
	//计算最大值+最小值
	min, max := 0, 0
	arrLen := 0
	for _, v := range arr {
		arrLen++
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	d := max - min

	buckets := make([]bucket, arrLen)
	for _, v := range arr {
		index := ((v - min) * (arrLen - 1)) / d
		if buckets[index].min == 0 || buckets[index].min > v {
			buckets[index].min = v
		}
		if buckets[index].max == 0 || buckets[index].max < v {
			buckets[index].max = v
		}
	}

	fmt.Println(buckets)
	leftbu := buckets[0].max
	maxGap := 0
	for i := 1; i < arrLen; i++ {
		if buckets[i].min == 0 && buckets[i].max == 0 {
			continue
		}

		if buckets[i].min-leftbu > maxGap {
			maxGap = buckets[i].min - leftbu
		}
		leftbu = buckets[i].max
	}

	return maxGap
}

// 用栈实现队列
type BaseStack []int

func (this *BaseStack) Push(ele int) {
	*this = append(*this, ele)
}
func (this *BaseStack) Pop() int {
	last := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return last
}

type stackQueue struct {
	s1 BaseStack
	s2 BaseStack
}

func NewStackQueue() *stackQueue {
	return &stackQueue{
		s1: make(BaseStack, 0),
		s2: make(BaseStack, 0),
	}
}

func (this *stackQueue) In(ele int) {
	this.s1.Push(ele)
}
func (this *stackQueue) Out() int {
	if len(this.s2) == 0 {
		for len(this.s1) != 0 {
			this.s2.Push(this.s1.Pop())
		}
		if len(this.s2) == 0 {
			return 0
		}
	}
	return this.s2.Pop()
}
func ds11() {
	ds3 := NewStackQueue()
	ds3.In(1)
	ds3.In(2)
	ds3.In(3)
	fmt.Println(ds3.Out())
	fmt.Println(ds3.Out())
	fmt.Println(ds3.Out())
}

// 寻找全排列的下一个数 = 在一个整数包含的数字的全部组合中,找出一个大于且仅大于原数的新整数
// https://leetcode-cn.com/problems/next-permutation/
// 12345 -> 12354
// 12435 -> 12453
func ds12() {
	nums := []int{1, 2, 4, 3, 2}

	findTransferPoint := func(a []int) int {
		for i := len(a) - 1; i > 0; i-- {
			if a[i] > a[i-1] {
				return i
			}
		}
		return 0
	}

	//1.找到逆序区域的索引
	index := findTransferPoint(nums)
	if index == 0 {
		sort.Ints(nums)
		return
	}

	//2.吧逆序区域的索引的前一位和逆序区域中大于他的最小数字交换位置
	minIndex := index
	for i := index + 1; i < len(nums); i++ {
		if nums[i] > nums[index-1] && nums[i] < nums[minIndex] {
			minIndex = i
		}
	}
	nums[index-1], nums[minIndex] = nums[minIndex], nums[index-1]

	//3.把逆序区域顺序排序
	if index < len(nums)-1 {
		sort.Ints(nums[index:])
	}
	fmt.Println(nums)
}

// 删除k个数字后的最小值
func ds13() {
	nums := []int{3, 0, 2, 0, 0}
	k := 1
	fmt.Println(DelKMin(nums, k))
	fmt.Println(DelKMin1(nums, k))

	nums1 := []int{5, 4, 1, 2, 7, 0, 9, 3, 6}
	k1 := 3
	fmt.Println(DelKMin(nums1, k1))
	fmt.Println(DelKMin1(nums1, k1))
}

// 实现方法1
func DelKMin(nums []int, k int) []int {
	if len(nums) == k {
		nums = []int{0}
		return nums
	}

	//记录上一次的位置
	lastIndex := 0
	for k > 0 {
		hasCut := false

		//从左到右,找到左边数字 > 右边的位置,删除该左位置的值
		for i := lastIndex; i < len(nums); i++ {
			if i+1 > len(nums)-1 {
				break
			}
			if nums[i] > nums[i+1] {
				if i-1 > 0 {
					lastIndex = i - 1
				}
				hasCut = true
				a := make([]int, 0)
				a = append(a, nums[:i]...)
				a = append(a, nums[i+1:]...)
				for a[0] == 0 {
					a = a[1:]
				}
				nums = a
				break
			}
		}

		//没找到则删除最后一个
		if !hasCut {
			nums = nums[:len(nums)-1]
			lastIndex = len(nums) - 1
		}
		k--
	}
	return nums
}

// 实现方法2 -- 栈实现
func DelKMin1(nums []int, k int) []int {
	if len(nums) == k {
		return []int{0}
	}
	stack := BaseStack{}
	for i, v := range nums {
		if i == 0 {
			stack.Push(v)
			continue
		}

		if stack[len(stack)-1] > v && k > 0 {
			k--
			stack.Pop()
		}

		stack.Push(v)
	}

	for stack[0] == 0 {
		stack = stack[1:]
	}

	return stack
}

// 两个整数相加
func ds14() {
	fmt.Println(Sum("1234", "12345"))
}
func Sum(a, b string) string {
	aLen := len(a)
	if len(b) > aLen {
		aLen = len(b)
	}
	arr1 := make([]int, aLen+1)
	arr2 := make([]int, aLen+1)

	j := 0
	for i := len(a) - 1; i >= 0; i-- {
		arr1[j], _ = strconv.Atoi(string(a[i]))
		j++
	}
	j = 0
	for i := len(b) - 1; i >= 0; i-- {
		arr2[j], _ = strconv.Atoi(string(b[i]))
		j++
	}
	arr3 := make([]int, aLen+1)
	for i := range arr1 {
		sum := arr1[i] + arr2[i]
		if sum > 10 {
			sum = sum % 10
			if i+1 > aLen {
				arr3 = append(arr3, 1)
			} else {
				arr3[i+1]++
			}

		}
		arr3[i] = arr3[i] + sum
	}

	s := ""
	for i := len(arr3) - 1; i >= 0; i-- {
		if arr3[i] != 0 {
			s += strconv.Itoa(arr3[i])
		}
	}

	return s
}

// 动态规划
func ds15() {
	w := 10 //工人个数
	//金矿含金量
	g := []int{
		400, 500, 200, 300, 350,
	}
	//挖金矿需要的工人数量
	p := []int{
		5, 5, 3, 4, 3,
	}

	fmt.Println(Wa(w, len(g), p, g))
	fmt.Println(Wa1(w, len(g), p, g))
	fmt.Println(Wa3(w, len(g), p, g))
}

func Ma(a, b int) int {
	if b > a {
		a = b
	}
	return a
}

//动态规划-递归
// w 工人的数量
// n 金矿的数量
// p 挖金矿需要的工人数量
// g 挖金矿的含金量

// F(n,w) 表示 n个金矿,w个工人,挖矿的含金量最大时的函数
// 边界1: 金矿 = 0 或 工人 = 0 的时候 F(n,w) = F(0,0) = 0

// 情况1:当剩下的工人不足挖金矿需要的工人数量时(n >= 1, w < p[n-1]) =>  F(n,w) = F(n-1,w)
// 常规情况: (n >= 1, w >= p[n-1]) => F(n,w) = max(F(n-1,w),F(n-1,w-p[n-1]) + g[n-1])
func Wa(w, n int, p, g []int) int {
	//边界值---------
	//挖矿工人为0 或者 金矿数量为0
	if w == 0 || n == 0 {
		return 0
	}
	//挖矿工人数量 < 第n-1个金矿需要的工人数量
	if w < p[n-1] {
		return Wa(w, n-1, p, g)
	}
	//边界值---------

	//情况2, 不挖最后一个金矿
	b := Wa(w, n-1, p, g)

	//情况1, 挖最后一个金矿
	a := Wa(w-p[n-1], n-1, p, g) + g[n-1]

	return Ma(a, b)
}

// 动态规划-自底向上-空间换时间
// i : 金矿数量
// j : 工人数量
// 时间复杂度O(nw)
// 空间复杂度O(nw)
func Wa1(w, n int, p, g []int) int {
	table := make([][]int, len(g)+1)
	for i := range table {
		table[i] = make([]int, w+1)
	}
	//p[i-1] : 第i-1个金矿需要的人数
	//i : 第几个金矿
	//j : 工人的数量
	for i := 1; i <= len(g); i++ {
		for j := 1; j <= w; j++ {
			//如果工人的数量 < 挖矿需要的工人数量
			if j < p[i-1] {
				table[i][j] = table[i-1][j]
			} else {
				table[i][j] = Ma(table[i-1][j], table[i-1][j-p[i-1]]+g[i-1])
			}
		}
	}

	return table[len(g)][w]
}

// 时间复杂度O(nw)
// 空间复杂度O(w)
func Wa3(w, n int, p, g []int) int {
	result := make([]int, w+1)
	for i := 1; i < len(g); i++ {
		for j := w; j >= 1; j-- {
			if j >= p[i-1] {
				result[j] = Ma(
					result[j],
					result[j-p[i-1]]+g[i-1],
				)
			}
		}
	}
	return result[w]
}
