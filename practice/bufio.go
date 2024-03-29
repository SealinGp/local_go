package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.2.md
缓冲 读写数据(包含命令行,文件读写)
写文件
 程序->内存缓冲->文件
读文件
 程序<-内存缓冲<-文件

from:https://zhuanlan.zhihu.com/p/73690883
bufio是通过缓冲来提高效率,bufio提供了缓冲区(分配一块内存),
读写均在缓冲区中,最后在读写文件,降低访问本地磁盘的次数,提高效率

程序读取的内容大小:B1 bytes
缓冲区大小:B2 bytes
文件大小:B3 bytes
读:
 B2 >= B1时:
 文件->缓冲区(读取B2大小的内容)->程序(接收B1大小的内容)
 缓冲区剩余:(B2-B1)bytes
 B2 < B1时:
 文件->程序(读取B1大小的内容)

程序写入的内容大小:B1 bytes
写:
 B2 >= B1时:
 程序->缓冲区(写入B1大小的内容)->文件(写入B1大小的内容)
 B2 < B1时:
 程序->文件(写入B1大小内容)

可以用于做一个简单的项目结构脚手架
*/
var buifoFuncs = map[string]func(){
	"buf1": buf1,
	"buf2": buf2,
	"buf3": buf3,
	"buf4": buf4,
	"buf5": buf5,
	"buf6": buf6,
	"buf7": buf7,
	"buf8": buf8,
}

//读取用户输入方法1
func buf1() {
	var (
		firstName, lastName string
		//s string
		//i int
		//f32 float32
		//input  = "56.12 / 5212 / Go"
		//format = "%f / %d / %s"
	)

	fmt.Println("enter full name:")
	//扫描来自标准输入的文本,将空格分隔的值一次存放到后续的参数内,直到碰到换行
	fmt.Scanln(&firstName, &lastName)

	fmt.Println("hi", firstName, lastName)

	//?
	//fmt.Sscanf(input,format,&f32,&i,&s)
	//fmt.Println(f32,i,s)
}

//读取用户输入方法2:缓冲读取
func buf2() {
	var (
		inputReader *bufio.Reader
		input       string
		err         error
	)
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("enter some input:")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("input is", input)
	}
}

func buf3() {
	/*
		linux下
		\n : 换行(回车)
		\r : 空格
	*/
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("enter input:")
	input, err := inputReader.ReadString('S')

	if err != nil {
		fmt.Println("sry,procedure error!msg:", err)
		return
	}

	fmt.Println("input len:", len(input)-strings.Count(input, "\r")-strings.Count(input, "\n"))

	//计算每一行 以空格分隔的有多少个单词
	a := strings.Split(input, "\n")
	words := make([]string, 0)
	for _, v := range a {
		word := strings.Split(v, " ")
		words = append(words, word...)
	}
	fmt.Println("words len:", len(words), words)

	fmt.Println("lines len:", strings.Count(input, "\n")+1)
}

/*
文件读取

文件句柄 os.File{}
标准输入 os.Stdin
标准输出 os.Stdout
*/
func buf4() {
	//打开文件句柄 inputF 为 *os.File
	inputF, inputErr := os.Open("array.go")
	if inputErr != nil {
		fmt.Println(inputErr)
		return
	}
	defer inputF.Close()

	inputRe := bufio.NewReader(inputF)
	for {
		//碰到'\n'(回车符)为标识符,算一行
		inputStr, readErr := inputRe.ReadString('\n')
		fmt.Println(inputStr)

		//判断读到文件末尾跳出
		if readErr == io.EOF {
			break
		}
	}
}

//带缓冲的读取
func buf5() {
	inputF, inputErr := os.Open("a.deb")
	if inputErr != nil {
		fmt.Println(inputErr)
		return
	}
	defer inputF.Close()

	inputRe := bufio.NewReader(inputF)
	return
	buf := make([]byte, 2048)
	for {
		n, err := inputRe.Read(buf)
		//当读取到末尾的时候,err会为EOF 并且 n 为 0
		if err != nil && err != io.EOF {
			fmt.Println("error:", err)
			break
		}
		if n == 0 {
			break
		}
	}

	//读取出来的顺序是乱的
	fmt.Println(string(buf))
}

//整个读取出来
func buf6() {
	//读取
	buf, err := ioutil.ReadFile("array.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	//写出到其他文件
	err = ioutil.WriteFile("a.go", buf, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

// https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.2.md
// csv文件读取
// 练习
func buf7() {
	file, err := os.Open("product.deb")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	type product struct {
		title    string
		price    float64
		quantity int
	}
	var products []product
	for {
		rowString, err := reader.ReadString('\n')
		pro := product{}
		str := strings.Split(rowString, ";")
		if len(str) >= 3 {
			pro.title = str[0]
			if v, er := strconv.ParseFloat(str[1], 64); er == nil {
				pro.price = v
			}
			if v, er := strconv.Atoi(str[2]); er == nil {
				pro.quantity = v
			}
		}

		products = append(products, pro)
		if err != nil && err != io.EOF {
			fmt.Println("error:", err)
			break
		}
		if err == io.EOF {
			break
		}
	}

	fmt.Println(products)
}
func buf8() {
	buf := bytes.Buffer{}
	for i := 0; i < 5; i++ {
		buf.WriteString(strconv.Itoa(i))
	}
	s := buf.String()
	fmt.Println(s)
	buf.Reset()
	fmt.Println(buf.String())
}
