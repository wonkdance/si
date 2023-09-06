// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: test.proto

package protos

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

// The request message containing the user's name.
type StudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int64  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *StudentRequest) Reset() {
	*x = StudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentRequest) ProtoMessage() {}

func (x *StudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentRequest.ProtoReflect.Descriptor instead.
func (*StudentRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *StudentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StudentRequest) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

// The response message containing the greetings
type StudentEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age         int64                  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	DateTime    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=dateTime,proto3" json:"dateTime,omitempty"`
	DoubleValue float64                `protobuf:"fixed64,4,opt,name=doubleValue,proto3" json:"doubleValue,omitempty"`
}

func (x *StudentEntity) Reset() {
	*x = StudentEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentEntity) ProtoMessage() {}

func (x *StudentEntity) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentEntity.ProtoReflect.Descriptor instead.
func (*StudentEntity) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *StudentEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StudentEntity) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *StudentEntity) GetDateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DateTime
	}
	return nil
}

func (x *StudentEntity) GetDoubleValue() float64 {
	if x != nil {
		return x.DoubleValue
	}
	return 0
}

type StudentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    int64            `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message   *string          `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	Count     *int64           `protobuf:"varint,3,opt,name=count,proto3,oneof" json:"count,omitempty"`
	Offset    *int64           `protobuf:"varint,4,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
	Limit     *int64           `protobuf:"varint,5,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Document  *StudentEntity   `protobuf:"bytes,6,opt,name=document,proto3,oneof" json:"document,omitempty"`
	Documents []*StudentEntity `protobuf:"bytes,7,rep,name=documents,proto3" json:"documents,omitempty"`
}

func (x *StudentReply) Reset() {
	*x = StudentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentReply) ProtoMessage() {}

func (x *StudentReply) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentReply.ProtoReflect.Descriptor instead.
func (*StudentReply) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

func (x *StudentReply) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *StudentReply) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *StudentReply) GetCount() int64 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *StudentReply) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *StudentReply) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *StudentReply) GetDocument() *StudentEntity {
	if x != nil {
		return x.Document
	}
	return nil
}

func (x *StudentReply) GetDocuments() []*StudentEntity {
	if x != nil {
		return x.Documents
	}
	return nil
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a,
	0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xaf, 0x02, 0x0a, 0x0c, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x2f, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x48, 0x04, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0x33, 0x0a, 0x07, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0f, 0x2e, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x28,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_test_proto_goTypes = []interface{}{
	(*StudentRequest)(nil),        // 0: StudentRequest
	(*StudentEntity)(nil),         // 1: StudentEntity
	(*StudentReply)(nil),          // 2: StudentReply
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_test_proto_depIdxs = []int32{
	3, // 0: StudentEntity.dateTime:type_name -> google.protobuf.Timestamp
	1, // 1: StudentReply.document:type_name -> StudentEntity
	1, // 2: StudentReply.documents:type_name -> StudentEntity
	0, // 3: Student.Read:input_type -> StudentRequest
	2, // 4: Student.Read:output_type -> StudentReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentRequest); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentEntity); i {
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
		file_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentReply); i {
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
	file_test_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
