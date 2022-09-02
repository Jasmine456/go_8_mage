package apps

import "github.com/gin-gonic/gin"

//IOC容器,管理者所有的对象


//代表他是一个IOC的Service Object
var(
	//保存着所有的对象
	services = map[string]Service{}
	httpServices = map[string]HttpService{}
)

//代表他是一个ioc 的Service Object
type Service interface {
	Init() error
	Name() string
}

//作为一个实现了Gin HTTP Handler服务，提供一个路由注册功能
type HttpService interface {
	Service
	Registry(gin.IRouter)
}

func RegistryHttp(svc HttpService){
	httpServices[svc.Name()]=svc
}

//获取对象
func GetHttpService(name string) any{
	if v,ok:=httpServices[name];ok{
		return v
	}
	panic("http service"+name+"not registied")
}

//对象注册
func Registry(svc Service){
	services[svc.Name()]=svc


////	断言服务类型
//	if v,ok:=svc.(blog.Service);ok{
//		blogService=v
//	}
//	if v,ok:=svc.(tag.Service);ok{
//		tagService=v
//	}
}

//获取对象
func GetService(name string) any{
	if v,ok:=services[name];ok{
		return v
	}
	panic("service"+name+"not registied")
}


//初始化所有已经注册过来的实例
func Init() error{
	//初始化service
	for i:=range services{
		if err:=services[i].Init();err!=nil{
			return err
		}
	}

	return nil
}

func InitHttpService(rootRouter gin.IRouter)error{
	//初始化http service
	for i:=range httpServices{
		api:=httpServices[i]
		if err:=api.Init();err!=nil{
			return err
		}
		//blogAPI.Registry(v1.Group("/blog"))
		api.Registry(rootRouter.Group(api.Name()))
	}

	return nil
}