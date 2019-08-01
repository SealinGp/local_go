package main

import (
	"fmt"
)
//不同范围的变量可以重名
var a = "G"
func init() {
	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
}
func main() {
	n()
	m()
	n()
}

func n()  {
	fmt.Println(a)
}
func m()  {
	a := "O"
	fmt.Println(a)
}
