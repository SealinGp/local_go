package simple

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
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
func (*Ref)TwoSum() {
	nums      := []int{2, 7, 11, 15}
	target    := 9

	numIndex  := make([]int, 0)
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
func (*Ref)Reverse()  {
	//转string
	x  := rand.New(rand.NewSource(time.Now().UnixNano())).Int31()
	if x == 0 {
		log.Fatal(x)
	}

	negative := false
	xStr     := strconv.Itoa(int(x))
	if x < 0 {
		negative = true
		xStr = xStr[1:]
	}
	xStrLen  := len(xStr)
	var xByte []byte

	//string反转
	for i := xStrLen-1;i >= 0;i-- {
		if i == xStrLen-1 && xStr[i] == 48 {
			continue
		}
		xByte = append(xByte,xStr[i])
	}
	xStr = string(xByte)

	//转int
	result,Err := strconv.Atoi(xStr)
	if Err != nil {
		log.Fatal(Err.Error())
	}
	if negative {
		result = -result
	}
	if result < -2147483648 || result > 2147483647  {
		result = 0
	}
	fmt.Println(x)
	fmt.Println(result)
}
func (*Ref)Reverse2()  {
	//转string
	//x  := rand.New(rand.NewSource(time.Now().UnixNano())).Int31()
	x      := -120

	negative := false
	if x < 0 {
		negative = true
		x        = -x
	}
	result := 0
	xStr   := strconv.Itoa(x)
	xLen   := len(xStr)
	intMax := int(math.Pow(2,31))
	xMax   := intMax - 1
	xMin   := -intMax
	for i := xLen;i > 0; i-- {
		result += (x%10)*int(math.Pow10(i-1))
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
func (*Ref)IsPalindrome()  {
	x     := 12121

	isPalindrome := false
	xReverse     := 0
	xReverseS    := []int{}
	xLeft        := x
	for {
		remainder := xLeft % 10
		xReverseS  = append(xReverseS,remainder)
		if xLeft = xLeft/10;xLeft == 0 {
			break
		}
	}

	xReverseSLen := len(xReverseS)
	for i,v := range xReverseS {
		xReverse += v*int(math.Pow10(xReverseSLen-i-1))
	}

	isPalindrome = x == xReverse
	fmt.Println(isPalindrome)
}
func (*Ref)IsPalindrome1()  {
	x     := 1212
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