package simple

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
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

/**
链接：https://www.nowcoder.com/questionTerminal/66ca0e28f90c42a196afd78cc9c496ea
来源：牛客网
原理：ip地址的每段可以看成是一个0-255的整数,把每段拆分成一个二进制形式组合起来,然后把这个二进制数转变成
一个长整数。
举例：一个ip地址为10.0.3.193
每段数字             相对应的二进制数
10                   00001010
0                    00000000
3                    00000011
193                  11000001
组合起来即为：00001010 00000000 00000011 11000001,转换为10进制数就是：167773121,即该IP地址转换后的数字就是它了
的每段可以看成是一个0-255的整数,需要对IP地址进行校验
 */
func (*Ref)IPv4ToInt64()  {
	ipv4 := "192.168.1.0"
	fmt.Println(IPv4ToInt(ipv4))
}
func IPv4ToInt(ipv4 string) int64 {
	ipv4StrSli    := strings.Split(ipv4,".")
	ipv4StrSliLen := len(ipv4StrSli)

	var ipv4Int64 int64
	var j = 3
	for i := 0; i < ipv4StrSliLen; i++ {
		v, _  := strconv.Atoi(ipv4StrSli[i])

		//ipv4Int64
		start :=  j * 8
		for v != 0 {
			ipv4Int64 = ipv4Int64 + int64((v % 2) << start)
			v /= 2
			start++
		}
		j--
	}
	return ipv4Int64
}
/**
zego面试题
每次输入一个int数,一共输入100w次,请输出结果为前100大的数
提示:最小堆
 */
func (*Ref)Top100()  {
	max        := 100
	inputCount := int(math.Pow10(6))
	var input int
	top100 := make([]int,max,max)
	for i := 0; i < inputCount; i++ {

		fmt.Scanf("%d",&input)
		if i <= max-1 {
			top100[i] = input
			continue
		}

		if i == max {
			sort.Ints(top100)
		}
		if input < top100[0] {
			continue
		}
		for j := range top100  {
			if j >= 1 {
				top100[j-1] = top100[j]
			}
			if input <= top100[j] && j-1 >= 0 {
				top100[j-1] = input
				break
			}
			if j + 1 > max-1  {
				top100[j] = input
				break
			}
		}
	}
	fmt.Println(top100)
}