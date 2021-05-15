package main

import (
	"fmt"
	"math"
	"os"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/08.2.md
8.map
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
		"mat1": mat1,
	}
	if nil == funs[n] {
		fmt.Println("func", n, "unregistered")
		return
	}
	funs[n]()
}

func mat1() {
	fmt.Println(math.Pow10(3))

	fmt.Println(math.Pow(-2, 2))

}
