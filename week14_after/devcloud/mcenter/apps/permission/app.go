package permission

import "github.com/infraboard/mcube/http/request"

const (
	AppName = "permission"
)

type Service interface {
	RPCServer
}

// NewCheckPermissionrequest todo
func NewCheckPermissionRequest() *CheckPermissionRequest {
	return &CheckPermissionRequest{
		Page: request.NewPageRequest(100, 1),
	}
}
