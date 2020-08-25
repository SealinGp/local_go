package main

import (
	"context"
	"fmt"
	"time"
)

/*
https://draveness.me/golang/concurrency/golang-context.html
context 上下文使用(控制 多个goroutine执行超时/错误时 提前结束 当前goroutine 以及触发后面的goroutine都取消)
*/

func main() {
	ctx,cf := context.WithTimeout(context.Background(),time.Second*1)
	ctx,cf1 := context.WithCancel(ctx)
	defer cf()
	defer cf1()


	go doSth(ctx,time.Millisecond*1500)
	select {
	case <-ctx.Done():
		fmt.Println("main",ctx.Err())
	}
}

func doSth(ctx context.Context,dur time.Duration)  {
	select {
	case <-ctx.Done():
		fmt.Println("handle",ctx.Err())
	case <-time.After(dur):
		fmt.Println("process request with",dur)
	}
}