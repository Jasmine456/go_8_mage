package main

import "fmt"

/*
实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error

*/

func product(args ...float64)(float64,error){
	//计算连乘
	//pro := float64(1.0)
	var pro float64=1
	for _,ele :=  range args{
		//pro = ele*pro
		pro *= ele
	}
	// 计算倒数
	if pro == 0{
		// 构造error的两种方式
		// return 0，errors.New("divide by 0")
		return 0,fmt.Errorf("divide by 0")
	}else{

		return 1/pro,nil
	}

}

func main(){
	tmp,error:= product(0)
	fmt.Println(tmp,error)
}