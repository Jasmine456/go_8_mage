package main

import "fmt"

//演示基础数据类型的默认值
func default_value(){
	var a int
	var b byte
	var f float32 = 2.4578
	var t bool
	var s string
	var r rune
	var arr []int

	fmt.Printf("default value of int %d\n",a)
	fmt.Printf("default value of byte %d\n",b)
	fmt.Printf("default value of float32 %.2f, %.3e,%g\n",f,f,f)
	fmt.Printf("default value of bool %t\n",t)
	fmt.Printf("default value of string [%s]\n",s)
	fmt.Printf("default value of rune %d [%c]\n",r,r)
	fmt.Printf("default value of slice %v,arr is nil %t\n",arr,arr==nil)

}

func scope(){
	var b byte =255
	var a int8= 127 //2^7-1
	var c uint8 = 255 //2^8-1
	var f float32 = 2.4578
	var t bool
	var s string
	var r rune
	var arr []int
	var m complex64
	m = complex(2,6)
	fmt.Printf("%T %T\n",real(m),imag(m)) // %t %T
	fmt.Printf("real %f imag %f \n",real(m),imag(m))
	fmt.Printf("%T\n",m)


	fmt.Println(a,c)
	fmt.Printf("default value of byte %d\n",b)
	fmt.Printf("default value of float32 %.2f, %.3e,%g\n",f,f,f)
	fmt.Printf("default value of bool %t\n",t)
	fmt.Printf("default value of string [%s]\n",s)
	fmt.Printf("default value of rune %d [%c]\n",r,r)
	fmt.Printf("default value of slice %v,arr is nil %t\n",arr,arr==nil)
}

type User struct{
	Name string
	Age int
}

func (self User) Hello(){
	fmt.Printf("my name is %s\n",self.Name)
}
type ms map[string]int

func (self ms)Say(){
	fmt.Printf("%s\n",self["hello"])
}

type mss = map[string]int


func main(){
	default_value()
}
