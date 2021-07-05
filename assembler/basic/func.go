package basic

/*
Go汇编语言函数
https://github.com/chai2010/advanced-go-programming-book/blob/master/ch3-asm/ch3-04-func.md

Go函数控制流
https://github.com/chai2010/advanced-go-programming-book/blob/master/ch3-asm/ch3-05-control-flow.md

函数定义:
TEXT symbol(SB),[flags,] $framesize(-argsize)

TEXT指令 ·函数名,flags标志(可选) 函数帧大小(可选参数大小)
函数的名字后面是(SB)，表示是函数名符号相对于SB伪寄存器的偏移量
flags:
  NOSPLIT(不会生成或包含栈分裂代码)
  WRAPPER(这个是一个包装函数，在panic或runtime.caller等某些处理函数帧的地方不会增加函数帧计数)
  NEEDCTXT(上下文参数,一般用于闭包函数)
*/
func Swap(a, b int) (int, int)
