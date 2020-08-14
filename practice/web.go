package main

import (
	"encoding/xml"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"fmt"
)

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
		"web1"  : web1,
		"web2"  : web2,
		"web3"  : web3,
		"web4"  : web4,
		"web5"  : web5,
		"web6"  : web6,
		"web7"  : web7,
		"web8"  : web8,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

const (
	host = "127.0.0.1"
	port = "8081"
)

func web1()  {
	http.HandleFunc("/",web1_HelloServer)
	addr := net.JoinHostPort(host,port)
	fmt.Println("listening..."+addr)
	err  := http.ListenAndServe(addr,nil)
	if err != nil {
		log.Fatal("listenAndServe Error:",err.Error())
	}
}
func web1_HelloServer(w http.ResponseWriter,r *http.Request)  {
	//fmt.Println("HelloServer func")
	_,_ = fmt.Fprint(w,"Hello :",r.URL.Path[1:])
}

type hello struct {}
func (hello)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprint(w,"Hello:",r.URL.Path[1:])
}
func web2()  {
	addr := net.JoinHostPort(host,port)
	err  := http.ListenAndServe(addr,hello{})
	if err != nil {
		log.Fatal(err)
	}
}

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.3.md
var (
	urls = []string{
		"http://www.google.com/",
		"http://golang.org/",
		"http://blog.golang.org/",
	}
)
func web3()  {
	for _,url := range urls {
		resp,err := http.Head(url)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(resp.Status)
		}
	}
}

func web4()  {
	res,err := http.Get("http://www.google.com")
	wcheckError(err)
	data,err := ioutil.ReadAll(res.Body)
	wcheckError(err) 
	fmt.Println(string(data))
}

//xml
type Status struct {
	Text string
}
type User struct {
	XMLName xml.Name
	Status Status
}
func web5()  {
	resp,_ := http.Get("http://twitter.com/users/Googland.xml")
	u := User{xml.Name{"","user"},Status{""}}
	data,err := ioutil.ReadAll(resp.Body)
	wcheckError(err)
	xml.Unmarshal(data,&u)
}

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.4.md
//一个简单的网页应用
const formTpl = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="submit" value="submit"/>
		</form>
	</body></html>
`
func web6()  {
	http.HandleFunc("/form",formHandle)
	if err := http.ListenAndServe(":3005",nil);err != nil {
		panic(err)
	}
}
func formHandle(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w,formTpl)
	case "POST":
		io.WriteString(w,"in:" + r.FormValue("in"))
	}
}

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.5.md
//参考13.5节中的闭包抓取panic错误来使得网页应用 防止跌机,挂掉
func web7()  {
	http.HandleFunc("/",panicCatch(web7_1handle))
	if err := http.ListenAndServe(":3005",nil);err != nil {
		log.Fatal(err.Error())
	}
}
func web7_1handle(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	str := "hi ,you access:" + r.URL.Path
	w.Write([]byte(str))
}
func panicCatch(han func(w http.ResponseWriter,r *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover();err != nil {
				fmt.Println(request.RemoteAddr,"panic catched!",err)
			}
		}()
		han(writer,request)
	}
}


func wcheckError(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.7.md
//探索template pkg
type Person struct {
	Name string
	NonExportedAgeField string
}
func web8()  {
	t   := template.New("hello")
	t,_  = t.Parse(fmt.Sprintln("hello {{.Name}}{{.NonExportedAgeField}}"))
	per := Person{"Sea","Zhang"}
	if err := t.Execute(os.Stdout,per) ;err != nil {
		fmt.Println(err.Error())
	}
}