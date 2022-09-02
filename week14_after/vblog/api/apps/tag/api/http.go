package api

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go_8_mage/week14/vblog/api/apps"
	"go_8_mage/week14/vblog/api/apps/tag"
	"go_8_mage/week14/vblog/api/protocol/auth"
)

func NewHTTPAPI() *HTTPAPI {
	return &HTTPAPI{}
}

func (h *HTTPAPI) Name() string {
	return tag.AppName
}

// HTTPAPI 定义用来对外暴露HTTP服务，注册给协议Server(HTTP Server,Gin)
type HTTPAPI struct {
	service tag.Service
	log     logger.Logger
}

func (h *HTTPAPI) Init() error {
	h.service = apps.GetService(tag.AppName).(tag.Service)
	h.log = zap.L().Named("api.tag")
	return nil
}

// URI 注册给 Gin
func (h *HTTPAPI) Registry(r gin.IRouter) {
	r.Use(auth.BasicAuth)
	//管理员的接口，需要认证
	r.POST("/", h.AddTag)
	r.DELETE("/:id", h.RemoveTag)
	r.GET("/:blog_id", h.QueryBlogTag)
}

func init() {
	apps.RegistryHttp(&HTTPAPI{})
}
