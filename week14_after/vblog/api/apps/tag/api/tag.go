package api

import (
	"github.com/gin-gonic/gin"
	"go_8_mage/week14/vblog/api/apps/tag"
	"net/http"
	"strconv"
)


//文章添加Tag
func (h *HTTPAPI) AddTag(c *gin.Context){
	//	接收请求参数
	req:=tag.NewAddTagRequest()
	if err:= c.BindJSON(&req.Tags);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":http.StatusBadRequest,
			"message":err.Error(),
		})
		return
	}

	set,err:=h.service.AddTag(c.Request.Context(),req)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,set)
}

//文章移除Tag
func (h *HTTPAPI) RemoveTag(c *gin.Context){
	//	接收请求参数
	req:=tag.NewRemoveTagRequest()
	bint,err:=strconv.Atoi(c.Param("id"))
	//if err:= c.BindJSON(&req.TagIds);err!=nil{
	//	c.JSON(http.StatusBadRequest,gin.H{
	//		"code":http.StatusBadRequest,
	//		"message":err.Error(),
	//	})
	//	return
	//}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	req.TagIds = []int{bint}

	set,err:=h.service.RemoveTag(c.Request.Context(),req)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,set)
}

//查询标签
func (h *HTTPAPI) QueryBlogTag(c *gin.Context){
	//blog_id 怎么传递  1. url路径 2. query string
	req:=tag.NewQueryTagRequest()
	bint,err:=strconv.Atoi(c.Param("blog_id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":http.StatusBadRequest,
			"message":err.Error(),
		})
		return
	}
	req.BlogId=bint
	set,err:= h.service.QueryTag(c.Request.Context(),req)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,set)
}