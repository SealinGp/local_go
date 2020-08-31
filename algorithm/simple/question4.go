package simple

import (
	"fmt"
)

//面试算法题

//1.将一个100W的数组分成长度相等的2组数组,要求左边的数组每个元素<右边的数组中的元素
func (*Ref)M1()  {
	arr := []int{6,5,4,3,2,1}
	fs(arr,0,len(arr)-1)
	fmt.Println(arr[:len(arr)/2],arr[len(arr)/2:])
}
func fs(arr []int,start,end int)  {
	if start >= end {
		return
	}
	privotIndex := part(arr,start,end)
	fs(arr,start,privotIndex-1)
	fs(arr,privotIndex+1,end)
}

func part(arr []int,startIndex,endIndex int) int {
	privot     := arr[startIndex]
	left,right := startIndex,endIndex


	for right != left {
		for left < right && arr[right] > privot {
			right--
		}
		for left < right && arr[left] <= privot {
			left++
		}
		if left < right {
			arr[left],arr[right] = arr[right],arr[left]
		}
	}

	arr[startIndex],arr[left] = arr[left],privot
	return left
}

//随机数问题
//https://leetcode-cn.com/problems/implement-rand10-using-rand7/
func (*Ref)R10()  {
	c := 0
	for  {
		a := rand7()
		b := rand7()
		c = (a-1)*7 + b
		if c <= 40 {
			break
		}
	}
	fmt.Println((c % 10) + 1)
}


//倒水问题 gcd?
//https://leetcode-cn.com/problems/water-and-jug-problem/
func (*Ref)PourWater()  {

}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2zsx1/
func (*Ref)Mp()  {
	prices := []int{7,6,4,3,1}

	maxPro := 0
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0  {
			maxPro += tmp
		}
	}
	fmt.Println(maxPro)
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/
func (*Ref)Rotate()  {
	nums     := []int{-1,-100,3,99}
	k        := 2

	if k <= 0 {
		return
	}
	L    := len(nums)
	move := L

	curI  := 0
	tmp   := nums[curI]
	first := true
	for move > 0 {
		newI       := (curI + k)%L

		if first {
			tmp         = nums[newI]
			nums[newI]  = nums[curI]
			first = false
		} else {
			tmp1       := nums[newI]
			nums[newI]  = tmp
			tmp         = tmp1
		}
		curI        = newI

		move--
	}

	//for newI :=  range nums {
	//	oldI := (newI - k) % L
	//	if oldI < 0 {
	//		oldI += L
	//	}
	//	fmt.Print(nums[oldI])
	//}
	fmt.Println(nums)
}

func (*Ref)ContainsDup(nums []int) bool {
	tmp  := make(map[int]int)
	for _, v := range nums {
		tmp[v]++
		if tmp[v] > 1 {
			return true
		}
	}
	return false
}