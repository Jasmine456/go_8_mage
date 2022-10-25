package cmd

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/app"
	"github.com/spf13/cobra"
	"github.com/go_8_mage/week14_after/devcloud/mcenter/apps/domain"
	"github.com/go_8_mage/week14_after/devcloud/mcenter/apps/user"
)

var (
	createTableFilePath string
)

// initCmd represents the start command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "mcenter 服务初始化",
	Long:  "mcenter 服务初始化",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化全局变量
		if err := loadGlobalConfig(confType); err != nil {
			return err
		}

		//初始化全局日志配置
		if err:=loadGlobalLogger();err!=nil{
			return err
		}

		//初始化全局app
		if err:=app.InitAllApp();err!=nil{
			return err
		}

		//创建一个主账号 admin
		us:=app.GetInternalApp(user.AppName).(user.Service)
		fmt.Println(us)
		createUserReq:=user.NewCreateUserRequest()
		createUserReq.CreateBy=user.CREATE_BY_ADMIN
		createUserReq.Domain=domain.DEFAULT_DOMAIN
		createUserReq.Username="admin"
		createUserReq.Password="123456"
		createUserReq.Type=user.TYPE_SUPPER
		u,err:=us.CreateUser(context.Background(),createUserReq)

		if err != nil {
			return err
		}
		fmt.Printf("初始化管理员用户: %s 成功\n", u.Spec.Username)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
