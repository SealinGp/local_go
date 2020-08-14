package main
import (
	"fmt"
	"os"
	"regexp"
)

/**
https://github.com/cdoco/learn-regex-zh

特殊字符
$ 匹配以前面的子表达式结尾
()标记一个子表达式的开始和结束为止
? 匹配前面的子表达式 0或1 个字符 (通配符号)
* 匹配前面的子表达式 >= 0 个字符 (通配符号)
+ 匹配前面的子表达式 >= 1 个字符  (通配符号)
. 匹配除了换行符\n之外的所有单字符
| 匹配或
{} 限定符表达式

限定符 指定表达式出现多少次满足匹配
{n}   匹配前面的子表达式 n 次
{n,}  匹配前面的子表达式 >= n 次
{n,m} 匹配前面的子表达式 n~m  次


{1,2} 匹配从1到2个字符
(xyz) 字符组,按照确切的顺序匹配字符xyz
^() 以括号里面的开始
()$ 以括号里面的结束
[^] 否定字符类。匹配除了方括号中的字符之外的其他字符
[]  表示字符或的范围,无序,如[a-zA-Z0-9_-] 表示包含a-z或A-Z或0-9或_或-的字符


非打印字符
\n 匹配一个换行符号
\r 匹配一个回车符号(windows: 每行结尾是\r\n mac: 每行结尾是\r linux:每行结尾是\n)
\f 匹配一个分页符号
\s 匹配任何空白字符,包括空格,制表,换页符,等价于[\f\n\r\t\v]
\S 匹配任何非空白字符,等价于[^\f\n\r\t\v]
\t 匹配一个制表符,等价于\x09 和\cl
\v 匹配一个垂直制表符,等价于\x0b 和\cK
\w 匹配一个在[a-zA-Z0-9_](所有字母和数字的字符)
\W 匹配一个非字母和数字的字符[^\w]
\d 匹配一个数字[0-9]
\D 匹配一个非数字[^\d]




不支持的匹配规则
\cx 匹配由x指明的控制字符,x的值必须为A-Z或a-z之一,否则将c视为原义字符
\cM 匹配由一个Control-M或回车符号

标记
/i 匹配不区分大小写
/g 全局搜索,搜索整个输入字符串中的所有匹配
/m 多行匹配.会匹配输入字符串每一行
*/
func main() {
	fun := map[string]func(){
		"regex1":regex1,
		"regex2":regex2,
		"regex3":regex3,
	}
	fun[os.Args[1]]()
}

//1.基本匹配
func regex1()  {
	re("the cat sat on the mat",`cat`)
}

//2.元字符
func regex2()  {
	//.
	re("The car parked in the garage.",`.ar`)
	//字符集
	re("The car parked in the garage.",`[tT]he`)
	//否定字符集
	re("The car parked in the garage.",`[^c]ar`)
	//*
	re("The car parked in the garage.",`[a-z]*`)
	//\s
	re("The fat cat sat on the cat",`\s*cat\s*`)
	//+
	re("The fat cat sat on the mat.",`c.+t`)
	//?
	re("The fat cat sat on the mat.",`T?he`)
	//{}
	re("The number was 9.9997 but we rounded it off to 10.01",`[0-9]{2,}`)
	//|
	re("The car parked in the garage.",`(T|t)he|car`)
	//转义字符
	re("The fat cat sat on the mat.",`(f|c|m)at\.?`)
	//$
	re("The fat cat sat on the mat.",`(at\.?)$`)
	//\w
	re("The fat cat sat on the mat.",`\w+`)
}

//常用正则表达式
func regex3()  {
	//电话号码
	re("0755-123",`^+?[\d\s-]{3,}$`)
}

func re(str,pattern string)  {
	reg  := regexp.MustCompile(pattern)
	str1 := reg.FindAllString(str,-1)
	fmt.Println(str1,len(str1))
}
