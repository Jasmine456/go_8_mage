package main

import (
	"fmt"
	"reflect"
)

//*int 参数类型为 指向整数的指针，而不是整数
func showMemoryAddress(x *int){
	//本身就是指针，打印地址不需要 &这个符号，如果想使用指针指向的变量的值，而不是其内存地址，可在指针变量前面加上星号
	fmt.Println(x)
	return
}

func main(){
	k:=40
	fmt.Println(k)
	fmt.Println(reflect.TypeOf(k))// 检查变量类型

	//获取变量在计算机内存中的地址，可在变量名前加上&字符。
	fmt.Println(&k)

	//&k 引用的是变量k的值，值所在的内存地址
	showMemoryAddress(&k) //返回的地址是相同的

}
