package main

import (
	"container/list"
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
		"l1" : l1,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

func l1()  {
	ln := list.New()
	for i := 0 ; i < 5; i++ {
		ln.PushBack(i)
	}

	for e := ln.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}