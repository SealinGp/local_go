package main;

import (
	"fmt"
);
/*
*variable_type 指针变量声明 
&variable_name 引用变量的内存地址

 var variable_name *varibale_type = &variable_name; 指针变量指向变量地址,默认nill(空指针)
*/
func main() {
	var i1 int8  = 1;
	var i2 *int8 = &i1;		//pointer variable declare, defaule value : nil(空指针)	
	//i2 = &i1 = 指针所指的内存位置

	fmt.Printf("point address:%x \n",&i1);
	fmt.Printf("point address:%x \n",i2);
	fmt.Printf("point                      address value:%d \n",*i2);

	strua := stru1{name:"v_kenqzhang",age:123};
	var strub *stru1 = &strua;
	// strub := &strua;

	fmt.Printf("strua:%s \n",strub.name);
	fmt.Printf("strua:%d \n",strub.age);
}

type stru1 struct {
	name string
	age uint8
}