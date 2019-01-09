package main;
import (
	"fmt"
);

func main() {
	args := os.Args;
	execute(args[1]);	
}
func execute(n string) {
	funs := map[string]func() {
		"channel1" : channel1,
	};	
	funs[n]();
}

func channel1() {
	
}