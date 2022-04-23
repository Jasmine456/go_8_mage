package main

import (
	"fmt"
	//"golang.org/x/tools/go/analysis/passes/stringintconv/testdata/src/a"
	"strings"
)

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


type byt = uint8

func str_func(){
	s := "hello, how  are you"
	fmt.Printf("%t\n",strings.HasPrefix(s,"he"))
}

func string_joint(){
	s1,s2,s3 := "aaa","bbb","ccc"
	S1 := s1+ " "+ s2 + " "+s3
	S2 := fmt.Sprintf("%s %s %s",s1,s2,s3)
	S3 := strings.Join([]string{s1,s2,s3}," ")

	sb :=strings.Builder{}
	sb.WriteString(s1)
	sb.WriteString(" ")
	sb.WriteString(s2)
	sb.WriteString(" ")
	sb.WriteString(s3)
	sb.WriteString(" ")
	S4 := sb.String()
	fmt.Println(S1,S2,S3,S4)

}

func main(){
	//string_joint()
	var a int
	var b bool
	var c [5]int
	var u User
	fmt.Printf("%p\n",&a)
	fmt.Printf("%p\n",&b)
	fmt.Printf("%p\n",&c)
	fmt.Printf("%p\n",&u)
}
