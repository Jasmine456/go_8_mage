package api

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go_8_mage/week14/vblog/api/apps"
	"go_8_mage/week14/vblog/api/apps/user"
	"go_8_mage/week14/vblog/api/conf"
)

func NewHTTPAPI() *HTTPAPI {
	return &HTTPAPI{
		log:zap.L().Named("api.user"),
	}
}

func (h *HTTPAPI)Name()string{
	return user.AppName
}

// HTTPAPI 定义用来对外暴露HTTP服务，注册给协议Server(HTTP Server,Gin)
type HTTPAPI struct {
	log logger.Logger
	conf *conf.Config
}

//把HTTP Api 托管到IOC
//init时获取依赖注入
func (h *HTTPAPI) Init() error{
	h.log=zap.L().Named("api.user")
	h.conf = conf.C()
	return nil
}



// URI 注册给 Gin
func (h *HTTPAPI) Registry(r gin.IRouter) {
	r.POST("/auth",h.Auth)
}

func init(){
	apps.RegistryHttp(&HTTPAPI{})
}
