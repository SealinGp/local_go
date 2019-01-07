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

func main() {

	// bool_type();

	// number_type();	

	// string_type();

	other_type();
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
	var ui2 byte   = 5; 	//uint8
	var ui3 uint16 = 6; 	//0~65535
	var ui4 uint32 = 7;		//0~4294967295
	var ui5 rune   = 7;		//uint32

	var ui6 uint64 = 8;		//0~18446744073709551615
	fmt.Println(ui1,ui2,ui3,ui4,ui5,ui6);

	var f1 float32    = 0.1;//IEEE-754 32位实数
	var f2 float64    = 0.2;//IEEE-754 64位虚数
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

	if b2 {
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
 	
 	//array [length]type{val,...}
	var a1 = [...]string{"a","b"};
	fmt.Println(a1[0]);

	//struct
	s1 := stru{age:12,name:"v_kenqzhang"};
	s2 := stru{"v_kenqzhang1",12};
	fmt.Println(s1,s2,s2.name);	

	//channel	
}
type stru struct {
	name string
	age int
}
