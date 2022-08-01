package controller

import (
	"github.com/gin-gonic/gin"
	"go_8_mage/week14/project_arch/mvc/dao"
)

func CreateBlog(ctx *gin.Context){
	//业务处理逻辑，也写在这里

	//调用Dao完成数据的入库
	dao.SaveBlog()
}

