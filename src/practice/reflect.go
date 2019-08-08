package main

import (
	"fmt"
	"os"
	"reflect"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/11.10.md
reflect
*/
func init() {
	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
}
func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("lack param ?func=xxx")
		return
	}

	execute(args[1])
}

func execute(n string) {
	funs := map[string]func(){
		"ref1" : ref1,
		"ref2" : ref2,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

func ref1()  {
	type myInt int
	var m myInt = 5
	v := reflect.ValueOf(m)
	fmt.Println(
		v,
		v.Type(),
		v.Kind(),
		v.Interface(),
		v.Interface().(myInt),
		reflect.TypeOf(m),
	)
}

//反射获取对象中的信息(包含属性,方法,标签)
type t struct {
	s1,s2,s3 string  "abc"
}
func (t1 t)String(b,c string) string  {
	return t1.s1 + t1.s2 + t1.s3 + b + string(c)
}
func ref2()  {
	var s interface{} = t{"s1_","s2_","s3_"}

	attributes    := reflect.TypeOf(s)
	value         := reflect.ValueOf(s)
	for i := 0; i < attributes.NumField(); i++ {
		fmt.Println(
			//对象的属性,属性类型,属性标签
			attributes.Field(i).Name,attributes.Field(i).Type,attributes.Field(i).Tag,

			//对象的属性值
			value.Field(i),

			value.NumMethod(),
		)
	}

	//调用对象中的方法,并给参数,若输入参数,则给nil
	s1 := "bbb"
	s2 := "_ccc"
	sv1 := reflect.ValueOf(s1)
	sv2 := reflect.ValueOf(s2)

	res := value.MethodByName("String").Call([]reflect.Value{sv1,sv2})
	//res := value.Method(0).Call([]reflect.Value{sv})
	fmt.Println(res)
}