# 支持GRPC的脚手架

## 项目初始化

安装脚手架工具:
```
go install github.com/infraboard/mcube/cmd/mcube@latest
```

使用脚手架生成框架代码
```sh
$ mcube.exe project init 
? 请输入项目包名称: gitee.com/go-course/go8/micro/tools/demo
? 是否接入权限中心[keyauth] No
? 选择数据库类型: MongoDB
? MongoDB服务地址,多个地址使用逗号分隔: (127.0.0.1:27017)

? MongoDB服务地址,多个地址使用逗号分隔: 127.0.0.1:27017
? 认证数据库名称: demo
? 认证用户: demo
? 生成样例代码 Yes
? 选择HTTP框架: go-restful
项目初始化完成, 项目结构如下: 
├───.gitignore (287b)
├───.mcube.yaml (209b)        
├───.vscode
│       └───settings.json (233b)
├───README.md (4207b)
├───apps
│       ├───all
│       │       ├───api.go (152b)
│       │       ├───impl.go (183b)
│       │       └───internal.go (107b)
│       └───book
│               ├───api
│               │       ├───book.go (2264b)
│               │       └───http.go (2254b)
│               ├───app.go (2213b)
│               ├───impl
│               │       ├───book.go (1678b)
│               │       ├───dao.go (3075b)
│               │       └───impl.go (843b)
│               └───pb
│                       └───book.proto (2364b)
├───client
│       ├───client.go (991b)
│       ├───client_test.go (656b)
│       └───config.go (164b)
├───cmd
│       ├───init.go (465b)
│       ├───root.go (1261b)
│       └───start.go (3795b)
├───conf
│       ├───config.go (3654b)
│       ├───load.go (720b)
│       └───log.go (365b)
├───docs
│       ├───README.md (15b)
│       └───schema
│               └───tables.sql (849b)
├───etc
│       ├───config.env (470b)
│       ├───config.toml (311b)
│       └───unit_test.env (17b)
├───go.mod (47b)
├───main.go (104b)
├───makefile (2973b)
├───protocol
│       ├───grpc.go (1348b)
│       └───http.go (2935b)
├───swagger
│       └───docs.go (725b)
└───version
        └───version.go (628b)
```

下载依赖
```
go tidy
```

## Copy依赖的Protobuf文件

```
# 使用 pb指令来完成protobuf文件copy, 但是windows由于路径文件不行
$ make pb
mkdir: created directory 'common'
mkdir: created directory 'common/pb'
mkdir: created directory 'common/pb/github.com'
mkdir: created directory 'common/pb/github.com/infraboard'
mkdir: created directory 'common/pb/github.com/infraboard/mcube'
mkdir: created directory 'common/pb/github.com/infraboard/mcube/pb'
cp: cannot stat 'E:Golang/pkg/mod/github.com/infraboard/mcube@v1.9.0/pb/*': No such file or directory
make: *** [makefile:64: pb] Error 1
```

需要对protobuf做版本管理: mcube这个库没有石油版本 git tag, 比如当前mcube 1.9, copy 1.9版本的 protobuf文件，

怎么获取对应版本的protobuf文件
```sh
# 获取当前项目依赖cmube的版本
$ go list -m "github.com/infraboard/mcube" | cut -d' ' -f2
v1.9.0

# 找到该版本的源码文件, 源码文件里面有protobuf文件, 也就是对应1.9版本的protobuf文件
$ ls /d/go_work/pkg/mod/github.com/infraboard | grep mcube
mcube@v1.0.3
mcube@v1.0.4
mcube@v1.0.5
mcube@v1.1.0
# 找到了对应版本的protobuf文件
$ ls /d/go_work/pkg/mod/github.com/infraboard/mcube\@v1.9.0/pb/

# 把这些文件copy到当前项目来
$ cp -r /d/go_work/pkg/mod/github.com/infraboard/mcube\@v1.9.0/pb/* common/pb/github.com/infraboard/mcube/pb

# 清除.go的文件，只要protobuf源文件
rm -rf common/pb/github.com/infraboard/mcube/pb/*/*.go

# 查看protobu是否copy过来
$ ls common/pb/github.com/infraboard/mcube/pb/
example  http  page  request  resource  response
```


## 编译 protobuf


make install: Install depence go package
```
# protoc-gen-go
@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# protoc-gen-go-grpc
@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# protoc-go-inject-tag 注入tag的插件
# 如果你要为你的生成后的结构体补充自定义标签: protoc-go-inject-tag
# 在字段上方加入注解: // @gotags: json:"create_at" bson:"create_at"
# protoc-go-inject-tag扫描注解，修改生成的struct 的tag
@go install github.com/favadi/protoc-go-inject-tag@latest
```

?, 工程里面生成一个样例代码, 这样样例代码是有protobuf文件, 为了样例跑起来，需要编译protobuf文件

make gen
```
    # 编译protobuf文件, apps/<模块名称>/pb/*.proto
	@protoc -I=. -I=common/pb --go_out=. --go_opt=module=${PKG} --go-grpc_out=. --go-grpc_opt=module=${PKG} apps/*/pb/*.proto
	@go fmt ./...

    # protobuf的tag注入库: protoc-go-inject-tag, 通过执行下面的插件来二次处理protobuf, 完成标签的重写
	@protoc-go-inject-tag -input=apps/*/*.pb.go
    # 为emnu类型 添加string描述
	@mcube generate enum -p -m apps/*/*.pb.go
```

## 安装数据库

```
docker pull mongo
docker run -itd -p 27017:27017 mongo
```

进入mongo的本地shell 添加用户: demo用户
```sh
$ docker exec -it 7913dd518691  mongo
> use demo
switched to db demo
> db.createUser({user: "demo", pwd: "123456", roles: [{ role: "dbOwner", db: "demo" }]})
Successfully added user: {
        "user" : "demo",
        "roles" : [
                {
                        "role" : "dbOwner",
                        "db" : "demo"
                }
        ]
}
```

确认配置文件里面的mongo配置: etc/config.toml
```toml
[mongodb]
endpoints = ["127.0.0.1:27017"]
username = "demo"
password = "123456"
database = "demo"
```


## 启动服务

```sh
$ make run
2022-10-01T17:33:51.182+0800    INFO    [INIT]  cmd/start.go:154        log level: debug
2022-10-01T17:33:51.215+0800    INFO    [CLI]   cmd/start.go:92 loaded grpc app: [book]
2022-10-01T17:33:51.216+0800    INFO    [CLI]   cmd/start.go:93 loaded http app: [book]
2022-10-01T17:33:51.216+0800    INFO    [CLI]   cmd/start.go:95 loaded internal app: []
2022-10-01T17:33:51.219+0800    INFO    [GRPC Service]  protocol/grpc.go:54     GRPC 服务监听地址: 127.0.0.1:18050
2022-10-01T17:33:51.219+0800    INFO    [HTTP Service]  protocol/http.go:78     Get the API using http://127.0.0.1:8050/apidocs.json
2022-10-01T17:33:51.220+0800    INFO    [HTTP Service]  protocol/http.go:81     HTTP服务启动成功, 监听地址: 127.0.0.1:8050
```

## 查看API文档

https://petstore.swagger.io/  输入： http://127.0.0.1:8050/apidocs.json



## 新技术

+ web框架切换成了 go-restful, 专门用于开发restful API的框架, k8s API server基于此框架开发的
+ mysql ---> mongodb 存储的数据 可能会出现非结构化的, 比如文件

### MongoDB

本身有点像ORM, 有一个 bson标签, 存入数据库时候 对应的bson的一个key
```go
type CreateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 创建人
	// @gotags: json:"create_by" bson:"create_by"
	CreateBy string `protobuf:"bytes,1,opt,name=create_by,json=createBy,proto3" json:"create_by" bson:"create_by"`
	// 名称
	// @gotags: json:"name" bson:"name" validate:"required"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" bson:"name" validate:"required"`
	// 作者
	// @gotags: json:"author" bson:"author" validate:"required"
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author" bson:"author" validate:"required"`
}
```

### GoRestful


```go
// RouteBuilder is a helper to construct Routes.
type RouteBuilder struct {
	rootPath                         string
	currentPath                      string
	produces                         []string
	consumes                         []string
	httpMethod                       string        // required
	function                         RouteFunction // required
	filters                          []FilterFunction
	conditions                       []RouteSelectionConditionFunction
	allowedMethodsWithoutContentType []string // see Route

	typeNameHandleFunc TypeNameHandleFunction // required

	// documentation, swagger集成时的文档信息
	doc                     string
	notes                   string
	operation               string
	readSample, writeSample interface{}
	parameters              []*Parameter
	errorMap                map[int]ResponseError
	defaultResponse         *ResponseError
	metadata                map[string]interface{}
	extensions              map[string]interface{}
	deprecated              bool
	contentEncodingEnabled  *bool
}
```


```go
func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"books"}

	ws.Route(ws.POST("").To(h.CreateBook)
		// Doc("create a book").
		// Metadata(restfulspec.KeyOpenAPITags, tags).
		// Reads(book.CreateBookRequest{}).
		// Writes(response.NewData(book.Book{})))

	ws.Route(ws.GET("/").To(h.QueryBook)
		// Doc("get all books").
		// Metadata(restfulspec.KeyOpenAPITags, tags).
		// Metadata("action", "list").
		// Reads(book.CreateBookRequest{}).
		// Writes(response.NewData(book.BookSet{})).
		// Returns(200, "OK", book.BookSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeBook)
		// Doc("get a book").
		// Param(ws.PathParameter("id", "identifier of the book").DataType("integer").DefaultValue("1")).
		// Metadata(restfulspec.KeyOpenAPITags, tags).
		// Writes(response.NewData(book.Book{})).
		// Returns(200, "OK", response.NewData(book.Book{})).
		// Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.UpdateBook)
		// Doc("update a book").
		// Param(ws.PathParameter("id", "identifier of the book").DataType("string")).
		// Metadata(restfulspec.KeyOpenAPITags, tags).
		// Reads(book.CreateBookRequest{}))

	ws.Route(ws.PATCH("/{id}").To(h.PatchBook)
		// Doc("patch a book").
		// Param(ws.PathParameter("id", "identifier of the book").DataType("string")).
		// Metadata(restfulspec.KeyOpenAPITags, tags).
		// Reads(book.CreateBookRequest{}))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteBook)
		// Doc("delete a book").
		// Metadata(restfulspec.KeyOpenAPITags, tags).
		// Param(ws.PathParameter("id", "identifier of the book").DataType("string")))
}
```
