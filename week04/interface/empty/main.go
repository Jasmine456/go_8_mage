package main

import "fmt"

var (
	i interface{}//空接口类型用interface{}表示，注意有{}
	a int
)

//该函数可以接受任意类型的数据类型
func foo(x interface{}){
	fmt.Printf("arg type is %T,value is %v\n",x,x)
}

func sum(slice []interface{})float64  {
	rect :=0.0
	for _,ele := range slice{
	//	在Type Switch 语句中的case子句中不能使用fallthrough
		switch v:=ele.(type) {//隐式的在每个case中声明了一个变量v， .(type) 只能用在switch后面
		case int://v已被转为int类型
			fmt.Printf("ele if int,value is %d\n",v)
			rect += float64(v)
		case float32:// v已被转为float32类型
			fmt.Printf("ele if float32,value is %f\n",v)
			rect +=float64(v)
		case float64:
			fmt.Printf("ele is float64,value is %f\n",v)
			rect+=v
		case int8,int32,byte:// 如果case后面跟多种type，则v还是interface{}类型，go语言中byte是uint8的别名
			fmt.Printf("ele is %T,value is %d\n",v,v)
		//	rect + =float64(v)//由于类型有多个，不能使用float64()强制类型转换
		case string:
			fmt.Printf("ele is string,value is %s\n",v)
		}
	}
	return rect
}





func main(){
//	空接口没有定义任何方法，因此任意类型都实现了空接口
	a = 6
	i = a

	foo(i)
	foo(a)
	fmt.Println()

//	map的key和value都可以是interface{}类型
	mmap := make(map[interface{}]interface{},10)
	mmap["a"] = 1
	mmap["b"] = "A"
	mmap["c"] = 0.2
	mmap[9] =18
	for k,v := range  mmap{
		fmt.Printf("key type %T %v,value %T %v\n",k,k,v,v)
	}
	fmt.Println()

//	类型断言
	if v,ok := i.(int);ok{ //如果断言成功，则ok为true，v是具体的类型
		fmt.Printf("i是int类型，其值为%d\n",v)
	}else{
		fmt.Println("i不是int类型")
	}
	if v,ok := i.(float32); ok{
		fmt.Printf("i是float类型，其值为%f\n",v)
	}else {
		fmt.Println("i不是float类型")
	}
	// 当要判断的类型比较多时，就需要写很多if-else，更好的方法是使用switch i.(type)
	fmt.Println()

	slice := make([]interface{},0,10)
	slice = append(slice,1)
	slice = append(slice,"A")
	slice = append(slice,0.2)
	slice = append(slice,byte(100))
	fmt.Printf("sum of slice is %f\n",sum(slice))
	fmt.Println()


}
