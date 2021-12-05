package main

import (
	"fmt"
	"os"
)

var funcs map[string]func()

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("lack param ?func=xxx")
		return
	}

	addAllFuncs()

	funcName := args[1]
	if f, ok := funcs[funcName]; ok {
		f()
		return
	}
}

func addAllFuncs() {
	addFuncs(grpcFuncs)
	addFuncs(map[string]func(){
		"TestBuf": TestBuf,
	})
	addFuncs(map[string]func(){
		"QuicFunc": QuicFunc,
	})

}

func addFuncs(fcs map[string]func()) {
	for k, v := range fcs {
		funcs[k] = v
	}
}
