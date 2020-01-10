package main
/*
#include <stdio.h>
void SayHello(const char* s);
void SayHello(const char* s) {
    puts(s);
}
void printint(int v) {
    printf("printint: %d\n",v);
}
struct Test {
    int i;
    float f;
};
enum DAY {
    Mon = 1,Tue,Wed,Thu,Fri,Sat,Sun
};
 */
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)
//https://github.com/chai2010/advanced-go-programming-book/blob/master/ch2-cgo/readme.md
func main() {
	C.SayHello(C.CString("hello world"))

	v := 42
	C.printint(C.int(v))

	//访问c定义的struct
	var t C.struct_Test
	fmt.Println(t.i)
	fmt.Println(t.f)

	var d C.enum_DAY = C.Mon
	fmt.Println(d)
	fmt.Println(C.Tue)

	//指针转换
	var a1 *int
	var b1 *float32
	newA1 := (*float32)(unsafe.Pointer(a1))
	newB1 := (*int)(unsafe.Pointer(b1))
	fmt.Println(reflect.TypeOf(newA1))
	fmt.Println(reflect.TypeOf(newB1))
}

