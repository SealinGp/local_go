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
		"err1" : err1,
		"err2" : err2,
		"err3" : err3,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
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
func err2()  {
	pro(err1)
}
func pro(g func())  {
	defer func() {
		if err := recover();err != nil {
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
func err3()  {
	err3_1()
	fmt.Println("call err3")
}
func err3_1()  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover err3_1",r)
		}
	}()
	fmt.Println("call err3_1")
	err3_2(0)

	fmt.Println("err3_1 return")
}
func err3_2(i int)  {
	if i > 3 {
		fmt.Println("panicing")
		panic(fmt.Sprintf("%v",i))
	}
	defer fmt.Println("defer in err3_2",i)
	fmt.Println("call err3_2",i)
	err3_2(i + 1)
}