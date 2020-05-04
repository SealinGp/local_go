package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func main() {
	go func() {
		if err := http.ListenAndServe(":9876",nil);err != nil {
			println(err)
			os.Exit(1)
		}
	}()
	outCh := make(chan int)

	// 死代码，永不读取
	go func() {
		if false {
			<-outCh
		}
		select {}
	}()

	//开启100个goroutine/秒,并且goroutine内阻塞不退出
	tick := time.Tick(time.Second/100)
	i    := 0
	for range tick {
		i++
		fmt.Println(i)
		alloc1(outCh)
	}
}

//goroutine内阻塞不退出
func alloc1(outCh chan<- int)  {
	go func() {
		defer fmt.Println("alloc-fm exit")

		//分配内存,假装用一下
		buf := make([]byte,(1<<20)*10)
		_   = len(buf)
		fmt.Println("alloc done")
		outCh <- 0
	}()
}