package main

import (
	"context"
	"log"
	"time"
)

var ctxFuncs = map[string]func(){
	"ctx1": ctx1,
	"ctx2": ctx2,
	"ctx3": ctx3,
}

// https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/
// context.Background 是上下文默认值,所有其他的上下文都是由它衍生出来的
// context.TODO 应该还是在不确定应该使用哪种上下文时使用 二者都是同一个类型
func ctx1() {
	//假设一次http请求设置的超时时间为1s
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//假设用来处理这次http请求的时间为500ms
	go handle(ctx, 500*time.Millisecond)

	select {
	case <-ctx.Done():
		log.Println("main", ctx.Err())
	}
}
func handle(ctx context.Context, dur time.Duration) {
	select {
	case <-ctx.Done():
		log.Println("handle", ctx.Err())
	case <-time.After(dur):
		log.Println("process request with", dur)
	}
	//process request with 500ms
}

func ctx2() {
	//假设一次http请求设置的超时时间为1s
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//假设用来处理这次http请求的时间为1500ms > 1s 因此该请求超时,在main goroutine 和 这个子goroutine中 的ctx都会超时
	go handle(ctx, 1500*time.Millisecond)

	select {
	case <-ctx.Done():
		log.Println("main", ctx.Err())
	}
}

func ctx3() {
	//父上下文,也就当前这个
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel2()

	//从父上下文中创建子上下文,并传值
	ctx3 := context.WithValue(ctx2, "a", "b")
	go handle(ctx2, 2*time.Second)
	go handle(ctx3, time.Second)
	select {
	case <-ctx3.Done():
		log.Println(ctx3.Value("a"), "ctx3")
	case <-ctx2.Done():
		log.Println("ctx2 done")
	}
}
