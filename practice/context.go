package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
https://draveness.me/golang/concurrency/golang-context.html
context 上下文使用(控制 多个goroutine执行超时/错误时 提前结束 当前goroutine 以及触发后面的goroutine都取消)
*/

func main() {
	Cond()
}
func Cond() {
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go listen(c, i)
	}
	time.Sleep(time.Second * 1)

	go broadcast(c)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
func listen(c *sync.Cond, i int) {
	c.L.Lock()
	c.Wait()
	log.Println("listen", i)
	c.L.Unlock()
}
func broadcast(c *sync.Cond) {
	c.L.Lock()
	c.Broadcast()
	c.L.Unlock()
}

func Ctx() {
	ctx, cf := context.WithTimeout(context.Background(), time.Second*1)
	ctx = context.WithValue(ctx, "k", "v")
	defer cf()

	go doSth(ctx, time.Second*4)
	select {
	case <-ctx.Done():
		log.Println("main", ctx.Err())
	}
}

func doSth(ctx context.Context, dur time.Duration) {
	select {
	case <-ctx.Done():
		log.Println("handle", ctx.Err())
	case <-time.After(dur):
		log.Println("process request with", dur)
	}
}
