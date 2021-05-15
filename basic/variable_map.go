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
		"map1":       map1,
		"map2":       map2,
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

type QmInfo struct {
	Name string
	Area string
}

//https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-for-range/
//1.map循环无序,k每次循环不一定相同
//2.把数组中的指针变量传递到另外一个map或者数组时,传递的是循环里面变量v的地址,而v的地址里面的值随着循环是会拷贝数组里面的元素进来的
func map1() {
	a := []QmInfo{
		{Name: "name1", Area: "area1"},
		{Name: "name2", Area: "area2"},
		{Name: "name3", Area: "area3"},
		{Name: "name4", Area: "area4"},
		{Name: "name5", Area: "area5"},
	}

	m1 := make(map[string]*QmInfo)
	m2 := make(map[string]*QmInfo)

	for i, v := range a {
		//把新变量v的地址给到m1[v.Name]
		m1[v.Name] = &v

		//数组元素变量的地址给到m2[v.Name]
		m2[v.Name] = &a[i]
	}

	//最终变量v的地址里面的值一定是数组里面的最后一个变量的值
	for k, v := range m1 {
		fmt.Println(k, "m1:", v, "---", "m2:", m2[k])
	}
}

//循环永动机?
func map2() {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr)
	return
}
