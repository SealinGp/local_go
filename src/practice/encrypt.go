package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.12.md
数据加密
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
		"enc1" : enc1,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

//sha1
func enc1()  {
	//加密test
	ha := sha1.New()
	io.WriteString(ha,"test") //等价于 ha.Write([]byte("test"))
	b := []byte{}
	s := ha.Sum(b)
	fmt.Printf("%x\n", s)
	fmt.Printf("%d\n", ha.Sum(b))

	ha.Reset()
	data := []byte("we shall overcome")
	n,err := ha.Write(data)
	if n != len(data) || err != nil {
		fmt.Println(err)
		return
	}
	checksum := ha.Sum(b)
	fmt.Printf("%x",checksum)
}
