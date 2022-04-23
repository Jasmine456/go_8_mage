package main

import (
	"fmt"
	"strings"
)

func fuzhi() {
	s1 := "sdfs sdf sdfsdf 中加 有"
	s2 := "dsfsdf '\"sdf\tsdfdsf\n \\中加 右"
	s3 := `sdfdsf '"sdf dsfdsf
\中加 右"'`
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

func string_function(){
	s := "abcfsdfsd中国人"
	fmt.Println(len(s))
	fmt.Println(strings.Split(s,"abc"))
	fmt.Println(strings.Contains(s,"abc"))
	fmt.Println(strings.Contains(s,"abc"))
	fmt.Println(strings.Index(s,"abc"))
	fmt.Println(strings.LastIndex(s,"abc"))
	fmt.Println(strings.HasPrefix(s,"abc"))
	fmt.Println(strings.HasSuffix(s,"abc"))
	fmt.Println(strings.HasSuffix(s,"abc"))


	s1 := "中"
	fmt.Println(len(s1))
	var arr []byte = []byte(s1)
	fmt.Println(arr)

}

func concat(){
	s1 := "abc"
	s2 := "efg"
	s3:= "中2q"
	space := " "

	s4 := s1+s2+s3
	fmt.Println(s4)

	s5 := fmt.Sprintf("%s%s%s",s1,s2,s3)
	fmt.Println(s5)

	arr := []string{s1,s2,s3}
	s6 := strings.Join(arr,"")
	fmt.Println(s6)

	sb := strings.Builder{}
	sb.WriteString(s1)
	sb.WriteString(space)
	sb.WriteString(s2)
	sb.WriteString(space)
	sb.WriteString(s3)
	s7 := sb.String()
	fmt.Println(s7)

}

func string_diceng(){
	s := "a中"
	arr := []byte(s) //强制类型转换
	brr := []rune(s)

	fmt.Println(arr)
	fmt.Println(brr)
}

func main() {
	//fuzhi()
	////string_function()
	//concat()
	string_diceng()
}
