package main

import "fmt"

func main() {
	//s1 := "hello"
	//s2 := "王晓静"
	//
	//for _, ele := range s1 {
	//	fmt.Println(ele)
	//}
	//fmt.Println()
	//for _, ele := range s2 {
	//	fmt.Println(ele)
	//}
	//
	//s3:=[]byte{20,22}
	//s4:=[]rune{55,55,55,56}
	//fmt.Println(s3,s4)
	//for _,ele:=range s3{
	//	fmt.Println(ele)
	//}
	s1:="My name is 焦明"
	//s1:="  My"
	arr := []byte(s1)
	brr := []rune(s1)
	fmt.Printf("last byte %d\n",arr[len(arr)-1])//string可以转换为[]byte或[]rune类型
	fmt.Printf("last byte %c\n",arr[len(arr)-1])//byte或rune可以转为string
	fmt.Printf("last rune %d\n",brr[len(brr)-1])
	fmt.Printf("last rune %c\n",brr[len(brr)-1])

	L:=len(s1)
	fmt.Printf("string len %d byte array len %d rune array len %d\n",L,len(arr),len(brr))
	for _,ele:=range s1{
		fmt.Printf("%c",ele)//string中的每個元素是字符
	}
	fmt.Println()
	for i:=0;i<L;i++{
		fmt.Printf("%d",s1[i])//【i】前面应该出现数组或切片，这里自动把string转成了【】byte 而不是[]rune
	}


}
