package common //同一个目录下只能存在一个包，即student.go和teacher.go的package名必须一致 。

import (
	"fmt"
	"go_8_mage/week06/package/common/math"
)

func init(){
	fmt.Println("enter package/common/teacher") //可以导入下级目录
}

type Teacher struct {
	WorkAge int
	students []Student
}

func (teacker *Teacher) Examine() int{
	arr := []int{1,2,3,4,5,6}
	sum := math.Sum(arr)
	fmt.Printf("sum is %d\n",sum)
	return sum
}

