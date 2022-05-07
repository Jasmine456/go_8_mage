package main

import "fmt"

/*
3.有一个长度为20的int数组，分别用两个for循环求前半部分的和和后半部分的积，要求只遍历一次数组
 */

func main(){
	arr := [20]int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19}
	len_arr:=len(arr)
	half_len_arr:=(len_arr/2)
	//fmt.Println(half_len_arr,len_arr)
	var sum,product int
	for i:=0;i<half_len_arr;i++{
		sum+=arr[i]
	}
	for j:=half_len_arr;j<len_arr;j++{
		//fmt.Println(arr[j])
		//product *=arr[j]
		product = product+arr[j]
	}
	fmt.Println(sum,product)
	//for i,ele := range arr{
	//	sum +=ele
	//
	//}
}
