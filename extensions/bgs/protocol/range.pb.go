// DO NOT EDIT: this file was auto-generated by Gophercraft/protoss

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: bgs/low/pb/client/global_extensions/range.proto

package protocol

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

type UnsignedIntRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min *uint64 `protobuf:"varint,1,opt,name=min" json:"min,omitempty"`
	Max *uint64 `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
}

func (x *UnsignedIntRange) Reset() {
	*x = UnsignedIntRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_global_extensions_range_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsignedIntRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsignedIntRange) ProtoMessage() {}

func (x *UnsignedIntRange) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_global_extensions_range_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsignedIntRange.ProtoReflect.Descriptor instead.
func (*UnsignedIntRange) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_global_extensions_range_proto_rawDescGZIP(), []int{0}
}

func (x *UnsignedIntRange) GetMin() uint64 {
	if x != nil && x.Min != nil {
		return *x.Min
	}
	return 0
}

func (x *UnsignedIntRange) GetMax() uint64 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

type SignedIntRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min *int64 `protobuf:"varint,1,opt,name=min" json:"min,omitempty"`
	Max *int64 `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
}

func (x *SignedIntRange) Reset() {
	*x = SignedIntRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_global_extensions_range_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedIntRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedIntRange) ProtoMessage() {}

func (x *SignedIntRange) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_global_extensions_range_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedIntRange.ProtoReflect.Descriptor instead.
func (*SignedIntRange) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_global_extensions_range_proto_rawDescGZIP(), []int{1}
}

func (x *SignedIntRange) GetMin() int64 {
	if x != nil && x.Min != nil {
		return *x.Min
	}
	return 0
}

func (x *SignedIntRange) GetMax() int64 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

type FloatRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min *float32 `protobuf:"fixed32,1,opt,name=min" json:"min,omitempty"`
	Max *float32 `protobuf:"fixed32,2,opt,name=max" json:"max,omitempty"`
}

func (x *FloatRange) Reset() {
	*x = FloatRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_global_extensions_range_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FloatRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloatRange) ProtoMessage() {}

func (x *FloatRange) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_global_extensions_range_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FloatRange.ProtoReflect.Descriptor instead.
func (*FloatRange) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_global_extensions_range_proto_rawDescGZIP(), []int{2}
}

func (x *FloatRange) GetMin() float32 {
	if x != nil && x.Min != nil {
		return *x.Min
	}
	return 0
}

func (x *FloatRange) GetMax() float32 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

var File_bgs_low_pb_client_global_extensions_range_proto protoreflect.FileDescriptor

var file_bgs_low_pb_client_global_extensions_range_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22,
	0x36, 0x0a, 0x10, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x22, 0x34, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x49, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x22, 0x30, 0x0a,
	0x0a, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x42,
	0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f,
	0x70, 0x68, 0x65, 0x72, 0x63, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x62, 0x67, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
}

var (
	file_bgs_low_pb_client_global_extensions_range_proto_rawDescOnce sync.Once
	file_bgs_low_pb_client_global_extensions_range_proto_rawDescData = file_bgs_low_pb_client_global_extensions_range_proto_rawDesc
)

func file_bgs_low_pb_client_global_extensions_range_proto_rawDescGZIP() []byte {
	file_bgs_low_pb_client_global_extensions_range_proto_rawDescOnce.Do(func() {
		file_bgs_low_pb_client_global_extensions_range_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgs_low_pb_client_global_extensions_range_proto_rawDescData)
	})
	return file_bgs_low_pb_client_global_extensions_range_proto_rawDescData
}

var file_bgs_low_pb_client_global_extensions_range_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bgs_low_pb_client_global_extensions_range_proto_goTypes = []interface{}{
	(*UnsignedIntRange)(nil), // 0: bgs.protocol.UnsignedIntRange
	(*SignedIntRange)(nil),   // 1: bgs.protocol.SignedIntRange
	(*FloatRange)(nil),       // 2: bgs.protocol.FloatRange
}
var file_bgs_low_pb_client_global_extensions_range_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bgs_low_pb_client_global_extensions_range_proto_init() }
func file_bgs_low_pb_client_global_extensions_range_proto_init() {
	if File_bgs_low_pb_client_global_extensions_range_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bgs_low_pb_client_global_extensions_range_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsignedIntRange); i {
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
		file_bgs_low_pb_client_global_extensions_range_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedIntRange); i {
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
		file_bgs_low_pb_client_global_extensions_range_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FloatRange); i {
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
			RawDescriptor: file_bgs_low_pb_client_global_extensions_range_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bgs_low_pb_client_global_extensions_range_proto_goTypes,
		DependencyIndexes: file_bgs_low_pb_client_global_extensions_range_proto_depIdxs,
		MessageInfos:      file_bgs_low_pb_client_global_extensions_range_proto_msgTypes,
	}.Build()
	File_bgs_low_pb_client_global_extensions_range_proto = out.File
	file_bgs_low_pb_client_global_extensions_range_proto_rawDesc = nil
	file_bgs_low_pb_client_global_extensions_range_proto_goTypes = nil
	file_bgs_low_pb_client_global_extensions_range_proto_depIdxs = nil
}
