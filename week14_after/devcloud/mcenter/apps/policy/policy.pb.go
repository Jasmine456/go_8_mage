// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/policy/pb/policy.proto

package policy

import (
	role "github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/role"
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

type PolicyType int32

const (
	// CustomPolicy (custom) 用户自己定义的策略
	PolicyType_CUSTOM PolicyType = 0
	// BuildInPolicy (build_in) 系统内部逻辑, 不允许用户看到并修改
	PolicyType_BUILD_IN PolicyType = 1
)

// Enum value maps for PolicyType.
var (
	PolicyType_name = map[int32]string{
		0: "CUSTOM",
		1: "BUILD_IN",
	}
	PolicyType_value = map[string]int32{
		"CUSTOM":   0,
		"BUILD_IN": 1,
	}
)

func (x PolicyType) Enum() *PolicyType {
	p := new(PolicyType)
	*p = x
	return p
}

func (x PolicyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PolicyType) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_policy_pb_policy_proto_enumTypes[0].Descriptor()
}

func (PolicyType) Type() protoreflect.EnumType {
	return &file_apps_policy_pb_policy_proto_enumTypes[0]
}

func (x PolicyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PolicyType.Descriptor instead.
func (PolicyType) EnumDescriptor() ([]byte, []int) {
	return file_apps_policy_pb_policy_proto_rawDescGZIP(), []int{0}
}

// Policy 权限策略
type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 策略ID
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 策略定义
	// @gotags: bson:"sepc" json:"spec"
	Spec *CreatePolicyRequest `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec" bson:"sepc"`
	// 关联的角色对象
	// @gotags: bson:"-" json:"role,omitempty"
	Role *role.Role `protobuf:"bytes,13,opt,name=role,proto3" json:"role,omitempty" bson:"-"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_policy_pb_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_apps_policy_pb_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_apps_policy_pb_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Policy) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Policy) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Policy) GetSpec() *CreatePolicyRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Policy) GetRole() *role.Role {
	if x != nil {
		return x.Role
	}
	return nil
}

// CreatePolicyRequest 创建策略的请求
type CreatePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 创建者
	// @gotags: bson:"create_by" json:"create_by"
	CreateBy string `protobuf:"bytes,1,opt,name=create_by,json=createBy,proto3" json:"create_by" bson:"create_by"`
	// 策略所属域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 范围
	// @gotags: bson:"namespace" json:"namespace" validate:"lte=120"
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace" bson:"namespace" validate:"lte=120"`
	// 比如namepsace更小的维度, 比如 你namespace还有很多小的组
	// 研发云这个项目, group 用于控制环境(env)
	// @gotags: bson:"group" json:"group"
	Group string `protobuf:"bytes,4,opt,name=group,proto3" json:"group" bson:"group"`
	// 范围控制
	// @gotags: bson:"scope" json:"scope"
	Scope string `protobuf:"bytes,5,opt,name=scope,proto3" json:"scope" bson:"scope"`
	// 用户
	// @gotags: bson:"username" json:"username" validate:"required,lte=120"
	Username string `protobuf:"bytes,6,opt,name=username,proto3" json:"username" bson:"username" validate:"required,lte=120"`
	// 角色名称
	// @gotags: bson:"role_id" json:"role_id" validate:"required,lte=40"
	RoleId string `protobuf:"bytes,7,opt,name=role_id,json=roleId,proto3" json:"role_id" bson:"role_id" validate:"required,lte=40"`
	// 策略过期时间
	// @gotags: bson:"expired_time" json:"expired_time"
	ExpiredTime int64 `protobuf:"varint,8,opt,name=expired_time,json=expiredTime,proto3" json:"expired_time" bson:"expired_time"`
	// 策略的类型
	// @gotags: bson:"type" json:"type"
	Type PolicyType `protobuf:"varint,9,opt,name=type,proto3,enum=go8.devcloud.mcenter.policy.PolicyType" json:"type" bson:"type"`
}

func (x *CreatePolicyRequest) Reset() {
	*x = CreatePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_policy_pb_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyRequest) ProtoMessage() {}

func (x *CreatePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_policy_pb_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyRequest.ProtoReflect.Descriptor instead.
func (*CreatePolicyRequest) Descriptor() ([]byte, []int) {
	return file_apps_policy_pb_policy_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePolicyRequest) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *CreatePolicyRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreatePolicyRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreatePolicyRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *CreatePolicyRequest) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *CreatePolicyRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreatePolicyRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *CreatePolicyRequest) GetExpiredTime() int64 {
	if x != nil {
		return x.ExpiredTime
	}
	return 0
}

func (x *CreatePolicyRequest) GetType() PolicyType {
	if x != nil {
		return x.Type
	}
	return PolicyType_CUSTOM
}

type PolicySet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*Policy `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *PolicySet) Reset() {
	*x = PolicySet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_policy_pb_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicySet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicySet) ProtoMessage() {}

func (x *PolicySet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_policy_pb_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicySet.ProtoReflect.Descriptor instead.
func (*PolicySet) Descriptor() ([]byte, []int) {
	return file_apps_policy_pb_policy_proto_rawDescGZIP(), []int{2}
}

func (x *PolicySet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PolicySet) GetItems() []*Policy {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_policy_pb_policy_proto protoreflect.FileDescriptor

var file_apps_policy_pb_policy_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x62,
	0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67,
	0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x17, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x44, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x33,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67,
	0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x22, 0xa9, 0x02, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x27, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x5c, 0x0a, 0x09, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x39, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x38, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2a, 0x26, 0x0a,
	0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43,
	0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x49, 0x4c, 0x44,
	0x5f, 0x49, 0x4e, 0x10, 0x01, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x61, 0x73, 0x6d, 0x69, 0x6e, 0x65, 0x34, 0x35, 0x36, 0x2f, 0x67,
	0x6f, 0x5f, 0x38, 0x5f, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x77, 0x65, 0x65, 0x6b, 0x31, 0x34, 0x5f,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_policy_pb_policy_proto_rawDescOnce sync.Once
	file_apps_policy_pb_policy_proto_rawDescData = file_apps_policy_pb_policy_proto_rawDesc
)

func file_apps_policy_pb_policy_proto_rawDescGZIP() []byte {
	file_apps_policy_pb_policy_proto_rawDescOnce.Do(func() {
		file_apps_policy_pb_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_policy_pb_policy_proto_rawDescData)
	})
	return file_apps_policy_pb_policy_proto_rawDescData
}

var file_apps_policy_pb_policy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_policy_pb_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_policy_pb_policy_proto_goTypes = []interface{}{
	(PolicyType)(0),             // 0: go8.devcloud.mcenter.policy.PolicyType
	(*Policy)(nil),              // 1: go8.devcloud.mcenter.policy.Policy
	(*CreatePolicyRequest)(nil), // 2: go8.devcloud.mcenter.policy.CreatePolicyRequest
	(*PolicySet)(nil),           // 3: go8.devcloud.mcenter.policy.PolicySet
	(*role.Role)(nil),           // 4: go8.devcloud.mcenter.role.Role
}
var file_apps_policy_pb_policy_proto_depIdxs = []int32{
	2, // 0: go8.devcloud.mcenter.policy.Policy.spec:type_name -> go8.devcloud.mcenter.policy.CreatePolicyRequest
	4, // 1: go8.devcloud.mcenter.policy.Policy.role:type_name -> go8.devcloud.mcenter.role.Role
	0, // 2: go8.devcloud.mcenter.policy.CreatePolicyRequest.type:type_name -> go8.devcloud.mcenter.policy.PolicyType
	1, // 3: go8.devcloud.mcenter.policy.PolicySet.items:type_name -> go8.devcloud.mcenter.policy.Policy
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_apps_policy_pb_policy_proto_init() }
func file_apps_policy_pb_policy_proto_init() {
	if File_apps_policy_pb_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_policy_pb_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_apps_policy_pb_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePolicyRequest); i {
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
		file_apps_policy_pb_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicySet); i {
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
			RawDescriptor: file_apps_policy_pb_policy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_policy_pb_policy_proto_goTypes,
		DependencyIndexes: file_apps_policy_pb_policy_proto_depIdxs,
		EnumInfos:         file_apps_policy_pb_policy_proto_enumTypes,
		MessageInfos:      file_apps_policy_pb_policy_proto_msgTypes,
	}.Build()
	File_apps_policy_pb_policy_proto = out.File
	file_apps_policy_pb_policy_proto_rawDesc = nil
	file_apps_policy_pb_policy_proto_goTypes = nil
	file_apps_policy_pb_policy_proto_depIdxs = nil
}