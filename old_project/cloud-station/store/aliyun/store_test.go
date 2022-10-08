package aliyun_test

import (
	//"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/stretchr/testify/assert"
	"go_8_mage/old_project/cloud-station/store"
	"go_8_mage/old_project/cloud-station/store/aliyun"
	"os"
	"testing"
	"github.com/AlecAivazis/survey/v2"
)

var (
	uploader store.Uploader
)

var(
	Accesskey = os.Getenv("ALI_AK")
	AccessSecret=os.Getenv("ALI_SK")
	OssEndpoint=os.Getenv("ALI_OSS_ENDPOINT")
	BucketName=os.Getenv("ALI_BUCKET_NAME")
)

func TestUpload(t *testing.T){
	//使用 assert 编写测试用例的断言语句
	// 通过New获取一个断言实例
	should:=assert.New(t)

	err:=uploader.Upload(BucketName,"test.txt","store_test.go")

	//如果你不用断言库， if err ==nil{}
	//封装：assert，他把连个对象的比较 断言封装成了方法
	//should.Equal(BucketName,"aaa")

	if should.NoError(err){
	//	没有Error 开启下一个步骤
		t.Log("upload ok")
	}
}

func TestUploadError(t *testing.T){
	//使用 assert 编写测试用例的断言语句
	// 通过New获取一个断言实例
	should:=assert.New(t)

	err:=uploader.Upload(BucketName,"test.txt","store_test.goxxx")
	should.Error(err,"open store_test.goxxx: The system cannot find the file specified.")
}



// 通过init 编写uploader实例化逻辑
func init(){

	ali,err:=aliyun.NewAliOssStore(&aliyun.Options{
		OssEndpoint,
		Accesskey,
		AccessSecret,
	})
	if err != nil {
		panic(err)
	}
	uploader=ali

}
