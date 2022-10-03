package main

import "log"

func main() {

	arr := [...]int{0, 0}
	log.Printf("%v", arr)
	v(arr)
	log.Printf("%v", arr)
}

func v(arr [2]int) {
	arr[0] = 1
}
