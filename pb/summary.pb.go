// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: summary.proto

package pb

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

type SummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockCode string `protobuf:"bytes,1,opt,name=StockCode,proto3" json:"StockCode,omitempty"`
}

func (x *SummaryRequest) Reset() {
	*x = SummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryRequest) ProtoMessage() {}

func (x *SummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryRequest.ProtoReflect.Descriptor instead.
func (*SummaryRequest) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{0}
}

func (x *SummaryRequest) GetStockCode() string {
	if x != nil {
		return x.StockCode
	}
	return ""
}

type SummaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockCode     string `protobuf:"bytes,1,opt,name=StockCode,proto3" json:"StockCode,omitempty"`
	PreviousPrice int64  `protobuf:"varint,2,opt,name=PreviousPrice,proto3" json:"PreviousPrice,omitempty"`
	OpenPrice     int64  `protobuf:"varint,3,opt,name=OpenPrice,proto3" json:"OpenPrice,omitempty"`
	HighestPrice  int64  `protobuf:"varint,4,opt,name=HighestPrice,proto3" json:"HighestPrice,omitempty"`
	LowestPrice   int64  `protobuf:"varint,5,opt,name=LowestPrice,proto3" json:"LowestPrice,omitempty"`
	ClosePrice    int64  `protobuf:"varint,6,opt,name=ClosePrice,proto3" json:"ClosePrice,omitempty"`
	AveragePrice  int64  `protobuf:"varint,7,opt,name=AveragePrice,proto3" json:"AveragePrice,omitempty"`
	Volume        int64  `protobuf:"varint,8,opt,name=Volume,proto3" json:"Volume,omitempty"`
	Value         int64  `protobuf:"varint,9,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *SummaryResponse) Reset() {
	*x = SummaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryResponse) ProtoMessage() {}

func (x *SummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryResponse.ProtoReflect.Descriptor instead.
func (*SummaryResponse) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{1}
}

func (x *SummaryResponse) GetStockCode() string {
	if x != nil {
		return x.StockCode
	}
	return ""
}

func (x *SummaryResponse) GetPreviousPrice() int64 {
	if x != nil {
		return x.PreviousPrice
	}
	return 0
}

func (x *SummaryResponse) GetOpenPrice() int64 {
	if x != nil {
		return x.OpenPrice
	}
	return 0
}

func (x *SummaryResponse) GetHighestPrice() int64 {
	if x != nil {
		return x.HighestPrice
	}
	return 0
}

func (x *SummaryResponse) GetLowestPrice() int64 {
	if x != nil {
		return x.LowestPrice
	}
	return 0
}

func (x *SummaryResponse) GetClosePrice() int64 {
	if x != nil {
		return x.ClosePrice
	}
	return 0
}

func (x *SummaryResponse) GetAveragePrice() int64 {
	if x != nil {
		return x.AveragePrice
	}
	return 0
}

func (x *SummaryResponse) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *SummaryResponse) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_summary_proto protoreflect.FileDescriptor

var file_summary_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6f, 0x68, 0x6c, 0x63, 0x22, 0x2e, 0x0a, 0x0e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xab, 0x02, 0x0a, 0x0f, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x48,
	0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x4c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x4c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0x4a, 0x0a, 0x0e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x14, 0x2e, 0x6f, 0x68, 0x6c, 0x63, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6f, 0x68, 0x6c, 0x63, 0x2e, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x61,
	0x6b, 0x61, 0x72, 0x69, 0x61, 0x77, 0x61, 0x68, 0x79, 0x75, 0x2f, 0x67, 0x6f, 0x2d, 0x6f, 0x68,
	0x6c, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_summary_proto_rawDescOnce sync.Once
	file_summary_proto_rawDescData = file_summary_proto_rawDesc
)

func file_summary_proto_rawDescGZIP() []byte {
	file_summary_proto_rawDescOnce.Do(func() {
		file_summary_proto_rawDescData = protoimpl.X.CompressGZIP(file_summary_proto_rawDescData)
	})
	return file_summary_proto_rawDescData
}

var file_summary_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_summary_proto_goTypes = []interface{}{
	(*SummaryRequest)(nil),  // 0: ohlc.SummaryRequest
	(*SummaryResponse)(nil), // 1: ohlc.SummaryResponse
}
var file_summary_proto_depIdxs = []int32{
	0, // 0: ohlc.SummaryService.Summary:input_type -> ohlc.SummaryRequest
	1, // 1: ohlc.SummaryService.Summary:output_type -> ohlc.SummaryResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_summary_proto_init() }
func file_summary_proto_init() {
	if File_summary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_summary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryRequest); i {
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
		file_summary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryResponse); i {
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
			RawDescriptor: file_summary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_summary_proto_goTypes,
		DependencyIndexes: file_summary_proto_depIdxs,
		MessageInfos:      file_summary_proto_msgTypes,
	}.Build()
	File_summary_proto = out.File
	file_summary_proto_rawDesc = nil
	file_summary_proto_goTypes = nil
	file_summary_proto_depIdxs = nil
}
