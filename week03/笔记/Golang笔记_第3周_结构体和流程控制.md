# Golang笔记\_第三周\_结构体和流程控制





## 目录

[TOC]

> - if语句
> - switch语句
> - for循环
> - break与 continue
> - goto语句与Label





### if语句



#### 常用示例

```
if 5> 9 {
	fmt.Println("5>9")
}
```



> - 如果逻辑表达式成立，就会执行{}里的内容
> - 逻辑表达式不需要加（）
> - “{”必须跟在逻辑表达式后面，不能另起一行



#### 复杂的逻辑表达式 if

```GO
if c,d,e := 5,9,2;c<d && (c>e||c>3){//初始化多个局部变量，复杂的逻辑表达式
    fmt.Println("fit") 
}
```



> - 逻辑表达式中可以含有变量或常量
> - if句子中允许包含1个（仅1个）分号，在分号前初始化一些局部变量(即只在if块内可见)



#### if规范

if-else代码演示

![image-20220422143407754](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC3%E5%91%A8_%E7%BB%93%E6%9E%84%E4%BD%93%E5%92%8C%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6.assets/image-20220422143407754.png)





if嵌套最好不要超过 3层，规避多层嵌套



#### 多层if更换方式代码示例

![image-20220422143717052](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC3%E5%91%A8_%E7%BB%93%E6%9E%84%E4%BD%93%E5%92%8C%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6.assets/image-20220422143717052.png)



#### if map的值不存在代码示例

![image-20220422150136474](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC3%E5%91%A8_%E7%BB%93%E6%9E%84%E4%BD%93%E5%92%8C%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6.assets/image-20220422150136474.png)



#### if代码实操---公交车道

![image-20220422150044379](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC3%E5%91%A8_%E7%BB%93%E6%9E%84%E4%BD%93%E5%92%8C%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6.assets/image-20220422150044379.png)









### switch语句

#### 语法示例

![image-20220422150232452](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC3%E5%91%A8_%E7%BB%93%E6%9E%84%E4%BD%93%E5%92%8C%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6.assets/image-20220422150232452.png)





#### 规范

> - switch-case-default可能模拟 if-else if-else,但只能实现相等判断
> - switch和case后面可以跟常量、变量或函数表达式，只要它们表示的数据类型相同就行
> - case后面可以跟多个值，只要有一个值满足就行







#### 空的switch

> - switch后带表达式时，switch-case只能模拟相等的情况；如果switch后不带表达式，case后就可以跟任意的条件表达式。



```GO
switch{
    case add(5)>10;
    	fmt.Println("right")
    default:
    	fmt.Println("wrong")
}
```







#### switch Type

![image-20220422150706624](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC3%E5%91%A8_%E7%BB%93%E6%9E%84%E4%BD%93%E5%92%8C%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6.assets/image-20220422150706624.png)





#### fallthrough 强制执行下一个case（或default）

```
func fall_throth(age int) {
	fmt.Printf("您的年龄是%d, 您可以：\n", age)
	switch {
	case age > 50:
		fmt.Println("出任国家首脑")
		fallthrough
	case age > 25:
		fmt.Println("生育子女")
		fallthrough
	case age > 22:
		fmt.Println("结婚")
		fallthrough
	case age > 18:
		fmt.Println("开车")
		fallthrough
	case age > 16:
		fmt.Println("参加工作")
	case age > 15:
		fmt.Println("上高中")
		fallthrough
	case age > 3:
		fmt.Println("上幼儿园")
	}
}
```



### for语句

#### 常用示例

```GO
arr := []int{1,2,3,4,5}
for i :=0;i< len(arr);i++{
    fmt.Printf("%d: %d\n",i,arr[i])
}
```



for初始化局部变量，条件表达式；后续操作

```go
for sum,i :=0,0;i< len(arr) && sum<100;sum,i=sum*1,i+1
```



#### 语法总结

> - 局部变量值尽在for块内可见
> - 初始化变量可以放在for上面
> - 后续操作可以放在for块内部
> - 只有条件判断时，前后的分号可以不要
> - for{} 是一个无限循环



#### for range

> - 遍历数组或切片
>   - for i,ele : =range arr
> - 遍历string
>   - for i,ele : =range "胜多负少" // ele是rune类型
> - 变量map，go不保证遍历的顺序
>   - for key，value :=range m
> - 遍历channel，遍历前一定要先close
>   - for ele := range ch
>   - for range 拿到的是数据的拷贝



#### for嵌套

矩阵乘法需要用到三层for循环嵌套

![image-20220505161559584](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC3%E5%91%A8_%E7%BB%93%E6%9E%84%E4%BD%93%E5%92%8C%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6.assets/image-20220505161559584.png)

#### 待完善代码







### break与continue

> - break与continue用于控制 for循环的代码流程，并且只针对最靠近自己的外层for循环
> - break：退出for循环，且本轮break下面的代码不再执行
> - continue：本轮continue下面的代码不再执行，进入for循环的下一轮





```go
//break和continue都是针对for循环的，不针对if或switch
//break和continue都是针对套在自己外面的最靠里的那层for循环，不针对更外层的for循环（除非使用Label）
func complex_break_continue() {
	const SIZE = 5
	arr := [SIZE][SIZE]int{}
	for i := 0; i < SIZE; i++ {
		fmt.Printf("开始检查第%d行\n", i)
		if i%2 == 1 {
			for j := 0; j < SIZE; j++ {
				fmt.Printf("开始检查第%d列\n", j)
				if arr[i][j]%2 == 0 {
					continue //针对第二层for循环
				}
				fmt.Printf("将要检查第%d列\n", j+1)
			}
			break //针对第一层for循环
		}
	}
}
```





### goto与Label

```
var i int = 4
MY_LABEL:
	i += 3
	fmt.Println(i)
	goto MY_LABEL //返回定义MY_LABEL的那一行，把代码再执行一遍（会进入一个无限循环）
```



```GO
if i%2 == 0 {
	goto L1 //Label指示的是某一行代码，并没有圈定一个代码块，所以goto L1也会执行L2后的代码
} else {
	goto L2//先使用Label
}
L1: 
	i += 3
L2: //后定义Label。Label定义后必须在代码的某个地方被使用
	i *= 3
```



 goto与Label结合可以实现break的功能，甚至比break更强大。



```
for i := 0; i < SIZE; i++ {
L2:
for j := 0; j < SIZE; j++ {
	goto L1
}
}
L1:
xxx
```



- break、continue与Label结合使用可以跳转到更外层的for循环。
- continue和break针对的Label必须写在for前面，而goto可以针对任意位置的Label。

```
func break_label() {
	const SIZE = 5
	arr := [SIZE][SIZE]int{}
L1:
	for i := 0; i < SIZE; i++ {
	L2:
		fmt.Printf("开始检查第%d行\n", i)

		if i%2 == 1 {
		L3:
			for j := 0; j < SIZE; j++ {
				fmt.Printf("开始检查第%d列\n", j)
				if arr[i][j]%3 == 0 {
					break L1 //直接退出最外层的fot循环
				} else if arr[i][j]%3 == 1 {
					goto L2 //continue和break针对的Label必须写在for前面，而goto可以针对任意位置的Label
				} else {
					break L3
				}
			}
		}
	}
}
```









---



### 结构体



#### 结构体创建、访问与修改

##### 定义结构体

```go
type user struct {
	id int
	score float32
	enrollment time.Time
	name,addr string // 多个字段类型相同时可以简写到一行里
}
```



##### 声明和初始化结构体

```go
var u user //声明，会用相应类型的默认值初始化struct里的每一个字段
u = user{} //用相应类型的默认值初始化struct里的每一个字段
u = user{id: 3, name: "zcy"} //赋值初始化
u = user{4, 100.0, time.Now(), "zcy", "beijing"} //赋值初始化，可以不写字段名，但需要跟结构体定义里的字段顺序一致
```

##### 访问和修改结构体

```go
u.enrollment = time.Now() //给结构体的成员变量赋值
fmt.Printf("id=%d, enrollment=%v, name=%s\n", u.id, u.enrollment, u.name)//访问结构体的成员变量
```



##### 成员方法

```go
//可以把user理解为hello函数的参数，即hello(u user, man string)
func (u user) hello(man string) {
    fmt.Println("hi " + man + ", my name is " + u.name)
}
//函数里不需要访问user的成员，可以传匿名，甚至_也不传
func (_ user) think(man string) {
    fmt.Println("hi " + man + ", do you know my name?")
}
```



##### 为自定义类型添加方法

```go
type UserMap map[int]user // 自定义类型
// 可以给自定义类型添加任意方法
func (um UserMap) GetUser(id int)User{
    return um[id]
}
```



##### 结构体的可见性

> - go语言关于可见的统一规则：大写字母开头跨package也可以访问，否则只能本package内部访问。
> - 结构体名称以大写开头时，apckage外部可见，在此前提下，结构体中以大写开头在成员变量或成员方法在package外部也可见



##### 匿名结构体

```GO
var stu struct { //声明stu是一个结构体，但这个结构体是匿名的
	Name string
	Addr string
}
stu.Name = "zcy"
stu.Addr = "bj"
```

匿名结构体通常用于只使用一次的情况



结构体中含有匿名成员

```GO
type Student struct{
    Id int
    string // 匿名字段
    float32 //直接使用数据类型作为字段名，所以匿名字段中不能出现重复的数据类型
}
var stu = Student{Id: 1, string: "zcy", float32: 79.5}
fmt.Printf("anonymous_member string member=%s float member=%f\n", stu.string, stu.float32)   //直接使用数据类型访问匿名成员
```





#### 结构体指针

---

##### 创建结构体指针

```go
var u User
user := &u //通过取址符&得到指针
user = &User{ //直接创建结构体指针
    Id: 3, Name: "zcy", addr: "beijing",
}
user = new(User) //通过new()函数实体化一个结构体，并返回其指针
```



##### 构造函数

```GO
//构造函数。返回指针是为了避免值拷贝
func NewUser(id int, name string) *User {
	return &User{
		Id: id,
		Name: name,
		addr: "China",
		Score: 59,
	}
}
```



##### 方法接收指针

```GO
//user传的是值，即传的是整个结构体的拷贝。在函数里修改结构体不会影响原来的结构体
func hello(u user, man string) {
    u.name = "杰克"
    fmt.Println("hi " + man + ", my name is " + u.name)
}
//传的是user指针，在函数里修改user的成员会影响原来的结构体
func hello2(u *user, man string) {
    u.name = "杰克"
    fmt.Println("hi " + man + ", my name is " + u.name)
}
//把user理解为hello()的参数，即hello(u user, man string)
func (u user) hello(man string) {
    u.name = "杰克"
    fmt.Println("hi " + man + ", my name is " + u.name)
}
//可以理解为hello2(u *user, man string)
func (u *user) hello2(man string) {
    u.name = "杰克"
    fmt.Println("hi " + man + ", my name is " + u.name)
}

```



#### 结构体嵌套

```GO
type user struct {
    name string
    sex byte
}
type paper struct {
    name string
    auther user //结构体嵌套
}
p := new(paper)
p.name = "论文标题"
p.auther.name = "作者姓名"
p.auther.sex = 0

type vedio struct {
    length int
    name string
    user//匿名字段,可用数据类型当字段名
}
```



##### 结构体嵌套时字段名冲突的问题

```go
v := new(vedio)
v.length = 13
v.name = "视频名称"
v.user.sex = 0 //通过字段名逐级访问
v.sex = 0 //对于匿名字段也可以跳过中间字段名，直接访问内部的字段名
v.user.name = "作者姓名" //由于内部、外部结构体都有name这个字段，名字冲突了，所以需要指定中间字段名
```





#### 深拷贝和浅拷贝

```GO
type User struct {
	Name string
}
type Vedio struct {
	Length int
	Author User
}
```

&#8195;&#8195;Go语言里的赋值都会发生值拷贝。  

![avatar](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC3%E5%91%A8_%E7%BB%93%E6%9E%84%E4%BD%93%E5%92%8C%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6.assets/deep_copy.png)  

```Go
type User struct {
	Name string
}
type Vedio struct {
	Length int
	Author *User
}
```

![avatar](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC3%E5%91%A8_%E7%BB%93%E6%9E%84%E4%BD%93%E5%92%8C%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6.assets/shallow_copy.png)  

- 深拷贝，拷贝的是值，比如Vedio.Length。
- 浅拷贝，拷贝的是指针，比如Vedio.Author。
- 深拷贝开辟了新的内存空间，修改操作不影响原先的内存。
- 浅拷贝指向的还是原来的内存空间，修改操作直接作用在原内存空间上。

&#8195;&#8195;传slice，对sclice的3个字段进行了拷贝，拷贝的是底层数组的指针，所以修改底层数组的元素会反应到原数组上。  

```Go
users := []User{{Name: "康熙"}}
func update_users(users []User) {
    users[0].Name = "光绪"
}
```

































