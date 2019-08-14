package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.4.md
解析命令行,获取参数,生成使用文档
*/
func init() {
	//fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
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
		"flag1" : flag1,
		"flag2" : flag2,
		"flag3" : flag3,
		"flag4" : flag4,
		"flag5" : flag5,
		"flag6" : flag6,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

//定义了一个默认值是 false的 -n选项, : -n=false
var (
	NewLine = flag.Bool("n",false,"print newline")
	Help    = flag.Bool("h",false,"print this help")
	Newline = "\n"
)
func flag1()  {
	//扫描参数列表,并设置flag
	flag.Parse()

	output := ""
	//flag.NArg 返回参数的数量
	for i := 0; i < flag.NArg() ; i++ {
		if flag.Arg(i) == "-h" {
			flag.PrintDefaults()
			return
		}
		output +=flag.Arg(i) + Newline
	}
	os.Stdout.WriteString(output)
}

//用buffer读取文件
func flag2()  {
	flag.Parse()

	//参数除了func的参数(第一个)后,若无其他参数,则命令行输入=输出
	if flag.NArg() == 0 || flag.NArg() == 1 {
		cat(bufio.NewReader(os.Stdin))
	}

	//命令行若有其他参数,若为文件,则读取输出
	for i := 0; i < flag.NArg() ; i++ {
		if i == 0 {
			continue
		}
		f,err := os.Open(flag.Arg(i))
		if err != nil {
			os.Stdout.WriteString(err.Error()+Newline)
			continue
		}
		cat(bufio.NewReader(f))
		f.Close()
	}

}
//读取文件,输出到命令行
func cat(r *bufio.Reader)  {
	for {
		buf,err := r.ReadString('\n')

		//读取报错
		if err != nil && err != io.EOF {
			os.Stdout.WriteString("error:" + err.Error() + Newline)
			break
		}

		os.Stdout.WriteString(buf + Newline)
		//文件末尾
		if err == io.EOF {
			break
		}
	}
}

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.6.md
//用切片读写文件
func flag3()  {
	file,err := os.Open("a.txt")
	if err != nil {
		//os.Stdout.WriteString(err.Error())
		return
	}
	cat2(file)
}
func cat2(file *os.File)  {
	const NBUF  = 512
	//数组
	var buf [NBUF]byte
	for {
		switch nr,err := file.Read(buf[:]);true {
		case nr < 0:  //err != os.EOF && err != nil
			os.Stderr.WriteString("cat: error reading: " + err.Error())
			os.Exit(1)
		case nr == 0: //err = os.EOF
			return
		case nr > 0: //err = nil
			buf[nr-1] = '\n'
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				os.Stderr.WriteString("cat: error writing: " + ew.Error())
			}
		}
	}
}
//用切片读取文件2
func flag4()  {
	flag.Parse()
	if flag.NArg() <= 1 {
		fmt.Println("no file")
		return
	}
	for i := 0; i < flag.NArg(); i++ {
		if i == 0 {
			continue
		}
		f,err := os.Open(flag.Arg(i))
		if f == nil {
			os.Stderr.WriteString(err.Error())
			os.Exit(1)
		}
		cat3(f)
		f.Close()
	}
}
func cat3(file *os.File)  {
	const NBUF  = 512
	var buf [NBUF]byte
	for {
		switch nr,err := file.Read(buf[:]);true {
		case nr < 0 :
			os.Stderr.WriteString("cat: error reading: " + err.Error())
			os.Exit(1)
		case nr == 0 :
			return
		case nr > 0 :
			buf[nr-1] = '\n'
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				os.Stderr.WriteString("cat: error writing: " + ew.Error())
			}
		}
	}
}

/*
fmt.Fprintf 写入内容到输出
数据I/O的模型,3个文件描述符
os.Stdout 标准输出流
os.Stderr 标准错误流
os.Stdin  标准输入流
*/
func flag5()  {
	fmt.Fprintf(os.Stdout,"%s\n","-- unbuffered")
	buf := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(buf,"%s\n","buffered")

	buf.Flush()
}

/*
练习
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.8.md
*/
func flag6()  {
	inputFile, _  := os.OpenFile("a.txt",os.O_WRONLY,0666)
	//先写入需要读取的数据
	str := "ab123asfasfas\ngh456fasgsdsdgash\ngh789jklfjklahjkjk"
	bw  := bufio.NewWriter(inputFile)
	bw.WriteString(str)
	bw.Flush()
	inputFile.Close()

	inputFile,_   = os.Open("a.txt")
	outputFile,_ := os.OpenFile("a1.txt",os.O_WRONLY|os.O_CREATE,0666)
	defer inputFile.Close()
	defer outputFile.Close()

	inputReader  := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString , _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("end")
			break
		} else if readerError != nil {
			fmt.Println(readerError)
			return
		}

		//第3~第5 [3-1,5-1+1) 左闭右开
		outputString := string(inputString[2:5]) + "\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	//缓冲流写入需要刷新
	e := outputWriter.Flush()
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println("Conversion done")
}