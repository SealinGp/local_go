package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

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
		"sig1": sig1,
		"sig2": sig2,
	}
	funs[n]()
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
