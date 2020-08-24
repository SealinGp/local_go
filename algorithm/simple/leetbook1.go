package simple

import "fmt"

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
