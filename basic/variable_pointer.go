package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
*variable_type 指针变量声明
&variable_name 引用变量的内存地址

 var variable_name *varibale_type = &variable_name; 指针变量指向变量地址,默认nill(空指针)
*/

var pointerFuncs = map[string]func(){
	"pointer1": pointer1,
	"pointer2": pointer2,
	"pointer3": pointer3,
	"pointer4": pointer4,
	"pointer5": pointer5,
	"pointer6": pointer6,
}

// pointer basic
func pointer1() {
	var i1 int8 = 1
	var i1Pointer *int8 = &i1 //pointer variable declare, defaule value : nil(空指针)

	//i1Pointer = &i1 = 指针所指的内存位置
	fmt.Println("i1:", i1)
	fmt.Println("i1Pointer:", i1Pointer)
	fmt.Println("*i1Pointer:", *i1Pointer)
	fmt.Println("*i1Pointer++")
	*i1Pointer++
	fmt.Println("i1:", i1)
	fmt.Println("--------------------")

	type stru1 struct {
		name string
		age  uint8
	}
	stru := stru1{name: "v_kenqzhang", age: 123}
	struPointer := &stru
	fmt.Println("stru:", stru)
	fmt.Println("struPointer:", struPointer)
	fmt.Println("struPointer.name:", struPointer.name)
	struPointer.name = "v_sshyu"
	fmt.Println("*struPointer:", *struPointer)
	fmt.Println("stru:", stru)

	type stru1_1 struct {
		name *string
		age  uint8
	}
	str := "v_kenqzhang"
	stru1_ := stru1_1{name: &str, age: 22}
	stru1Pointer := stru1_
	fmt.Println("stru1_", stru1_)
	fmt.Println("stru1Pointer", stru1Pointer)
	*stru1_.name = "v_sshyu"
	fmt.Println("stru1Pointer.name", *stru1Pointer.name)
}

/*
指针使用:*xxx 指针声明&xxx

*/
type stru2 struct {
	a string
	b string
}
type interface1 interface {
	exchange()
}

func (st *stru2) exchange() {
	mid := st.a
	st.a = st.b
	st.b = mid
}
func pointer2() {
	st := &stru2{"a", "b"}
	st.exchange()
	fmt.Println(*st)
	fmt.Println(st.a)

	st1 := stru2{"a", "b"}
	st1.exchange()
	fmt.Println(st1)
	fmt.Println(st1.a)
}

type stru3 struct {
	a string
	b string
}
type interface2 interface {
	exchange2()
}

func (st stru3) exchange2() stru3 {
	mid := st.a
	st.a = st.b
	st.b = mid
	return st
}
func pointer3() {
	st := stru3{"a", "b"}
	st1 := st.exchange2()
	fmt.Println(st)
	fmt.Println(st1)
}

/**
https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html
指针变量:存储的值是变量的地址,自身有一个地址

1.栈和指针
协程内存分为2块
[堆内存] : 存放全局变量
[栈内存] : 存放函数栈帧,分为[活动内存][非活动内存]
[活动内存] : 函数调用时在活动内存中进行,调用完成返回后在[非活动内存]区域
2个函数栈帧之间的调用存在一个转换过程(活动内存和非活动内存区域之间的转换)

2.变量逃逸分析(可使用 go build -gcflags "-m -m" 追踪代码中是否存在逃逸的变量)
变量逃逸: 函数栈帧中的局部变量 因 与函数栈帧之间共享 导致该局部变量从函数栈内存逃逸到堆内存中去
可能导致变量逃逸的情况: 1.多个函数栈帧之间共享变量地址 2.使用interface作为接受参数 3.使用make关键词的长度不确定

3.内存管理(go test -run none -bench banmark函数名 -benchtime 测试的秒数 -benchmem)
*/

func pointer4() {
	u1 := createUserV1()
	u2 := createUserV2()
	println("u1 address:", &u1, "u2 address:", &u2, "u2 value", u2)
}

type user struct {
	name  string
	email string
}

func createUserV1() user {
	v1 := user{"a", "b"}
	println("v1 address:", &v1)
	return v1
}
func createUserV2() *user {
	v2 := user{"a1", "b1"}
	println("v2 address:", &v2)
	return &v2
}

//逃逸分析的可能情况: 1.会导致编译器不知道分配内存多少导致该变量逃逸到堆内存中去
func pointer5() {
	pointer5_1(10)
}
func pointer5_1(size int) {
	b := make([]byte, size)
	b = append(b, 'a')
	fmt.Println(string(b))
}

func pointer6() {
	//unsafe.Pointer: 可以跟任意指针类型互转,可以跟uintptr类型互转
	v1 := uint(12)
	v2 := int(13)
	fmt.Println(reflect.TypeOf(&v1), reflect.TypeOf(&v2))
	p := &v1
	fmt.Println(reflect.TypeOf(p)) //*uint
	//*int -> *uint,类型转换
	p = (*uint)(unsafe.Pointer(&v2))
	fmt.Println(reflect.TypeOf(p)) //*uint

	//uintptr : 内置类型,能存储指针的整型,可用于指针运算,在64bit平台上底层的数据类型是:typedef uint64 uintptr;
	type x struct {
		a bool
		b int16
		c []int
	}
	x1 := x{}
	//pb := &x1.b
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x1)) + unsafe.Offsetof(x1.b)))
	*pb = 42
	fmt.Println(x1.b)
}
