package main

import (
	"fmt"
	"reflect"
)

func main() {


}
//2.倒水问题
//3.开车,距离,加油问题



func test()  {
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