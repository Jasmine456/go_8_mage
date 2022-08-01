# Golang笔记\_第6周\_包与工程化和常用标准库



## 包与工程化

### 用 go mod管理工程

初始化项目：

```GO
go mod init $module_name
```



$module_name和目录名可以不一样。上述命令会生成 go.mod文件，该文件内容形如：

```GO
module go-course

go 1.17

require (
    github.com/ethereum/go-ethereum v1.10.8
    github.com/gin-gonic/gin v1.7.4
)
```



go依次从当前项目、$GOROOT 、$GOPATH下寻找依赖包。

1. 从当前go文件所在的目录逐级向上查找go.mod文件（假设go.mod位于目录$mode_path下），里面定义了module_name，则引入包的路径为module_name/包相对于$mode_path的路径。
2. go标准库提供的包在$GOROOT/src下。
3. 第三方依赖包在$GOPATH/pkg/mod下。



​	从go1.7开始，go get 只负责下载第三方依赖包，并把它加到go.mod 文件里，由go install 负责安装二进制文件。

> - go get github.com/mailru/easyjson会在$GOPATH/pkg/mod目录下生成github.com/mailru/easyjson目录。
> - go install github.com/mailru/easyjson/easyjson会在$GOPATH/bin下生成easyjson二进制可执行文件。
>   &#8195;&#8195;go mod tidy通过扫描当前项目中的所有代码来添加未被记录的依赖至go.mod文件或从go.mod文件中删除不再被使用的依赖。 



### 包引入规则

包的声明

> - GO文件再第一行声明 package xxx。
> - 在包声明的上面可写关于包的注释，包注释也可以专门写在doc.go 里
> - 包名跟目录名可以不同。
> - 同一个目录下，所有go文件的包名必须一致
> - 包的引用   可以直接使用同目录下其他go文件里的变量、函数、结构体
> - 跨目录使用则需要变量前加入包名，并且引入包所在的目录。

```GO
import  "go-course/package"
mypackage.Add()     //mypackage是包名，它所在的目录是go-course/package
```



> - 在import块里可以引用父目录，也可以引用子目录。
> - 引用关系不能构成一个环
> - 在import的目录前面可以给包起一个别名

```GO
imoprt asd "go-course/package"
asd.Add()
```



### init调用链

​		main函数是go程序的唯一入口，所以main函数只能存在一个。main函数必须位于main包中。在main函数执行之前会先执行inti()函数。在一个目录，甚至一个go文件里，init()可以重复定义。引入其他包时，相应包里的init()函数也会在main()函数之前被调用。



![init链](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC6%E5%91%A8_%E5%8C%85%E4%B8%8E%E5%B7%A5%E7%A8%8B%E5%8C%96%E5%92%8C%E5%B8%B8%E7%94%A8%E6%A0%87%E5%87%86%E5%BA%93.assets/init%E9%93%BE.png)





```GO
import _ "net/httppprof"
```

​		在目录前一个_,代码里却没有显示地使用这个包里的函数或变量，实际事项执行这个包里的init()函数。



### 可见性

> - 以小写字母开头命名的函数、变量、结构体只能在本包内访问。
> - 以大写字母开头命名的函数、变量、结构体在其他包中也可以访问。
> - 如果结构体名字以答谢字母开头，而其成员变量、成员方法以小写字母开头，则这样的成员只能在本包内访问。



GO中命名为internal的package，只有该package的上一级package才可以访问该package的内容。如下图c目录(internal的上一级目录) 及其子孙目录之间可以任意import，但a目录和b目录不能import internal及其下属的所有目录。



![path](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC6%E5%91%A8_%E5%8C%85%E4%B8%8E%E5%B7%A5%E7%A8%8B%E5%8C%96%E5%92%8C%E5%B8%B8%E7%94%A8%E6%A0%87%E5%87%86%E5%BA%93.assets/path.png)

![internal](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC6%E5%91%A8_%E5%8C%85%E4%B8%8E%E5%B7%A5%E7%A8%8B%E5%8C%96%E5%92%8C%E5%B8%B8%E7%94%A8%E6%A0%87%E5%87%86%E5%BA%93.assets/internal.png)







## 常用标准库

### 数学计算

数学常量

```GO
math.E	//自然对数的底，2.718281828459045
math.Pi	//圆周率，3.141592653589793
math.Phi	//黄金分割，长/短，1.618033988749895
math.MaxInt	//9223372036854775807
uint64(math.MaxUint)	//得先把MaxUint转成uint64才能输出，18446744073709551615
math.MaxFloat64	//1.7976931348623157e+308
math.SmallestNonzeroFloat64	//最小的非0且正的浮点数，5e-324
```

NaN(Not a Number)

```GO
f := math.NaN()
math.IsNaN(f)
```



常用函数

```GO
math.Ceil(1.1)  //向上取整，2
math.Floor(1.9)	//向下取整，1。 math.Floor(-1.9)=-2
math.Trunc(1.9)	//取整数部分，1
math.Modf(2.5)	//返回整数部分和小数部分，2  0.5
math.Abs(-2.6)	//绝对值，2.6
math.Max(4, 8)	//取二者的较大者，8
math.Min(4, 8)	//取二者的较小者，4
math.Mod(6.5, 3.5)	//x-Trunc(x/y)*y结果的正负号和x相同，3
math.Sqrt(9)		//开平方，3
math.Cbrt(9)		//开三次方，2.08008
```



三角函数

```go
math.Sin(1)
math.Cos(1)
math.Tan(1)
math.Tanh(1)
```



对数和指数

```GO
math.Log(5)	//自然对数，1.60943
math.Log1p(4)	//等价于Log(1+p)，确保结果为正数，1.60943
math.Log10(100)	//以10为底数，取对数，2
math.Log2(8)	//以2为底数，取对数，3
math.Pow(3, 2)	//x^y，9
math.Pow10(2)	//10^x，100
math.Exp(2)	//e^x，7.389
```



随机数生成器

```GO
//创建一个Rand
source := rand.NewSource(1) //seed相同的情况下，随机数生成器产生的数列是相同的
rander := rand.New(source)
for i := 0; i < 10; i++ {
    fmt.Printf("%d ", rander.Intn(100))
}
fmt.Println()
source.Seed(1) //必须重置一下Seed
rander2 := rand.New(source)
for i := 0; i < 10; i++ {
    fmt.Printf("%d ", rander2.Intn(100))
}
fmt.Println()

//使用全局Rand
rand.Seed(1)                //如果对两次运行没有一致性要求，可以不设seed
fmt.Println(rand.Int())     //随机生成一个整数
fmt.Println(rand.Float32()) //随机生成一个浮点数
fmt.Println(rand.Intn(100)) //100以内的随机整数，[0,100)
fmt.Println(rand.Perm(100)) //把[0,100)上的整数随机打乱
arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
rand.Shuffle(len(arr), func(i, j int) { //随机打乱一个给定的slice
    arr[i], arr[j] = arr[j], arr[i]
})
fmt.Println(arr)
```



### 时间函数

时间的解析和格式化

```GO
TIME_FMT := "2006-01-02 15:04:05"
now := time.Now()
ts := now.Format(TIME_FMT)
loc,_ = time.LoadLocation("Asia/Shanghai")
t,_ = time.ParseInLocation(TIME_FMT, ts, loc)
```

时间运算

```GO
diff1 := t1.Sub(t0) //计算t1跟t0的时间差，返回类型是time.Duration
diff2 := time.Since(t0) //计算当前时间跟t0的时间差，返回类型是time.Duration
diff3 := time.Duration(3 * time.Hour) //Duration表示两个时刻之间的距离
t4 := t0.Add(diff3) 
t4.After(t0)    //true
```



时间的属性

```GO
t0.Unix(), t0.UnixMilli(), t0.UnixMicro(), t0.UnixNano()
t2.Year(), t2.Month(), t2.Day(), t2.YearDay()
t2.Weekday().String(), t2.Weekday()
t1.Hour(), t1.Minute(), t1.Second()
```

定时执行

```GO
tm := time.NewTimer(3 * time.Second)
<-tm.C //阻塞3秒钟
//do something
tm.Stop()

//或者用：
<-time.After(3 * time.Second) //阻塞3秒钟
```



周期执行

```GO
tk := time.NewTicker(1 * time.Second)
for i := 0; i < 10; i++ {
    <-tk.C //阻塞1秒钟
    //do something
}
tk.Stop()
```





### I/O操作

格式化输出

| 输出格式 |                         输出内容                         |
| :------: | :------------------------------------------------------: |
|    %t    |                    单词 true 或 false                    |
|    %b    |                       表示为二进制                       |
|    %d    |                       表示为十进制                       |
|    %e    |  （=%.6e）有 6 位小数部分的科学计数法，如 -1234.456e+78  |
|    %f    |         （=%.6f）有 6 位小数部分，如 123.456123          |
|    %g    | 根据实际情况采用 %e 或 %f 格式（获得更简洁、准确的输出） |
|    %s    |                直接输出字符串或者字节数组                |
|    %v    |                     值的默认格式表示                     |
|   %+v    |           类似 %v，但输出结构体时会添加字段名            |
|   %#v    |                     值的 Go 语法表示                     |
|    %Т    |                  值的类型的 Go 语法表示                  |



标准输入

```GO
fmt.Println("please input two word")
var word1 string 
var word2 string
fmt.Scan(&word1, &word2) //读入多个单词，空格分隔。如果输入了更多单词会被缓存起来，丢给下一次scan

fmt.Println("please input an int")
var i int
fmt.Scanf("%d", &i) //类似于Scan，转为特定格式的数据
```



打开文件

```GO
func os.Open(name string) (*os.File, error)
fout, err := os.OpenFile("data/verse.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
```



​		os.O_WRONLY以只写的方式打开文件，os.O_TRUNC把文件之前的内容先清空掉，os.O_CREATE如果文件不存在则先创建，0666新建文件的权限设置。

读文件

```GO
cont := make([]byte, 10)
fin.Read(cont) //读出len(cont)个字节，返回成功读取的字节数
fin.ReadAt(cont, int64(n)) //从指定的位置开始读len(cont)个字节
fin.Seek(int64(n), 0) //重新定位。whence: 0从文件开头计算偏移量，1从当前位置计算偏移量，2到文件末尾的偏移量
```



```GO
reader := bufio.NewReader(fin) //读文件文件建议用bufio.Reader
for { //无限循环
    if line, err := reader.ReadString('\n'); err != nil { //指定分隔符
        if err == io.EOF {
            if len(line) > 0 { //如果最后一行没有换行符，则此时最后一行就存在line里
                fmt.Println(line)
            }
            break //已读到文件末尾
        } else {
            fmt.Printf("read file failed: %v\n", err)
        }
    } else {
        line = strings.TrimRight(line, "\n") //line里面是包含换行符的，需要去掉
        fmt.Println(line)
    }
}
```



写文件

```GO
defer fout.Close() //别忘了关闭文件句柄
writer := bufio.NewWriter(fout)
writer.WriteString("明月多情应笑我")
writer.WriteString("\n") //需要手动写入换行符
```

创建文件/目录

```GO
os.Create(name string)//创建文件
os.Mkdir(name string, perm fs.FileMode)//创建目录
os.MkdirAll(path string, perm fs.FileMode)//增强版Mkdir，沿途的目录不存在时会一并创建
os.Rename(oldpath string, newpath string)//给文件或目录重命名，还可以实现move的功能
os.Remove(name string)//删除文件或目录，目录不为空时才能删除成功
os.RemoveAll(path string)//增强版Remove，所有子目录会递归删除
```



遍历目录

```GO
if fileInfos, err := ioutil.ReadDir(path); err != nil {
	return err
} else {
    for _, fileInfo := range fileInfos {
    fmt.Println(fileInfo.Name())
    if fileInfo.IsDir() { //如果是目录，就递归子遍历
        walk(filepath.Join(path, fileInfo.Name}
    }
}
```



默认的log输出到控制台

```GO
log.Printf("%d+%d=%d\n", 3, 4, 3+4)
log.Println("Hello Golang")
log.Fatalln("Bye, the world") //日志输出后会执行os.Exit(1)
```

指定日志输出到文件

```GO
cmd_path, err := exec.LookPath(“df”) //查看系统命令所在的目录，确保命令已安装
cmd := exec.Command("df", "-h") //相当于命令df -h，注意Command的每一个参数都不能包含空格
output, err := cmd.Output() //cmd.Output()运行命令并获得其输出结果
cmd = exec.Command("rm", "./data/test.log")
cmd.Run() //如果不需要获得命令的输出，直接调用cmd.Run()即可
```



### 编码

​		json是go标准库里自带的序列化工具，使用了反射，效率比较低。

easyjson只针对预先定义好的json结构体对输入的json字符串进行纯字符串的截取，并将对应的json字段赋值给结构体。easyjson-all xxx.go 生成go文件中定义的结构体对应的解析，xxx.go所在的package不能是main。

```GO
func easyjson.Marshal(v easyjson.Marshaler) ([]byte, error)
func easyjson.Unmarshal(data []byte, v easyjson.Unmarshaler) error
```



sonic是字节跳动开源的json序列化工具包，号称性能强过easyjson、jsoniter，使用起来非常方便。

```GO
import "github.com/bytedance/sonic"

// Marshal
output, err := sonic.Marshal(&data) 
// Unmarshal
err := sonic.Unmarshal(input, &data) 
```



​		base64经常在http环境下用来传输较长的信息。任意byte数组都可以采用base64编码转为字符串，并且可以反解回byte数组。编码和解码的方法是公开、确定的，base64不属于加密算法。

```GO
func (*base64.Encoding).EncodeToString(src []byte) string
func (*base64.Encoding).DecodeString(s string) ([]byte, error)
```



​		compress包下实现了zlib、bzip、gip、lzw等压缩算法。

```GO
writer := zlib.NewWriter(fout)//压缩
writer.Write(bytes)
reader, err := zlib.NewReader(fin) //解压
io.Copy(os.Stdout, reader) 
```

