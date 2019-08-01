package main

import (
	"fmt"
	"os"
	"strings"

	//"strconv"
	"time"
)

const (
	timeLayOut = "2006:01:02:15:04:05"
)

/*
tar打包文件
*/
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
func execute(funcN string)  {
	funcMap := map[string]func(){
		"time1" : time1,
	}
	funcMap[funcN]()
}
func time1()  {
	now := time.Now()
	fmt.Println(
		strings.Split(now.Format(timeLayOut),":"),
	)
}