package blog

import "context"

const(
	AppName="blog"
)

type Service interface {
	//创建文章
	//ctx 接口的上下文，ctx用于区分业务功能数据和非业务功能数据
	//业务功能数据： 用户提交的数据，创建Blog的CreateBlogRequest
	//需要添加接口Trace，RequestId
	//非业务功能数据：RequestId，需要从上下文中传递
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)

	//	更新文章
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)

	//	文章的删除
	// 为什么删除后，还要返回数据，方便前端和事件总线使用
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)

	//	文章列表
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)

	//	文章详情
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)

	//	发布接口
	UpdateBlogStatus(context.Context, *UpdateBlogStatusRequest) (*Blog, error)
}
