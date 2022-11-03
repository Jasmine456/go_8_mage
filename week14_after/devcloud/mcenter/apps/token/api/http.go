package api

import (
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/token"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/book"
)

var (
	h = &handler{}
)

type handler struct {
	service token.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(book.AppName)
	h.service = app.GetGrpcApp(token.AppName).(token.Service)
	return nil
}

func (h *handler) Name() string {
	return token.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"令牌管理"}

	ws.Route(ws.POST("/").To(h.IssueToken).
		Doc("颁发令牌(登录)").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(token.IssueTokenRequest{}).
		Writes(response.NewData(token.Token{})))

	ws.Route(ws.DELETE("/").To(h.RevolkToken).
		Doc("撤销令牌(登出)").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PATCH("/").To(h.ChangeNamespace).
		Doc("切换空间").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(token.ChangeNamespaceRequest{}))

	ws.Route(ws.GET("/").To(h.ValidateToken).
		Doc("验证令牌").
		Metadata(restfulspec.KeyOpenAPITags,tags).
		Reads(token.ValidateTokenRequest{}).
		Writes(response.NewData(token.Token{})).
		Returns(200,"OK",token.Token{}))
}

func init() {
	app.RegistryRESTfulApp(h)
}
