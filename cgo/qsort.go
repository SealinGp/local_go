package main

/**
#include <stdlib.h>
typedef int (*qsort_cmp_func_t)(const void* a,const void* b);
extern int _cgo_qsort_compare(void* a,void* b);
*/
import "C"
import (
	"log"
	"reflect"
	"unsafe"
)

//https://github.com/chai2010/advanced-go-programming-book/blob/master/ch2-cgo/ch2-06-qsort.md
func Slice(slice interface{}, less func(a, b int) bool, g go_qsort_compare_info) {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		log.Fatal("slice param required")
	}
	if sv.Len() == 0 {
		return
	}
	g.Lock()
	defer g.Unlock()
	g.base = unsafe.Pointer(sv.Index(0).Addr().Pointer())
	g.elemnum = sv.Len()
	g.elemsize = int(sv.Type().Elem().Size())
	g.less = less

	C.qsort(
		g.base,
		C.size_t(g.elemnum),
		C.size_t(g.elemsize),
		C.qsort_cmp_func(C._cgo_qsort_compare),
	)
}
