// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/role/pb/rpc.proto

package role

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

// QueryRoleRequest 列表查询
type QueryRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// @gotags: json:"type"
	Type *RoleType `protobuf:"varint,2,opt,name=type,proto3,enum=go8.devcloud.mcenter.role.RoleType,oneof" json:"type"`
	// 属于那个域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain"`
}

func (x *QueryRoleRequest) Reset() {
	*x = QueryRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRoleRequest) ProtoMessage() {}

func (x *QueryRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRoleRequest.ProtoReflect.Descriptor instead.
func (*QueryRoleRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRoleRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryRoleRequest) GetType() RoleType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return RoleType_BUILDIN
}

func (x *QueryRoleRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

// DescribeRoleRequest role详情
type DescribeRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// @gotags: json:"name,omitempty" validate:"required,lte=64"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" validate:"required,lte=64"`
	// @gotags: bson:"with_permissions" json:"with_permissions"
	WithPermissions bool `protobuf:"varint,3,opt,name=with_permissions,json=withPermissions,proto3" json:"with_permissions" bson:"with_permissions"`
	// @gotags: bson:"type" json:"type"
	Type RoleType `protobuf:"varint,4,opt,name=type,proto3,enum=go8.devcloud.mcenter.role.RoleType" json:"type" bson:"type"`
}

func (x *DescribeRoleRequest) Reset() {
	*x = DescribeRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRoleRequest) ProtoMessage() {}

func (x *DescribeRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRoleRequest.ProtoReflect.Descriptor instead.
func (*DescribeRoleRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeRoleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DescribeRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DescribeRoleRequest) GetWithPermissions() bool {
	if x != nil {
		return x.WithPermissions
	}
	return false
}

func (x *DescribeRoleRequest) GetType() RoleType {
	if x != nil {
		return x.Type
	}
	return RoleType_BUILDIN
}

// QueryPermissionRequest 查询用户权限
type QueryPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// @gotags: json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace"`
	// @gotags: json:"username"
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username"`
	// 查询角色权限
	// @gotags: json:"role_id"
	RoleId string `protobuf:"bytes,4,opt,name=role_id,json=roleId,proto3" json:"role_id"`
	// 忽略数据
	// @gotags: json:"skip_items"
	SkipItems bool `protobuf:"varint,5,opt,name=skip_items,json=skipItems,proto3" json:"skip_items"`
}

func (x *QueryPermissionRequest) Reset() {
	*x = QueryPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPermissionRequest) ProtoMessage() {}

func (x *QueryPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPermissionRequest.ProtoReflect.Descriptor instead.
func (*QueryPermissionRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *QueryPermissionRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryPermissionRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryPermissionRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *QueryPermissionRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *QueryPermissionRequest) GetSkipItems() bool {
	if x != nil {
		return x.SkipItems
	}
	return false
}

// DescribeRoleRequest role详情
type DescribePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribePermissionRequest) Reset() {
	*x = DescribePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePermissionRequest) ProtoMessage() {}

func (x *DescribePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePermissionRequest.ProtoReflect.Descriptor instead.
func (*DescribePermissionRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *DescribePermissionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteRoleRequest role删除
type DeleteRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id" validate:"required,lte=64"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" validate:"required,lte=64"`
	// @gotags: json:"delete_policy"
	DeletePolicy bool `protobuf:"varint,2,opt,name=delete_policy,json=deletePolicy,proto3" json:"delete_policy"`
}

func (x *DeleteRoleRequest) Reset() {
	*x = DeleteRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleRequest) ProtoMessage() {}

func (x *DeleteRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRoleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteRoleRequest) GetDeletePolicy() bool {
	if x != nil {
		return x.DeletePolicy
	}
	return false
}

type AddPermissionToRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 创建者ID
	// @gotags: json:"create_by" validate:"required"
	CreateBy string `protobuf:"bytes,3,opt,name=create_by,json=createBy,proto3" json:"create_by" validate:"required"`
	// @gotags: json:"role_id" validate:"required,lte=64"
	RoleId string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id" validate:"required,lte=64"`
	// @gotags: json:"permissions" validate:"required"
	Permissions []*Spec `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions" validate:"required"`
}

func (x *AddPermissionToRoleRequest) Reset() {
	*x = AddPermissionToRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPermissionToRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPermissionToRoleRequest) ProtoMessage() {}

func (x *AddPermissionToRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPermissionToRoleRequest.ProtoReflect.Descriptor instead.
func (*AddPermissionToRoleRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *AddPermissionToRoleRequest) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *AddPermissionToRoleRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *AddPermissionToRoleRequest) GetPermissions() []*Spec {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type RemovePermissionFromRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"role_id" validate:"required,lte=64"
	RoleId string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id" validate:"required,lte=64"`
	// @gotags: json:"remove_all"
	RemoveAll bool `protobuf:"varint,2,opt,name=remove_all,json=removeAll,proto3" json:"remove_all"`
	// @gotags: json:"permission_id"
	PermissionId []string `protobuf:"bytes,3,rep,name=permission_id,json=permissionId,proto3" json:"permission_id"`
}

func (x *RemovePermissionFromRoleRequest) Reset() {
	*x = RemovePermissionFromRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePermissionFromRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePermissionFromRoleRequest) ProtoMessage() {}

func (x *RemovePermissionFromRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePermissionFromRoleRequest.ProtoReflect.Descriptor instead.
func (*RemovePermissionFromRoleRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *RemovePermissionFromRoleRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *RemovePermissionFromRoleRequest) GetRemoveAll() bool {
	if x != nil {
		return x.RemoveAll
	}
	return false
}

func (x *RemovePermissionFromRoleRequest) GetPermissionId() []string {
	if x != nil {
		return x.PermissionId
	}
	return nil
}

type UpdatePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// permission id
	// @gotags: json:"id" validate:"required,lte=64"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" validate:"required,lte=64"`
	// 维度
	// @gotags: json:"label_key"
	LabelKey string `protobuf:"bytes,2,opt,name=label_key,json=labelKey,proto3" json:"label_key"`
	// 适配所有值
	// @gotags: json:"match_all"
	MatchAll bool `protobuf:"varint,3,opt,name=match_all,json=matchAll,proto3" json:"match_all"`
	// 标识值
	// @gotags: json:"label_values"
	LabelValues []string `protobuf:"bytes,4,rep,name=label_values,json=labelValues,proto3" json:"label_values"`
}

func (x *UpdatePermissionRequest) Reset() {
	*x = UpdatePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePermissionRequest) ProtoMessage() {}

func (x *UpdatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePermissionRequest.ProtoReflect.Descriptor instead.
func (*UpdatePermissionRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePermissionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePermissionRequest) GetLabelKey() string {
	if x != nil {
		return x.LabelKey
	}
	return ""
}

func (x *UpdatePermissionRequest) GetMatchAll() bool {
	if x != nil {
		return x.MatchAll
	}
	return false
}

func (x *UpdatePermissionRequest) GetLabelValues() []string {
	if x != nil {
		return x.LabelValues
	}
	return nil
}

var File_apps_role_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_role_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65,
	0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x1a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65,
	0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x10,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62,
	0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x77,
	0x69, 0x74, 0x68, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x37,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x67,
	0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xc2, 0x01, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x73, 0x6b, 0x69, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2b, 0x0a, 0x19,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x22, 0x95, 0x01, 0x0a, 0x1a, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12,
	0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x7e, 0x0a, 0x1f, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46,
	0x72, 0x6f, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x17,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x6c,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x6c,
	0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x32, 0xa7, 0x03, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x5c, 0x0a, 0x09,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2b, 0x2e, 0x67, 0x6f, 0x38, 0x2e,
	0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x12, 0x5f, 0x0a, 0x0c, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2e, 0x2e, 0x67, 0x6f, 0x38,
	0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x6f, 0x38,
	0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x6e, 0x0a, 0x0f, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x31,
	0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x71, 0x0a, 0x12, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x34, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65,
	0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x49,
	0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x61, 0x73,
	0x6d, 0x69, 0x6e, 0x65, 0x34, 0x35, 0x36, 0x2f, 0x67, 0x6f, 0x5f, 0x38, 0x5f, 0x6d, 0x61, 0x67,
	0x65, 0x2f, 0x77, 0x65, 0x65, 0x6b, 0x31, 0x34, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2f, 0x64,
	0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_apps_role_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_role_pb_rpc_proto_rawDescData = file_apps_role_pb_rpc_proto_rawDesc
)

func file_apps_role_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_role_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_role_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_role_pb_rpc_proto_rawDescData)
	})
	return file_apps_role_pb_rpc_proto_rawDescData
}

var file_apps_role_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_apps_role_pb_rpc_proto_goTypes = []interface{}{
	(*QueryRoleRequest)(nil),                // 0: go8.devcloud.mcenter.role.QueryRoleRequest
	(*DescribeRoleRequest)(nil),             // 1: go8.devcloud.mcenter.role.DescribeRoleRequest
	(*QueryPermissionRequest)(nil),          // 2: go8.devcloud.mcenter.role.QueryPermissionRequest
	(*DescribePermissionRequest)(nil),       // 3: go8.devcloud.mcenter.role.DescribePermissionRequest
	(*DeleteRoleRequest)(nil),               // 4: go8.devcloud.mcenter.role.DeleteRoleRequest
	(*AddPermissionToRoleRequest)(nil),      // 5: go8.devcloud.mcenter.role.AddPermissionToRoleRequest
	(*RemovePermissionFromRoleRequest)(nil), // 6: go8.devcloud.mcenter.role.RemovePermissionFromRoleRequest
	(*UpdatePermissionRequest)(nil),         // 7: go8.devcloud.mcenter.role.UpdatePermissionRequest
	(*request.PageRequest)(nil),             // 8: infraboard.mcube.page.PageRequest
	(RoleType)(0),                           // 9: go8.devcloud.mcenter.role.RoleType
	(*Spec)(nil),                            // 10: go8.devcloud.mcenter.role.Spec
	(*RoleSet)(nil),                         // 11: go8.devcloud.mcenter.role.RoleSet
	(*Role)(nil),                            // 12: go8.devcloud.mcenter.role.Role
	(*PermissionSet)(nil),                   // 13: go8.devcloud.mcenter.role.PermissionSet
	(*Permission)(nil),                      // 14: go8.devcloud.mcenter.role.Permission
}
var file_apps_role_pb_rpc_proto_depIdxs = []int32{
	8,  // 0: go8.devcloud.mcenter.role.QueryRoleRequest.page:type_name -> infraboard.mcube.page.PageRequest
	9,  // 1: go8.devcloud.mcenter.role.QueryRoleRequest.type:type_name -> go8.devcloud.mcenter.role.RoleType
	9,  // 2: go8.devcloud.mcenter.role.DescribeRoleRequest.type:type_name -> go8.devcloud.mcenter.role.RoleType
	8,  // 3: go8.devcloud.mcenter.role.QueryPermissionRequest.page:type_name -> infraboard.mcube.page.PageRequest
	10, // 4: go8.devcloud.mcenter.role.AddPermissionToRoleRequest.permissions:type_name -> go8.devcloud.mcenter.role.Spec
	0,  // 5: go8.devcloud.mcenter.role.RPC.QueryRole:input_type -> go8.devcloud.mcenter.role.QueryRoleRequest
	1,  // 6: go8.devcloud.mcenter.role.RPC.DescribeRole:input_type -> go8.devcloud.mcenter.role.DescribeRoleRequest
	2,  // 7: go8.devcloud.mcenter.role.RPC.QueryPermission:input_type -> go8.devcloud.mcenter.role.QueryPermissionRequest
	3,  // 8: go8.devcloud.mcenter.role.RPC.DescribePermission:input_type -> go8.devcloud.mcenter.role.DescribePermissionRequest
	11, // 9: go8.devcloud.mcenter.role.RPC.QueryRole:output_type -> go8.devcloud.mcenter.role.RoleSet
	12, // 10: go8.devcloud.mcenter.role.RPC.DescribeRole:output_type -> go8.devcloud.mcenter.role.Role
	13, // 11: go8.devcloud.mcenter.role.RPC.QueryPermission:output_type -> go8.devcloud.mcenter.role.PermissionSet
	14, // 12: go8.devcloud.mcenter.role.RPC.DescribePermission:output_type -> go8.devcloud.mcenter.role.Permission
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_apps_role_pb_rpc_proto_init() }
func file_apps_role_pb_rpc_proto_init() {
	if File_apps_role_pb_rpc_proto != nil {
		return
	}
	file_apps_role_pb_role_proto_init()
	file_apps_role_pb_permission_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_role_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRoleRequest); i {
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
		file_apps_role_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRoleRequest); i {
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
		file_apps_role_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPermissionRequest); i {
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
		file_apps_role_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePermissionRequest); i {
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
		file_apps_role_pb_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleRequest); i {
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
		file_apps_role_pb_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPermissionToRoleRequest); i {
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
		file_apps_role_pb_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePermissionFromRoleRequest); i {
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
		file_apps_role_pb_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePermissionRequest); i {
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
	file_apps_role_pb_rpc_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_role_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_role_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_role_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_role_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_role_pb_rpc_proto = out.File
	file_apps_role_pb_rpc_proto_rawDesc = nil
	file_apps_role_pb_rpc_proto_goTypes = nil
	file_apps_role_pb_rpc_proto_depIdxs = nil
}