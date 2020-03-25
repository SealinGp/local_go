package simple

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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

// https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/
func (*Ref)MaxDiv()  {
	//假设str1 为长的那个 str2为短的那个
	str1 := "LEET"
	str2 := "CODE"

	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	
	str3 := ""
	maxL := 0

	for i := range str2 {
		if  str2[:i+1] == str1[:i+1] &&
			strings.ReplaceAll(str1,str1[:i+1],"") == "" &&
			strings.ReplaceAll(str2,str1[:i+1],"") == "" {
			if maxL < len(str1[:i+1]) {
				str3 = str1[:i+1]
				maxL = len(str1[:i+1])
			}
		}
	}

	fmt.Println(str3)
	fmt.Println(maxL)
}

// https://leetcode-cn.com/problems/longest-palindromic-substring/
func (*Ref)LPalidromic() {
	s := "cyyoacmjwjubfkzrrbvquqkwhsxvmytmjvbborrtoiyotobzjmohpadfrvmxuagbdczsjuekjrmcwyaovpiogspbslcppxojgbfxhtsxmecgqjfuvahzpgprscjwwutwoiksegfreortttdotgxbfkisyakejihfjnrdngkwjxeituomuhmeiesctywhryqtjimwjadhhymydlsmcpycfdzrjhstxddvoqprrjufvihjcsoseltpyuaywgiocfodtylluuikkqkbrdxgjhrqiselmwnpdzdmpsvbfimnoulayqgdiavdgeiilayrafxlgxxtoqskmtixhbyjikfmsmxwribfzeffccczwdwukubopsoxliagenzwkbiveiajfirzvngverrbcwqmryvckvhpiioccmaqoxgmbwenyeyhzhliusupmrgmrcvwmdnniipvztmtklihobbekkgeopgwipihadswbqhzyxqsdgekazdtnamwzbitwfwezhhqznipalmomanbyezapgpxtjhudlcsfqondoiojkqadacnhcgwkhaxmttfebqelkjfigglxjfqegxpcawhpihrxydprdgavxjygfhgpcylpvsfcizkfbqzdnmxdgsjcekvrhesykldgptbeasktkasyuevtxrcrxmiylrlclocldmiwhuizhuaiophykxskufgjbmcmzpogpmyerzovzhqusxzrjcwgsdpcienkizutedcwrmowwolekockvyukyvmeidhjvbkoortjbemevrsquwnjoaikhbkycvvcscyamffbjyvkqkyeavtlkxyrrnsmqohyyqxzgtjdavgwpsgpjhqzttukynonbnnkuqfxgaatpilrrxhcqhfyyextrvqzktcrtrsbimuokxqtsbfkrgoiznhiysfhzspkpvrhtewthpbafmzgchqpgfsuiddjkhnwchpleibavgmuivfiorpteflholmnxdwewj"

	isPa := func(str string, le int) bool {
		for i := 0; i < le/2; i++ {
			if str[i] != str[le-i-1] {
				return false
			}
		}
		return true
	}

	m := ""
	l := 0
	sL := len(s)
	for i := 0; i < sL; i++ {
		for j := i + 1; j <= sL; j++ {
			Le := len(s[i:j])
			if Le > l && isPa(s[i:j], Le) {
				l = len(s[i:j])
				m = s[i:j]
			}
		}
	}
	fmt.Println(m)
}
func (*Ref)FindMid2()  {
}

func (*Ref)CompressString()  {
	S := "aabcccccaaa"
	SL := len(S)
	sc := ""
	var v1 rune
	v1L := 0
	for i:= 0; i < SL + 1; i++ {
		if (i == SL) || (v1 != rune(S[i]) && v1L > 0) {
			sc += string(v1) + strconv.Itoa(v1L)
			v1L = 0
		}

		if i == SL {
			continue
		}
		v1 = rune(S[i])
		v1L++
	}


	if SL <= len(sc) {
		sc = S
	}
	fmt.Println(sc)
}

// https://leetcode-cn.com/problems/surface-area-of-3d-shapes/
func (*Ref)SurfaceArea()  {
	grid := [][]int{
		{2,2,2},{2,1,2},{2,2,2},
	}

	s     := 0
	for i := range grid {
		for j,v := range grid[i]  {
			curS := 0
			if v > 0 {
				curS = 6*v - 2*(v-1)
			}
			s = s + curS

			//情况1 j相邻
			prevj := j-1
			previ := i
			if prevj >= 0 {
				min := grid[previ][prevj]
				if v < min {
					min = v
				}
				s = s - min * 2
			}

			//情况2 i相邻
			prevj = j
			previ = i-1
			if previ >= 0 {
				min := grid[previ][prevj]
				if v < min {
					min = v
				}
				s = s - min * 2
			}
		}
	}

	log.Println(s)
	//log.Println(prevAll)
}