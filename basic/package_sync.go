//doc:https://golang.org/pkg/sync/
package main

import (
	"fmt"
	"sync"
	"time"
)

var syncFuncs = map[string]func(){
	"mutex_sync": mutex_sync,
	"sync1":      sync1,
	"sync2":      sync2,
	"sync3":      sync3,
}

/*互斥锁,注:首次使用后不可复制,func(表示属于哪个结构体) name(输入参数)(返回参数){}
共享资源(变量)
 Type Mutex struct {}
    func (m *Mutex)Lock()	//锁住m,若m阻塞到m解锁
    func (m *Mutex)Unlock() //解锁m,若m未加锁会导致错误
 Type WaitGroup struct {}
 	添加int类型的delta到WaitGroup计数器中,
	若计数器=0,则释放在等待时阻止的所有goroutine,
	若计数器<0,则报错
    func (wg *WaitGroup) Add(delta int)

    将WaitGroup计数器-1
	func (wg *WaitGroup) Done()

	等待直到WaitGroup=0
	func (wg *WaitGroup) Wait()
 def:
    var m sync.Mutex;
	m := sync.Mutex{};
匿名函数:
 可以被赋于一个变量
 匿名函数的闭包最后一个括号表示对匿名函数的调用
 a := func(){}();
*/
func mutex_sync() {
	var mutex sync.Mutex    //控制协程内的工作范围
	var wait sync.WaitGroup //控制程序等待运行结束

	fmt.Println("locked")
	mutex.Lock()

	for i := 1; i <= 3; i++ {
		wait.Add(1)

		go func(i int) {

			//锁开始
			mutex.Lock()
			fmt.Println("Lock:", i)

			time.Sleep(time.Second)

			//锁结束
			fmt.Println("unlock:", i)
			mutex.Unlock()

			defer wait.Done()
		}(i)
	}

	time.Sleep(time.Second)

	fmt.Println("Unlocked")
	mutex.Unlock()

	wait.Wait() //等待所有协程运行完毕
}

/*
refUrl : https://tour.go-zh.org/concurrency/9
共享读,独占写, 写写互斥, 写读互斥, 读读不互斥
*/
func sync1() {
	c := SafeCount{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.set("key1")
	}

	time.Sleep(time.Second)
	fmt.Println(c.get("key1"))
	fmt.Println("after defer")
}

type SafeCount struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCount) set(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}
func (c *SafeCount) get(key string) int {
	c.mux.Lock()
	defer func() {
		fmt.Println("???")
		c.mux.Unlock()
	}()
	return c.v[key]
}

func sync2() {
	c := SafeCount1{v: make(map[string][]int)}
	for i := 0; i < 1000; i++ {
		go c.set("key1", i)
	}

	time.Sleep(time.Second) //预估gorountine 执行时间,等待其执行结束
	fmt.Println(c.get("key1"))
	fmt.Println("after func get defer")
}

type SafeCount1 struct {
	v   map[string][]int
	mux sync.Mutex
}

func (c *SafeCount1) set(key string, index int) {
	c.mux.Lock()
	c.v[key] = append(c.v[key], index)
	c.mux.Unlock()
}
func (c *SafeCount1) get(key string) []int {
	c.mux.Lock()
	defer func() {
		fmt.Println("after get")
		c.mux.Unlock()
	}()
	return c.v[key]
}
func (c *SafeCount1) getLen(key string) int {
	c.mux.Lock()
	defer func() {
		fmt.Println("???")
		c.mux.Unlock()
	}()
	return len(c.v[key])
}

func sync3() {
}
