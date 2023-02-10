package middleware

import (
	"fmt"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/permission"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/token"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/client/rpc"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"net/http"
)

// 给服务端提供的Restful接口的 认证鉴权中间件
func NewHttpAudit() (*HttpAudit, error) {
	return &HttpAudit{
		log: zap.L().Named("audit.http"),
		// 从全局变量获取
		client: rpc.C(),
	}, nil
}

//// FilterFunction definitions must call ProcessFilter on the FilterChain to pass on the control and eventually call the RouteFunction
//type FilterFunction func(*Request, *Response, *FilterChain)

type HttpAudit struct {
	log logger.Logger
	//	基于rpc的客户端
	client *rpc.ClientSet
}

// 是否开启权限的控制，交给中间件使用方去决定
func (a *HttpAudit) GoRestfulHttpAutherFun(req *restful.Request, resp *restful.Response, next *restful.FilterChain) {

	//测试
	//var b=req.Request.Body
	//fmt.Printf("%v\n,%t\n,%T\n",b,b,b)

	if req.Request.Body == http.NoBody {
		response.Failed(resp.ResponseWriter, fmt.Errorf("body 为空"))
		return
	}

	// 请求拦截
	route := req.SelectedRoute()
	meta := route.Metadata()
	a.log.Debug("route meta: ", meta)

	isAuth, ok := meta[label.Auth]
	// 有认证标签，并且开启了认证
	if ok && isAuth.(bool) {
		//获取用户Token,Token放在Header Authorization
		ak := token.GetTokenFromHTTPHeader(req.Request)

		//2. 调用GRPC校验用户Token合法性
		tk, err := a.client.Token().ValidateToken(req.Request.Context(), token.NewValidateTokenRequest(ak))
		if err != nil {
			response.Failed(resp.ResponseWriter, err)
			return
		}

		//	是不是需要返回用户的认证信息：哪个人，在哪个空间下面， token本身的信息
		req.SetAttribute("tk", tk)

		//	判断用户权限
		isPerm, ok := meta[label.Permission]
		if ok && isPerm.(bool) {
			ci, err := a.client.ClientInfo(req.Request.Context())
			if err != nil {
				response.Failed(resp.ResponseWriter, err)
				return
			}

			//	调用鉴权接口
			check := permission.NewCheckPermissionRequest()
			check.Domain = tk.Domain
			check.Namespace = tk.Namespace
			check.Username = tk.Username
			check.ServiceId = ci.Id
			check.Path = fmt.Sprintf("%s.%s", route.Method(), route.Path())

			perm, err := a.client.Permission().CheckPermission(req.Request.Context(), check)
			if err != nil {
				response.Failed(resp.ResponseWriter, err)
				return
			}
			a.log.Debugf("permission check pass: %s", perm)
		}

	}
	// next flow
	next.ProcessFilter(req, resp)

}
