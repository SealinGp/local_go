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
	addFuncs(map[string]func(){
		"lissajous":  Lissajous,
		"tar1":       tar1,
		"Buffer":     Buffer,
		"test":       test,
		"Tabwriter":  Tabwriter,
		"clockwall":  clockwall,
		"unsafeFunc": unsafeFunc,
	})
	addFuncs(reqFuncs)
	addFuncs(sigFuncs)
	addFuncs(strFuncs)
	addFuncs(timeFuncs)

}

func addFuncs(fcs map[string]func()) {
	for k, v := range fcs {
		funcs[k] = v
	}
}
