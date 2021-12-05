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
		"chan1":      chan1,
		"Cond":       Cond,
		"cop1":       cop1,
		"characters": characters,
		"netcat2":    netcat2,
		"reverb1":    reverb1,
		"reverb2":    reverb2,
		"netcat3":    netcat3,
		"IMServer":   IMServer,
		"IMClient":   IMClient,
		"mail":       mail,
		"pack1":      pack1,
		"tickerFunc": tickerFunc,
		"udpFunc":    udpFunc,
	})
	addFuncs(atomicFuncs)
	addFuncs(arrFuncs)
	addFuncs(channelFuncs)
	addFuncs(cmdFuncs)
	addFuncs(compressFuncs)
	addFuncs(ctxFuncs)
	addFuncs(encryptFuncs)
	addFuncs(errFuncs)
	addFuncs(flagFuncs)
	addFuncs(forFuncs)
	addFuncs(fFuncs)
	addFuncs(gobFuncs)
	addFuncs(runtineFuncs)
	addFuncs(implFuncs)
	addFuncs(jsonFuncs)
	addFuncs(listFuncs)
	addFuncs(lockFuncs)
	addFuncs(mapFuncs)
	addFuncs(mathFuncs)
	addFuncs(noneFuncs)
	addFuncs(recFuncs)
	addFuncs(reqFuns)
	addFuncs(rpcFuns)
	addFuncs(rsaFuncs)
	addFuncs(scopeFuncs)
	addFuncs(selFuncs)
	addFuncs(structFuns)
	addFuncs(switchFuncs)
	addFuncs(tcpFuncs)
	addFuncs(unsafeFuncs)
	addFuncs(webFuncs)
	addFuncs(xmlFuncs)

}

func addFuncs(fcs map[string]func()) {
	for k, v := range fcs {
		funcs[k] = v
	}
}
