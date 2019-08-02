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
		"func1":func1,
		"func2":func2,
	}
	funcMap[funcN]()
}

//命名 & 非命名返回值
func func1()  {
	i1, i2 := func1_1(10)
	i3, i4 := func1_2(10)
	fmt.Println(
		i1,i2,
		i3,i4,
	)
}
func func1_1(i1 int) (int,int)  {
	return i1*10,i1*20
}
func func1_2(i1 int) (i2 int, i3 int)  {
	i2, i3 = i1*10,i1*20
	return
}

//多参数组
func func2() ()  {
	func2_1("a","b","c")

	strArr := []string{"d","e","f"}
	func2_1(strArr...)
}
func func2_1(s2 ...string)  {
	fmt.Println(s2)
}