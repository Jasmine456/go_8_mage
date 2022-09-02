package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger/zap"
	"go_8_mage/week14/vblog/api/conf"
	"net/http"
)

func BasicAuth(c *gin.Context) {
	//	处理请求：检查BasicAuth
	// username:password
	zap.L().Debugf("basic auth:%s",c.Request.Header.Get("Authorization"))

	//	1.获取用户的用户密码
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "username or password not conrect",
		})
		return
	}
	//	和系统配置的用户密码进行比对
	zap.L().Debugf("auth user:%s", username)
	ac := conf.C().Auth
	if !(username == ac.Username && password == ac.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "username or password not conrect",
		})
		return
	}

	//	处理响应：无
	//	继续路由到后面的Handler
	c.Next()
}
