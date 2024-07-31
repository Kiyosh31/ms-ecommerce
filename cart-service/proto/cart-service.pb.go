// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: cart-service/proto/cart-service.proto

package cartPb

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

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId   string     `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Total    float64    `protobuf:"fixed64,3,opt,name=total,proto3" json:"total,omitempty"`
	Products []*Product `protobuf:"bytes,4,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_service_proto_cart_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_cart_service_proto_cart_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_cart_service_proto_cart_service_proto_rawDescGZIP(), []int{0}
}

func (x *Cart) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cart) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Cart) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Cart) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string  `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Quantity  float32 `protobuf:"fixed32,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price     float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_service_proto_cart_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_cart_service_proto_cart_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_cart_service_proto_cart_service_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Product) GetQuantity() float32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type SingleCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Cart    *Cart  `protobuf:"bytes,2,opt,name=cart,proto3" json:"cart,omitempty"`
}

func (x *SingleCartResponse) Reset() {
	*x = SingleCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_service_proto_cart_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleCartResponse) ProtoMessage() {}

func (x *SingleCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_service_proto_cart_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleCartResponse.ProtoReflect.Descriptor instead.
func (*SingleCartResponse) Descriptor() ([]byte, []int) {
	return file_cart_service_proto_cart_service_proto_rawDescGZIP(), []int{2}
}

func (x *SingleCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SingleCartResponse) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type MultipleCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Cart    []*Cart `protobuf:"bytes,2,rep,name=cart,proto3" json:"cart,omitempty"`
}

func (x *MultipleCartResponse) Reset() {
	*x = MultipleCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_service_proto_cart_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleCartResponse) ProtoMessage() {}

func (x *MultipleCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_service_proto_cart_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleCartResponse.ProtoReflect.Descriptor instead.
func (*MultipleCartResponse) Descriptor() ([]byte, []int) {
	return file_cart_service_proto_cart_service_proto_rawDescGZIP(), []int{3}
}

func (x *MultipleCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MultipleCartResponse) GetCart() []*Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type CreateCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart *Cart `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
}

func (x *CreateCartRequest) Reset() {
	*x = CreateCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_service_proto_cart_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCartRequest) ProtoMessage() {}

func (x *CreateCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_service_proto_cart_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCartRequest.ProtoReflect.Descriptor instead.
func (*CreateCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_service_proto_cart_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCartRequest) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type GetCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	CartId string `protobuf:"bytes,2,opt,name=cartId,proto3" json:"cartId,omitempty"`
}

func (x *GetCartRequest) Reset() {
	*x = GetCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_service_proto_cart_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartRequest) ProtoMessage() {}

func (x *GetCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_service_proto_cart_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartRequest.ProtoReflect.Descriptor instead.
func (*GetCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_service_proto_cart_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetCartRequest) GetCartId() string {
	if x != nil {
		return x.CartId
	}
	return ""
}

type GetCartsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetCartsRequest) Reset() {
	*x = GetCartsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_service_proto_cart_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartsRequest) ProtoMessage() {}

func (x *GetCartsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_service_proto_cart_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartsRequest.ProtoReflect.Descriptor instead.
func (*GetCartsRequest) Descriptor() ([]byte, []int) {
	return file_cart_service_proto_cart_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetCartsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_cart_service_proto_cart_service_proto protoreflect.FileDescriptor

var file_cart_service_proto_cart_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x61, 0x72, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x61, 0x72, 0x74, 0x50, 0x62, 0x22,
	0x71, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x50, 0x62,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x22, 0x59, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x50, 0x0a,
	0x12, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a,
	0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x50, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22,
	0x52, 0x0a, 0x14, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x20, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x50, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63,
	0x61, 0x72, 0x74, 0x22, 0x35, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x50, 0x62, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x40, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xd7, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x50, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x50, 0x62, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x50, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x50, 0x62, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x50, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x50, 0x62, 0x2e, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4b, 0x69, 0x79, 0x6f, 0x73, 0x68, 0x33, 0x31, 0x2f, 0x6d, 0x73, 0x2d, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cart_service_proto_cart_service_proto_rawDescOnce sync.Once
	file_cart_service_proto_cart_service_proto_rawDescData = file_cart_service_proto_cart_service_proto_rawDesc
)

func file_cart_service_proto_cart_service_proto_rawDescGZIP() []byte {
	file_cart_service_proto_cart_service_proto_rawDescOnce.Do(func() {
		file_cart_service_proto_cart_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_service_proto_cart_service_proto_rawDescData)
	})
	return file_cart_service_proto_cart_service_proto_rawDescData
}

var file_cart_service_proto_cart_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cart_service_proto_cart_service_proto_goTypes = []any{
	(*Cart)(nil),                 // 0: cartPb.Cart
	(*Product)(nil),              // 1: cartPb.Product
	(*SingleCartResponse)(nil),   // 2: cartPb.SingleCartResponse
	(*MultipleCartResponse)(nil), // 3: cartPb.MultipleCartResponse
	(*CreateCartRequest)(nil),    // 4: cartPb.CreateCartRequest
	(*GetCartRequest)(nil),       // 5: cartPb.GetCartRequest
	(*GetCartsRequest)(nil),      // 6: cartPb.GetCartsRequest
}
var file_cart_service_proto_cart_service_proto_depIdxs = []int32{
	1, // 0: cartPb.Cart.products:type_name -> cartPb.Product
	0, // 1: cartPb.SingleCartResponse.cart:type_name -> cartPb.Cart
	0, // 2: cartPb.MultipleCartResponse.cart:type_name -> cartPb.Cart
	0, // 3: cartPb.CreateCartRequest.cart:type_name -> cartPb.Cart
	4, // 4: cartPb.CartService.CreateCart:input_type -> cartPb.CreateCartRequest
	5, // 5: cartPb.CartService.GetCart:input_type -> cartPb.GetCartRequest
	6, // 6: cartPb.CartService.GetAllCarts:input_type -> cartPb.GetCartsRequest
	2, // 7: cartPb.CartService.CreateCart:output_type -> cartPb.SingleCartResponse
	2, // 8: cartPb.CartService.GetCart:output_type -> cartPb.SingleCartResponse
	3, // 9: cartPb.CartService.GetAllCarts:output_type -> cartPb.MultipleCartResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cart_service_proto_cart_service_proto_init() }
func file_cart_service_proto_cart_service_proto_init() {
	if File_cart_service_proto_cart_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cart_service_proto_cart_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Cart); i {
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
		file_cart_service_proto_cart_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Product); i {
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
		file_cart_service_proto_cart_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SingleCartResponse); i {
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
		file_cart_service_proto_cart_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MultipleCartResponse); i {
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
		file_cart_service_proto_cart_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCartRequest); i {
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
		file_cart_service_proto_cart_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetCartRequest); i {
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
		file_cart_service_proto_cart_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetCartsRequest); i {
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
			RawDescriptor: file_cart_service_proto_cart_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_service_proto_cart_service_proto_goTypes,
		DependencyIndexes: file_cart_service_proto_cart_service_proto_depIdxs,
		MessageInfos:      file_cart_service_proto_cart_service_proto_msgTypes,
	}.Build()
	File_cart_service_proto_cart_service_proto = out.File
	file_cart_service_proto_cart_service_proto_rawDesc = nil
	file_cart_service_proto_cart_service_proto_goTypes = nil
	file_cart_service_proto_cart_service_proto_depIdxs = nil
}
