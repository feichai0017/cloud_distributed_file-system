// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: proxy.proto

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

type SingleAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Params []byte `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"` // json encoded
}

func (x *SingleAction) Reset() {
	*x = SingleAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleAction) ProtoMessage() {}

func (x *SingleAction) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleAction.ProtoReflect.Descriptor instead.
func (*SingleAction) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *SingleAction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SingleAction) GetParams() []byte {
	if x != nil {
		return x.Params
	}
	return nil
}

type ReqExec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sequence    bool            `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`       // whether to execute in sequence
	Transaction bool            `protobuf:"varint,2,opt,name=transaction,proto3" json:"transaction,omitempty"` // whether to execute in a transaction
	ResultType  int32           `protobuf:"varint,3,opt,name=resultType,proto3" json:"resultType,omitempty"`   // 0: every sql result will be returned, 1: only return the last result
	Actions     []*SingleAction `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`          // one or more actions
}

func (x *ReqExec) Reset() {
	*x = ReqExec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqExec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqExec) ProtoMessage() {}

func (x *ReqExec) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqExec.ProtoReflect.Descriptor instead.
func (*ReqExec) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *ReqExec) GetSequence() bool {
	if x != nil {
		return x.Sequence
	}
	return false
}

func (x *ReqExec) GetTransaction() bool {
	if x != nil {
		return x.Transaction
	}
	return false
}

func (x *ReqExec) GetResultType() int32 {
	if x != nil {
		return x.ResultType
	}
	return 0
}

func (x *ReqExec) GetActions() []*SingleAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

type ResExec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"` // executed result, json encoded
}

func (x *ResExec) Reset() {
	*x = ResExec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResExec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResExec) ProtoMessage() {}

func (x *ResExec) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResExec.ProtoReflect.Descriptor instead.
func (*ResExec) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *ResExec) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResExec) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ResExec) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proxy_proto protoreflect.FileDescriptor

var file_proxy_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x64, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x22, 0x3a, 0x0a, 0x0c, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x45, 0x78, 0x65, 0x63, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a,
	0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x43, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x45, 0x78, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0x69, 0x0a, 0x0e, 0x44, 0x42, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x62, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x45, 0x78, 0x65, 0x63, 0x1a, 0x21, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x62,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x45, 0x78, 0x65, 0x63, 0x22, 0x00, 0x42,
	0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2d, 0x69, 0x6e, 0x2d, 0x63, 0x6e, 0x2f, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69,
	0x61, 0x6c, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_proto_rawDescOnce sync.Once
	file_proxy_proto_rawDescData = file_proxy_proto_rawDesc
)

func file_proxy_proto_rawDescGZIP() []byte {
	file_proxy_proto_rawDescOnce.Do(func() {
		file_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_proto_rawDescData)
	})
	return file_proxy_proto_rawDescData
}

var file_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proxy_proto_goTypes = []any{
	(*SingleAction)(nil), // 0: go.micro.service.dbproxy.SingleAction
	(*ReqExec)(nil),      // 1: go.micro.service.dbproxy.ReqExec
	(*ResExec)(nil),      // 2: go.micro.service.dbproxy.ResExec
}
var file_proxy_proto_depIdxs = []int32{
	0, // 0: go.micro.service.dbproxy.ReqExec.actions:type_name -> go.micro.service.dbproxy.SingleAction
	1, // 1: go.micro.service.dbproxy.DBProxyService.ExecuteAction:input_type -> go.micro.service.dbproxy.ReqExec
	2, // 2: go.micro.service.dbproxy.DBProxyService.ExecuteAction:output_type -> go.micro.service.dbproxy.ResExec
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proxy_proto_init() }
func file_proxy_proto_init() {
	if File_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SingleAction); i {
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
		file_proxy_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ReqExec); i {
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
		file_proxy_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ResExec); i {
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
			RawDescriptor: file_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proxy_proto_goTypes,
		DependencyIndexes: file_proxy_proto_depIdxs,
		MessageInfos:      file_proxy_proto_msgTypes,
	}.Build()
	File_proxy_proto = out.File
	file_proxy_proto_rawDesc = nil
	file_proxy_proto_goTypes = nil
	file_proxy_proto_depIdxs = nil
}
