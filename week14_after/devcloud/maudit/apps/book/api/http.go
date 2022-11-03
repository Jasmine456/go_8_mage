package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/go_8_mage/week14_after/devcloud/maudit/apps/book"
)

var (
	h = &handler{}
)

type handler struct {
	service book.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(book.AppName)
	h.service = app.GetGrpcApp(book.AppName).(book.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return book.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"books"}
	ws.Route(ws.POST("").To(h.CreateBook).
		Doc("create a book").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 开启认证
		Metadata(label.Auth, true).
		// 开启鉴权
		Metadata(label.Permission, true).
		// 该接口操作的资源名称
		Metadata(label.Resource, h.Name()).
		// 标注该资源的具体动态
		Metadata(label.Action, label.Create.Value()).
		//

		Reads(book.CreateBookRequest{}).
		Writes(response.NewData(book.Book{})))

	ws.Route(ws.GET("/").To(h.QueryBook).
		Doc("get all books").
		//开启认证
		Metadata(label.Auth, true).
		//开启鉴权
		Metadata(label.Permission, true).

		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 该接口操作的资源名称
		Metadata(label.Resource, h.Name()).
		// 标注该资源的具体动态
		Metadata(label.Action, label.List.Value()).
		Reads(book.CreateBookRequest{}).
		Writes(response.NewData(book.BookSet{})).
		Returns(200, "OK", book.BookSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeBook).
		Doc("get a book").
		//开启认证
		Metadata(label.Auth, true).
		//开启鉴权
		Metadata(label.Permission, true).
		Param(ws.PathParameter("id", "identifier of the book").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 该接口操作的资源名称
		Metadata(label.Resource, h.Name()).
		// 标注该资源的具体动态
		Metadata(label.Action, label.Get.Value()).
		Writes(response.NewData(book.Book{})).
		Returns(200, "OK", response.NewData(book.Book{})).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.UpdateBook).
		Doc("update a book").
		//开启认证
		Metadata(label.Auth, true).
		//开启鉴权
		Metadata(label.Permission, true).
		Param(ws.PathParameter("id", "identifier of the book").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 该接口操作的资源名称
		Metadata(label.Resource, h.Name()).
		// 标注该资源的具体动态
		Metadata(label.Action, label.Update.Value()).
		Reads(book.CreateBookRequest{}))

	ws.Route(ws.PATCH("/{id}").To(h.PatchBook).
		Doc("patch a book").
		//开启认证
		Metadata(label.Auth, true).
		//开启鉴权
		Metadata(label.Permission, true).
		Param(ws.PathParameter("id", "identifier of the book").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 该接口操作的资源名称
		Metadata(label.Resource, h.Name()).
		// 标注该资源的具体动态
		Metadata(label.Action, label.Update.Value()).
		Reads(book.CreateBookRequest{}))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteBook).
		Doc("delete a book").
		//开启认证
		Metadata(label.Auth, true).
		//开启鉴权
		Metadata(label.Permission, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 该接口操作的资源名称
		Metadata(label.Resource, h.Name()).
		// 标注该资源的具体动态
		Metadata(label.Action, label.Delete.Value()).
		Param(ws.PathParameter("id", "identifier of the book").DataType("string")))
}

func init() {
	app.RegistryRESTfulApp(h)
}
