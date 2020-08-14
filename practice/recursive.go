package main

import (
	"fmt"
	"os"
	"time"
)

func init() {
	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
}
func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("lack param ?func=xxx")
		return
	}

	execute(args[1])
}
func execute(funcN string)  {
	funcMap := map[string]func(){
		"rec1" : rec1,
		"rec2" : rec2,
		"rec3" : rec3,
	}
	funcMap[funcN]()
}

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/06.6.md
//斐波那契数列 前2个数为1,从第3个数开始,每个数为前两个数之和
func rec1()  {
	s      := time.Now()

	result := 0
	for i := 0; i <= 40; i++ {
		result = fib1(i)
		fmt.Println(result)
	}

	fmt.Println(time.Now().Sub(s))
}
func fib1(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		//1 + 1 = 2
		//(1+1) + 1
		res = fib1(n-1) + fib1(n-2)
	}
	return
}

//闭包(不使用递归) 生成斐波那契数列
func rec2()  {
	s      := time.Now()
	fib    := fib2(11)
	println(fib)
	fmt.Println(time.Now().Sub(s))
}
func fib2(n int) int {
	fibNums := []int{}

	fib     := func(n int) int {
		if n <= 1 {
			fibNums = append(fibNums,1)
		} else {
			len1   := len(fibNums)
			num    := fibNums[len1-1] + fibNums[len1-2]
			fibNums = append(fibNums,num)
		}
		return fibNums[len(fibNums)-1]
	}
	res := 1
	for i := 0; i <= n; i++ {
		res = fib(i)
	}
	return res
}

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/06.12.md
//缓存
func rec3()  {
	s      := time.Now()
	var result uint64

	for i := 0; i < LIM; i++ {
		result = fib3(i)
		fmt.Println(result)
	}

	fmt.Println(time.Now().Sub(s))
}
const LIM = 41
var fibs [LIM]uint64
func fib3(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fib3(n-1) + fib3(n-2)
	}
	fibs[n] = res
	return
}