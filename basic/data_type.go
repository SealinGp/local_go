package main

import (
	"fmt"
	"os"
)

/*
1 bit(二进制位) 一个二进制位,一个二进制位 表示两种状态,0或1
8 bit 0000 0000 ~ 1111 1111

ASCII   : 英语字符 和 二进制位 的映射

http://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html
Unicode : 这是一种所有符号的编码,Unicode 只是一个符号集,它只规定了符号的二进制代码,却没有规定这个二进制代码应该如何存储
UTF-8 就是在互联网上使用最广的一种 Unicode 的实现方式

1 byte(字节) = 8 bit(比特)
1 Kb        = 1024 byte
1 Mb        = 1024 Kb
1 Gb        = 1024 Mb

fmt.Printf()格式字符
 %d   : 十进制输出 (负数带符号)
 %o   : 八进制输出 (无前缀0)
 %x|X : 十六进制输出无符号整数(无前缀Ox)
 %u   : 十进制输出无符号整数(unsigned)
 %f   : 小数形式输出单,双精度实数
 %e|E : 指数形式输出单,双精度实数
 %g|G : 以%f或%e中较短的输出宽度输出单、双精度实数
 %c   : 单个字符输出
 %s   : 字符串输出
 %v   : 其他类型的变量
 %b   : 以二进制输出（%08b）
= 深拷贝

引用变量
值类型(int,float,bool,string,slice,struct)的变量的值存储在栈中
go中,指针属于引用类型,被引用的变量会存储在堆中,便于垃圾回收,且内存空间比
栈更大

_ 的作用 (https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/04.4.md)
_ 用于抛弃值,如5在 _, b = 5, 7 中被抛弃,实际上他是一个只写变量,不可读取到他的值

init 函数
在每个包完成初始化后自动执行,优先级>main

位运算

二元运算符
只可用于整数类型的变量,需要他们拥有等长位模式时
与 & : 将与计算的结果,true 转为1, false转为0
  1 & 1  true   1
  1 & 0  false  0
  0 & 1  false  0
  0 & 0  false  0
或 | : 将或计算的结果,true 转为1, false转为0
  1 | 1  true  1
  1 | 0  true  1
  0 | 1  true  1
  0 | 0  false 0
异或 ^ : 不同为1,相同为0
  1 ^ 1 0
  1 ^ 0 1
  0 ^ 1 1
  0 ^ 0 0
位清除 &^ : 将指定位置上的值设置为0

一元运算符
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/04.5.md

按位补足 ^
^10 = -01 ^ 10 = -11

位左移 << :
左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。
其功能把"<<"左边的运算数的各二进位全部左移若干位，
由"<<"右边的数指定移动的位数，高位丢弃，低位补0。

优先级  运算符
 7     ^ !
 6     * /(对于整数运算,则结果依然为整数) %(求余数) << >> & &^
 5     + - | ^
 4     == != < <= >= >
 3     <-
 2     &&
 1     ||

转义字符:
\n : 换行
\r : 回车
\t : tab键

rune类型中,一个中文占一个字节(相当于java中的char类型),
ascii 对应的转换为
rune -> string  65 -> 'A'
r1 := 'A'
r1Str := string(r1)

string类型中,一个中文占3个字节
*/
//执行顺序 全局变量初始化->init函数执行->main函数执行->defer函数执行
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
		"bool_type":   bool_type,
		"number_type": number_type,
		"string_type": string_type,
		"other_type":  other_type,
		"type1":       type1,
		"type2":       type2,
	}
	funs[n]()
}

/*
//
	int int8 int16 int32(rune) int64

	uint8(byte) uint16 uint32 uint64

	float32 float64 complex64 complex128

    rune:
		unicode code point
		(表示unicode编码的int值,相当于java中的char,go本身是utf-8编码的,1个unicode=2个字节),
		alias for int32
		[]rune(s string) 可以将s转为 unicode code point
    byte:
		raw data, alias for uint8
占用内存
1 int8  = 1 uint8 = 1 byte = 8 bit
1 int16 = 1 uint16 = 2 byte
1 int32 = 1 uint32 = 4 byte
1 int64 = 8 byte
*/
func number_type() {
	var i1 int = 1<<31 - 1   //int32                                    [-2^31,2^31-1]
	var i4 int32 = 1<<31 - 1 //-2147483648~2147483647                   [-2^31,2^31-1]
	var ui5 rune = 7         //int32                  [-2^31,2^31-1]

	var i2 int8 = 1<<7 - 1   //-128~127                                 [-2^7,2^7-1]
	var i3 int16 = 1<<15 - 1 //-32768~32767                             [-2^15,2^15-1]
	var i5 int64 = 1<<63 - 1 //-9223372036854775808~9223372036854775807 [-2^63,2^63-1]
	fmt.Println(i1, i2, i3, i4, i5)

	var ui1 uint8 = 1 << 8 //0~255(2^8)
	var ui2 byte = 1 << 8  //0~255

	var ui3 uint16 = 6 //0~65535                [0,2^16]
	var ui4 uint32 = 7 //0~4294967295           [0,2^32]
	var ui6 uint64 = 8 //0~18446744073709551615 [0,2^64]
	fmt.Println(ui1, ui2, ui3, ui4, ui5, ui6)

	var f1 float32 = 0.1    //IEEE-754 32位实数(小数点后6位)
	var f2 float64 = 0.2    //IEEE-754 64位虚数(小数点后15位)
	var f3 complex64 = 0.3  //32位实数和虚数
	var f4 complex128 = 0.4 //64位实数和虚数
	fmt.Println(f1, f2, f3, f4)
}

func bool_type() {
	var b1 bool = true
	var b2 bool = false

	if b1 {
		fmt.Println(b1)
	}

	if !b2 {
		fmt.Println(b2)
	}
}

func string_type() {
	var s1 string = "string1"
	fmt.Println(s1)
}

/*
 pointer,array,struct,
 channel,函数类型,切片类型,interface,map
*/
func other_type() {
	//pointer details see variable_pointer.go for reference
	var i int = 1
	fmt.Println(&i)

	//array [length]type{val,...} 数组长度不可变,不可追加元素(不可使用append())
	var a1 = [...]string{"a", "b"}
	for i := 0; i < len(a1); i++ {
		fmt.Println(a1[i])
	}

	//slice 切片(长度可变的数组)
	//见variable_slice.go

	//struct
	type obj struct {
		name string
		age  int
	}
	var s3 obj = obj{"candyxu", 40}
	s1 := obj{age: 12, name: "v_kenqzhang"}
	s2 := obj{"v_kenqzhang1", 12}
	fmt.Println(s1, s2, s2.name, s3)

	//channel
	//不带长度的管道(双向,<-,->单向)
	/*var ch1 chan string;
		var ch2 chan int;
		var chs []chan string;
		ch4 := make(chan string);
		ch5 := make(chan int);
	    ch1<- "1";
	    ch2<- 1;
	    chs[0]<- "1";

		//带长度的管道
		ch6 := make([]chan string,5);*/
}

func type1() {
	s := "中文检查"
	rs := []rune(s)
	fmt.Println(
		len(rs),
		string(rs[:2]),
	) //2 中文

	fmt.Println(
		len(s),
		s[:2],
	) //12(3*4) ??
}
func type2() {
	var a float64
	a = 12.2
	fmt.Println(fmt.Sprint(a))
}
