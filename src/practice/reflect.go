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

//动态执行函数
type ref struct {}
func execute(n string) {
	ref    := NewRef()
	refVal := reflect.ValueOf(&ref).Elem()
	//nil为输入参数
	refVal.MethodByName(n).Call(nil)
}
func NewRef() *ref {
	return &ref{}
}

func (r *ref)Ref1()  {
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
func (r *ref)Ref2()  {
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

//通过反射设置值
func (r *ref)Ref3()  {
	//普通值
	a  := 123
	av := reflect.ValueOf(&a).Elem()
	av.SetInt(456)
	fmt.Println(a)

	//结构值(属性名必须大写(public),否则不可设置)
	type s struct {
		Name string
		Age uint8
	}
	b := s{}
	bv := reflect.ValueOf(&b).Elem()
	bv.Field(0).SetString("abc")
	var ua uint8 = 12
	bv.Field(1).Set(reflect.ValueOf(ua))
	fmt.Println(b)
}