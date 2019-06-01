//doc:https://golang.org/pkg/sync/
package main;
import (
	"fmt"
	"os"
	"sync"
	"time"

	// "path/filepath"
)

func main() {
	gocache := os.Getenv("XDG_CACHE_HOME");	
	fmt.Println(gocache);
	return;
	args := os.Args;
	execute(args[1]);	
}
func execute(n string) {
	funs := map[string]func() {
		"mutex_sync" : mutex_sync,
	};	
	funs[n]();
}

/*互斥锁,注:首次使用后不可复制,func(表示属于哪个结构体) name(输入参数)(返回参数){}
 Type Mutex struct {}
    func (m *Mutex)Lock()	//锁住m,若m阻塞到m解锁
    func (m *Mutex)Unlock()//解锁m,若m未加锁会导致错误
 Type WaitGrout struct {}
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
	var mutex sync.Mutex;   //控制协程内的工作范围
	var wait sync.WaitGroup;//控制程序等待运行结束

	fmt.Println("locked");
	mutex.Lock();

	for i := 1; i <= 3; i++ {
		wait.Add(1);

		fmt.Println("outside not lock:",i);
		go func (i int) {
			fmt.Println("inside not lock:",i);

			//锁开始
			mutex.Lock();
			fmt.Println("Lock:",i);

			time.Sleep(time.Second);

			//锁结束
			fmt.Println("unlock:",i);
			mutex.Unlock();

			defer wait.Done();
		}(i);
	}

	time.Sleep(time.Second);
	fmt.Println("Unlocked");
	mutex.Unlock();				//等待所有协程运行完毕

	wait.Wait();
}