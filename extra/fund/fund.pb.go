// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: fund/fund.proto

package fund

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateFundV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string       `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Fund        *FundMessage `protobuf:"bytes,2,opt,name=fund,proto3" json:"fund,omitempty"`
}

func (x *CreateFundV1Request) Reset() {
	*x = CreateFundV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fund_fund_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFundV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFundV1Request) ProtoMessage() {}

func (x *CreateFundV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_fund_fund_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFundV1Request.ProtoReflect.Descriptor instead.
func (*CreateFundV1Request) Descriptor() ([]byte, []int) {
	return file_fund_fund_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFundV1Request) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *CreateFundV1Request) GetFund() *FundMessage {
	if x != nil {
		return x.Fund
	}
	return nil
}

type CreateFundV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateFundV1Response) Reset() {
	*x = CreateFundV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fund_fund_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFundV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFundV1Response) ProtoMessage() {}

func (x *CreateFundV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_fund_fund_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFundV1Response.ProtoReflect.Descriptor instead.
func (*CreateFundV1Response) Descriptor() ([]byte, []int) {
	return file_fund_fund_proto_rawDescGZIP(), []int{1}
}

type GetFundsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFundsV1Request) Reset() {
	*x = GetFundsV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fund_fund_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFundsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFundsV1Request) ProtoMessage() {}

func (x *GetFundsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_fund_fund_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFundsV1Request.ProtoReflect.Descriptor instead.
func (*GetFundsV1Request) Descriptor() ([]byte, []int) {
	return file_fund_fund_proto_rawDescGZIP(), []int{2}
}

type GetFundsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Funds []*GetFundMessage `protobuf:"bytes,1,rep,name=funds,proto3" json:"funds,omitempty"`
}

func (x *GetFundsV1Response) Reset() {
	*x = GetFundsV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fund_fund_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFundsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFundsV1Response) ProtoMessage() {}

func (x *GetFundsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_fund_fund_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFundsV1Response.ProtoReflect.Descriptor instead.
func (*GetFundsV1Response) Descriptor() ([]byte, []int) {
	return file_fund_fund_proto_rawDescGZIP(), []int{3}
}

func (x *GetFundsV1Response) GetFunds() []*GetFundMessage {
	if x != nil {
		return x.Funds
	}
	return nil
}

type GetFundV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *GetFundV1Request) Reset() {
	*x = GetFundV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fund_fund_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFundV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFundV1Request) ProtoMessage() {}

func (x *GetFundV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_fund_fund_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFundV1Request.ProtoReflect.Descriptor instead.
func (*GetFundV1Request) Descriptor() ([]byte, []int) {
	return file_fund_fund_proto_rawDescGZIP(), []int{4}
}

func (x *GetFundV1Request) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type GetFundV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fund *GetFundMessage `protobuf:"bytes,1,opt,name=fund,proto3" json:"fund,omitempty"`
}

func (x *GetFundV1Response) Reset() {
	*x = GetFundV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fund_fund_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFundV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFundV1Response) ProtoMessage() {}

func (x *GetFundV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_fund_fund_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFundV1Response.ProtoReflect.Descriptor instead.
func (*GetFundV1Response) Descriptor() ([]byte, []int) {
	return file_fund_fund_proto_rawDescGZIP(), []int{5}
}

func (x *GetFundV1Response) GetFund() *GetFundMessage {
	if x != nil {
		return x.Fund
	}
	return nil
}

type FundMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *FundMessage) Reset() {
	*x = FundMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fund_fund_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundMessage) ProtoMessage() {}

func (x *FundMessage) ProtoReflect() protoreflect.Message {
	mi := &file_fund_fund_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundMessage.ProtoReflect.Descriptor instead.
func (*FundMessage) Descriptor() ([]byte, []int) {
	return file_fund_fund_proto_rawDescGZIP(), []int{6}
}

func (x *FundMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FundMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetFundMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GetFundMessage) Reset() {
	*x = GetFundMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fund_fund_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFundMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFundMessage) ProtoMessage() {}

func (x *GetFundMessage) ProtoReflect() protoreflect.Message {
	mi := &file_fund_fund_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFundMessage.ProtoReflect.Descriptor instead.
func (*GetFundMessage) Descriptor() ([]byte, []int) {
	return file_fund_fund_proto_rawDescGZIP(), []int{7}
}

func (x *GetFundMessage) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetFundMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetFundMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_fund_fund_proto protoreflect.FileDescriptor

var file_fund_fund_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x75, 0x6e, 0x64, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2b, 0x0a, 0x04, 0x66, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x75, 0x6e, 0x64,
	0x2e, 0x46, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x66, 0x75,
	0x6e, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x46, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x75, 0x6e,
	0x64, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x05, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x22, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x75,
	0x6e, 0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x43,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x66,
	0x75, 0x6e, 0x64, 0x22, 0x43, 0x0a, 0x0b, 0x46, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x56, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46,
	0x75, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x32, 0xcb, 0x02, 0x0a, 0x04, 0x46, 0x75, 0x6e, 0x64, 0x12, 0x71, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x56, 0x31, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e,
	0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x75,
	0x6e, 0x64, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x65, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x56, 0x31, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x66, 0x75, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x56, 0x31, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01,
	0x2a, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x66,
	0x75, 0x6e, 0x64, 0x12, 0x69, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x56,
	0x31, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fund_fund_proto_rawDescOnce sync.Once
	file_fund_fund_proto_rawDescData = file_fund_fund_proto_rawDesc
)

func file_fund_fund_proto_rawDescGZIP() []byte {
	file_fund_fund_proto_rawDescOnce.Do(func() {
		file_fund_fund_proto_rawDescData = protoimpl.X.CompressGZIP(file_fund_fund_proto_rawDescData)
	})
	return file_fund_fund_proto_rawDescData
}

var file_fund_fund_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_fund_fund_proto_goTypes = []interface{}{
	(*CreateFundV1Request)(nil),  // 0: proto.fund.CreateFundV1Request
	(*CreateFundV1Response)(nil), // 1: proto.fund.CreateFundV1Response
	(*GetFundsV1Request)(nil),    // 2: proto.fund.GetFundsV1Request
	(*GetFundsV1Response)(nil),   // 3: proto.fund.GetFundsV1Response
	(*GetFundV1Request)(nil),     // 4: proto.fund.GetFundV1Request
	(*GetFundV1Response)(nil),    // 5: proto.fund.GetFundV1Response
	(*FundMessage)(nil),          // 6: proto.fund.FundMessage
	(*GetFundMessage)(nil),       // 7: proto.fund.GetFundMessage
}
var file_fund_fund_proto_depIdxs = []int32{
	6, // 0: proto.fund.CreateFundV1Request.fund:type_name -> proto.fund.FundMessage
	7, // 1: proto.fund.GetFundsV1Response.funds:type_name -> proto.fund.GetFundMessage
	7, // 2: proto.fund.GetFundV1Response.fund:type_name -> proto.fund.GetFundMessage
	0, // 3: proto.fund.Fund.CreateFundV1:input_type -> proto.fund.CreateFundV1Request
	4, // 4: proto.fund.Fund.GetFundV1:input_type -> proto.fund.GetFundV1Request
	2, // 5: proto.fund.Fund.GetFundsV1:input_type -> proto.fund.GetFundsV1Request
	1, // 6: proto.fund.Fund.CreateFundV1:output_type -> proto.fund.CreateFundV1Response
	5, // 7: proto.fund.Fund.GetFundV1:output_type -> proto.fund.GetFundV1Response
	3, // 8: proto.fund.Fund.GetFundsV1:output_type -> proto.fund.GetFundsV1Response
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_fund_fund_proto_init() }
func file_fund_fund_proto_init() {
	if File_fund_fund_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fund_fund_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFundV1Request); i {
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
		file_fund_fund_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFundV1Response); i {
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
		file_fund_fund_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFundsV1Request); i {
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
		file_fund_fund_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFundsV1Response); i {
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
		file_fund_fund_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFundV1Request); i {
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
		file_fund_fund_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFundV1Response); i {
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
		file_fund_fund_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FundMessage); i {
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
		file_fund_fund_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFundMessage); i {
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
			RawDescriptor: file_fund_fund_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fund_fund_proto_goTypes,
		DependencyIndexes: file_fund_fund_proto_depIdxs,
		MessageInfos:      file_fund_fund_proto_msgTypes,
	}.Build()
	File_fund_fund_proto = out.File
	file_fund_fund_proto_rawDesc = nil
	file_fund_fund_proto_goTypes = nil
	file_fund_fund_proto_depIdxs = nil
}
