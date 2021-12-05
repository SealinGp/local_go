package main

import (
	"fmt"
)

var defineFuncs = map[string]func(){
	"define1": define1,
}

func define1() {
	//variable
	var i1 int8 = 1
	var i2 = 1
	i3 := 1
	//这种声明一般用于全局变量
	var (
		i4 int8 = 2
		i5 int8 = 2
	)

	//const variable
	const s1 string = "abc"
	const s2 = "abc"
	const (
		s3        = "abc"
		s4 string = "abc"
	)

	fmt.Println(i1, i2, i3, i4, i5, s1, s2, s3, s4)
}
