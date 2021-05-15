package main

import (
	"fmt"
	"os"
)

//不同范围的变量可以重名
var (
	a  = "G"
	a1 = "G"
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
func execute(funcN string) {
	funcMap := map[string]func(){
		"scope1": scope1,
		"scope2": scope2,
	}
	funcMap[funcN]()
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
