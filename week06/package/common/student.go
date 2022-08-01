package common

import (
	"fmt" //在$GOROOT/src目录下有fmt包
	wsx "go_8_mage/week06/package" //从当前go文件所在的目录逐级向上查找go.mod文件(假设go.mod位于)
	"go_8_mage/week06/package/common/math"
	"gonum.org/v1/gonum/stat" //第三方依赖包 在$GOPATH/pkg/mod下
)

func init(){
	fmt.Println("enter package/common/student")
}

type Student struct {
	Score float32
	wsx.User
}

func (stu *Student) varc(arr []float64)float64{
	math.Sum([]int{1,2,3})
	v := stat.Variance(arr,nil)
	fmt.Printf("Variance is %.4f\n",v)
	return v

}
