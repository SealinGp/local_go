package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

/*
匿名,闭包函数
*/
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

func execute(n string) {
	funs := map[string]func(){
		"noneFunc1": noneFunc1,
		"noneFunc2": noneFunc2,
		"noneFunc3": noneFunc3,
		"noneFunc4": noneFunc4,
	}
	funs[n]()
}

func noneFunc1() {
	fmt.Println(f())
}
func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func noneFunc2() {
	a1 := f1(2)
	b1 := a1(3)
	fmt.Println(b1)
}
func f1(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func noneFunc3() {
	a := f2()         //x = 0
	fmt.Println(a(1)) //1 x = 1
	fmt.Println(a(2)) //3 2+1=3
	fmt.Println(a(3)) //6 3+3=6
}
func f2() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

//使用闭包调试
func noneFunc4() {
	t := time.Now()
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(file, line)
	}
	where() // /root/www/sea/local_go/src/practice/noneFunc.go 76

	time.Sleep(2 * time.Second)
	fmt.Println(time.Now().Sub(t))
}
