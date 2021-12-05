package main

import (
	"fmt"
)

//不同范围的变量可以重名
var (
	a  = "G"
	a1 = "G"
)

var scopeFuncs = map[string]func(){
	"scope1": scope1,
	"scope2": scope2,
}

func scope1() {
	scope1_n()
	scope1_m()
	scope1_n()
}
func scope1_n() {
	fmt.Println(a)
}
func scope1_m() {
	a := "O"
	fmt.Println(a)
}

func scope2() {
	scope2_n()
	scope2_m()
	scope2_n()
}

func scope2_n() {
	fmt.Println(a1)
}
func scope2_m() {
	a1 = "O"
	fmt.Println(a1)
}
