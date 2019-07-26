package main;

import(
	"fmt"
	"os"	

	"archive/tar"
	"bytes"	
	"io"
	"log"	
);

/*
tar打包文件
*/
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
		"tar1" : tar1,
		"Buffer" : Buffer,
		"test" : test,
	};	
	funs[n]();		
}

type file struct{
	Name string
	Body string
};

/*
ref:https://golang.org/pkg/bytes/
bytes.Buffer类型是一个可变大小的字节缓冲区,有Read,Write方法
*/
func Buffer() {
	var b bytes.Buffer;
	b.Write([]byte("Hello "));
	fmt.Fprintf(&b,"word!");
	b.WriteTo(os.Stdout);	
}

func test() {
	a := "abc";		
	fmt.Println([]byte(a));
}

/*
例子:https://golang.org/pkg/archive/tar/
*/
func tar1()  {
	/*
	https://golang.org/pkg/bytes/
	bytes.Buffer:
	type Buffer 
	    func (b *Buffer) Write(p []bype) (n int, err error);
	*/ 
	var buf bytes.Buffer;

	//tar.NewWriter(w io.Writer) *Writer 
	/*
	io.Writer :
	type Writer interface {
	   Write(p []byte) (n int, err error)
	}
	*/
	tw := tar.NewWriter(&buf);

	files := []file {
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	};

    //文件写入
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	//文件读取
	tr := tar.NewReader(&buf)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}

func tar2() {
	
}