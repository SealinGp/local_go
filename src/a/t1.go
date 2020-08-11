package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)
var (
	domain    = "http://www.xyw234.com/"
	formUrls  = []string{
		domain + "forum.php?mod=forumdisplay&fid=334", //shenzhen QM url (fid=334)
	}
	Cookies = "uFLb_2132_saltkey=dN2PhTP8; uFLb_2132_lastvisit=1592916904; PHPSESSID=t9uhgkrfpofr6b7h6frepuq2i6; uFLb_2132_ulastactivity=f4a70%2FXobXU44LIOP47MpSdoO%2FK4rAL8%2FGTyV7HBBHXCFhnJp%2B6r; uFLb_2132_auth=605dHSJfMiIIkVrywT%2BI1wiw3RGaxJfYMB0pfEx740pW%2BH1DBc3Hh0S7%2Fqgj1ixavxgFWBLiPmMoKfMpPlLfSbAoWQ; uFLb_2132_lastcheckfeed=58139%7C1592920891; uFLb_2132_sid=hAwR6r; uFLb_2132_lip=104.128.93.177%2C1592921022; uFLb_2132_visitedfid=334; uFLb_2132_viewid=tid_36183; uFLb_2132_st_p=58139%7C1592921108%7C4e0b5ab58685f0febd27bfab949be7a2; uFLb_2132_lastact=1592922076%09forum.php%09forumdisplay; uFLb_2132_st_t=58139%7C1592922076%7Cd4738fcdf63099183887bbcc755bc903; uFLb_2132_forum_lastvisit=D_334_1592922076"
)

type QmInfo struct {
	Name    string
	Area    string
	Price   string
	Project string  //项目
	AssessRes string  //评价
	WeChat  string
	Phone   string
}

func main() {
	publish := &pub{
		mu:   sync.RWMutex{},
		subs: make(map[sub]topicFunc),
	}
	var wg sync.WaitGroup
	sub1 := publish.Sub(func(topic interface{}) bool {
		return true
	})
	sub2 := publish.Sub(func(topic interface{}) bool {
		if s,ok := topic.(string);ok {
			return strings.Contains(s,"golang")
		}
		return false
	})


	wg.Add(2)
	go func() {
		defer wg.Done()
		for msg := range sub1 {
			fmt.Println("sub1:",msg)
		}
	}()
	go func() {
		defer wg.Done()
		for msg := range sub2 {
			fmt.Println("sub2:",msg)
		}
	}()

	publish.SendTopic("hello")
	publish.SendTopic("hello golang")
	publish.Close()

	wg.Wait()
}
type sub chan interface{}
type topicFunc func(topic interface{}) bool
type pub struct {
	mu sync.RWMutex
	subs map[sub]topicFunc
}

func (this *pub)Sub(tf topicFunc) chan interface{} {
	curPub := make(chan interface{})

	this.mu.Lock()
	defer this.mu.Unlock()
	this.subs[curPub] = tf
	return curPub
}
func (this *pub)SendTopic(v interface{})  {
	this.mu.RLock()
	defer this.mu.RUnlock()
	var wg sync.WaitGroup
	for sub,top := range this.subs {
		if top(v) {
			wg.Add(1)
			go this.sendTopic(sub,v,&wg)
		}
	}
	wg.Wait()
}
func (this *pub)sendTopic(s sub,v interface{},wg *sync.WaitGroup)  {
	defer wg.Done()
	select {
	case s <- v:
		fmt.Println("send success!")
	case <-time.After(time.Second*30):
		fmt.Println("send time out after 30s")
	}
}
func (this *pub)Evit(sub chan interface{})  {
	this.mu.Lock()
	this.mu.Unlock()
	delete(this.subs,sub)
	close(sub)
}
func (this *pub)Close()  {
	this.mu.Lock()
	this.mu.Unlock()
	for sub := range this.subs {
		delete(this.subs,sub)
		close(sub)
	}
}

func Pipei()  {
	//匹配到所有详情页
	AllDetailPages := []string{
		"forum.php?mod=viewthread&tid=36183",
	}
	//for _, v := range formUrls {
	//	lists := re(getContent(v),`forum.php\?mod=viewthread&tid=\d+`)
	//	if len(lists) > 0 {
	//		AllDetailPages = append(AllDetailPages,lists...)
	//	}
	//}

	Qms := make([]QmInfo,0)
	for _, v := range AllDetailPages {
		content := getContent(domain + v + "&extra=page=1")
		info    := QmInfo{}

		name    := re(content,`常用昵称： \p{Han}+</span>`)
		if len(name) > 0 {
			info.Name = strings.ReplaceAll(name[0],"</span>","")
		}
		area    := re(content,`所在区域： \p{Han}+</span>`)
		if len(area) > 0 {
			info.Area = strings.ReplaceAll(area[0],"</span>","")
		}
		price    := re(content,`消费价格： \p{Han}+</span>`)
		if len(price) > 0 {
			info.Price = strings.ReplaceAll(price[0],"</span>","")
		}
		project  := re(content,`服务项目： \p{Han}+</span>`)
		if len(project) > 0 {
			info.Project = strings.ReplaceAll(project[0],"</span>","")
		}
		assess := re(content,`评价：[\p{Han}a-zA-Z\d\+\|]+</div>`)
		if len(assess) > 0 {
			info.AssessRes = strings.ReplaceAll(assess[0],"</div>","")
		}
		weChat := re(content,`微信号码： [\p{Han}\sa-zA-Z\d]+</div>`)
		if len(weChat) > 0 {
			info.WeChat = strings.ReplaceAll(weChat[0],"</div>","")

		} else {
			//TODO
			//获取WeChat的函数
		}

		phone := re(content,`联系电话： [\p{Han}\sa-zA-Z\d]+</div>`)
		if len(phone) > 0 {
			info.Phone = strings.ReplaceAll(phone[0],"</div>","")
		} else {
			//TODO
			//获取联系方式的函数
		}


		Qms = append(Qms,info)

	}
	//str := `<em>[<a href="forum.php?mod=forumdisplay&fid=334&amp;filter=typeid&amp;typeid=197">龙岗区</a>]</em> <a href="forum.php?mod=viewthread&amp;tid=36183&amp;extra=page%3D1" onclick="atarget(this)" class="s xst">小御姐范的36C妍妍</a>`
	//re(str,`href=(\"([^<>"\']*)\"|\'([^<>"\']*)\')`)
	fmt.Println(Qms)
}

//正则匹配
func re(str,pattern string) []string  {
	reg  := regexp.MustCompile(pattern)
	return reg.FindAllString(str,-1)
}

//获取内容
func getContent(url string) string {
	req,_   := http.NewRequest("GET",url,nil)
	Cookies  = strings.TrimSpace(Cookies)
	Cookies1 := strings.Split(Cookies,";")
	for _, v := range Cookies1  {
		i := strings.IndexRune(v,'=')
		req.AddCookie(&http.Cookie{
			Name  : strings.TrimSpace(v[:i]),
			Value : strings.TrimSpace(v[i+1:]),
			Path  : "/",
			Domain:"www.xyw234.com",
		})
	}
	resp,_ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	respBy,_ := ioutil.ReadAll(resp.Body)
	return string(respBy)
}