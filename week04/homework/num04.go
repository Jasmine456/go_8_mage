package main

import "fmt"

/*
实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，
interface{}允许是4种类型：float32、float64、int、byte

 */

func square(num interface{}) interface{}{
	var result interface{}
	switch v:= num.(type) {
	case int:
		result=v*v
	case byte:
		result=v*v
	case float64:
		result=v*v
	case float32:
		result=v*v
	default:
		fmt.Println("请重新选择传入的类型，目前只允许是4种类型：float32、float64、int、byte")
		return nil
	}
	return result
}

func main(){
	fmt.Println(square(true))
}