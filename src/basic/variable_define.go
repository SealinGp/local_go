package main;

import (
	"fmt"
	"os"
);

func init() {
  fmt.Println("Content-Type:text/plain;charset=utf-8\n\n");
}
func main() {
	args := os.Args;
    if len(args) <= 1 {
    	fmt.Println("url lack param ?func=xxx");
    	return;
    }

	execute(args[1]);
}
func execute(n string) {
	funs := map[string]func() {
		"define1"   : define1,
	};	
	funs[n]();
}
func define1() {
	//variable
	var i1 int8 = 1;
	var i2 	    = 1;
	i3 		   := 1;
	//这种声明一般用于全局变量
	var (
		i4 int8 = 2
		i5 int8 = 2
	);

	//const variable
	const s1 string = "abc";
	const s2 = "abc";
	const (
		s3 = "abc"
		s4 string = "abc"
	)

	fmt.Println(i1,i2,i3,i4,i5,s1,s2,s3,s4);
}