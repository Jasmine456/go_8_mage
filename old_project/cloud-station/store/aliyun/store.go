package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go_8_mage/old_project/cloud-station/store"
)

var (
	// 检查 对象是否实现了接口的强制的约束
	_ store.Uploader = &AliOssStore{}
	//_ store.Uploader = (*AliOssStore)(nil)
)

type Options struct {
	Endpoint, AccessKey, AccessSecret string
}

func (o *Options) Validate() error {
	if o.Endpoint == "" || o.AccessKey == "" || o.AccessSecret == "" {
		return fmt.Errorf("endpoint,access_key access_secret has one empty")
	}
	return nil
}

// AliOssStore对象的构造函数
func NewAliOssStore(opts *Options) (*AliOssStore, error) {
	c, err := oss.New(opts.Endpoint, opts.AccessKey, opts.AccessSecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{
		client: c,
		listener: NewProgressListener(),
	}, nil
}

type AliOssStore struct {
	// 阿里云 OSS client，私有变量，不运行外部
	client *oss.Client
	//	 依赖listener的实现
	listener oss.ProgressListener
}

func (s *AliOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	//	2. 获取bucket对象
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return err
	}
	//	3. 上传文件到该bucket
	if err := bucket.PutObjectFromFile(objectKey, fileName,oss.Progress(s.listener)); err != nil {
		return err
	}

	//	4. 打印下载链接
	downloadURL, err := bucket.SignURL(fileName, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载URL:%s\n", downloadURL)
	fmt.Println("请在1天之内下载。")
	return nil
}
