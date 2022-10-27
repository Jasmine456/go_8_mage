package protocol

import (
	"context"
	"fmt"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/client/rpc/auth"
	"net/http"
	"time"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcube/app"

	"github.com/go_8_mage/week14_after/devcloud/maudit/conf"
	"github.com/go_8_mage/week14_after/devcloud/maudit/swagger"
)

// NewHTTPService 构建函数
func NewHTTPService() *HTTPService {

	r := restful.DefaultContainer
	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	// http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/emicklei/Projects/swagger-ui/dist"))))

	// Optionally, you may need to enable CORS for the UI to work.
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"*"},
		CookiesAllowed: false,
		Container:      r}
	r.Filter(cors.Filter)

	//补充一个认证中间件
	//1. 可以把该认真中间件写在maudit内部
	//2. 如果有100个微服务都需要接入到用户中心，都需要认证，每个模块独立维护一套 会带来很大维护成本，抽象成一个公共库
	//3. 公共库：怎么保证公开哭的可用性，服务端和客户端能力拆开，需要单独验证客户端的 能否正常
	//4. 服务端自己维护自己的中间件，服务端就是中间件的提供方 同时通过服务端
	am, err := auth.NewAuther("127.0.0.1:18050")
	if err != nil {
		panic(err)
	}
	r.Filter(am.GoRestfulAutherFun)

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.HTTP.Addr(),
		Handler:           r,
	}

	return &HTTPService{
		r:      r,
		server: server,
		l:      zap.L().Named("HTTP Service"),
		c:      conf.C(),
	}
}

// HTTPService http服务
type HTTPService struct {
	r      *restful.Container
	l      logger.Logger
	c      *conf.Config
	server *http.Server
}

func (s *HTTPService) PathPrefix() string {
	return fmt.Sprintf("/%s/api", s.c.App.Name)
}

// Start 启动服务
func (s *HTTPService) Start() error {
	// 装置子服务路由
	app.LoadRESTfulApp(s.PathPrefix(), s.r)

	// API Doc
	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: swagger.Docs}
	s.r.Add(restfulspec.NewOpenAPIService(config))
	s.l.Infof("Get the API using http://%s%s", s.c.App.HTTP.Addr(), config.APIPath)

	// 启动 HTTP服务
	s.l.Infof("HTTP服务启动成功, 监听地址: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service is stopped")
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil
}

// Stop 停止server
func (s *HTTPService) Stop() error {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Errorf("graceful shutdown timeout, force exit")
	}
	return nil
}

//注册功能列表
func (s *HTTPService) registryEndpoint() {
	wss := s.r.RegisteredWebServices()
	for i := range wss {
		routes := wss[i].Routes()
		eps := tools.TransferRoutesToEndpoints(routes)
		fmt.Println(eps)
		//for j:=range routes{
		//	r:=routes[j]
		//
		//	fmt.Println(r.Metadata,r.Path,r.Metadata,r.Operation)
		//}

	}
}
