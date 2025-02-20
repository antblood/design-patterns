// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: subject.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The message containing the state change information.
type AddRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	mi := &file_subject_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subject_proto_msgTypes[0]
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
	return file_subject_proto_rawDescGZIP(), []int{0}
}

// The response message containing confirmation of the notification.
type AddResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	mi := &file_subject_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subject_proto_msgTypes[1]
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
	return file_subject_proto_rawDescGZIP(), []int{1}
}

// The message containing the state change information.
type RemoveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveRequest) Reset() {
	*x = RemoveRequest{}
	mi := &file_subject_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRequest) ProtoMessage() {}

func (x *RemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subject_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRequest.ProtoReflect.Descriptor instead.
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return file_subject_proto_rawDescGZIP(), []int{2}
}

// The response message containing confirmation of the notification.
type RemoveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveResponse) Reset() {
	*x = RemoveResponse{}
	mi := &file_subject_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveResponse) ProtoMessage() {}

func (x *RemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subject_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveResponse.ProtoReflect.Descriptor instead.
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return file_subject_proto_rawDescGZIP(), []int{3}
}

// The message containing the state change information.
type LatestStateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LatestStateRequest) Reset() {
	*x = LatestStateRequest{}
	mi := &file_subject_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LatestStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatestStateRequest) ProtoMessage() {}

func (x *LatestStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subject_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatestStateRequest.ProtoReflect.Descriptor instead.
func (*LatestStateRequest) Descriptor() ([]byte, []int) {
	return file_subject_proto_rawDescGZIP(), []int{4}
}

// The response message containing confirmation of the notification.
type LatestStateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	State         string                 `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LatestStateResponse) Reset() {
	*x = LatestStateResponse{}
	mi := &file_subject_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LatestStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatestStateResponse) ProtoMessage() {}

func (x *LatestStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subject_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatestStateResponse.ProtoReflect.Descriptor instead.
func (*LatestStateResponse) Descriptor() ([]byte, []int) {
	return file_subject_proto_rawDescGZIP(), []int{5}
}

func (x *LatestStateResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type UpdateStateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	State         string                 `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStateRequest) Reset() {
	*x = UpdateStateRequest{}
	mi := &file_subject_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStateRequest) ProtoMessage() {}

func (x *UpdateStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subject_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStateRequest.ProtoReflect.Descriptor instead.
func (*UpdateStateRequest) Descriptor() ([]byte, []int) {
	return file_subject_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateStateRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type UpdateStateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStateResponse) Reset() {
	*x = UpdateStateResponse{}
	mi := &file_subject_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStateResponse) ProtoMessage() {}

func (x *UpdateStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subject_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStateResponse.ProtoReflect.Descriptor instead.
func (*UpdateStateResponse) Descriptor() ([]byte, []int) {
	return file_subject_proto_rawDescGZIP(), []int{7}
}

type BroadcastRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BroadcastRequest) Reset() {
	*x = BroadcastRequest{}
	mi := &file_subject_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BroadcastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRequest) ProtoMessage() {}

func (x *BroadcastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subject_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRequest.ProtoReflect.Descriptor instead.
func (*BroadcastRequest) Descriptor() ([]byte, []int) {
	return file_subject_proto_rawDescGZIP(), []int{8}
}

type BroadcastResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BroadcastResponse) Reset() {
	*x = BroadcastResponse{}
	mi := &file_subject_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BroadcastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastResponse) ProtoMessage() {}

func (x *BroadcastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subject_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastResponse.ProtoReflect.Descriptor instead.
func (*BroadcastResponse) Descriptor() ([]byte, []int) {
	return file_subject_proto_rawDescGZIP(), []int{9}
}

var File_subject_proto protoreflect.FileDescriptor

var file_subject_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b,
	0x0a, 0x13, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x2a, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12,
	0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xce, 0x02, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x30, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x13, 0x2e, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12,
	0x16, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x48, 0x0a, 0x0b, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x1b, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x12, 0x19, 0x2e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x6f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2d, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_subject_proto_rawDescOnce sync.Once
	file_subject_proto_rawDescData []byte
)

func file_subject_proto_rawDescGZIP() []byte {
	file_subject_proto_rawDescOnce.Do(func() {
		file_subject_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_subject_proto_rawDesc), len(file_subject_proto_rawDesc)))
	})
	return file_subject_proto_rawDescData
}

var file_subject_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_subject_proto_goTypes = []any{
	(*AddRequest)(nil),          // 0: subject.AddRequest
	(*AddResponse)(nil),         // 1: subject.AddResponse
	(*RemoveRequest)(nil),       // 2: subject.RemoveRequest
	(*RemoveResponse)(nil),      // 3: subject.RemoveResponse
	(*LatestStateRequest)(nil),  // 4: subject.LatestStateRequest
	(*LatestStateResponse)(nil), // 5: subject.LatestStateResponse
	(*UpdateStateRequest)(nil),  // 6: subject.UpdateStateRequest
	(*UpdateStateResponse)(nil), // 7: subject.UpdateStateResponse
	(*BroadcastRequest)(nil),    // 8: subject.BroadcastRequest
	(*BroadcastResponse)(nil),   // 9: subject.BroadcastResponse
}
var file_subject_proto_depIdxs = []int32{
	0, // 0: subject.Subject.Add:input_type -> subject.AddRequest
	2, // 1: subject.Subject.Remove:input_type -> subject.RemoveRequest
	4, // 2: subject.Subject.LatestState:input_type -> subject.LatestStateRequest
	6, // 3: subject.Subject.UpdateState:input_type -> subject.UpdateStateRequest
	8, // 4: subject.Subject.Broadcast:input_type -> subject.BroadcastRequest
	1, // 5: subject.Subject.Add:output_type -> subject.AddResponse
	3, // 6: subject.Subject.Remove:output_type -> subject.RemoveResponse
	5, // 7: subject.Subject.LatestState:output_type -> subject.LatestStateResponse
	7, // 8: subject.Subject.UpdateState:output_type -> subject.UpdateStateResponse
	9, // 9: subject.Subject.Broadcast:output_type -> subject.BroadcastResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_subject_proto_init() }
func file_subject_proto_init() {
	if File_subject_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_subject_proto_rawDesc), len(file_subject_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_subject_proto_goTypes,
		DependencyIndexes: file_subject_proto_depIdxs,
		MessageInfos:      file_subject_proto_msgTypes,
	}.Build()
	File_subject_proto = out.File
	file_subject_proto_goTypes = nil
	file_subject_proto_depIdxs = nil
}
