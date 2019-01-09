package main;
import(
	"fmt"
)

/*
Go语言追求简洁优雅,所以,Go语言不支持传统
的 try…catch…finally 这种异常,因为Go语
言的设计者们认为,将异常与控制结构混在一起
会很容易使得代码变得混乱.因为开发者很容易
滥用异常,甚至一个小小的错误都抛出一个异常.
在Go语言中,使用多值返回来返回错误.不要用
异常代替错误,更不要用来控制流程.在极个别
的情况下,也就是说.遇到真正的异常的情况下
(比如除数为0了).才使用Go中引入的Exception
处理：defer,panic,recover.

defer :
  说明:析构函数,添加函数结束时执行的语句,defer可以多次,这样形成一个defer栈,后defer的语句在函数返回时将先被调用.
  使用:defer func_name();
panic :
  说明:panic会导致程序挂掉,panic程序挂掉后,执行defer->再向上传递,
  	   在调试程序时,通过 panic 来打印堆栈,方便定位错误.
  使用:panic(msg);
recover:
  说明:

*/
func main() {
	show_error();
	
	// throw_error("错误!");

	/*Try(func() {
		panic("throw error!");
	},func (e interface{}) {
		fmt.Println(e);
	})*/
}

func throw_error(errMsg string) {
	//catch error msg
	defer func(){
		if err := recover();err != nil {
			fmt.Println(err);
		}
	}()

	//throw error msg
	panic(errMsg);
}

/*
继承 go内置 error interface中的Error方法
type error interface {
    Error() string
}
*/
type DivideError struct {
	dividee int32
	divider int32
}
func (de *DivideError) Error() string {
	str := `
	cannot proceed,the divider is zero.
	dividee: %d
	divider: %d
`;
	return fmt.Sprintf(str,de.dividee,de.divider);
}
func show_error() {
	//main函数执行完毕后被调用的函数 defer
	defer defer_();

	de := DivideError{12,123};	
	fmt.Println(de);	
	fmt.Println(&de);	
	fmt.Println(de.Error());	
}
//main函数执行完毕后被调用的函数 defer
func defer_() {
	fmt.Println("------------");
}

/*
	golang中实现try...catch...用法
*/
func Try(fun func(),handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
            handler(err)
        }
	}();
	fun();
}