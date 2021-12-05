package main

import (
	"fmt"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/08.2.md
8.map
*/

var mapFuncs = map[string]func(){
	"map1": map1,
	"map2": map2,
}

/*
永远使用make来初始化Map
判断map中的key是否存在
*/
func map1() {
	//对于map类型来说,make的第二个参数是cap
	m1 := make(map[string]string, 10)
	m1["abc1"] = "abc1"
	m1["abc2"] = "abc2"

	val, ok := m1["abc1"]
	if !ok {
		fmt.Println("key not exists")
		return
	}
	fmt.Println(val, len(m1))

	//delete,若key不存在,不会出现错误
	delete(m1, "abc2")
}

//map reverse
func map2() {
	reverse := func(m map[string]string) (rm map[string]string) {
		//不Make会出错 nil map不可赋值
		rm = make(map[string]string, len(m))
		for k, v := range m {
			rm[v] = k
		}
		return
	}

	m1 := map[string]string{
		"a": "a1",
		"b": "b1",
		"c": "c1",
	}
	m2 := reverse(m1)
	fmt.Println(m2)
}
