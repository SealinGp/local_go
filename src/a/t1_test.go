package main_test

import (
	"testing"
)

func BenchmarkFib1(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Fib1(11)
	}
}
func BenchmarkFib2(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Fib2(11)
	}
}
func BenchmarkFib3(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Fib3(11)
	}
}


func Fib2(n int) int {
	var c []int
	fib1 := func(a int) int {
		if lenc := len(c); lenc >= 2 {
			a = c[lenc-1] + c[lenc-2]
		}
		c = append(c, a)
		return c[len(c)-1]
	}
	res := 1
	for i := 0; i <= n; i++ {
		if i <= 1 {
			res = fib1(1)
		} else {
			res = fib1(i)
		}
	}
	return res
}

func Fib3(n int) int {
	fibNums := []int{}
	fib     := func(n int) int {
		if n <= 1 {
			fibNums = append(fibNums,1)
		} else {
			len1   := len(fibNums)
			num    := fibNums[len1-1] + fibNums[len1-2]
			fibNums = append(fibNums,num)
		}
		return fibNums[len(fibNums)-1]
	}
	res := 1
	for i := 0; i <= n; i++ {
		res = fib(i)
	}
	return res
}

func Fib1(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib1(n-1) + Fib1(n-2)
}