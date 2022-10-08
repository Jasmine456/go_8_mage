package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go_8_mage/old_project/cloud-station-git/store"
	"github.com/go-playground/validator"
)

//使用Validate的单个实例，它会缓存结构信息
var(
	validate = validator.New()
)

//构造函数
func NewUploader(endpoint,accessKey,secretKey string) (store.Uploader,error) {
	a:=&ali{
		Endpoint: endpoint,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
	if err:=a.validate();err!=nil{
		return nil,err
	}

	return a,nil
}

type ali struct {
	Endpoint  string `validate:"required,url"`
	AccessKey string `validate:"required"`
	SecretKey string `validate:"required"`
}

func (a *ali) validate() error{
	return validate.Struct(a)
}

func (a *ali) UploadFile(bucketName, objectKey, localFilePath string) error {
	bucket,err:=a.GetBucket(bucketName)
	if err!=nil{
		return err
	}

	err=bucket.PutObjectFromFile(objectKey,localFilePath)
	if err != nil {
		return fmt.Errorf("upload file to bucket: %s error,%s",bucketName,err)
	}
	singedURL,err:=bucket.SignURL(objectKey,oss.HTTPGet,60*60*24)
	if err != nil {
		return fmt.Errorf("SignURL error,%s",err)
	}
	fmt.Printf("下载链接：%s\n",singedURL)
	fmt.Println("\n注意：文件下载有效期为1天，中转站保存时间为3天，请及时下载")
	return nil

	//return fmt.Errorf("not impl")
}

func (a *ali) GetBucket(bucketName string) (*oss.Bucket, error) {
	if bucketName == "" {
		return nil, fmt.Errorf("upload bucket name required")
	}

	//	New client
	client, err := oss.New(a.Endpoint, a.AccessKey, a.SecretKey)
	if err != nil {
		return nil, err
	}

	//	Get bucket
	bucket,err:=client.Bucket(bucketName)
	if err!=nil{
		return nil,err
	}
	return bucket,nil
}
