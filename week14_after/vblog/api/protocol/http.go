package protocol

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go_8_mage/week14/vblog/api/apps"
	_ "go_8_mage/week14/vblog/api/apps"
	_ "go_8_mage/week14/vblog/api/apps/all"
	//blog_impl "go_8_mage/week14/vblog/api/apps/blog/impl"
	//tag_impl "go_8_mage/week14/vblog/api/apps/tag/impl"
	"go_8_mage/week14/vblog/api/conf"
	"net/http"
	"time"
)

//HTTP服务，需要有简单的地址和端口,从全局配置中获取Server监听参数
func NewHTTP() *HTTP {
	r:=gin.Default()
	return &HTTP{
		log:zap.L().Named("server.http"),
		router: r,
		server: &http.Server{
			ReadHeaderTimeout: 60 * time.Second,
			ReadTimeout:       60 * time.Second,
			WriteTimeout:      60 * time.Second,
			IdleTimeout:       60 * time.Second,
			MaxHeaderBytes:    1 << 20, //1M
			//处理HTTP协议的逻辑，HTTP中间件，是一个路由(框架，Gin，...)与处理(自己)
			Handler:r,
			Addr:              conf.C().App.HTTP.Addr(),
		},
	}
}

//HTTP 服务对象
type HTTP struct {
	router *gin.Engine
	server *http.Server
	log logger.Logger
}

func (h *HTTP) Start() error {
	////没有Gin产生任何关系，业务模块的URL还没注册
	////blog 具体实现
	//svr:=blog_impl.NewImpl()
	////初始化
	//if err:=svr.Init();err!=nil{
	//	panic(err)
	//}
	////
	//tsvr:=tag_impl.NewImpl(svr)
	////初始化
	//if err:=tsvr.Init();err!=nil{
	//	panic(err)
	//}

	if err:= apps.Init();err!=nil{
		return err
	}

	v1:=h.router.Group("/vblog/api/v1")

	////方便Nginx /vblog --> xxx:8080
	////注册blog服务的Restful API
	//blogAPI:=blog_api.NewHTTPAPI()
	//if err:=blogAPI.Init();err!=nil{
	//	return err
	//}
	//blogAPI.Registry(v1.Group("/blog"))

	////注册Tag服务的API
	//tagAPI:=tag_api.NewHTTPAPI()
	//if err:=tagAPI.Init();err!=nil{
	//	return err
	//}
	//tagAPI.Registry(v1.Group("/tag"))
	apps.InitHttpService(v1)
	//打印下日志
	h.log.Infof("http server serve on: %s",h.server.Addr)

	if err:=h.server.ListenAndServe();err!=nil{
		//处理正常退出情况
		if err == http.ErrServerClosed{
			return nil
		}
		return fmt.Errorf("server ListenAndServe error,%s",err)
	}
	return nil
}

func (h *HTTP) Stop(ctx context.Context){
	h.log.Infof("server graceful shutdown ...")
	//HTTP Server 优雅关闭
	//支持ctx，10分钟 请求都没退出，做超时设置
	if err:=h.server.Shutdown(ctx);err!=nil{
		h.log.Warnf("shutdown error,%s",err)
	} else{
		h.log.Infof("server graceful shutdown ok")
	}
}

