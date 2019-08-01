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
		"keyword1": keyword1,
	}
	funs[n]()
}

func keyword1() {
	var kWord = [...]string{
		//key words
		"break", "default", "func", "interface",
		"select", "case", "defer", "go", "map",
		"struct", "chan", "else", "goto", "package",
		"switch", "const", "fallthrough", "if","range",
		 "type", "continue", "for", "import", "return", "var",

		 //func,type
		 "append","bool","byte","cap","close","complex","complex64",
		 "complex128","uint16","copy","false","float32","float64",
		 "imag","int","int8","int16","uint32","uint32","int64","iota","len",
		 "make","new","nil","panic","uint64","print","println","real","recover",
		 "string","true","uint","uint8","uintptr",
	}
	kLen := len(kWord)
	for i := 0; i < kLen; i++ {
		fmt.Println(kWord[i])
	}
}
