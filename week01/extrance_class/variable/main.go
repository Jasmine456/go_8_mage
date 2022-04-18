package main

import "fmt"

func constant(){
	{
		const PI float32 = 3.14
		fmt.Printf("PI=%f\n",PI)
	}
	{
		const (
			PI = 3.14
			E = 2.71
		)
		fmt.Printf("PI=%f,E=%f\n",PI,E)
	}
	{
		const(
			a = 100
			b // 100跟上一行的值相同
			c // 100
		)
	}
}
