package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := "S"
	v := reflect.ValueOf(&i)
	v.Elem().SetString("??")
	fmt.Println(i)
}