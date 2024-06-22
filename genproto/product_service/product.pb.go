// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: product.proto

package product_service

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

type ProductPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProductPrimaryKey) Reset() {
	*x = ProductPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPrimaryKey) ProtoMessage() {}

func (x *ProductPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPrimaryKey.ProtoReflect.Descriptor instead.
func (*ProductPrimaryKey) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NameUz          string  `protobuf:"bytes,1,opt,name=name_uz,json=nameUz,proto3" json:"name_uz,omitempty"`
	NameRu          string  `protobuf:"bytes,2,opt,name=name_ru,json=nameRu,proto3" json:"name_ru,omitempty"`
	NameEn          string  `protobuf:"bytes,3,opt,name=name_en,json=nameEn,proto3" json:"name_en,omitempty"`
	DescriptionUz   string  `protobuf:"bytes,4,opt,name=description_uz,json=descriptionUz,proto3" json:"description_uz,omitempty"`
	DescriptionRu   string  `protobuf:"bytes,5,opt,name=description_ru,json=descriptionRu,proto3" json:"description_ru,omitempty"`
	DescriptionEn   string  `protobuf:"bytes,6,opt,name=description_en,json=descriptionEn,proto3" json:"description_en,omitempty"`
	Active          bool    `protobuf:"varint,7,opt,name=active,proto3" json:"active,omitempty"`
	OrderNo         int32   `protobuf:"varint,8,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	InPrice         float32 `protobuf:"fixed32,9,opt,name=in_price,json=inPrice,proto3" json:"in_price,omitempty"`
	OutPrice        float32 `protobuf:"fixed32,10,opt,name=out_price,json=outPrice,proto3" json:"out_price,omitempty"`
	LeftCount       int32   `protobuf:"varint,11,opt,name=left_count,json=leftCount,proto3" json:"left_count,omitempty"`
	DiscountPercent float32 `protobuf:"fixed32,12,opt,name=discount_percent,json=discountPercent,proto3" json:"discount_percent,omitempty"`
	CategoryId      string  `protobuf:"bytes,13,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *CreateProduct) Reset() {
	*x = CreateProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProduct) ProtoMessage() {}

func (x *CreateProduct) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProduct.ProtoReflect.Descriptor instead.
func (*CreateProduct) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProduct) GetNameUz() string {
	if x != nil {
		return x.NameUz
	}
	return ""
}

func (x *CreateProduct) GetNameRu() string {
	if x != nil {
		return x.NameRu
	}
	return ""
}

func (x *CreateProduct) GetNameEn() string {
	if x != nil {
		return x.NameEn
	}
	return ""
}

func (x *CreateProduct) GetDescriptionUz() string {
	if x != nil {
		return x.DescriptionUz
	}
	return ""
}

func (x *CreateProduct) GetDescriptionRu() string {
	if x != nil {
		return x.DescriptionRu
	}
	return ""
}

func (x *CreateProduct) GetDescriptionEn() string {
	if x != nil {
		return x.DescriptionEn
	}
	return ""
}

func (x *CreateProduct) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *CreateProduct) GetOrderNo() int32 {
	if x != nil {
		return x.OrderNo
	}
	return 0
}

func (x *CreateProduct) GetInPrice() float32 {
	if x != nil {
		return x.InPrice
	}
	return 0
}

func (x *CreateProduct) GetOutPrice() float32 {
	if x != nil {
		return x.OutPrice
	}
	return 0
}

func (x *CreateProduct) GetLeftCount() int32 {
	if x != nil {
		return x.LeftCount
	}
	return 0
}

func (x *CreateProduct) GetDiscountPercent() float32 {
	if x != nil {
		return x.DiscountPercent
	}
	return 0
}

func (x *CreateProduct) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type GetProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug            string  `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	NameUz          string  `protobuf:"bytes,3,opt,name=name_uz,json=nameUz,proto3" json:"name_uz,omitempty"`
	NameRu          string  `protobuf:"bytes,4,opt,name=name_ru,json=nameRu,proto3" json:"name_ru,omitempty"`
	NameEn          string  `protobuf:"bytes,5,opt,name=name_en,json=nameEn,proto3" json:"name_en,omitempty"`
	DescriptionUz   string  `protobuf:"bytes,6,opt,name=description_uz,json=descriptionUz,proto3" json:"description_uz,omitempty"`
	DescriptionRu   string  `protobuf:"bytes,7,opt,name=description_ru,json=descriptionRu,proto3" json:"description_ru,omitempty"`
	DescriptionEn   string  `protobuf:"bytes,8,opt,name=description_en,json=descriptionEn,proto3" json:"description_en,omitempty"`
	Active          bool    `protobuf:"varint,9,opt,name=active,proto3" json:"active,omitempty"`
	OrderNo         int32   `protobuf:"varint,10,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	InPrice         float32 `protobuf:"fixed32,11,opt,name=in_price,json=inPrice,proto3" json:"in_price,omitempty"`
	OutPrice        float32 `protobuf:"fixed32,12,opt,name=out_price,json=outPrice,proto3" json:"out_price,omitempty"`
	LeftCount       int32   `protobuf:"varint,13,opt,name=left_count,json=leftCount,proto3" json:"left_count,omitempty"`
	DiscountPercent float32 `protobuf:"fixed32,14,opt,name=discount_percent,json=discountPercent,proto3" json:"discount_percent,omitempty"`
	CreatedAt       string  `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       string  `protobuf:"bytes,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GetProduct) Reset() {
	*x = GetProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProduct) ProtoMessage() {}

func (x *GetProduct) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProduct.ProtoReflect.Descriptor instead.
func (*GetProduct) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *GetProduct) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetProduct) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *GetProduct) GetNameUz() string {
	if x != nil {
		return x.NameUz
	}
	return ""
}

func (x *GetProduct) GetNameRu() string {
	if x != nil {
		return x.NameRu
	}
	return ""
}

func (x *GetProduct) GetNameEn() string {
	if x != nil {
		return x.NameEn
	}
	return ""
}

func (x *GetProduct) GetDescriptionUz() string {
	if x != nil {
		return x.DescriptionUz
	}
	return ""
}

func (x *GetProduct) GetDescriptionRu() string {
	if x != nil {
		return x.DescriptionRu
	}
	return ""
}

func (x *GetProduct) GetDescriptionEn() string {
	if x != nil {
		return x.DescriptionEn
	}
	return ""
}

func (x *GetProduct) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *GetProduct) GetOrderNo() int32 {
	if x != nil {
		return x.OrderNo
	}
	return 0
}

func (x *GetProduct) GetInPrice() float32 {
	if x != nil {
		return x.InPrice
	}
	return 0
}

func (x *GetProduct) GetOutPrice() float32 {
	if x != nil {
		return x.OutPrice
	}
	return 0
}

func (x *GetProduct) GetLeftCount() int32 {
	if x != nil {
		return x.LeftCount
	}
	return 0
}

func (x *GetProduct) GetDiscountPercent() float32 {
	if x != nil {
		return x.DiscountPercent
	}
	return 0
}

func (x *GetProduct) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetProduct) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UpdateProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NameUz          string  `protobuf:"bytes,1,opt,name=name_uz,json=nameUz,proto3" json:"name_uz,omitempty"`
	NameRu          string  `protobuf:"bytes,2,opt,name=name_ru,json=nameRu,proto3" json:"name_ru,omitempty"`
	NameEn          string  `protobuf:"bytes,3,opt,name=name_en,json=nameEn,proto3" json:"name_en,omitempty"`
	DescriptionUz   string  `protobuf:"bytes,4,opt,name=description_uz,json=descriptionUz,proto3" json:"description_uz,omitempty"`
	DescriptionRu   string  `protobuf:"bytes,5,opt,name=description_ru,json=descriptionRu,proto3" json:"description_ru,omitempty"`
	DescriptionEn   string  `protobuf:"bytes,6,opt,name=description_en,json=descriptionEn,proto3" json:"description_en,omitempty"`
	Active          bool    `protobuf:"varint,7,opt,name=active,proto3" json:"active,omitempty"`
	OrderNo         int32   `protobuf:"varint,8,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	InPrice         float32 `protobuf:"fixed32,9,opt,name=in_price,json=inPrice,proto3" json:"in_price,omitempty"`
	OutPrice        float32 `protobuf:"fixed32,10,opt,name=out_price,json=outPrice,proto3" json:"out_price,omitempty"`
	LeftCount       int32   `protobuf:"varint,11,opt,name=left_count,json=leftCount,proto3" json:"left_count,omitempty"`
	DiscountPercent float32 `protobuf:"fixed32,12,opt,name=discount_percent,json=discountPercent,proto3" json:"discount_percent,omitempty"`
	Id              string  `protobuf:"bytes,13,opt,name=id,proto3" json:"id,omitempty"`
	CategoryId      string  `protobuf:"bytes,14,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *UpdateProduct) Reset() {
	*x = UpdateProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProduct) ProtoMessage() {}

func (x *UpdateProduct) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProduct.ProtoReflect.Descriptor instead.
func (*UpdateProduct) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProduct) GetNameUz() string {
	if x != nil {
		return x.NameUz
	}
	return ""
}

func (x *UpdateProduct) GetNameRu() string {
	if x != nil {
		return x.NameRu
	}
	return ""
}

func (x *UpdateProduct) GetNameEn() string {
	if x != nil {
		return x.NameEn
	}
	return ""
}

func (x *UpdateProduct) GetDescriptionUz() string {
	if x != nil {
		return x.DescriptionUz
	}
	return ""
}

func (x *UpdateProduct) GetDescriptionRu() string {
	if x != nil {
		return x.DescriptionRu
	}
	return ""
}

func (x *UpdateProduct) GetDescriptionEn() string {
	if x != nil {
		return x.DescriptionEn
	}
	return ""
}

func (x *UpdateProduct) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *UpdateProduct) GetOrderNo() int32 {
	if x != nil {
		return x.OrderNo
	}
	return 0
}

func (x *UpdateProduct) GetInPrice() float32 {
	if x != nil {
		return x.InPrice
	}
	return 0
}

func (x *UpdateProduct) GetOutPrice() float32 {
	if x != nil {
		return x.OutPrice
	}
	return 0
}

func (x *UpdateProduct) GetLeftCount() int32 {
	if x != nil {
		return x.LeftCount
	}
	return 0
}

func (x *UpdateProduct) GetDiscountPercent() float32 {
	if x != nil {
		return x.DiscountPercent
	}
	return 0
}

func (x *UpdateProduct) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProduct) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type GetListProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListProductRequest) Reset() {
	*x = GetListProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListProductRequest) ProtoMessage() {}

func (x *GetListProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListProductRequest.ProtoReflect.Descriptor instead.
func (*GetListProductRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{4}
}

func (x *GetListProductRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListProductRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListProductRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    int64         `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	Products []*GetProduct `protobuf:"bytes,2,rep,name=Products,proto3" json:"Products,omitempty"`
}

func (x *GetListProductResponse) Reset() {
	*x = GetListProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListProductResponse) ProtoMessage() {}

func (x *GetListProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListProductResponse.ProtoReflect.Descriptor instead.
func (*GetListProductResponse) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{5}
}

func (x *GetListProductResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListProductResponse) GetProducts() []*GetProduct {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x67, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x23, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa5, 0x03, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x5f,
	0x75, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x55, 0x7a,
	0x12, 0x17, 0x0a, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x75, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65,
	0x45, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x75, 0x7a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x7a, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75,
	0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x69, 0x6e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x65, 0x66, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0xe3, 0x03,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67,
	0x12, 0x17, 0x0a, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x75, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x55, 0x7a, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x72, 0x75, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x75, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x7a, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x55, 0x7a, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x75, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x65, 0x66, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x6c, 0x65, 0x66, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xb5, 0x03, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x75, 0x7a,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x55, 0x7a, 0x12, 0x17,
	0x0a, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x75, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x5f,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x45, 0x6e,
	0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x75, 0x7a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x7a, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x12, 0x25,
	0x0a, 0x0e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x69, 0x6e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x65, 0x66, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x29, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x6a, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0xb1, 0x03, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x1e, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12,
	0x49, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_product_proto_goTypes = []interface{}{
	(*ProductPrimaryKey)(nil),      // 0: product_service_go.ProductPrimaryKey
	(*CreateProduct)(nil),          // 1: product_service_go.CreateProduct
	(*GetProduct)(nil),             // 2: product_service_go.GetProduct
	(*UpdateProduct)(nil),          // 3: product_service_go.UpdateProduct
	(*GetListProductRequest)(nil),  // 4: product_service_go.GetListProductRequest
	(*GetListProductResponse)(nil), // 5: product_service_go.GetListProductResponse
	(*empty.Empty)(nil),            // 6: google.protobuf.Empty
}
var file_product_proto_depIdxs = []int32{
	2, // 0: product_service_go.GetListProductResponse.Products:type_name -> product_service_go.GetProduct
	1, // 1: product_service_go.ProductService.Create:input_type -> product_service_go.CreateProduct
	0, // 2: product_service_go.ProductService.GetByID:input_type -> product_service_go.ProductPrimaryKey
	4, // 3: product_service_go.ProductService.GetList:input_type -> product_service_go.GetListProductRequest
	3, // 4: product_service_go.ProductService.Update:input_type -> product_service_go.UpdateProduct
	0, // 5: product_service_go.ProductService.Delete:input_type -> product_service_go.ProductPrimaryKey
	2, // 6: product_service_go.ProductService.Create:output_type -> product_service_go.GetProduct
	2, // 7: product_service_go.ProductService.GetByID:output_type -> product_service_go.GetProduct
	5, // 8: product_service_go.ProductService.GetList:output_type -> product_service_go.GetListProductResponse
	2, // 9: product_service_go.ProductService.Update:output_type -> product_service_go.GetProduct
	6, // 10: product_service_go.ProductService.Delete:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPrimaryKey); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProduct); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProduct); i {
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
		file_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProduct); i {
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
		file_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListProductRequest); i {
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
		file_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListProductResponse); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
