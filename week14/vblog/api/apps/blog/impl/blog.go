package impl

import (
	"context"
	"go_8_mage/week14/vblog/api/apps/blog"
)

//创建文章
func (i *Impl)CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error){
	return nil,nil
}

//	更新文章
func (i *Impl)UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.Blog, error){
	return nil,nil
}

//	文章的删除
func (i *Impl)DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) (*blog.Blog, error){
	return nil,nil
}

//	文章列表
func (i *Impl)QueryBlog(ctx context.Context, req *blog.QueryBlogRequest) (*blog.BlogSet, error){
	return nil,nil
}

//	文章详情
func (i *Impl)DescribeBlog(ctx context.Context, req *blog.DescribeBlogRequest) (*blog.Blog, error){
	return nil,nil
}

//	发布接口
func (i *Impl)UpdateBlogStatus(ctx context.Context, req *blog.UpdateBlogStatusRequest) (*blog.Blog, error){
	return nil,nil
}


