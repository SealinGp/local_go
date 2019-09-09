package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

/*
string package
*/
func init() {
	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
}
func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("lack param ?func=xxx")
		return
	}

	execute(args[1])
}

func execute(n string) {
	funs := map[string]func(){
		"str1" : str1,
		"conv" : conv,
		"str2" : str2,
		"str3" : str3,
	}
	funs[n]()
}
func str1()  {
	str1 := " ,a,b,c,D,E,F, "
	fmt.Println(
		strings.HasPrefix(str1," ,"),
		strings.HasSuffix(str1,", "),
		strings.Contains(str1,"c,"),

		strings.Count(str1,"E"),

		strings.Repeat(str1,2),

		strings.ToLower(str1),
		strings.ToUpper(str1),

		strings.Replace(str1,"cD","Dc",-1),

		strings.TrimSpace(str1),
		strings.Trim(str1," "),
		strings.TrimLeft(str1," "),
		strings.TrimRight(str1," "),

		strings.Split(str1,","),
		strings.Join(strings.Split(str1,","),"_"),
	)

	//首个索引
	if Index := strings.Index(str1,"cD") ;-1 != Index {
		fmt.Println(Index)
	}
	//最后一个索引
	if Index := strings.LastIndex(str1,"cd"); -1 != Index {
		fmt.Println(Index)
	}
}

//与字符串相关的类型转换都是通过包strconv实现的
func conv()  {
	var f1 float64 = 2.3
	var s1 string  = "12"
	fmt.Println(
		//number -> string
		strconv.Itoa(5),
		strconv.FormatFloat(f1,'b',2,64),
	)
	//string -> int
	if sN, err := strconv.Atoi(s1);err != nil {
		fmt.Println(err)
		//终止程序的执行
		os.Exit(-1)
	} else {
		fmt.Println(sN)
	}


	//string -> float
	s2       := "12.3"
	if sN1, err := strconv.ParseFloat(s2,32); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sN1)
	}
}

func str2()  {
	//当前字母往后推,第14位
	rot13 := func(r rune) rune {
		var r1 rune
		switch {
			case r >= 'A' && r <= 'Z':
				r1 = 'A' + (r-'A'+13)%26
			case r >= 'a' && r <= 'z':
				r1 = 'a' + (r-'a'+13)%26
			default:
				r1 = r
		}
		return r1
	}
	fmt.Println(strings.Map(rot13, "Twas brillig and the slithy gopher..."))
}
func str3()  {
	fmt.Println(strings.SplitAfter("a,b,c",","))
}