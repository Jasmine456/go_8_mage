package impl_test

import (
	"context"
	"github.com/infraboard/mcube/logger/zap"
	"go_8_mage/week14/vblog/api/apps"
	"go_8_mage/week14/vblog/api/apps/blog"
	"go_8_mage/week14/vblog/api/apps/tag"
	"go_8_mage/week14/vblog/api/apps/tag/impl"
	"go_8_mage/week14/vblog/api/conf"
	"testing"

	blogImpl "go_8_mage/week14/vblog/api/apps/blog/impl"
	_ "go_8_mage/week14/vblog/api/apps/all"
)

//svc 实现 tag.Service
var svc tag.Service

func TestAddTag(t *testing.T){
	req:=tag.NewAddTagRequest()
	req.AddTag(&tag.CreateTagRequest{
		BlogId: 1,
		Key: "Language",
		Value: "Golang",
	})
	set,err:= svc.AddTag(context.Background(),req)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(set)
}

func TestQueryTag(t *testing.T){
	req:=tag.NewQueryTagRequest()
	req.AddTagId(1)
	set,err:=svc.QueryTag(context.Background(),req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestRemoveTag(t *testing.T){
	req:=tag.NewRemoveTagRequest()
	req.AddTagId(1)
	set,err:=svc.RemoveTag(context.Background(),req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}



func init(){
	//加载配置
	if err:=conf.LoadConfigFromEnv();err!=nil{
		panic(err)
	}

	zap.DevelopmentSetup()

	////blog service
	//blogIns:=blogImpl.NewImpl()
	//if err:=blogIns.Init();err!=nil{
	//	panic(err)
	//}
	////tag service
	//svc=impl.NewImpl(blogIns)
	//if err:= svc.Init();err!=nil{
	//	panic(err)
	//}

	//IOC容器实例对象初始化
	if err:=apps.Init();err!=nil{
		panic(err)
	}


	svc = apps.GetService(tag.AppName).(tag.Service)
}
