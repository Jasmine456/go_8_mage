package tag

import "context"

type Service interface {
	Query(context.Context,*QueryRequest)(*TagSet,error)
}

func NewQueryRequest(blogId int) *QueryRequest{
	return &QueryRequest{
		Limte:20,
	}
}


type QueryRequest struct {
	BlogId int
	Limte int
}
