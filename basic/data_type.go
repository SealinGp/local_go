package main;

import(
	"fmt"
);
/*
1 b(字节byte) = 8 bit(比特)
1 Kb          = 1024 b
1 Mb          = 1024 Kb
1 Gb          = 1024 Mb

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
*/
//执行顺序 全局变量初始化->init函数执行->main函数执行->defer函数执行
func init() {
  fmt.Println("Content-Type:text/plain;charset=utf-8\n\n");
}
func main() {
    
	// bool_type();

	number_type();	

	// string_type();

	// other_type();
}

/*
	int int8 int16 int32 int64

	uint8(byte) uint16 uint32(rune) uint64	

	float32 float64 complex64 complex128
*/
func number_type() {	
	var i1 int   = 1;		//int32
	var i2 int8  = 2;		//-128~127
	var i3 int16 = 3;   	//-32768~32767
	var i4 int32 = 3;		//-2147483648~2147483647
	var i5 int64 = 4;		//-9223372036854775808~9223372036854775807	
	fmt.Println(i1,i2,i3,i4,i5);

	var ui1 uint8  = 5; 	//0~255
	var ui2 byte   = 5; 	//0~255
	var ui3 uint16 = 6; 	//0~65535
	var ui4 uint32 = 7;		//0~4294967295
	var ui5 rune   = 7;		//uint32
	var ui6 uint64 = 8;		//0~18446744073709551615
	fmt.Println(ui1,ui2,ui3,ui4,ui5,ui6);

	var f1 float32    = 0.1;//IEEE-754 32位实数(小数点后6位)
	var f2 float64    = 0.2;//IEEE-754 64位虚数(小数点后15位)
	var f3 complex64  = 0.3;//32位实数和虚数
	var f4 complex128 = 0.4;//64位实数和虚数
	fmt.Println(f1,f2,f3,f4);
}

func bool_type() {
	var b1 bool = true;
	var b2 bool = false;

	if b1 {
		fmt.Println(b1);
	}

	if !b2 {
		fmt.Println(b2);
	}
}

func string_type() {	
	var s1 string = "string1";
	fmt.Println(s1);
}


/*
	 pointer,array,struct,
	 channel,函数类型,切片类型,interface,map	 
*/
func other_type() {
	//pointer details see variable_pointer.go for reference
	var i int = 1;
	fmt.Println(&i);
 	
 	//array [length]type{val,...} 数组长度不可变,不可追加元素(不可使用append())
	var a1 = [...]string{"a","b"};
	for i := 0; i < len(a1); i++ {
		fmt.Println(a1[i]);
	}

	//slice 切片(长度可变的数组)
	//见variable_slice.go
	

	//struct
	type obj struct {
		name string
		age int
	}
	s1 := obj{age:12,name:"v_kenqzhang"};
	s2 := obj{"v_kenqzhang1",12};
	fmt.Println(s1,s2,s2.name);	

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

