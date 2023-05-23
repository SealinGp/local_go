package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

var funcs = map[string]func(){
	"linkTable": linkTable,
	"stack":     stack,
	"queue":     queue,
	"hashTable": hashTable,
	"dsu":       dsu,
	"Heap":      Heap,
	"ds1":       ds1,
	"ds2":       ds2,
	"ds3":       ds3,
	"ds4":       ds4,
	"ds5":       ds5,
	"ds6":       ds6,
	"ds7":       ds7,
	"ds8":       ds8,
	"ds9":       ds9,
	"ds10":      ds10,
	"ds11":      ds11,
	"ds12":      ds12,
	"ds13":      ds13,
	"ds14":      ds14,
	"ds15":      ds15,
}

func main() {

	args := os.Args
	if len(args) <= 1 {
		fmt.Println("lack param ?func=xxx")
		return
	}

	funcName := args[1]
	if f, ok := funcs[funcName]; ok {
		f()
		return
	}
}

func addAllFuncs() {
	addFuncs(struct1Funcs)
	addFuncs(structFuncs)
}

func addFuncs(fcs map[string]func()) {
	for k, v := range fcs {
		funcs[k] = v
	}
}
