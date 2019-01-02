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



func main () {
  //prin()

  fmt.Println(add(1,2))
}

func prin() {
/*
  package use 
     xx.X() //ucfirst
*/
  fmt.Println("number : ",rand.Intn(10))
  fmt.Printf("printf: %g \n",math.Sqrt(7))
}

func add(a int,b int) int {
  return a + b;
}