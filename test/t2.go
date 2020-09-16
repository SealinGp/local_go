package main

import (
	"fmt"
	"strconv"
	"sync"
)
func main() {
	wg_test()
}
//1.waitGroup?
//Add为什么要放在go关键字外面(放在wait之前)? 因为有可能wait执行结束了,add都还没添加进去
//i的值为什么会超过9 ? 因为在i++完了之后,还没来得及做i<10的比较时,协程就已经打印了
func wg_test()  {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}


//2.http请求按顺序处理?

//3.限流的 sync.Mutex(Lock的时候挂起) 跟 atomic.CompareAndSwap 区别?

//4.死锁问题? sync.waitGroup or channel

//goroutine 交替打印
func switchPrint()  {
	c1 := make(chan string)
	c2 := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 0; i < 3; i++ {
			c1 <- strconv.Itoa(i)

			s := <-c2
			fmt.Print(s)
		}
	}()
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 'a'; i < 'd'; i++ {
			v := <-c1

			fmt.Print(v)

			c2 <- string(i)
		}
	}()
	wg.Wait()
	fmt.Println("")
}