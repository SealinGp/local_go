package main

import (
	"fmt"
	"os"
)

var funcs map[string]func()

func init() {
	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
}

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
	addFuncs(dataTypeFuncs)
	addFuncs(designModelFuncs)
	addFuncs(funcErrHandleFuncs)
	addFuncs(funcInterFuncs)
	addFuncs(funcRecurisiveFuncs)
	addFuncs(funcRoutine)
	addFuncs(funcSelectFuncs)
	addFuncs(judgeFuncs)
	addFuncs(keywordFuncs)
	addFuncs(syncFuncs)
	addFuncs(regFuncs)
	addFuncs(channelFuns)
	addFuncs(defineFuncs)
	addFuncs(mapFuncs)
	addFuncs(pointerFuncs)
	addFuncs(rangeFuncs)
	addFuncs(sliceFuncs)
	addFuncs(map[string]func(){
		"HeapMem": HeapMem,
	})

}

func addFuncs(fcs map[string]func()) {
	for k, v := range fcs {
		funcs[k] = v
	}
}
