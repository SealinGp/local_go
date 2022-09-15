package main

import (
	"log"
	"math"
)

func main() {
	// log.Printf("%v", getMin(")))()()()())())())))))())))))))())()))()()()))))))((("))
	log.Printf("%v", stockPairs([]int32{
		407,
		1152,
		403,
		1419,
		689,
		1029,
		108,
		128,
		1307,
		300,
		775,
		622,
		730,
		978,
		526,
		943,
		127,
		566,
		869,
		715,
		983,
		820,
		1394,
		901,
		606,
		497,
		98,
		1222,
		843,
		600,
		1153,
		302,
		1450,
		1457,
		973,
		1431,
		217,
		936,
		958,
		1258,
		970,
		1155,
		1061,
		1341,
		657,
		333,
		1151,
		790,
		101,
		588,
		263,
		101,
		534,
		747,
		405,
		585,
		111,
		849,
		695,
		1256,
		1508,
		139,
		336,
		1430,
		615,
		1295,
		550,
		783,
		575,
		992,
		709,
		828,
		1447,
		1457,
		738,
		1024,
		529,
		406,
		164,
		994,
		1008,
		50,
		811,
		564,
		580,
		952,
		768,
		863,
		1225,
		251,
		1032,
		1460}, 1558))

	//1,46,

}

// 92
// 1558

func stockPairs(stocksProfit []int32, target int64) int32 {
	var pairNum int32

	pairsMap := make(map[int32]bool)

	for _, profit := range stocksProfit {
		_, ok := pairsMap[profit]

	}

	for i := 0; i < len(stocksProfit); i++ {
		v := stocksProfit[i]
		dst := target - int64(v)

		_, ok := pairsMap[v]
		if dst < 0 || ok {
			continue
		}

		for j := i + 1; j < len(stocksProfit); j++ {
			_, ok := pairsMap[stocksProfit[j]]
			if ok {
				continue
			}

			if dst == int64(stocksProfit[j]) {
				pairsMap[v] = true
				pairsMap[stocksProfit[j]] = true
				pairNum++
			}
		}
	}

	return pairNum
}

func getMin(s string) int32 {
	var left int32
	var right int32
	for _, c := range s {
		if c == '(' {
			left++
		}

		if c == ')' {
			right++
		}
	}

	log.Printf("left:%v right:%v s:%v", left, right, len(s))

	ret := int32(math.Abs(float64(left) - float64(right)))
	return ret
}
