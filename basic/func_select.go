package main

import (
	"fmt"
)

var funcSelectFuncs = map[string]func(){
	"code1": code1,
	"code2": code2,
	"code3": code3,
	"code4": code4,
	"code5": code5,
}

func code1() {
	var c1 chan uint8
	var c2 chan uint8
	var c3 chan uint8

	select {
	case v1 := <-c1:
		fmt.Printf("received %v from c1\n", v1)
	case v2 := <-c2:
		fmt.Printf("received %v from c2\n", v2)
	case c3 <- 23:
		fmt.Printf("sent %v to c3\n", 23)
	default:
		fmt.Printf("no one was ready to communicate\n") //
	}
	/*
			类似于switch语句
		除 default 外,若只有一个 case 语句评估通过，那么就执行这个case里的语句；
		除 default 外,若有多个 case 语句评估通过，那么通过伪随机的方式随机选一个；
		如果 default 外的 case 语句都没有通过评估，那么执行 default 里的语句；
		如果没有 default，那么 代码块会被阻塞，指导有一个 case 通过评估；否则一直阻塞

		ref:https://www.cnblogs.com/jianxinzhou/p/3931893.html
		linux下的select 通过检查管道是否阻塞,来进行监听
	*/
}
func code2() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received") //
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent") //
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity") //
	}
}

//带缓存的管道-------------------------------
func code3() {
	messages := make(chan string, 1)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received") //
	}

	msg := "hi world"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg) //
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message1", msg) //
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func code4() {
	messages := make(chan string, 1)
	signals := make(chan bool)

	messages <- "test"
	select {
	case msg := <-messages:
		fmt.Println("received message", msg) //
	default:
		fmt.Println("no message received")
	}

	msg := "hi world"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg) //
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg) //
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

//fibonacci 斐波那契数列 例子
func code5() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fib(c, quit)
}

func fib(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			/*
				i = 0,c = 0,x = 0,y = 1;
				fmt.Println(c);//0

				i = 1,c = 1,x = 1, y = 1;
				fmt.Println(c);//1

				i = 2,c = 1,x = 1, y = 2;
				fmt.Println(c);//1

				i = 3,c = 2,x = 2, y = 3;
				fmt.Println(c);//2
			*/
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
