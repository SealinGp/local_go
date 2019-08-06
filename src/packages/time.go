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
time packages
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
	a := 5
	switch a  {
	case 5:
		fmt.Println("a = 5")
	case 6,7:
		fmt.Println("a = 6 or 7")
	default:
		fmt.Println("a = none")
	}

	/*a1 := []uint8{5,6,7}
	var c,d uint8
	switch c,d = a1[0],a1[1] {
	case c == 5:
		fmt.Println(a1[0])
	case d == 6:
		fmt.Println(a1[1])
	}*/
}