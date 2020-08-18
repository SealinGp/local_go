package main

import (
	"fmt"
	"os"
)

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
		"create_slice": create_slice,
		"for_slice":    for_slice,
		"fun_slice":    fun_slice,
		"fs1":    fs1,
	}
	funs[n]()
}
func test() int {
	var ret int
	return ret
}

var a int

/*
slice:
 切片 = 动态数组(长度可变)
长度跟容量:
ref : https://my.oschina.net/joeyjava/blog/504249
如果通过make函数创建Slice的时候指定了容量参数,
那内存管理器会根据指定的容量的值先划分一块内存空间,
然后才在其中存放有数组元素,多余部分处于空闲状态,在Slice上追加元素的时候,
首先会放到这块空闲的内存中,如果添加的参数个数超过了容量值,
内存管理器会重新划分一块容量值为原容量值*2大小的内存空间


切片 或 数组 长度len和容量cap
数组: 长度和容量相等,不可append
切片: 长度和容量可不相等,可append
例: s := make([]int,2,3) 的切片,索引赋值对长度,append对容量,超容量给双倍
s[0] = 1 //1,2  len=2 cap=3
s[1] = 2 //1,2  len=2 cap=3
s = append(s,3) //1,2,3   len=3 cap=3  (s[2] = 3 //报错)
s = append(s,4)//1,2,3,4  len=4 cap=6
*/
//声明切片
func create_slice() {
	//直接创建数组切片 []type,length,capability(maxLength) 长度 容量(最大长度(可选))
	/*var sli3 = make([]int, 5, 10)
	sli3[1] = 3
	fmt.Println(sli3)
	fmt.Println("----------------------")*/

	//基于切片创建切片
	var sli1  = [5]string{"a","b","c","d","e"}
	var sli2  = sli1[0:2]  //startIndex:endIndex <=> [startIndex,endIndex) 左闭右开
	var sli21 = sli1[:2]   //不填写为下限 <=> 0:2 <=> [0,2)
	var sli22 = sli1[1:]   //不填写为上限 <=> 1:maxIndex+1 <=> [1,5)
	var sli23 = sli1[3:]   //不填写为上限 <=> 3:len
	sli24    := sli1[1:3]
	sli2[0] = "abc"
	fmt.Println(sli1)
	fmt.Println(sli2)
	fmt.Println(sli21)
	fmt.Println(sli22)
	fmt.Println(sli23)
	fmt.Println(sli24)//b,c
}

//遍历切片foreach
func for_slice() {
	var sli1 [3]string = [3]string{"a", "b", "c"}
	for i, v := range sli1 {
		fmt.Println(i, v)
	}
}

//切片函数
func fun_slice() {
	//len() cap()
	var sli1 = make([]string, 3, 5)
	sli1[0] = "abc"
	sli1[1] = "abc"
	sli1[2] = "abc"
	echo_str_slice(sli1)
	fmt.Println("---------------")

	//append(slice,val):入栈 copy()
	sli1 = append(sli1, "0")
	echo_str_slice(sli1)

	var sli2 = make([]string, len(sli1), cap(sli1))
	copy(sli2, sli1)
	echo_str_slice(sli2)
}

func echo_str_slice(sli []string) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)
}
func echo_int_slice(sli []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)
}

//切片扩容机制
//从源码看: 1.期望容量 > 旧容量*2 ,则新容量=期望容量 2.期望容量 < 旧容量*2 && 旧容量<1024 则新容量=旧容量*2 3.期望容量 < 旧容量*2 && 旧容量>=1024 新容量=旧容量*1.25
//从实际看: 根据切片类型,扩容机制有所不同
//https://juejin.im/post/6844903812331732999
func fs1()  {
	//byte
	a := []byte{1,8}
	a = append(a,1,1,1)
	fmt.Println(cap(a)) //8

	//int expectCap+1
	b := []int{23,51}
	b = append(b,4,5,6)
	fmt.Println(cap(b)) //6

	//int32
	c := []int32{1,23}
	c = append(c,2,5,6)
	fmt.Println(cap(c)) //8

	//struct
	type D struct {
		age byte
		name string
	}
	d := []D{
		{1,"123"},
		{2,"234"},
	}
	d = append(d,D{3,"456"},D{4,"567"},D{4,"678"})
	fmt.Println(cap(d)) //5
}