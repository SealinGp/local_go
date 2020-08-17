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

	mt := reflect.TypeOf((*interface{})(nil)).Elem()
	a1 := a{}
	aT := reflect.TypeOf(&a1)
	fmt.Println(aT.Implements(mt))

	v1 := reflect.ValueOf(Add)
	if v1.Kind() == reflect.Func {

	}
}
type a struct {

}

func Add()  {

}