package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type Blog struct {
	Title   string
	Content string
	Summary string
}

func main() {
	//额外的配置，监听的地址和端口，比如 localhost:80
	addr := os.Getenv("APP_ADDRESS")

	//	加载我们API Controller
	r := gin.New()
	r.POST("/vblog/api/v1/blogs", func(ctx *gin.Context) {
		// 1.获取用户的请求参数
		// 通过Query可以获取URL参数
		// http协议 query string: ?page_size=20&page_number=1
		// ctx.Request.URL.Query()
		// psStr := ctx.Query("page_size")
		// pnStr := ctx.Query("page_number")

		// 通过bind可以获取用户通过body传入的参数
		// 直接读取HTTP协议的body数据
		// ctx.Request.Body.Read()
		// 需要知道数据格式: 通过HTTP规范的 Content-Type Header
		//     "application/json" --> JSON binding
		//     "application/xml"  --> XML binding
		// 已经知道body里面的数据格式: Content-Type:application/json
		b := new(Blog)
		// json.Unmarshal(payload, b)
		// 反序列化  []bytes <---> obj
		// 后面所有的操作，基于obj完成
		if err := ctx.Bind(b); err != nil {
			// 返回给客户端
			ctx.JSON(400, map[string]string{"error": err.Error()})
			return
		}
		// 基于对象obj(b) 进行业务逻辑出来

		// 去数据库获取数据(DAO Data Access Object, 数据层) 数据库里面的数据(Row) ---> Object, ORM
		// SQL: Insert 语句, insert blog (title,content) VALUES (?, ?)

		// 需要返回给接口调用方,进行业务加工的(Controller, Bussise 业务层)
		// 获取Summary
		b.Summary = b.Content[:100]

		// 通过HTTP协议返回调用结果数据
		ctx.JSON(200, b)
	})

//	启动web框架
	if err:=r.Run(addr);err!=nil{
		log.Println(err.Error())
		os.Exit(1)
	}

}
