package api

import (
	"github.com/gin-gonic/gin"
	"go_8_mage/week14/vblog/api/apps/blog"
	"net/http"
	"strconv"
)

//必须要符合Gin的Handler签名 func(*context)
// POST /vblog/api/v1/ HTTP Body:JSON DATA{"title_img": "xxx", "title_name":"xxx"}

////	获取用户请求
//payload, err := io.ReadAll(c.Request.Body)
//if err != nil {
//	//响应头的处理
//	c.Writer.Header().Set("X-Request-Id", "xxxx")
//	//返回 Status code
//	c.Writer.WriteHeader(http.StatusBadRequest)
//	//HTTP Response Body
//	c.Writer.Write([]byte(`ok`))
//}
//defer c.Request.Body.Close()
//

//请求的参数URL /vblog/api/v1/?keywords=hello
//c.Request.URL.Query().Get("keywords")
//路由定义 /vblog/api/v1/:xxx  /vblog/api/v1/abc  Gin路由库
//abc
//c.Param("xxx")

//	请求的参数还有在Header里面 Request Header
//	Header：X-Oauth-Token：xxxx
//	c.Request.Header.Get("X-Oauth-Token")

//obj := newxxxObject()
//json.Unmarshal(payload,obj)
func (h *HTTPAPI) CreateBlog(c *gin.Context) {
	//	Web框架（http协议框架）Gin 获取来自HTTP协议的用户请求
	//	HTTP：有一个头 Content-Type,表示Body里面的数据的格式 "application/json"
	//	Request Object <---> HTTP Body
	req:=blog.NewCreateBlogRequest()
	if err:=c.BindJSON(req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":http.StatusBadRequest,
			"message":err.Error(),
		})
		return
	}

	//获取用户
	username,_,_:=c.Request.BasicAuth()
	req.Author=username


	ins,err:=h.service.CreateBlog(c.Request.Context(),req)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,ins)
}

func (h *HTTPAPI) QueryBlog(c *gin.Context) {
	//Gin的逻辑怎么写
	//c.Request.URL.Query().Get("keywords")  page_size,page_number
	//req:=blog.NewQueryBlogRequest()
	//req.Keywords=c.Query("keywords")
	//req.PageSize= c.Query("page_size")
	//req.PageNumber= c.Query("page_number")
	//time.Sleep(3*time.Second)
	req:=blog.NewQueryBlogRequestFromHTTP(c.Request)
	set,err:=h.service.QueryBlog(c.Request.Context(),req)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	////补充Tag标签，为查询所有Blog都补充标签
	//for i:= range set.Items{
	//	req:=tag.NewQueryTagRequest()
	//	req.BlogId= set.Items[i].Id
	//	tset,err:= h.tag.QueryTag(c.Request.Context(),req)
	//	if err!=nil{
	//		c.JSON(http.StatusInternalServerError,gin.H{
	//			"code":http.StatusInternalServerError,
	//			"message":err.Error(),
	//		})
	//		return
	//	}
	//	set.Items[i].Tags=tset.Items
	//}

	c.JSON(http.StatusOK,set)
}

//文章详情
func (h *HTTPAPI) DescribeBlog(c *gin.Context) {
	//获取URL路径参数 路由定义 /vblog/api/v1/:xxx  /vblog/api/v1/abc  Gin路由库
	//"/:id" /abc id=abc
	blogIdStr := c.Param("id")
	bid,err:=strconv.Atoi(blogIdStr)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	req:=blog.NewDescribeBlogRequest(bid)
	ins,err:=h.service.DescribeBlog(c.Request.Context(),req)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,ins)
}

func (h *HTTPAPI) DeleteBlog(c *gin.Context) {
	//time.Sleep(4*time.Second)
	//"/:id" /abc id=abc
	blogIdStr := c.Param("id")
	bid,err:=strconv.Atoi(blogIdStr)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	req:=blog.NewDeleteBlogRequest(bid)
	ins,err:=h.service.DeleteBlog(c.Request.Context(),req)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,ins)
}

func (h *HTTPAPI) UpdateBlogStatus(c *gin.Context) {
	req:= blog.NewDefaultUpdateBlogStatusRequest()
	if err:=c.BindJSON(req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":http.StatusBadRequest,
			"message":err.Error(),
		})
		return
	}

	//"/:id" /abc id=abc
	// 覆盖掉来自body里面的id参数
	blogIdStr := c.Param("id")
	bid,err:=strconv.Atoi(blogIdStr)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	req.Id=bid
	ins,err:=h.service.UpdateBlogStatus(c.Request.Context(),req)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,ins)
}


func (h *HTTPAPI) PutBlog(c *gin.Context) {
	//"/:id" /abc id=abc
	// 覆盖掉来自body里面的id参数
	blogIdStr := c.Param("id")
	bid,err:=strconv.Atoi(blogIdStr)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	req:=blog.NewPutUpdateBlogRequest(bid)

	//读取来自body的json参数
	if err:=c.BindJSON(req.CreateBlogRequest);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	ins,err:=h.service.UpdateBlog(c.Request.Context(),req)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,ins)
}
func (h *HTTPAPI) PatchBlog(c *gin.Context) {
	//"/:id" /abc id=abc
	// 覆盖掉来自body里面的id参数
	blogIdStr := c.Param("id")
	bid,err:=strconv.Atoi(blogIdStr)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	req:=blog.NewPatchUpdateBlogRequest(bid)
	if err:=c.BindJSON(req.CreateBlogRequest);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	ins,err:=h.service.UpdateBlog(c.Request.Context(),req)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,ins)
}

