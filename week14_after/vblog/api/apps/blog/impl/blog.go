package impl

import (
	"context"
	"fmt"
	"github.com/imdario/mergo"
	"github.com/infraboard/mcube/exception"
	"go_8_mage/week14/vblog/api/apps/blog"
	"go_8_mage/week14/vblog/api/apps/tag"
	"time"
)

//创建文章
func (i *Impl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {
	//校验对象的合法性
	if err := req.Validate(); err != nil {
		//工程是否需要定义统一异常，不定义 都可以字符串的形式返回

		return nil, exception.NewBadRequest("validate create blog request error,%s", err)
	}
	//创建对象
	ins := blog.NewCreateBlog(req)

	//把对象入库，id 是自增主键，是由数据库计算
	//LastId insert
	//ins.Id = LastId()
	fmt.Println(ins)
	if err := i.save(ctx, ins); err != nil {
		return nil, err
	}
	return ins, nil
}

//	更新文章
func (i *Impl) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.Blog, error) {

	//根据blog id 获取blog对象
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(req.BlogId))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case blog.UPDATE_MODE_PUT:
		//	对象的全量更新
		ins.CreateBlogRequest = req.CreateBlogRequest
	case blog.UPDATE_MODE_PATCH:
		//	局部更新,有很多要更新的字段
		//ins.CreateBlogRequest.TitleName = req.TitleName

		//	有没有能进行 object merge lib
		// PATCH old <---- new old
		if err:=mergo.MergeWithOverwrite(ins.CreateBlogRequest,req.CreateBlogRequest);err!=nil{
			return nil,err
		}
	default:
		return nil, exception.NewBadRequest("update mode note support %s", req.UpdateMode)
	}
	//实例对象更新完成
	//检查update参数的合法性，需要入库之前检查，因为有Patch操作
	if err:=ins.CreateBlogRequest.Validate();err!=nil{
		return nil,exception.NewBadRequest("validate request error,%s",err)
	}

	//数据库更新
	if err:=i.DB().WithContext(ctx).Updates(ins).Error;err!=nil{
		return nil,err
	}

	return ins,err
}

//	文章的删除
func (i *Impl) DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) (*blog.Blog, error) {

	//delete的时候，需要返回blog对象，先查询blog
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(req.Id))
	if err != nil {
		return nil, err
	}

	//query := i.DB().Where("id=?",req.Id)
	if err := i.DB().Delete(ins); err != nil {
		return nil, err.Error
	}
	return ins, nil

	return nil, nil
}

//	文章列表
func (i *Impl) QueryBlog(ctx context.Context, req *blog.QueryBlogRequest) (*blog.BlogSet, error) {

	set := blog.NewBlogSet()
	query := i.DB()

	//条件拼接
	//支持关键字查询，.*
	if req.Keywords != "" {
		query = query.Where(
			"title_name LIKE ? OR content LIKE ?",
			"%"+req.Keywords+"%",
			"%"+req.Keywords+"%",
			)
	}

	//如何查询总条数：COUNT(*)
	if err := query.Count(&set.Total).Error; err != nil {
		return nil, err
	}

	//SQL 如何支持分页 LIMIT <offset>,<limit> 并且倒序排列，让最新创建的条目 放置在最前面
	query = query.Offset(req.Offset()).Limit(req.PageSize).Order("create_at DESC")

	//需要查询分页数据
	if err := query.WithContext(ctx).Scan(&set.Items).Error; err != nil {
		return nil, err
	}

	//补充Tag标签，为查询所有Blog都补充标签
	for index:= range set.Items{
		req:=tag.NewQueryTagRequest()
		req.BlogId= set.Items[index].Id
		tset,err:= i.tag.QueryTag(ctx,req)
		if err!=nil{
			return nil, err
		}
		set.Items[index].Tags=tset.Items
	}

	return set, nil
}

//	文章详情
func (i *Impl) DescribeBlog(ctx context.Context, req *blog.DescribeBlogRequest) (*blog.Blog, error) {
	//默认blog示例
	ins := blog.NewCreateBlog(blog.NewCreateBlogRequest())

	query := i.DB().Where("id=?", req.Id)
	if err := query.Find(ins).Error; err != nil {
		return nil, err
	}

	//注意处理404
	if ins.Id == 0 {
		return nil, exception.NewNotFound("blog %d not found", req.Id)
	}

	return ins, nil
}

//	发布接口
func (i *Impl) UpdateBlogStatus(ctx context.Context, req *blog.UpdateBlogStatusRequest) (*blog.Blog, error) {

	//需要判断该对象是否存在，SQL EXIST
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(req.Id))
	if err != nil {
		return nil, err
	}

	//修改状态
	if ins.Status == req.Status{
		return nil,exception.NewBadRequest("status aready %s",req.Status)
	}
	ins.Status = req.Status
	if ins.Status == blog.STATUS_PUBLISHD{
		ins.PublishAt = time.Now().Unix()
	}

	//入库保存
	//i.DB().WithContext(ctx).UpdateColumns("")
	// UPDATE `blog` SET `create_at`=1659237545,`publish_at`=1659749715,`title_img`='更新URL',`title_name`='xxxx',`content`='测试更新',`status`=1 WHERE `id` = 1
	if err:= i.DB().WithContext(ctx).Save(ins).Error;err!=nil{
		return nil,err
	}
	return ins, nil
}
