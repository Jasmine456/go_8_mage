package aliyun

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	bucketName    = "devcloud-station-jasmine"
	objectKey     = "store.go"
	localFilePath = "store.go"

	endpoint = "http://oss-cn-hangzhou.aliyuncs.com"
	ak       = os.Getenv("ALI_AK")
	sk       = os.Getenv("ALI_SK")
)

func TestUploadFile(t *testing.T){
	//fmt.Println("hello test detail log")
	should:=assert.New(t)

	uploader,err := NewUploader(endpoint,ak,sk)
	if should.NoError(err){
		err=uploader.UploadFile(bucketName,objectKey,localFilePath)
		should.NoError(err)
	}
}
