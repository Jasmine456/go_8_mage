package main

/*
实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error

*/

func product(others ...float64)interface{}{
	pro := float64(1.0)
	for _,ele :=  range others{
		pro = ele*pro
	}

	if pro == 0{
		return "error"
	}else{
		c := 1/pro
		return c
	}

}
