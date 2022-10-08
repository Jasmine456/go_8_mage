# 阿里云盘上传下载功能 用go的cli实现（面向对象版本2

interface.go
定义接口
```go
type Uploader interface{
	Upload(bucketName string,objectKey string,fileName string)error
}
```


