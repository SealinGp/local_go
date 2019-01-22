package main;
import (
	"os"
	"fmt"
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
	};	
	funs[n]();
}

func channel1() {
	fmt.Println("123");
}