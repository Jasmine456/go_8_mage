package main

import (
	"fmt"
	"math"
	"strings"
)

//输出一个int32对应的二进制表示
func BinaryFormat(n int32) string {
	a := uint32(n)  //有符号转为无符号
	sb := strings.Builder{}  //把多个短字符串拼接成一个长字符串
	c := uint32(math.Pow(2, 31))  //最高位是1，其他位全是0
	for i := 0; i < 32; i++ {
		if a&c !=0{
			sb.WriteString("1")   //往sb后面拼接“1”
		} else {
			sb.WriteString("0") //往sb后面拼接“0”
		}
		c >>= 1 //"1" 往右移一位
	}
	return sb.String()//返回长度为32的字符串
}

func main()  {
	//c := uint32(math.Pow(2, 31))
	//d := uint32(math.Pow(2, 2))
	//fmt.Printf("c is %b, d is %b\n",c,d)
	fmt.Println(BinaryFormat(0))
	fmt.Println(BinaryFormat(100))
	fmt.Println(BinaryFormat(-100))
	fmt.Println(BinaryFormat(342553432))

}
