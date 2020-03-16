package simple

import (
	"fmt"
	"strconv"
)

//https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
func (*Ref)FindMid()  {
	nums1, nums2 := []int{3},[2]int{-2,-1}

	nums1C  := nums1[:]
	nums1CL:= len(nums1)
	i1     := 0
	for _, v2 := range nums2 {
		for ; i1 < nums1CL; i1++ {
			if v2 < nums1C[i1] {
				left   := make([]int,i1+1)
				copy(left,nums1C[:i1])
				left[i1] = v2
				left     = append(left,nums1C[i1:]...)
				nums1C   = left
				nums1CL  = len(nums1C)
				i1++
				break
			} else if i1 == nums1CL-1 {
				nums1C  = append(nums1C,v2)
				nums1CL = len(nums1C)
				i1++
				break
			}
		}
	}

	if nums1CL == 0 && len(nums2) > 0 {
		nums1C  = nums2[:]
		nums1CL = len(nums2)
	}

	var mid float64
	if nums1CL % 2 == 0 {
		index := nums1CL/2
		mid = (float64(nums1C[index]) + float64(nums1C[index-1])) / 2
	} else {
		mid = float64(nums1C[nums1CL / 2])
	}

	fmt.Println(mid)
}
func (*Ref)FindMid2()  {
}

func (*Ref)CompressString()  {
	S  := "aabcccccaaa"

	//SL := len(S)
	sc := ""

	var v1 rune
	v1L := 0
	for _, v := range S {
		if v1 != v && v1L != 0 {
			sc += string(v1) + strconv.Itoa(v1L)
			if v1 == 'c' {
				fmt.Println(string(v))
			}
			v1L = 0
		}

		v1 = v
		v1L++
	}

	//if SL < len(sc) {
	//	sc = S
	//}

	//fmt.Println(sc)
	fmt.Println(S)
}