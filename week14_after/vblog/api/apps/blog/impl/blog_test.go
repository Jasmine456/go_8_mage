package impl_test

import (
	"context"
	"github.com/infraboard/mcube/exception"
	"github.com/stretchr/testify/assert"
	"go_8_mage/week14/vblog/api/apps"
	"go_8_mage/week14/vblog/api/apps/blog"
	"go_8_mage/week14/vblog/api/apps/blog/impl"
	"go_8_mage/week14/vblog/api/conf"
	"testing"

	_ "go_8_mage/week14/vblog/api/apps/all"
)


var blogService blog.Service

func TestCreateBlog(t *testing.T){
	req:=blog.NewCreateBlogRequest()
	req.TitleName="vblog"
	req.Content="develop vblog system"
	ins,err:=blogService.CreateBlog(context.Background(),req)
	if err != nil {
		t.Fatalf(err.Error())
		err.(exception.APIException).ErrorCode()
	}
	t.Log(ins)
}

func TestImpl_QueryBlog(t *testing.T) {
	req:=blog.NewQueryBlogRequest()
	req.Keywords = "Txxx"
	ins,err:=blogService.QueryBlog(context.Background(),req)
	if err != nil {
		t.Fatalf(err.Error())
		err.(exception.APIException).ErrorCode()
	}
	t.Log(ins)
}

func TestDescribeBlog(t *testing.T) {
	req:=blog.NewDescribeBlogRequest(1)
	ins,err:=blogService.DescribeBlog(context.Background(),req)
	if err != nil {
		t.Fatalf(err.Error())
		err.(exception.APIException).ErrorCode()
	}
	t.Log(ins)
}

func TestDeleteBlog(t *testing.T) {
	req:=blog.NewDeleteBlogRequest(3)
	ins,err:=blogService.DeleteBlog(context.Background(),req)
	if err != nil {
		t.Fatalf(err.Error())
		err.(exception.APIException).ErrorCode()
	}
	t.Log(ins)
}

func TestUpdateBlog(t *testing.T) {
	req:=blog.NewPutUpdateBlogRequest(1)
	//req.TitleName="xxx"
	//req.Content="测试更新"

	ins,err:=blogService.UpdateBlog(context.Background(),req)
	if err != nil {
		t.Fatalf(err.Error())
		err.(exception.APIException).ErrorCode()
	}
	t.Log(ins)
}
func TestUpdateBlogPatch(t *testing.T) {
	req:=blog.NewPatchUpdateBlogRequest(2)
	req.TitleImg="patch 局部更新"


	ins,err:=blogService.UpdateBlog(context.Background(),req)
	if err != nil {
		t.Fatalf(err.Error())
		err.(exception.APIException).ErrorCode()
	}
	t.Log(ins)
}

func TestUpdateBlogStatus(t *testing.T) {
	req:=blog.NewUpdateBlogStatusRequest(1,blog.STATUS_PUBLISHD)
	ins,err:=blogService.UpdateBlogStatus(context.Background(),req)
	if err != nil {
		t.Fatalf(err.Error())
	}

	//用assert 断言库来做 两个值的比较
	should:=assert.New(t)
	//if ins.Status!=blog.STATUS_PUBLISHD
	if should.NotEqual(blog.STATUS_PUBLISHD,ins.Status){
		t.Fatal(err.Error())
	}
	t.Log(ins)
}



func init(){
	//加载配置
	if err:=conf.LoadConfigFromEnv();err!=nil{
		panic(err)
	}

	//IOC容器实例对象初始化
	if err:=apps.Init();err!=nil{
		panic(err)
	}


	blogService = apps.GetService(blog.AppName).(blog.Service)
}