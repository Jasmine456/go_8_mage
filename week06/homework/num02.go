package main

import (
	"fmt"
	"time"
)

/*
2. 我们是每周六上课，输出我们未来4次课的上课日期（不考虑法定假日）
 */

const (
	TIME_FMT = "2006-01-02 15:04:05"
	DATE_FMT = "20060102"
)

func main(){
	now:=time.Now()
	fmt.Println(int(now.Weekday()))
	diff:=6-int(now.Weekday())
	oneday := time.Duration(24*time.Hour)
	d := time.Duration(time.Duration(diff)*oneday)
	num01 := now.Add(d)
	fmt.Printf("第 %d 周上课的时间为 %s\n",1,num01.Format(TIME_FMT))
	var date time.Time
	for i:=1;i<4;i++{
		j:=time.Duration(i)*7*oneday
		date =num01.Add(time.Duration(j))
		fmt.Printf("第 %d 周上课的时间为 %s\n",i+1,date.Format(TIME_FMT))
	}
}