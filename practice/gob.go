package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.11.md
适用于两服务端均为 基于go 的 RPCs服务 之间的数据传输
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
		"gob1": gob1,
		"gob2": gob2,
		"gob3": gob3,
	}
	if nil == funs[n] {
		fmt.Println("func", n, "unregistered")
		return
	}
	funs[n]()
}

type P struct {
	X, Y, Z int
	Name    string
}
type Q struct {
	X, Y, Z *int32
	Name    string
}

func gob1() {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network) //write
	dec := gob.NewDecoder(&network) //read

	err := enc.Encode(P{3, 4, 5, "pythagoras"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(q.Name, *q.X, *q.Y, *q.Z)
}

type Address struct {
	T  string
	C  string
	Co string
}
type VCard struct {
	FN string
	LN string
	Ad []*Address
	Re string
}

var content string

func gob2() {
	vc := VCard{
		FN: "Jan",
		LN: "Ker",
		Ad: []*Address{
			&Address{"t1", "c1", "co1"},
			&Address{"t2", "c2", "co2"},
		},
		Re: "none",
	}
	file, _ := os.OpenFile("vcard.deb", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ok")
}

func gob3() {
	var vc VCard
	file, err := os.Open("vcard.deb")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	dec := gob.NewDecoder(file)
	err = dec.Decode(&vc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(
		vc, vc.Ad[0],
	)
}
