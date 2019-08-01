package main

import (
	"fmt"
	"sort"
	"strconv"
)

func init() {
	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
}

/*
type i interface 是一种存储方法的类型
当某个类型t实现了接口类型I中的所有方法的时候,说明该类型t实现了I接口

why?
 1.writing generic algorithm （泛型编程）
 2.hiding implementation detail （隐藏具体实现）
 3.providing interception points
*/
func main() {
	// nokia := Nokia{number:13129551271};
	//以下为两种实现方法

	//实现Phone接口
	// var phone Phone;
	// phone = new(arg);
	// phone.call(nokia);

	//同时也将方法call存入arg struct中,可直接调用
	// a := arg{};
	// a.call(nokia);

	//a := arg{};
	// callPhone(a,nokia);

	// ArrObj();
	test()
}

//struct
type Nokia struct {
	number uint64
}

//save function call() in arg struct
//该类型实现了phone接口
type arg struct {
}

//interface
type Phone interface {
	call(Nokia)
}

//execute the interface
//类型a 实现了 phone(接口)类型
func (a arg) call(no Nokia) {
	// fmt.Printf("calling %d...\n",no.number);

	str := "calling " + strconv.FormatUint(no.number, 10) + "..."
	fmt.Println(str)

	/*str := "calling %d...\n";
	str = fmt.Sprintf(str,no.number);
	fmt.Println(str);*/
}

func callPhone(p Phone, no Nokia) {
	p.call(no)
}

type obj struct {
	name string
	age  int8
}
type ArObj []obj

//数组对象
func ArrObj() {
	a := [2]obj{
		{"v_kenqzhang", 11},
		{"xavierma", 18},
	}
	fmt.Println(a)
}

/*
1.泛型编程(writing generic algorithm)
sort.Sort() 函数的参数是一个Interface,包含了3个方法 Len(),Swap(),Less()
只要类型ArObj实现了这三个方法,就可以使用sort.Sort(ArObj)函数
*/
func (a ArObj) Len() int {
	return len(a)
}
func (a ArObj) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

//索引i对应的元素是否在索引j的后面
func (a ArObj) Less(i, j int) bool {
	return a[i].age < a[j].age
}
func test() {
	ab := ArObj{
		{"bob", 31},
		{"john", 42},
		{"Michael", 17},
		{"jenny", 26},
	}
	sort.Sort(ab)
	fmt.Println(ab)

	people := []obj{
		{"bob", 31},
		{"john", 42},
		{"Michael", 17},
		{"jenny", 26},
	}
	sort.Sort(ArObj(people)) //ArObj() 表示类型转换为ArObj
	fmt.Println(people)
}

/*
2.隐藏具体实现(hiding implementation detail)
*/
