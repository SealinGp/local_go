package main

import (
	"certs-master/config"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	timeLayOut = "2006-01-02 15:04:05"
)

/*
time packages
*/
func init() {
	//fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
}
func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("lack param ?func=xxx")
		return
	}

	execute(args[1])
}
func execute(funcN string)  {
	funcMap := map[string]func(){
		"time1" : time1,
		"time2" : time2,
		"time3" : time3,
		"time4" : time4,
		"time5" : time5,
	}
	funcMap[funcN]()
}
func time1()  {
	now := time.Now()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a1 := r.Intn(26)
	a2 := string(rune(int('A') + a1)) + "_" + fmt.Sprint(time.Now().Unix())

	fmt.Println(
		strings.Split(now.Format(timeLayOut),":"),
		int(now.Month()),
		a2,
	)


	a := 5
	switch a  {
	case 5:
		fmt.Println("a = 5")
	case 6,7:
		fmt.Println("a = 6 or 7")
	default:
		fmt.Println("a = none")
	}

	/*a1 := []uint8{5,6,7}
	var c,d uint8
	switch c,d = a1[0],a1[1] {
	case c == 5:
		fmt.Println(a1[0])
	case d == 6:
		fmt.Println(a1[1])
	}*/
}
func time2()  {
	t := time.Time{}
	ts := "2019-11-07T23:59:59+08:00"
	t.UnmarshalText([]byte(ts))
	fmt.Println(t.Format(timeLayOut))
}

func time3()  {
	startTime  := "2019-02-01T00:00:00+08:00"
	startTime1 := &time.Time{}
	endTime    := "2019-04-30T23:59:59+08:00"
	endTime1   := &time.Time{}
	if err := startTime1.UnmarshalText([]byte(startTime)); err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := endTime1.UnmarshalText([]byte(endTime)); err != nil {
		fmt.Println(err.Error())
		return
	}

	timeUnit    := "week"
	var err error
	switch timeUnit {
	case "week":
		startTime1,err = getWeekDay(first,startTime1)
		if err != nil {
			fmt.Println("get start week day err:",err.Error())
			return
		}

		endTime1,err = getWeekDay(last,endTime1)
		if err != nil {
			fmt.Println("get start week day err:",err.Error())
			return
		}
	case "month":
		startTime1,err = getMonthDay(first,startTime1)
		if err != nil {
			fmt.Println("parse month start time err:"+err.Error())
			return
		}

		endTime1,err = getMonthDay(last,endTime1)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println(startTime1.Format(timeLayOut))
	fmt.Println(endTime1.Format(timeLayOut))
}

func time4()  {
	a,_ := time.ParseDuration("24h")
	b   := time.Hour * 24
	fmt.Println(a == b)
}

type WeekDayType int
const (
	first WeekDayType = iota  //获取给定时间的周一/月头
	last                      //获取给定时间的周日/月尾
)
//获取给定时间的当前周的 周一/周日 对应的日期
func getWeekDay(dayType WeekDayType,t *time.Time) (*time.Time,error) {
	var (
		currentWeekDay = t.Weekday()
		dayNum            int
		HourMinSecond     string
		compareWeekDay    time.Weekday
		compareWeekDayAbs time.Weekday
	)

	if dayType == first {
		HourMinSecond  = " 00:00:00"
		compareWeekDay = time.Monday
	} else {
		HourMinSecond = " 23:59:59"
		compareWeekDay = time.Sunday
	}
	for i := 0; i < 7; i++ {
		compareWeekDayAbs = currentWeekDay
		if dayType == first && compareWeekDayAbs < 0 {
			compareWeekDayAbs = 7 + compareWeekDayAbs
		}
		if dayType == last && compareWeekDayAbs > 6 {
			compareWeekDayAbs = compareWeekDayAbs - 7
		}
		if compareWeekDayAbs == compareWeekDay {
			break
		}
		if dayType == first {
			currentWeekDay--
		} else {
			currentWeekDay++
		}
		dayNum++
	}
	if dayType == first {
		dayNum = -dayNum
	}
	weekDay := t.AddDate(0,0,dayNum)
	var err error
	weekDay,err = time.Parse(config.DefaultTimeLayOut,weekDay.Format("2006-01-02") + HourMinSecond)
	return &weekDay,err
}
//获取给定时间的当月的 第一天/最后一天 对应的日期
func getMonthDay(dayType WeekDayType,t *time.Time) (*time.Time,error) {
	dayHourMinSecond := ""
	if dayType == first {
		dayHourMinSecond = "01 00:00:00"
	} else {
		dayHourMinSecond = "01 23:59:59"
	}
	
	monthDay,err := time.Parse(timeLayOut,t.Format("2006-01-") + dayHourMinSecond)
	if err != nil {
		return &monthDay,err
	}
	if dayType == last {
		monthDay = monthDay.AddDate(0,1,-1)
	}
	return &monthDay,nil
}

//获取给定时间的 当月最后一天的日期 = 当月有多少天
func getlastDay(t *time.Time) (time.Time,error) {
	monthEndTime,err := time.Parse(timeLayOut,t.Format("2006-01-") + "01" + t.Format(" 15:04:05"))
	if err != nil {
		return monthEndTime,err
	}
	monthEndTime = monthEndTime.AddDate(0,1,-1)
	return monthEndTime,nil
}

func time5()  {
	a := []interface{}{}
	b := []interface{}{
		"1","2","3",
	}
	a = append(a,&b)
	var c interface{} = a
	for index,v := range c.([]interface{})  {
		fmt.Println(index,v)
	}
}