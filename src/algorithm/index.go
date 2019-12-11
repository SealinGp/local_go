package main

import (
	"algorithm/simple"
	"fmt"
	"os"
	"reflect"
	"time"
)

func main() {
	startTime := time.Now()
	defer func() {
		d := time.Now().Sub(startTime)
		fmt.Println("processTime:",d)
	}()


	args     := os.Args
	if len(args) <= 1 {
		fmt.Println("lack param ?func=xxx")
		return
	}

	ref    := simple.Ref{}
	refVal := reflect.ValueOf(&ref).Elem()
	refVal.MethodByName(args[0]).Call(nil)
}