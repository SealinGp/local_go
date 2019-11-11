package main

import (
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
	fmt.Println("HelloServer func")
	fmt.Fprint(w,"Hello :",r.URL.Path[1:])
}

type web2t struct {}
func (web2t)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprint(w,"Hello:",r.URL.Path[1:])
}
func web2()  {
	addr := net.JoinHostPort(host,port)
	err  := http.ListenAndServe(addr,web2t{})
	if err != nil {
		log.Fatal(err)
	}
}
