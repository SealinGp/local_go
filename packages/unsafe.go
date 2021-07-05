package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//获取操作数在内存中的字节大小
	fmt.Println(unsafe.Sizeof(float64(0)))
	fmt.Println(unsafe.Sizeof(int64(0)))

	//
	//unsafe.Offsetof()

	//
	//unsafe.Alignof()

	x1 := x{b: 0}
	pb := &x1.b
	*pb = 1
	fmt.Println(*pb, x1.b)

	//
	pb1 := (*uint16)(
		unsafe.Pointer(
			uintptr(unsafe.Pointer(&x1)) + unsafe.Offsetof(x1.b),
		),
	)
	*pb1 = 2
	fmt.Println(x1.b)
}

type x struct {
	a bool
	b int16
	c []int
}

func Floate64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}
