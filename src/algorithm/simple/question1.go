package simple

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
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
/*func (*Ref)Water() {
	x, y, z := 3, 5, 4
	//remain_x: x中的水量
	remain_x, remain_y := x, y

	//存储已经搜索过的所有的remain_x/remain_y 的状态
	stack := [][]int{
		{0, 0},
	}

	for stack != nil {

	}
}*/

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

// https://leetcode-cn.com/problems/available-captures-for-rook/
func (*Ref)NRC()  {

	//R : [2,3]  //y轴位置 [7,3] ~ [0,3]
	board := [][]byte{
		{'.','.','.','.','.','.','.','.'},
		{'.','.','.','p','.','.','.','.'},
		{'.','.','.','R','.','.','.','p'},
		{'.','.','.','.','.','.','.','.'},
		{'.','.','.','.','.','.','.','.'},
		{'.','.','.','p','.','.','.','.'},
		{'.','.','.','.','.','.','.','.'},
		{'.','.','.','.','.','.','.','.'},
	}
	//是不是白车
	ifWhiteRook := func(b byte) bool {
		return b == 'R'
	}
	//是不是黑卒
	ifBlackPawn := func(b byte) bool {
		return b == 'p'
	}
	//是不是空格
	ifEmpty := func(b byte) bool {
		return b == '.'
	}
	//白车一次吃的黑卒的数量
	nums := 0

	//1.找白车的位置[x1,y1]
	x1, y1 := 0, 0
	out:for x := range board  {
		for y := range board[x]  {
			if ifWhiteRook(board[x][y]) {
				x1, y1 = x, y
				break out
			}
		}
	}

	//找x轴上的黑卒的位置[x2,y2] [0,y1] ~ [8,y1]
	for i := 0; i < 8; i++ {
		if ifBlackPawn(board[i][y1]) {

			x2,y2 := i, y1
			leftX  := x2
			rightX := x1
			if rightX < leftX {
				leftX, rightX = rightX, leftX
			}

			//[x2,y2] ~ [x1,y1]之间必须全为'.',否则不能吃到黑卒
			for i1 := leftX + 1; i1 < rightX ; i1++ {
				if !ifEmpty(board[i1][y2]) {
					nums--
					break
				}
			}
			//假设找到了
			nums++
		}
	}

	//找y轴上的黑卒的位置[x2,y2] [x1,0] ~ [x1,8]
	for j := 0; j < 8; j++ {
		if ifBlackPawn(board[x1][j]) {
			x2,y2 := x1,j

			topY    := y2
			bottomY := y1
			if topY < bottomY {
				topY, bottomY = bottomY,topY
			}

			//[x2,y2] ~ [x1,y1]之间必须全为'.',否则不能吃到黑卒
			for j1 := bottomY + 1; j1 < topY ; j1++ {
				if !ifEmpty(board[x2][j1]) {
					nums--
					break
				}
			}
			//假设找到了
			nums++
		}
	}


	log.Println(nums)
}

// https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/
func (*Ref)HGSX()  {
	deck := []int{
		1,2,3,4,4,3,2,1,
	}
	//分y组,每组x个, x >= 2时返回true
	//1 =< y <= deckL/2
	//1.x * y = deckL : x = deckL / y

	deckL := len(deck)
	ifw   := false
	if deckL < 1 {
		return
	}

	//计算每个整数的个数
	tmp := make(map[int]int)
	for _,d := range deck  {
		tmp[d]++
	}

	//找出分组的可能性
	for y := 1; y <= deckL/2; y++{
		if deckL % y == 0 {
			x   := deckL/y
			if x < 2 {
				continue
			}

			//每个整数a的个数v 满足 v%x == 0 才可分组成功
			manzu := true
			for _,v := range tmp {
				if v % x != 0 {
					manzu = false
					break
				}
			}
			if manzu {
				ifw = true
				break
			}
		}
	}

	log.Println(ifw)
}

// https://leetcode-cn.com/problems/short-encoding-of-words/
func (*Ref)MLE()  {
	v := []string{}
	a := `["bxrodah","injscbv","batvd","kkhtn","ejtptcj","abcnrz","lfoiyv","zmbuzem","biusr","smiort","jgytozv","ovfywf","hsdbaso","uvgtti","iozrjc","ehyma","hdvfpok","mmnkldn","wyhyzu","urrsv","lwvjqo","ntgnh","deckcte","btoxi","wxtsgu","pqpleg","yeamt","oqkyuh","eqerff","pjwhcbh","lteuzzk","rakudyb","sffyiwh","kgndhg","xuyqwdl","pgdffb","zlkiqjs","ekhkyn","vwrii","psuzf","aesakrw","wqiwql","labyxjf","sqzkvj","vymxjj","hdhfs","pxtuk","buyfzq","cemogsf","cyedcxm","iyzgipr","rxuba","buuwd","lwxyml","asppsty","tfkbkd","wndcq","ymuxucw","zmvqnz","ovpdkf","lopgbgr","allupcw","bosewz","jthic","qprlkud","uqaakyr","gnpei","mmizuuk","rjtyt","nqqghp","iarae","pgelplm","auozx","upctqzb","bqore","fsxgk","dccbk","bnjrsgx","mfcut","ifywcw","yuoewkq","rufliba","tyjbx","zvsuglc","xsyki","kthaejn","hcbkshi","dcerd","xkbdn","cjpuum","feibejn","ffier","hedlt","qgprst","sldeydu","qildg","leqajo","fxknf","cimgbnt","ytdhqet","zfynbku","stoeok","vbkpex","yrvzoi","pkfuop","qhynoji","wcgvbqe","ydefst","maiwzcl","ldfjyf","grfubze","akdinjy","xvojay","vxkmxkc","qamswut","ilkwwhr","kwyvnq","ifvhfcx","hdela","tezylp","blgtg","llbnit","igsdxfo","iedin","njljm","mifiqfa","jdpywdv","xjdkurf","flyck","htslrm","dadwda","btigwi","uudjmye","ghmvljd","ftpcy","cxnulqq","iqifv","fymqvy","clzltyq","vnutx","fzvnyin","oezmoyp","pmwuydh","vbkha","okvizm","qoseznu","qawnk","hgreqe","syrzxc","tpifez","sguzt","rrdkrt","xziqcix","qdyhua","asdxy","vruri","wkapfkr","cemcygx","aizapj","sgzqbt","qimbzko","lukdni","bqgau","wbxwf","oifqeuz","nrxymea","jfghsnh","aujcbw","fkikpw","vdcwlf","xaqdpo","ijfpho","rtwjor","xxsxqb","uviosh","uwcsk","dfsqws","verltrk","gfefjml","bgyqyiu","oagqwy","mqjbfaw","xypala","kytnm","djbkmx","nsqjjac","mijvgo","okmpj","tjdza","wvmzs","zkjkjs","tayali","sbkmk","ydnxw","diaovq","wdqcz","yxjlvbm","wswdq","eutvzz","jrhja","pobeybk","svzegoe","nrgjsm","xcahqb","sgvmfqd","itgfyu","cbeqt","zmsxo","appxc","xyjak","gfcdqk","cqmjs","ywnewza","ktgjvl","xblna","pqqft","zphyif","aavloy","frlhdfq","odkzpj","ozfzr","kytxywu","ssrzi","hwetxek","djntz","cjvwme","aenmqri","blnekyf","ftror","btnscbg","nwhpdg","nrnmrbv","prjsgbj","mskxd","anxldu","mdwbn","uvgjifr","zbigei","evoqf","fmsrno","kfrzgg","dwhgo","wmogp","ifgyv","jjkeqjm","qxlrwun","insjvqj","jlomam","zjiyow","mxxpd","zauhwl","rofzd","jknyrsw","qdirr","yvwdp","uqclps","sggbkdy","lngqvtp","ukgutmh","fdtqhnj","wfpry","gimboui","ejcqu","bwgrbcn","rjitbzj","pjcbj","zmizwc","qymof","jklda","ydhzurn","aukquy","peezcqm","tlgaapy","dmnuowo","anzaxj","feyxg","agslyjq","vswqbm","qmfvg","lgsklb","ccfmeib","jedta","rshsdli","roazd","mddflzn","gmklkx","vovro","hdgoz","nndhga","grijs","urvrmd","zgjwn","qluokyn","ueauer","khejr","uxnua","dflqvg","ootqo","plkxdh","ofvheff","vezfrfj","wtvuuwy","dnqzgvf","ojecs","kbnxfxk","fgdgmd","bljnvyt","fdwpb","nehyyd","qytprev","ovvpj","entjs","fzyfoi","ojhycw","xzlpl","kczei","usvcsn","wniszq","itivc","ixklaxe","hxwdkc","gzopm","hotqdhn","pddxn","zfteijc","smpntqp","tjehf","safdbmc","lnbby","hbahfud","iuefhqb","msjzik","xzvhpyi","tnwgb","bdfchqy","lzxvu","oikwv","ewvaary","xgghujx","hecqblo","kmuavb","telqjy","nkjhmth","evzva","yxqdh","hfotn","cbltrue","ilbapdd","mpkhlty","bayrhf","qhjatlq","satoq","okyeq","zqllt","vqafo","rmasf","lbtpn","zbvghnx","pfkneg","mjydg","tfcfqj","eoevh","ljlgn","idrir","vgext","iejujm","hzabfsl","kqzlwlw","cvnbev","fitjzki","mnscnp","gnzczt","silaafj","qrvmd","wvszpks","jprnq","rhlbjz","pgivo","wlonmx","nnzeu","upxgoqf","ikreh","bizpj","tkzwszc","lububwr","wmvwt","xesfxcx","hqakbo","gpiitdt","wchhe","sjtcxt","jikubc","nloyg","lqygn","vuxqbif","gqydau","wxqxzj","ptfoftq","kvbaf","utmqm","dnsbja","fvufwez","kcaioa","lbfyajc","kxbvud","eduwvuc","ajbrdb","buibz","mgfla","cbflk","tapyt","havzca","fivnq","efofg","cshqqw","lblrj","csejne","dsdwkv","wliez","zggdp","ujdcxs","bbkul","rbgag","mrxdhtg","zmqgduf","gwjblr","cmyvk","iisblq","homomzk","phxpu","rpncnrm","telvbbq","dyqyr","twnupz","lrasieq","rrtiofs","xbuxn","cztuk","qzync","hcbmyuo","joiidxq","iptmn","tqwchnr","yfntld","yzftxl","hcrxgv","sgore","otpnj","vxnnp","khgbzo","qhizrg","oukfgup","gcgrr","uvowitq","asgqmj","bqvqx","fkjzpv","tuztpkz","hrprosx","akxwc","zsswtqa","nddje","nrngaj","cfufmm","uejroib","lmngezv","mnunzje","jnyhr","jrxnot","alaukfc","obfkup","xuwvdun","batidu","roleatw","hehyiay","eqvemdj","hxbjlw","gmljgl","dcefhpp","zgxyhz","rgatve","tdfdyyl","iatkd","nrvnis","evjwppw","jhdkz","krczt","nwdagp","hnqti","xtqqhn","jpabgod","igxihhi","ivpzao","ztgetd","upfjt","ibbyg","vfwjp","odnafv","lkbzd","tdtjsod","kuzmulu","uquieg","mxcwtum","cakhufs","dlnue","slomrw","hhnsie","bilnpyl","kztlznh","jynoymm","servsr","rmzzh","bshrn","wyuegvb","sdpjie","sgjdnd","hcajr","gdmjj","lgrmlnp","tcgkt","ejymia","vfjwfx","riskhc","kfdjm","hvdfilu","ymjje","gzeluio","ungod","fidxbu","pgvmqj","qkupp","ugajlm","qbcwf","msibbxf","ggxjw","ciaghu","qilbnon","lxgxfss","apgomj","umqwe","rsbndl","mnigehp","oevktmv","hmduhst","gtxds","veata","idsje","fpgtwai","apdsmuy","hitpb","uuecxw","ksckmfc","lekozhm","alsgjn","eiwjio","gxqeva","tyuzpbp","knzbj","evnvy","fovty","wblana","hnfgycl","grlwdx","bpszwxi","bqivb","xjkqccf","tmvck","prmfw","yrzlf","apihmw","atkhd","ulzzrha","wuiwpcx","cxhkfh","daunxlc","avinya","mwvwyaa","rokdbx","vzetdx","cvztvwh","umcft","dikyktg","kiobmtc","obprd","uvhqu","cecwsf","frnife","gkdbd","gudrxk","jmtes","xcvcmp","nzjlvs","jqjnntw","urdyxz","xkwgxss","dzszwj","lgkdmeg","nfoncmf","iftolsg","mmjgydf","fnrwjg","fflet","kkikx","nmaim","suknl","qczbbb","yctnxb","lrjxhj","fxyma","rkksxuy","kjewbj","ofwcbp","kpair","sgdgwcz","xovkgd","xcyxeaa","okrvpxh","fkvtz","iotean","xigmgpw","ntuctm","ytorbw","ygrbabw","wvcedge","ksekbfy","cvfbbyj","ehoawf","andnkxl","kcwfcln","rlhofob","lmfkyr","cfrxk","nddbgp","dsblvlp","hfvhbrr","ampjvw","etwbmif","mdvftdq","eblss","lqlvgxi","uiyfoq","ahwixk","vbpmicp","ivqegul","ydfdbu","fjggcct","yemjsgk","ouqud","gdzob","ghgje","czmfv","ektfe","oixruoo","hppbo","msojeli","pplbuc","dysto","fiaihu","ikujcmd","brdptz","bkjny","dcxle","ayzvfao","dhjihbk","lvggerl","fgrkn","gkivc","vkkttr","dbrtroq","pjfcusj","cinvj","vcdtgg","izzmtz","fztazu","iyuof","blmvdp","ygrfa","tidaf","opoutr","rlryhsr","yixkpuy","wezwpvd","ovdxf","ogspxs","hrgkukz","yqdtb","iaacv","zntck","pyojju","oecwtyt","qzhkh","sxjgcmz","mjwyfj","hlyyxec","rjemv","jqjto","ksqlvmp","uyrov","uswci","wggnzx","aefhc","tyfrlem","gkjas","stsuj","wnprqr","zfpfc","ybkcgxw","ystnth","kiwchiv","ojpvjz","lgpkdv","hoasw","lwvfoe","abxju","ouoeorm","lbpqt","msychc","nowip","kmhncli","otawf","mlzwp","zbmrg","yejrlyl","aidwaq","utbxa","dftzszr","yqmzxs","mliha","wqybjzf","siuarf","sikpt","dfqmbr","pauvt","mrggidv","csjfret","ejexa","konrrsl","zrdcii","twvsgxt","kkrhmwy","fwqik","hbgtdr","ezceq","wfxruo","dvjirgv","sdtkc","uxphrx","zficlv","evjjp","psheu","ljsjlfj","mwybn","uxcpg","pkvasp","stgxq","qbfpm","btnbadf","wqura","qfwgb","sfvfm","jlrwg","dtrxnst","tdlsspx","jdyof","foehi","axygkno","xjsiz","gujwjnb","wmworph","zlmsasd","jdcwqpd","szslaqy","fzxxm","urfdl","dnhpj","lqfsylq","scpmyra","qhazd","lkfoto","xvjrc","kshihv","tddvd","ztvyhvz","bwauiqe","hvdlkvu","ebnxmsq","mhhrvph","sznznie","sekveyo","ouisr","ahvsy","ohvhywj","wpbsk","ohmid","cgkrirx","nybiim","ggexr","lqqzayd","anvoj","ztscqp","nlsuotr","aclnvk","gjokyws","ablonc","aposd","youywmx","xaqdxph","qouqzp","pmcuff","hiqid","zevcwdv","gvipopo","ccgor","nkprjfq","iqvvdot","ypegl","liuibdu","yqmkm","weosf","dyluq","mvidimt","dxfpfmb","sruvp","jvgcig","visap","pcoojtr","ueezd","sepnlat","tpnfw","xnlyhbg","nticmy","ccowr","yfbggg","mpggqcw","ucolx","vhkeh","pxfvbqs","cbiaufq","cfvqybl","xtwvnq","tddkmn","otuxrtr","sahhi","elbxw","jwqbs","wgtik","mozacn","agxwpky","vplsmsf","vdjxp","wusop","xythvjv","vztaeon","znhhqq","vwdnn","vuvct","pqatl","pzwjt","corgzih","osbazmd","qvtawuq","qslpsdg","gbhvwmv","zlfxst","mktie","mlyvtpc","dxcxc","kurdpj","dakmtq","hfheula","ehamzyf","cperd","mskra","hihxhcb","ueyavzo","decgosg","fsbjdou","nvdzrwd","brlebf","nrkuds","tweiba","ksgkkt","krjexbn","cooajp","wkhjazt","zdiwy","xyijant","ekvpgg","ksscf","ujfvf","rwbuz","luvcogn","luous","pniaup","fdcbti","locthgt","xbqhdq","atgju","indjic","gqgkwa","mltour","vxkqvgl","yvjaos","kafof","lnxzyo","eqhnzu","kcycx","gayaz","acbtrk","nagnsb","vspbxp","udssjyr","srkmjx","xoprrui","czitbhh","hddqdn","pmiihyp","ndsmzh","mrncqge","oyuzx","fsahau","ztboe","gfdwe","skjfvgh","biimg","zglzs","mtdvfxf","gatond","kkcnm","pbgsbet","wzjrvcy","oizpkr","lsqfevr","rmkeosj","tkonjd","wygrty","zpxds","sgdwc","tftadrq","koiim","pmtme","uwfgo","abatasp","xcpyoe","iycmuiy","essyp","lugvk","kwmaek","yqiklfq","drydy","acfawu","wqmrmki","uckfwa","sylgk","nexzuoh","azqoown","mxuuoz","vqknme","uitgs","qaiyxqb","zjfll","hptpep","zosrgpj","wwrym","tuxna","wvishlo","xwojh","lzahy","pexfrzb","brltuj","fghut","bwpeib","clihctt","rbotev","cqkek","rwxqjce","hxedmx","ekowdj","ngjhxsh","fagrvwp","hiciueq","pdpyc","ojzule","mydwp","rfjvxj","gpbujk","mzhltd","lyemvb","bumzh","bekejst","wjjvjet","zgmuyt","gpdcqdh","awgzz","hoaeqzz","czvnkx","uoengv","umtgu","ikidx","wodokn","vuoywj","lbirgk","ypiqf","kplexu","dvwhv","tjjrrqr","vwmyx","cyvpu","bultr","popyjka","brlftz","lkkft","zsxiq","ginwqt","fbvciw","fjxtwv","emhxxd","lqbhzkg","lxsmr","wmiva","qgqrajs","zjcqba","znusf","xbpbn","mokbo","iebtm","reuoj","bdjdym","zputz","oiwpxk","ztzirri","rhgqkwp","louymi","ywovpsb","ezgpp","kqkugz","gcisp","uoetkk","qlidr","stuinea","ldqwuc","vzycdzv","zivfj","fsqls","uklezd","tweym","kwfxpzr","cqzrre","xmcbul","jzntdx","lkunway","fzuye","pecmsbk","ipkbc","gmwwau","cgfque","xgryh","snolr","rvidskm","gcnquh","xicknvt","ofqyihc","atylqsd","tefxj","dcpwzb","rxnloj","zvhoe","fgcyqp","ugatgrw","jciqar","dbysodw","vuuozv","gpqndu","kwqespi","mrxkb","oaiyi","ixyyrek","qufqjsy","fqulk","tpbzad","ypvznyn","escwvz","ptbmuhl","qfncv","giplyo","jxqtvzu","yvecbld","ihyzzkg","mqrjpk","ljftz","khxia","epcdpl","ynasjcq","jrksca","qqlsid","ghwmgld","dhirvfg","hctwla","owome","iibbjr","egjusya","wdghbb","rixilx","bmxhbxc","pegskz","srorl","ezvcgk","xobtez","kznoy","vvgms","einapez","kxujb","lxpnze","yywdbb","ieomgy","glqzkj","tddty","cbflar","azckmw","jrcls","dpwlb","ymntd","hzgiifc","jvglvf","dqopi","ktfeh","vqhjx","mmfwrak","acutyss","sitmhy","uicejxf","mkettn","cpdnu","odlqb","auvle","rkueogt","tlwbcxf","gsouvn","zonemh","nyksj","lpisxg","ldozs","zznsu","frraeen","cetjn","aoztut","lfpai","zprza","jiimgi","yngco","xanbtjr","xwkiwbx","tewclva","kefxt","watggo","rodjvxj","hyzbvw","swolxde","udfgn","wgvwa","xxhuw","tayrba","jcmmq","laqet","iekvehq","smootw","ifoyi","dccqp","ozkgzj","wxyhoy","zyyez","nieckoo","ewygj","geicytl","rqomq","bnxjnln","zmqwy","abeoy","elrbwqi","wpwglth","vtohjwu","ogqvzk","sednera","arkvuqq","boxdn","uzhpnhh","ntnygei","rsgqir","hupnds","zzyrr","mqowqj","kcldpbp","ndfoj","txymv","cxamllh","lzsohc","bksiwx","lpnab","dxbdksm","vrevzbm","csjkyo","desos","lujorcz","fpmktkm","kmdqncy","tkhon","agmvro","cwsurw","qsbgzz","qxsme","nxrng","fdaao","enffol","adyzt","xjagvnd","hiwems","xqedgqb","uulgo","cbcnl","dukqkz","pgvclh","hszecge","soaprrs","vralj","wachfq","qjuxrcj","vuwyxy","okkbi","cdlfl","wwzvrc","wjzcwi","zbbrb","gqzxgd","zuhfab","gjevjj","ioncvnn","thpmnzc","tklercp","hirykrm","xkikhcd","wrpqdk","rvnhg","zkziqy","aldwhd","dujvgxw","fqwzhj","jenkmz","idkcbc","xgbxv","mrurnpy","acltfuk","mdtco","yrqxn","hntcxg","nlyfsew","uztap","zucarxh","ajaod","tdlvxz","yrwxbn","xhzmmaa","iodcw","exhmtg","napen","quwvjx","bjavdwg","serqgkg","ohqesc","uvfwh","itcuoi","gknjx","lispmwy","gdmhxg","twvhigq","epqaheq","cspvf","tjzqe","smnlemk","kuufdl","uqrga","egcrcw","xhjflb","zdehco","pnpekyn","fbhmii","gdxmp","fftjcma","rrsaiux","iacquvj","ywlaox","kqciv","kxbegy","jurtqj","zihts","wzbwsa","sxzdxk","vmdlsvp","vznobz","flnbd","jvgcqm","amesqe","nrpwq","lorpy","lzjnzf","hvtuoo","wbntzlr","zmbchn","xoemiu","lxcvc","jxuub","bakjfd","rukala","xzgyd","qjbmeja","myckkyp","ztfrn","veigkh","fgced","vqvfz","uihnudn","geikyd","qtnjiwu","lpytbxp","ehwnt","owsdr","viehch","cskip","qrgwrsy","sfulobn","piwepyc","bvezw","iwzusic","heayql","adbchea","zenczyd","weoao","dudkax","uzghy","ggqvzc","glrax","lykmp","uotfxl","gbtdj","aeunj","mqzbagr","cixog","wpzxa","pdqzkov","vrvfq","sczonao","pfcgnt","djhsmjq","jibtys","bqzywc","ebbcbl","xusmfk","pewbx","jvwtck","tpfhtm","mpmjyv","nyzvwb","bfqnimj","rjmcodt","wkwqrcl","lhoul","psrqq","zjgpfl","oeclb","gakxtp","kdczy","mlwbzf","yovpi","wmvvh","iipupag","dxlmrul","wvaopyr","icqhrq","fwdwm","fklito","mmbsy","bmnrobi","axenxn","deyvs","aohdcl","rbqgf","bjdmki","qmbyx","qvgrf","jpbzt","stlehna","ntmpjvn","xitmvr","qelonf","gdnmyxw","jukckcb","pezauf","zkiik","ieqhny","dnsbgek","lxqokhz","bpbbdvd","rrrmnt","uzgge","lfzrcvj","wcsqx","uwgnn","tbieqez","grapl","pwummap","xaqgn","cajpmkc","pugsohx","pwcpo","xukxswe","bqdovgf","rhjxh","zgsvr","kqmbh","ebqjegk","bubdw","askohq","ieytkhc","mnawc","yunzk","dovqb","llhjba","kctrt","jzcuek","lfzvao","stwfzk","bwzynj","ikgwwg","fduto","nejkg","czyqc","illsrt","lhnkmpz","qqnwih","mbnzzfb","azflm","fgntjg","kbwgy","ijrgr","frnaozq","oiayl","qpzoau","awutn","hsubhkx","gnbvrl","rujodr","egkngk","nxzzv","uyevlah","fejhjri","qfchdh","rgweyq","knocd","ztefi","zuehzmy","zshxyib","ljbthca","quajc","vetran","dpvjxh","svaprrg","uksltzi","vvqnj","hcufj","cihosd","execr","bhirei","nnmrjun","bwavdc","pvvduw","jnyld","fhxtup","etcchic","yqnyi","jcfkimq","bwfdru","jgbqj","bonvd","qypbbqp","yanpj","znsow","pqynzj","qljno","mqfkp","egqbx","wodwxum","gpkgeqr","jtbhuny","yhzwu","pmbokms","xayorl","sbuab","axaklvx","znxfng","tiskjad","pssgdec","xrkoxuh","sssjcg","xcgfpkc","izrbx","ixypza","lfzssco","wrzqbz","ginxtod","phfii","tvidve","eqixznt","zvuykq","zudkd","dkfpvvq","hffwofk","dpgivqt","socyp","uklvaaj","lzovbn","hdgsxt","xporpei","zppwuz","ivaji","vtvgt","crbgr","razgdlr","vqrjq","pfesusz","zkteh","tblkhse","ncpau","apjzr","csyxoyw","rhezl","omqofm","uylzaw","fcsmld","aeaklhp","naummx","pezlqu","rlsxyc","vkicwic","whodrul","iewbzv","agkbg","unaonkl","lolgzbf","ybnvkr","torewb","vfeesxt","krxlj","ifahunt","wvitbya","cxquw","hkwbdls","fzmai","lttjadu","sfjqd","mjwfuc","dkqivy","mlrra","iutwrh","wfqzgf","mghnocn","mvabbq","xoxwom","roxmwfo","knvma","kppav","bpqpwp","xuihdw","fpimqkj","hfendm","mfrod","pfuzzdi","tpwjt","dsjlmwr","chlvt","ugunfs","bqupxlu","uqbwlu","lpgtq","qgrvb","gkrtqw","nfoqt","ovpmp","mrnsw","ppesk","ogcfjrb","mkxtajz","fulotuh","fncfo","chzzdvx","zclvvt","jepfaa","mbbmfr","syrtuv","uqmjid","dqsju","ukkvdra","xwaejz","wiykqn","sgeyt","yqzjv","ayojwy","xwumpyu","lcptjr","feovz","mowbx","kmzzyo","fezhn","zzxkgl","nffmxmx","nqait","luctl","czuxhjq","rqumlcb","wmrkivi","pxwzc","usoxagq","cidam","gmhlht","tkxwle","tscsi","pnbcw","ctbbfvl","khfzn","rxrobrd","uwyhkk","uuufwi","awvoiun","smaoam","jfbvgt","ffepme","ezsgli","pngjzc","exgkvb","hccsa","akcuiv","fekke","ojzjgt","ahhht","dulbw","xnecvj","otihgp","cdusknq","ppuce","oxeig","yvawsyu","lgtjsg","veaivz","xdkbnrh","cmuqd","ratlckz","xssbubg","blcixic","bawhz","hppnxp","hfewe","jpdcol","bsecwwh","wdevg","lziad","esvvty","sdyxxcv","gfplyvx","edppso","yfmhyy","haaxeao","qoqzkr","glkaoj","lsznj","qfqdy","whxjhjw","bkgrvn","vbbbrt","tybpoe","xjsmsh","mftoe","eohkflj","upwiem","xwhxo","yhrlzr","cquxyi","beupd","qrurus","ftzklm","abtog","ierpun","lvmpbp","qdqqcj","hdafx","fqglai","lrtrfua","latqxdu","zclnyri","hawxcr","jtnfvw","mubsczg","tqzyp","qpbisk","hjkbqu","jwmrsai","dmmtvmk","vphpjeh","euzymy","rplif","xmkmd","qdfhbff","sbugncv","eulte","gfborb","ugydm","ftwzgz","otzsiw","grafjlu","huvzy","ocxxjqm","ugyyfwq","ydvzr","acuygg","pwyhtc","hiaure","frasv","vbcvh","qlpsji","vvmuhx","dshyo","zolzf","hjtwyo","ycyztl","isixa","hiskwt","sdocf","wbqrs","hxktznr","gehec","uitar","gyiuq","nopyjkp","kaecqg","dhrloau","nnwmum","kpunxl","vlgqah","xaqijru","vjhuvh","gucrmja","njtnn","pfxfs","wpdlgdx","csxvc","sdzsx","veaeop","wdpubgz","buuaz","zberxo","tywikh","erxjzny","yltfas","shpbt","brhlth","rzmqosz","klvqwh","szvsacf","qcmmcwn","cyfzf","evwvu","lghbxe","mayqzcr","fndss","gqlnwsj","ofiscbi","togfs","ufbhc","mccir","lfhtm","pradz","dfqynfe","wxajfl","twtofj","qdejyap","dkjmvym","ytwqfuu","qoagsrk","chstp","tmdtvj","jybks","lucjfme","uiawh","dwgixlc","nuyoo","zztuzq","nrfrtqh","gfkfqfd","ctygsi","tbdztx","auagjb","eqtohll","azswd","mbfzv","zdtlhj","vgxvmvj","eswyx","xiiwm","gdjxthr","utcrz","veybyu","lcqedd","uncqy","ukcdml","dmatsz","monmrdc","vldmukn","wcvjfm","wrrzg","vthcv","tscrbrs","diocyuj","uxxew","rhhkuz","fjcrdiy","xhbvxuj","jcphq","tswpsu","sionh","lwiffoo","owlzrjr","qsvoko","qnworfs","qcxsgza","ekspb","osfbr","gsqcp","wsmqbwt","upxowu","kxwbipl","colmwdy","uxefw","caznqq","beltdxo","hwsbj","lwdtwl","yfdblmp","mlapxw","qinxsk","fymgjg","xbdpwn","tbauzz","lsbnwy","wqfksuj","fqgolg","rzqkos","lbaiqxb","escvlmg","zddlfn","jqrmbi","mqbxlyw","vqfukp","sfayeg","kxpxw","mxpfmw","exavuh","tqkaw","pabcd","grazud","nojsnxj","kznfxb","ncobwr","lcopum","qkrot","dchrr","wpzyvgy","vopbxf","ywyxbxs","zawfwua","ihqbhbw","qdzwnbz","onexy","fwxxa","snczd","whmrvl","gzwars","pnbfyyj","suhddk","yynvi","ubojvyd","klmmdc","phjarz","tuxdgq","ofxlrcr","hifgd","xctxnyr","eaadwz","nawgyo","kldqkq","jmsdtjm","taxxqpl","nukgwkg","vhnrhs","lpsmsk","qfptbkv","newtlen","llhahbx","xhles","jicvazl","vgodxds","trhmsru","ojswjok","gvwhte","lidxsq","iwodtsw","rznacne","pageb","qgrpas","qtvtg","otrqt","pxggabo","gbhidqv","rchkac","pxcip","hadulr","hxtxo","aagsmg","vonpa","cxvgn","jdznofs","xlchba","zrwlem","ceyffo","rzsymx","illfd","jsxjhi","mydir","sqtapve","fvtgt","njtsfm","tgzfaz","zbpqfs","ghmpgbl","iofobe","mclvf","lhrkk","frzekrg","mkhyoht","elixw","bivaxk","yhwfdd","nnfwrxf","ouqrtr","xnzxlc","zqfgzj","djpgpr","uhledcc","zxrzc","cklizo","wmltmow","noyucgl","plafl","rvluub","ptovdyw","ysiabsb","yfqiogz","pnopt","rcehki","yofru","guxevbv","mtiox","gobzzgf","yfyrww","zeszrob","jjqikre","ayiqif","ffvbst","ayckj","eycjw","dkvyzwr","vkwqrgf","rqicx","giwqcr","pmdzqw","rexodj","fzmltmx","ipfskur","qbdqiow","fqkixoa","levyca","qrhdp","cyksjxe","xxkqyg","jbcye","trttgf","gjqvu","oqiqrt","yxjbrr","iacng","hmoru","zcexy","iuwynb","nlmcld"]`
	json.Unmarshal([]byte(a),&v)

	words := v
	wordsMap := make(map[string]bool,len(words))
	for _, w := range words {
		wordsMap[w] = true
	}

	for _,w := range words  {
		for i := 1; i < len(w); i++ {
			_,e := wordsMap[w[i:]]
			if e {
				delete(wordsMap,w[i:])
			}
		}
	}

	l := 0
	for w := range wordsMap {
		l += len(w) + 1
	}
	log.Println(l)
}

// https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
/**
//https://blog.csdn.net/u011500062/article/details/72855826

m : 每次删除第m个数字
n : 一共有n个数字
假设最后剩下的数字 在序列长度为n的数组中 索引为x
找出每次删除的时候 n,m,x的关系即可解出方程
 */
func (*Ref)LR()  {
	n, m := 70866, 116922

	x  := 0
	tn := 1
	for i := 0;i <= n-1; i++ {
		x = (x + m) % tn
		tn++
	}

	log.Println(x)
}
// https://leetcode-cn.com/problems/sort-an-array/
func (*Ref)Sort()  {
	nums := []int{
		5,2,3,1,
	}

	nums = S(nums,len(nums))
	log.Println(nums)
}
func S(nums []int,L int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid   := nums[0]
	left  := []int{}
	right := []int{}
	for i := 1; i < L; i++ {
		if nums[i] < mid {
			left = append(left,nums[i])
		} else {
			right = append(right,nums[i])
		}
	}

	la := []int{}
	la = append(la,S(left,len(left))...)
	la = append(la,mid)
	la = append(la,S(right,len(right))...)
	return la
}

func (*Ref)MDAS()  {
	seq := "(()(())())"
	C   := 0
	D   := []int{}


	for i := 0;i < len(seq); i++ {
		if seq[i] == '(' {
			C++
			D = append(D,C%2)
		} else {
			D = append(D,C%2)
			C--
		}
	}

	log.Println(D)
}

// m × n
// 1:活细胞 0:死细胞

// = 3 活
// > 3 || < 2 死
//卷积运算
// https://leetcode-cn.com/problems/game-of-life/
func (*Ref)GOF()  {
	board := [][]int{
		{0,1,0},
		{0,0,1},
		{1,1,1},
		{0,0,0},
	}

	n    := len(board)
	m    := 0
	if n >= 1 {
		m  = len(board[0])
	}

	nextBoard := [][]int{}
	for i := range board  {
		rows := []int{}
		for j := range board[i]  {
			liveCells := 0

			for i1 := i-1; i1 <= i+1; i1++ {
				for j1 := j-1; j1 <= j+1 ; j1++ {
					//不包括自己
					if i1 == i && j1 == j {
						continue
					}
					//不能超过边界
					if j1 < 0 || j1 > m-1 {
						continue
					}
					if i1 < 0 || i1 > n-1 {
						continue
					}

					liveCells += board[i1][j1]
				}
			}


			nextCell := 0
			//假设为死细胞
			if board[i][j] == 0 && liveCells == 3 {
				nextCell = 1
			} else if board[i][j] == 1 {
				//活细胞
				nextCell = 1
				if liveCells > 3 || liveCells < 2 {
					nextCell = 0
				}
			}

			rows = append(rows,nextCell)
		}
		nextBoard = append(nextBoard,rows)
	}

	copy(board,nextBoard)
	log.Println(board)
}

//https://leetcode-cn.com/problems/lru-cache/
type LRUCache struct {
	kv   map[int]*node
	mu   sync.Mutex
	ca   int
	keys []int
}
type node struct {
	value      int
}
func Constructor(capacity int) LRUCache {
	return LRUCache{
		kv: make(map[int]*node,capacity),
		mu: sync.Mutex{},
		ca:capacity,
		keys:make([]int,capacity),
	}
}
func (this *LRUCache) Get(key int) int {
	this.mu.Lock()
	defer func() {
		this.mu.Unlock()
	}()
	if this.kv[key] == nil {
		return -1
	}
	this.keysSort(key,false)

	return this.kv[key].value
}
func (this *LRUCache) Put(key int, value int)  {
	this.mu.Lock()
	defer func() {
		this.mu.Unlock()
	}()

	if this.kv[key] != nil {
		this.kv[key].value = value
		this.keysSort(key,false)
		return
	}
	//超出阀值,删除旧的
	if len(this.kv) >= this.ca {
		delete(this.kv,this.keys[0])
	}
	this.keysSort(key,true)

	this.kv[key] = &node{
		value:      value,
	}
	return
}
func (this *LRUCache)keysSort(key int,delFirst bool)  {
	findI := this.ca
	if delFirst {
		findI = -1
	}
	for i,v := range this.keys {
		if v == key {
			findI = i
		}
		if i > findI && i-1 >= 0 {
			this.keys[i-1] = this.keys[i]
		}
	}
	this.keys[this.ca-1] = key
	return
}

func (*Ref)LRUCache()  {
	a := Constructor(3)
	a.Put(1,1)
	a.Put(2,2)
	a.Put(3,3)
	a.Get(2)
	a.Get(1)
	a.Put(4,4)
	fmt.Println(a.kv,a.keys)
}
func sort(k int,ks []int)  {
}