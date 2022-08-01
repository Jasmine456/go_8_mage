package http

import (
	"github.com/gin-gonic/gin"
	"go_8_mage/week14/project_arch/ddd/blog"
)

func NewHandler(svr blog.Service) *Handler {
	return &Handler{
		service: svr,
	}
}

type Handler struct {
	//这里的实现，可以是具体实现，也可以是mock对象
	//也可以任意切换具体的实现 基于MongoDB/Mysql/es/etcd/hbase
	service blog.Service
}

//把该业务的handler注册给根路由
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/vblog/api/v1/blogs", h.CreateBlog)
}

func (h *Handler) CreateBlog(ctx *gin.Context) {
	h.service.CreateBlog(nil, nil)
}
