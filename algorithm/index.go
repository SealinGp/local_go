package main

import (
	"algorithm/simple"
	"fmt"
	"os"
	"reflect"
	"time"
)

func main() {
	startTime := time.Now()
	defer func() {
		d := time.Now().Sub(startTime)
		fmt.Println("processTime:",d)
	}()


	args     := os.Args
	if len(args) <= 1 {
		fmt.Println("lack param ?func=xxx")
		return
	}

	ref    := simple.Ref{}
	refVal := reflect.ValueOf(&ref)
	refVal.MethodByName(args[1]).Call(nil)
}

// 算法的时间复杂度 O()
// T(n)是语句执行的次数的函数;随着n的增加,T(n)增长最慢的算法,我们称为最优算法;即算法的时间复杂度小

// 假设 写一个 1+2+3+...+100的程序
// 算法1
func Index1() int {
	sum,n := 0,100            //执行1次
	for i := 1; i <= n ; i++ {//执行n+1次
		sum += 1
	}
	// 算法1时间复杂度位O(n)
	return sum
}
// 算法2
func Index2() int {
	sum,n := 0,100   //执行1次
	sum = (1 + n) * (n/2)//执行1次
	// 算法2时间复杂度位O(1)
	return sum
}

// 数据结构
//