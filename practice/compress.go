package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.2.md
读写数据到文件
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
		"com1": com1,
		"com2": com2,
		"com3": com3,
		"com4": com4,
		"com5": com5,
		"com6": com6,
	}
	if nil == funs[n] {
		fmt.Println("func", n, "unregistered")
		return
	}
	funs[n]()
}

//使用缓冲写入内容到文件
func com1() {
	/*
		     打开文件句柄
			os.O_WRONLY : 只写
			os.O_RDONLY : 只读
			os.O_CREATE : 文件不存在则创建
			os.O_TRUNC :  文件若存在则将该文件长度截为0
	*/
	outputF, err := os.OpenFile("output.deb", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputF.Close()

	//打开一个缓冲区
	writer := bufio.NewWriter(outputF)
	ouputStr := "hello asdasf"

	//缓冲区写入内容
	for i := 0; i < 10; i++ {
		writer.WriteString(ouputStr)
	}

	//清空缓冲区,冲出内容到文件
	writer.Flush()
}

//不使用缓冲写入内容到文件
func com2() {
	_, absPath, line, _ := runtime.Caller(1)
	f, _ := os.OpenFile("test.deb", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer f.Close()
	f.WriteString("abc1")
	fmt.Println(absPath, line)
}

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.2.md
//练习12.4
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() {
	f, e := os.OpenFile(p.Title, os.O_CREATE|os.O_WRONLY, 0666)
	if e != nil {
		os.Stdout.WriteString(e.Error())
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	w.Write(p.Body)

	e = w.Flush()
	if e != nil {
		os.Stdout.WriteString(e.Error())
		return
	}
}
func (page *Page) load() {
	buf, err := ioutil.ReadFile(page.Title)
	if err != nil {
		os.Stdout.WriteString(err.Error())
		return
	}

	os.Stdout.WriteString(string(buf) + "\n")
}
func com3() {
	p := Page{
		Title: "a.deb",
		Body:  []byte("adasfafsagd"),
	}

	p.save()
	p.load()
}

//文件拷贝
func com4() {
	w, err := copyF("b.deb", "a.deb")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(w)
}
func copyF(dst, src string) (written int64, err error) {
	/*//方法1
	//读取
	buf,err := ioutil.ReadFile(src)
	if err != nil {
		return
	}
	//写出到其他文件
	err = ioutil.WriteFile(dst,buf,0644)
	if err != nil {
		return
	}*/

	srcF, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcF.Close()

	dstF, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dstF.Close()

	return io.Copy(dstF, srcF)
}

//文件压缩
func com5() {
	e := Compress([]string{"array.go", "cmd.go"}, "1.zip")
	if e != nil {
		fmt.Println(e.Error())
	}
}
func com6() {
	DeCompress("1.zip")
}

func Compress(src []string, dst string) error {
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(dstFile)
	defer zw.Close()

	for _, v := range src {
		f, err := os.Stat(v)
		if err != nil {
			return err
		}

		zwc, err := zw.Create(f.Name())
		if err != nil {
			return err
		}
		contents, err := ioutil.ReadFile(v)
		_, err = zwc.Write(contents)
		if err != nil {
			return err
		}

	}
	return nil
}
func DeCompress(zipFile string) {
	r, _ := zip.OpenReader(zipFile)
	defer r.Close()
	for _, zf := range r.File {
		fmt.Println(zf.Name)
	}
}
