---

---

# Golang笔记\_第2周\_数据类型



2022-4-21 15:29 开始记录go笔记

> 种一棵树最好的时间是十年前，其次，是现在。

当前是在slice 课程，就先从slice开始记录，前面的后期再补充



## 目录

[TOC]



#### 基础数据类型

![image-20220423102826620](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC2%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B.assets/image-20220423102826620.png)









#### 复合数据类型

![image-20220423094728436](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC2%E5%91%A8_%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B.assets/image-20220423094728436.png)

#### 自定义类型、类型别名

```
类型别名
type byte = uint8
type rune = int32
type semaphore = uint8

自定义类型
type user struct{name string；age int} // 用分号把多行代码隔开
type signal uint8
type ms map[string]string
type add func(a,b int)int
```





#### 字符串



##### 字符串赋值

```GO
s1 := "My name is jasmine☺"
s2 := "He say:\"I'm fine.\"\\Thank\tyou.\\" //包含转义字符
s3 := `here is first line.

there is third line.
`//反引号里的转义字符无效。反引号里的内容原封不动的输出，包括空百符和换行符
```



##### 字符串常用操作

| 方法                                 | 介绍           |      |
| ------------------------------------ | -------------- | ---- |
| len(str)                             | 求长度         |      |
| strings.Split                        | 分割           |      |
| strings.Contains                     | 判断是否包含   |      |
| strings.HasPrefix，strings.HasSuffix | 前缀、后缀判断 |      |
| strings.Index()，strings.LastIndex() | 子串出现的位置 |      |





##### 字符串拼接

> - 加号连接
> - func fmt.Sprintf{format string, a ...interface{}} string
> - func strings.Join(elems []string,sep string)string
> - 当有大量的string 需要拼接时，用strings.Builder效率最高





#### byte和rune

> - string中每个元素叫”字符“，字符有两种
>   - byte：1个字节，代表ASCII码的一个字符
>   - rune：4个字节，代表一个UTF-8字符，一个汉字可用一个rune表示
> - string底层是byte数组，string的长促就是改byte数组的长度，UTF-8编码下一个汉字占3个byte，即一个汉字占3个长度
> - string可以转换为[]byte或[]rune类型
> - string是常量，不能修改其中的字符















#### Printf%字符的定义

| 字符 | 含义       | 其他       |
| ---- | ---------- | ---------- |
| %d   | 整型       |            |
| %s   | string     |            |
| %t   | bool       |            |
| %f   | float32/64 |            |
| %p   | 指针       |            |
| %v   | 引用类型   | %v %+v %#v |
|      |            |            |







#### 数组（array）





#### 切片(slice)

![avatar](Golang笔记_第2周_数据类型.assets/image-20220421155901609.png)



##### 切片的定义

###### 概念待补充

```

```





###### 切片源码定义

```go
type slice struct{
    array unsafe.Pointer
    len int
    cap int
}
```



##### 初始化切片

```GO
var s []int //切片声明
s = []int{}// 初始化，len=cap=0
s = make([]int,3)// 初始化，len=cap=3
s = make([]int,3,5)// 初始化，len=3，cap=5
s = []int{1,2,3,4,5}// 初始化，len=cap=5
s2d := [][]int{
    {1},{2,3},//二维数组各行的列数相等，当二维切片各行的len可以不等
}


// 二维数组的用途   地图


```



###### 二维数组的各行的len可以不相等代码演示

![image-20220421155106593](Golang笔记_第2周_数据类型.assets/image-20220421155106593.png)



###### 一维切片初始化的代码演示

![image-20220421155643527](Golang笔记_第2周_数据类型.assets/image-20220421155643527.png)





##### append操作



###### 概念解释

> - 切片相对于数组最大的特点就是可以追加元素，可以自动扩容
>
> - 追加的元素放到预留的内存空间里，同时len加1
>
> - 如果预留空间已用完，则会重新申请一块更大的内存空间，capacity变成之前的2倍（cap<1024)或1.25倍(cap>1024).
>
>   把原内存空间的数据拷贝过来，在新内存空间上执行append操作



######  append操作代码演示

![image-20220421160552135](Golang笔记_第2周_数据类型.assets/image-20220421160552135.png)



###### append要拿到返回值

不然切片的append的操作不会改变切片的值

python中是s.append(num)就添加了

但是在go语言中必须要要用原切片 s= append（s，100）接受这个值，才能append成功



![image-20220421160735695](Golang笔记_第2周_数据类型.assets/image-20220421160735695.png)





###### 验证slice的cap是否是呈两倍递增

![image-20220421161016399](Golang笔记_第2周_数据类型.assets/image-20220421161016399.png)





![image-20220421161349896](Golang笔记_第2周_数据类型.assets/image-20220421161349896.png)





##### 烧脑问题：子切片和母切片的问题

截取子切片

![image-20220421162512735](Golang笔记_第2周_数据类型.assets/image-20220421162512735.png)



![image-20220421162746641](Golang笔记_第2周_数据类型.assets/image-20220421162746641.png)

![image-20220421162922309](Golang笔记_第2周_数据类型.assets/image-20220421162922309.png)





###### 字母切片的代码演示

![image-20220421163720696](Golang笔记_第2周_数据类型.assets/image-20220421163720696.png)





##### 切片传参

名词解释

> - go语言函数传参，传的都是值，即传切片会把切片的{arrayPoint,len,cap}这三个字段拷贝一份传进来
> - 由于传的是底层的数组的指针，所以可以直接修改底层数组里的元素





![image-20220421163922872](Golang笔记_第2周_数据类型.assets/image-20220421163922872.png)



###### 切片传参代码演示

![image-20220421164304737](Golang笔记_第2周_数据类型.assets/image-20220421164304737.png)





#### map

##### 底层概念图示

![image-20220421164422259](Golang笔记_第2周_数据类型.assets/image-20220421164422259.png)













##### map的初始化

![image-20220421164922572](Golang笔记_第2周_数据类型.assets/image-20220421164922572.png)





```
#概念文字 
# TODO
```





##### 添加和删除key

![image-20220421165125131](Golang笔记_第2周_数据类型.assets/image-20220421165125131.png)



```

```



##### 根据key找value

![image-20220421165811839](Golang笔记_第2周_数据类型.assets/image-20220421165811839.png)



##### 遍历map

map的值是无序的，但是值都是一个环，只是开始的值有变化

```go
for key,value := range m{
	fmt.Printf("%s=%d\n",key,value)
}
```



##### 代码演示

```go
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
```





#### 管道(chanel)也叫通道



##### 底层原理图

![image-20220421173249851](Golang笔记_第2周_数据类型.assets/image-20220421173249851.png)



##### 声明和初始化

```go
var ch chan int // 声明
ch = make(chan int,8) //初始化，环形队列里可容纳8个int
```



##### 管道的接收和取操作

- ch <- 1 //往管道里写入(send)数据
- ch <- 2
- ch <- 3
- v := <-ch //从管道里取走(recv)数据
- v = <-ch



##### 只读只写管道定义

```
read_only := make (<-chan int)   //定义只读的channel
write_only := make (chan<- int)   //定义只写的channel
```

&#8195;&#8195;定义只读和只写的channel意义不大，一般用于在参数传递中。  

```
//只能向channel里写数据 
func send(c chan<- int) { 
    c <- 1 
} 

//只能取channel中的数据 
func recv(c <-chan int) {
	_ = <-c
}

//返回一个只读channel
func (c *Context) Done() <-chan struct{} {
    return nil
}
```





##### 遍历管道

```go
close(ch) // 遍历前必须先关闭管道，禁止再写入元素



// 遍历管道里剩下的元素

for ele := range ch{
	fmt.Println(ele)
}
```



##### 代码演示

```go
package main

import "fmt"

func main(){
	var ch chan int
	fmt.Printf("ch is nil %t\n",ch == nil)
	fmt.Printf("len of ch is %d\n", len(ch))

	ch = make(chan int,10)
	fmt.Printf("len of ch is %d\n",len(ch))
	fmt.Printf("cap of ch is %d\n",cap(ch))
	for i := 0;i<10;i++{
		ch <- 3
	}
	fmt.Printf("len of ch is %d\n",len(ch))
	<-ch
	<-ch
	ch <- 3
	fmt.Printf("len of ch is %d\n",len(ch))
	fmt.Printf("cap of ch is %d\n",cap(ch))

	//两种循环方式是等价的，遍历管道
	close(ch)
	L:=len(ch)
	for i := 0;i<L;i++{
		ele := <-ch
		fmt.Println(ele)
	}
	fmt.Println("----------------------")

	//for ele := range ch{
	//	fmt.Println(ele)
	//}
	fmt.Printf("len of ch is %d\n", len(ch))
}

```





#### 引用类型总结

> - slice（数组），map（hastable）和channel（环形队列）是go语言里的3种引用类型，都可以通过**make**函数来进行初始化（申请内存分配）
> - 因为它们都包含一个指向底层数据结构的指针，所以称之为"引用"类型
> - 引用类型未初始化时都是nil，可以对他们执行len()函数，返回0
>
> 









#### 作业

![image-20220421181054551](Golang笔记_第2周_数据类型.assets/image-20220421181054551.png)