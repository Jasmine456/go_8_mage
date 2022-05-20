# Golang笔记_第5周_面向对象和反射

[toc]

## 面向对象

### 面向对象的概念

---

洗衣服过程剖析：

1. 给洗衣机里加脏衣服和洗衣粉
2. 启动洗衣机
3. 洗衣机自动注水，然后滚动
4. 脏衣服从黑颜色变成白颜色
5. 洗衣机自动停止



用面向过程的思想实现代码。

```GO
//准备洗衣服
//输入参数：
//powder 洗衣机里放多少洗衣粉
//closes 洗衣机里放多少衣服
//clean 衣服是否是干净的
//返回值：
//洗衣机是否开启
//准备洗多少衣服
func prepare(powder int, closes int, clean bool) (bool, int) {
	if powder <= 0 || closes <= 0 || clean == true {
		return false, 0
	}
	return true, closes
}

//开始洗衣服
//输入参数：
//washer_state 洗衣机是否开启
//closes 准备洗多少衣服
//返回值：
//衣服是否是干净的
//洗了多少衣服
//洗衣机是否开启
func wash(washer_state bool, closes int) (bool, int, bool) {
	if washer_state == false {
		return false, 0, false
	} else {
		fmt.Println("注水")
		fmt.Println("滚动")
		fmt.Println("关机")
		return true, closes, false
	}
}

//检查最终状态
//输入参数：
//clean 衣服是否是干净的
//closes 洗了多少衣服
//washer_state 洗衣机是否开启
func check(clean bool, closes int, washer_state bool) {
	if clean && closes > 0 {
		fmt.Printf("洗干净了%d件衣服\n", closes)
		if washer_state {
			fmt.Println("你忘关洗衣机了")
		}
	} else {
		fmt.Println("洗衣失败")
	}
}

//整个洗衣服的过程
func WashProcedure(powder, closes int) {
	washer_state := false
	clean := false

	washer_state, closes = prepare(powder, closes, clean)
	clean, closes, washer_state = wash(washer_state, closes)
	check(clean, closes, washer_state)
}
```







面向过程编程整个过程分为若干步，每一步对应一个函数，函数之间要传递大量的函数。

面向对象编程把大量参数封装到一个结构体里面，给结构体赋予方法，方法里面去修改结构体的成员变量。

go语言面向对象的好处： 打包参数，继承，面向接口编程



```GO
//洗衣机
type Washer struct {
	State  bool
	Powder int
}

//衣服
type Closes struct {
	Clean bool
}

func (washer *Washer) prepare(closes []*Closes) error {
	if washer.State == true || washer.Powder <= 0 || len(closes) <= 0 {
		return errors.New("请确保在关机的状态下加入适量衣物和洗衣粉")
	}
	return nil
}

func (washer *Washer) wash(closes []*Closes) error {
	if err := washer.prepare(closes); err != nil {
		return err
	}

	fmt.Println("开机")
	washer.State = true

	//检查是否有脏衣服
	clean := true
	for _, ele := range closes {
		if ele.Clean == false {
			clean = false
			break
		}
	}
	if clean {
		washer.State = false
		return errors.New("所有衣服都是干净的，不需要洗")
	}

	//开始洗衣服
	fmt.Println("注水")
	fmt.Println("滚动")
	fmt.Println("关机")
	washer.State = false
	for _, ele := range closes {
		ele.Clean = true
	}
	return nil
}

func (washer *Washer) check(err error, closes []*Closes) {
	if err != nil {
		fmt.Printf("洗衣失败:%v\n", err)
	} else {
		fmt.Printf("洗干净了%d件衣服\n", len(closes))
		if washer.State == true {
			fmt.Println("你忘关洗衣机了")
		}
	}
}
```



### 构造函数

定义User结构体

```GO
type User struct {
    Name string //""表示未知
    Age int //-1表示未知
    Sex byte //1男，2女，3未知
}
```

- u := User{}构造一个空的User，个字段都取相应数据类型的默认值
- up := new(User)构造一个空的User，并返回其指针。



自定义构造函数

```GO
func NewDefaultUser() *User {
    return &User{
        Name: "",
        Age: -1,
        Sex: 3,
    }
}
```



```GO
func NewUser(name string, age int, sex byte) *User {
    return &User{
        Name: name,
        Age: age,
        Sex: sex,
    }
}
```



单例模式，确保在并发的情况下，整个进程里只会创建struct的一个实例。

```GO
var (
    sUser *User
    uOnce sync.Once
)
func GetUserInstance() *User {
    uOnce.Do(func() { //确保即使在并发的情况下，下面的3行代码在整个go进程里只会被执行一次
    if sUser == nil {
        sUser = NewDefaultUser()
        }
    })
    return sUser
}

//调用GetUserInstance()得到的是同一个User实例
su1 := GetUserInstance()
su2 := GetUserInstance()
//修改su1会影响su2
```





### 继承和重写

通过嵌入匿名结构体，变相实现“继承”的功能，因为访问匿名成员时可以跳过成员名直接访问它的内部成员。

```GO
type Plane struct {
	color string
}
type Bird struct {
	Plane 
}
bird := Bird {}
bird.Plane.color
bird.color
```



重写

```GO
func (plane Plane) fly() int {
	return 500
}

//重写父类(Plane)的fly方法
func (bird Bird) fly() int {
	return bird.Plane.fly()+100 //调用父类的方法
}
```

正规来讲，Go语言不支持继承，他只是支持组合。

```GO
type Plane struct {}
type Car struct{}
//Bird组合了Plane和Car的功能
type Bird struct {
	Plane 
	Car
}
```





### 泛型

在有泛型之前，同样的功能需要不同的参数类型单独实现一个函数

```GO
func add4int(a, b int) int {
	return a + b
}
func add4float32(a, b float32) float32 {
	return a + b
}
func add4string(a, b string) string {
	return a + b
}
```



使用泛型

```GO
type Addable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | complex64 | complex128 | string
}
func add[T Addable](a,b T)T{
	return a+b
}
```

&#8195;&#8195;go1.18正式支持泛型，但是Go语言之父Rob Pike建议不在Go 1.18的标准库中使用泛型。  





## 反射

### 反射介绍

反射就是在运行期间(不是编译期间) 探知对象的类型信息和内存结构、更新变量、调用它们的方法。

#### 反射的使用场景：

- 函数的参数类型是interface{},需要在运行时对原始类型进行判断，针对不同的类型采取不同的处理方式，比如json.Marshal(v interface{})
- 在运行时根据某些条件动态决定调用哪些函数，比如根据配置文件执行相应的算子函数

Go标准库里的json序列化就使用了反射

```GO
type User struct {
    Name string
    Age int
    Sex byte `json:"gender"`
}
user := User{
    Name: "钱钟书",
    Age: 57,
    Sex: 1,
}
json.Marshal(user)  //返回 {"Name":"钱钟书","Age":57,"gender":1}
```



#### 反射的弊端：

> - 代码难以阅读，难以维护
> - 编译期间不能发现类型错误，覆盖测试难度很大，有些bug需要到线上运行很长时间才能发现，可能会造成严重的使用后果
> - 反射性能很差，通常比正常代码慢一到两个数量级，在对性能要求很高，或大量反复调用的代码块里建议不要食用反射。





### 反射的基础类型

![reflect](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC5%E5%91%A8_%E9%9D%A2%E5%90%91%E5%AF%B9%E8%B1%A1%E5%92%8C%E5%8F%8D%E5%B0%84.assets/reflect.png)



reflect.Type用于获取类型相关的信息

```GO
type Type interface {
	Method(int) Method  //第i个方法
	MethodByName(string) (Method, bool) //根据名称获取方法
	NumMethod() int  //方法的个数
	Name() string   //获取结构体名称
	PkgPath() string //包路径
	Size() uintptr  //占用内存的大小
	String() string  //获取字符串表述
	Kind() Kind  //数据类型
	Implements(u Type) bool  //判断是否实现了某接口
	AssignableTo(u Type) bool  //能否赋给另外一种类型
	ConvertibleTo(u Type) bool  //能否转换为另外一种类型
	Elem() Type  //解析指针
	Field(i int) StructField  //第i个成员
	FieldByIndex(index []int) StructField  //根据index路径获取嵌套成员
	FieldByName(name string) (StructField, bool)  //根据名称获取成员
	FieldByNameFunc(match func(string) bool) (StructField, bool)  //
	Len() int  //容器的长度
	NumIn() int  //输出参数的个数
	NumOut() int  //返回参数的个数
}
```



通过reflect.Value获取、修改原始数据类型里的值

```GO
type Value struct {
	// 代表的数据类型
	typ *rtype
	// 指向原始数据的指针
	ptr unsafe.Pointer
}
```





### 反射API

#### reflect.Type

通过TypeOf() 得到Type类型

```
typeI := reflect.TypeOf(1)       
typeS := reflect.TypeOf("hello") 
fmt.Println(typeI)               //int
fmt.Println(typeS)               //string

typeUser := reflect.TypeOf(&common.User{}) 
fmt.Println(typeUser)                     //*common.User
fmt.println（typeUser.Elem()) //common.User,Elem()对指针类型进行解析
fmt.Println(typeUser.Kind())                 //ptr
fmt.Println(typeUser.Elem().Kind())    //struct
```



#### 指针Type转为非指针Type

```GO
typeUser := reflect.TypeOf(&common.User{}) 
typeUser2 := reflect.TypeOf(common.User{})
assert.IsEqual(typeUser.Elem(), typeUser2)
```



#### 获取struct成员变量的信息

```GO
typeUser := reflect.TypeOf(common.User{}) //需要用struct的Type，不能用指针的Type
fieldNum := typeUser.NumField()           //成员变量的个数
for i := 0; i < fieldNum; i++ {
	field := typeUser.Field(i)
	fmt.Printf("%d %s offset %d anonymous %t type %s exported %t json tag %s\n", i,
		field.Name,            //变量名称
		field.Offset,          //相对于结构体首地址的内存偏移量，string类型会占据16个字节
		field.Anonymous,       //是否为匿名成员
		field.Type,            //数据类型，reflect.Type类型
		field.IsExported(),    //包外是否可见（即是否以大写字母开头）
		field.Tag.Get("json")) //获取成员变量后面``里面定义的tag
}
fmt.Println()

//可以通过FieldByName获取Field
if nameField, ok := typeUser.FieldByName("Name"); ok {
	fmt.Printf("Name is exported %t\n", nameField.IsExported())
}
//也可以根据FieldByIndex获取Field
thirdField := typeUser.FieldByIndex([]int{2}) //参数是个slice，因为有struct嵌套的情况
fmt.Printf("third field name %s\n", thirdField.Name)
```





#### 内存对齐

```Go
type A struct {
	sex    bool   //由于由于下一个成员age刚好占8个字节，所以内存对齐机制会让sex占用8B个字节(向8的整倍数对齐)
	age    int64  //offset  8B
	weight uint16 //offset  16B
} //整个结构体占24个字节
t := reflect.TypeOf(A{})
fmt.Println(t.Size()) //24B

fieldNum := t.NumField() //成员变量的个数，包括未导出成员
for i := 0; i < fieldNum; i++ {
	field := t.Field(i)
	fmt.Printf("%s offset %d\n",
		field.Name,   //变量名称
		field.Offset, //相对于结构体首地址的内存偏移量，string类型会占据16个字节
	)
}

type B struct {
	weight uint16 //uint16占2个字节
	sex    bool   //offset 2B
	age    int64  //offset  8B。虽然weight和sex加起来占3个字节，但由于age要占满8个字节，所以weight和sex加起来实际上占8个字节(向8的整倍数对齐)
} //整个结构体占16个字节
t = reflect.TypeOf(B{})
fmt.Println(t.Size()) //24B

fieldNum = t.NumField() //成员变量的个数，包括未导出成员
for i := 0; i < fieldNum; i++ {
	field := t.Field(i)
	fmt.Printf("%s offset %d\n",
		field.Name,   //变量名称
		field.Offset, //相对于结构体首地址的内存偏移量，string类型会占据16个字节
	)
}
```

#### 获取struct成员方法的信息

```Go
typeUser := reflect.TypeOf(common.User{})
methodNum := typeUser.NumMethod() //成员方法的个数。接收者为指针的方法【不】包含在内
for i := 0; i < methodNum; i++ {
	method := typeUser.Method(i)
	fmt.Printf("method name:%s ,type:%s, exported:%t\n", method.Name, method.Type, method.IsExported())
}
fmt.Println()

typeUser2 := reflect.TypeOf(&common.User{})
methodNum = typeUser2.NumMethod() //成员方法的个数。接收者为指针或值的方法【都】包含在内，也就是说值实现的方法指针也实现了（反之不成立）
for i := 0; i < methodNum; i++ {
	method := typeUser2.Method(i)
	fmt.Printf("method name:%s ,type:%s, exported:%t\n", method.Name, method.Type, method.IsExported())
}
```

#### 获取函数的信息

```Go
func Add(a, b int) int {
	return a + b
}

typeFunc := reflect.TypeOf(Add) //获取函数类型
fmt.Printf("is function type %t\n", typeFunc.Kind() == reflect.Func)
argInNum := typeFunc.NumIn()   //输入参数的个数
argOutNum := typeFunc.NumOut() //输出参数的个数
for i := 0; i < argInNum; i++ {
	argTyp := typeFunc.In(i)
	fmt.Printf("第%d个输入参数的类型%s\n", i, argTyp)
}
for i := 0; i < argOutNum; i++ {
	argTyp := typeFunc.Out(i)
	fmt.Printf("第%d个输出参数的类型%s\n", i, argTyp)
}
```

#### 判断类型是否实现了某接口

```Go
//通过reflect.TypeOf((*<interface>)(nil)).Elem()获得接口类型。因为People是个接口不能创建实例，所以把nil强制转为*common.People类型
typeOfPeople := reflect.TypeOf((*common.People)(nil)).Elem()
fmt.Printf("typeOfPeople kind is interface %t\n", typeOfPeople.Kind() == reflect.Interface)
t1 := reflect.TypeOf(common.User{})
t2 := reflect.TypeOf(&common.User{})
//如果值类型实现了接口，则指针类型也实现了接口；反之不成立
fmt.Printf("t1 implements People interface %t\n", t1.Implements(typeOfPeople))
```

#### reflect.Value

##### 如果获得Value

通过ValueOf()得到Value。

```Go
iValue := reflect.ValueOf(1)
sValue := reflect.ValueOf("hello")
userPtrValue := reflect.ValueOf(&common.User{
	Id:     7,
	Name:   "杰克逊",
	Weight: 65,
	Height: 1.68,
})
fmt.Println(iValue)       //1
fmt.Println(sValue)       //hello
fmt.Println(userPtrValue) //&{7 杰克逊  65 1.68}
```

##### Value转为Type

```Go
iType := iValue.Type()
sType := sValue.Type()
userType := userPtrValue.Type()
//在Type和相应Value上调用Kind()结果一样的
fmt.Println(iType.Kind() == reflect.Int, iValue.Kind() == reflect.Int, iType.Kind() == iValue.Kind())  
fmt.Println(sType.Kind() == reflect.String, sValue.Kind() == reflect.String, sType.Kind() == sValue.Kind()) 
fmt.Println(userType.Kind() == reflect.Ptr, userPtrValue.Kind() == reflect.Ptr, userType.Kind() == userPtrValue.Kind())
```

##### 指针Value和非指针Value互相转换

```Go
userValue := userPtrValue.Elem()                    //Elem() 指针Value转为非指针Value
fmt.Println(userValue.Kind(), userPtrValue.Kind())  //struct ptr
userPtrValue3 := userValue.Addr()                   //Addr() 非指针Value转为指针Value
fmt.Println(userValue.Kind(), userPtrValue3.Kind()) //struct ptr
```

##### 得到Value对应的原始数据

通过Interface()函数把Value转为interface{}，再从interface{}强制类型转换，转为原始数据类型。或者在Value上直接调用Int()、String()等一步到位。

```Go
fmt.Printf("origin value iValue is %d %d\n", iValue.Interface().(int), iValue.Int())
fmt.Printf("origin value sValue is %s %s\n", sValue.Interface().(string), sValue.String())
user := userValue.Interface().(common.User)
fmt.Printf("id=%d name=%s weight=%.2f height=%.2f\n", user.Id, user.Name, user.Weight, user.Height)
user2 := userPtrValue.Interface().(*common.User)
fmt.Printf("id=%d name=%s weight=%.2f height=%.2f\n", user2.Id, user2.Name, user2.Weight, user2.Height)
```

##### 空Value的判断

```Go
var i interface{} //接口没有指向具体的值
v := reflect.ValueOf(i)
fmt.Printf("v持有值 %t, type of v is Invalid %t\n", v.IsValid(), v.Kind() == reflect.Invalid)

var user *common.User = nil
v = reflect.ValueOf(user) //Value指向一个nil
if v.IsValid() {
	fmt.Printf("v持有的值是nil %t\n", v.IsNil()) //调用IsNil()前先确保IsValid()，否则会panic
}

var u common.User //只声明，里面的值都是0值
v = reflect.ValueOf(u)
if v.IsValid() {
	fmt.Printf("v持有的值是对应类型的0值 %t\n", v.IsZero()) //调用IsZero()前先确保IsValid()，否则会panic
}
```

##### 可寻址

```Go
v1 := reflect.ValueOf(2) //不可寻址
var x int
v2 := reflect.ValueOf(x)                           //不可寻址
v3 := reflect.ValueOf(&x)                          //不可寻址
v4 := v3.Elem()                                    //可寻址
fmt.Printf("v1 is addressable %t\n", v1.CanAddr()) //false
fmt.Printf("v2 is addressable %t\n", v2.CanAddr()) //false
fmt.Printf("v3 is addressable %t\n", v3.CanAddr()) //false
fmt.Printf("v4 is addressable %t\n", v4.CanAddr()) //true

slice := make([]int, 3, 5)
v5 := reflect.ValueOf(slice)                       //不可寻址
v6 := v5.Index(0)                                  //可寻址
fmt.Printf("v5 is addressable %t\n", v5.CanAddr()) //false
fmt.Printf("v6 is addressable %t\n", v6.CanAddr()) //true   sliceValue里的每一个元素是可寻址的

mp := make(map[int]bool, 5)
v7 := reflect.ValueOf(mp)                          //不可寻址
fmt.Printf("v7 is addressable %t\n", v7.CanAddr()) //false
```

##### 通过Value修改原始数据的值

```Go
var i int = 10
var s string = "hello"
user := common.User{
	Id:     7,
	Name:   "杰克逊",
	Weight: 65.5,
	Height: 1.68,
}

valueI := reflect.ValueOf(&i)
if valueI.CanAddr() { //false
	valueI.SetInt(8)
}
if valueI.Elem().CanAddr() { //必须是可寻址的，才能调用Set进行修改
	valueI.Elem().SetInt(8)
}

valueS := reflect.ValueOf(&s)
valueUser := reflect.ValueOf(&user)

valueS.Elem().SetString("golang") //不能在指针Value上调用Set系列函数
// valueS.Elem().SetInt(8)//会panic
valueUser.Elem().FieldByName("Weight").SetFloat(68.0) //FieldByName()通过Name返回类的成员变量。不能在指针Value上调用FieldByName
fmt.Printf("i=%d, s=%s, user.Weight=%.1f\n", i, s, user.Weight)
```

强调一下，要想修改原始数据的值，给ValueOf传的必须是指针，而指针Value不能调用Set和FieldByName方法，所以得先通过Elem()转为非指针Value。  
未导出成员的值不能通过反射进行修改。  

```Go
addrValue := valueUser.Elem().FieldByName("addr")
if addrValue.CanAddr() { //true
	fmt.Println("通过FieldByName获得的value是可寻址的")
}
if addrValue.CanSet() {
	addrValue.SetString("北京")
} else {
	fmt.Println("addr是未导出成员，不可Set") //以小写字母开头的成员相当于是私有成员
}
```

##### 通过Value修改Slice

```Go
users := make([]*common.User, 1, 5) //len=1，cap=5
users[0] = &common.User{
	Id:     7,
	Name:   "杰克逊",
	Weight: 65.5,
	Height: 1.68,
}

sliceValue := reflect.ValueOf(&users) //准备通过Value修改users，所以传users的地址
if sliceValue.Elem().Len() > 0 {      //取得slice的长度
	sliceValue.Elem().Index(0).Elem().FieldByName("Name").SetString("令狐一刀")
	fmt.Printf("1st user name change to %s\n", users[0].Name)
}
```

甚至可以修改slice的cap，新的cap必须位于原始的len到cap之间，即只能把cap改小。

```Go
sliceValue.Elem().SetCap(3)
```

通过把len改大，可以实现向slice中追加元素的功能。

```Go
sliceValue.Elem().SetLen(2)
//调用reflect.Value的Set()函数修改其底层指向的原始数据
sliceValue.Elem().Index(1).Set(reflect.ValueOf(&common.User{
	Id:     8,
	Name:   "李达",
	Weight: 80,
	Height: 180,
}))
fmt.Printf("2nd user name %s\n", users[1].Name)
```

##### 修改map

Value.SetMapIndex()函数：往map里添加一个key-value对。  
Value.MapIndex()函数： 根据Key取出对应的map。  

```Go
u1 := &common.User{
	Id:     7,
	Name:   "杰克逊",
	Weight: 65.5,
	Height: 1.68,
}
u2 := &common.User{
	Id:     8,
	Name:   "杰克逊",
	Weight: 65.5,
	Height: 1.68,
}
userMap := make(map[int]*common.User, 5)
userMap[u1.Id] = u1

mapValue := reflect.ValueOf(&userMap)                                                         //准备通过Value修改userMap，所以传userMap的地址
mapValue.Elem().SetMapIndex(reflect.ValueOf(u2.Id), reflect.ValueOf(u2))                      //SetMapIndex 往map里添加一个key-value对
mapValue.Elem().MapIndex(reflect.ValueOf(u1.Id)).Elem().FieldByName("Name").SetString("令狐一刀") //MapIndex 根据Key取出对应的map
for k, user := range userMap {
	fmt.Printf("key %d name %s\n", k, user.Name)
}
```

##### 调用函数

```Go
valueFunc := reflect.ValueOf(Add) //函数也是一种数据类型
typeFunc := reflect.TypeOf(Add)
argNum := typeFunc.NumIn()            //函数输入参数的个数
args := make([]reflect.Value, argNum) //准备函数的输入参数
for i := 0; i < argNum; i++ {
	if typeFunc.In(i).Kind() == reflect.Int {
		args[i] = reflect.ValueOf(3) //给每一个参数都赋3
	}
}
sumValue := valueFunc.Call(args) //返回[]reflect.Value，因为go语言的函数返回可能是一个列表
if typeFunc.Out(0).Kind() == reflect.Int {
	sum := sumValue[0].Interface().(int) //从Value转为原始数据类型
	fmt.Printf("sum=%d\n", sum)
}
```

##### 调用成员方法

```Go
common.User{
	Id:     7,
	Name:   "杰克逊",
	Weight: 65.5,
	Height: 1.68,
}
valueUser := reflect.ValueOf(&user)              //必须传指针，因为BMI()在定义的时候它是指针的方法
bmiMethod := valueUser.MethodByName("BMI")       //MethodByName()通过Name返回类的成员变量
resultValue := bmiMethod.Call([]reflect.Value{}) //无参数时传一个空的切片
result := resultValue[0].Interface().(float32)
fmt.Printf("bmi=%.2f\n", result)

//Think()在定义的时候用的不是指针，valueUser可以用指针也可以不用指针
thinkMethod := valueUser.MethodByName("Think")
thinkMethod.Call([]reflect.Value{})

valueUser2 := reflect.ValueOf(user)
thinkMethod = valueUser2.MethodByName("Think")
thinkMethod.Call([]reflect.Value{})
```

#### 创建对象

##### 创建struct

```Go
user :=t := reflect.TypeOf(common.User{})
value := reflect.New(t) //根据reflect.Type创建一个对象，得到该对象的指针，再根据指针提到reflect.Value
value.Elem().FieldByName("Id").SetInt(10)
user := value.Interface().(*common.User) //把反射类型转成go原始数据类型Call([]reflect.Value{})
```

##### 创建slice

```Go
var slice []common.User
sliceType := reflect.TypeOf(slice)
sliceValue := reflect.MakeSlice(sliceType, 1, 3)
sliceValue.Index(0).Set(reflect.ValueOf(common.User{
	Id:     8,
	Name:   "李达",
	Weight: 80,
	Height: 180,
}))
users := sliceValue.Interface().([]common.User)
fmt.Printf("1st user name %s\n", users[0].Name)
```

##### 创建map

```Go
var userMap map[int]*common.User
mapType := reflect.TypeOf(userMap)
// mapValue:=reflect.MakeMap(mapType)
mapValue := reflect.MakeMapWithSize(mapType, 10)

user := &common.User{
	Id:     7,
	Name:   "杰克逊",
	Weight: 65.5,
	Height: 1.68,
}
key := reflect.ValueOf(user.Id)
mapValue.SetMapIndex(key, reflect.ValueOf(user))                    //SetMapIndex 往map里添加一个key-value对
mapValue.MapIndex(key).Elem().FieldByName("Name").SetString("令狐一刀") //MapIndex 根据Key取出对应的map
userMap = mapValue.Interface().(map[int]*common.User)
fmt.Printf("user name %s %s\n", userMap[7].Name, user.Name)
```

reflect包里除了MakeSlice()和MakeMap()，还有MakeChan()和MakeFunc()。

### 自行实现json序列化

&#8195;&#8195;所谓序列化即把struct实例转为string。比如定义了User和Book两个struct。

```Go
type User struct {
    Name string
    Age int
    Sex byte `json:"gender"`
}
type Book struct {
    ISBN string `json:"isbn"`
    Author User `json:"author"`
    Keywords []string `json:"kws"`
}
```

&#8195;&#8195;Book的实例序列化后为  

```json
{
	"isbn": "4243547567",
	"author": {
		"Name": "钱钟书",
		"Age": 57,
		"gender": 1
	},
	"kws": ["爱情", "民国", "留学"]
}
```

&#8195;&#8195;序列化实现思路：

- 从内向外、从简单到复杂地考虑序列化问题。
  1. 如果要序列化一个int、float、string，很简单。
  2. 如果要序列化一个slice，则在第1步的基础上用[]括起来。
  3. 如果要序列化一个struct，FieldName直接打印出来，FieldValue的序列化可以参考第1、2步。
  4. 如果struct内部还嵌套了struct，则递归调用第3步。
- 通过反射解析struct，得到json key和struct FieldName的对应关系。
- 如果FieldValue是基本的值类型，则通过反射给FieldValue赋值很简单。
- 如果FieldValue是slice类型，则需要通过反射先创建一个slice，再给slice里的每个元素赋值。
- 如果FieldValue是是内嵌struct，则递归调用反序列化函数，给FieldValue赋值。
- 如果FieldValue是是内嵌struct指针，则需要创建内嵌struct对应的实例（申请内存空间），再递归调用反序列化函数，给FieldValue赋值。





























































