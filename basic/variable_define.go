package main;

import (
	"fmt"
);


func main() {
	//variable
	var i1 int8 = 1;
	var i2 	    = 1;
	i3 		   := 1;

	//const variable
	const s1 string = "abc";
	const s2 = "abc";
	const (
		s3 = "abc"
		s4 = "abc"
	)

	fmt.Println(i1,i2,i3,s1,s2,s3,s4);
}