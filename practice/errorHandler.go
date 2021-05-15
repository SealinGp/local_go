package main

import (
	"errors"
	"fmt"
	"os"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/13.0.md
错误处理
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
		"err1": err1,
		"err2": err2,
		"err3": err3,
		"err4": err4,
		"err5": err5,
	}
	if nil == funs[n] {
		fmt.Println("func", n, "unregistered")
		return
	}
	funs[n]()
}
func err1() {
	err := errors.New("error")
	panic("pan")
	fmt.Printf(err.Error())
}

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/13.5.md
//抓取panic错误
func err2() {
	pro(err1)
}
func pro(g func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	g()
}

/*
call err3_1
call err3_2 0
call err3_2 1
call err3_2 2
call err3_2 3
panicing
defer err3_2 3
defer err3_2 2
defer err3_2 1
defer err3_2 0
recover err3_1 4
call err3
*/
func err3() {
	err3_1()
	fmt.Println("call err3")
}
func err3_1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover err3_1", r)
		}
	}()
	fmt.Println("call err3_1")
	err3_2(0)

	fmt.Println("err3_1 return")
}
func err3_2(i int) {
	if i > 3 {
		fmt.Println("panicing")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("defer in err3_2", i)
	fmt.Println("call err3_2", i)
	err3_2(i + 1)
}

//defer的坑
//https://www.jianshu.com/p/9a7364762714?tdsourcetag=s_pctim_aiomsg
//defer的函数的参数是在执行defer时计算的，defer的函数中的变量的值是在函数执行时计算的
//同一个函数里面的defer,栈结构,最后写的先执行,不同函数的defer,先执行的函数的defer先执行
func err4() {
	err4_1() //  x = 0
	err4_2() //  x = 7
}
func err4_1() (x int) {
	defer fmt.Println("in err4_1 defer x=", x)
	x = 7
	return 9
}
func err4_2() (x int) {
	x = 7
	defer fmt.Println("in err4_2 defer x=", x)
	return 9
}

func err5() {
	err5_1() // x=9
	err5_2() // n=0 x=9
}
func err5_1() (x int) {
	defer func() {
		fmt.Println("in err5_1 defer x=", x)
	}()
	x = 7
	return 9
}
func err5_2() (x int) {
	defer func(n int) {
		fmt.Println("in err5_2 n=", n)
		fmt.Println("in err5_2 x=", x)
	}(x)
	x = 7
	return 9
}
