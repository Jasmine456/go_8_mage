package main

import "fmt"

/*
6.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，
求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
 */


type Student struct{
	Name string
	English float32
	Chinese float32
	Math float32
}


func main(){
	var s []Student
	s = []Student{}

	var arg,three_arg_ch,three_arg_en,three_arg_ma,sum_ch,sum_en,sum_ma float32
	flunk := []string{}
	for i:=0;i<len(s);i++{
		arg=(s[i].Chinese+s[i].English+s[i].Math)/3
		fmt.Printf("%s同学的三门课的平均分为%.2f",s[i].Name,arg)
		sum_ch += s[i].Chinese
		sum_en += s[i].English
		sum_ma += s[i].Math
		if arg<60{
			flunk=append(flunk,s[i].Name)
		}
	}
	three_arg_en=(sum_en/float32(len(s)))
	three_arg_ma=(sum_ma/float32(len(s)))
	three_arg_ch=(sum_ch/float32(len(s)))
	fmt.Printf("全班同学平均分低于60的有 %d位",len(flunk))



}
