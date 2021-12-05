package main

import (
	"container/list"
	"fmt"
)

var listFuncs = map[string]func(){
	"l1": l1,
}

func l1() {
	ln := list.New()
	for i := 0; i < 5; i++ {
		ln.PushBack(i)
	}

	for e := ln.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
