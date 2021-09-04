// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.17.3
// source: github.com/eolinker/eosc/service/ctl.proto

package service

import (
	proto "github.com/golang/protobuf/proto"
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

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BroadcastIP   string `protobuf:"bytes,1,opt,name=broadcastIP,proto3" json:"broadcastIP,omitempty"`
	BroadcastPort int32  `protobuf:"varint,2,opt,name=broadcastPort,proto3" json:"broadcastPort,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_github_com_eolinker_eosc_service_ctl_proto_rawDescGZIP(), []int{0}
}

func (x *JoinRequest) GetBroadcastIP() string {
	if x != nil {
		return x.BroadcastIP
	}
	return ""
}

func (x *JoinRequest) GetBroadcastPort() int32 {
	if x != nil {
		return x.BroadcastPort
	}
	return 0
}

type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string    `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Code string    `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Info *NodeInfo `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_github_com_eolinker_eosc_service_ctl_proto_rawDescGZIP(), []int{1}
}

func (x *JoinResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *JoinResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *JoinResponse) GetInfo() *NodeInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeKey string `protobuf:"bytes,1,opt,name=nodeKey,proto3" json:"nodeKey,omitempty"`
	NodeID  int32  `protobuf:"varint,2,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_github_com_eolinker_eosc_service_ctl_proto_rawDescGZIP(), []int{2}
}

func (x *NodeInfo) GetNodeKey() string {
	if x != nil {
		return x.NodeKey
	}
	return ""
}

func (x *NodeInfo) GetNodeID() int32 {
	if x != nil {
		return x.NodeID
	}
	return 0
}

type LeaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *NodeInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *LeaveRequest) Reset() {
	*x = LeaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRequest) ProtoMessage() {}

func (x *LeaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRequest.ProtoReflect.Descriptor instead.
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return file_github_com_eolinker_eosc_service_ctl_proto_rawDescGZIP(), []int{3}
}

func (x *LeaveRequest) GetInfo() *NodeInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type LeaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *LeaveResponse) Reset() {
	*x = LeaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveResponse) ProtoMessage() {}

func (x *LeaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveResponse.ProtoReflect.Descriptor instead.
func (*LeaveResponse) Descriptor() ([]byte, []int) {
	return file_github_com_eolinker_eosc_service_ctl_proto_rawDescGZIP(), []int{4}
}

func (x *LeaveResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LeaveResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_github_com_eolinker_eosc_service_ctl_proto_rawDescGZIP(), []int{5}
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_github_com_eolinker_eosc_service_ctl_proto_rawDescGZIP(), []int{6}
}

type ClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterRequest) Reset() {
	*x = ClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterRequest) ProtoMessage() {}

func (x *ClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterRequest.ProtoReflect.Descriptor instead.
func (*ClusterRequest) Descriptor() ([]byte, []int) {
	return file_github_com_eolinker_eosc_service_ctl_proto_rawDescGZIP(), []int{7}
}

type ClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterResponse) Reset() {
	*x = ClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterResponse) ProtoMessage() {}

func (x *ClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterResponse.ProtoReflect.Descriptor instead.
func (*ClusterResponse) Descriptor() ([]byte, []int) {
	return file_github_com_eolinker_eosc_service_ctl_proto_rawDescGZIP(), []int{8}
}

type StartConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config []*EnvConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
}

func (x *StartConfig) Reset() {
	*x = StartConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartConfig) ProtoMessage() {}

func (x *StartConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartConfig.ProtoReflect.Descriptor instead.
func (*StartConfig) Descriptor() ([]byte, []int) {
	return file_github_com_eolinker_eosc_service_ctl_proto_rawDescGZIP(), []int{9}
}

func (x *StartConfig) GetConfig() []*EnvConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type EnvConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EnvConfig) Reset() {
	*x = EnvConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvConfig) ProtoMessage() {}

func (x *EnvConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvConfig.ProtoReflect.Descriptor instead.
func (*EnvConfig) Descriptor() ([]byte, []int) {
	return file_github_com_eolinker_eosc_service_ctl_proto_rawDescGZIP(), []int{10}
}

func (x *EnvConfig) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *EnvConfig) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_github_com_eolinker_eosc_service_ctl_proto protoreflect.FileDescriptor

var file_github_com_eolinker_eosc_service_ctl_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c,
	0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2f, 0x65, 0x6f, 0x73, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x63, 0x74, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x55, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x49, 0x50, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x62,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x5b, 0x0a, 0x0c,
	0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x3c, 0x0a, 0x08, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x22, 0x35, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x35,
	0x0a, 0x0d, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a, 0x0b, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x45, 0x6e, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0x33, 0x0a, 0x09, 0x45, 0x6e, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x7d, 0x0a, 0x0a, 0x43, 0x74, 0x69,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12,
	0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38,
	0x0a, 0x05, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2f,
	0x65, 0x6f, 0x73, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_eolinker_eosc_service_ctl_proto_rawDescOnce sync.Once
	file_github_com_eolinker_eosc_service_ctl_proto_rawDescData = file_github_com_eolinker_eosc_service_ctl_proto_rawDesc
)

func file_github_com_eolinker_eosc_service_ctl_proto_rawDescGZIP() []byte {
	file_github_com_eolinker_eosc_service_ctl_proto_rawDescOnce.Do(func() {
		file_github_com_eolinker_eosc_service_ctl_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_eolinker_eosc_service_ctl_proto_rawDescData)
	})
	return file_github_com_eolinker_eosc_service_ctl_proto_rawDescData
}

var file_github_com_eolinker_eosc_service_ctl_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_github_com_eolinker_eosc_service_ctl_proto_goTypes = []interface{}{
	(*JoinRequest)(nil),     // 0: service.JoinRequest
	(*JoinResponse)(nil),    // 1: service.JoinResponse
	(*NodeInfo)(nil),        // 2: service.NodeInfo
	(*LeaveRequest)(nil),    // 3: service.LeaveRequest
	(*LeaveResponse)(nil),   // 4: service.LeaveResponse
	(*InfoRequest)(nil),     // 5: service.InfoRequest
	(*InfoResponse)(nil),    // 6: service.InfoResponse
	(*ClusterRequest)(nil),  // 7: service.ClusterRequest
	(*ClusterResponse)(nil), // 8: service.ClusterResponse
	(*StartConfig)(nil),     // 9: service.StartConfig
	(*EnvConfig)(nil),       // 10: service.EnvConfig
}
var file_github_com_eolinker_eosc_service_ctl_proto_depIdxs = []int32{
	2,  // 0: service.JoinResponse.info:type_name -> service.NodeInfo
	2,  // 1: service.LeaveRequest.info:type_name -> service.NodeInfo
	10, // 2: service.StartConfig.config:type_name -> service.EnvConfig
	0,  // 3: service.CtiService.Join:input_type -> service.JoinRequest
	3,  // 4: service.CtiService.Leave:input_type -> service.LeaveRequest
	1,  // 5: service.CtiService.Join:output_type -> service.JoinResponse
	4,  // 6: service.CtiService.Leave:output_type -> service.LeaveResponse
	5,  // [5:7] is the sub-list for method output_type
	3,  // [3:5] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_eolinker_eosc_service_ctl_proto_init() }
func file_github_com_eolinker_eosc_service_ctl_proto_init() {
	if File_github_com_eolinker_eosc_service_ctl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinResponse); i {
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
		file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
		file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveRequest); i {
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
		file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveResponse); i {
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
		file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
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
		file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
		file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterRequest); i {
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
		file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterResponse); i {
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
		file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartConfig); i {
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
		file_github_com_eolinker_eosc_service_ctl_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvConfig); i {
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
			RawDescriptor: file_github_com_eolinker_eosc_service_ctl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_eolinker_eosc_service_ctl_proto_goTypes,
		DependencyIndexes: file_github_com_eolinker_eosc_service_ctl_proto_depIdxs,
		MessageInfos:      file_github_com_eolinker_eosc_service_ctl_proto_msgTypes,
	}.Build()
	File_github_com_eolinker_eosc_service_ctl_proto = out.File
	file_github_com_eolinker_eosc_service_ctl_proto_rawDesc = nil
	file_github_com_eolinker_eosc_service_ctl_proto_goTypes = nil
	file_github_com_eolinker_eosc_service_ctl_proto_depIdxs = nil
}