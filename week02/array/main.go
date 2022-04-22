package main

import "fmt"

func basic(){
	var arr1 [5]int =[5]int{}
	var arr2 = [5]int{}
	var arr3 = [5]int{1,2,4}
	var arr4 = [5]int{2:8,3:8}
	var arr5 =[...]int{1,2,4}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
}

func arrPoint(arr *[5]int){
	fmt.Printf("%p\n",arr)
	arr[0]+=10
	fmt.Println(arr[0])



}

func array2d(){
	var arr1 = [5][3]int{{1,2},{3},{6,3,8}}
	//var arr2 = [...][3]int{{1,2},{6,3,8}}
	//fmt.Println(len(arr1))
	//fmt.Println(len(arr2))
	//fmt.Println(cap(arr1))
	//fmt.Println(cap(arr2))

	//for i ,row := range arr1{
	//	fmt.Printf("%T\n",row)
	//	for j,n := range row{
	//		fmt.Printf("%d %d %d\n",i,j,n)
	//	}
	//}

	for i:=0;i<len(arr1);i++{
		for j:=0;j<len(arr1[i]);j++{
			fmt.Printf("%d %d %d\n",i,j,arr1[i][j])
		}
	}
}


func for_range(){
	arr := [...]int{1,2,3,4,5}
	for i ,ele := range arr{// ele是arr里元素的拷贝
		fmt.Printf("%d %d\n",i,ele)
		arr[i] += 8
		fmt.Printf("%d %d\n",i,ele)
	}
}

func main(){
	//for_range()
	//array2d()
	//basic()
	//var crr [5]int = [5]int{1,2,3,6,9}
	//fmt.Printf("%p\n",&crr)
	//fmt.Printf("%p\n",&crr[0])
	//fmt.Printf("%p\n",&crr[1])
	//arrPoint(&crr)
	//fmt.Println(crr[0])
	var a int = 9
	var b *int
	b = &a //取址
	*b =3 // 解析指针

}