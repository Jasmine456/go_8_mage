package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//从GET请求的URL中获取参数
func url(engine *gin.Engine) {
	engine.GET("student", func(ctx *gin.Context) {
		name := ctx.Query("name")
		addr := ctx.DefaultQuery("addr", "China") //如果没传addr参数，则默认为China
		ctx.String(http.StatusOK, name+"\tlive in\t"+addr)
	})
}

//从Restful风格的url中获取参数
func restful(engine *gin.Engine) {
	engine.GET("/student/:name/*addr", func(ctx *gin.Context) {
		name := ctx.Param("name")
		addr := ctx.Param("addr")
		ctx.String(http.StatusOK, name+"live in"+addr)
	})
}

// 从post表单中获取参数
func post(engine *gin.Engine) {
	engine.POST("/student", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		addr := ctx.DefaultPostForm("addr", "China")
		ctx.String(http.StatusOK, name+"\tlive in\t"+addr)
	})
}

//上传单个文件
func upload_file(engine *gin.Engine) {
	//	 限制表单上传大小为8M，默认上限是32M
	engine.MaxMultipartMemory = 8 << 20
	engine.POST("/upload",func(ctx *gin.Context){
		file,err:=ctx.FormFile("file")
		if err!=nil{
			fmt.Printf("get file error %v\n",err)
			ctx.String(http.StatusInternalServerError,"upload file failed")
		} else{
			ctx.SaveUploadedFile(file,"./data/"+file.Filename)// 把用户上传的文件存到data目录下
			ctx.String(http.StatusOK,file.Filename)
		}
	})
}


//上传多个文件
func upload_multi_file(engin *gin.Engine){
	engin.POST("/upload_files",func(ctx *gin.Context){
		form,err:=ctx.MultipartForm() //MultipartForm 中不只包含多个文件
		if err !=nil{
			ctx.String(http.StatusBadRequest,err.Error())
		} else {
			// 从MultipartForm中获取上传的文件
			files:=form.File["files"]
			for _,file:=range files{
				ctx.SaveUploadedFile(file,"./data/"+file.Filename) // 把用户上传的文件存到data目录下
			}
			ctx.String(http.StatusOK,"upload"+strconv.Itoa(len(files))+"files")
		}
	})
}

type Student struct {
	Name string
	Addr string
}


func main(){
	engine:=gin.Default()

	url(engine) //http://localhost:5656/student?name=jasmine&addr=bj
	restful(engine)//http://localhost:5656/student?name=jasmine&addr=bj
	post(engine)
	upload_file(engine)
	upload_multi_file(engine)


	engine.Run(":5656")
}