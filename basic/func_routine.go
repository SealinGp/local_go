package main;
import(
	"fmt"
	"time"
);

/*
go 并发
进程:
 (面向进程设计的)UNIX/Linux version<Linux2.4: 进程是程序的实体
 (面向线程设计的)win/mac os/Linux version>=Linux2.6: 进程是线程的容器
 程序本身只是指令,数据,以及组织形式的描述,进程是程序的真正实例
 多个进程 可能 与同一个程序有关,每个进程可 同步(顺序) 或 异步(平行) 的方式运行
 现代计算机系统可在同一段时间内以进程的形式将多个程序加载到存储器中,并借由时间共享(时分复用),
 以在一个cpu(处理器)上表现出同时运行的感觉,
线程:
协程:
*/ 
func main() {
	go say2("world");
	say2("hello");

	// for i := 0; i < 5; i++ {
		// say("hello ");		
	// }
}
func say(msg string) {
	fmt.Println(msg);
}
func say2(s string) {
    for i := 0; i < 5; i++ {    	
        time.Sleep(100 * time.Millisecond);//模拟IO阻塞操作
        fmt.Println(s);
    }
}
