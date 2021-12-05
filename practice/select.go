package main

import (
	"fmt"

	//"runtime"

	"time"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.5.md
通道,超时,计时器
1 s  = 1000 ms(毫秒)
1 ms = 1000 us(微妙)
1 us = 1000 ns(纳秒)
像圆周率π一样,e为数学中的常数, e = 2.71828182856904523536...
*/
//func init() {
//	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
//}
var selFuncs = map[string]func(){
	"sel1": sel1,
	"sel2": sel2,
	"sel3": sel3,
	"sel4": sel4,
}

func sel1() {
	/*
		 * 以d为周期给返回的通道发送时间,d单位纳秒,作用通常用于限制请求频率
		   chan := time.Tick(d)
		   rate_per_sec := 10
		   var dur Duration = 1e9
		   chRate := time.Tick(dur)
		   for req := range requests {
		       <- chRate // rate limit our Service.Method RPC calls,每 1e9 纳秒阻塞一次,限制调用
		       go client.Call("Service.Method", req, ...)
		   }

		* 若周期请求数量暴增,可使用带缓冲的通道和计时器对象来承载


		* 定时器Timer 跟 计时器Ticker 结构体很像,但Timer 在规定时间后,只发送一次时间
		* time.After(d duration)<-chan Time 在时间d 后,当前时间被发到返回的通道中, 因此它和NewTimer(d).C 是等价的
		* 下列例子(任务超时),也同样说明了select中default的作用
	*/
	start := time.Now()

	tick := time.Tick(time.Millisecond * 200) //200ms 执行一次(1s内执行1000/200=5,5-1 = 4次,实际上只执行了4次,最后一次还没执行就结束了)
	boom := time.After(time.Second * 1)       //1s    后结束(1000ms)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
			//任务结束的信号
		case <-boom:
			fmt.Println("BOOM!")
			fmt.Println(time.Since(start))
			return
			//在每200ms周期内都执行这个,满足200ms就执行一次其他任务
		default:
			fmt.Println("default")
			time.Sleep(5e7)
		}
	}
}

//1.简单超时模式,从通道ch中接收数据,但是最多等待1s(1e9 = 1s,约等于)
func sel2() {
	start := time.Now()
	timeout := make(chan bool, 1) //注意:超时管道设置缓冲大小,避免协程死锁,确保超时的通道可以被垃圾回收
	ch := make(chan int)
	//超时设置
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	//接收数据的协程任务
	go func() {
		time.Sleep(time.Second * 2)
		ch <- 1
	}()

	//没有执行其他的default任务,直接等待至超时
	select {
	case v := <-ch:
		fmt.Println("get ch finished:", v)
		break
	case <-timeout:
		fmt.Println("time out")
		break
	}

	//有执行其他的default任务,执行其他的default任务至超时,并且超时时间不一定是规定的时间,当default跟timeout都符合时,select会随机选择一个case
	/*f : for {
		select{
		case v := <-ch:
			fmt.Println("get ch finished:",v)
			break f
		case <-timeout:
			fmt.Println("time out")
			break f
		default:
			fmt.Println("default task")
		}
	}*/
	fmt.Println(time.Since(start))
}

//2.取耗时很长的同步调用
/**
注意缓冲大小设置为 1 是必要的，可以避免协程死锁以及确保超时的通道可以被垃圾回收。
此外，需要注意在有多个 case 符合条件时， select 对 case 的选择是伪随机的，
如果上面的代码稍作修改如下，则 select 语句可能不会在定时器超时信号到来时立刻选中
time.After(timeoutNs) 对应的 case，因此协程可能不会严格按照定时器设置的时间结束
*/
func sel3() {
	ch := make(chan int, 1)

	//执行任务
	go func() {
		time.Sleep(time.Second * 2)
		ch <- 1
	}()

	select {
	case v := <-ch:
		fmt.Println("task finished", v)
	case <-time.After(time.Second * 1):
		fmt.Println("after 1s time out")
		break
	}

	/*L:
	for {
		select {
		case <-ch:
			// do something
		case <-time.After(timeoutNs):
			// call timed out
			break L
		}
	}*/
}

//第三种形式
func sel4() {
	/**
	假设程序从多个复制的数据库同时读取,只需要一个答案,需要接收首先到达的答案,Query 函数获取数据的连接切片
	并请求,并行请求每一个数据库并返回收到的第一个响应
	func Query(conns []Conn,query string) Result {
		ch := make(chan Result,1)
		for _, conn := range conns {
			go func(c Conn) {
				select {
					case ch <- c.DoQuery(query):
					default:
				}
	 		}(conn)
		}

		//结果通道必须带缓冲,保证第一个发送进来的数据有位置可以存放,确保放入的首个数据总会成功
		//所以第一个到达的值会被获取跟执行的顺序无关,正在执行的协程总是可以使用runtime.Goexit()来停止
		return <-ch
	}

	在应用中缓存数据:
	应用程序中用到了来自数据库/常见的数据存储 的数据时,经常会把数据缓存到内存中,因为从数据库中获取数据的操作代价很高
	如果数据库中的值不发生变化,就没有问题,如果值有变化,我们需要一个机制来周期性的从数据库重新读取这些值:缓存的值
	就不可用/过期了
	*/
}
