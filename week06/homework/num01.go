package main

import (
	"fmt"
	"time"
)

/*
1. 把字符串1998-10-01 08:10:00解析成time.Time，再格式化成字符串199810010810
 */

var (
	loc *time.Location
)
const (
	TIME_FMT = "2006-01-02 15:04:05"
	DATE_FMT = "200601021504"
)

func init() {
	loc, _ = time.LoadLocation("Asia/Shanghai")
}

func main(){

	str := "1998-10-01 08:10:00"
	if str_t,err := time.ParseInLocation(TIME_FMT,str,loc);err !=nil{
		fmt.Printf("parseInlocation method failed %v\n",err)
	}else{
		fmt.Printf("%v,%T\n",str_t,str_t)
		//	再格式化成字符串199810010810
		str_new:=str_t.Format(DATE_FMT)
		fmt.Println(str_new)
	}

}