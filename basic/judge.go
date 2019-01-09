package main;

import (
	"fmt"
);

func main() {
	// if_sentences();
	// for_sentences();
	switch_sentences();
}


func if_sentences() {	
	var b bool = true;
	var a bool = false;
	var c bool = true;
	if b || a {//true
		fmt.Println(b); //true
	}

	if b && a {//false
		fmt.Println(a);
	}

	if b &&
	   a &&
	   c {//false
		fmt.Println(c);
	}
}

func for_sentences() {
	for i := 0; i < 10; i++ {
		fmt.Println(i);
	}

	//while
	i := 10;	
	for i>0 {
		i--;
		fmt.Println(i);
	}
}

func switch_sentences() {
	condition := true;	
	switch condition != true {
		case false:
			fmt.Println(false);
		case true:
			fmt.Println(true);
		default:
			fmt.Println("error!");
	}
	return;
}