package api

import (
	"github.com/gin-gonic/gin"
	"go_8_mage/week14/vblog/api/apps/user"
	"net/http"
)

func (h *HTTPAPI) Auth(c *gin.Context){
	//用戶參數
	req:=user.NewAuthRequest()
	if err:=c.BindJSON(req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":http.StatusBadRequest,
			"message":err.Error(),
		})
		return
	}

	if req.Username == h.conf.Auth.Username && req.Password == h.conf.Auth.Password{
		c.JSON(http.StatusOK,gin.H{
			"code":0,
			"data":req.Username,
		})
		return
	}

	c.JSON(http.StatusUnauthorized,gin.H{
		"code":http.StatusUnauthorized,
		"message":"用户名或者密码不正确",
	})
	return
}
