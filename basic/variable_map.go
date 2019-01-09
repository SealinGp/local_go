package main;

import (
	"fmt"
);

func main() {
	// create_map();

	// for_map();

	fun_map();
}

func create_map() {
	var m1 map[string]string = make(map[string]string);
	m1["name"] = "v_kenqzhang";

	m2 := make(map[string]string);
	m2 ["name"] = "sealingp";

	fmt.Println(m1);
	fmt.Println(m2);
	seprate_line();
}

func for_map() {
	m1 := make(map[string]string);
	m1["k1"] = "v1";
	m1["k2"] = "v2";

	for k := range m1 {
		fmt.Println(k);
	}

	for k,v := range m1 {
		fmt.Println(k,v);
	}
	seprate_line();
}

func fun_map() {
	//delete
	m1 := make(map[string]string);
	m1["a"]  = "b";
	m1["a1"] = "b1";

	//map_obj,key
	delete(m1,"a1");
	fmt.Println(m1);
	seprate_line();
}

func seprate_line() {
	fmt.Println("--------------");
}