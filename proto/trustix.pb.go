// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: trustix.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	schema "github.com/tweag/trustix/schema"
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

type KeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,req,name=Key" json:"Key,omitempty"`
}

func (x *KeyRequest) Reset() {
	*x = KeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustix_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRequest) ProtoMessage() {}

func (x *KeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trustix_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRequest.ProtoReflect.Descriptor instead.
func (*KeyRequest) Descriptor() ([]byte, []int) {
	return file_trustix_proto_rawDescGZIP(), []int{0}
}

func (x *KeyRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type EntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     []byte                      `protobuf:"bytes,1,req,name=Key" json:"Key,omitempty"`
	Entries map[string]*schema.MapEntry `protobuf:"bytes,2,rep,name=Entries" json:"Entries,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (x *EntriesResponse) Reset() {
	*x = EntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustix_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntriesResponse) ProtoMessage() {}

func (x *EntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trustix_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntriesResponse.ProtoReflect.Descriptor instead.
func (*EntriesResponse) Descriptor() ([]byte, []int) {
	return file_trustix_proto_rawDescGZIP(), []int{1}
}

func (x *EntriesResponse) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *EntriesResponse) GetEntries() map[string]*schema.MapEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type LogValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogName *string `protobuf:"bytes,1,req,name=LogName" json:"LogName,omitempty"`
	Value   []byte  `protobuf:"bytes,2,req,name=Value" json:"Value,omitempty"`
}

func (x *LogValueResponse) Reset() {
	*x = LogValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustix_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogValueResponse) ProtoMessage() {}

func (x *LogValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trustix_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogValueResponse.ProtoReflect.Descriptor instead.
func (*LogValueResponse) Descriptor() ([]byte, []int) {
	return file_trustix_proto_rawDescGZIP(), []int{2}
}

func (x *LogValueResponse) GetLogName() string {
	if x != nil && x.LogName != nil {
		return *x.LogName
	}
	return ""
}

func (x *LogValueResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type LogValueDecision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogNames   []string `protobuf:"bytes,1,rep,name=LogNames" json:"LogNames,omitempty"`
	Value      []byte   `protobuf:"bytes,2,req,name=Value" json:"Value,omitempty"`
	Confidence *int32   `protobuf:"varint,3,req,name=Confidence" json:"Confidence,omitempty"`
}

func (x *LogValueDecision) Reset() {
	*x = LogValueDecision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustix_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogValueDecision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogValueDecision) ProtoMessage() {}

func (x *LogValueDecision) ProtoReflect() protoreflect.Message {
	mi := &file_trustix_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogValueDecision.ProtoReflect.Descriptor instead.
func (*LogValueDecision) Descriptor() ([]byte, []int) {
	return file_trustix_proto_rawDescGZIP(), []int{3}
}

func (x *LogValueDecision) GetLogNames() []string {
	if x != nil {
		return x.LogNames
	}
	return nil
}

func (x *LogValueDecision) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *LogValueDecision) GetConfidence() int32 {
	if x != nil && x.Confidence != nil {
		return *x.Confidence
	}
	return 0
}

type DecisionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Decision *LogValueDecision `protobuf:"bytes,1,req,name=Decision" json:"Decision,omitempty"`
	// Non-matches (hash mismatch)
	Mismatches []*LogValueResponse `protobuf:"bytes,2,rep,name=Mismatches" json:"Mismatches,omitempty"`
	// Full misses (lognames missing log entry entirely)
	Misses []string `protobuf:"bytes,3,rep,name=Misses" json:"Misses,omitempty"`
}

func (x *DecisionResponse) Reset() {
	*x = DecisionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trustix_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecisionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecisionResponse) ProtoMessage() {}

func (x *DecisionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trustix_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecisionResponse.ProtoReflect.Descriptor instead.
func (*DecisionResponse) Descriptor() ([]byte, []int) {
	return file_trustix_proto_rawDescGZIP(), []int{4}
}

func (x *DecisionResponse) GetDecision() *LogValueDecision {
	if x != nil {
		return x.Decision
	}
	return nil
}

func (x *DecisionResponse) GetMismatches() []*LogValueResponse {
	if x != nil {
		return x.Mismatches
	}
	return nil
}

func (x *DecisionResponse) GetMisses() []string {
	if x != nil {
		return x.Misses
	}
	return nil
}

var File_trustix_proto protoreflect.FileDescriptor

var file_trustix_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x1a, 0x15, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2f, 0x6d, 0x61, 0x70, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1e, 0x0a, 0x0a, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x22,
	0xab, 0x01, 0x0a, 0x0f, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c,
	0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x3f, 0x0a, 0x07, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78,
	0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x45, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x42, 0x0a,
	0x10, 0x4c, 0x6f, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x07, 0x4c, 0x6f, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x64, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x67, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0c,
	0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08,
	0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x4c, 0x6f, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x44, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x4d, 0x69, 0x73, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69,
	0x78, 0x2e, 0x4c, 0x6f, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0a, 0x4d, 0x69, 0x73, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x4d, 0x69, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x4d, 0x69, 0x73, 0x73, 0x65, 0x73, 0x32, 0x90, 0x02, 0x0a, 0x12, 0x54, 0x72, 0x75, 0x73, 0x74,
	0x69, 0x78, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x64, 0x52, 0x50, 0x43, 0x12, 0x36, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x69, 0x78, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x13, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69,
	0x78, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x06, 0x44, 0x65, 0x63, 0x69, 0x64,
	0x65, 0x12, 0x13, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78,
	0x2e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c, 0x44, 0x65, 0x63, 0x69, 0x64, 0x65, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x13, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2e, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74,
	0x69, 0x78, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x77, 0x65, 0x61, 0x67, 0x2f, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x69, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_trustix_proto_rawDescOnce sync.Once
	file_trustix_proto_rawDescData = file_trustix_proto_rawDesc
)

func file_trustix_proto_rawDescGZIP() []byte {
	file_trustix_proto_rawDescOnce.Do(func() {
		file_trustix_proto_rawDescData = protoimpl.X.CompressGZIP(file_trustix_proto_rawDescData)
	})
	return file_trustix_proto_rawDescData
}

var file_trustix_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_trustix_proto_goTypes = []interface{}{
	(*KeyRequest)(nil),       // 0: trustix.KeyRequest
	(*EntriesResponse)(nil),  // 1: trustix.EntriesResponse
	(*LogValueResponse)(nil), // 2: trustix.LogValueResponse
	(*LogValueDecision)(nil), // 3: trustix.LogValueDecision
	(*DecisionResponse)(nil), // 4: trustix.DecisionResponse
	nil,                      // 5: trustix.EntriesResponse.EntriesEntry
	(*schema.MapEntry)(nil),  // 6: MapEntry
}
var file_trustix_proto_depIdxs = []int32{
	5, // 0: trustix.EntriesResponse.Entries:type_name -> trustix.EntriesResponse.EntriesEntry
	3, // 1: trustix.DecisionResponse.Decision:type_name -> trustix.LogValueDecision
	2, // 2: trustix.DecisionResponse.Mismatches:type_name -> trustix.LogValueResponse
	6, // 3: trustix.EntriesResponse.EntriesEntry.value:type_name -> MapEntry
	0, // 4: trustix.TrustixCombinedRPC.Get:input_type -> trustix.KeyRequest
	0, // 5: trustix.TrustixCombinedRPC.GetStream:input_type -> trustix.KeyRequest
	0, // 6: trustix.TrustixCombinedRPC.Decide:input_type -> trustix.KeyRequest
	0, // 7: trustix.TrustixCombinedRPC.DecideStream:input_type -> trustix.KeyRequest
	1, // 8: trustix.TrustixCombinedRPC.Get:output_type -> trustix.EntriesResponse
	1, // 9: trustix.TrustixCombinedRPC.GetStream:output_type -> trustix.EntriesResponse
	4, // 10: trustix.TrustixCombinedRPC.Decide:output_type -> trustix.DecisionResponse
	4, // 11: trustix.TrustixCombinedRPC.DecideStream:output_type -> trustix.DecisionResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_trustix_proto_init() }
func file_trustix_proto_init() {
	if File_trustix_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trustix_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRequest); i {
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
		file_trustix_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntriesResponse); i {
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
		file_trustix_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogValueResponse); i {
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
		file_trustix_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogValueDecision); i {
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
		file_trustix_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecisionResponse); i {
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
			RawDescriptor: file_trustix_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trustix_proto_goTypes,
		DependencyIndexes: file_trustix_proto_depIdxs,
		MessageInfos:      file_trustix_proto_msgTypes,
	}.Build()
	File_trustix_proto = out.File
	file_trustix_proto_rawDesc = nil
	file_trustix_proto_goTypes = nil
	file_trustix_proto_depIdxs = nil
}
