# Golang笔记\_第4周\_Go函数编程和接口编程





[toc]

## 函数

### 函数的基本形式

```go
// 函数定义 a b是形参
func argf(a int,b int){
    a = a+b
}
var x ,y int =3,6
argf(x,y)// 函数调用，x，y是实参
```



> - 形参是函数内部的局部变量，实参的值会拷贝给形参
> - 参数定义时的第一个大括号不能另起一行
> - 形参可以有0个或多个
> - 参数类型相同时可以只写一次，比如argf(a,b int)
> - 在函数内部修改形参的值，实参的值不受影响
> - 如果想通过函数修改实参，就需要指针类型



```GO
func argf(a,b *int){
    *a=*a+*b
    *b = 888
}
var x,y int = 3,6
argf(&x,&y)
```



slice 、map、channel都是引用类型，它们作为函数参数是其实跟普通struct没什么区别，都是对struct内部的各个字段做一次拷贝传到函数内部。

```GO
package main

import "fmt"

func slice_arg_1(arr []int) { //slice作为参数，实际上是把slice的arrayPointer、len、cap拷贝了一份传进来
	arr[0] = 1           //修改底层数据里的首元素
	arr = append(arr, 1) //arr的len和cap发生了变化，不会影响实参
}

func main() {
	arr := []int{8}
	slice_arg_1(arr)
	fmt.Println(arr[0])   //1
	fmt.Println(len(arr)) //1
}
```



关于函数返回值

- 可以返回0个或多个参数。
- 可以在func行直接声明要返回的变量。
- return后面的语句不会执行。
- 无返回参数时return可以不写。   

```Go
func returnf(a, b int) (c int) { //返回变量c已经声明好了
    a = a + b
    c = a //直接使用c
    return //由于函数要求有返回值，即使给c赋过值了，也需要显式写return
}
```

&#8195;&#8195;不定长参数实际上是slice类型。  

```Go
func variable_ength_arg(a int, other ...int) int { 
    sum := a
    for _, ele := range other {//不定长参数实际上是slice类型
        sum += ele
    }
    fmt.Printf("len %d cap %d\n", len(other), cap(other))
    return sum
}
variable_ength_arg(1)
variable_ength_arg(1,2,3,4)
```

&#8195;&#8195;append函数接收的就是不定长参数。  

```Go
arr = append(arr, 1, 2, 3)
arr = append(arr, 7)
arr = append(arr)
slice := append([]byte("hello "), "world"...) //...自动把"world"转成byte切片，等价于[]byte("world")...
slice2 := append([]rune("hello "), []rune("world")...) //需要显式把"world"转成rune切片
```

&#8195;&#8195;在很多场景下string都隐式的转换成了byte切片，而非rune切片，比如"a中"[1]是228而非"中"。
递归函数  

```Go
func Fibonacci(n int) int {
    if n == 0 || n == 1 {
        return n //凡是递归，一定要有终止条件，否则会进入无限循环
    }
    return Fibonacci(n-1) + Fibonacci(n-2) //递归调用函数自身
}
```

### 匿名函数

&#8195;&#8195;函数也是一种数据类型。

```Go
func function_arg1(f func(a, b int) int, b int) int { //f参数是一种函数类型
	a := 2 * b
	return f(a, b)
}

type foo func(a, b int) int //foo是一种函数类型
func function_arg2(f foo, b int) int { //参数类型看上去简洁多了
    a := 2 * b
    return f(a, b)
}

type User struct {
    Name string
    bye foo //bye的类型是foo，而foo代表一种函数类型
    hello func(name string) string //使用匿名函数来声明struct字段的类型
}

ch := make(chan func(string) string, 10)
ch <- func(name string) string {  //使用匿名函数
	return "hello " + name
}
```

### 闭包

&#8195;&#8195;闭包（Closure）是引用了自由变量的函数，自由变量将和函数一同存在，即使已经离开了创造它的环境。闭包复制的是原对象的指针。  

```Go
package main

import "fmt"

//闭包（Closure）是引用了自由变量的函数。自由变量将和函数一同存在，即使已经离开了创造它的环境。
func sub() func() {
	i := 10
	fmt.Printf("%p\n", &i)
	b := func() {
		fmt.Printf("i addr %p\n", &i) //闭包复制的是原对象的指针
		i--                           //b函数内部引用了变量i
		fmt.Println(i)
	}
	return b //返回了b函数，变量i和b函数将一起存在，即使已经离开函数sub()
}

// 外部引用函数参数局部变量
func add(base int) func(int) int {
	return func(i int) int {
		fmt.Printf("base addr %p\n", &base)
		base += i
		return base
	}
}

func main() {
	b := sub()
	b()
	b()
	fmt.Println()

	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2)) //11,13
	// 此时tmp1和tmp2不是一个实体了
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2)) //101,103
}
```

### 延迟调用defer

- defer用于注册一个延迟调用（在函数返回之前调用）。
- defer典型的应用场景是释放资源，比如关闭文件句柄，释放数据库连接等。
- 如果同一个函数里有多个defer，则后注册的先执行。
- defer后可以跟一个func，func内部如果发生panic，会把panic暂时搁置，当把其他defer执行完之后再来执行这个。
- defer后不是跟func，而直接跟一条执行语句，则相关变量在注册defer时被拷贝或计算。

```Go
func basic() {
    fmt.Println("A")
    defer fmt.Println(1) fmt.Println("B")
    //如果同一个函数里有多个defer，则后注册的先执行
    defer fmt.Println(2)
    fmt.Println("C")
}
```

```Go
func defer_exe_time() (i int) {
	i = 9
	defer func() { //defer后可以跟一个func
		fmt.Printf("first i=%d\n", i) //打印5，而非9。充分理解“defer在函数返回前执行”的含义，不是在“return语句前执行defer”
	}()
	defer func(i int) {
		fmt.Printf("second i=%d\n", i) //打印9
	}(i)
	defer fmt.Printf("third i=%d\n", i) //defer后不是跟func，而直接跟一条执行语句，则相关变量在注册defer时被拷贝或计算
	return 5
}
```

### 异常处理

&#8195;&#8195;go语言没有try catch，它提倡返回error。  

```Go
func divide(a, b int) (int, error) {
    if b == 0 {
        return -1, errors.New("divide by zero")
    }
    return a / b, nil
}
if res, err := divide(3, 0); err != nil {//函数调用方判断error是否为nil
    fmt.Println(err.Error())
}
```

&#8195;&#8195;Go语言定义了error这个接口，自定义的error要实现Error()方法。  

```Go
type PathError struct {    //自定义error
    path string
    op string
    createTime string
    message string
}
func (err PathError) Error() string {    //error接口要求实现Error() string方法
	return err.createTime + ": " + err.op + " " + err.path + " " + err.message
}
```

何时会发生panic:  

- 运行时错误会导致panic，比如数组越界、除0。
- 程序主动调用panic(error)。

panic会执行什么：  

1. 逆序执行当前goroutine的defer链（recover从这里介入）。
2. 打印错误信息和调用堆栈。
3. 调用exit(2)结束整个进程。  

&#8195;&#8195;recover会使程序从panic中恢复，并返回panic value。recover所在的函数后续的代码不会执行，但函数可以正常返回。在未发生panic时调用recover，会返回nil。recover()必须在defer中才能生效。  

```Go
func soo() {
	fmt.Println("enter soo")

	defer func() { //去掉这个defer试试，看看panic的流程。把这个defer放到soo函数末尾试试。把这个defer移到main()里试试。
		//recover必须在defer中才能生效
		if panic_value := recover(); panic_value != nil {
			fmt.Printf("soo函数中发生了panic:%v\n", panic_value)
		}
	}()
	fmt.Println("regist recover")

	defer fmt.Println("hello")
	defer func() {
		n := 0
		_ = 3 / n //除0异常，发生panic，下一行的defer没有注册成功
		defer fmt.Println("how are you")
	}()
}
```

```Go
package main

import (
	"fmt"
)

func B() {
	// defer func() { //方式一，recover()在B()函数里，则在B()函数中panic后面的代码不会执行。不影响BBBBBBB的打印
	// 	if panicValue := recover(); panicValue != nil {
	// 		fmt.Printf("panic info %v\n", panicValue)
	// 	}
	// }()
	panic(5)
}

func main() {
	defer func() { //方式二，recover()在main()函数里，则在main()函数中panic后面的代码不会执行。BBBBBBB不会打印出来
		if panicValue := recover(); panicValue != nil {
			fmt.Printf("panic info %v\n", panicValue)
		}
	}()

	B()
	fmt.Println("BBBBBBB")
}
```





## 面向接口编程

### 接口的基本概念



#### 定义

&#8195;&#8195;接口是一组行为规范的集合。  

```Go
type Transporter interface { //定义接口。通常接口名以er结尾
    //接口里面只定义方法，不定义变量
    move(src string, dest string) (int, error) //方法名 (参数列表) 返回值列表
    whistle(int) int //参数列表和返回值列表里的变量名可以省略
}
```

&#8195;&#8195;只要结构体拥有接口里声明的所有方法，就称该结构体“实现了接口”。一个struct可以同时实现多个接口。  

```Go
type Car struct { //定义结构体时无需要显式声明它要实现什么接口
    price int
}

func (car Car) move(src string, dest string) (int, error) {
    return car.price, nil
}
func (car Car) whistle(n int) int {
    return n
}
```



#### 底层结构

&#8195;&#8195;接口值有两部分组成, 一个指向该接口的具体类型的指针和另外一个指向该具体类型真实数据的指针。  

```Go
car := Car{"宝马", 100}
var transporter Transporter
transporter = car
```

![image-20220511171831672](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC4%E5%91%A8_Go%E5%87%BD%E6%95%B0%E5%92%8C%E6%8E%A5%E5%8F%A3%E7%BC%96%E7%A8%8B.assets/image-20220511171831672.png)



#### 接口的使用  

```Go
func transport(src, dest string, transporter Transporter) error {
	 _,err := transporter.move(src, dest)
	return err
}
var car Car		//Car实现了Transporter接口
var ship Shiper	// Shiper实现了Transporter接口
transport("北京", "天津", car)
transport("北京", "天津", ship)
```



#### 接口的赋值  

```Go
func (car Car) whistle(n int) int {…}//方法接收者是值
func (ship *Shiper) whistle(n int) int {…} //方法接收者用指针，则实现接口的是指针类型
car := Car{}
ship := Shiper{}
var transporter Transporter
transporter = car 
transporter = &car     //值实现的方法，指针同样也实现了
transporter = &ship
```

### 接口嵌入

```Go
type Transporter interface {
	whistle(int) int
}
type Steamer interface {
    Transporter //接口嵌入。相当于Transporter接口定义的行为集合是Steamer的子集
    displacement() int
}
```



### 空接口

&#8195;&#8195;空接口类型用interface{}表示，注意有{}。

```Go
var i interface{} 
```

&#8195;&#8195;空接口没有定义任何方法，因此任意类型都实现了空接口。

```Go
var a int = 5
i = a
```

```Go
func square(x interface{}){} //该函数可以接收任意数据类型
```

&#8195;&#8195;slice的元素、map的key和value都可以是空接口类型。map中的key可以是任意能够用==操作符比较的类型，不能是函数、map、切片，以及包含上述3中类型成员变量的的struct。map的value可以是任意类型。  



### 类型断言

```Go
if v, ok := i.(int); ok {//若断言成功，则ok为true，v是具体的类型
	fmt.Printf("i是int类型，其值为%d\n", v)
} else {
	fmt.Println("i不是int类型")
}
```

&#8195;&#8195;当要判断的类型比较多时，就需要写很多if-else，更好的方法是使用switch i.(type)。  

```Go
switch v := i.(type) {    //隐式地在每个case中声明了一个变量v
case int:  //v已被转为int类型
	fmt.Printf("ele is int, value is %d\n", v)
	//在 Type Switch 语句的 case 子句中不能使用fallthrough
case float64:      //v已被转为float64类型
	fmt.Printf("ele is float64, value is %f\n", v)
case int8, int32, byte: //如果case后面跟多种type，则v还是interface{}类型
	fmt.Printf("ele is %T, value is %d\n", v, v)
}
```



```GO
switch v := ele.(type) { //隐式地在每个case中声明了一个变量v。.(type)只能用在switch后面
		case int: //v已被转为int类型
			fmt.Printf("ele is int, value is %d\n", v)
			rect += float64(v)
		case float32: //v已被转为float32类型
			fmt.Printf("ele is float32, value is %f\n", v)
			rect += float64(v)
		case float64:
			fmt.Printf("ele is float64, value is %f\n", v)
			rect += v
		case int8, int32, byte: //如果case后面跟多种type，则v还是interface{}类型。go语言中byte是uint8的别名
			fmt.Printf("ele is %T, value is %d\n", v, v)
			// rect += float64(v)//由于类型有多个，不能使用float64()强制类型转换
		case string:
			fmt.Printf("ele is string, value is %s\n", v)
		}
```







### 面向接口编程

电商推荐流程  
![image-20220511172838733](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC4%E5%91%A8_Go%E5%87%BD%E6%95%B0%E5%92%8C%E6%8E%A5%E5%8F%A3%E7%BC%96%E7%A8%8B.assets/image-20220511172838733.png)



为每一个步骤定义一个接口。  

```Go
type Recaller interface {
    Recall(n int) []*common.Product //生成一批推荐候选集
}
type Sorter interface {
    Sort([]*common.Product) []*common.Product //传入一批商品，返回排序之后的商品
}
type Filter interface {
    Filter([]*common.Product) []*common.Product //传入一批商品，返回过滤之后的商品
}
type Recommender struct {
    Recallers []recall.Recaller
    Sorter sort.Sorter
    Filters []filter.Filter
}
```

使用纯接口编写推荐主流程。

```Go
func (rec *Recommender) Rec() []*common.Product {
	RecallMap := make(map[int]*common.Product, 100)
	//顺序执行多路召回
	for _, recaller := range rec.Recallers {
		products := recaller.Recall(10) //统一设置每路最多召回10个商品
		for _, product := range products {
			RecallMap[product.Id] = product //把多路召回的结果放到map里，按Id进行排重
		}
	}
	//把map转成slice
	RecallSlice := make([]*common.Product, 0, len(RecallMap))
	for _, product := range RecallMap {
		RecallSlice = append(RecallSlice, product)
	}
	SortedResult := rec.Sorter.Sort(RecallSlice) //对召回的结果进行排序
	//顺序执行多种过滤规则
	FilteredResult := SortedResult
	for _, filter := range rec.Filters {
		FilteredResult = filter.Filter(FilteredResult)
	}
	return FilteredResult
}
```

&#8195;&#8195;面向接口编程，在框架层面全是接口。具体的实现由不同的开发者去完成，每种实现单独放到一个go文件里，大家的代码互不干扰。通过配置选择采用哪种实现，也方便进行效果对比。  

