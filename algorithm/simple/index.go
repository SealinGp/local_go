package main

import (
	"fmt"
	"time"
	"os"
)

func init() {
  fmt.Println("Content-Type:text/plain;charset=utf-8\n\n");
}
func main() {
	startTime := time.Now();
	args      := os.Args;
    if len(args) <= 1 {
    	fmt.Println("lack param ?func=xxx");
    	echoTime(startTime);
    	return;
    }

	execute(args[1]);
	echoTime(startTime);
}
func execute(n string) {
	funs := map[string]func() {
		"twoSum" : twoSum,
	};		
	funs[n]();		
}
func echoTime(startTime time.Time) {
	since := time.Since(startTime).String();
    fmt.Println("processTime:",since);
}
/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。不能重复利用这个数组中同样的元素。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum

著作权归领扣网络所有.商业转载请联系官方授权,非商业转载请注明出处.
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
func twoSum() {
	nums      := []int{2,7,11,15};
	target    := 9;
	numIndex  := make([]int,0);		
	valueLeft := 0;

	for index,value := range nums {
		valueLeft = target - value;
		sli       := nums[index+1:];

		index2,find := in_slice(valueLeft,sli);
		if find {			
			numIndex = append(numIndex,index,index2+(index+1));			
		}		
	}

	fmt.Println(numIndex);
}

func in_slice(sliceElement int,intSlice []int) (int,bool) {
	find      := false;
	findIndex := 0;
	for index,element := range intSlice {
		if element == sliceElement {
			find      = true;
			findIndex = index;
			break;
		}
	}
	return findIndex,find;
}