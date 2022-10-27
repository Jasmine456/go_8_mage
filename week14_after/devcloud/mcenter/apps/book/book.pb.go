// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/book/pb/book.proto

package book

import (
	"github.com/infraboard/mcube/http/request"
	//request "github.com/infraboard/mcube/http/request"
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

// Book todo
type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 唯一ID
	// @gotags: json:"id" bson:"_id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 录入时间
	// @gotags: json:"create_at" bson:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: json:"update_at" bson:"update_at"
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 更新人
	// @gotags: json:"update_by" bson:"update_by"
	UpdateBy string `protobuf:"bytes,4,opt,name=update_by,json=updateBy,proto3" json:"update_by" bson:"update_by"`
	// 书本信息
	// @gotags: json:"data" bson:"data"
	Data *CreateBookRequest `protobuf:"bytes,5,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_book_pb_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_apps_book_pb_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_apps_book_pb_book_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Book) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Book) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Book) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *Book) GetData() *CreateBookRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 创建人
	// @gotags: json:"create_by" bson:"create_by"
	CreateBy string `protobuf:"bytes,1,opt,name=create_by,json=createBy,proto3" json:"create_by" bson:"create_by"`
	// 名称
	// @gotags: json:"name" bson:"name" validate:"required"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" bson:"name" validate:"required"`
	// 作者
	// @gotags: json:"author" bson:"author" validate:"required"
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author" bson:"author" validate:"required"`
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_book_pb_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_book_pb_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_apps_book_pb_book_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBookRequest) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *CreateBookRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type QueryBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 关键字参数
	// @gotags: json:"keywords"
	Keywords string `protobuf:"bytes,2,opt,name=keywords,proto3" json:"keywords"`
}

func (x *QueryBookRequest) Reset() {
	*x = QueryBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_book_pb_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBookRequest) ProtoMessage() {}

func (x *QueryBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_book_pb_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBookRequest.ProtoReflect.Descriptor instead.
func (*QueryBookRequest) Descriptor() ([]byte, []int) {
	return file_apps_book_pb_book_proto_rawDescGZIP(), []int{2}
}

func (x *QueryBookRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryBookRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

// BookSet todo
type BookSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页时，返回总数量
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 一页的数据
	// @gotags: json:"items"
	Items []*Book `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *BookSet) Reset() {
	*x = BookSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_book_pb_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookSet) ProtoMessage() {}

func (x *BookSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_book_pb_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookSet.ProtoReflect.Descriptor instead.
func (*BookSet) Descriptor() ([]byte, []int) {
	return file_apps_book_pb_book_proto_rawDescGZIP(), []int{3}
}

func (x *BookSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *BookSet) GetItems() []*Book {
	if x != nil {
		return x.Items
	}
	return nil
}

type DescribeBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// book id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribeBookRequest) Reset() {
	*x = DescribeBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_book_pb_book_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeBookRequest) ProtoMessage() {}

func (x *DescribeBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_book_pb_book_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeBookRequest.ProtoReflect.Descriptor instead.
func (*DescribeBookRequest) Descriptor() ([]byte, []int) {
	return file_apps_book_pb_book_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeBookRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// book id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 更新模式
	// @gotags: json:"update_mode"
	UpdateMode request1.UpdateMode `protobuf:"varint,2,opt,name=update_mode,json=updateMode,proto3,enum=infraboard.mcube.request.UpdateMode" json:"update_mode"`
	// 更新人
	// @gotags: json:"update_by"
	UpdateBy string `protobuf:"bytes,3,opt,name=update_by,json=updateBy,proto3" json:"update_by"`
	// 更新时间
	// @gotags: json:"update_at"
	UpdateAt int64 `protobuf:"varint,4,opt,name=update_at,json=updateAt,proto3" json:"update_at"`
	// 更新的书本信息
	// @gotags: json:"data"
	Data *CreateBookRequest `protobuf:"bytes,5,opt,name=data,proto3" json:"data"`
}

func (x *UpdateBookRequest) Reset() {
	*x = UpdateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_book_pb_book_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRequest) ProtoMessage() {}

func (x *UpdateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_book_pb_book_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return file_apps_book_pb_book_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBookRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBookRequest) GetUpdateMode() request1.UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return request1.UpdateMode(0)
}

func (x *UpdateBookRequest) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *UpdateBookRequest) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *UpdateBookRequest) GetData() *CreateBookRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// book id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DeleteBookRequest) Reset() {
	*x = DeleteBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_book_pb_book_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequest) ProtoMessage() {}

func (x *DeleteBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_book_pb_book_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return file_apps_book_pb_book_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBookRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_apps_book_pb_book_proto protoreflect.FileDescriptor

var file_apps_book_pb_book_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x65, 0x6d, 0x6f, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x04, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5c, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x66, 0x0a, 0x10, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x22, 0x46, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xd6, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32,
	0xbf, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x3c, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1b, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1e, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x3b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x3b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4a, 0x61, 0x73, 0x6d, 0x69, 0x6e, 0x65, 0x34, 0x35, 0x36, 0x2f, 0x67, 0x6f, 0x5f, 0x38, 0x5f,
	0x6d, 0x61, 0x67, 0x65, 0x2f, 0x77, 0x65, 0x65, 0x6b, 0x31, 0x34, 0x5f, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_book_pb_book_proto_rawDescOnce sync.Once
	file_apps_book_pb_book_proto_rawDescData = file_apps_book_pb_book_proto_rawDesc
)

func file_apps_book_pb_book_proto_rawDescGZIP() []byte {
	file_apps_book_pb_book_proto_rawDescOnce.Do(func() {
		file_apps_book_pb_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_book_pb_book_proto_rawDescData)
	})
	return file_apps_book_pb_book_proto_rawDescData
}

var file_apps_book_pb_book_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_apps_book_pb_book_proto_goTypes = []interface{}{
	(*Book)(nil),                // 0: demo.book.Book
	(*CreateBookRequest)(nil),   // 1: demo.book.CreateBookRequest
	(*QueryBookRequest)(nil),    // 2: demo.book.QueryBookRequest
	(*BookSet)(nil),             // 3: demo.book.BookSet
	(*DescribeBookRequest)(nil), // 4: demo.book.DescribeBookRequest
	(*UpdateBookRequest)(nil),   // 5: demo.book.UpdateBookRequest
	(*DeleteBookRequest)(nil),   // 6: demo.book.DeleteBookRequest
	(*request.PageRequest)(nil), // 7: infraboard.mcube.page.PageRequest
	(request1.UpdateMode)(0),    // 8: infraboard.mcube.request.UpdateMode
}
var file_apps_book_pb_book_proto_depIdxs = []int32{
	1,  // 0: demo.book.Book.data:type_name -> demo.book.CreateBookRequest
	7,  // 1: demo.book.QueryBookRequest.page:type_name -> infraboard.mcube.page.PageRequest
	0,  // 2: demo.book.BookSet.items:type_name -> demo.book.Book
	8,  // 3: demo.book.UpdateBookRequest.update_mode:type_name -> infraboard.mcube.request.UpdateMode
	1,  // 4: demo.book.UpdateBookRequest.data:type_name -> demo.book.CreateBookRequest
	1,  // 5: demo.book.Service.CreateBook:input_type -> demo.book.CreateBookRequest
	2,  // 6: demo.book.Service.QueryBook:input_type -> demo.book.QueryBookRequest
	4,  // 7: demo.book.Service.DescribeBook:input_type -> demo.book.DescribeBookRequest
	5,  // 8: demo.book.Service.UpdateBook:input_type -> demo.book.UpdateBookRequest
	6,  // 9: demo.book.Service.DeleteBook:input_type -> demo.book.DeleteBookRequest
	0,  // 10: demo.book.Service.CreateBook:output_type -> demo.book.Book
	3,  // 11: demo.book.Service.QueryBook:output_type -> demo.book.BookSet
	0,  // 12: demo.book.Service.DescribeBook:output_type -> demo.book.Book
	0,  // 13: demo.book.Service.UpdateBook:output_type -> demo.book.Book
	0,  // 14: demo.book.Service.DeleteBook:output_type -> demo.book.Book
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_apps_book_pb_book_proto_init() }
func file_apps_book_pb_book_proto_init() {
	if File_apps_book_pb_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_book_pb_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_apps_book_pb_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookRequest); i {
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
		file_apps_book_pb_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBookRequest); i {
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
		file_apps_book_pb_book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookSet); i {
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
		file_apps_book_pb_book_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeBookRequest); i {
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
		file_apps_book_pb_book_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookRequest); i {
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
		file_apps_book_pb_book_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookRequest); i {
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
			RawDescriptor: file_apps_book_pb_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_book_pb_book_proto_goTypes,
		DependencyIndexes: file_apps_book_pb_book_proto_depIdxs,
		MessageInfos:      file_apps_book_pb_book_proto_msgTypes,
	}.Build()
	File_apps_book_pb_book_proto = out.File
	file_apps_book_pb_book_proto_rawDesc = nil
	file_apps_book_pb_book_proto_goTypes = nil
	file_apps_book_pb_book_proto_depIdxs = nil
}
