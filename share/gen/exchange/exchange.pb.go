// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: exchange/exchange.proto

package exchange

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

type ExchangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum            int64  `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
	InputCurrency  string `protobuf:"bytes,2,opt,name=inputCurrency,proto3" json:"inputCurrency,omitempty"`
	OutputCurrency string `protobuf:"bytes,3,opt,name=outputCurrency,proto3" json:"outputCurrency,omitempty"`
}

func (x *ExchangeRequest) Reset() {
	*x = ExchangeRequest{}
	mi := &file_exchange_exchange_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExchangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRequest) ProtoMessage() {}

func (x *ExchangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_exchange_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRequest.ProtoReflect.Descriptor instead.
func (*ExchangeRequest) Descriptor() ([]byte, []int) {
	return file_exchange_exchange_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeRequest) GetSum() int64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

func (x *ExchangeRequest) GetInputCurrency() string {
	if x != nil {
		return x.InputCurrency
	}
	return ""
}

func (x *ExchangeRequest) GetOutputCurrency() string {
	if x != nil {
		return x.OutputCurrency
	}
	return ""
}

type ExchangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int64 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *ExchangeResponse) Reset() {
	*x = ExchangeResponse{}
	mi := &file_exchange_exchange_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExchangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeResponse) ProtoMessage() {}

func (x *ExchangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_exchange_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeResponse.ProtoReflect.Descriptor instead.
func (*ExchangeResponse) Descriptor() ([]byte, []int) {
	return file_exchange_exchange_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangeResponse) GetSum() int64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

var File_exchange_exchange_proto protoreflect.FileDescriptor

var file_exchange_exchange_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x22, 0x71, 0x0a, 0x0f, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x26,
	0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x24, 0x0a, 0x10, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x32, 0x4e, 0x0a, 0x09,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x08, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x73, 0x73, 0x61,
	0x64, 0x64, 0x61, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2d, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x72, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x3b, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exchange_exchange_proto_rawDescOnce sync.Once
	file_exchange_exchange_proto_rawDescData = file_exchange_exchange_proto_rawDesc
)

func file_exchange_exchange_proto_rawDescGZIP() []byte {
	file_exchange_exchange_proto_rawDescOnce.Do(func() {
		file_exchange_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_exchange_exchange_proto_rawDescData)
	})
	return file_exchange_exchange_proto_rawDescData
}

var file_exchange_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_exchange_exchange_proto_goTypes = []any{
	(*ExchangeRequest)(nil),  // 0: exchange.ExchangeRequest
	(*ExchangeResponse)(nil), // 1: exchange.ExchangeResponse
}
var file_exchange_exchange_proto_depIdxs = []int32{
	0, // 0: exchange.Exchanger.Exchange:input_type -> exchange.ExchangeRequest
	1, // 1: exchange.Exchanger.Exchange:output_type -> exchange.ExchangeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_exchange_exchange_proto_init() }
func file_exchange_exchange_proto_init() {
	if File_exchange_exchange_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_exchange_exchange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exchange_exchange_proto_goTypes,
		DependencyIndexes: file_exchange_exchange_proto_depIdxs,
		MessageInfos:      file_exchange_exchange_proto_msgTypes,
	}.Build()
	File_exchange_exchange_proto = out.File
	file_exchange_exchange_proto_rawDesc = nil
	file_exchange_exchange_proto_goTypes = nil
	file_exchange_exchange_proto_depIdxs = nil
}