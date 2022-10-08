package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"go_8_mage/old_project/cloud-station/store"
	"go_8_mage/old_project/cloud-station/store/aliyun"
	"go_8_mage/old_project/cloud-station/store/aws"
	"go_8_mage/old_project/cloud-station/store/tx"
)

var(
	ossProvier string
	ossEndpoint,accessKey,accessSecret,bucketName,uploadFile string
)

var (
	default_ak ="LTAI5tSk4B4ERmJw8ZMMqJyY"
	default_sk="4EWBkoA1TaygQPQ9duHaJEtEIYoQVZ"
)

var UploadCmd = &cobra.Command{
	Use:"upload",
	Long:"upload 文件上传",
	Short: "upload 文件上传",
	Example:"upload -f filename",
	RunE: func(cmd *cobra.Command,args []string)error{
		var (
			uploader store.Uploader
			err error
		)
		switch ossProvier {
		case "aliyun":
			aliOpts:=&aliyun.Options{
				Endpoint: ossEndpoint,
				AccessKey: accessKey,
				AccessSecret: accessSecret,
			}
			setAliDefault(aliOpts)
			uploader,err=aliyun.NewAliOssStore(aliOpts)

		case "tx":
			uploader=tx.NewTxOssStore()
		case "aws":
			uploader=aws.NewAwsOssStore()
		default:
			return fmt.Errorf("not support oss storage provider")
		}
		if err!=nil{
			return err
		}

		//使用Upload来上传文件
		return uploader.Upload(bucketName,uploadFile,uploadFile)
	},
}

func setAliDefault(opts *aliyun.Options){
	if opts.AccessKey==""{
		opts.AccessKey=default_ak
	}
	if opts.AccessSecret == ""{
		opts.AccessSecret=default_sk
	}
}

func init(){
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvier,"provider","p","aliyun","oss storage provier [aliyun/tx/aws]")
	f.StringVarP(&ossEndpoint,"endpoint","e","oss-cn-hangzhou.aliyuncs.com","oss storage provier endpoint")
	f.StringVarP(&accessKey,"access_key","k","","oss storage provier ak")
	f.StringVarP(&accessSecret,"access_secret","s","","oss storage provier sk")
	f.StringVarP(&bucketName,"bucket_name","b","devcloud-station-jasmine","oss storage provier bucketName")
	f.StringVarP(&uploadFile,"upload_file","f","","upload file name")

	RootCmd.AddCommand(UploadCmd)
}
