package main

//iota 网上练习参考链接
//https://www.cnblogs.com/bigdataZJ/p/go-iota.html

import (
	"fmt"
	//"go_8_mage/week01/extrance_class/util"
)

func constant() {
	{
		const PI float32 = 3.14
		fmt.Printf("PI=%f\n", PI)
	}
	{
		const (
			PI = 3.14
			E  = 2.71
		)
		fmt.Printf("PI=%f,E=%f\n", PI, E)
	}
	{
		const (
			a = 100
			b // 100跟上一行的值相同
			c // 100
		)
		fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
	}
	{
		const (
			a = iota //0
			b        //1
			c        //2
			d        //3
		)
		fmt.Printf("a=%d,b=%d,c=%d,d=%d\n", a, b, c, d)
	}
	{
		const (
			a = iota //0
			b = 50   //50
			c = iota //2
			d        //3
		)
		fmt.Printf("a=%d,b=%d,c=%d,d=%d\n", a, b, c, d)
	}
	{
		const (
			a = iota //0
			b        //1
			_        //2
			d        //3
		)
		fmt.Printf("a=%d,b=%d,d=%d\n", a, b, d)
	}
	{
		const (
			a = iota //0
			b = 50   //50
			c = iota //2
			d        //3
		)
		fmt.Printf("a=%d,b=%d,c=%d,d=%d\n", a, b, c, d)
	}
	{
		const (
			_  = iota
			KB = 1 << (10 * iota)
			MB = 1 << (10 * iota)
			GB = 1 << (10 * iota)
			TB = 1 << (10 * iota)
			EB = 1 << (10 * iota)
		)
		fmt.Printf("KB=%d,MB=%d,GB=%d,TB=%d,EB=%d\n", KB, MB, GB, TB, EB)
	}
	{
		const (
			a, b = iota + 1, iota + 2 //1,2
			c, d                      //2,3
			e, f                      //3,4
		)
		fmt.Printf("a=%d,b=%d,c=%d,d=%d,e=%d,f=%d\n", a, b, c, d, e, f)
	}
	{
		const (
			AA = iota     //0
			BB            //1
			_             //2
			DD = iota + 1 //3
			EE            //5
		)
		fmt.Printf("AA=%d,BB=%d,DD=%d,EE=%d\n", AA, BB, DD, EE)
	}
}

//literal 字面量
func literal() {
	fmt.Printf("%t\n", 04 == 4.00)      //用到了一个整型字面量和浮点型字面量
	fmt.Printf("%v\n", .4i)             //虚数字面量0.4i
	fmt.Printf("%t\n", '\u4f17' == '众') //Unicode和rune字面量
	fmt.Printf("%c\n", '众')             //众
	fmt.Printf("%c\n", '\u4f17')        //众
	fmt.Printf("Hello\nWorld\n!\n")     //字符串字面量
}

//scope 作用域
func scope() {
	//fmt.Printf("name=%s,Sex=%s,PI=%f\n", name, Sex, util.PI)
	//同package下的全局变量可以直接访问
	//不同package下的全局变量前面加上package名称也可以访问，但是变量名必须以大写字母开始

	var a = 9       //函数内定义的局部变量只能在本函数内访问
	var name = "43" // 内部声明的变量可以跟外部声明的变量有冲突，以内部的为准
	fmt.Printf("a=%d, name=%s\n", a, name)
	{ // {}圈定了一个内部作用域，内部域可以访问外部域的变量，但反过来不行
		fmt.Printf("a=%d\n", a) //访问外部域的
		a := 7                  //声明并定义了一个内部域的a
		fmt.Printf("a=%d\n", a)
		//a := 9 //在一个作用域内不能重复声明同一个变量
		a = 9
		fmt.Printf("a=%d\n", a)
		b := 9
		fmt.Printf("a=%d\n", b)
	}

	{
		//fmt.Printf("b=%d\n",b) //不能跨域访问b
		b := 9 // 在本作用域内，声明并定义了一个变量b
		fmt.Printf("b=%d\n", b)
	}
	//fmt.Printf("b=%d\n",b)  // 不能访问内部域的变量
}

func pointer() {
	var a int
	p := &a //p 用%p和%d输出得结果是一样的，只是形式不一样(十六进制和十进制)
	fmt.Printf("%p %p\n", p, &a)
	fmt.Printf("%d 0x%x\n", p, p)
}

func main() {
	constant()
	literal()
	scope()
	pointer()
}
