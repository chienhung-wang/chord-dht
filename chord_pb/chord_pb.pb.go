// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: chord_pb/chord_pb.proto

package chord_pb

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

type GetFingersResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids   [][]byte `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	Addrs []string `protobuf:"bytes,2,rep,name=addrs,proto3" json:"addrs,omitempty"`
}

func (x *GetFingersResp) Reset() {
	*x = GetFingersResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_pb_chord_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFingersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFingersResp) ProtoMessage() {}

func (x *GetFingersResp) ProtoReflect() protoreflect.Message {
	mi := &file_chord_pb_chord_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFingersResp.ProtoReflect.Descriptor instead.
func (*GetFingersResp) Descriptor() ([]byte, []int) {
	return file_chord_pb_chord_pb_proto_rawDescGZIP(), []int{0}
}

func (x *GetFingersResp) GetIds() [][]byte {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GetFingersResp) GetAddrs() []string {
	if x != nil {
		return x.Addrs
	}
	return nil
}

type KeyVal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *KeyVal) Reset() {
	*x = KeyVal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_pb_chord_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyVal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyVal) ProtoMessage() {}

func (x *KeyVal) ProtoReflect() protoreflect.Message {
	mi := &file_chord_pb_chord_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyVal.ProtoReflect.Descriptor instead.
func (*KeyVal) Descriptor() ([]byte, []int) {
	return file_chord_pb_chord_pb_proto_rawDescGZIP(), []int{1}
}

func (x *KeyVal) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyVal) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_pb_chord_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_chord_pb_chord_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_chord_pb_chord_pb_proto_rawDescGZIP(), []int{2}
}

func (x *Key) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type NodeId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NodeId) Reset() {
	*x = NodeId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_pb_chord_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeId) ProtoMessage() {}

func (x *NodeId) ProtoReflect() protoreflect.Message {
	mi := &file_chord_pb_chord_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeId.ProtoReflect.Descriptor instead.
func (*NodeId) Descriptor() ([]byte, []int) {
	return file_chord_pb_chord_pb_proto_rawDescGZIP(), []int{3}
}

func (x *NodeId) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type NodeAddr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Addr string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *NodeAddr) Reset() {
	*x = NodeAddr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_pb_chord_pb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeAddr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeAddr) ProtoMessage() {}

func (x *NodeAddr) ProtoReflect() protoreflect.Message {
	mi := &file_chord_pb_chord_pb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeAddr.ProtoReflect.Descriptor instead.
func (*NodeAddr) Descriptor() ([]byte, []int) {
	return file_chord_pb_chord_pb_proto_rawDescGZIP(), []int{4}
}

func (x *NodeAddr) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *NodeAddr) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type FindFindSuccessorResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found bool      `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	Addr  *NodeAddr `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *FindFindSuccessorResp) Reset() {
	*x = FindFindSuccessorResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_pb_chord_pb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindFindSuccessorResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindFindSuccessorResp) ProtoMessage() {}

func (x *FindFindSuccessorResp) ProtoReflect() protoreflect.Message {
	mi := &file_chord_pb_chord_pb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindFindSuccessorResp.ProtoReflect.Descriptor instead.
func (*FindFindSuccessorResp) Descriptor() ([]byte, []int) {
	return file_chord_pb_chord_pb_proto_rawDescGZIP(), []int{5}
}

func (x *FindFindSuccessorResp) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

func (x *FindFindSuccessorResp) GetAddr() *NodeAddr {
	if x != nil {
		return x.Addr
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Empty bool `protobuf:"varint,1,opt,name=empty,proto3" json:"empty,omitempty"`
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chord_pb_chord_pb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_chord_pb_chord_pb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_chord_pb_chord_pb_proto_rawDescGZIP(), []int{6}
}

func (x *Empty) GetEmpty() bool {
	if x != nil {
		return x.Empty
	}
	return false
}

var File_chord_pb_chord_pb_proto protoreflect.FileDescriptor

var file_chord_pb_chord_pb_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x6f, 0x72, 0x64,
	0x5f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x68, 0x6f, 0x72, 0x64,
	0x5f, 0x70, 0x62, 0x22, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x22, 0x2c, 0x0a,
	0x06, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x17, 0x0a, 0x03, 0x4b,
	0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x18, 0x0a, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32,
	0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x22, 0x55, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x12, 0x26, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0x1d, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xe6, 0x03, 0x0a, 0x09, 0x43, 0x68, 0x6f,
	0x72, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x12,
	0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x64,
	0x64, 0x72, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x10, 0x2e, 0x63, 0x68, 0x6f, 0x72,
	0x64, 0x5f, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x1a, 0x1f, 0x2e, 0x63, 0x68,
	0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x46, 0x69, 0x6e, 0x64, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x12, 0x0f, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x12, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x1a, 0x0f, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x73, 0x12, 0x0f, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70,
	0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x70, 0x46, 0x69, 0x6e, 0x64, 0x12,
	0x10, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06, 0x4d, 0x61, 0x70, 0x47, 0x65,
	0x74, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79,
	0x1a, 0x10, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x06, 0x4d, 0x61, 0x70, 0x50, 0x75, 0x74, 0x12, 0x10,
	0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c,
	0x1a, 0x10, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x09, 0x4d, 0x61, 0x70, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79,
	0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x22,
	0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chord_pb_chord_pb_proto_rawDescOnce sync.Once
	file_chord_pb_chord_pb_proto_rawDescData = file_chord_pb_chord_pb_proto_rawDesc
)

func file_chord_pb_chord_pb_proto_rawDescGZIP() []byte {
	file_chord_pb_chord_pb_proto_rawDescOnce.Do(func() {
		file_chord_pb_chord_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_chord_pb_chord_pb_proto_rawDescData)
	})
	return file_chord_pb_chord_pb_proto_rawDescData
}

var file_chord_pb_chord_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_chord_pb_chord_pb_proto_goTypes = []interface{}{
	(*GetFingersResp)(nil),        // 0: chord_pb.GetFingersResp
	(*KeyVal)(nil),                // 1: chord_pb.KeyVal
	(*Key)(nil),                   // 2: chord_pb.Key
	(*NodeId)(nil),                // 3: chord_pb.NodeId
	(*NodeAddr)(nil),              // 4: chord_pb.NodeAddr
	(*FindFindSuccessorResp)(nil), // 5: chord_pb.FindFindSuccessorResp
	(*Empty)(nil),                 // 6: chord_pb.Empty
}
var file_chord_pb_chord_pb_proto_depIdxs = []int32{
	4,  // 0: chord_pb.FindFindSuccessorResp.addr:type_name -> chord_pb.NodeAddr
	4,  // 1: chord_pb.ChordNode.Join:input_type -> chord_pb.NodeAddr
	3,  // 2: chord_pb.ChordNode.FindSuccessor:input_type -> chord_pb.NodeId
	6,  // 3: chord_pb.ChordNode.GetPredecessor:input_type -> chord_pb.Empty
	4,  // 4: chord_pb.ChordNode.Notify:input_type -> chord_pb.NodeAddr
	6,  // 5: chord_pb.ChordNode.GetFingers:input_type -> chord_pb.Empty
	3,  // 6: chord_pb.ChordNode.HelpFind:input_type -> chord_pb.NodeId
	2,  // 7: chord_pb.ChordNode.MapGet:input_type -> chord_pb.Key
	1,  // 8: chord_pb.ChordNode.MapPut:input_type -> chord_pb.KeyVal
	2,  // 9: chord_pb.ChordNode.MapDelete:input_type -> chord_pb.Key
	4,  // 10: chord_pb.ChordNode.Join:output_type -> chord_pb.NodeAddr
	5,  // 11: chord_pb.ChordNode.FindSuccessor:output_type -> chord_pb.FindFindSuccessorResp
	4,  // 12: chord_pb.ChordNode.GetPredecessor:output_type -> chord_pb.NodeAddr
	6,  // 13: chord_pb.ChordNode.Notify:output_type -> chord_pb.Empty
	0,  // 14: chord_pb.ChordNode.GetFingers:output_type -> chord_pb.GetFingersResp
	4,  // 15: chord_pb.ChordNode.HelpFind:output_type -> chord_pb.NodeAddr
	1,  // 16: chord_pb.ChordNode.MapGet:output_type -> chord_pb.KeyVal
	1,  // 17: chord_pb.ChordNode.MapPut:output_type -> chord_pb.KeyVal
	2,  // 18: chord_pb.ChordNode.MapDelete:output_type -> chord_pb.Key
	10, // [10:19] is the sub-list for method output_type
	1,  // [1:10] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_chord_pb_chord_pb_proto_init() }
func file_chord_pb_chord_pb_proto_init() {
	if File_chord_pb_chord_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chord_pb_chord_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFingersResp); i {
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
		file_chord_pb_chord_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyVal); i {
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
		file_chord_pb_chord_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
		file_chord_pb_chord_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeId); i {
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
		file_chord_pb_chord_pb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeAddr); i {
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
		file_chord_pb_chord_pb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindFindSuccessorResp); i {
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
		file_chord_pb_chord_pb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_chord_pb_chord_pb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chord_pb_chord_pb_proto_goTypes,
		DependencyIndexes: file_chord_pb_chord_pb_proto_depIdxs,
		MessageInfos:      file_chord_pb_chord_pb_proto_msgTypes,
	}.Build()
	File_chord_pb_chord_pb_proto = out.File
	file_chord_pb_chord_pb_proto_rawDesc = nil
	file_chord_pb_chord_pb_proto_goTypes = nil
	file_chord_pb_chord_pb_proto_depIdxs = nil
}
