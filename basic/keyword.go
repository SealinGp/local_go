package main;

import (
  "fmt"
);

func init() {
  fmt.Println("Content-Type:text/plain;charset=utf-8\n\n");
}
func main() {	
    var kWord = [...]string{"break","default","func","interface","select","case","defer","go","map","struct","chan","else","goto","package","switch","const","fallthrough","if","range","type","continue","for","import","return","var"}
	kLen := len(kWord)
    for i := 0;i < kLen; i++ {
        fmt.Println(kWord[i])
    }
}

