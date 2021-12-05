package main

import (
	"fmt"
	"unsafe"
)

//https://www.flysnow.org/2017/07/06/go-in-action-unsafe-pointer.html
var unsafeFuncs = map[string]func(){
	"unsafe1": unsafe1,
	"unsafe2": unsafe2,
}

func unsafe1() {
	i := 10
	ip := &i
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(i, *ip, *fp)
}

type user struct {
	name string
	age  int
}

func unsafe2() {
	u := new(user)
	fmt.Println(*u)

	name := (*string)(unsafe.Pointer(u))
	*name = "张3"
	age := (*int)(unsafe.Pointer(
		uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)),
	)
	*age = 20

	fmt.Println(*u)
}
