// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/user/pb/rpc.proto

package user

import (
	request "github.com/infraboard/mcube/http/request"
	request1 "github.com/infraboard/mcube/pb/request"
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

type DESCRIBE_BY int32

const (
	// 通过UserId查询用户
	DESCRIBE_BY_USER_ID DESCRIBE_BY = 0
	// 通过Username查询用户
	DESCRIBE_BY_USER_NAME DESCRIBE_BY = 1
)

// Enum value maps for DESCRIBE_BY.
var (
	DESCRIBE_BY_name = map[int32]string{
		0: "USER_ID",
		1: "USER_NAME",
	}
	DESCRIBE_BY_value = map[string]int32{
		"USER_ID":   0,
		"USER_NAME": 1,
	}
)

func (x DESCRIBE_BY) Enum() *DESCRIBE_BY {
	p := new(DESCRIBE_BY)
	*p = x
	return p
}

func (x DESCRIBE_BY) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DESCRIBE_BY) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_user_pb_rpc_proto_enumTypes[0].Descriptor()
}

func (DESCRIBE_BY) Type() protoreflect.EnumType {
	return &file_apps_user_pb_rpc_proto_enumTypes[0]
}

func (x DESCRIBE_BY) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DESCRIBE_BY.Descriptor instead.
func (DESCRIBE_BY) EnumDescriptor() ([]byte, []int) {
	return file_apps_user_pb_rpc_proto_rawDescGZIP(), []int{0}
}

// DescribeUserRequest 查询用户详情
type DescribeUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 查询方式
	// @gotags: json:"describe_by"
	DescribeBy DESCRIBE_BY `protobuf:"varint,1,opt,name=describe_by,json=describeBy,proto3,enum=go8.devcloud.mcenter.DESCRIBE_BY" json:"describe_by"`
	// 用户账号id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 用户账号
	// @gotags: json:"username"
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username"`
}

func (x *DescribeUserRequest) Reset() {
	*x = DescribeUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeUserRequest) ProtoMessage() {}

func (x *DescribeUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeUserRequest.ProtoReflect.Descriptor instead.
func (*DescribeUserRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *DescribeUserRequest) GetDescribeBy() DESCRIBE_BY {
	if x != nil {
		return x.DescribeBy
	}
	return DESCRIBE_BY_USER_ID
}

func (x *DescribeUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DescribeUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// QueryUserRequest 获取子账号列表
type QueryUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 用户所属Domain
	// @gotags: json:"domain" validate:"required"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain" validate:"required"`
	// 账号提供方
	// @gotags: json:"provider"
	Provider *PROVIDER `protobuf:"varint,3,opt,name=provider,proto3,enum=go8.devcloud.mcenter.PROVIDER,oneof" json:"provider"`
	// 用户类型
	// @gotags: json:"type"
	Type *TYPE `protobuf:"varint,4,opt,name=type,proto3,enum=go8.devcloud.mcenter.TYPE,oneof" json:"type"`
	// 通过Id
	// @gotags: json:"user_ids"
	UserIds []string `protobuf:"bytes,5,rep,name=user_ids,json=userIds,proto3" json:"user_ids"`
	// 是否获取数据
	// @gotags: json:"skip_items"
	SkipItems bool `protobuf:"varint,8,opt,name=skip_items,json=skipItems,proto3" json:"skip_items"`
	// 关键字查询
	// @gotags: json:"keywords"
	Keywords string `protobuf:"bytes,9,opt,name=keywords,proto3" json:"keywords"`
}

func (x *QueryUserRequest) Reset() {
	*x = QueryUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserRequest) ProtoMessage() {}

func (x *QueryUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserRequest.ProtoReflect.Descriptor instead.
func (*QueryUserRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *QueryUserRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryUserRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *QueryUserRequest) GetProvider() PROVIDER {
	if x != nil && x.Provider != nil {
		return *x.Provider
	}
	return PROVIDER_LOCAL
}

func (x *QueryUserRequest) GetType() TYPE {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return TYPE_SUB
}

func (x *QueryUserRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *QueryUserRequest) GetSkipItems() bool {
	if x != nil {
		return x.SkipItems
	}
	return false
}

func (x *QueryUserRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

// UpdatePasswordRequest 密码更新(需要知道原密码) 用户自己使用
type UpdatePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户Id
	// @gotags: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	// 旧密码
	// @gotags: json:"old_pass"
	OldPass string `protobuf:"bytes,2,opt,name=old_pass,json=oldPass,proto3" json:"old_pass"`
	// 新密码
	// @gotags: json:"new_pass"
	NewPass string `protobuf:"bytes,3,opt,name=new_pass,json=newPass,proto3" json:"new_pass"`
	// 是否重置
	// @gotags: json:"is_reset"
	IsReset bool `protobuf:"varint,4,opt,name=is_reset,json=isReset,proto3" json:"is_reset"`
	// 重置原因
	// @gotags: json:"reset_reason"
	ResetReason string `protobuf:"bytes,5,opt,name=reset_reason,json=resetReason,proto3" json:"reset_reason"`
}

func (x *UpdatePasswordRequest) Reset() {
	*x = UpdatePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePasswordRequest) ProtoMessage() {}

func (x *UpdatePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePasswordRequest.ProtoReflect.Descriptor instead.
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePasswordRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdatePasswordRequest) GetOldPass() string {
	if x != nil {
		return x.OldPass
	}
	return ""
}

func (x *UpdatePasswordRequest) GetNewPass() string {
	if x != nil {
		return x.NewPass
	}
	return ""
}

func (x *UpdatePasswordRequest) GetIsReset() bool {
	if x != nil {
		return x.IsReset
	}
	return false
}

func (x *UpdatePasswordRequest) GetResetReason() string {
	if x != nil {
		return x.ResetReason
	}
	return ""
}

// 重置密码, 用户忘了自己的密码, 管理员使用
type ResetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户名
	// @gotags: json:"user_id"
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	// 新密码
	// @gotags: json:"new_pass"
	NewPass string `protobuf:"bytes,3,opt,name=new_pass,json=newPass,proto3" json:"new_pass"`
	// 是否重置
	// @gotags: json:"is_reset"
	IsReset bool `protobuf:"varint,4,opt,name=is_reset,json=isReset,proto3" json:"is_reset"`
	// 重置原因
	// @gotags: json:"reset_reason"
	ResetReason string `protobuf:"bytes,5,opt,name=reset_reason,json=resetReason,proto3" json:"reset_reason"`
}

func (x *ResetPasswordRequest) Reset() {
	*x = ResetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordRequest) ProtoMessage() {}

func (x *ResetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordRequest.ProtoReflect.Descriptor instead.
func (*ResetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *ResetPasswordRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ResetPasswordRequest) GetNewPass() string {
	if x != nil {
		return x.NewPass
	}
	return ""
}

func (x *ResetPasswordRequest) GetIsReset() bool {
	if x != nil {
		return x.IsReset
	}
	return false
}

func (x *ResetPasswordRequest) GetResetReason() string {
	if x != nil {
		return x.ResetReason
	}
	return ""
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户账号id
	// @gotags: json:"user_ids" validate:"required,lte=60"
	UserIds []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids" validate:"required,lte=60"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteUserRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

// UpdateUserRequest todo
type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 更新模式
	// @gotags: json:"update_mode"
	UpdateMode request1.UpdateMode `protobuf:"varint,1,opt,name=update_mode,json=updateMode,proto3,enum=infraboard.mcube.request.UpdateMode" json:"update_mode"`
	// 用户Id
	// @gotags: json:"user_id" validate:"required,lte=120"
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id" validate:"required,lte=120"`
	// profile 账号profile
	// @gotags: json:"profile"
	Profile *Profile `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUserRequest) GetUpdateMode() request1.UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return request1.UpdateMode(0)
}

func (x *UpdateUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateUserRequest) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

var File_apps_user_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_user_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65,
	0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x17,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01,
	0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x38,
	0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x42, 0x59, 0x52, 0x0a, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc4, 0x02, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x3f, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x67,
	0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x48, 0x00, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x38, 0x2e,
	0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x54, 0x59, 0x50, 0x45, 0x48, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x6b, 0x69, 0x70, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x73, 0x6b, 0x69, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xa4, 0x01, 0x0a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65,
	0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65,
	0x77, 0x50, 0x61, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0x88, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x2e,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0xac,
	0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2a, 0x29, 0x0a,
	0x0b, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x42, 0x59, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x32, 0xb0, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43,
	0x12, 0x55, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x29, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x6f,
	0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67,
	0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x42, 0x49, 0x5a, 0x47, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x61, 0x73, 0x6d, 0x69, 0x6e,
	0x65, 0x34, 0x35, 0x36, 0x2f, 0x67, 0x6f, 0x5f, 0x38, 0x5f, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x77,
	0x65, 0x65, 0x6b, 0x31, 0x34, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x76, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_user_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_user_pb_rpc_proto_rawDescData = file_apps_user_pb_rpc_proto_rawDesc
)

func file_apps_user_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_user_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_user_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_user_pb_rpc_proto_rawDescData)
	})
	return file_apps_user_pb_rpc_proto_rawDescData
}

var file_apps_user_pb_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_user_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apps_user_pb_rpc_proto_goTypes = []interface{}{
	(DESCRIBE_BY)(0),              // 0: go8.devcloud.mcenter.DESCRIBE_BY
	(*DescribeUserRequest)(nil),   // 1: go8.devcloud.mcenter.DescribeUserRequest
	(*QueryUserRequest)(nil),      // 2: go8.devcloud.mcenter.QueryUserRequest
	(*UpdatePasswordRequest)(nil), // 3: go8.devcloud.mcenter.UpdatePasswordRequest
	(*ResetPasswordRequest)(nil),  // 4: go8.devcloud.mcenter.ResetPasswordRequest
	(*DeleteUserRequest)(nil),     // 5: go8.devcloud.mcenter.DeleteUserRequest
	(*UpdateUserRequest)(nil),     // 6: go8.devcloud.mcenter.UpdateUserRequest
	(*request.PageRequest)(nil),   // 7: infraboard.mcube.page.PageRequest
	(PROVIDER)(0),                 // 8: go8.devcloud.mcenter.PROVIDER
	(TYPE)(0),                     // 9: go8.devcloud.mcenter.TYPE
	(request1.UpdateMode)(0),      // 10: infraboard.mcube.request.UpdateMode
	(*Profile)(nil),               // 11: go8.devcloud.mcenter.Profile
	(*User)(nil),                  // 12: go8.devcloud.mcenter.User
	(*UserSet)(nil),               // 13: go8.devcloud.mcenter.UserSet
}
var file_apps_user_pb_rpc_proto_depIdxs = []int32{
	0,  // 0: go8.devcloud.mcenter.DescribeUserRequest.describe_by:type_name -> go8.devcloud.mcenter.DESCRIBE_BY
	7,  // 1: go8.devcloud.mcenter.QueryUserRequest.page:type_name -> infraboard.mcube.page.PageRequest
	8,  // 2: go8.devcloud.mcenter.QueryUserRequest.provider:type_name -> go8.devcloud.mcenter.PROVIDER
	9,  // 3: go8.devcloud.mcenter.QueryUserRequest.type:type_name -> go8.devcloud.mcenter.TYPE
	10, // 4: go8.devcloud.mcenter.UpdateUserRequest.update_mode:type_name -> infraboard.mcube.request.UpdateMode
	11, // 5: go8.devcloud.mcenter.UpdateUserRequest.profile:type_name -> go8.devcloud.mcenter.Profile
	1,  // 6: go8.devcloud.mcenter.RPC.DescribeUser:input_type -> go8.devcloud.mcenter.DescribeUserRequest
	2,  // 7: go8.devcloud.mcenter.RPC.QueryUser:input_type -> go8.devcloud.mcenter.QueryUserRequest
	12, // 8: go8.devcloud.mcenter.RPC.DescribeUser:output_type -> go8.devcloud.mcenter.User
	13, // 9: go8.devcloud.mcenter.RPC.QueryUser:output_type -> go8.devcloud.mcenter.UserSet
	8,  // [8:10] is the sub-list for method output_type
	6,  // [6:8] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_apps_user_pb_rpc_proto_init() }
func file_apps_user_pb_rpc_proto_init() {
	if File_apps_user_pb_rpc_proto != nil {
		return
	}
	file_apps_user_pb_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_user_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeUserRequest); i {
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
		file_apps_user_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserRequest); i {
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
		file_apps_user_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePasswordRequest); i {
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
		file_apps_user_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordRequest); i {
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
		file_apps_user_pb_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
		file_apps_user_pb_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
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
	file_apps_user_pb_rpc_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_user_pb_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_user_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_user_pb_rpc_proto_depIdxs,
		EnumInfos:         file_apps_user_pb_rpc_proto_enumTypes,
		MessageInfos:      file_apps_user_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_user_pb_rpc_proto = out.File
	file_apps_user_pb_rpc_proto_rawDesc = nil
	file_apps_user_pb_rpc_proto_goTypes = nil
	file_apps_user_pb_rpc_proto_depIdxs = nil
}
