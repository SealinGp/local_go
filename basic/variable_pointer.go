package main;

import (
	"fmt"
);
/*
*variable_type 指针变量声明 
&variable_name 引用变量的内存地址

 var variable_name *varibale_type = &variable_name; 指针变量指向变量地址,默认nill(空指针)
*/
func init() {
  fmt.Println("Content-Type:text/plain;charset=utf-8\n\n");
}
func main() {
	var i1 int8  = 1;
	var i2 *int8 = &i1;		//pointer variable declare, defaule value : nil(空指针)

	//i2 = &i1 = 指针所指的内存位置
    fmt.Println("i1:",i1);
    fmt.Println("&i1:",&i1);
    fmt.Println("i2:",i2);
    fmt.Println("*i2:",*i2);
    fmt.Println("--------------------");
	
	type stru1 struct {
		name string
		age uint8
	}
	// var stru stru1         = stru1{name:"v_kenqzhang",age:123};
	// var struPointer *stru1 = &stru;
    stru        := stru1{name:"v_kenqzhang",age:123};
    struPointer := &stru;
	fmt.Println("stru:",stru);
	fmt.Println("&stru:",&stru);
	fmt.Println("&stru.name:",&stru.name);
	fmt.Println("struPointer:",struPointer);
	fmt.Println("struPointer.name:",struPointer.name);
	fmt.Println("*struPointer:",*struPointer);
}