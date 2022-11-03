package api

import (
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/token"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) IssueToken(r *restful.Request, w *restful.Response) {
	req := token.NewIssueTokenRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	tk, err := h.service.IssueToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, tk)
}

// Refresh Token必须传递, 判断撤销者身份
func (u *handler) RevolkToken(r *restful.Request, w *restful.Response) {
	qs := r.Request.URL.Query()
	req := token.NewRevolkTokenRequest()
	// 从HTTP Header中获取Access token
	req.AccessToken = token.GetTokenFromHTTPHeader(r.Request)
	// 从Query string获取Refresh Token
	req.RefreshToken = qs.Get("refresh_token")

	ins, err := h.service.RevolkToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, ins)
}

func (u *handler) ChangeNamespace(r *restful.Request, w *restful.Response) {
	req := token.NewChangeNamespaceRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.ChangeNamespace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}

func (u *handler)ValidateToken(r *restful.Request,w *restful.Response){
	tk:=r.Request.Header.Get(token.VALIDATE_TOKEN_HEADER_KEY)
	req:=token.NewValidateTokenRequest(tk)

	resp,err:=h.service.ValidateToken(r.Request.Context(),req)
	if err != nil {
		response.Failed(w.ResponseWriter,err)
		return
	}
	response.Success(w.ResponseWriter,resp)
}