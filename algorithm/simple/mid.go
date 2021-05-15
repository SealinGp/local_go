package simple

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//1.倒水 2.车油满最大行程距离,油量
//https://leetcode-cn.com/problems/water-and-jug-problem/
func (*Ref) CMW() {
	fmt.Println(cal(3, 7))
}
func cal(x, y float32) float32 {
	z := y - x
	return x + z*rand01()
}
func rand01() float32 {
	rand.Seed(time.Now().Unix())
	return rand.Float32()
}

func (*Ref) DS() {
	s := "3[z]2[2[y]pq4[2[jk]e1[f]]]ef"
	lefti := 0
S:
	for i, s1 := range s {
		if s1 == '[' {
			lefti = i
		}
		if s1 == ']' {
			startNum := lefti - 1
			for {
				startNum--
				if startNum < 0 {
					startNum++
					break
				}

				if s[startNum] >= 'a' && s[startNum] <= 'z' {
					startNum++
					break
				}
				if s[startNum] == '[' {
					startNum++
					break
				}
			}

			num1, _ := strconv.Atoi(s[startNum:lefti])
			old := s[startNum : i+1]
			news := strings.Repeat(s[lefti+1:i], num1)
			s = strings.Replace(s, old, news, 1)
			goto S
		}
	}
	fmt.Println(s)
}
