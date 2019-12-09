package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
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
		"Md5En" : Md5En,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

//sha1
func enc1()  {
	SHA1  := sha1.New()
	Msg   :=  "test"
	MsgBy := []byte(Msg)
	Key   := "key"
	KeyBy := []byte(Key)

	//加密
	_,err := SHA1.Write(MsgBy)
	if err != nil {
		fmt.Println(err)
		return
	}
	checksum    := SHA1.Sum(nil)
	checksumStr := fmt.Sprintf("%x",checksum)
	fmt.Println(checksumStr)


	//hmac的sha1
	HmacWithSha1 := hmac.New(sha1.New,KeyBy)
	//HmacWithSha1.Write(MsgBy)
	checksum1    := HmacWithSha1.Sum(nil)
	checksum1Str := fmt.Sprintf("%x",checksum1)
	fmt.Println(checksum1Str)
}

/**
https://studygolang.com/articles/2283
md5加密
*/
func Md5En()  {
	msg := "121"
	m := md5.New()
	m.Write([]byte(msg))
	by := m.Sum(nil)
	fmt.Println(hex.EncodeToString(by))
}
