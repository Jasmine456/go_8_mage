package api

import (
	"fmt"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/token"
	"github.com/emicklei/go-restful/v3"
	"github.com/go_8_mage/week14_after/devcloud/maudit/apps/book"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) CreateBook(r *restful.Request, w *restful.Response) {
	req := book.NewCreateBookRequest()
	if req ==nil{
		response.Failed(w.ResponseWriter, fmt.Errorf("请添加创建book参数"))
		return
	}
	//需要拿到用户身份信息
	//通过中间件已经知道，怎么才能知道中间件的返回结果
	//在go restful 提供一个map attributes map[string]interface{} //for storing request-scoped values
	//用于存储请求的 是一些上下文信息，认证信息 就放到上下文中的 map，通过Request的Attribute来写入Map r.SetAttribute()
	tk := r.Attribute("tk")

	//引入mcenter依赖库，go mod tidy下载依赖
	if tk != nil {
		if v := tk.(*token.Token); v != nil {
			fmt.Println(v.UserId)
		}
	}
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.CreateBook(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, set)
}

func (h *handler) QueryBook(r *restful.Request, w *restful.Response) {
	req := book.NewQueryBookRequestFromHTTP(r.Request)
	set, err := h.service.QueryBook(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}

func (h *handler) DescribeBook(r *restful.Request, w *restful.Response) {
	req := book.NewDescribeBookRequest(r.PathParameter("id"))
	ins, err := h.service.DescribeBook(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, ins)
}

func (h *handler) UpdateBook(r *restful.Request, w *restful.Response) {
	req := book.NewPutBookRequest(r.PathParameter("id"))

	if err := r.ReadEntity(req.Data); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.UpdateBook(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}

func (h *handler) PatchBook(r *restful.Request, w *restful.Response) {
	req := book.NewPatchBookRequest(r.PathParameter("id"))

	if err := r.ReadEntity(req.Data); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.UpdateBook(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}

func (h *handler) DeleteBook(r *restful.Request, w *restful.Response) {
	req := book.NewDeleteBookRequestWithID(r.PathParameter("id"))
	set, err := h.service.DeleteBook(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}
