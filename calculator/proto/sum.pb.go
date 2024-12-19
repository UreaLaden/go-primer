// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.0--rc2
// source: sum.proto

package proto

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

type AddRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Num1          uint32                 `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2          uint32                 `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	mi := &file_sum_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sum_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_sum_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetNum1() uint32 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *AddRequest) GetNum2() uint32 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type AddResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        uint32                 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	mi := &file_sum_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sum_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_sum_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetResult() uint32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_sum_proto protoreflect.FileDescriptor

var file_sum_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x34, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x32, 0x22, 0x25, 0x0a,
	0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sum_proto_rawDescOnce sync.Once
	file_sum_proto_rawDescData = file_sum_proto_rawDesc
)

func file_sum_proto_rawDescGZIP() []byte {
	file_sum_proto_rawDescOnce.Do(func() {
		file_sum_proto_rawDescData = protoimpl.X.CompressGZIP(file_sum_proto_rawDescData)
	})
	return file_sum_proto_rawDescData
}

var file_sum_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sum_proto_goTypes = []any{
	(*AddRequest)(nil),  // 0: calculator.AddRequest
	(*AddResponse)(nil), // 1: calculator.AddResponse
}
var file_sum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sum_proto_init() }
func file_sum_proto_init() {
	if File_sum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sum_proto_goTypes,
		DependencyIndexes: file_sum_proto_depIdxs,
		MessageInfos:      file_sum_proto_msgTypes,
	}.Build()
	File_sum_proto = out.File
	file_sum_proto_rawDesc = nil
	file_sum_proto_goTypes = nil
	file_sum_proto_depIdxs = nil
}
