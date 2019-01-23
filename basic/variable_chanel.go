package main;
import (
	"os"
	"fmt"
	"time"
);


func main() {
	args := os.Args;
	if len(args) <= 1 {
		fmt.Println("函数名未指定");
		return ;
	}

	execute(args[1]);
}
func execute(n string) {
	funs := map[string]func() {
		"channel1" : channel1,
		"channel2" : channel2,
		"channel3" : channel3,
		"channel4" : channel4,
		"channel5" : channel5,
		"channel6" : channel6,
		"channel7" : channel7,
	};	
	funs[n]();
}

//refurl: https://colobu.com/2016/04/14/Golang-Channels/

/*
 作用:协程之间通信的方式
 定义方式
	var ch1 chan int;    				//双向管道
	var ch1 chan<- int;  				//单向写
	var ch1 <-chan int;  				//单向读

	ch2 := make(chan int,capacity int); //capacity 容量/缓存
	容量设置的时候,当容量满的时候才会发生blocking(阻塞)
*/

//buffered channels 缓存管道,可以避免阻塞
func channel1() {
	ch := make(chan int,100);
    v  := 1;
    ch <- v;
    v1:= <-ch;

    fmt.Println(v1);
}

func channel2() {	
	c := make(chan int);

	//函数执行完毕后关闭管道
	defer close(c);
    
    //创建协程并调用
    go func() { c <- 3 + 4 }();

    i := <-c;
    fmt.Println(i);
    //报panic错误,因为管道在协程还没写入7的时候程序就结束了,管道关闭,导致panic错误
}

//blocking 阻塞
func channel3() {
	s := []int{7,2,8,-9,4,0};
	c := make(chan int);

	//入栈7+2+8=17
	go channle3_sum(s[0:len(s)/2],c);
	//入栈-9+4+0=-5
	go channle3_sum(s[len(s)/2:len(s)],c);

	//出栈 -5,17 等待管道写入数据后(等待协程执行结束),读取出来
	x, y := <-c, <-c;
	fmt.Println(x, y, x+y);
}

func channle3_sum(a []int,c chan int) {
	sum := 0;
	for _,v := range a {
		sum += v;
	}
	c <- sum;
}

//range 处理channel
func channel4() {
	go func() {
		time.Sleep(1*time.Hour);
	}();

	c := make(chan int);

	go func() {
		for i := 0; i < 10; i++ {
			c <- i;
		}

		//若在管道写入完毕后不关闭管道,则程序会一直阻塞在for..range
		close(c);
	}();

	for i := range c {
		fmt.Println(i);
	}

	fmt.Println("Finished");
}

//select 处理channel
func channel5() {
	c := make(chan int);
	quit := make(chan int);

	go func() {
		for i := 0; i < 10; i++ {
			//c读
			fmt.Println(<-c);
		}
		//quit写
		quit <- 0;
	}();

	channel5_fibonacci(c,quit);
}

func channel5_fibonacci(c,quit chan int) {
	x, y := 0, 1;

	//死循环等待协程写入quit数据后读取跳出循环
	for {
		select {
			//i=0,x=0,y=1    0
			//i=1,x=1,y=1    1
			//i=2,x=1,y=2    1
			//i=3,x=2,y=3    2
			//i=4,x=3,y=5    3
		case c <- x:
			x, y = y, x+y;
		case <-quit:
			fmt.Println("quit");
			return;
		}
	}
}

/*
timeout
time.After(t int) 在时间t后返回一个单向可读的channel
*/
func channel6() {
	c1 := make(chan string,1);
	go func() {
		time.Sleep(time.Second * 2);
		c1 <- "result 1";
	}();

	select {
	case res := <-c1:
		fmt.Println(res);
	case t1 := <-time.After(time.Second * 1):
		fmt.Println("timeout 1",t1);
	}
}

/*
time.NewTimer和ticker
time.NewTimer(t int):定时器,在t时间后返回一个单向读时间channel
*/
func channel7() {
	//2s后返回单向可读的时间channel
	timer1 := time.NewTimer(time.Second * 2);

	fmt.Println("now ",time.Now().Format("2006-01-02 15:04:05"));

	//阻塞2s
	time2 := <-timer1.C;
	
	fmt.Println("2 second later ",time2.Format("2006-01-02 15:04:05"));	
}