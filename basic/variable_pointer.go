package main

import (
	"fmt"
)

func main() {
	var i1 int8  = 1;
	var i2 *int8 = &i1;		//pointer variable declare, defaule value : nil(空指针)	
	fmt.Printf("point address:%x \n",i2);
	fmt.Printf("point address value:%d \n",*i2);	
}