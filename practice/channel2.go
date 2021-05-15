package main

import (
	"fmt"
	"os"
)

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
		"chan1": chan1,
	}
	if nil == funs[n] {
		fmt.Println("func", n, "unregistered")
		return
	}
	funs[n]()
}

func chan1() {

}
