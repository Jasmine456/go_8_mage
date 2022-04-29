package main

import (
	"fmt"
	"math/rand"
)

/*
创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，
利用map统计该切片里有多少个互不相同的元素。

 */


func main(){
	//var s []int
	s := make([]int,0,10)
	m := make(map[int]int,100)
	count:=100
	// 将100个随机数存到 切片 s中
	for i:=0;i<count;i++{
		s = append(s, rand.Intn(128))
	}
	//fmt.Printf("%v,%d,%d\n",s,len(s),cap(s))

	// 将切片按照0-99的索引 存到一个map m
	for i,j := range s{
		//fmt.Println(i,j)
		m[j]=i
	}
	//for key,value:= range m{
	//	fmt.Printf("%d,%d\n",key,value)
	//}
	fmt.Printf("统计该切片里有 %d 个互不相同的元素",len(m))

}