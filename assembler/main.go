package main

/*
Go 汇编
https://github.com/chai2010/advanced-go-programming-book/blob/master/ch3-asm/ch3-01-basic.md
https://chai2010.cn/advanced-go-programming-book/ch3-asm/ch3-01-basic.html

常量:
$1          //10进制
$0xf4f8fcff //16进制
$1.5        //浮点
$'a'        //字符
$"abc"      //字符串
VarName<> :表示私有变量
flags     :变量标识符(DUPOK|RODATA(只读)|NOPTR(无指针))
DUPOK表示该变量对应的标识符可能有多个，在链接时只选择其中一个即可（一般用于合并相同的常量字符串，减少重复数据占用的空间
width : 内存地址宽度

常用寄存器: 16位 32位 64位
累加寄存器: AX EAX RAX
基址寄存器: BX EBX RBX
计数寄存器: CX ECX RCX
数据寄存器: DX EDX RDX
堆栈基指针: BP EBP RBP
变址寄存器: SI ESI RSI
堆栈顶指针: SP ESP RSP
指令寄存器: IP EIP RIP

指令:
movb(8位) movw(16位) movl(32位) movq(64位)

定义全局变量:
1.声明变量名,变量大小
GLOBL ·VarName(SB),[flags...,]$width
2.声明变量值,变量地址
DATA ·VarName+offset(SB)/width,$VarVal

函数定义:
TEXT指令 ·函数名,flags标志(可选) 函数帧大小(可选参数大小)
TEXT symbol(SB),[flags,] $framesize(-argsize)
flags:
  NOSPLIT(不会生成或包含栈分裂代码)
  WRAPPER(这个是一个包装函数，在panic或runtime.caller等某些处理函数帧的地方不会增加函数帧计数)
  NEEDCTXT(上下文参数,一般用于闭包函数)

*/
import (
	"fmt"

	"github.com/SealinGp/local_go/assembler/basic"
)

func main() {
	println(basic.Id)
	fmt.Println("Num:", basic.Num)
	fmt.Println("BoolValue:", basic.BoolValue)
	fmt.Println("TrueValue", basic.TrueValue)
	fmt.Println("FalseValue:", basic.FalseValue)
	fmt.Println("Int32Value:", basic.Int32Value)
	fmt.Println("Uint32Value:", basic.Uint32Value)
	fmt.Println("Float32Value:", basic.Float32Value)
	fmt.Println("Float64Value:", basic.Float64Value)
	fmt.Println("ReadOnlyInt:", basic.ReadOnlyInt)
	fmt.Println("Name:", basic.Name)
}
