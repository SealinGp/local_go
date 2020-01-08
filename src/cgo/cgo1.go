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
 */
import "C"
func main() {
	C.SayHello(C.CString("hello world"))

	v := 42
	C.printint(C.int(v))
}

