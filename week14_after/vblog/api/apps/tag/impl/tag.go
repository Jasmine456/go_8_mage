package impl

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/exception"
	"go_8_mage/week14/vblog/api/apps/blog"
	"go_8_mage/week14/vblog/api/apps/tag"
)

//文章添加Tag
func (i *Impl) AddTag(ctx context.Context,req *tag.AddTagRequest)(*tag.TagSet,error){

	//参数校验<只校验的参数的有无>
	if err:= req.Validate();err!=nil{
		return nil,exception.NewBadRequest("validate AddTagRequest error,%s",err)
	}

	//校验对象是否存在<blog_id>对象的blog是否真的存在
	//性能的问题的写法
	bids:=req.BlogIds()
	i.log.Debugf("query blogs: %v",bids)
	for _,bid:=range bids{
		// 20 bid 是不是就要和数据库交互 20次
		_,err:=i.blog.DescribeBlog(ctx,blog.NewDescribeBlogRequest(bid))
		if err != nil {
			return nil,err
		}
	}


	//构造对象
	set := tag.NewTagSetFromAddTagRequest(req)

	//保存入库,只执行Insert操作
	if err:= i.DB().Create(set.Items).Error;err!=nil{
		return nil,err
	}

	return set,nil
}

//根据TagId查询 Tag的列表
func (i *Impl) QueryTag(ctx context.Context,req *tag.QueryTagRequest)(*tag.TagSet,error){
	set:=tag.NewTagSet()

	query:=i.DB()

	if req.BlogId!=0{
		query = query.Where("blog_id=?",req.BlogId)
	}
	if len(req.TagIds) >0{
		//SELECT * FROM tag WHERE id IN (1,2);
		query=i.DB().Where("id IN ?",req.TagIds)
	}




	if err:=query.Scan(&set.Items).Error;err!=nil{
		return nil,err
	}

	return set,nil
}


//文章移除Tag
func (i *Impl)  RemoveTag(ctx context.Context,req *tag.RemoveTagRequest)(*tag.TagSet,error){
	queryReq:=tag.NewQueryTagRequest()
	queryReq.AddTagId(req.TagIds...)
	ts,err:=i.QueryTag(ctx,queryReq)
	if err != nil {
		return nil,err
	}
	if len(ts.Items) !=len(req.TagIds){
		return nil,fmt.Errorf("has tag not exits")
	}

	//删除这些tag
	//i.DB().Where("id IN ?",req.TagIds).Delete(ts)
	if err:=i.DB().Delete(ts.Items).Error;err!=nil{
		return nil, err
	}


	return ts,nil
}
