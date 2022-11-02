package tools

import (
	"fmt"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/endpoint"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
)

// 用于route转换成 Entry
func TransferRoutesToEndpoints(routes []restful.Route) (endpoints []*endpoint.Entry) {
	for _, r := range routes {
		var resource, action string
		var authEnabled, permEnabled bool
		if r.Metadata != nil {
			if v, ok := r.Metadata[label.Resource]; ok {
				resource, _ = v.(string)
			}
			if v, ok := r.Metadata[label.Action]; ok {
				action, _ = v.(string)
			}
			if v, ok := r.Metadata[label.Auth]; ok {
				authEnabled, _ = v.(bool)
			}
			if v, ok := r.Metadata[label.Permission]; ok {
				permEnabled, _ = v.(bool)
			}
		}
		endpoints = append(endpoints, &endpoint.Entry{
			FunctionName:     r.Operation,
			Resource:         resource,
			Path:             fmt.Sprintf("%s.%s", r.Method, r.Path),
			Method:           r.Method,
			AuthEnable:       authEnabled,
			PermissionEnable: permEnabled,
			Labels:           map[string]string{"action": action},
		})
	}
	return
}
