package cmd

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"
	"go_8_mage/week14/vblog/api/conf"
	"go_8_mage/week14/vblog/api/protocol"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	confType string
	confFile string
)

// StartCmd represents the base command when called without any subcommands
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动API Server",
	Long:  "启动API Server",
	RunE: func(cmd *cobra.Command, args []string) error {
		//编写程序启动逻辑

		//加载配置
		if err := loadConfig(); err != nil {
			return err
		}

		//初始化全局变量
		loadGlobal()

		//需要监听来自os的信号,比如你取消了或者终止了服务
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

		// http server 启动是阻塞的
		http := protocol.NewHTTP()

		//使用阻塞channel 来等待协程退出
		wg:=&sync.WaitGroup{}
		wg.Add(1)
		go func() {
			//多个Goroutine 同时执行的，有可能还没来得及+1，wg就退出了
			//wg.Add(1)

			defer wg.Done()
			//	启动一个Goroutine再后台，处理来自操作系统的信号
			for v := range ch {
				zap.L().Infof("receive signal: %s,stop service", v)

				switch v {
				case syscall.SIGHUP:
					if err:=loadConfig();err!=nil{
						zap.L().Errorf("reload config error,%s",err)
					}
				default:
					//	优雅关闭HTTP服务
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()
					http.Stop(ctx)
				}

				//退出循环,Goroutine 退出
				return
			}
		}()

		if err:=http.Start();err!=nil{
			return err
		}

		//等待程序优雅关闭完成
		wg.Wait()

		return nil
	},
}

func loadConfig() error {
	switch confType {
	case "env":
		return conf.LoadConfigFromEnv()
	case "file":
		return conf.LoadConfigFromToml(confFile)
	default:
		return fmt.Errorf("not supported config type,%s", confType)
	}
}

func loadGlobal() {
	//全局日志对象
	zap.DevelopmentSetup()
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confType, "config-type", "t", "file", "the service config type [file/env/etcd")
	StartCmd.PersistentFlags().StringVarP(&confFile, "config-file", "f", "etc/config.toml", "the service config from file")

	RootCmd.AddCommand(StartCmd)
}
