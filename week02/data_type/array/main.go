package main

import "fmt"

func array1d() {
	var arr1 [5]int = [5]int{}
	var arr2 = [5]int{}
	var arr3 = [5]int{3, 2}
	var arr4 = [5]int{2: 15, 4: 30}
	var arr5 = [...]int{3, 2, 6, 5, 4}
	var arr6 = [...]struct {
		name string
		age  int
	}{{"Tom", 18}, {"Jim", 20}}
	fmt.Printf("arr1=%#v\n", arr1)
	fmt.Printf("arr1=%#v\n", arr2)
	fmt.Printf("arr1=%#v\n", arr3)
	fmt.Printf("arr1=%#v\n", arr4)
	fmt.Printf("arr1=%#v\n", arr5)
	fmt.Printf("arr1=%#v\n", arr6)
	//	通过index访问数组里的元素
	fmt.Printf("arr5[0]=%d\n", arr5[0])
	fmt.Printf("arr5[len(arr5-1]=%d\n", arr5[len(arr5)-1])
	// 遍历数组里的元素
	for i, ele := range arr5 {
		fmt.Printf("index=%d,element=%d\n", i, ele)
	}
	// 数组的长度是确定不变，capacity和length相等
	fmt.Printf("len(arr1)=%d\n", len(arr1))
	fmt.Printf("cap(arr1)=%d\n", cap(arr1))

}

func array2d() {
	var arr1 = [5][3]int{{1}, {2, 3}}   //5行3列，只给前2行赋值，且前2行的所有列还没有赋满
	var arr2 = [...][3]int{{1}, {2, 3}} //第一维可以用...推测，第二位不能用...
	// 根据index访问数组里的元素
	fmt.Printf("arr[1][1]=%d\n", arr1[1][1])
	fmt.Printf("arr[4][2]=%d\n", arr1[4][2]) // 最后一个元素
	//	遍历二维数组
	for row, array := range arr2 {
		for col, ele := range array {
			fmt.Printf("arr2[%d][%d]=%d\n", row, col, ele)
		}
	}
	// 对于多维数组，其cap和len指第一维的长度
	fmt.Printf("len(arr1)=%d\n", len(arr1))
	fmt.Printf("cap(arr1)=%d\n", cap(arr1))
}

//参数必须是长度为5的int型数组(注意长度必须是5）
//
//go语言中函数传参都是按值传递，即传递数组实际上穿的是数字的拷贝
func update_array1(arr [5]int) {
	fmt.Printf("array in function,address is %p\n", &arr[0])
	arr[0] = 888
}

//传数组的指针
func update_array2(arr *[5]int) {
	fmt.Printf("array in function,address is %p\n", &((*arr)[0]))
	arr[0] = 888 // 因为传的是数组的指针，所以直接在原来的内存空间上进行修改
}

//for range 取得的是值拷贝
func for_range() {
	fmt.Println("--------------------")
	arr := [...]int{1, 2, 3}
	for i, ele := range arr { //ele是arr中元素的拷贝
		arr[i] += 8 //修改arr里的元素，不影响ele
		fmt.Printf("%d %d %d\n", i, arr[i], ele)
		ele += 1 //修改ele不影响arr
		fmt.Printf("%d %d %d\n", i, arr[i], ele)
	}


	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d %d\n", i, arr[i])
	}
}

func main() {
	//array1d()
	//array2d()

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("数组的地址：%p\n", &arr)
	fmt.Printf("第一个元素的地址：%p\n", &arr[0])
	fmt.Printf("第二个元素的地址：%p\n", &arr[1])
	update_array1(arr)
	fmt.Printf("arr=%v\n", arr)
	update_array2(&arr)
	fmt.Printf("arr=%v\n", arr)
	for_range()

}
