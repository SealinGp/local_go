package main

import (
	"bytes"
	"fmt"
	"os"
)

/*
array
*/
func init() {
	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
}
func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("lack param ?func=xxx")
		return
	}

	execute(args[1])
}
func execute(n string) {
	funs := map[string]func(){
		"arr1" : arr1,
		"arr2" : arr2,
		"arr3" : arr3,
		"slice1" : slice1,
		"by1" : by1,
		"slice2" : slice2,
		"slice3" : slice3,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

//数组声明方式的区别
func arr1()  {
	var a1 [5]int     //[5]int
	a2 := new([5]int) //*[5]int
	a1  = *a2
	a1[2] = 100

	fmt.Println(a1)
	fmt.Println(a2)
}

//多维数组
func arr2()  {
	var a [2][2][2]float64

	f  := [2]float64{2.1,2.2}
	f2 := [2][2]float64{f,f}
	a[0] = f2
	a[1] = f2
	fmt.Println(a[0])
}

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/07.1.md
把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：
传递数组的指针
使用数组的切片
 */

//传递数组的指针
func arr3()  {
	a1 := [3]float64{1.0,2.0,3.0}
	s  := arr3_1(&a1)
	fmt.Println(s)
}
func arr3_1(a *[3]float64) (sum float64)  {
	for _, v := range a  {
		sum += v
	}
	return
}

//传递数组的切片

//切片(引用赋值)
func slice1()  {
	s1 := []int{1,2,3}
	s2 := s1
	s2[0] = 2
	fmt.Println(s1,s2)

	s3 := [3]int{1,2,3}
	s4 := s3[:]
	s4[0] = 2
	fmt.Println(s3,s4)

	/*
	new 和 make的区别
	new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址
	make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel
	*/
	s5 := make([]int,3,5)
	s6 := new([5]int)[0:3]
	fmt.Println(s5,s6)

	//使用append后,就是复制赋值了
	s7   := [3]int{1,2,3}
	s8   := s7[:]
	s8[0] = 3
	s8    = append(s8,4)
	s8[0] = 2
	fmt.Println(s7,s8)

	//
	s9  := []byte{'p','o','e','m'}
	s10 := s9[2:] //e,m
	s10[1] = 't'  //e,t
	fmt.Println(string(s9),string(s10))
	var p rune = 'p'
	fmt.Println(p,string(p))
}

func by1()  {
	//缓冲byte类型的缓冲器
	buf1 := bytes.NewBufferString("hello")
	buf2 := bytes.NewBuffer([]byte("hello"))
	buf3 := bytes.NewBuffer([]byte{'h','e','l','l','o'})
	fmt.Println(
		buf1.String(),
		buf2.String(),
		buf3.String(),
	)

	buf4 := bytes.NewBufferString("")
	buf5 := bytes.NewBuffer([]byte{})
	buf6 := bytes.Buffer{}
	buf6.WriteByte('a')
	fmt.Println(
		buf4.String(),
		buf5.String(),
		buf6.String(),
	)

	//write,writeString,WriteByte
	s := []byte{' ','w','o','r','l','d'}
	buf4.WriteString("hello")
	buf4.Write(s)
	buf4.WriteByte('?')
	buf4.WriteRune('a')
	fmt.Println(buf4.String())
}

func slice2()  {
	sum := func(f []float32) (s float32) {
		for _,v := range f  {
			s += v
		}
		return
	}
	f := []float32{2.1,2.2}
	fmt.Println(sum(f))
}

//切片重组(reslice) [indexStart:indexEnd] len = indexEnd - indexStart
func slice3()  {
	a := [5]int{9,8,7,6}
	a1 := a[2:3]
	fmt.Println(a1)
	a1 = a[2:4]
	fmt.Println(a1)
}