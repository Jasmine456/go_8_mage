package auth

import (
	"fmt"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/token"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/user"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

//// FilterFunction definitions must call ProcessFilter on the FilterChain to pass on the control and eventually call the RouteFunction
//type FilterFunction func(*Request, *Response, *FilterChain)

func NewAuther() *Auther {
	return &Auther{
		log: zap.L().Named("auther.http"),
		tk:  app.GetInternalApp(token.AppName).(token.Service),
	}
}

type Auther struct {
	log logger.Logger
	tk  token.Service
}

func (a *Auther) GoRestfulAutherFun(req *restful.Request, resp *restful.Response, next *restful.FilterChain) {
	// 请求拦截
	meta := req.SelectedRoute().Metadata()
	a.log.Debug("route meta: ", meta)

	isAuth, ok := meta[label.Auth]
	//有认证标签，并且开启了认证
	if ok && isAuth.(bool) {
		//	获取token
		ak := token.GetTokenFromHTTPHeader(req.Request)

		tk, err := a.tk.ValidateToken(req.Request.Context(), token.NewValidateTokenRequest(ak))
		if err != nil {
			response.Failed(resp.ResponseWriter, err)
			return
		}

		v, ok := meta[label.Allow]
		if ok {
			ut := v.(user.TYPE)
			// 权限的编号来判断
			if tk.UserType < ut {
				response.Failed(resp.ResponseWriter, fmt.Errorf("permssion deny:%s,required:%s", tk.UserType, ut))
				return
			}
		}

	}

	// next flow
	next.ProcessFilter(req, resp)

	//	响应拦截
	//a.log.Debug(resp)

}
