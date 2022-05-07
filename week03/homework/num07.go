package main

import "fmt"

/*
7.定义一个5行4列的float数组，把里面的元素全部放到一个一维切片里
 */
func main(){
	arr := [5][4]float32{}
	s := []float32{}
	fmt.Println(len(arr),cap(arr))
	for row,array := range arr{
		for col,ele := range array{
			fmt.Printf("arr[%d][%d]=%d",row,col,ele)
			s=append(s,ele)
		}
	}
	fmt.Println(s)
}
