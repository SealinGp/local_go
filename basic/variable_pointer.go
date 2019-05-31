package main;

import (
	"os"
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
	args := os.Args;
    if len(args) <= 1 {
    	fmt.Println("lack param ?func=xxx");
    	return;
    }

	execute(args[1]);
}
func execute(n string) {
	funs := map[string]func() {
		"pointer1" : pointer1,
		"pointer2" : pointer2,
		"pointer3" : pointer3,
	};	
	funs[n]();		
}

// pointer basic
func pointer1() {
	var i1 int8         = 1;
	var i1Pointer *int8 = &i1;		//pointer variable declare, defaule value : nil(空指针)

	//i1Pointer = &i1 = 指针所指的内存位置
    fmt.Println("i1:",i1);
    fmt.Println("i1Pointer:",i1Pointer);
    fmt.Println("*i1Pointer:",*i1Pointer);
    fmt.Println("*i1Pointer++");
    *i1Pointer++;
    fmt.Println("i1:",i1);
    fmt.Println("--------------------");

	
	type stru1 struct {
		name string
		age uint8
	};	
    stru        := stru1{name:"v_kenqzhang",age:123};
    struPointer := &stru; 
	fmt.Println("stru:",stru);;	
	fmt.Println("struPointer:",struPointer);
	fmt.Println("struPointer.name:",struPointer.name);
	struPointer.name = "v_sshyu";
	fmt.Println("*struPointer:",*struPointer);
	fmt.Println("stru:",stru);

	type stru1_1 struct {
		name *string
		age uint8
	}
	str          := "v_kenqzhang";
	stru1_       := stru1_1{name:&str,age:22};
	stru1Pointer := stru1_;
	fmt.Println("stru1_",stru1_);	
	fmt.Println("stru1Pointer",stru1Pointer);	
	*stru1_.name = "v_sshyu";
	fmt.Println("stru1Pointer.name",*stru1Pointer.name);
}



/*
指针使用:*xxx 指针声明&xxx

*/
type stru2 struct {
	a string
	b string
};
type interface1 interface {
	exchange();
}
func (st *stru2)exchange() {
	mid := st.a;
	st.a = st.b;
	st.b = mid;	
}
func pointer2() {
	st := &stru2{"a","b"};
	st.exchange();
	fmt.Println(*st);
	fmt.Println(st.a);

	st1 := stru2{"a","b"};
	st1.exchange();
	fmt.Println(st1);
	fmt.Println(st1.a);
}

type stru3 struct {
	a string
	b string
};
type interface2 interface {
	exchange2();
}
func (st stru3)exchange2() (stru3) {
	mid := st.a;
	st.a = st.b;
	st.b = mid;	
	return st;
}
func pointer3() {
	st := stru3{"a","b"};
	st1 := st.exchange2();
	fmt.Println(st);
	fmt.Println(st1);
}