package main

import (
	"bytes"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math"
	"os"
	"reflect"
	"strings"
)

/*
7.array slice []bytes string
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
		"slice4" : slice4,
		"slice5" : slice5,
		"app1" : app1,
		"slice6" : slice6,
		"slice7" : slice7,
		"slice8" : slice8,
		"slice9" : slice9,
		"slice10" : slice10,
		"slice11" : slice11,
		"slice12" : slice12,
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
	new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址(数组、结构体和所有的值类型)
	make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel
	*/
	s5 := make([]int,3,5)
	s6 := new([5]int)[0:3]
	fmt.Println(s5,s6)

	//使用append后,就是复制赋值了
	s7   := [3]int{1,2,3} //[1,2,3]
	s8   := s7[:]         //[1,2,3]
	s8[0] = 3             //[3,2,3]
	s8    = append(s8,4) //[3,2,3,4]
	s8[0] = 2 //[2,2,3,4]
	fmt.Println(s7,s8)//[3,2,3] [2,2,3,4]

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
	buf4.Write([]byte("abc"))
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

//切片重组(reslice) 左闭右开[)[indexStart:indexEnd] len = indexEnd - indexStart
 func slice3()  {
	a  := [5]int{9,8,7,6}

	a1 := a[2:3] //[7]
	fmt.Println(a1)

	a1 = a[2:4] //[7,6]
	fmt.Println(a1)

	a1 = append(a1,10)
	fmt.Println(a1,a)//[7,6,10] [9,8,7,6,10]
}

/*
切片复制
copy(dst,src) 将 src切片 复制到 dst切片 上(对应位置),
返回复制的数量
*/
func slice4()  {
	sl_from := []int{1,2,3,4}
	sl_to   := make([]int,10)
	n := copy(sl_to,sl_from)
	fmt.Println(sl_to,n,len(sl_to))
}

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/07.5.md
切片append扩容
*/
func AppendByte(slice []byte,data ...byte) ([]byte) {
	sLen := len(slice)
	totalLen := sLen + len(data)
	//扩容,新建一个内存
	if totalLen > cap(slice) {
		//扩容的长度为原本长度+要推进去的切片的长度
		newSlice := make([]byte,(totalLen+1)*2)
		copy(newSlice,slice)
		slice = newSlice
	}
	slice = slice[0:totalLen]
	copy(slice[sLen:totalLen],data)
	return slice
}

//过滤出偶数的切片
func slice5()  {
	s  := []int{1,2,3,4}
	fn := func(i int) (ou bool) {
		if i%2 == 0 {
			ou = true
		}
		return
	}
	a := Filter(s,fn)
	fmt.Println(a)
}
func Filter(s []int,fn func(int) bool) (s1 []int) {
	for _,si := range s  {
		if fn(si) {
			s1 = append(s1,si)
		}
	}
	return
}

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/07.6.md
//append函数用法
func app1() {
	//1.切片追加
	s1 := make([]int, 0)
	s2 := []int{1, 2, 3}
	s1 = append(s1, s2...) //<=> append(s1,s2[0],s2[1],s2[2])
	fmt.Println(s1)

	//2.复制
	s3 := make([]int, len(s1)*2)
	copy(s3[len(s1)-1:], s1)
	fmt.Println(s3)

	//3.删除位于索引i的元素
	i := 2
	s3 = append(s3[:i], s3[i+1:]...)
	fmt.Println(s3)

	//4.i~j 切除
	j := 3
	s3 = append(s3[:i], s3[j:]...)
	fmt.Println(s3)

	//5.切片拓展
	s3 = append(s3, make([]int, 3)...)
	fmt.Println(s3)

	//6.索引i插入元素
	s3 = append(s3[:i-1], 3)
	fmt.Println(s3)

	//7.索引i追加长度为j的新切片
	s3 = append(s3[:i], []int{1, 3, 4}...)
	fmt.Println(s3)

	//切片和垃圾回收
	//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/07.6.md
}

func slice6()  {
	str := "abcdefgf"
	//efgf + abcd
	a   := str[len(str)/2:] + str[:len(str)/2]
	fmt.Println(str[len(str)/2:],str[:len(str)/2],a)
	fmt.Println(str[len(str)/2:len(str)])
}

//翻转slice []byte string
func slice7()  {
	str    := "Google"
	strSli := []byte(str)
	L      := len(strSli)
	half   := 0
	//长度为奇数需要处理一下,中间的那个数不需要翻转
	if L%2 != 0 {
		half = int(math.Ceil(float64(L/2)))
	} else {
		half = L/2
	}
	j      := L-1
	var a byte
	for i := range strSli[:half]  {
		a = strSli[i]
		strSli[i] = strSli[j]
		strSli[j] = a
		j--
	}
	str = string(strSli)
	fmt.Println(str)
}

func slice8()  {
	f := func(ar... string) string {
		if len(ar) > 0 {
			return ar[0]
		}
		return ""
	}

	b := f()
	fmt.Println(b)
	c := f("1")
	fmt.Println(c)
}

func slice9()  {
	//a := "13129551272"
	//fmt.Println(a[len(a)-6:])


	v5 := uuid.Must(uuid.NewV4())

	fmt.Println(strings.Replace(v5.String(),"-","",-1)[:uuid.Size])
}
func slice10()  {
	b1 := 2
	b  := []int{1,2,3}
	a,ok := in_slice(b1,b)
	fmt.Println(a,ok)
}
func slice11()  {
	s := []int{1, 2, 3, 4, 5, 6}
	s = s[1:4]
	fmt.Println(s)//2,3,4

	s = s[:5]
	fmt.Println(s)//2,3,4,5,6

	s = s[1:]//3,4,5,6
	fmt.Println(s)

	//[low:high:max] max: 最大容量则只能是索引 max-1 处的元素
	fmt.Println(s[0:1:3],cap(s[0:1:2]))
}

func in_slice(sliceElement interface{}, Slices interface{}) (int,bool) {
	index  := -1
	exists := false
	value  := reflect.ValueOf(Slices)
	if value.Kind() != reflect.Slice {
		return index, false
	}

	for i := 0; i < value.Len(); i++ {
		if reflect.DeepEqual(sliceElement,value.Index(i).Interface()) {
			index  = i
			exists = true
			break
		}
	}

	return index,exists
}

func slice12()  {
	a  := [2]int{1,2}
	a1 := a[:]
	for _, v := range a1 {
		a1 = append(a1,2)
		fmt.Println(v)
	}
	fmt.Println(a1)

	a2 := a[:]
	a2L:= len(a2)
	i  := 0
	for ;i < a2L;i++ {
		if i <= 1 {
			a2  = append(a2,2)
			a2L = len(a2)
		}
		fmt.Println(a2[i])
	}
	fmt.Println(a2)
}