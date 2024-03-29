

[toc]



# 基础部分实战项目--特征抽取
## 什么是特征抽取
&#8195;&#8195;电商领域给用户推荐商品，计算用户点击每个商品的概率，按概率从大到小排序。概率计算公式：$\hat{y}=𝑤_1𝑥_1+𝑤_2𝑥_2+𝑤_3 𝑥_3+\ldots$，w是权重，x非0即1，代表各个特征具体的取值，比如  

![image-20220606155024842](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC7%E5%91%A8_%E7%89%B9%E5%BE%81%E6%8A%BD%E5%8F%96%20%E7%AD%96%E7%95%A5%E7%B3%BB%E7%BB%9F%E9%85%8D%E7%BD%AE%E5%8C%96%E7%AE%A1%E7%90%86.assets/image-20220606155024842.png)



连续特征离散化:   
1. 分箱法。比如年龄，指定分割点18,25,30,35,40
2. 取对数。比如销量对2取对数
3. 取一个字段。比如时间，取其Hour
4. ……  

给特征分配编号  
- 离散化之后，有些特征的取值范围仍然是不确定的，比如销量的上限是不确定的，所以 log_2 销量 也是不确定的。无法顺序地给为每个特征的每个取值进行编号。
- 索性通过哈希函数把特征取值映射到[0, 2^64)上，并以此作为其编号。几乎可以认为特征取值不同，其编号就不同。
- 不同的特征可能存在相同的取值（比如采用分箱法离散化的特征），所以特征和特征取值要一并喂给哈希函数。  

&#8195;&#8195;所谓特征抽取，就是给定商品的原始特征，提取出模型所需要的特征，进行预处理（比如离散化），最终把每个特征表示成一个或多个uint64。  

## 特征抽取的设计思路
功能设计  
- 商品struct是事先定义好的。
    - type Product struct{Sales int; Tags []string}
- 可用的离散化方法和哈希函数是事先定义好的。
- 通过配置文件来指定商品的各个字段应该采用哪种离散化方法和哪个哈希函数。
- 配置可以灵活改动，以快速验证不同特征、不同预处理方法对模型效果的影响。  

代码设计  
基础结构体  
```Go
type Location struct {
	Province string
	City string
}

type User struct {
	Name string
	Age int
	Gender byte
	Address *Location
}

type Product struct {
    Id int
    Name string
    Sales int //销量
    Feedback float32 //好评率
    Seller *User //商家
    OnShelfTime time.Time //上架时间
    Tags []string
}
```
基础接口  
```Go
//离散化接口
type Discretizer interface {
        Discretize(i interface{}) string
}

//哈希接口
type Transformer interface {
        Hash(string, int) uint64
}
```
特征转换配置类   
```Go
type FeatureConfig struct {
    Path string //从Product中取得字段
    DiscretizeFunc Discretizer
    HashFunc Transformer
}
```
配置文件  
```
[
    {
        "id": 8,
        "path": "Feedback",
        "discretize": "bin 0.1,0.2,0.3,0.4,0.5,0.6,0.7,0.8,0.9",
        "hash": "farm"
    },
    {
        "id": 9,
        "path": "OnShelfTime",
        "discretize": "hour",
        "hash": "farm"
    }
]
```
&#8195;&#8195;使用反射根据path拿到对应Field的取值，然后送给离散化和哈希函数。discretize指示使用哪个离散化函数以及对应的参数，hash指示使用哪个哈希函数。  
## 性能测试  
单元测试  
```Go
func TestStrCat(b *testing.T) {
	hello := "hello"
	golang := "golang"
	fmt.Printf("%s %s\n", hello, golang)
}
```
```Shell
go test -v go_test.go -timeout=20m -count=1
```
- -v 打印详情测试信息
- -timeout 默认10分钟超时
- -count 函数运行几次  

基准测试  
```Go
func BenchmarkStrCat(b *testing.B) {
    hello := "hello"
    golang := "golang"
    for i := 0; i < b.N; i++ {
        fmt.Printf("%s %s\n", 	hello, golang)
    }
}
```
```Shell
go test -bench=StrCat -run=^$ -benchmem -benchtime=2s -cpuprofile=data/cpu.prof -memprofile=data/mem.prof
```
- -bench 正则指定运行哪些基准测试
- -run 正则指定运行哪些单元测试
- -benchmem 输出内存分配情况
- -benchtime 每个函数运行多长时间  

测试代码规范  
- 单元测试和基准测试必须放在以_test.go为后缀的文件里。
- 单元测试函数以Test开头，基准测试函数以Benchmark开头。
- 单元测试以*testing.T为参数，函数无返回值。
- 基准测试以*testing.B为参数，函数无返回值。  

proof是可视化性能分析工具，提供以下功能：
1. CPU Profiling：按一定频率采集CPU使用情况。
2. Memory Profiling：监控内存使用情况，检查内存泄漏。
3. Goroutine Profiling：对正在运行的Goroutine进行堆栈跟踪和分析，检查协程泄漏。  

监控CPU使用命令go tool pprof data/cpu.prof。进入交互界面后常用的命令有：  
- topn：列出最耗计算资源的前n个函数
- list func：列出某函数里每一行代码消耗多少计算资源
- peek func：列出某函数里最耗计算资源的前几个子函数  

![image-20220606155119577](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC7%E5%91%A8_%E7%89%B9%E5%BE%81%E6%8A%BD%E5%8F%96%20%E7%AD%96%E7%95%A5%E7%B3%BB%E7%BB%9F%E9%85%8D%E7%BD%AE%E5%8C%96%E7%AE%A1%E7%90%86.assets/image-20220606155119577.png)

pprof结果可以在浏览器上进行可视化。  
```Go
go tool pprof -http=:8080 data/cpu.prof
```
![image-20220606155107746](Golang%E7%AC%94%E8%AE%B0_%E7%AC%AC7%E5%91%A8_%E7%89%B9%E5%BE%81%E6%8A%BD%E5%8F%96%20%E7%AD%96%E7%95%A5%E7%B3%BB%E7%BB%9F%E9%85%8D%E7%BD%AE%E5%8C%96%E7%AE%A1%E7%90%86.assets/image-20220606155107746.png)