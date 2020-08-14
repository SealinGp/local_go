package basic
/*
Go汇编语言变量定义
https://github.com/chai2010/advanced-go-programming-book/blob/master/ch3-asm/ch3-03-const-and-var.md

Go汇编语言定义变量无法指定类型信息,因此需要先通过Go语言声明变量的类型.
以下是在Go语言中声明的几个bool类型变量.
在Go语言中声明的变量不能含有初始化语句,对应pkg1_amd64.s为汇编环境定义对应的变量初始化

常量:
$1          //10进制
$0xf4f8fcff //16进制
$1.5        //浮点
$'a'        //字符
$"abc"      //字符串
VarName<> :表示私有变量
flags     :变量标识符(DUPOK|RODATA(只读)|NOPTR(无指针))
DUPOK表示该变量对应的标识符可能有多个，在链接时只选择其中一个即可（一般用于合并相同的常量字符串，减少重复数据占用的空间

定义全局变量:
1.声明变量名,变量大小
GLOBL ·VarName(SB),[flags...,]$width
2.声明变量值,变量地址
DATA ·VarName+offset(SB)/width,$VarVal

*/
var (
	Id int
	Num [2]int
	BoolValue bool
	TrueValue bool
	FalseValue bool
	Int32Value int32
	Uint32Value uint32
	Float32Value float32
	Float64Value float64
	M map[string]int
	Ch chan int
	ReadOnlyInt int
	Name string
	NameData string
)