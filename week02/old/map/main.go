package main

import "fmt"

func main(){
	var m map[string]int
	fmt.Println(len(m))
	m = make(map[string]int) //等价于m = make(map[string]int,0)
	fmt.Println(len(m))
	m = make(map[string]int,10) //cap=10
	fmt.Println(len(m))
	m=map[string]int{"A":3,"B":2,"C":1,"M":5,"N":7}
	fmt.Println(len(m))
	m["D"]=18

	delete(m,"B")

	key:="a"
	v,ok := m[key]
	if ok{
		fmt.Println(v)
	}else{
		fmt.Printf("%s这个key不存在\n",key)
	}


	for key,value := range m{
		// 不能修改map
		value +=2
		//可以修改!!!,值拷贝的问题
		//m[key] +=2
		//fmt.Printf("key:%s value:%d\n",key,value)
		fmt.Printf("value1:%d,value2:%d\n",value,m[key])
	}
	fmt.Println(m)

}
