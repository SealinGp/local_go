package main

import (
	"fmt"
	"math"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/08.2.md
8.map
*/

var mathFuncs = map[string]func(){
	"mat1": mat1,
}

func mat1() {
	fmt.Println(math.Pow10(3))

	fmt.Println(math.Pow(-2, 2))

}
