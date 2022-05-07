package main

import "fmt"

/*
arr:=make([]int,3,4);brr:=append(arr,1)问arr里的元素是什么？arr和brr会相互影响吗？写代码验证一下
 */

func main(){
	arr := make([]int,3,4)
	brr:=append(arr,1)

//	问arr里的元素是什么？
//3个0

//	arr和brr会相互影响吗？
// 两个切片在len或者cap容量不超过4时，是会互相影响的，因为brr 引用了arr的内存地址
//但是append追加元素超过4时，两个切片就不再互相影响，内存地址就分离开来了
	fmt.Println(brr)
	fmt.Println(arr)
	arr = append(arr,2)
	fmt.Printf("%v,%p\n",brr,brr)
	fmt.Printf("%v,%p\n",arr,arr)
	arr = append(arr,3)
	fmt.Printf("%v,%p\n",brr,brr)
	fmt.Printf("%v,%p\n",arr,arr)


}
