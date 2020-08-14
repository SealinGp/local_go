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
		"switch3" : switch3,
	}
	funcMap[funcN]()
}

//fallthrough 找到符合的条件,该条件后面的一个case会执行,注意:此关键词不可在type-switch判断时使用
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

/*
type-switch类型断言
判断该接口变量的类型 t,ok := 接口变量.(结构变量)

若结构变量实现接口变量的方法为 *指针形式的话,那么这里应该是
t,ok := 接口变量.(*结构变量)
*/
func switch3()  {
	var s s2
	s = s1{a1:"abc"}

	//类型断言1
	switch t := s.(type) {
		case s1:
			fmt.Println("s1",t)
		default:
			fmt.Println("def",t)
	}

	//类型断言2
	if _,ok := s.(s2);ok {
		fmt.Println("s2 type")
	}
}
type s1 struct {
	a1 string
}
type s2 interface {
	s2F()
}
func (s s1)s2F()  {
	fmt.Println(s.a1)
}