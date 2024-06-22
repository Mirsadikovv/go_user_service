// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: seller.proto

package user_service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type SellerPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SellerPrimaryKey) Reset() {
	*x = SellerPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellerPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellerPrimaryKey) ProtoMessage() {}

func (x *SellerPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_seller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellerPrimaryKey.ProtoReflect.Descriptor instead.
func (*SellerPrimaryKey) Descriptor() ([]byte, []int) {
	return file_seller_proto_rawDescGZIP(), []int{0}
}

func (x *SellerPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SellerGmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gmail string `protobuf:"bytes,1,opt,name=gmail,proto3" json:"gmail,omitempty"`
}

func (x *SellerGmail) Reset() {
	*x = SellerGmail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellerGmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellerGmail) ProtoMessage() {}

func (x *SellerGmail) ProtoReflect() protoreflect.Message {
	mi := &file_seller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellerGmail.ProtoReflect.Descriptor instead.
func (*SellerGmail) Descriptor() ([]byte, []int) {
	return file_seller_proto_rawDescGZIP(), []int{1}
}

func (x *SellerGmail) GetGmail() string {
	if x != nil {
		return x.Gmail
	}
	return ""
}

type CreateSeller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone  string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Email  string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ShopId string `protobuf:"bytes,4,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *CreateSeller) Reset() {
	*x = CreateSeller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSeller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSeller) ProtoMessage() {}

func (x *CreateSeller) ProtoReflect() protoreflect.Message {
	mi := &file_seller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSeller.ProtoReflect.Descriptor instead.
func (*CreateSeller) Descriptor() ([]byte, []int) {
	return file_seller_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSeller) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateSeller) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateSeller) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSeller) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

type Seller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Phone     string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	ShopId    string `protobuf:"bytes,5,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	CreatedAt string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Seller) Reset() {
	*x = Seller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Seller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Seller) ProtoMessage() {}

func (x *Seller) ProtoReflect() protoreflect.Message {
	mi := &file_seller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Seller.ProtoReflect.Descriptor instead.
func (*Seller) Descriptor() ([]byte, []int) {
	return file_seller_proto_rawDescGZIP(), []int{3}
}

func (x *Seller) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Seller) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Seller) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Seller) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Seller) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

func (x *Seller) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Seller) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UpdateSeller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Phone  string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Email  string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Name   string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	ShopId string `protobuf:"bytes,5,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *UpdateSeller) Reset() {
	*x = UpdateSeller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSeller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSeller) ProtoMessage() {}

func (x *UpdateSeller) ProtoReflect() protoreflect.Message {
	mi := &file_seller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSeller.ProtoReflect.Descriptor instead.
func (*UpdateSeller) Descriptor() ([]byte, []int) {
	return file_seller_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSeller) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSeller) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateSeller) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateSeller) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateSeller) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

type GetSeller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Phone     string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	ShopId    string `protobuf:"bytes,5,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	CreatedAt string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GetSeller) Reset() {
	*x = GetSeller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSeller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSeller) ProtoMessage() {}

func (x *GetSeller) ProtoReflect() protoreflect.Message {
	mi := &file_seller_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSeller.ProtoReflect.Descriptor instead.
func (*GetSeller) Descriptor() ([]byte, []int) {
	return file_seller_proto_rawDescGZIP(), []int{5}
}

func (x *GetSeller) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetSeller) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetSeller) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetSeller) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetSeller) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

func (x *GetSeller) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetSeller) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetListSellerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit  uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListSellerRequest) Reset() {
	*x = GetListSellerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListSellerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListSellerRequest) ProtoMessage() {}

func (x *GetListSellerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seller_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListSellerRequest.ProtoReflect.Descriptor instead.
func (*GetListSellerRequest) Descriptor() ([]byte, []int) {
	return file_seller_proto_rawDescGZIP(), []int{6}
}

func (x *GetListSellerRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetListSellerRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListSellerRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListSellerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count   int64     `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Sellers []*Seller `protobuf:"bytes,2,rep,name=sellers,proto3" json:"sellers,omitempty"`
}

func (x *GetListSellerResponse) Reset() {
	*x = GetListSellerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seller_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListSellerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListSellerResponse) ProtoMessage() {}

func (x *GetListSellerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_seller_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListSellerResponse.ProtoReflect.Descriptor instead.
func (*GetListSellerResponse) Descriptor() ([]byte, []int) {
	return file_seller_proto_rawDescGZIP(), []int{7}
}

func (x *GetListSellerResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListSellerResponse) GetSellers() []*Seller {
	if x != nil {
		return x.Sellers
	}
	return nil
}

var File_seller_proto protoreflect.FileDescriptor

var file_seller_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a,
	0x0b, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x67, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0xaf, 0x01, 0x0a, 0x06,
	0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x77, 0x0a,
	0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0xb2, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x58, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x5d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x07, 0x73, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x73, 0x32, 0xb3, 0x03, 0x0a, 0x0d, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x1a, 0x14, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a,
	0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x49, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_seller_proto_rawDescOnce sync.Once
	file_seller_proto_rawDescData = file_seller_proto_rawDesc
)

func file_seller_proto_rawDescGZIP() []byte {
	file_seller_proto_rawDescOnce.Do(func() {
		file_seller_proto_rawDescData = protoimpl.X.CompressGZIP(file_seller_proto_rawDescData)
	})
	return file_seller_proto_rawDescData
}

var file_seller_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_seller_proto_goTypes = []interface{}{
	(*SellerPrimaryKey)(nil),      // 0: user_service.SellerPrimaryKey
	(*SellerGmail)(nil),           // 1: user_service.SellerGmail
	(*CreateSeller)(nil),          // 2: user_service.CreateSeller
	(*Seller)(nil),                // 3: user_service.Seller
	(*UpdateSeller)(nil),          // 4: user_service.UpdateSeller
	(*GetSeller)(nil),             // 5: user_service.GetSeller
	(*GetListSellerRequest)(nil),  // 6: user_service.GetListSellerRequest
	(*GetListSellerResponse)(nil), // 7: user_service.GetListSellerResponse
	(*empty.Empty)(nil),           // 8: google.protobuf.Empty
}
var file_seller_proto_depIdxs = []int32{
	3, // 0: user_service.GetListSellerResponse.sellers:type_name -> user_service.Seller
	2, // 1: user_service.SellerService.Create:input_type -> user_service.CreateSeller
	0, // 2: user_service.SellerService.GetByID:input_type -> user_service.SellerPrimaryKey
	6, // 3: user_service.SellerService.GetList:input_type -> user_service.GetListSellerRequest
	4, // 4: user_service.SellerService.Update:input_type -> user_service.UpdateSeller
	0, // 5: user_service.SellerService.Delete:input_type -> user_service.SellerPrimaryKey
	1, // 6: user_service.SellerService.GetByGmail:input_type -> user_service.SellerGmail
	3, // 7: user_service.SellerService.Create:output_type -> user_service.Seller
	3, // 8: user_service.SellerService.GetByID:output_type -> user_service.Seller
	7, // 9: user_service.SellerService.GetList:output_type -> user_service.GetListSellerResponse
	3, // 10: user_service.SellerService.Update:output_type -> user_service.Seller
	8, // 11: user_service.SellerService.Delete:output_type -> google.protobuf.Empty
	0, // 12: user_service.SellerService.GetByGmail:output_type -> user_service.SellerPrimaryKey
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_seller_proto_init() }
func file_seller_proto_init() {
	if File_seller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_seller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellerPrimaryKey); i {
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
		file_seller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellerGmail); i {
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
		file_seller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSeller); i {
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
		file_seller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Seller); i {
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
		file_seller_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSeller); i {
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
		file_seller_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSeller); i {
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
		file_seller_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListSellerRequest); i {
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
		file_seller_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListSellerResponse); i {
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
			RawDescriptor: file_seller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_seller_proto_goTypes,
		DependencyIndexes: file_seller_proto_depIdxs,
		MessageInfos:      file_seller_proto_msgTypes,
	}.Build()
	File_seller_proto = out.File
	file_seller_proto_rawDesc = nil
	file_seller_proto_goTypes = nil
	file_seller_proto_depIdxs = nil
}
