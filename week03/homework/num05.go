package main

import "fmt"

/*
5.给定月份，判断属于哪个季节。分别用if和switch实现
*/

func if_season(month uint8)string {
	if month <= 3 {
		fmt.Printf("%d属于春季", month)
	} else if month <= 6 {
		fmt.Printf("%d属于夏季", month)
	} else if month <= 9 {
		fmt.Printf("%d属于秋季", month)
	} else if month <= 12 {
		fmt.Printf("%d属于冬季", month)
	} else {
		fmt.Printf("%d输入月份为非法数字，请输出正确的月份，1-12！")
	}
	return ""
}

func switch_season(month uint8)string {
	switch {
	case month <= 3:
		fmt.Printf("%d属于春季", month)
	case month <= 6:
		fmt.Printf("%d属于夏季", month)
	case month <=9:
		fmt.Printf("%d属于秋季",month)
	case month <=12:
		fmt.Printf("%d属于冬季",month)
	default:
		fmt.Printf("%d输入月份为非法数字，请输出正确的月份，1-12！")
	}
	return ""
}

func main(){
	fmt.Println(if_season(1))
	fmt.Println(switch_season(5))
}
