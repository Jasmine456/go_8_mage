package blog


import "context"

type Service interface {
	//	为什么是这样
	//	接口的定义 是站在接口的使用方定义的
	//	接口一定要有很高的兼容性（不能随便接口的函数前面）
	CreateBlog(context.Context,*CreateBlogRequest) (*Blog,error)
}