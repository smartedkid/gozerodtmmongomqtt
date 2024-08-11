// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.19.4
// source: order_srv.proto

package order_srv

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_srv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_order_srv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_order_srv_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_srv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_order_srv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_order_srv_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type AddOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid  int64 `protobuf:"varint,1,opt,name=Userid,proto3" json:"Userid,omitempty"`
	GoodsId int64 `protobuf:"varint,2,opt,name=GoodsId,proto3" json:"GoodsId,omitempty"`
	Pay     int64 `protobuf:"varint,3,opt,name=Pay,proto3" json:"Pay,omitempty"`
	Count   int64 `protobuf:"varint,4,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *AddOrderRequest) Reset() {
	*x = AddOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_srv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrderRequest) ProtoMessage() {}

func (x *AddOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_srv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrderRequest.ProtoReflect.Descriptor instead.
func (*AddOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_srv_proto_rawDescGZIP(), []int{2}
}

func (x *AddOrderRequest) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *AddOrderRequest) GetGoodsId() int64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *AddOrderRequest) GetPay() int64 {
	if x != nil {
		return x.Pay
	}
	return 0
}

func (x *AddOrderRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type AddOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddOrderResponse) Reset() {
	*x = AddOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_srv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrderResponse) ProtoMessage() {}

func (x *AddOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_srv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrderResponse.ProtoReflect.Descriptor instead.
func (*AddOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_srv_proto_rawDescGZIP(), []int{3}
}

func (x *AddOrderResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_order_srv_proto protoreflect.FileDescriptor

var file_order_srv_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x72, 0x76, 0x22, 0x1d, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x1e, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x6b, 0x0a, 0x0f, 0x41,
	0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x50, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x50,
	0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x90, 0x02, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x72, 0x76, 0x12, 0x2f, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x49,
	0x6e, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x72, 0x76, 0x2e,
	0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x47, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x64,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x72, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_srv_proto_rawDescOnce sync.Once
	file_order_srv_proto_rawDescData = file_order_srv_proto_rawDesc
)

func file_order_srv_proto_rawDescGZIP() []byte {
	file_order_srv_proto_rawDescOnce.Do(func() {
		file_order_srv_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_srv_proto_rawDescData)
	})
	return file_order_srv_proto_rawDescData
}

var file_order_srv_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_order_srv_proto_goTypes = []interface{}{
	(*Request)(nil),          // 0: order_srv.Request
	(*Response)(nil),         // 1: order_srv.Response
	(*AddOrderRequest)(nil),  // 2: order_srv.AddOrderRequest
	(*AddOrderResponse)(nil), // 3: order_srv.AddOrderResponse
}
var file_order_srv_proto_depIdxs = []int32{
	0, // 0: order_srv.Order_srv.Ping:input_type -> order_srv.Request
	2, // 1: order_srv.Order_srv.Order:input_type -> order_srv.AddOrderRequest
	2, // 2: order_srv.Order_srv.TranInsOrder:input_type -> order_srv.AddOrderRequest
	2, // 3: order_srv.Order_srv.TranDelOrder:input_type -> order_srv.AddOrderRequest
	1, // 4: order_srv.Order_srv.Ping:output_type -> order_srv.Response
	3, // 5: order_srv.Order_srv.Order:output_type -> order_srv.AddOrderResponse
	3, // 6: order_srv.Order_srv.TranInsOrder:output_type -> order_srv.AddOrderResponse
	3, // 7: order_srv.Order_srv.TranDelOrder:output_type -> order_srv.AddOrderResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_srv_proto_init() }
func file_order_srv_proto_init() {
	if File_order_srv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_srv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_order_srv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_order_srv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrderRequest); i {
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
		file_order_srv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrderResponse); i {
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
			RawDescriptor: file_order_srv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_srv_proto_goTypes,
		DependencyIndexes: file_order_srv_proto_depIdxs,
		MessageInfos:      file_order_srv_proto_msgTypes,
	}.Build()
	File_order_srv_proto = out.File
	file_order_srv_proto_rawDesc = nil
	file_order_srv_proto_goTypes = nil
	file_order_srv_proto_depIdxs = nil
}
