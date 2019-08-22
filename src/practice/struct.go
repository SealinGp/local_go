package main

import (
	"fmt"
	"os"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/10.0.md
struct 跟 method
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
		"struct1" : struct1,
		"struct2" : struct2,
		"struct3" : struct3,
		"struct4" : struct4,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

type b struct {
	b1 string
	b2 int
	b3 rune
}
func struct1()  {
	var b1 *b
	b1 = new(b)
	b1.b1  = "abc"
	b1.b2  = 2
	b1.b3  = 'A'
	fmt.Println(b1)

	b2 := new(b)
	b2.b1  = "abc"
	b2.b2  = 2
	b2.b3  = 'A'
	fmt.Println(b2)

	b3 := &b{b1:"abc",b2:2,b3:'A'}
	fmt.Println(b3)
	//表达式new(Type) 和 &Type{} 是等价的
}

//递归结构体,二叉树
type Tree struct {
	le *Tree       //左节点
	data float64
	ri *Tree       //右节点
}

//带标签的结构体,reflect 包可以获取tag
type TagType struct {
	field1 bool   "An important answer"
	field2 string "the name of the thing"
	field3 int    "how much there are"
}

//匿名结构体,此时变量名为其类型
type innerS struct {
	in1 int
	in2 int
}
//这里表示outerS 继承了 innerS
type outerS struct {
	b int
	c float32
	int       "等价于 int int"
	innerS    "等价于 innerS innerS"
}
func struct2()  {
	out := new(outerS)
	out.b = 6
	out.c = 7.5
	out.int = 60

	out.innerS.in1 = 5
	out.innerS.in2 = 10
	//out.in1 = 5
	//out.in2 = 10
	fmt.Println(out)
}

//方法

//结构体方法
type me struct {
	a string
	b string
}
//非结构体方法
type I []int
func struct3()  {
	m := &me{a:"abc",b:"def"}
	m.set("aaa","bbb")
	fmt.Println(m)
	m.echo()

	i := I{1,2,3}
	s := i.sum()
	fmt.Println(s)
}
//m为一个指针,若不需要用m,可以用_代替或不写 (_ *me) 或 (*me)
func (m *me)set(a,b string)  {
	m.a = a
	m.b = b
}
func (m me)echo()  {
	fmt.Println(m)
}
func (i I)sum() (s int) {
	for _,v := range i  {
		s += v
	}
	return
}

func struct4()  {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}
