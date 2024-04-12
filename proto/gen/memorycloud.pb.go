// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: proto/proto/memorycloud.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     int64                  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	FilePath   string                 `protobuf:"bytes,2,opt,name=FilePath,proto3" json:"FilePath,omitempty"`
	FileWeight string                 `protobuf:"bytes,3,opt,name=FileWeight,proto3" json:"FileWeight,omitempty"`
	Date       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=Date,proto3" json:"Date,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_memorycloud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_memorycloud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_proto_memorycloud_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *CreateRequest) GetFileWeight() string {
	if x != nil {
		return x.FileWeight
	}
	return ""
}

func (x *CreateRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   int64  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=FileName,proto3" json:"FileName,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_memorycloud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_memorycloud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_proto_memorycloud_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *DeleteRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status          string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	RemainingWeight string `protobuf:"bytes,2,opt,name=RemainingWeight,proto3" json:"RemainingWeight,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_memorycloud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_memorycloud_proto_msgTypes[2]
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
	return file_proto_proto_memorycloud_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Response) GetRemainingWeight() string {
	if x != nil {
		return x.RemainingWeight
	}
	return ""
}

var File_proto_proto_memorycloud_proto protoreflect.FileDescriptor

var file_proto_proto_memorycloud_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x57, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x44,
	0x61, 0x74, 0x65, 0x22, 0x43, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x0f,
	0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x32, 0x87, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x3b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1a, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x18, 0x5a, 0x16, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_proto_memorycloud_proto_rawDescOnce sync.Once
	file_proto_proto_memorycloud_proto_rawDescData = file_proto_proto_memorycloud_proto_rawDesc
)

func file_proto_proto_memorycloud_proto_rawDescGZIP() []byte {
	file_proto_proto_memorycloud_proto_rawDescOnce.Do(func() {
		file_proto_proto_memorycloud_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_proto_memorycloud_proto_rawDescData)
	})
	return file_proto_proto_memorycloud_proto_rawDescData
}

var file_proto_proto_memorycloud_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_proto_memorycloud_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),         // 0: memorycloud.CreateRequest
	(*DeleteRequest)(nil),         // 1: memorycloud.DeleteRequest
	(*Response)(nil),              // 2: memorycloud.Response
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_proto_proto_memorycloud_proto_depIdxs = []int32{
	3, // 0: memorycloud.CreateRequest.Date:type_name -> google.protobuf.Timestamp
	0, // 1: memorycloud.MemoryCloud.Create:input_type -> memorycloud.CreateRequest
	1, // 2: memorycloud.MemoryCloud.Delete:input_type -> memorycloud.DeleteRequest
	2, // 3: memorycloud.MemoryCloud.Create:output_type -> memorycloud.Response
	2, // 4: memorycloud.MemoryCloud.Delete:output_type -> memorycloud.Response
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_proto_memorycloud_proto_init() }
func file_proto_proto_memorycloud_proto_init() {
	if File_proto_proto_memorycloud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_proto_memorycloud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_proto_proto_memorycloud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_proto_proto_memorycloud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_proto_memorycloud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_proto_memorycloud_proto_goTypes,
		DependencyIndexes: file_proto_proto_memorycloud_proto_depIdxs,
		MessageInfos:      file_proto_proto_memorycloud_proto_msgTypes,
	}.Build()
	File_proto_proto_memorycloud_proto = out.File
	file_proto_proto_memorycloud_proto_rawDesc = nil
	file_proto_proto_memorycloud_proto_goTypes = nil
	file_proto_proto_memorycloud_proto_depIdxs = nil
}
