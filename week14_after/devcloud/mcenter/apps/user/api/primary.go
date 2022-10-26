package api

import (
	"fmt"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/user"
)

var (
	primaryHandler = &primary{}
)

//子账号管理
type primary struct {
	service user.Service
	log     logger.Logger
}

func (h *primary) Config() error {
	h.log = zap.L().Named(user.AppName)
	h.service = app.GetGrpcApp(user.AppName).(user.Service)
	return nil
}

func (h *primary) Name() string {
	return "user/primary"
}

func (h *primary) Version() string {
	return "v1"
}


// 需要提供的Restful接口
// 哪些接口 是和哪些类型的用户使用
// 只能主账号采用调用
func (h *primary) Registry(ws *restful.WebService) {
	tags := []string{"子账号管理"}
	ws.Route(ws.POST("").To(h.CreateUser).
		Doc("创建子账号(主账号)").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}).
		Writes(user.User{}))

	ws.Route(ws.GET("/").To(h.QueryUser).
		Metadata(label.Resource, "子账号").
		Metadata(label.Auth, true).
		Metadata(label.Allow, user.TYPE_PRIMARY).
		Doc("查询子账号列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}).
		Writes(response.NewData(user.User{})))

	ws.Route(ws.GET("/{id}").To(h.DescribeUser).
		Metadata(label.Resource, "子账号").
		Metadata(label.Auth, true).
		Doc("查询子用户详细").
		Param(ws.PathParameter("id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(response.NewData(user.User{})).
		Returns(200, "OK", response.NewData(user.User{})).
		Returns(404, "Not Found", nil))

}

// 创建子账号(主账号)
func (h *primary) CreateUser(r *restful.Request, w *restful.Response) {
	// 通过HTTP 获取用户请求数据
	req := user.NewCreateUserRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	req.Type = user.TYPE_SUB
	req.CreateBy = user.CREATE_BY_ADMIN

	set, err := h.service.CreateUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, set)
}

// 查询子账号
func (h *primary) QueryUser(r *restful.Request, w *restful.Response) {
	req := user.NewQueryUserRequestFromHTTP(r.Request)

	// 控制参数范围
	req.WithType(user.TYPE_SUB)
	ins, err := h.service.QueryUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, ins)
}

func (h *primary) DescribeUser(r *restful.Request, w *restful.Response) {
	// url path 获取变量id  {id}
	req := user.NewDescriptUserRequestWithId(r.PathParameter("id"))
	ins, err := h.service.DescribeUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	if !ins.Spec.Type.Equal(user.TYPE_SUB) {
		response.Failed(w.ResponseWriter, fmt.Errorf("not an sub account"))
		return
	}

	// restful返回的数据, 需要脱敏的，但是rpc访问的数据 不需要脱敏
	ins.Desensitize()

	response.Success(w.ResponseWriter, ins)
}
