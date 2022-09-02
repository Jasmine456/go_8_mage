package api

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go_8_mage/week14/vblog/api/apps"
	"go_8_mage/week14/vblog/api/apps/blog"
	"go_8_mage/week14/vblog/api/apps/tag"
	"go_8_mage/week14/vblog/api/protocol/auth"
)

func NewHTTPAPI() *HTTPAPI {
	return &HTTPAPI{}
}

func (h *HTTPAPI)Name()string{
	return blog.AppName
}

// HTTPAPI 定义用来对外暴露HTTP服务，注册给协议Server(HTTP Server,Gin)
type HTTPAPI struct {
	//	业务逻辑：依赖后端Blog Service的实现
	service blog.Service
	tag tag.Service
	//	打印日志 log.Print,logrus,zap(Uber高性能）
	//	定义了Logger接口，可以容纳多种Log库的实现
	// 选择使用原生的logrus 或者zap 或者直接使用log
	// 替换成标准逻辑： https://pkg.go.dev/go.uber.org/zap
	log logger.Logger
}

//把HTTP Api 托管到IOC
//init时获取依赖注入
func (h *HTTPAPI) Init() error{
	h.service = apps.GetService(blog.AppName).(blog.Service)
	h.tag=apps.GetService(tag.AppName).(tag.Service)
	h.log=zap.L().Named("api.blog")
	return nil
}



// URI 注册给 Gin
func (h *HTTPAPI) Registry(r gin.IRouter) {
	//开发接口，不需要认证
	r.GET("/",h.QueryBlog)
	r.GET("/:id",h.DescribeBlog)

	r.Use(auth.BasicAuth)
	//管理员的接口，需要认证
	r.POST("/",h.CreateBlog)
	r.DELETE("/:id",h.DeleteBlog)
	r.PUT("/:id",h.PutBlog)
	r.PATCH("/:id",h.PatchBlog)
	r.POST("/:id/status",h.UpdateBlogStatus)

	h.log.Debugf("registry http handler complete")
}

func init(){
	apps.RegistryHttp(&HTTPAPI{})
}
