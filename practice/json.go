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
var jsonFuncs = map[string]func(){
	"json1": json1,
	"json2": json2,
}

type Address struct {
	Type    string
	City    string
	Country string
}
type VCard struct {
	FName   string
	LName   string
	Address []*Address
	Remark  string
}

func json1() {
	addresses := []*Address{
		&Address{"T1", "C1", "CO1"},
		&Address{"T2", "C2", "co2"},
	}
	vc := VCard{"Zhang", "Sea", addresses, "re"}

	//json_encode并输出
	js, _ := json.Marshal(vc)
	fmt.Println(vc)
	fmt.Println(string(js))

	//json_decode
	var v VCard
	json.Unmarshal(js, &v)
	for _, add := range v.Address {
		fmt.Println(add)
	}
	fmt.Println(*v.Address[0])

	//json_encode并写入文件
	file, _ := os.OpenFile("t.deb", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	//encoder json写入文件,decoder读取文件json数据
	//将json_encode写入文件流
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		fmt.Println(err)
	}
}

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.9.md

json 与 go 类型对应
bool    对应 JSON 的 boolean
float64 对应 JSON 的 number
string  对应 JSON 的 string
nil     对应 JSON 的 null
*/
func json2() {
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)

	//在不知道json结构的情况下进行decode,并判断类型
	//var j map[string]interface{}
	j := make(map[string]interface{})
	err := json.Unmarshal(b, &j)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//类型断言,t = v 的内容,t的类型在case中
	for k, v := range j {
		switch t := v.(type) {
		case string:
			fmt.Println("string:", k, v)
		case float64:
			fmt.Println("float64", k, v)
		case int:
			fmt.Println("int", k, v)
		case []string:
			fmt.Println("[]string:", k, v)
		case []interface{}:
			fmt.Print("[]interface: ", k)
			//正确的循环
			for _, v2 := range t {
				fmt.Print(" ", v2)
			}
			fmt.Println("")
			//错误的循环
			//for k1,v2 := range v  {
			//
			//}
		default:
			fmt.Println("default", k, v, t)
		}
	}

	//知道json结构的情况下可以定义一个struct
	type FamilyMember struct {
		Name    string
		Age     int
		Parents []string
	}
	fm := FamilyMember{}
	json.Unmarshal(b, &fm)
	fmt.Println(fm)
}
