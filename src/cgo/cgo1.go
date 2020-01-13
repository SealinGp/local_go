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
	"log"
	"os"
	"reflect"
	"sync"
	"unsafe"
)
//https://github.com/chai2010/advanced-go-programming-book/blob/master/ch2-cgo/readme.md
func main() {
	if len(os.Args) <= 1 {
		log.Fatal("func required.")
	}
	fun := map[string]func() {
		"cgo1":cgo1,
		"cgo2":cgo2,
	}
	fun[os.Args[1]]()
}

//basic
func cgo1()  {
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

//https://github.com/chai2010/advanced-go-programming-book/blob/master/ch2-cgo/ch2-06-qsort.md
type go_qsort_compare_info struct{
	base unsafe.Pointer
	elemnum int
	elemsize int
	less func(a,b int) bool
	sync.Mutex
}
func cgo2()  {
	values := []int64{42,9,101,95,27,25}
	g := go_qsort_compare_info{
		base:     nil,
		elemnum:  0,
		elemsize: 0,
		less:     nil,
		Mutex:    sync.Mutex{},
	}
	Slice(
		values, func(a, b int) bool {
			return values[a] < values[b]
		},
		g,
	)
}

//https://github.com/chai2010/advanced-go-programming-book/blob/master/ch2-cgo/ch2-07-memory.md
func cgo3()  {
	
}

