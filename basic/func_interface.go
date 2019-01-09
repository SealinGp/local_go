package main;
import(
	"fmt"
	"strconv"
);

func main() {
	var phone Phone;
	nokia := Nokia{number:13129551271};

	//实现Phone接口
	phone = new(arg);
	phone.call(nokia);

	//同时也将方法call存入arg struct中,可直接调用
	a := arg{};
	a.call(nokia);
}

//struct
type Nokia struct {
	number uint64
}

//save function call() in arg struct
type arg struct {

}
//interface
type Phone interface {
	call(no Nokia)
}


//execute the interface
func (a arg) call(no Nokia) {
	// fmt.Printf("calling %d...\n",no.number);

	str := "calling " + strconv.FormatUint(no.number,10) + "...";
	fmt.Println(str);

	/*str := "calling %d...\n";
	str = fmt.Sprintf(str,no.number);
	fmt.Println(str);*/
}

