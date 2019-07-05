package main

import (
	"fmt"
	"os"
)

func init() {
  fmt.Println("Content-Type:text/plain;charset=utf-8\n\n");
}
func main() {
	args := os.Args;
    if len(args) <= 1 {
    	fmt.Println("lack param ?func=xxx");
    	return;
    }

	execute(args[1]);
}
func execute(n string) {
	funs := map[string]func() {
		"create_slice" : create_slice,
		"for_slice"  : for_slice,
		"fun_slice"  : fun_slice,
	};	
	funs[n]();		
}
func test() int {
	var ret int;	
	return ret;
}
var a int;
/*
slice:
 切片 = 动态数组(长度可变)
*/
//声明切片
func create_slice() {	
	//直接创建数组切片 []type,length,capability(maxLength) 长度 容量(最大长度(可选))
	var sli3 = make([]int,5,10);
	sli3[1]  = 3;
	fmt.Println(sli3);
	fmt.Println("----------------------");

	//基于切片创建切片
	var sli1 [2]int = [2]int{1,2};
	var sli2 = sli1[0:2]; //startIndex:length
	var sli21= sli1[:2];  //不填写为下限
	var sli22= sli1[0:];  //不填写为上限
	var sli23= sli1[:];   //不填写为上限
	sli2[0]  = 3;	
	fmt.Println(sli1);
	fmt.Println(sli2);
	fmt.Println(sli21);
	fmt.Println(sli22);
	fmt.Println(sli23);
}

//遍历切片foreach
func for_slice() {
	var sli1 [3]string = [3]string{"a","b","c"};	
	for i, v := range sli1 {
		fmt.Println(i,v);
	}
}

//切片函数
func fun_slice() {
	//len() cap()
	var sli1 = make([]string,3,5);
	sli1[0] = "abc";
	sli1[1] = "abc";
	sli1[2] = "abc";
	echo_str_slice(sli1);
	fmt.Println("---------------");

	//append(slice,val):入栈 copy()
	sli1 = append(sli1,"0");	
	echo_str_slice(sli1);

	var sli2 = make([]string,len(sli1),cap(sli1));
	copy(sli2,sli1);
	echo_str_slice(sli2);
}

func echo_str_slice(sli []string) {
	 fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli);
}
func echo_int_slice(sli []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli);
}