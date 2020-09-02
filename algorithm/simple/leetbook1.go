package simple

import (
	"container/list"
	"fmt"
	"sort"
)

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
func (*Ref)Nums()  {
	nums  := []int{0,1,1,2,2}


	len1  := removeDuplicates(nums)

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
	left  := 0
	right := left+1
	for right < numsL && left < numsL {
		if nums[left] == nums[right] {
			right++
		} else {
			if right != left + 1 {
				nums[left+1],nums[right] = nums[right],nums[left+1]
			}
			left++
			right++
		}
	}
	return left+1
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2zsx1/
func (*Ref)Mp2()  {

}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x21ib6/
func (*Ref)SN()  {
	nums := []int{2,2,1}

	s := 0
	for i := range nums {
		s = s ^ nums[i]
	}
	fmt.Println(s)
}

//https://leetcode-cn.com/problems/combination-sum/
func (*Ref)S1()  {
	candidates := []int{2,3,5}
	target     := 8

	fmt.Println(combinationSum(candidates,target))
}
func combinationSum(candidates []int, target int) [][]int {
	current  := list.New()
	results  := [][]int{}
	dfs1(candidates,0,current,target, func(a []int) {
		results = append(results,a)
	})
	return results
}
//深度优先搜索
func dfs1(candidates []int,index int,current *list.List,remainTarget int,fun func(a []int))  {
	if remainTarget < 0 {
		return
	}
	if remainTarget == 0 {
		f := current.Front()
		arr := []int{}
		for f != nil {
			arr = append(arr,f.Value.(int))
			f = f.Next()
		}
		fun(arr)
		return
	}


	for i := index; i < len(candidates); i++ {
		current.PushBack(candidates[i])
		dfs1(candidates,i,current,remainTarget - candidates[i],fun)
		current.Remove(current.Back())
	}
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2y0c2/
func (*Ref)Intersect()  {
	fmt.Println(intersect([]int{7,2,2,4,7,0,3,4,5},[]int{3,9,8,6,1,9}))
	fmt.Println(intersect1([]int{7,2,2,4,7,0,3,4,5},[]int{3,9,8,6,1,9}))
}
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	nums1L := len(nums1)
	nums2L := len(nums2)

	arr   := []int{}
	left  := 0
	right := 0
	for left < nums1L && right < nums2L {
		for left < nums1L && right < nums2L && nums1[left] < nums2[right] {
			left++
		}
		for left < nums1L && right < nums2L && nums2[right] < nums1[left] {
			right++
 		}

		if left < nums1L && right < nums2L && nums1[left] == nums2[right] {
			arr = append(arr,nums1[left])
			left++
			right++
		}
	}
	return arr
}
func intersect1(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1,nums2 = nums2,nums1
	}
	tmp := make(map[int]int)
	for _, v := range nums1 {
		tmp[v]++
	}
	arr := []int{}
	for _, v := range nums2 {
		if _, ok := tmp[v];ok {
			arr = append(arr,v)
			tmp[v]--
			if tmp[v] <= 0 {
				delete(tmp,v)
			}
		}
	}
	return arr
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2cv1c/
func (*Ref)PlusOne()  {
	fmt.Println(plusOne([]int{9}))
}
func plusOne(digits []int) []int {
	for i := len(digits)-1; i >= 0; i--  {
		digits[i]++
		if digits[i] < 10 {
			break
		}
		digits[i] = digits[i] % 10
	}

	if digits[0] == 0 {
		digits = append([]int{1},digits...)
	}
	return digits
}