package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/cgi"
	"net/http/fcgi"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"

	"reflect"
	"strings"
	//"log"
	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
)

type FastCGI struct{}

func init() {
	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
}
func main() {
	args := os.Args
	if len(args) < 1 {
		fmt.Println("lack func u want to run.")
		return
	}
	execute(args[1])
}

func execute(func1 string) {
	funs := map[string]func(){
		"net1": net1,
		"net2": net2,
		"net3": net3,
		"net4": net4,
		"net5": net5,
	}
	funs[func1]()
}

type addr struct {
	host string
	port string
}

func net1() {
	/*
	 func JoinHostPort(host,port string) string
	 返回由host,port组成的地址
	*/
	add := addr{
		host: "golang.org",
		port: "80",
	}
	address := net.JoinHostPort(add.host, add.port)
	fmt.Println(address)

}

/*
https://www.infoq.cn/article/golang-standard-library-part02
CGI服务器
*/
func net2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handle     := new(cgi.Handler)
		handle.Path = "/usr/local/go/bin/go"
		handle.Dir  = "/root/www/sea/local_go/src"
		script     := "/root/www/sea/local_go/src" + r.URL.Path

		args       := []string{"run", script}
		handle.Args = append(handle.Args, args...)
		handle.ServeHTTP(w, r)
	})
	http.ListenAndServe(":8989", nil)
}

/*
FastCGI服务器(创建后,需要在nginx中配置,参考php fastcgi 的配置内容)
*/
func net3() {
	listener, _ := net.Listen("tcp", "120.0.0.1:8989")
	srv         := new(FastCGI)
	fcgi.Serve(listener, srv)
	select {}
}

func (s *FastCGI) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello,fastcgi"))
}

/*
http 服务器
*/
func net4() {
	http.HandleFunc("/hello", hello)
	http.Handle("/handle/", http.HandlerFunc(say))

	go func() {
		http.ListenAndServe(":8989", nil)
	}()


	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal,os.Interrupt,syscall.SIGTERM)
	<-exitSignal
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello"))
}

type Handlers struct {
}

func (h *Handlers) ResAction(w http.ResponseWriter, req *http.Request) {
	fmt.Println("res")
	w.Write([]byte("res"))
}

func say(w http.ResponseWriter, req *http.Request) {
	pathInfo := strings.Trim(req.URL.Path, "/") //前后去/号
	parts    := strings.Split(pathInfo, "/")
	fmt.Println(strings.Join(parts, "|"))
	action   := ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	fmt.Println(action)
	handle := &Handlers{}

	//ValueOf 返回一个新值，初始化为存储在接口 i 中的具体值。ValueOf(nil) 返回零值。
	controller := reflect.ValueOf(handle)

	/*
		func (v Value) MethodByName(name string) Value
		MethodByName 返回与给定名称的v的方法对应的函数值。
		对返回函数调用的参数不应包含接收方; 返回的函数将始终使用 v 作为接收者。如果没有找到方法，它返回零值
	*/
	method := controller.MethodByName(action)
	r      := reflect.ValueOf(req)
	wr     := reflect.ValueOf(w)
	method.Call([]reflect.Value{wr, r})
}

//开启守护进程
//ref: https://phpjieshuo.com/archives/121/
//原理: 父进程开启一个子进程,然后父进程退出,那么此时子进程就变成了守护进程
func net5() {
	//需要守护的进程
	a := func() {
		log.Println(os.Getpid())
		select{}
	}

	//将其守护进程化
	daemonize1(a)
}

func daemonize1(process func())  {
	if os.Getenv("daemonize") == "1" {
		fmt.Println("child process running")
		process()
		return
	}
	sc, _     := filepath.Abs(os.Args[0])
	cmd       := exec.Command(sc,os.Args[1:]...)
	cmd.Stdin  = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = []string{"daemonize=1"}
	e     := cmd.Start()
	if e != nil {
		log.Fatal(e.Error())
	}
	fmt.Println("parent process finished")
}

func daemonize2(process func())  {
	if os.Getenv("daemonize") == "1" {
		fmt.Println("child process running")
		process()
		return
	}

	sc, _ := filepath.Abs(os.Args[0])
	cmd1 := &exec.Cmd{
		Path:         sc,
		Args:         append([]string{sc},os.Args[1:]...),
		Stdin:        os.Stdin,
		Stdout:       os.Stdout,
		Stderr:       os.Stderr,
		Env:[]string{"daemonize=1"},
	}
	e := cmd1.Start()
	if e != nil {
		log.Fatal(e.Error())
	}
	fmt.Println("i'm parent process finished.")
}

