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

// https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters/
// 1.字母只能用一次
func (*Ref)CountC()  {
	words   := []string{"cat","bt","hat","tree"}
	chars  := "atach"

	charts1 := chars
	remNum  := 0
	for _, word := range words {
		//是否掌握了
		notRem := false
		//每次拼写
		charts1 = chars

		for _, w := range word {
			i := strings.IndexRune(charts1,w)
			if i == -1 {
				notRem = true
				break
			}

			//每个字母只能用一次
			charts1 = charts1[:i] + charts1[i+1:]
		}
		if !notRem {
			remNum += len(word)
		}
	}

	fmt.Println(remNum)
}

// https://leetcode-cn.com/problems/longest-palindrome/
func (*Ref)LP()  {
	
}
// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
func (*Ref)ZXDKGS()  {
	arr  := []int{3,2,1}
	k    := 2

	arr1 := []int{}
	for k > 0 {
		min  := arr[0]
		minI := 0
		arrL := len(arr)
		for i := 1;i < arrL; i++  {
			if min > arr[i] {
				min  = arr[i]
				minI = i
			}
		}

		tmp := []int{}
		tmp = append(tmp,arr[:minI]...)
		tmp = append(tmp,arr[minI+1:]...)
		arr = tmp
		k--

		arr1 = append(arr1,min)
	}

	log.Println(arr1)
}

// https://leetcode-cn.com/problems/water-and-jug-problem/
func (*Ref)Water()  {
	x, y, z            := 3, 5, 4
	//remain_x: x中的水量
	remain_x, remain_y := x, y

	//存储已经搜索过的所有的remain_x/remain_y 的状态
	stack := [][]int{
		{0,0},
	}

	for stack != nil {

	}
}