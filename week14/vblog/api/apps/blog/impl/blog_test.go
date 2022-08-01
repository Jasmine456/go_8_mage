package impl_test

import (
	"context"
	"go_8_mage/week14/vblog/api/apps/blog"
	"go_8_mage/week14/vblog/api/apps/blog/impl"
	"testing"
)


var blogService blog.Service

func TestCreateBlog(t *testing.T){
	req:=blog.NewCreateBlogRequest()
	ins,err:=blogService.CreateBlog(context.Background(),req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(ins)
}

func init(){
	blogService = impl.NewImpl()
}