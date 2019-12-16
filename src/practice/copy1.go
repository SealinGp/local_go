package main

import (
	"fmt"
	"github.com/jinzhu/copier"
	"os"
	"time"
)

/*
7.array slice []bytes string
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
		"cop1" : cop1,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

type cops1 struct {
	A time.Time
	B string
}
type cops2 struct {
	A string
	B string
}
func cop1()  {
	c1 := cops1{
		A:time.Now(),
		B:"b1",
	}
	c2 := cops2{
		A:"2014-12-16 00:00:00",
		B:"b2",
	}
	err := copier.Copy(&c1,&c2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(c1)
}