package main

import "fmt"

func default_value(){
	var a int = 6685454654
	var b byte = 255
	var f float32 = 465464.654564
	var t bool = true
	var s string = "中" // 双引号是string
	var r rune = '中'//int32 单引号是一个字符，背后默默地会转为rune
	var d rune = '☺'
	var p *int = &a //&表示取后面额变量的地址
	var q *string = &s

	fmt.Printf("a=%d\n",a) // 0
	fmt.Printf("b=%d\n",b) // 0
	fmt.Printf("f=%f\n",f) // 0.0
	fmt.Printf("t=%t\n",t) //false
	fmt.Printf("s=[%s]\n",s) // ""
	fmt.Printf("r=%d\n",r) // 0
	fmt.Printf("d=%d\n",d) // 0
	fmt.Printf("d=%d %x\n",p,p)  // %d 十进制 %b 二进制 %x十六进制
	fmt.Printf("d=%d %x\n",q,q)
	fmt.Printf("p=%p\n",p) // %p十六进制
	fmt.Printf("q=%p\n",q)

}

func main(){
	default_value()
}
