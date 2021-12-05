package main

import (
	"fmt"
)

var forFuncs = map[string]func(){
	"for1":  for1,
	"for2":  for2,
	"for3":  for3,
	"for4":  for4,
	"for5":  for5,
	"for6":  for6,
	"for7":  for7,
	"for8":  for8,
	"for9":  for9,
	"for10": for10,
	"for11": for11,
}

//goto
func for1() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//goto 和 标签(一般为全大写)来模拟循环体
	i := 0
START:
	fmt.Println(i)
	i++
	if i < 5 {
		goto START
	}
}

//redeclare variable
func for2() {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Println(v)
		v = 5
	}
}

//dead loop
func for3() {
	for i := 0; ; i++ {
		fmt.Println(i)
	}
}
func for4() {
	for i := 0; ; {
		fmt.Println(i)
		i++
	}
}
func for5() {
	for i := 0; i < 3; {
		fmt.Println("Value of i:", i)
	}
}
func for6() {
	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}
}

//multiple condition 多条件
func for7() {
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s+"a" {
		fmt.Println(i, j, s)
	}
}

//label continue break
func for8() {
LABEL1:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if j == 2 {
				continue LABEL1
			}
			fmt.Println(i, j)
		}
	}
}
func for8_1() {
LABEL1:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if j == 2 {
				break LABEL1
			}
			fmt.Println(i, j)
		}
	}
}
func for9() {
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if j >= 2 {
				continue
			}
			fmt.Println(i, j)
		}
	}

	for i := 0; i <= 3; i++ {
	LA1:
		for j := 0; j <= 3; j++ {
			if j >= 2 {
				continue LA1
			}
			fmt.Println(i, j)
		}
	}
}
func for10() {
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if j == 2 {
				break
			}
			fmt.Println(i, j)
		}
	}

	for i := 0; i <= 3; i++ {
	LA1:
		for j := 0; j <= 3; j++ {
			if j == 2 {
				break LA1
			}
			fmt.Println(i, j)
		}
	}
}

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/05.6.md
func for11() {
	i := 0
	for { //since there are no checks, this is an infinite loop
		if i >= 3 {
			break
		}
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}
}
