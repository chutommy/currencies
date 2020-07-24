// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: currency.proto

package currency

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// GetCurrencyRequest defines the request message for the GetCurrency and
// the SubscribeCurrency calls.
type GetCurrencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name stands for the currency code for the currency.
	// The Name value is not case sensitive.
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *GetCurrencyRequest) Reset() {
	*x = GetCurrencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrencyRequest) ProtoMessage() {}

func (x *GetCurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_currency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrencyRequest.ProtoReflect.Descriptor instead.
func (*GetCurrencyRequest) Descriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{0}
}

func (x *GetCurrencyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// GetCurrencyResponse defines the response message for the GetCurrency call
// and the StreamingSubscribeResponse message.
type GetCurrencyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name stands for the currency code for the currency.
	// Every Name values are capitalized.
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// Country holds the name of the country where the currency came from.
	Country string `protobuf:"bytes,2,opt,name=Country,proto3" json:"Country,omitempty"`
	// Description is the full name of the currency.
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	// Change represents the latest currency change in the percentages.
	Change float32 `protobuf:"fixed32,4,opt,name=Change,proto3" json:"Change,omitempty"`
	// RateUSD is the exchange rates between the currency and the USD.
	// Both currency values are taken from the lastest source update.
	RateUSD float32 `protobuf:"fixed32,5,opt,name=RateUSD,proto3" json:"RateUSD,omitempty"`
	// UpdatedAt is the time of the last update of the currency in the source.
	UpdatedAt string `protobuf:"bytes,6,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *GetCurrencyResponse) Reset() {
	*x = GetCurrencyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrencyResponse) ProtoMessage() {}

func (x *GetCurrencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_currency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrencyResponse.ProtoReflect.Descriptor instead.
func (*GetCurrencyResponse) Descriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{1}
}

func (x *GetCurrencyResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetCurrencyResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetCurrencyResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetCurrencyResponse) GetChange() float32 {
	if x != nil {
		return x.Change
	}
	return 0
}

func (x *GetCurrencyResponse) GetRateUSD() float32 {
	if x != nil {
		return x.RateUSD
	}
	return 0
}

func (x *GetCurrencyResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// StreamingSubscribeResponse defines the response message for
// the SubscribeCurrency call. It holds either GetCurrencyResponse
// or the Status error.
type StreamingSubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*StreamingSubscribeResponse_GetCurrencyResponse
	//	*StreamingSubscribeResponse_Error
	Message isStreamingSubscribeResponse_Message `protobuf_oneof:"message"`
}

func (x *StreamingSubscribeResponse) Reset() {
	*x = StreamingSubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingSubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingSubscribeResponse) ProtoMessage() {}

func (x *StreamingSubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_currency_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingSubscribeResponse.ProtoReflect.Descriptor instead.
func (*StreamingSubscribeResponse) Descriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{2}
}

func (m *StreamingSubscribeResponse) GetMessage() isStreamingSubscribeResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *StreamingSubscribeResponse) GetGetCurrencyResponse() *GetCurrencyResponse {
	if x, ok := x.GetMessage().(*StreamingSubscribeResponse_GetCurrencyResponse); ok {
		return x.GetCurrencyResponse
	}
	return nil
}

func (x *StreamingSubscribeResponse) GetError() *status.Status {
	if x, ok := x.GetMessage().(*StreamingSubscribeResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isStreamingSubscribeResponse_Message interface {
	isStreamingSubscribeResponse_Message()
}

type StreamingSubscribeResponse_GetCurrencyResponse struct {
	// Get_currency_response defines the response message with
	// the data about the currency.
	GetCurrencyResponse *GetCurrencyResponse `protobuf:"bytes,1,opt,name=GetCurrencyResponse,proto3,oneof"`
}

type StreamingSubscribeResponse_Error struct {
	// Error defines the error status of the problem which occured.
	Error *status.Status `protobuf:"bytes,2,opt,name=Error,proto3,oneof"`
}

func (*StreamingSubscribeResponse_GetCurrencyResponse) isStreamingSubscribeResponse_Message() {}

func (*StreamingSubscribeResponse_Error) isStreamingSubscribeResponse_Message() {}

// GetRateRequest defines the request meesage for the GetRate call.
// It needs the base currency and the destination currency.
type GetRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base represents the base currency fot the exchange rate.
	Base string `protobuf:"bytes,1,opt,name=Base,proto3" json:"Base,omitempty"`
	// Destination represents the destination currency fot the exchange rate.
	Destination string `protobuf:"bytes,2,opt,name=Destination,proto3" json:"Destination,omitempty"`
}

func (x *GetRateRequest) Reset() {
	*x = GetRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currency_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRateRequest) ProtoMessage() {}

func (x *GetRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_currency_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRateRequest.ProtoReflect.Descriptor instead.
func (*GetRateRequest) Descriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{3}
}

func (x *GetRateRequest) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *GetRateRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

// GetRateResponse defines the response message for the GetRate call.
// It holds only the xchange rate of the request's base and destination.
type GetRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Rate is the result exchange rate.
	Rate float32 `protobuf:"fixed32,1,opt,name=Rate,proto3" json:"Rate,omitempty"`
}

func (x *GetRateResponse) Reset() {
	*x = GetRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currency_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRateResponse) ProtoMessage() {}

func (x *GetRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_currency_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRateResponse.ProtoReflect.Descriptor instead.
func (*GetRateResponse) Descriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{4}
}

func (x *GetRateResponse) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

var File_currency_proto protoreflect.FileDescriptor

var file_currency_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x61, 0x74, 0x65, 0x55, 0x53, 0x44, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x52, 0x61, 0x74, 0x65, 0x55, 0x53, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x9d, 0x01, 0x0a, 0x1a,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x13, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x42, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x42, 0x61, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x52, 0x61, 0x74, 0x65, 0x32, 0xbd, 0x01, 0x0a, 0x08, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x49, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x2c, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_currency_proto_rawDescOnce sync.Once
	file_currency_proto_rawDescData = file_currency_proto_rawDesc
)

func file_currency_proto_rawDescGZIP() []byte {
	file_currency_proto_rawDescOnce.Do(func() {
		file_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_currency_proto_rawDescData)
	})
	return file_currency_proto_rawDescData
}

var file_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_currency_proto_goTypes = []interface{}{
	(*GetCurrencyRequest)(nil),         // 0: GetCurrencyRequest
	(*GetCurrencyResponse)(nil),        // 1: GetCurrencyResponse
	(*StreamingSubscribeResponse)(nil), // 2: StreamingSubscribeResponse
	(*GetRateRequest)(nil),             // 3: GetRateRequest
	(*GetRateResponse)(nil),            // 4: GetRateResponse
	(*status.Status)(nil),              // 5: google.rpc.Status
}
var file_currency_proto_depIdxs = []int32{
	1, // 0: StreamingSubscribeResponse.GetCurrencyResponse:type_name -> GetCurrencyResponse
	5, // 1: StreamingSubscribeResponse.Error:type_name -> google.rpc.Status
	0, // 2: Currency.GetCurrency:input_type -> GetCurrencyRequest
	0, // 3: Currency.SubscribeCurrency:input_type -> GetCurrencyRequest
	3, // 4: Currency.GetRate:input_type -> GetRateRequest
	1, // 5: Currency.GetCurrency:output_type -> GetCurrencyResponse
	2, // 6: Currency.SubscribeCurrency:output_type -> StreamingSubscribeResponse
	4, // 7: Currency.GetRate:output_type -> GetRateResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_currency_proto_init() }
func file_currency_proto_init() {
	if File_currency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_currency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrencyRequest); i {
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
		file_currency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrencyResponse); i {
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
		file_currency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingSubscribeResponse); i {
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
		file_currency_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRateRequest); i {
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
		file_currency_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRateResponse); i {
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
	file_currency_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*StreamingSubscribeResponse_GetCurrencyResponse)(nil),
		(*StreamingSubscribeResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_currency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_currency_proto_goTypes,
		DependencyIndexes: file_currency_proto_depIdxs,
		MessageInfos:      file_currency_proto_msgTypes,
	}.Build()
	File_currency_proto = out.File
	file_currency_proto_rawDesc = nil
	file_currency_proto_goTypes = nil
	file_currency_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurrencyClient is the client API for Currency service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrencyClient interface {
	// GetCurrency provides the current data about one certain currency.
	// The data holds the currency code, country of oritign, the description,
	// the last currency value change in percentages, the exchange rate to USD
	// and the time of the last update.
	GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error)
	// SubscribeCurrency works as the GetCurrency call, except that
	// it does not send a response instantly but wait until the database
	// changes some of its value, then it sends all subscribed currency data
	// to each client.
	SubscribeCurrency(ctx context.Context, opts ...grpc.CallOption) (Currency_SubscribeCurrencyClient, error)
	// GetRate calculates the exchange rates between the base and
	// the destination. The service takes the latest data
	// from the source.
	GetRate(ctx context.Context, in *GetRateRequest, opts ...grpc.CallOption) (*GetRateResponse, error)
}

type currencyClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyClient(cc grpc.ClientConnInterface) CurrencyClient {
	return &currencyClient{cc}
}

func (c *currencyClient) GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error) {
	out := new(GetCurrencyResponse)
	err := c.cc.Invoke(ctx, "/Currency/GetCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyClient) SubscribeCurrency(ctx context.Context, opts ...grpc.CallOption) (Currency_SubscribeCurrencyClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Currency_serviceDesc.Streams[0], "/Currency/SubscribeCurrency", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencySubscribeCurrencyClient{stream}
	return x, nil
}

type Currency_SubscribeCurrencyClient interface {
	Send(*GetCurrencyRequest) error
	Recv() (*StreamingSubscribeResponse, error)
	grpc.ClientStream
}

type currencySubscribeCurrencyClient struct {
	grpc.ClientStream
}

func (x *currencySubscribeCurrencyClient) Send(m *GetCurrencyRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *currencySubscribeCurrencyClient) Recv() (*StreamingSubscribeResponse, error) {
	m := new(StreamingSubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *currencyClient) GetRate(ctx context.Context, in *GetRateRequest, opts ...grpc.CallOption) (*GetRateResponse, error) {
	out := new(GetRateResponse)
	err := c.cc.Invoke(ctx, "/Currency/GetRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyServer is the server API for Currency service.
type CurrencyServer interface {
	// GetCurrency provides the current data about one certain currency.
	// The data holds the currency code, country of oritign, the description,
	// the last currency value change in percentages, the exchange rate to USD
	// and the time of the last update.
	GetCurrency(context.Context, *GetCurrencyRequest) (*GetCurrencyResponse, error)
	// SubscribeCurrency works as the GetCurrency call, except that
	// it does not send a response instantly but wait until the database
	// changes some of its value, then it sends all subscribed currency data
	// to each client.
	SubscribeCurrency(Currency_SubscribeCurrencyServer) error
	// GetRate calculates the exchange rates between the base and
	// the destination. The service takes the latest data
	// from the source.
	GetRate(context.Context, *GetRateRequest) (*GetRateResponse, error)
}

// UnimplementedCurrencyServer can be embedded to have forward compatible implementations.
type UnimplementedCurrencyServer struct {
}

func (*UnimplementedCurrencyServer) GetCurrency(context.Context, *GetCurrencyRequest) (*GetCurrencyResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCurrency not implemented")
}
func (*UnimplementedCurrencyServer) SubscribeCurrency(Currency_SubscribeCurrencyServer) error {
	return status1.Errorf(codes.Unimplemented, "method SubscribeCurrency not implemented")
}
func (*UnimplementedCurrencyServer) GetRate(context.Context, *GetRateRequest) (*GetRateResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetRate not implemented")
}

func RegisterCurrencyServer(s *grpc.Server, srv CurrencyServer) {
	s.RegisterService(&_Currency_serviceDesc, srv)
}

func _Currency_GetCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServer).GetCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Currency/GetCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServer).GetCurrency(ctx, req.(*GetCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Currency_SubscribeCurrency_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CurrencyServer).SubscribeCurrency(&currencySubscribeCurrencyServer{stream})
}

type Currency_SubscribeCurrencyServer interface {
	Send(*StreamingSubscribeResponse) error
	Recv() (*GetCurrencyRequest, error)
	grpc.ServerStream
}

type currencySubscribeCurrencyServer struct {
	grpc.ServerStream
}

func (x *currencySubscribeCurrencyServer) Send(m *StreamingSubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *currencySubscribeCurrencyServer) Recv() (*GetCurrencyRequest, error) {
	m := new(GetCurrencyRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Currency_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Currency/GetRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServer).GetRate(ctx, req.(*GetRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Currency_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Currency",
	HandlerType: (*CurrencyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrency",
			Handler:    _Currency_GetCurrency_Handler,
		},
		{
			MethodName: "GetRate",
			Handler:    _Currency_GetRate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeCurrency",
			Handler:       _Currency_SubscribeCurrency_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "currency.proto",
}
