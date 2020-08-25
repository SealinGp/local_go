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

func part(arr []int,left,right int) int {
	if len(arr) <= 1 {
		return left
	}
	tmpIndex := left
	tmp := arr[left]
	for left < right {
		for left < right && arr[right] > tmp {
			right--
		}
		for left < right && arr[left] <= tmp {
			left++
		}
		if left < right {
			arr[left],arr[right] = arr[right],arr[left]
		}
	}

	arr[tmpIndex],arr[left] = arr[left],tmp
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
