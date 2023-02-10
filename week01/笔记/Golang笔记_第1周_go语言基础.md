# Golang笔记\_第一周\_go语言基础



[TOC]





## 初始go语言和安装步骤

### 开发环境搭建

1. 下载。到https://studygolang.com/dl上下载最新的Go稳定版本。
2. 安装。对于Windows和macOS用户，直接双击即可安装，留意一下安装路径。对于Linux用户，直接解压安装包即可，比如你打算把go安装到/usr/local目录下，则使用命令
   tar zxvf goxxx.tar.gz –C /usr/local。这样go标准库及相关的可执行文件就安装到了/usr/local/go目录下，在后续的步骤中会把/usr/local/go赋给GOROOT环境变量。  
3. 准确GOPATH。在任意目录下创建一个空目录，将来用于存放go语言第三方库文件。比如你打算使用/data/go_path这个目录，则在Linux下使用命令mkdir -p /data/go_path。在GOPATH目录建3个子目录：src、bin、pkg。  
4. 配置环境变量。把第2步和第3步生成的目录分别赋给GOROOT和GOPATH环境变量，对于Linux和Mac用户在~/.bashrc文件中追加以下几行

```shell
export GOROOT=/usr/local/go
export GOPATH=/data/go_path
export PATH=$PATH:$GOROOT/bin: :$GOPATH/bin
```

&#8195;&#8195;PATH环境变量下的可执行文件在任意目录下都可以直接访问。  
&#8195;&#8195;对于Windows用户，编辑用户环境变量，新增GOROOT和GOPATH，把GOROOT/bin和GOPATH/bin添加到Path里。如下图  



&#8195;集成开发环境推荐GoLand和VSCode，后者是免费的。VSCode需要额外安装支持Go语言的插件，如下图  

### 第一个Go程序

```Go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}
```

&#8195;&#8195;main()函数是Go程序的唯一入口，且main()函数必须位于package main中。fmt是Go标准库中的一个package，该package下有一个Println()函数用于输出字符串。Go语言会依次从以下3个目录里查找依赖包：

1. 当前工作目录
2. $GOPATH/pkg/mod
3. $GOROOT/src

### Go命令介绍

```Shell
(base) zcymac:~ zcy$ go help
Go is a tool for managing Go source code.

Usage:

	go <command> [arguments]

The commands are:

	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         update packages to use new APIs
	fmt         gofmt (reformat) package sources
	generate    generate Go files by processing source
	get         add dependencies to current module and install them
	install     compile and install packages and dependencies
	list        list packages or modules
	mod         module maintenance
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

	buildconstraint build constraints
	buildmode       build modes
	c               calling between Go and C
	cache           build and test caching
	environment     environment variables
	filetype        file types
	go.mod          the go.mod file
	gopath          GOPATH environment variable
	gopath-get      legacy GOPATH go get
	goproxy         module proxy protocol
	importpath      import path syntax
	modules         modules, module versions, and more
	module-get      module-aware go get
	module-auth     module authentication using go.sum
	packages        package lists and patterns
	private         configuration for downloading non-public code
	testflag        testing flags
	testfunc        testing functions
	vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.
```

go help: 查看帮助文档。  

```Shell
go help build
```

go build: 对源代码和依赖的文件进行打包，生成可执行文件。  

```Shell
go build -o my_first_go_exe entrance_class/demo.go
```

go install: 编译并安装包或依赖，安装到$GOPATH/bin下。  

```Shell
go install entrance_class/demo.go
```

go get: 把依赖库添加到当前module中，如果本机之前从未下载过则先下载。

```Shell
go get github.com/tinylib/msgp 
```

以上命令会在$GOPATH/pkg/mod目录下会生成github.com/tinylib/msgp目录。  

```Shell
go install github.com/tinylib/msgp@latest 
```

以上命令会在$GOPATH/bin下生成msgp可执行文件。  
go mod init module_name
初始化一个Go项目。  
go mod tidy通过扫描当前项目中的所有代码来添加未被记录的依赖至go.mod文件或从go.mod文件中删除不再被使用的依赖。  
go run: 编译并运行程序。  
go test: 执行测试代码。  
go tool: 执行go自带的工具。go tool pprof对cpu、内存和协程进行监控；go tool trace跟踪协程的执行过程。  
go vet: 检查代码中的静态错误。  
go fmt: 对代码文件进行格式化，如果用了IDE这个命令就不需要了。

```Shell
go fmt entrance_class/demo.go
```

go doc: 查看go标准库或第三方库的帮助文档。 

```Shell 
go doc fmt
go doc gonum.org/v1/gonum/stat
```

go version: 查看go版本号。  
go env: 查看go环境信息。  























## go基础语法

### 标识符和关键字

---

go变量、常量、自定义类型、包、函数的命名方式必须遵循以下规则：

1. 首字符可以是任意Unicode字符或下划线
2. 首字符之外的部分可以是Unicode字符、下划线或数字。
3. 名字的长度无限制

> 理论上名字里可以有汉字，甚至可以全是汉字，但实际中不要这么做。





Go语言的关键字

``` GO
break default func interface select case defer go map struct chan else goto package switch const if range type continue for import return fallthrough var
```



常量

``` go
true false iota nil
```



数据类型

``` go
int  int8  int16  int32  int64  uint  uint8  uint16  uint32  uint64  uintptr float32  float64  complex128  complex64 bool byte  rune  string  error
```



函数

``` GO
make  len  cap  new  append  copy  close  delete  complex  real  imag  panic  recover
```



### 操作符和表达式

***

算术运算符

| 运算符 | 描述 |
| ------ | ---- |
| +      | 相加 |
| -      | 相减 |
| *      | 相乘 |
| /      | 相除 |
| %      | 求余 |



``` GO
//arithmetic 算术运算
func arithmetic(){
	var a float32 = 8
	var b float32 = 3
	var c float32 = a+b
	var d float32 = a-b
	var e float32 = a*b
	var f float32 = a/b
	fmt.Printf("a=%.3f,b=%.3f,c=%.3f,d=%.3f,e=%.3f,f=%.3f\n",a,b,c,d,e,f)
}
```



关系运算符

| 运算符 | 描述                                                      |
| -----: | --------------------------------------------------------- |
|     == | 检查两个值是否相等，如果相等返回True否则返回False         |
|     != | 检查两个值是否不相等，不过不相等返回True否则返回False     |
|      > | 检查左边值是否大于右边值，如果是返回True否则返回False     |
|     >= | 检查左边值是否大于等于右边值，如果是返回True否则返回False |
|      < | 检查左边值是否小于右边值，如果是返回True否则返回False     |
|     <= | 检查左边值是否小于等于右边值，如果是返回True否则返回False |

```GO
//relational 关系运算符
func relational() {
	var a float32 = 8
	var b float32 = 3
	var c float32 = 8
	fmt.Printf("a==b吗 %t\n", a == b)
	fmt.Printf("a!=b吗 %t\n", a != b)
	fmt.Printf("a>b吗 %t\n", a > b)
	fmt.Printf("a>=b吗 %t\n", a >= b)
	fmt.Printf("a<c吗 %t\n", a < b)
	fmt.Printf("a<=c吗 %t\n", a <= c)
}
```



逻辑运算符  

| 运算符 | 描述                                                         |
| :----: | :----------------------------------------------------------- |
|   &    | 逻辑 AND 运算符。 如果两边的操作数都是 True，则为 True，否则为 False |
|  \|\|  | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则为 True，否则为 False |
|   !    | 逻辑 NOT 运算符。 如果条件为 True，则为 False，否则为 True   |

```Go
//logistic 逻辑运算符
func logistic() {
	var a float32 = 8
	var b float32 = 3
	var c float32 = 8
	fmt.Printf("a>b && b>c吗 %t\n", a > b && b > c)
	fmt.Printf("a>b || b>c吗 %t\n", a > b || b > c)
	fmt.Printf("a>b不成立，对吗 %t\n", !(a > b))
	fmt.Printf("b>c不成立，对吗 %t\n", !(b > c))
}
```

位运算符

| 运算符 | 描述                                                         |
| :----: | :----------------------------------------------------------- |
|   &    | 参与运算的两数各对应的二进位相与（两位均为1才为1）           |
|   \|   | 参与运算的两数各对应的二进位相或（两位有一个为1就为1）       |
|   ^    | 参与运算的两数各对应的二进位相异或，当两对应的二进位相同时为0，不同时为1。作为一元运算符时表示按位取反，，符号位也跟着变 |
|   <<   | 左移n位就是乘以2的n次方。a<<b是把a的各二进位全部左移b位，高位丢弃，低位补0。通过左移，符号位可能会变 |
|   >>   | 右移n位就是除以2的n次方。a>>b是把a的各二进位全部右移b位，正数高位补0，负数高位补1 |

```Go
//bit_op 位运算
func bit_op() {
	fmt.Printf("os arch %s, int size %d\n", runtime.GOARCH, strconv.IntSize) //int是4字节还是8字节，取决于操作系统是32位还是64位
	var a int32 = 260
	fmt.Printf("260     %s\n", util.BinaryFormat(a))
	fmt.Printf("-260    %s\n", util.BinaryFormat(-a)) //负数用补码表示。在对应正数二进制表示的基础上，按拉取反，再末位加1
	fmt.Printf("260&4   %s\n", util.BinaryFormat(a&4))
	fmt.Printf("260|3   %s\n", util.BinaryFormat(a|3))
	fmt.Printf("260^7   %s\n", util.BinaryFormat(a^7))     //^作为二元运算符时表示异或
	fmt.Printf("^-260   %s\n", util.BinaryFormat(^-a))     //^作为一元运算符时表示按位取反，符号位也跟着变
	fmt.Printf("-260>>10 %s\n", util.BinaryFormat(-a>>10)) //正数高位补0，负数高位补1
	fmt.Printf("-260<<3 %s\n", util.BinaryFormat(-a<<3))   //负数左移，可能变成正数
	//go语言没有循环（无符号）左/右移符号   >>>  <<<
}
```

```GO
//输出一个int32对应的二进制表示
func BinaryFormat(n int32)string  {
	a := uint32(n)
	sb := strings.Builder{}
	c := uint32(math.Pow(2,31)) // 最高位上是1，其他位全是0
	for i := 0;i<32;i++{
		if a&c != 0{ //111 101
			sb.WriteString("1")
		}else{
			sb.WriteString("0")
		}
		c >>= 1 // "1"往右移一位
 	}
 	return sb.String()
}
```







赋值运算符

| 运算符 | 描述                                           |
| :----: | :--------------------------------------------- |
|   =    | 简单的赋值运算符，将一个表达式的值赋给一个左值 |
|   +=   | 相加后再赋值                                   |
|   -=   | 相减后再赋值                                   |
|   *=   | 相乘后再赋值                                   |
|   /=   | 相除后再赋值                                   |
|   %=   | 求余后再赋值                                   |
|  <<=   | 左移后赋值                                     |
|  >>=   | 右移后赋值                                     |
|   &=   | 按位与后赋值                                   |
|  \|=   | 按位或后赋值                                   |
|   ^=   | 按位异或后赋值                                 |

```Go
//assignment 赋值运算
func assignment() {
	var a, b int = 8, 3
	a += b
	fmt.Printf("a+=b %d\n", a)
	a, b = 8, 3
	a -= b
	fmt.Printf("a-=b %d\n", a)
	a, b = 8, 3
	a *= b
	fmt.Printf("a*=b %d\n", a)
	a, b = 8, 3
	a /= b
	fmt.Printf("a/=b %d\n", a)
	a, b = 8, 3
	a %= b
	fmt.Printf("a%%=b %d\n", a) //%在fmt里有特殊含意，所以需要前面再加个%转义一下
	a, b = 8, 3
	a <<= b
	fmt.Printf("a<<=b %d\n", a)
	a, b = 8, 3
	a >>= b
	fmt.Printf("a>>=b %d\n", a)
	a, b = 8, 3
	a &= b
	fmt.Printf("a&=b %d\n", a)
	a, b = 8, 3
	a |= b
	fmt.Printf("a|=b %d\n", a)
	a, b = 8, 3
	a ^= b
	fmt.Printf("a^=b %d\n", a)
}
```

### 变量、常量、字面量

#### 变量类型

|   类型   |                         go变量类型                         | fmt输出  |
| :------: | :--------------------------------------------------------: | :------: |
|   整型   | int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 |    %d    |
|  浮点型  |                      float32 float64                       | %f %e %g |
|  布尔型  |                            bool                            |    %t    |
|   指针   |                          uintptr                           |    %p    |
|   引用   |                     map slice channel                      |    %v    |
|   字节   |                            byte                            |    %c    |
| 任意字符 |                            rune                            |    %c    |
|  字符串  |                           string                           |    %s    |
|   错误   |                           error                            |    %v    |

#### 变量声明

&#8195;&#8195;Go语言变量必须先声明再使用，所谓使用指读取或修改。  
标题声明

```Go
var name string 
var age int 
var isOk bool
```

批量声明

```Go
var ( 
	name string 
	age int 
	isOk bool 
)
```

#### 变量初始化

&#8195;&#8195;如果声明后未显式初始化，数值型初始化0，字符串初始化为空字符串，布尔型初始化为false，引用类型、函数、指针、接口初始化为nil。

```Go
var a string="china"  //初始化一个变量
var a="china"  //类型推断为string
var a,b int=3,7  //初始化多个变量
var a,b="china",7  //初始化多个变量，每个变量都单独地执行类型推断     
```

&#8195;&#8195;函数内部的变量(非全局变量)可以通过:=声明并初始化。

```Go
a:=3
```

&#8195;&#8195;下划线表示匿名变量。匿名变量不占命名空间，不会分配内存，因此可以重复使用。

```Go
_=2+4
```

#### 常量

&#8195;&#8195;常量在定义时必须赋值，且程序运行期间其值不能改变。

```Go
const PI float32=3.14

const(
    PI=3.14
    E=2.71
)

const(
    a=100
    b	//100，跟上一行的值相同
    c	//100，跟上一行的值相同
)
```

iota

```Go
const(
    a=iota	//0
    b		//1
    c		//2
    d		//3
)

const(
    a=iota 	//0
    b		//1
    _		//2
    d		//3
)

const(
    a=iota 	//0
    b=30    
    c=iota 	//2
    d		//3
)

const(
    _=iota		// iota =0
    KB=1<<(10* iota) 	// iota =1
    MB=1<<(10* iota) 	// iota =2
    GB=1<<(10* iota) 	// iota =3
    TB=1<<(10* iota) 	// iota =4
)

const(
    a,b=iota+1, iota+2	//1,2  iota =0
     c,d			//2,3  iota =1
     e,f			//3,4  iota =2
)
```

#### 字面量

&#8195;&#8195;字面量--没有出现变量名，直接出现了值。基础类型的字面量相当于是常量。

```Go
fmt.Printf("%t\n", 04 == 4.00) //用到了整型字面量和浮点型字面量
fmt.Printf("%v\n", .4i) //虚数字面量 0.4i
fmt.Printf("%t\n", '\u4f17' == '众') //Unicode和rune字面量
fmt.Printf("Hello\nWorld\n!\n") //字符串字面量
```

### 变量作用域

&#8195;&#8195;对于全局变量，如果以大写字母开头，所有地方都可以访问，跨package访问时需要带上package名称；如果以小写字母开头，则本package内都可以访问。  
&#8195;&#8195;函数内部的局部变量，仅本函数内可以访问。{}可以固定一个作用域。内部声明的变量可以跟外部声明的变量有冲突，以内部的为准--就近原则。

```Go
var (
    A=3	//所有地方都可以访问
    b=4	//本package内可以访问
)

func foo(){
    b:=5  //本函数内可以访问
    {
        b:=6  //本作用域内可以访问
    }
}
```

### 注释与godoc

#### 注释的形式

- 单行注释，以//打头。
- 多行注释有2种形式：
  1. 连续多行以//打头，注意多行注释之间不能出现空行。
  2. 在段前使用/\*，段尾使用*/。
- 注释行前加缩进即可写go代码。
- 注释中给定的关键词。NOTE: 引人注意，TODO: 将来需要优化，Deprecated: 变量或函数强烈建议不要再使用。

```Go
//Add 2个整数相加
//返回和。
//
//NOTE: 注释可以有多行，但中间不能出现空行（仅有//不算空行）。
func Add(a, b int) int {
	return a + b
}

/*
Sub 函数使用示例：
  for i:=0;i<3;i++{
	  Sub(i+1, i)
  }
看到了吗？只需要行前缩进，注释里就可以写go代码，是不是很简单。
*/
func Sub(a, b int) int {
	return a - b
}

//TODO: Prod 该函数不能并发调用，需要优化
func Prod(a, b int) int {
	return a * b
}

//Deprecated: Div 不要再调用了
func Div(a, b int) int {
	return a / b
}
```

#### 注释的位置

&#8195;&#8195;针对行的注释在行上方或右侧。函数的上方在func xxx()上方。结构体的注释在type xxx struct上方。包注释在package xxx的上方。一个包只需要在一个地方写包注释，通常会专门写一个doc.go，里面只有一行package xxx和关于包的注释。

```Go
// FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:
//
//	s := strconv.FormatBool(true)
//	s := strconv.FormatFloat(3.1415, 'E', -1, 64)
//	s := strconv.FormatInt(-42, 16)
//	s := strconv.FormatUint(42, 16)
package fmt
```

#### go doc

&#8195;&#8195;go doc是go自带的命令。

```Shell
go doc entrance_class/util
```

上述命令查看entrance_class/util包的注释。

#### godoc

&#8195;&#8195;godoc是第三方工具，可以为项目代码导出网页版的注释文档。安装godoc命令如下

```Shell
go get -u golang.org/x/tools/cmd/godoc
go install golang.org/x/tools/cmd/godoc@latest
```

启动http服务：

```Shell
godoc -http=:6060
```

用浏览器访问http://127.0.0.1:6060 ，可以查看go标准库的文档。





















