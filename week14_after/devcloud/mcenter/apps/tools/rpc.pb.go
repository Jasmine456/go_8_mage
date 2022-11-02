// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/endpoint/pb/rpc.proto

package tools

import (
	request "github.com/infraboard/mcube/http/request"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RegistryRequest 服务注册请求, 验证服务的合法性, 需要客户端凭证
type RegistryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"client_id" validate:"required"
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id" validate:"required"`
	// @gotags: json:"client_secret" validate:"required"
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret" validate:"required"`
	// @gotags: json:"version" validate:"required,lte=32"
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version" validate:"required,lte=32"`
	// @gotags: json:"entries"
	Entries []*Entry `protobuf:"bytes,4,rep,name=entries,proto3" json:"entries"`
}

func (x *RegistryRequest) Reset() {
	*x = RegistryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryRequest) ProtoMessage() {}

func (x *RegistryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryRequest.ProtoReflect.Descriptor instead.
func (*RegistryRequest) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *RegistryRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *RegistryRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *RegistryRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *RegistryRequest) GetEntries() []*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

// RegistryReponse todo
type RegistryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"message"
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
}

func (x *RegistryResponse) Reset() {
	*x = RegistryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryResponse) ProtoMessage() {}

func (x *RegistryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryResponse.ProtoReflect.Descriptor instead.
func (*RegistryResponse) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *RegistryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// DescribeEndpointRequest todo
type DescribeEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribeEndpointRequest) Reset() {
	*x = DescribeEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeEndpointRequest) ProtoMessage() {}

func (x *DescribeEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeEndpointRequest.ProtoReflect.Descriptor instead.
func (*DescribeEndpointRequest) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeEndpointRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// QueryEndpointRequest 查询应用列表
type QueryEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 几个服务总共由哪些Endpoinit
	// 查询完 Endpoint, 根据Label进行匹配
	// @gotags: json:"service_ids"
	ServiceIds []string `protobuf:"bytes,2,rep,name=service_ids,json=serviceIds,proto3" json:"service_ids"`
	// @gotags: json:"path"
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path"`
	// @gotags: json:"method"
	Method string `protobuf:"bytes,4,opt,name=method,proto3" json:"method"`
	// @gotags: json:"function_name"
	FunctionName string `protobuf:"bytes,5,opt,name=function_name,json=functionName,proto3" json:"function_name"`
	// 根据资源名称和标签能匹配到对应的Endpoint
	// @gotags: json:"resources"
	Resources []string `protobuf:"bytes,6,rep,name=resources,proto3" json:"resources"`
	// @gotags: json:"labels"
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// @gotags: json:"permission_enable"
	PermissionEnable *bool `protobuf:"varint,8,opt,name=permission_enable,json=permissionEnable,proto3,oneof" json:"permission_enable"`
}

func (x *QueryEndpointRequest) Reset() {
	*x = QueryEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryEndpointRequest) ProtoMessage() {}

func (x *QueryEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryEndpointRequest.ProtoReflect.Descriptor instead.
func (*QueryEndpointRequest) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *QueryEndpointRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryEndpointRequest) GetServiceIds() []string {
	if x != nil {
		return x.ServiceIds
	}
	return nil
}

func (x *QueryEndpointRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *QueryEndpointRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *QueryEndpointRequest) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *QueryEndpointRequest) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *QueryEndpointRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *QueryEndpointRequest) GetPermissionEnable() bool {
	if x != nil && x.PermissionEnable != nil {
		return *x.PermissionEnable
	}
	return false
}

var File_apps_endpoint_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_endpoint_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f,
	0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f,
	0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65,
	0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a,
	0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x07,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x2c, 0x0a, 0x10,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xba, 0x03, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x57, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x30, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x10, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x14, 0x0a, 0x12,
	0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x32, 0xe2, 0x02, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x73, 0x0a, 0x10, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2e,
	0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x73, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x36, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x6f,
	0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x71, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x33, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x6f,
	0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x61, 0x73, 0x6d, 0x69, 0x6e, 0x65, 0x34, 0x35, 0x36,
	0x2f, 0x67, 0x6f, 0x5f, 0x38, 0x5f, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x77, 0x65, 0x65, 0x6b, 0x31,
	0x34, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_endpoint_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_endpoint_pb_rpc_proto_rawDescData = file_apps_endpoint_pb_rpc_proto_rawDesc
)

func file_apps_endpoint_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_endpoint_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_endpoint_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_endpoint_pb_rpc_proto_rawDescData)
	})
	return file_apps_endpoint_pb_rpc_proto_rawDescData
}

var file_apps_endpoint_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_endpoint_pb_rpc_proto_goTypes = []interface{}{
	(*RegistryRequest)(nil),         // 0: go8.devcloud.mcenter.endpoint.RegistryRequest
	(*RegistryResponse)(nil),        // 1: go8.devcloud.mcenter.endpoint.RegistryResponse
	(*DescribeEndpointRequest)(nil), // 2: go8.devcloud.mcenter.endpoint.DescribeEndpointRequest
	(*QueryEndpointRequest)(nil),    // 3: go8.devcloud.mcenter.endpoint.QueryEndpointRequest
	nil,                             // 4: go8.devcloud.mcenter.endpoint.QueryEndpointRequest.LabelsEntry
	(*Entry)(nil),                   // 5: go8.devcloud.mcenter.endpoint.Entry
	(*request.PageRequest)(nil),     // 6: infraboard.mcube.page.PageRequest
	(*Endpoint)(nil),                // 7: go8.devcloud.mcenter.endpoint.Endpoint
	(*EndpointSet)(nil),             // 8: go8.devcloud.mcenter.endpoint.EndpointSet
}
var file_apps_endpoint_pb_rpc_proto_depIdxs = []int32{
	5, // 0: go8.devcloud.mcenter.endpoint.RegistryRequest.entries:type_name -> go8.devcloud.mcenter.endpoint.Entry
	6, // 1: go8.devcloud.mcenter.endpoint.QueryEndpointRequest.page:type_name -> infraboard.mcube.page.PageRequest
	4, // 2: go8.devcloud.mcenter.endpoint.QueryEndpointRequest.labels:type_name -> go8.devcloud.mcenter.endpoint.QueryEndpointRequest.LabelsEntry
	0, // 3: go8.devcloud.mcenter.endpoint.RPC.RegistryEndpoint:input_type -> go8.devcloud.mcenter.endpoint.RegistryRequest
	2, // 4: go8.devcloud.mcenter.endpoint.RPC.DescribeEndpoint:input_type -> go8.devcloud.mcenter.endpoint.DescribeEndpointRequest
	3, // 5: go8.devcloud.mcenter.endpoint.RPC.QueryEndpoints:input_type -> go8.devcloud.mcenter.endpoint.QueryEndpointRequest
	1, // 6: go8.devcloud.mcenter.endpoint.RPC.RegistryEndpoint:output_type -> go8.devcloud.mcenter.endpoint.RegistryResponse
	7, // 7: go8.devcloud.mcenter.endpoint.RPC.DescribeEndpoint:output_type -> go8.devcloud.mcenter.endpoint.Endpoint
	8, // 8: go8.devcloud.mcenter.endpoint.RPC.QueryEndpoints:output_type -> go8.devcloud.mcenter.endpoint.EndpointSet
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_endpoint_pb_rpc_proto_init() }
func file_apps_endpoint_pb_rpc_proto_init() {
	if File_apps_endpoint_pb_rpc_proto != nil {
		return
	}
	file_apps_endpoint_pb_endpoint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_endpoint_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_endpoint_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_endpoint_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeEndpointRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_endpoint_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryEndpointRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_apps_endpoint_pb_rpc_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_endpoint_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_endpoint_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_endpoint_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_endpoint_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_endpoint_pb_rpc_proto = out.File
	file_apps_endpoint_pb_rpc_proto_rawDesc = nil
	file_apps_endpoint_pb_rpc_proto_goTypes = nil
	file_apps_endpoint_pb_rpc_proto_depIdxs = nil
}
