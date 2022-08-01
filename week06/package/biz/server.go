package main

import (
	"fmt"
	"go_8_mage/week06/package/common"
)

func init(){ //先执行这个init()
	fmt.Println("enter package/biz/server")
}

//在一个目录，甚至一个go文件里，init()可以重复定义
func main(){ // 再执行这个init()
	teacher := new(common.Teacher)
	teacher.Examine()
}


//enter package/user
//enter package/common/math/basic
//enter package/common/student
//enter package/common/teacher
//enter package/biz/server
//sum is 21