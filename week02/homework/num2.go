package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
2. 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。
输入的切片可能很短，也可能很长。
 */


func arr2string(arr []int) string{
	var str string
	sb := strings.Builder{}
	space := " "
	for i:=0;i< len(arr);i++{
		str = strconv.Itoa(arr[i])
		sb.WriteString(str)
		if i < len(arr){
			sb.WriteString(space)
		}
	}
	output := sb.String()
	return output
}

func main(){
	s := []int{4,5,6}
	fmt.Println(arr2string(s))
}