package main

import "fmt"

/*
截取子切片、apped返回新切片，都会导致内存共享，扩容后才会分离
 */
func main()  {
	s:= make([]int,3,5)
	//for i :=0;i<3;i++{
	//	s[i]=i+1
	//}
	sub_slice:=s[1:3]
	fmt.Printf("s len:%d,cap:%d\n",len(s),cap(s))
	fmt.Printf("sub_slice len:%d,cap:%d\n",len(sub_slice),cap(sub_slice))
	fmt.Printf("%p\n",&sub_slice[0])
	sub_slice = append(sub_slice, 1)
	fmt.Printf("%p\n",&sub_slice[0])
	//fmt.Println(s[3])
	fmt.Println(sub_slice[0])




}
