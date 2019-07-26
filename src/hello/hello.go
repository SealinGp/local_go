package main

/*
import use
 1.
   import "xx"
   import "xx"
 2.
   import (
	"xx"
	"xx"
   )
*/
import (
   "fmt"
   "math"
   "math/rand"
)

func init() {
  fmt.Println("Content-Type:text/plain;charset=utf-8\n\n");
}

func main () {
  //prin()

//  fmt.Println(add(1,2))
   index1();
//    index2()
      // fmt.Println(index3(3))
}

func prin() {
/*
  package use
     xx.X() //ucfirst
*/
  fmt.Println("number : ",rand.Intn(10));
  fmt.Printf("printf: %g \n",math.Sqrt(7));
}

func add(a int,b int) int {
  return a + b;
}
func index1() {
    var a int;
    var b int;
    a = 1;
    b = 2;

    c:= 3;
    d:= 5;
   fmt.Println("a : ", a);
   fmt.Println("b : ", b);

   fmt.Println("c : ", c);
   fmt.Println("d : ", d);
}


func index2() {
  x,y := change("world","hello");
  fmt.Println(x,y);
}
func change(x,y string) (string,string) {
    return y,x;
}

//不建议这样子使用
func index3(a int)(x,y int) {
  x = a+1;
  y = a-1;
  return
}

