package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.10.md
xml

<a at='av'>a1</a>

TagName:标签名  (a)
TagVal:标签值   (a1)
AttrName:属性名 (at)
AttrVal :属性值 (av)
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
		"xml1": xml1,
		"xml2": xml2,
	}
	if nil == funs[n] {
		fmt.Println("func", n, "unregistered")
		return
	}
	funs[n]()
}

var (
	x = "<Person a='abc'>" +
		"<FirstName>" + "Laura" + "</FirstName>" +
		"<LastName>" + "Lynn" + "</LastName>" +
		"<LN>" + "LN123" + "</LN>" +
		"</Persion>"
)

func xml1() {
	r := strings.NewReader(x)
	xd := xml.NewDecoder(r)
	for t, err := xd.Token(); err == nil; t, err = xd.Token() {
		switch token := t.(type) {
		//开始标签
		case xml.StartElement:

			//开始标签名
			fmt.Println("SE:", token.Name.Local)

			//标签属性
			for _, v := range token.Attr {
				fmt.Println(
					"attr :", v.Name.Local,
					"value:", v.Value,
				)
			}

		//结束标签
		case xml.EndElement:
			fmt.Println("EE:")

		//标签包裹的内容
		case xml.CharData:
			content := string([]byte(token))
			fmt.Println("CD:", content)

		//其他
		default:
			fmt.Println("DF:")
		}
	}
}

func xml2() {
	find := "Person/LN"
	val := getTagValByTagName(find, x, "/")
	fmt.Println(val)
}
func getTagValByTagName(tag, xmls, sep string) (val string) {
	r := strings.NewReader(xmls)
	xd := xml.NewDecoder(r)

	find1 := strings.Split(tag, sep)
	le := len(find1)
	i := 0
	if le <= 0 {
		return
	}
	for t, err := xd.Token(); err == nil; t, err = xd.Token() {
		switch token := t.(type) {
		//开始标签
		case xml.StartElement:
			if i < le && token.Name.Local == find1[i] {
				i++
			}
		//结束标签
		case xml.EndElement:
		//标签包裹的内容
		case xml.CharData:
			content := string([]byte(token))
			if i == le {
				val = content
				break
			}
		//其他
		default:
		}
	}

	return
}
