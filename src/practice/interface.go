package main

import (
	"fmt"
	"math"
	"os"
	"sort"
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
		"int4" : int4,
		"int5" : int5,
		"int6" : int6,
		"int7" : int7,
		"int8" : int8,
		"int9" : int9,
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

//java泛型 = go空接口
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
			default:
			fmt.Println("not found type")
		}
	}
}

type Appender interface {
	Append(int)
}
type Lener interface {
	Len()int
}
type List []int

func (l *List)Append(v int)  {
	//传进来的是指针(&l 地址),需要通过*l调用其值来 append
	*l = append(*l,v)
}
func (l List)Len() int {
	return len(l)
}
func CountInfo(a Appender, start,end int)  {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}
func LongEnough(l Lener) bool  {
	return l.Len()*10 > 42
}

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/11.6.md

type a int[]
(*a)func xx1() {}
(a)func xx2() {}

定义:
var a1 a 或 a1 = make(a,0) 或 a1 :=  a{}
可调用:
a1.xx1() //在函数xx1中修改a的值,会函数影响外的a的值
a1.xx2() //在函数xx1中修改a的值,不会函数影响外的a的值

定义
a1 := new(a) 或 a1 := &a{}
可调用:
a1.xx1() //在函数xx1中修改a的值,会函数影响外的a的值
a1.xx2() //在函数xx1中修改a的值,不会函数影响外的a的值
*/
func int5()  {
	//定义一个类型的值
	var lst List // <=> lst := make(List,0)
	fmt.Println(lst)
	//不可以使用CountInfo()因为CountInfo中使用的Append方法是定义在指针上的,而这里定义的是在值上
	if LongEnough(lst) {
		fmt.Println(lst,"long enough")
	}

	//定义一个类型的指针值
	plst := new(List)// <=> plst := &List{}
	CountInfo(plst,1,10)
	if LongEnough(plst) {
		fmt.Println("long enough",plst)
	}
}

//参考:https://blog.csdn.net/jiang_mingyi/article/details/81811217
//平行赋值
func int6() {
	//先算左边i,s[0] = 2,"Z" ZBC
	i := 1
	s := []string{"A", "B", "C"}
	i, s[i-1] = 2, "Z"
	fmt.Println(s)

	//a := []int{1,2,3,4}
	//defer func(a []int) {
	//	fmt.Println(a)
	//}(a)
	////引发了panic 赋值成功
	//a[0],a[4] = a[1],a[2]

	b := []int{1,2,3,4}
	defer func(a []int) {
		fmt.Println(a)
	}(b)
	//引发panic  赋值失败
	b[0],b[1] = b[2],b[4]

	/*
	总结
	平行赋值的执行顺序: 右索引,取址 左索引,取址
	*/
}

//对周一~周日进行排序
type day struct {
	num int "周几对应的数字大小"
	shortName string "缩写"
	longName string "全称"
}
type dayArr struct {
	data []*day
}
func (p *dayArr)Len() int {
	return len(p.data)
}
func (p *dayArr)Less(i,j int) bool {
	return p.data[i].num < p.data[j].num
}
func (p *dayArr)Swap(i,j int)  {
	p.data[i],p.data[j] = p.data[j],p.data[i]
}
func int7()  {
	d := dayArr{
		data:[]*day{
			{num:0,shortName:"SUN",longName:"Sunday"},
			{num:1,shortName:"MON",longName:"Monday"},
			{num:3,shortName:"WED",longName:"Wednesday"},
			{num:2,shortName:"TUE",longName:"Tuesday"},
			{num:4,shortName:"THU",longName:"Thursday"},
			{num:5,shortName:"FRI",longName:"Friday"},
			{num:6,shortName:"SAT",longName:"Saturday"},
		},
	}
	sort.Sort(&d)
	if !sort.IsSorted(&d) {
		fmt.Println("fail")
		return
	}
	
	for _,d := range d.data {
		fmt.Println(d)
	}
}

//通用类型 或 泛型
type Ele interface {}
type AnyArr []Ele
type intArr []int
func int8()  {
	b := AnyArr{"abc",123,'A',intArr{3,4}}
	for k,v := range b  {
		switch v.(type) {
			case int:
				fmt.Println("int",k,v)
			case string:
				fmt.Println("string",k,v)
			case rune:
				fmt.Println("rune",k,v)
			case intArr:
				fmt.Println("intArr",k,v)
			default:
				fmt.Println("unknown",k,v)
		}
	}
}

//接口到接口类型转换
type myPrint interface {
	print()
}
type myPrintS interface {
}
type myS struct {
	a interface{}
}
func (m myS) print() {
	fmt.Println(m.a)
}

//接口 动态类型 转换 interface myPrintS -> myPrint
func p(mp myPrintS)  {
	//接口变量类型断言
	if m,ok := mp.(myPrint);ok {
		m.print()
	} else {
		fmt.Println("not interface type myPrint,can't print")
	}
}
func int9()  {
	var m myPrintS
	ms := myS{a:"bacas"}
	m = ms
	p(m)
}