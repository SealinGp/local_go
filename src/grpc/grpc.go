package main

import (
	"errors"
	"fmt"
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
		//for  {
			con,err1 := listener.Accept()
			if err1 != nil {
				log.Println(err1.Error())
				//break
			}
			fmt.Println("accept success!")
			rpc.ServeConn(con)
			fmt.Println("serve conn success!")
			defer con.Close()
		//}
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

func grpc3()  {

}
