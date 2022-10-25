package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"

	"go_8_mage/week14_after/micro/tools/demo/apps/book"
)

//之前的Handler都是写Gin的，本项目用的"github.com/emicklei/go-restful/v3"
func (h *handler) CreateBook(r *restful.Request, w *restful.Response) {
	req := book.NewCreateBookRequest()

	//读取HTTP数据，c.BindJson(),r.ReadEntity() BindJson
	//gin ctx = Request + Response
	// go-restful = request(Http Request) response(HTTP http Response)
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
