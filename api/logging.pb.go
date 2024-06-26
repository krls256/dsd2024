// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: api/logging.proto

package api

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

type LoggingZeroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoggingZeroRequest) Reset() {
	*x = LoggingZeroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_logging_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggingZeroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggingZeroRequest) ProtoMessage() {}

func (x *LoggingZeroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_logging_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggingZeroRequest.ProtoReflect.Descriptor instead.
func (*LoggingZeroRequest) Descriptor() ([]byte, []int) {
	return file_api_logging_proto_rawDescGZIP(), []int{0}
}

type LoggingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *LoggingMessage) Reset() {
	*x = LoggingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_logging_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggingMessage) ProtoMessage() {}

func (x *LoggingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_logging_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggingMessage.ProtoReflect.Descriptor instead.
func (*LoggingMessage) Descriptor() ([]byte, []int) {
	return file_api_logging_proto_rawDescGZIP(), []int{1}
}

func (x *LoggingMessage) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *LoggingMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type AllText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text []string `protobuf:"bytes,1,rep,name=text,proto3" json:"text,omitempty"`
}

func (x *AllText) Reset() {
	*x = AllText{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_logging_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllText) ProtoMessage() {}

func (x *AllText) ProtoReflect() protoreflect.Message {
	mi := &file_api_logging_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllText.ProtoReflect.Descriptor instead.
func (*AllText) Descriptor() ([]byte, []int) {
	return file_api_logging_proto_rawDescGZIP(), []int{2}
}

func (x *AllText) GetText() []string {
	if x != nil {
		return x.Text
	}
	return nil
}

type LoggingStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success      bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *LoggingStatusResponse) Reset() {
	*x = LoggingStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_logging_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggingStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggingStatusResponse) ProtoMessage() {}

func (x *LoggingStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_logging_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggingStatusResponse.ProtoReflect.Descriptor instead.
func (*LoggingStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_logging_proto_rawDescGZIP(), []int{3}
}

func (x *LoggingStatusResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LoggingStatusResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_api_logging_proto protoreflect.FileDescriptor

var file_api_logging_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x5a, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34,
	0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x1d, 0x0a, 0x07, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x22, 0x56, 0x0a, 0x15, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x7a, 0x0a, 0x0e, 0x4c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a,
	0x03, 0x4c, 0x6f, 0x67, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x12, 0x17,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5a, 0x65, 0x72, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c,
	0x6c, 0x54, 0x65, 0x78, 0x74, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_logging_proto_rawDescOnce sync.Once
	file_api_logging_proto_rawDescData = file_api_logging_proto_rawDesc
)

func file_api_logging_proto_rawDescGZIP() []byte {
	file_api_logging_proto_rawDescOnce.Do(func() {
		file_api_logging_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_logging_proto_rawDescData)
	})
	return file_api_logging_proto_rawDescData
}

var file_api_logging_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_logging_proto_goTypes = []interface{}{
	(*LoggingZeroRequest)(nil),    // 0: api.LoggingZeroRequest
	(*LoggingMessage)(nil),        // 1: api.LoggingMessage
	(*AllText)(nil),               // 2: api.AllText
	(*LoggingStatusResponse)(nil), // 3: api.LoggingStatusResponse
}
var file_api_logging_proto_depIdxs = []int32{
	1, // 0: api.LoggingService.Log:input_type -> api.LoggingMessage
	0, // 1: api.LoggingService.All:input_type -> api.LoggingZeroRequest
	3, // 2: api.LoggingService.Log:output_type -> api.LoggingStatusResponse
	2, // 3: api.LoggingService.All:output_type -> api.AllText
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_logging_proto_init() }
func file_api_logging_proto_init() {
	if File_api_logging_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_logging_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggingZeroRequest); i {
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
		file_api_logging_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggingMessage); i {
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
		file_api_logging_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllText); i {
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
		file_api_logging_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggingStatusResponse); i {
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
			RawDescriptor: file_api_logging_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_logging_proto_goTypes,
		DependencyIndexes: file_api_logging_proto_depIdxs,
		MessageInfos:      file_api_logging_proto_msgTypes,
	}.Build()
	File_api_logging_proto = out.File
	file_api_logging_proto_rawDesc = nil
	file_api_logging_proto_goTypes = nil
	file_api_logging_proto_depIdxs = nil
}
