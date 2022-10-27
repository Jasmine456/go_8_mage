// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/service/pb/rpc.proto

package service

import (
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

type DescribeBy int32

const (
	// 通过service id查询应用详情
	DescribeBy_SERVICE_ID DescribeBy = 0
	// 通过service name查询应用详情
	DescribeBy_SERVICE_NAME DescribeBy = 1
	// 通过service client_id查询应用详情
	DescribeBy_SERVICE_CLIENT_ID DescribeBy = 2
)

// Enum value maps for DescribeBy.
var (
	DescribeBy_name = map[int32]string{
		0: "SERVICE_ID",
		1: "SERVICE_NAME",
		2: "SERVICE_CLIENT_ID",
	}
	DescribeBy_value = map[string]int32{
		"SERVICE_ID":        0,
		"SERVICE_NAME":      1,
		"SERVICE_CLIENT_ID": 2,
	}
)

func (x DescribeBy) Enum() *DescribeBy {
	p := new(DescribeBy)
	*p = x
	return p
}

func (x DescribeBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DescribeBy) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_service_pb_rpc_proto_enumTypes[0].Descriptor()
}

func (DescribeBy) Type() protoreflect.EnumType {
	return &file_apps_service_pb_rpc_proto_enumTypes[0]
}

func (x DescribeBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DescribeBy.Descriptor instead.
func (DescribeBy) EnumDescriptor() ([]byte, []int) {
	return file_apps_service_pb_rpc_proto_rawDescGZIP(), []int{0}
}

// ValidateCredentialRequest 校验服务凭证
type ValidateCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务客户端ID
	// @gotags: json:"client_id" yaml:"client_id" validate:"required,lte=100"
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id" yaml:"client_id" validate:"required,lte=100"`
	// 服务客户端凭证
	// @gotags: json:"client_secret" yaml:"client_secret" validate:"required,lte=100"
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret" yaml:"client_secret" validate:"required,lte=100"`
}

func (x *ValidateCredentialRequest) Reset() {
	*x = ValidateCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_service_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCredentialRequest) ProtoMessage() {}

func (x *ValidateCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_service_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCredentialRequest.ProtoReflect.Descriptor instead.
func (*ValidateCredentialRequest) Descriptor() ([]byte, []int) {
	return file_apps_service_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateCredentialRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ValidateCredentialRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

// DescribeMicroRequest 查询应用详情
type DescribeServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 查询详情的方式
	// @gotags: json:"describe_by" yaml:"describe_by"
	DescribeBy DescribeBy `protobuf:"varint,1,opt,name=describe_by,json=describeBy,proto3,enum=go8.devcloud.mcenter.service.DescribeBy" json:"describe_by" yaml:"describe_by"`
	// 服务客户端Id
	// @gotags: json:"client_id" yaml:"client_id"
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id" yaml:"client_id"`
	// 服务名称
	// @gotags: json:"name" yaml:"name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" yaml:"name"`
	// 服务Id
	// @gotags: json:"id" yaml:"id"
	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id" yaml:"id"`
}

func (x *DescribeServiceRequest) Reset() {
	*x = DescribeServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_service_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeServiceRequest) ProtoMessage() {}

func (x *DescribeServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_service_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeServiceRequest.ProtoReflect.Descriptor instead.
func (*DescribeServiceRequest) Descriptor() ([]byte, []int) {
	return file_apps_service_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeServiceRequest) GetDescribeBy() DescribeBy {
	if x != nil {
		return x.DescribeBy
	}
	return DescribeBy_SERVICE_ID
}

func (x *DescribeServiceRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *DescribeServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DescribeServiceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_apps_service_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_service_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f, 0x38,
	0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1d, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x19, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0xa4, 0x01, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x49, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x62,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65,
	0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42,
	0x79, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x45,
	0x0a, 0x0a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x12, 0x0e, 0x0a, 0x0a,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x15,
	0x0a, 0x11, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54,
	0x5f, 0x49, 0x44, 0x10, 0x02, 0x32, 0xeb, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x74, 0x0a,
	0x12, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x12, 0x37, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67,
	0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67,
	0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4a, 0x61, 0x73, 0x6d, 0x69, 0x6e, 0x65, 0x34, 0x35, 0x36, 0x2f, 0x67, 0x6f, 0x5f,
	0x38, 0x5f, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x77, 0x65, 0x65, 0x6b, 0x31, 0x34, 0x5f, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_service_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_service_pb_rpc_proto_rawDescData = file_apps_service_pb_rpc_proto_rawDesc
)

func file_apps_service_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_service_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_service_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_service_pb_rpc_proto_rawDescData)
	})
	return file_apps_service_pb_rpc_proto_rawDescData
}

var file_apps_service_pb_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_service_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apps_service_pb_rpc_proto_goTypes = []interface{}{
	(DescribeBy)(0),                   // 0: go8.devcloud.mcenter.service.DescribeBy
	(*ValidateCredentialRequest)(nil), // 1: go8.devcloud.mcenter.service.ValidateCredentialRequest
	(*DescribeServiceRequest)(nil),    // 2: go8.devcloud.mcenter.service.DescribeServiceRequest
	(*Service)(nil),                   // 3: go8.devcloud.mcenter.service.Service
}
var file_apps_service_pb_rpc_proto_depIdxs = []int32{
	0, // 0: go8.devcloud.mcenter.service.DescribeServiceRequest.describe_by:type_name -> go8.devcloud.mcenter.service.DescribeBy
	1, // 1: go8.devcloud.mcenter.service.RPC.ValidateCredential:input_type -> go8.devcloud.mcenter.service.ValidateCredentialRequest
	2, // 2: go8.devcloud.mcenter.service.RPC.DescribeService:input_type -> go8.devcloud.mcenter.service.DescribeServiceRequest
	3, // 3: go8.devcloud.mcenter.service.RPC.ValidateCredential:output_type -> go8.devcloud.mcenter.service.Service
	3, // 4: go8.devcloud.mcenter.service.RPC.DescribeService:output_type -> go8.devcloud.mcenter.service.Service
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_service_pb_rpc_proto_init() }
func file_apps_service_pb_rpc_proto_init() {
	if File_apps_service_pb_rpc_proto != nil {
		return
	}
	file_apps_service_pb_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_service_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateCredentialRequest); i {
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
		file_apps_service_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeServiceRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_service_pb_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_service_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_service_pb_rpc_proto_depIdxs,
		EnumInfos:         file_apps_service_pb_rpc_proto_enumTypes,
		MessageInfos:      file_apps_service_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_service_pb_rpc_proto = out.File
	file_apps_service_pb_rpc_proto_rawDesc = nil
	file_apps_service_pb_rpc_proto_goTypes = nil
	file_apps_service_pb_rpc_proto_depIdxs = nil
}
