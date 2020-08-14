package main
/*
Go 汇编
https://github.com/chai2010/advanced-go-programming-book/blob/master/ch3-asm/ch3-01-basic.md

*/
import (
	"assembler/basic"
	"fmt"
)

func main() {
	println(basic.Id)
	fmt.Println("Num:",basic.Num)
	fmt.Println("BoolValue:",basic.BoolValue)
	fmt.Println("TrueValue",basic.TrueValue)
	fmt.Println("FalseValue:",basic.FalseValue)
	fmt.Println("Int32Value:",basic.Int32Value)
	fmt.Println("Uint32Value:",basic.Uint32Value)
	fmt.Println("Float32Value:",basic.Float32Value)
	fmt.Println("Float64Value:",basic.Float64Value)
	fmt.Println("ReadOnlyInt:",basic.ReadOnlyInt)
	fmt.Println("Name:",basic.Name)
}