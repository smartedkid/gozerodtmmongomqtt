// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.19.4
// source: score_srv.proto

package score_srv

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
		mi := &file_score_srv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_score_srv_proto_msgTypes[0]
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
	return file_score_srv_proto_rawDescGZIP(), []int{0}
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
		mi := &file_score_srv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_score_srv_proto_msgTypes[1]
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
	return file_score_srv_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type AddUserScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid  int64 `protobuf:"varint,1,opt,name=Userid,proto3" json:"Userid,omitempty"`
	GoodsId int64 `protobuf:"varint,2,opt,name=GoodsId,proto3" json:"GoodsId,omitempty"`
	Score   int64 `protobuf:"varint,3,opt,name=Score,proto3" json:"Score,omitempty"`
}

func (x *AddUserScoreRequest) Reset() {
	*x = AddUserScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_srv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserScoreRequest) ProtoMessage() {}

func (x *AddUserScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_score_srv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserScoreRequest.ProtoReflect.Descriptor instead.
func (*AddUserScoreRequest) Descriptor() ([]byte, []int) {
	return file_score_srv_proto_rawDescGZIP(), []int{2}
}

func (x *AddUserScoreRequest) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *AddUserScoreRequest) GetGoodsId() int64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *AddUserScoreRequest) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type AddUserScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddUserScoreResponse) Reset() {
	*x = AddUserScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_srv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserScoreResponse) ProtoMessage() {}

func (x *AddUserScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_score_srv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserScoreResponse.ProtoReflect.Descriptor instead.
func (*AddUserScoreResponse) Descriptor() ([]byte, []int) {
	return file_score_srv_proto_rawDescGZIP(), []int{3}
}

func (x *AddUserScoreResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_score_srv_proto protoreflect.FileDescriptor

var file_score_srv_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x22, 0x1d, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x1e, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x5d, 0x0a, 0x13, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x41, 0x64,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xaf, 0x02, 0x0a,
	0x09, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x12, 0x2f, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x12, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73,
	0x72, 0x76, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x2e, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c,
	0x54, 0x72, 0x61, 0x6e, 0x49, 0x6e, 0x73, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x2e, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a,
	0x0c, 0x54, 0x72, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x2e,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_score_srv_proto_rawDescOnce sync.Once
	file_score_srv_proto_rawDescData = file_score_srv_proto_rawDesc
)

func file_score_srv_proto_rawDescGZIP() []byte {
	file_score_srv_proto_rawDescOnce.Do(func() {
		file_score_srv_proto_rawDescData = protoimpl.X.CompressGZIP(file_score_srv_proto_rawDescData)
	})
	return file_score_srv_proto_rawDescData
}

var file_score_srv_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_score_srv_proto_goTypes = []interface{}{
	(*Request)(nil),              // 0: score_srv.Request
	(*Response)(nil),             // 1: score_srv.Response
	(*AddUserScoreRequest)(nil),  // 2: score_srv.AddUserScoreRequest
	(*AddUserScoreResponse)(nil), // 3: score_srv.AddUserScoreResponse
}
var file_score_srv_proto_depIdxs = []int32{
	0, // 0: score_srv.Score_srv.Ping:input_type -> score_srv.Request
	2, // 1: score_srv.Score_srv.AddUserScore:input_type -> score_srv.AddUserScoreRequest
	2, // 2: score_srv.Score_srv.TranInsScore:input_type -> score_srv.AddUserScoreRequest
	2, // 3: score_srv.Score_srv.TranDelScore:input_type -> score_srv.AddUserScoreRequest
	1, // 4: score_srv.Score_srv.Ping:output_type -> score_srv.Response
	3, // 5: score_srv.Score_srv.AddUserScore:output_type -> score_srv.AddUserScoreResponse
	3, // 6: score_srv.Score_srv.TranInsScore:output_type -> score_srv.AddUserScoreResponse
	3, // 7: score_srv.Score_srv.TranDelScore:output_type -> score_srv.AddUserScoreResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_score_srv_proto_init() }
func file_score_srv_proto_init() {
	if File_score_srv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_score_srv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_score_srv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_score_srv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserScoreRequest); i {
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
		file_score_srv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserScoreResponse); i {
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
			RawDescriptor: file_score_srv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_score_srv_proto_goTypes,
		DependencyIndexes: file_score_srv_proto_depIdxs,
		MessageInfos:      file_score_srv_proto_msgTypes,
	}.Build()
	File_score_srv_proto = out.File
	file_score_srv_proto_rawDesc = nil
	file_score_srv_proto_goTypes = nil
	file_score_srv_proto_depIdxs = nil
}
