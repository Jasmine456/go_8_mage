package main

import (
	"github.com/gin-gonic/gin"
	"go_8_mage/week14/project_arch/mvc/controller"
)

//工程入口
//完成 依赖模块的拼装
func main(){
	r:=gin.New()
	r.POST("/vblog/api/v1/blogs",controller.CreateBlog())
}
