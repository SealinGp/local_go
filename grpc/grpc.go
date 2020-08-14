package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"os"
	"sync"
	"time"
)
//https://github.com/chai2010/advanced-go-programming-book/blob/master/ch4-rpc/ch4-01-rpc-intro.md

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("func required")
	}
	f := map[string]func(){
		"grpc1":grpc1,
		"grpc2":grpc2,
		"grpc3":grpc3,
		"grpc4":grpc4,
		"grpc5":grpc5,
		"grpc6":grpc6,
	}
	f[os.Args[1]]()
}

//https://github.com/chai2010/advanced-go-programming-book/blob/master/ch4-rpc/ch4-02-pb-intro.md


//https://github.com/chai2010/advanced-go-programming-book/blob/master/ch4-rpc/ch4-03-netrpc-hack.md
//achrive "watch" functional based on rpc
type KVStoreService struct {
	m map[string]string  //存储k-v数据库
	filter map[string]func(key string) //过滤器函数列表
	sm sync.Mutex
}
func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}
func (p *KVStoreService)Get(key string,value *string) error {
	p.sm.Lock()
	defer p.sm.Unlock()

	if v,ok := p.m[key];ok {
		*value = v
		return nil
	}
	return errors.New("not found")
}
func (p *KVStoreService)Set(kv [2]string,reply *struct{}) error {
	p.sm.Lock()
	defer p.sm.Unlock()

	key,value := kv[0],kv[1]
	if oldVal := p.m[key];oldVal != value {
		for _,fn := range p.filter  {
			fn(key)
		}
	}

	p.m[key] = value
	return nil
}
func (p *KVStoreService)Watch(timeout time.Duration,keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d",time.Now,rand.Int())
	ch := make(chan string,10) //同时监听10个

	p.sm.Lock()
	p.filter[id] = func(key string) {
		ch <- key
	}
	p.sm.Unlock()

	select {
	case <-time.After(timeout):
		return fmt.Errorf("timeout after %s",timeout.String())
	case key := <- ch:
		*keyChanged = key
		return nil
	}

	return nil
}

//rpc server
func grpc1()  {
	var (
		listener net.Listener
		done = make(chan bool,1)
	)
	fmt.Println("start to listening rpc services...")
	err := func() (e error) {
		if e = rpc.RegisterName("KVStoreService",NewKVStoreService()); e != nil {
			return e
		}
		fmt.Println("register KVStoreService success! start to listening :1234...")
		if listener,e  = net.Listen("tcp",":1234");e != nil {
			return e
		}
		fmt.Println("listening success! start to accept...")
		return nil
	}()
	if err != nil {
		log.Fatal(err.Error())
	}

	//accept 1 connect req
	go func() {
		fmt.Println("wait accept...")
		con,err1 := listener.Accept()
		if err1 != nil {
			log.Println(err1.Error())
		}
		fmt.Println("accept success!")
		rpc.ServeConn(con)
		fmt.Println("serve conn success!")
		defer con.Close()
		done <- true
	}()

	<-done
	fmt.Println("exit success.")
}
//rpc client
func grpc2()  {
	conn,err := net.Dial("tcp","127.0.0.1:1234")
	if err != nil {
		log.Fatal(err.Error())
	}
	client := rpc.NewClient(conn)
	doClientWork(client)
}
func doClientWork(client *rpc.Client)  {
	go func() {
		var keyChanged string
		err := client.Call("KVStoreService.Watch",time.Second * 30,&keyChanged)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("watch success,timeout:30s,watched:",keyChanged)
	}()

	err := client.Call("KVStoreService.Set",[2]string{"abc","abc-value"},new(struct{}))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Call KVStoreService.Set success!")
	time.Sleep(time.Second*30)
}

//type HelloService struct {}
//func (p *HelloService)Hello(request *String,reply *String) error {
//	reply.Value = "hello:" + request.GetValue()
//	return nil
//}

/**
4.3.3 反向rpc 服务端
使用场景: 在公司内网提供一个rcp服务,但是外网无法连接到内网的服务器,这个时候从内网链接到外网的tcp服务器,然后
基于tcp连接向外网提供服务
*/
func grpc3()  {
	rpc.Register(new(HelloService1))

	for {
		//连接外网的rpc服务端
		conn,_ := net.Dial("tcp",":1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}

		rpc.ServeConn(conn)
		conn.Close()
	}
}
//反向rpc 客户端
func grpc4()  {
	//建立一个tcp服务,然后等待内网连接
	listener, err := net.Listen("tcp",":1234")
	if err != nil {
		log.Fatal("ListenTCP error:",err)
	}

	//用于等待内网连接,阻塞的作用
	clientChan := make(chan *rpc.Client)

	go func() {
		for {
			//接受内网的连接
			conn,err := listener.Accept()
			if err != nil {
				log.Fatal("Accept error")
			}

			//连接成功后放入管道
			clientChan <- rpc.NewClient(conn)
		}
	}()

	//等待内网连接后执行内网rpc服务的方法
	doClientWork1(clientChan)
}
func doClientWork1(clientChan <-chan *rpc.Client)  {
	client := <- clientChan
	defer client.Close()

	var reply string
	err := client.Call("HelloService.Hello","hello",&reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

/*
4.3.4 上下文信息 + 简单登录状态的验证
为外网客户端提供每个单独的连接,即平行的多连接,上下文特性
*/
type HelloService1 struct {
	conn net.Conn
	isLogin bool
}
func (p *HelloService1)Hello1(request string,reply *string) error {
	if p.isLogin {
		return fmt.Errorf("pls login first!")
	}
	*reply = "hello:" + request + ", from " + p.conn.RemoteAddr().String()
	return nil
}
func (p *HelloService1)Login(request string,reply *string) error {
	//账号密码验证
	if request != "user:password" {
		return fmt.Errorf("auth failed")
	}
	log.Println("login.ok")
	p.isLogin = true
	return nil
}
//反向rpc客户端
func grpc5()  {
	listener,err := net.Listen("tcp",":1234")
	if err != nil {
		log.Fatal("Listen tcp error :",err)
	}

	for  {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error:",err)
		}

		go func() {
			defer conn.Close()

			p := rpc.NewServer()
			p.Register(&HelloService1{conn:conn})
			p.ServeConn(conn)
		}()
	}
}

// https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-04-grpc.html
// 4.4.2 grpc 入门

// https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-05-grpc-hack.html
// 4.5.1 grpc 进阶
// 1.go get -u google.golang.org/grpc && go get -u github.com/golang/protobuf/protoc-gen-go
// 2.write hello.proto
// 3.protoc --go_out=plugins=grpc:. hello.proto
// 4.go run grpc.go grpc6 (grpc server)
// 5.go run grpc.go grpc7 (grpc client)
type HelloServiceImpl struct {

}
func (p *HelloServiceImpl)Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{Value: "hello: " + args.GetValue()}
	return reply,nil
}

//grpc server
func grpc6()  {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer,new(HelloServiceImpl))

	lis, err := net.Listen("tcp",":1234")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("grpc server in :1234")
	grpcServer.Serve(lis)
}
//grpc client
func grpc7()  {
	conn, err := grpc.Dial("localhost:1234",grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value:"hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}

// https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-06-grpc-ext.html
// 4.6 grpc 和 protobuf 拓展
// 4.6.1 protobuf 验证器
// 4.6.2 grpc -> rest api