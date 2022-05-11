package main

import (
	"fmt"
	"time"
)

//a,b是形参，形参是函数内部的局部变量，实参的值会拷贝给形参
func arg1(a int, b int) { //注意大括号{不能另起一行
	a = a + b // 在函数内部修改形参的值，实参的值不受影响
	return    // 函数返回，return后面的语句不会再执行
	fmt.Println("我不会被输出")
}

func arg2(a, b int) { //参数类型相同时可以只写一次
	a = a + b
	//	不写return时，默认执行完最后一行代码函数返回
}

func arg3(a, b *int) { // 如果想通过函数修改实参，就需要指针类型
	*a = *a + *b
	*b = 888
}

func no_arg() { //函数可以没有参数，也没有返回值
	fmt.Println("欢迎开启Golang之旅")
}

func return1(a, b int) int { //函数需要返回一个int型数据
	a = a + b
	c := a //声明并初始化一个变量c
	return c
}

func return2(a, b int) (c int) {
	a = a + b
	c = a  // 直接使用c
	return // 由于函数要求有返回值，即使给c赋过值了，也需要显式写return
}

func return3() (int, int) { //可以没有形参，可以返回多个参数
	now := time.Now()
	//fmt.Println(now.Hour(),now.Minute())
	return now.Hour(), now.Minute()
}

//不定长参数
func variable_ength_arg(a int, other ...int) int {
	sum := a
	//	不定长参数实际上是slice类型
	for _, ele := range other {
		sum += ele
	}
	if len(other) > 0 {
		fmt.Printf("first ele %d len %d cap %d\n", other[0], len(other), cap(other))
	} else {
		fmt.Printf("len %d cap %d\n", len(other), cap(other))
	}
	return sum
}

func main() {
	//arg1(1, 2)

	var x, y int = 3, 6
	arg1(x, y)
	fmt.Printf("x=%d,y=%d\n", x, y)
	arg2(x, y)
	fmt.Printf("x=%d,y=%d\n", x, y)
	//& 是取地址符号，即取得某个变量的地址，如 &a
	//* 是指针运算符，可以表示一个变量是指针类型，也可以表示一个指针变量所指向的存储单元，也就是这个地址所存储的值。
	arg3(&x, &y)
	fmt.Printf("x=%d,y=%d\n", x, y)

	x, y = 3, 6
	fmt.Printf("return1 %d\n", return1(x, y))
	fmt.Printf("return2 %d\n", return2(x, y))
	hour, _ := return3() // 可以用_ 忽略返回值
	fmt.Printf("return3 %d\n", hour)
	//fmt.Println(return3)
	fmt.Println()

	//	不定长参数可以对应0个或多个实参
	fmt.Println(variable_ength_arg(1))
	fmt.Println(variable_ength_arg(1, 2))
	fmt.Println(variable_ength_arg(1, 2, 3))
	fmt.Println(variable_ength_arg(1, 2, 3, 4))
	arr := []int{4, 5, 6}
	fmt.Println(variable_ength_arg(1, arr...)) //slice（注意不能是数组）后面加... 可作为不定长参数对应的实参
	fmt.Println()

	//	append函数接收的就是不定长参数
	arr = append(arr, 1, 2, 3)
	arr = append(arr, 7)
	//arr = append(arr)
	fmt.Printf("new arr %v\n", arr)
	slice := append([]byte("hello"), "world"...) // ...自动把“world”转成byte切片，等价于[]byte("world")...
	fmt.Printf("slice %v\n", slice)
	slice2 := append([]rune("hello"), []rune("world")...) //需要显示把“world”转成rune切片
	fmt.Printf("slice2 %v\n", slice2)
	fmt.Println()

}
