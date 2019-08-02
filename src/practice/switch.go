package main

import (
	"fmt"
	"os"
)

/*
tar打包文件
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
func execute(funcN string)  {
	funcMap := map[string]func(){
		"switch1" : switch1,
		"switch2" : switch2,
	}
	funcMap[funcN]()
}

//fallthrough 找到符合的条件,该条件后面的一个case会执行
func switch1()  {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

func switch2()  {
	a := 5
	switch a  {
	case 5:
		fmt.Println("a = 5")
	case 6,7:
		fmt.Println("a = 6 or 7")
	default:
		fmt.Println("a = none")
	}
}