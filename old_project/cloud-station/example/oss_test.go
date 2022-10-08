package example_test

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"testing"
)

//定义一个全局的client,在包加载时初始化（init()函数）
var (
	client *oss.Client
)

var(
	Accesskey = os.Getenv("ALI_AK")
	AccessSecret=os.Getenv("ALI_SK")
	OssEndpoint=os.Getenv("ALI_OSS_ENDPOINT")
	BucketName=os.Getenv("ALI_BUCKET_NAME")
)

//测试阿里云ossSDK BucketList接口
func TestBucketList(t *testing.T){

	lsRes, err := client.ListBuckets()
	if err != nil {
		t.Log(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}

func TestUploadFile(t *testing.T){
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		t.Log(err)
	}

	// 上传文件到bucket中
	// 常见我们文件，需要创建一个文件夹
	// 云商OssServer会根据你的key 路径结构，自动的帮你创建目录
	// objectKey 上传到bucket里面的对象的名称（包含路径）
	// mkdir/test.go,oss server 会自动创建一个mydir的目录，mkdir -pv
	// 把测试当前这个文件上传到mydir下
	err = bucket.PutObjectFromFile("mkdir/test.go", "oss_test.go")
	if err != nil {
		t.Log(err)
	}
}

//初始化一个oss Client，等下给其他所有测试用例使用
func init(){
	c, err := oss.New(OssEndpoint, AccessSecret,Accesskey)
	fmt.Println(AccessSecret)
	if err != nil {
		// HandleError(err)
		panic(err)
	}
	client=c
}