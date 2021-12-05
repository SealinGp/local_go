package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/copier"
)

/*
7.array slice []bytes string
*/

type cops1 struct {
	A time.Time
	B string
}
type cops2 struct {
	A string
	B string
}

func cop1() {
	c1 := cops1{
		A: time.Now(),
		B: "b1",
	}
	c2 := cops2{
		A: "2014-12-16 00:00:00",
		B: "b2",
	}
	err := copier.Copy(&c1, &c2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(c1)
}
