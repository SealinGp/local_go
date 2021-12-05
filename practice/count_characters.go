package main

import (
	"fmt"
	"unicode/utf8"
)


func characters() {
	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Println("str1 len:", len(str1))
	fmt.Println("str1 RuneCount:", utf8.RuneCountInString(str1))

	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Println("str2 len:", len(str2))
	fmt.Println("str2 RuneCount:", utf8.RuneCountInString(str2))
}
