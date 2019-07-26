package main;

import (
	"os"
	"fmt"
);

func init() {
  fmt.Println("Content-Type:text/plain;charset=utf-8\n\n");
}
func main() {
	args := os.Args;
    if len(args) <= 1 {
    	fmt.Println("lack param ?func=xxx");
    	return;
    }

	execute(args[1]);
}
func execute(n string) {
	funs := map[string]func() {
		"keyword1"     : keyword1,
	};	
	funs[n]();		
}

func keyword1() {	
    var kWord = [...]string{"break","default","func","interface","select","case","defer","go","map","struct","chan","else","goto","package","switch","const","fallthrough","if","range","type","continue","for","import","return","var"}
	kLen := len(kWord)
    for i := 0;i < kLen; i++ {
        fmt.Println(kWord[i])
    }
}