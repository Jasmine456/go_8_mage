package operate

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
)

const (
	AppName = "operate"
)

var (
	validate = validator.New()
)

type Service interface {
	RPCServer
}

func NewOperateLog() *OperateLog {
	return &OperateLog{}
}

func (l *OperateLog) Validate() error {
	return validate.Struct(l)
}

func (req *QueryOperateLogRequest) Validate() error {
	return nil
}

func NewOperateLogSet() *OperateLogSet {
	return &OperateLogSet{
		Items: []*OperateLog{},
	}
}

func (s *OperateLogSet) Add(item *OperateLog) {
	s.Items = append(s.Items, item)
}

func NewQueryOperateLogRequest() *QueryOperateLogRequest {
	return &QueryOperateLogRequest{
		Page: request.NewDefaultPageRequest(),
	}
}
