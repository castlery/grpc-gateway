// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.0
// source: examples/internal/proto/examplepb/wrappers.proto

package examplepb

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Wrappers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringValue *wrappers.StringValue `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	Int32Value  *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=int32_value,json=int32Value,proto3" json:"int32_value,omitempty"`
	Int64Value  *wrappers.Int64Value  `protobuf:"bytes,3,opt,name=int64_value,json=int64Value,proto3" json:"int64_value,omitempty"`
	FloatValue  *wrappers.FloatValue  `protobuf:"bytes,4,opt,name=float_value,json=floatValue,proto3" json:"float_value,omitempty"`
	DoubleValue *wrappers.DoubleValue `protobuf:"bytes,5,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	BoolValue   *wrappers.BoolValue   `protobuf:"bytes,6,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	Uint32Value *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=uint32_value,json=uint32Value,proto3" json:"uint32_value,omitempty"`
	Uint64Value *wrappers.UInt64Value `protobuf:"bytes,8,opt,name=uint64_value,json=uint64Value,proto3" json:"uint64_value,omitempty"`
	BytesValue  *wrappers.BytesValue  `protobuf:"bytes,9,opt,name=bytes_value,json=bytesValue,proto3" json:"bytes_value,omitempty"`
}

func (x *Wrappers) Reset() {
	*x = Wrappers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_wrappers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wrappers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wrappers) ProtoMessage() {}

func (x *Wrappers) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_wrappers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wrappers.ProtoReflect.Descriptor instead.
func (*Wrappers) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_wrappers_proto_rawDescGZIP(), []int{0}
}

func (x *Wrappers) GetStringValue() *wrappers.StringValue {
	if x != nil {
		return x.StringValue
	}
	return nil
}

func (x *Wrappers) GetInt32Value() *wrappers.Int32Value {
	if x != nil {
		return x.Int32Value
	}
	return nil
}

func (x *Wrappers) GetInt64Value() *wrappers.Int64Value {
	if x != nil {
		return x.Int64Value
	}
	return nil
}

func (x *Wrappers) GetFloatValue() *wrappers.FloatValue {
	if x != nil {
		return x.FloatValue
	}
	return nil
}

func (x *Wrappers) GetDoubleValue() *wrappers.DoubleValue {
	if x != nil {
		return x.DoubleValue
	}
	return nil
}

func (x *Wrappers) GetBoolValue() *wrappers.BoolValue {
	if x != nil {
		return x.BoolValue
	}
	return nil
}

func (x *Wrappers) GetUint32Value() *wrappers.UInt32Value {
	if x != nil {
		return x.Uint32Value
	}
	return nil
}

func (x *Wrappers) GetUint64Value() *wrappers.UInt64Value {
	if x != nil {
		return x.Uint64Value
	}
	return nil
}

func (x *Wrappers) GetBytesValue() *wrappers.BytesValue {
	if x != nil {
		return x.BytesValue
	}
	return nil
}

var File_examples_internal_proto_examplepb_wrappers_proto protoreflect.FileDescriptor

var file_examples_internal_proto_examplepb_wrappers_proto_rawDesc = []byte{
	0x0a, 0x30, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x04,
	0x0a, 0x08, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x32, 0xbe, 0x09, 0x0a, 0x0f, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x38, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70,
	0x62, 0x2e, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x1a, 0x38, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x57, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3a, 0x01,
	0x2a, 0x12, 0x66, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x3a, 0x01, 0x2a, 0x12, 0x66, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22,
	0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x3a, 0x01,
	0x2a, 0x12, 0x66, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x62, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x3a, 0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x3a, 0x01,
	0x2a, 0x12, 0x66, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x57, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x3a,
	0x01, 0x2a, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x32, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_examples_internal_proto_examplepb_wrappers_proto_rawDescOnce sync.Once
	file_examples_internal_proto_examplepb_wrappers_proto_rawDescData = file_examples_internal_proto_examplepb_wrappers_proto_rawDesc
)

func file_examples_internal_proto_examplepb_wrappers_proto_rawDescGZIP() []byte {
	file_examples_internal_proto_examplepb_wrappers_proto_rawDescOnce.Do(func() {
		file_examples_internal_proto_examplepb_wrappers_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_internal_proto_examplepb_wrappers_proto_rawDescData)
	})
	return file_examples_internal_proto_examplepb_wrappers_proto_rawDescData
}

var file_examples_internal_proto_examplepb_wrappers_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_examples_internal_proto_examplepb_wrappers_proto_goTypes = []interface{}{
	(*Wrappers)(nil),             // 0: grpc.gateway.examples.internal.proto.examplepb.Wrappers
	(*wrappers.StringValue)(nil), // 1: google.protobuf.StringValue
	(*wrappers.Int32Value)(nil),  // 2: google.protobuf.Int32Value
	(*wrappers.Int64Value)(nil),  // 3: google.protobuf.Int64Value
	(*wrappers.FloatValue)(nil),  // 4: google.protobuf.FloatValue
	(*wrappers.DoubleValue)(nil), // 5: google.protobuf.DoubleValue
	(*wrappers.BoolValue)(nil),   // 6: google.protobuf.BoolValue
	(*wrappers.UInt32Value)(nil), // 7: google.protobuf.UInt32Value
	(*wrappers.UInt64Value)(nil), // 8: google.protobuf.UInt64Value
	(*wrappers.BytesValue)(nil),  // 9: google.protobuf.BytesValue
	(*empty.Empty)(nil),          // 10: google.protobuf.Empty
}
var file_examples_internal_proto_examplepb_wrappers_proto_depIdxs = []int32{
	1,  // 0: grpc.gateway.examples.internal.proto.examplepb.Wrappers.string_value:type_name -> google.protobuf.StringValue
	2,  // 1: grpc.gateway.examples.internal.proto.examplepb.Wrappers.int32_value:type_name -> google.protobuf.Int32Value
	3,  // 2: grpc.gateway.examples.internal.proto.examplepb.Wrappers.int64_value:type_name -> google.protobuf.Int64Value
	4,  // 3: grpc.gateway.examples.internal.proto.examplepb.Wrappers.float_value:type_name -> google.protobuf.FloatValue
	5,  // 4: grpc.gateway.examples.internal.proto.examplepb.Wrappers.double_value:type_name -> google.protobuf.DoubleValue
	6,  // 5: grpc.gateway.examples.internal.proto.examplepb.Wrappers.bool_value:type_name -> google.protobuf.BoolValue
	7,  // 6: grpc.gateway.examples.internal.proto.examplepb.Wrappers.uint32_value:type_name -> google.protobuf.UInt32Value
	8,  // 7: grpc.gateway.examples.internal.proto.examplepb.Wrappers.uint64_value:type_name -> google.protobuf.UInt64Value
	9,  // 8: grpc.gateway.examples.internal.proto.examplepb.Wrappers.bytes_value:type_name -> google.protobuf.BytesValue
	0,  // 9: grpc.gateway.examples.internal.proto.examplepb.WrappersService.Create:input_type -> grpc.gateway.examples.internal.proto.examplepb.Wrappers
	1,  // 10: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateStringValue:input_type -> google.protobuf.StringValue
	2,  // 11: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateInt32Value:input_type -> google.protobuf.Int32Value
	3,  // 12: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateInt64Value:input_type -> google.protobuf.Int64Value
	4,  // 13: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateFloatValue:input_type -> google.protobuf.FloatValue
	5,  // 14: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateDoubleValue:input_type -> google.protobuf.DoubleValue
	6,  // 15: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateBoolValue:input_type -> google.protobuf.BoolValue
	7,  // 16: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateUInt32Value:input_type -> google.protobuf.UInt32Value
	8,  // 17: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateUInt64Value:input_type -> google.protobuf.UInt64Value
	9,  // 18: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateBytesValue:input_type -> google.protobuf.BytesValue
	10, // 19: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateEmpty:input_type -> google.protobuf.Empty
	0,  // 20: grpc.gateway.examples.internal.proto.examplepb.WrappersService.Create:output_type -> grpc.gateway.examples.internal.proto.examplepb.Wrappers
	1,  // 21: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateStringValue:output_type -> google.protobuf.StringValue
	2,  // 22: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateInt32Value:output_type -> google.protobuf.Int32Value
	3,  // 23: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateInt64Value:output_type -> google.protobuf.Int64Value
	4,  // 24: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateFloatValue:output_type -> google.protobuf.FloatValue
	5,  // 25: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateDoubleValue:output_type -> google.protobuf.DoubleValue
	6,  // 26: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateBoolValue:output_type -> google.protobuf.BoolValue
	7,  // 27: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateUInt32Value:output_type -> google.protobuf.UInt32Value
	8,  // 28: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateUInt64Value:output_type -> google.protobuf.UInt64Value
	9,  // 29: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateBytesValue:output_type -> google.protobuf.BytesValue
	10, // 30: grpc.gateway.examples.internal.proto.examplepb.WrappersService.CreateEmpty:output_type -> google.protobuf.Empty
	20, // [20:31] is the sub-list for method output_type
	9,  // [9:20] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_examples_internal_proto_examplepb_wrappers_proto_init() }
func file_examples_internal_proto_examplepb_wrappers_proto_init() {
	if File_examples_internal_proto_examplepb_wrappers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_internal_proto_examplepb_wrappers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wrappers); i {
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
			RawDescriptor: file_examples_internal_proto_examplepb_wrappers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_examples_internal_proto_examplepb_wrappers_proto_goTypes,
		DependencyIndexes: file_examples_internal_proto_examplepb_wrappers_proto_depIdxs,
		MessageInfos:      file_examples_internal_proto_examplepb_wrappers_proto_msgTypes,
	}.Build()
	File_examples_internal_proto_examplepb_wrappers_proto = out.File
	file_examples_internal_proto_examplepb_wrappers_proto_rawDesc = nil
	file_examples_internal_proto_examplepb_wrappers_proto_goTypes = nil
	file_examples_internal_proto_examplepb_wrappers_proto_depIdxs = nil
}
