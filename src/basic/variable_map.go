package main

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
}
func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("url lack param ?func=xxx")
		return
	}

	execute(args[1])
}
func execute(n string) {
	funs := map[string]func(){
		"create_map": create_map,
		"for_map":    for_map,
		"fun_map":    fun_map,
	}
	funs[n]()
}

func create_map() {
	var m1 map[string]string = make(map[string]string)
	m1["name"] = "v_kenqzhang"

	m2 := make(map[string]string)
	m2["name"] = "sealingp"

	m3 := map[string]string{
		"name": "shamcleren",
	}

	fmt.Println(m1, m2, m3)
	seprate_line()
}

func for_map() {
	m1 := make(map[string]string)
	m1["k1"] = "v1"
	m1["k2"] = "v2"
	fmt.Println(m1)

	for k := range m1 {
		fmt.Println(k)
	}

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	for _, v := range m1 {
		fmt.Println(v)
	}
	seprate_line()
}

func fun_map() {
	//create
	m1 := make(map[string]string)
	m1["a"] = "b"
	m1["a1"] = "b1"

	//delete
	delete(m1, "a1")

	//update
	m1["a"] = "b1"

	/*
		check if exists
		这里的ok为类型断言
		map,interface,pointer,slice,func,chan 的默认值是 nil
	*/
	if v, ok := m1["a1"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println(ok)
	}

	fmt.Println(m1)
	seprate_line()
}

func seprate_line() {
	fmt.Println("--------------")
}
