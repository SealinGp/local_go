package main

import (
	"fmt"
	"math"
	"os"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/11.1.md
interface 和 reflect
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
		"int1" : int1,
		"int2" : int2,
		"int3" : int3,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

type I1 interface {
	Method1(a int) int
}
type str1 struct {

}
func (s *str1)Method1(a int) int {
	return a*2
}
func int1()  {
	s := str1{}
	r := s.Method1(2)
	fmt.Println(r)
}

type Simpler interface {
	Get() int
	Set(int)
}
type Simple struct {
	s int
}
func (si *Simple)Get() int {
	return si.s
}
func (si *Simple)Set(i int)  {
	si.s = i
}
func m1(S Simpler,i int) int {
	S.Set(i)
	return S.Get()
}
func int2()  {
	S := &Simple{s:5}
	a := m1(S,6)
	fmt.Println(a,S)
}

//类型断言:如何检测和转换接口变量的类型
type Square struct {
	side float32
}
type Circle struct {
	radius float32
}
type Shaper interface {
	Area() float32
}
func (sq *Square)Area() float32 {
	return sq.side*sq.side
}
func (c *Circle)Area() float32 {
	return c.radius * c.radius * math.Pi
}
func int3()  {
	sq1 := &Square{side:5}

	//sh必须为接口变量
	var sh Shaper
	sh = sq1

	//判断该接口变量的类型 t,ok := 接口变量.(*结构变量)
	if t,ok := sh.(*Square);ok {
		fmt.Println(t,"square")
	}
	if u,ok := sh.(*Circle);ok {
		fmt.Println(u,"circle")
	}

	switch t := sh.(type) {
		case *Square:
			fmt.Println("Square",t)
		case *Circle:
			fmt.Println("Circle",t)
		case nil:
			fmt.Println("nil",t)
		default:
	}
}

func int4()  {
	classifier(13,-14.3,"BELOGIUM",complex(1,2),nil,false)
}
func classifier(items ...interface{})  {
	for i,x := range items  {
		switch x.(type) {
			case bool:
			fmt.Println("bool",i)
			case float64:
			fmt.Println("float64",i)
			case int,int64:
			fmt.Println("int or int64",i)
			case nil:
			fmt.Println("nil",i)
		}
	}
}