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

var gobFuncs = map[string]func(){
	"gob1": gob1,
	"gob2": gob2,
	"gob3": gob3,
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

type Address1 struct {
	T  string
	C  string
	Co string
}
type VCard1 struct {
	FN string
	LN string
	Ad []*Address1
	Re string
}

var content string

func gob2() {
	vc := VCard1{
		FN: "Jan",
		LN: "Ker",
		Ad: []*Address1{
			&Address1{"t1", "c1", "co1"},
			&Address1{"t2", "c2", "co2"},
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
	var vc VCard1
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
