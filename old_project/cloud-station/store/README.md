# 阿里云盘上传下载功能 用go的cli实现（面向对象版本
interface.go
定义接口 
```go
type Uploader interface{
	Upload(bucketName string,objectKey string,fileName string)error
}
```

定义结构体，即抽象出一个对象
对象一 AliOssStore
        属性 client *oss.Client (一个阿里云oss client类型的client属性)
        属性 listener 后期进度条添加的属性(进度条 progress)

    1. 给对象定义 默认的构造函数,NewAliOssStore。
        AliOssStore的构造函数需要传入三个 固定参数，将这三个参数抽象出一个对象传入此构造方法
    2. 将该对象实现 Uploader接口，即实现 Uploader中的所有方法，当前只有Upload 1个方法 
    3. 编写测试用例

对象二 Options
        属性 Endpoint, AccessKey, AccessSecret string
    将oss clinet所需的三个参数整合为一个Options对象，三个参数定义为这个对象的属性
    1. validate验证参数的 方法

    


