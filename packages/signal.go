package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var sigFuncs = map[string]func(){
	"sig1": sig1,
	"sig2": sig2,
}

//监听所有停止信号
func sig1() {
	sigC := make(chan os.Signal)

	signal.Notify(sigC)
	fmt.Println("启动,pid:", os.Getpid())

	t := <-sigC
	fmt.Println("退出信号:", t)
}

func sig2() {
	pid := 4833
	err := syscall.Kill(pid, syscall.SIGTERM)
	if err != nil {
		fmt.Println("kill pid", pid, "error:", err.Error())
		return
	}
	fmt.Println("kill success!")
}
