package main

import "fmt"

func update_array(arr [5]int){
	fmt.Printf("array in function,address is %p\n",&arr[0])
	arr[0]=888
}

func update_array2(arr *[5]int){
	fmt.Printf("array in function,address is %p\n",&((*arr)[0]))
	fmt.Printf("array in function,address is %p\n",&arr[0])
	arr[0]=888
}

func update_slice(sli []int){
	fmt.Printf("slice in function,address is %p\n",&sli[0])
	sli[0]=999
}

func update_slice2(sli *[]int){
	fmt.Printf("slice in function,address is %p\n",&(*sli)[0])
	(*sli)[0]=999
}

func main(){
	arr:=[5]int{1,2,3}
	sli:=[]int{1,2,3}
	update_array(arr)
	fmt.Println(arr[0],&arr[0])
	update_array2(&arr)
	fmt.Println(arr[0],&arr[0])
	update_slice(sli)
	fmt.Println(sli[0])
	update_slice2(&sli)
	fmt.Println(sli[0])
}
