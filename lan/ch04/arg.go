package main

import (
	"fmt"
	"unsafe"
)

type MyStruct struct {
	i int
	j int
}

func myFunction1(ms *MyStruct) {
	ptr := unsafe.Pointer(ms)
	for i := 0; i < 2; i++ {
		c := (*int)(unsafe.Pointer((uintptr(ptr) + uintptr(8*i))))
		*c += i + 1
		fmt.Printf("[%p] %d\n", c, *c)
	}
}

func main() {
	a := &MyStruct{i: 40, j: 50}
	myFunction1(a)
	fmt.Printf("[%p] %v\n", a, a)
}
