package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.9.md
json
*/
func init() {
	fmt.Println("Content-Type:application/json;charset=utf-8\n\n")
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
		"json1" : json1,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

type Address struct {
	Type    string
	City    string
	Country string
}
type VCard struct {
	FName string
	LName string
	Address []*Address
	Remark string
}
func json1()  {
	addresses := []*Address{
		&Address{"T1","C1","CO1"},
		&Address{"T2","C2","co2"},
	}
	vc   := VCard{"Zhang","Sea",addresses,"re"}
	js,_ := json.Marshal(vc)
	fmt.Println(string(js))

	//json写入文件
	file, _ := os.OpenFile("t.json",os.O_CREATE|os.O_WRONLY,0666)
	defer file.Close()

	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		fmt.Println(err)
	}
}