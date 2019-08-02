package main

import (
	"fmt"
	"os"
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
	}
	funcMap[funcN]()
}

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/06.6.md
//斐波那契数列 前2个数为1,从第3个数开始,每个数为前两个数之和
func rec1()  {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fib1(i)
		fmt.Println(result)
	}
}
func fib1(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fib1(n-1) + fib1(n-2)
	}
	return
}