package endpoint

import (
	"fmt"
	"hash/fnv"
	"time"
)

const (
	AppName = "endpoint"
)

type Service interface {
	RPCServer
}

// Endpoints 功能列表
func (req *RegistryRequest) Endpoints(serviceID string) []*Endpoint {
	eps := make([]*Endpoint, 0, len(req.Entries))
	for i := range req.Entries {
		ep := &Endpoint{
			// 为该功能生成一个唯一ID, 这里的Path包含method.path
			// path 是url唯一建
			Id:        GenHashID(serviceID, req.Entries[i].Path),
			CreateAt:  time.Now().UnixMilli(),
			UpdateAt:  time.Now().UnixMilli(),
			ServiceId: serviceID,
			Version:   req.Version,
			Entry:     req.Entries[i],
		}
		eps = append(eps, ep)
	}
	return eps
}

// GenHashID hash id
func GenHashID(service, grpcPath string) string {
	hashedStr := fmt.Sprintf("%s-%s", service, grpcPath)
	h := fnv.New32a()
	h.Write([]byte(hashedStr))
	return fmt.Sprintf("%x", h.Sum32())
}

// NewRegistryResponse todo
func NewRegistryResponse(message string) *RegistryResponse {
	return &RegistryResponse{Message: message}
}

// NewDefaultEndpoint todo
func NewDefaultEndpoint() *Endpoint {
	return &Endpoint{
		Entry: &Entry{},
	}
}

// NewEndpointSet 实例化
func NewEndpointSet() *EndpointSet {
	return &EndpointSet{
		Items: []*Endpoint{},
	}
}

// Add 添加
func (s *EndpointSet) Add(e *Endpoint) {
	s.Items = append(s.Items, e)
}

// NewRegistryRequest 注册请求
func NewRegistryRequest(version string, entries []*Entry) *RegistryRequest {
	return &RegistryRequest{
		Version: version,
		Entries: entries,
	}
}
