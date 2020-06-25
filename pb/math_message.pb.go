// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: math_message.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The numbers are order
type Sum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A float32 `protobuf:"fixed32,1,opt,name=a,proto3" json:"a,omitempty"`
	B float32 `protobuf:"fixed32,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *Sum) Reset() {
	*x = Sum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sum) ProtoMessage() {}

func (x *Sum) ProtoReflect() protoreflect.Message {
	mi := &file_math_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sum.ProtoReflect.Descriptor instead.
func (*Sum) Descriptor() ([]byte, []int) {
	return file_math_message_proto_rawDescGZIP(), []int{0}
}

func (x *Sum) GetA() float32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Sum) GetB() float32 {
	if x != nil {
		return x.B
	}
	return 0
}

type NewSumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum *Sum `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *NewSumRequest) Reset() {
	*x = NewSumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewSumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewSumRequest) ProtoMessage() {}

func (x *NewSumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_math_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewSumRequest.ProtoReflect.Descriptor instead.
func (*NewSumRequest) Descriptor() ([]byte, []int) {
	return file_math_message_proto_rawDescGZIP(), []int{1}
}

func (x *NewSumRequest) GetSum() *Sum {
	if x != nil {
		return x.Sum
	}
	return nil
}

type NewSumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *NewSumResponse) Reset() {
	*x = NewSumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewSumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewSumResponse) ProtoMessage() {}

func (x *NewSumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_math_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewSumResponse.ProtoReflect.Descriptor instead.
func (*NewSumResponse) Descriptor() ([]byte, []int) {
	return file_math_message_proto_rawDescGZIP(), []int{2}
}

func (x *NewSumResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type Multiplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A float32 `protobuf:"fixed32,1,opt,name=a,proto3" json:"a,omitempty"`
	B float32 `protobuf:"fixed32,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *Multiplication) Reset() {
	*x = Multiplication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Multiplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Multiplication) ProtoMessage() {}

func (x *Multiplication) ProtoReflect() protoreflect.Message {
	mi := &file_math_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Multiplication.ProtoReflect.Descriptor instead.
func (*Multiplication) Descriptor() ([]byte, []int) {
	return file_math_message_proto_rawDescGZIP(), []int{3}
}

func (x *Multiplication) GetA() float32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Multiplication) GetB() float32 {
	if x != nil {
		return x.B
	}
	return 0
}

type NewMultiplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Multiplication *Multiplication `protobuf:"bytes,1,opt,name=multiplication,proto3" json:"multiplication,omitempty"`
}

func (x *NewMultiplicationRequest) Reset() {
	*x = NewMultiplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMultiplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMultiplicationRequest) ProtoMessage() {}

func (x *NewMultiplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_math_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMultiplicationRequest.ProtoReflect.Descriptor instead.
func (*NewMultiplicationRequest) Descriptor() ([]byte, []int) {
	return file_math_message_proto_rawDescGZIP(), []int{4}
}

func (x *NewMultiplicationRequest) GetMultiplication() *Multiplication {
	if x != nil {
		return x.Multiplication
	}
	return nil
}

type NewMultiplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *NewMultiplicationResponse) Reset() {
	*x = NewMultiplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMultiplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMultiplicationResponse) ProtoMessage() {}

func (x *NewMultiplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_math_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMultiplicationResponse.ProtoReflect.Descriptor instead.
func (*NewMultiplicationResponse) Descriptor() ([]byte, []int) {
	return file_math_message_proto_rawDescGZIP(), []int{5}
}

func (x *NewMultiplicationResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type FibonacciRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N int32 `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
}

func (x *FibonacciRequest) Reset() {
	*x = FibonacciRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciRequest) ProtoMessage() {}

func (x *FibonacciRequest) ProtoReflect() protoreflect.Message {
	mi := &file_math_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciRequest.ProtoReflect.Descriptor instead.
func (*FibonacciRequest) Descriptor() ([]byte, []int) {
	return file_math_message_proto_rawDescGZIP(), []int{6}
}

func (x *FibonacciRequest) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

type FibonacciResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *FibonacciResponse) Reset() {
	*x = FibonacciResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciResponse) ProtoMessage() {}

func (x *FibonacciResponse) ProtoReflect() protoreflect.Message {
	mi := &file_math_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciResponse.ProtoReflect.Descriptor instead.
func (*FibonacciResponse) Descriptor() ([]byte, []int) {
	return file_math_message_proto_rawDescGZIP(), []int{7}
}

func (x *FibonacciResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_math_message_proto protoreflect.FileDescriptor

var file_math_message_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x61, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x74, 0x68, 0x22, 0x21, 0x0a, 0x03, 0x53, 0x75,
	0x6d, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x61, 0x12,
	0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x62, 0x22, 0x2c, 0x0a,
	0x0d, 0x4e, 0x65, 0x77, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d, 0x61,
	0x74, 0x68, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x22, 0x28, 0x0a, 0x0e, 0x4e,
	0x65, 0x77, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2c, 0x0a, 0x0e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x01, 0x62, 0x22, 0x58, 0x0a, 0x18, 0x4e, 0x65, 0x77, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3c, 0x0a, 0x0e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x33, 0x0a,
	0x19, 0x4e, 0x65, 0x77, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x20, 0x0a, 0x10, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x6e, 0x22, 0x2b, 0x0a, 0x11, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63,
	0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0xd8, 0x01, 0x0a, 0x0b, 0x4d, 0x61, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x32, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e,
	0x4e, 0x65, 0x77, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x6d, 0x61, 0x74, 0x68, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x4e,
	0x65, 0x77, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x4e,
	0x65, 0x77, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x46, 0x69,
	0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x12, 0x16, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x46,
	0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_math_message_proto_rawDescOnce sync.Once
	file_math_message_proto_rawDescData = file_math_message_proto_rawDesc
)

func file_math_message_proto_rawDescGZIP() []byte {
	file_math_message_proto_rawDescOnce.Do(func() {
		file_math_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_math_message_proto_rawDescData)
	})
	return file_math_message_proto_rawDescData
}

var file_math_message_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_math_message_proto_goTypes = []interface{}{
	(*Sum)(nil),                       // 0: math.Sum
	(*NewSumRequest)(nil),             // 1: math.NewSumRequest
	(*NewSumResponse)(nil),            // 2: math.NewSumResponse
	(*Multiplication)(nil),            // 3: math.Multiplication
	(*NewMultiplicationRequest)(nil),  // 4: math.NewMultiplicationRequest
	(*NewMultiplicationResponse)(nil), // 5: math.NewMultiplicationResponse
	(*FibonacciRequest)(nil),          // 6: math.FibonacciRequest
	(*FibonacciResponse)(nil),         // 7: math.FibonacciResponse
}
var file_math_message_proto_depIdxs = []int32{
	0, // 0: math.NewSumRequest.sum:type_name -> math.Sum
	3, // 1: math.NewMultiplicationRequest.multiplication:type_name -> math.Multiplication
	1, // 2: math.MathService.Sum:input_type -> math.NewSumRequest
	4, // 3: math.MathService.Multiplication:input_type -> math.NewMultiplicationRequest
	6, // 4: math.MathService.Fibonacci:input_type -> math.FibonacciRequest
	2, // 5: math.MathService.Sum:output_type -> math.NewSumResponse
	5, // 6: math.MathService.Multiplication:output_type -> math.NewMultiplicationResponse
	7, // 7: math.MathService.Fibonacci:output_type -> math.FibonacciResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_math_message_proto_init() }
func file_math_message_proto_init() {
	if File_math_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_math_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sum); i {
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
		file_math_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewSumRequest); i {
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
		file_math_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewSumResponse); i {
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
		file_math_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Multiplication); i {
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
		file_math_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMultiplicationRequest); i {
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
		file_math_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMultiplicationResponse); i {
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
		file_math_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciRequest); i {
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
		file_math_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciResponse); i {
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
			RawDescriptor: file_math_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_math_message_proto_goTypes,
		DependencyIndexes: file_math_message_proto_depIdxs,
		MessageInfos:      file_math_message_proto_msgTypes,
	}.Build()
	File_math_message_proto = out.File
	file_math_message_proto_rawDesc = nil
	file_math_message_proto_goTypes = nil
	file_math_message_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MathServiceClient is the client API for MathService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MathServiceClient interface {
	Sum(ctx context.Context, in *NewSumRequest, opts ...grpc.CallOption) (*NewSumResponse, error)
	Multiplication(ctx context.Context, in *NewMultiplicationRequest, opts ...grpc.CallOption) (*NewMultiplicationResponse, error)
	Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (MathService_FibonacciClient, error)
}

type mathServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMathServiceClient(cc grpc.ClientConnInterface) MathServiceClient {
	return &mathServiceClient{cc}
}

func (c *mathServiceClient) Sum(ctx context.Context, in *NewSumRequest, opts ...grpc.CallOption) (*NewSumResponse, error) {
	out := new(NewSumResponse)
	err := c.cc.Invoke(ctx, "/math.MathService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathServiceClient) Multiplication(ctx context.Context, in *NewMultiplicationRequest, opts ...grpc.CallOption) (*NewMultiplicationResponse, error) {
	out := new(NewMultiplicationResponse)
	err := c.cc.Invoke(ctx, "/math.MathService/Multiplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathServiceClient) Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (MathService_FibonacciClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MathService_serviceDesc.Streams[0], "/math.MathService/Fibonacci", opts...)
	if err != nil {
		return nil, err
	}
	x := &mathServiceFibonacciClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MathService_FibonacciClient interface {
	Recv() (*FibonacciResponse, error)
	grpc.ClientStream
}

type mathServiceFibonacciClient struct {
	grpc.ClientStream
}

func (x *mathServiceFibonacciClient) Recv() (*FibonacciResponse, error) {
	m := new(FibonacciResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MathServiceServer is the server API for MathService service.
type MathServiceServer interface {
	Sum(context.Context, *NewSumRequest) (*NewSumResponse, error)
	Multiplication(context.Context, *NewMultiplicationRequest) (*NewMultiplicationResponse, error)
	Fibonacci(*FibonacciRequest, MathService_FibonacciServer) error
}

// UnimplementedMathServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMathServiceServer struct {
}

func (*UnimplementedMathServiceServer) Sum(context.Context, *NewSumRequest) (*NewSumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedMathServiceServer) Multiplication(context.Context, *NewMultiplicationRequest) (*NewMultiplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiplication not implemented")
}
func (*UnimplementedMathServiceServer) Fibonacci(*FibonacciRequest, MathService_FibonacciServer) error {
	return status.Errorf(codes.Unimplemented, "method Fibonacci not implemented")
}

func RegisterMathServiceServer(s *grpc.Server, srv MathServiceServer) {
	s.RegisterService(&_MathService_serviceDesc, srv)
}

func _MathService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewSumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/math.MathService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Sum(ctx, req.(*NewSumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathService_Multiplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewMultiplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Multiplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/math.MathService/Multiplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Multiplication(ctx, req.(*NewMultiplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathService_Fibonacci_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FibonacciRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MathServiceServer).Fibonacci(m, &mathServiceFibonacciServer{stream})
}

type MathService_FibonacciServer interface {
	Send(*FibonacciResponse) error
	grpc.ServerStream
}

type mathServiceFibonacciServer struct {
	grpc.ServerStream
}

func (x *mathServiceFibonacciServer) Send(m *FibonacciResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MathService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "math.MathService",
	HandlerType: (*MathServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _MathService_Sum_Handler,
		},
		{
			MethodName: "Multiplication",
			Handler:    _MathService_Multiplication_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Fibonacci",
			Handler:       _MathService_Fibonacci_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "math_message.proto",
}