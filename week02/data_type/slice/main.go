package main

import "fmt"

func slice_init() {
	var s []int //声明切片 len=cap=0
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = []int{} //初始化，len=cap=0
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = make([]int, 3) //初始化，len=cap=3
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = make([]int, 3, 5) //初始化，len=3 cap=5
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = []int{1, 2, 3, 4, 5} // 初始化，len=cap=5
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	fmt.Println("==================================")

	//	二维切片初始化的一种方式
	s2d := [][]int{
		{1},
		{2, 3}, // 二维数组各行的列数是相等的，但二维切片各行的len可以不相等
	}
	fmt.Printf("s2d len %d cap %d\n", len(s2d), cap(s2d))
	fmt.Printf("s2d[0] len %d cap %d\n", len(s2d[0]), cap(s2d[0]))
	fmt.Printf("s2d[1] len %d cap %d\n", len(s2d[1]), cap(s2d[1]))
	fmt.Println("===============================")

}

func slice_append(){
	arr := make([]int,3,6)
	brr := append(arr,8) //arr和brr共享地城数组，但他们的len不同
	brr[0] =9

	fmt.Printf("arr[0]=%d,cap of arr %d,len of arr %d\n",arr[0],cap(arr),len(arr))
	fmt.Printf("brr[0]=%d,cap of brr %d,len of brr %d\n",brr[0],cap(brr),len(brr))

	s := make([]int,3,5)
	for i:=0;i<3;i++{
		s[i] = i+1
	}// s=[1,2,3]
	fmt.Printf("s[0] address %p,s=%v\n",&s[0],s)

	/*
		capacity还够用，直接把追加的元素放到预留的内存空间上
	 */
	s = append(s,4,5) // 可以一次append多个元素
	fmt.Printf("s[0] address %p,s=%v\n",&s[0],s)
	/*
		capacity不够用了，得申请一片新的内存，把老数据先拷贝过来，在新内存上执行append操作
	*/
	s = append(s,6)
	fmt.Printf("s[0] address %p,s=%v\n",&s[0],s)
	fmt.Println("=======================")



}





func sub_slice(){
	/*
		截取一部分,创造子切片，此时子切片与母切片(或母数组)共享地城内存从空间，母切片的capacity子切片可能直接用
	*/
	s := make([]int,3,5)
	for i := 0;i<3;i++{
		s[i] = i+1
	}// s=[1,2,3]
	fmt.Printf("s[1] address %p\n",&s[1])
	sub_slice := s[1:3]
	fmt.Printf("len %d cap %d\n",len(sub_slice),cap(sub_slice))
	/*
		母切片的capacity还允许子切片执行append操作
	 */
	sub_slice = append(sub_slice,6,7) //可以一次append多个元素
	sub_slice[0] = 8
	fmt.Printf("s=%v,sub_slice=%v,s[1] address %p,sub_slice[0] address %p\n",s,sub_slice,&s[1],&sub_slice[0] )
	/*
		母切片的capacity用完了，子切片再执行append就得申请一片新的内存，把老数据先拷贝过来，在新内存上执行append操作
	 */


	arr := [5]int{1,2,3,4,5}
	fmt.Printf("arr[1] address %p\n",&arr[1])
	sub_slice = arr[1:3] // 从数组创造子切片，len=cap=2
	fmt.Printf("len %d cap %d\n",len(sub_slice),cap(sub_slice))
}


func main() {
	//slice_init()
	//slice_append()
	sub_slice()
}
