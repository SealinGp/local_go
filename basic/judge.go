package main

import (
	"fmt"
	"os"
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
		"if_sentences":     if_sentences,
		"for_sentences":    for_sentences,
		"switch_sentences": switch_sentences,
	}
	funs[n]()
}

func if_sentences() {
	var b bool = true
	var a bool = false
	var c bool = true
	if b || a { //true
		fmt.Println(b) //true
	}

	if b && a { //false
		fmt.Println(a)
	}

	if b &&
		a &&
		c { //false
		fmt.Println(c)
	}
}

func for_sentences() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//while
	i := 10
	for i > 0 {
		i--
		fmt.Println(i)
	}
}

func switch_sentences() {
	condition := true
	switch !condition {
	case false:
		fmt.Println(false)
	case true:
		fmt.Println(true)
	default:
		fmt.Println("error!")
	}
	return
}
