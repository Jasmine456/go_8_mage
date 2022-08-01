package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func boy(c *gin.Context) { //你所需要的东西全部封装在了gin.Context里面，包括http.Request和ResponseWrite
	c.String(http.StatusOK, "hi boy") // 通过gin.Context.String返回一个 text/plain类型的正文
}

func girl(c *gin.Context) {
	c.String(http.StatusOK, "hi girl")
}

func main() {
	//gin.SetMode(gin.ReleaseMode) //发布模式，默认是Debug模式
	engine := gin.Default() //默认的engin已经自带了Logger和Recovery两个中间件
	engine.GET("/", boy)
	engine.POST("/", girl)

	// 路由分组
	oldVersion := engine.Group("/v1")
	oldVersion.GET("student", boy) //http://localhost:5656/v1/student
	oldVersion.GET("teacher", boy) //http://localhost:5656/v1/teacher

	newVersion := engine.Group("/v2")
	newVersion.GET("student", girl)
	newVersion.GET("teacher", girl)

	engine.Run(":5656")
}
