// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: extender.proto

package service

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

type ExtendsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extends []*ExtendsBasicInfo `protobuf:"bytes,1,rep,name=extends,proto3" json:"extends,omitempty"`
}

func (x *ExtendsRequest) Reset() {
	*x = ExtendsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extender_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendsRequest) ProtoMessage() {}

func (x *ExtendsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_extender_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendsRequest.ProtoReflect.Descriptor instead.
func (*ExtendsRequest) Descriptor() ([]byte, []int) {
	return file_extender_proto_rawDescGZIP(), []int{0}
}

func (x *ExtendsRequest) GetExtends() []*ExtendsBasicInfo {
	if x != nil {
		return x.Extends
	}
	return nil
}

type ExtendsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg         string              `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Code        string              `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Extends     []*ExtendsInfo      `protobuf:"bytes,3,rep,name=extends,proto3" json:"extends,omitempty"`
	FailExtends []*ExtendsBasicInfo `protobuf:"bytes,4,rep,name=failExtends,proto3" json:"failExtends,omitempty"`
}

func (x *ExtendsResponse) Reset() {
	*x = ExtendsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extender_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendsResponse) ProtoMessage() {}

func (x *ExtendsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_extender_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendsResponse.ProtoReflect.Descriptor instead.
func (*ExtendsResponse) Descriptor() ([]byte, []int) {
	return file_extender_proto_rawDescGZIP(), []int{1}
}

func (x *ExtendsResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ExtendsResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ExtendsResponse) GetExtends() []*ExtendsInfo {
	if x != nil {
		return x.Extends
	}
	return nil
}

func (x *ExtendsResponse) GetFailExtends() []*ExtendsBasicInfo {
	if x != nil {
		return x.FailExtends
	}
	return nil
}

type ExtendsBasicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Group   string `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Project string `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Msg     string `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ExtendsBasicInfo) Reset() {
	*x = ExtendsBasicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extender_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendsBasicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendsBasicInfo) ProtoMessage() {}

func (x *ExtendsBasicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_extender_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendsBasicInfo.ProtoReflect.Descriptor instead.
func (*ExtendsBasicInfo) Descriptor() ([]byte, []int) {
	return file_extender_proto_rawDescGZIP(), []int{2}
}

func (x *ExtendsBasicInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExtendsBasicInfo) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *ExtendsBasicInfo) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ExtendsBasicInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ExtendsBasicInfo) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ExtendsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Group   string    `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	Project string    `protobuf:"bytes,4,opt,name=project,proto3" json:"project,omitempty"`
	Version string    `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Plugins []*Plugin `protobuf:"bytes,6,rep,name=plugins,proto3" json:"plugins,omitempty"`
}

func (x *ExtendsInfo) Reset() {
	*x = ExtendsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extender_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendsInfo) ProtoMessage() {}

func (x *ExtendsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_extender_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendsInfo.ProtoReflect.Descriptor instead.
func (*ExtendsInfo) Descriptor() ([]byte, []int) {
	return file_extender_proto_rawDescGZIP(), []int{3}
}

func (x *ExtendsInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExtendsInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExtendsInfo) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *ExtendsInfo) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ExtendsInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ExtendsInfo) GetPlugins() []*Plugin {
	if x != nil {
		return x.Plugins
	}
	return nil
}

type ExtendsInstallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extends []*ExtendsBasicInfo `protobuf:"bytes,1,rep,name=extends,proto3" json:"extends,omitempty"`
}

func (x *ExtendsInstallRequest) Reset() {
	*x = ExtendsInstallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extender_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendsInstallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendsInstallRequest) ProtoMessage() {}

func (x *ExtendsInstallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_extender_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendsInstallRequest.ProtoReflect.Descriptor instead.
func (*ExtendsInstallRequest) Descriptor() ([]byte, []int) {
	return file_extender_proto_rawDescGZIP(), []int{4}
}

func (x *ExtendsInstallRequest) GetExtends() []*ExtendsBasicInfo {
	if x != nil {
		return x.Extends
	}
	return nil
}

type ExtendsUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extends []*ExtendsBasicInfo `protobuf:"bytes,1,rep,name=extends,proto3" json:"extends,omitempty"`
}

func (x *ExtendsUpdateRequest) Reset() {
	*x = ExtendsUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extender_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendsUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendsUpdateRequest) ProtoMessage() {}

func (x *ExtendsUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_extender_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendsUpdateRequest.ProtoReflect.Descriptor instead.
func (*ExtendsUpdateRequest) Descriptor() ([]byte, []int) {
	return file_extender_proto_rawDescGZIP(), []int{5}
}

func (x *ExtendsUpdateRequest) GetExtends() []*ExtendsBasicInfo {
	if x != nil {
		return x.Extends
	}
	return nil
}

type ExtendsUninstallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg         string              `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Code        string              `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Extends     []*ExtendsBasicInfo `protobuf:"bytes,3,rep,name=extends,proto3" json:"extends,omitempty"`
	FailExtends []*ExtendsBasicInfo `protobuf:"bytes,4,rep,name=failExtends,proto3" json:"failExtends,omitempty"`
}

func (x *ExtendsUninstallResponse) Reset() {
	*x = ExtendsUninstallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extender_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendsUninstallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendsUninstallResponse) ProtoMessage() {}

func (x *ExtendsUninstallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_extender_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendsUninstallResponse.ProtoReflect.Descriptor instead.
func (*ExtendsUninstallResponse) Descriptor() ([]byte, []int) {
	return file_extender_proto_rawDescGZIP(), []int{6}
}

func (x *ExtendsUninstallResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ExtendsUninstallResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ExtendsUninstallResponse) GetExtends() []*ExtendsBasicInfo {
	if x != nil {
		return x.Extends
	}
	return nil
}

func (x *ExtendsUninstallResponse) GetFailExtends() []*ExtendsBasicInfo {
	if x != nil {
		return x.FailExtends
	}
	return nil
}

type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // {group}:{project}:{name}
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Group   string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	Project string `protobuf:"bytes,4,opt,name=project,proto3" json:"project,omitempty"`
	Render  string `protobuf:"bytes,5,opt,name=render,proto3" json:"render,omitempty"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extender_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_extender_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_extender_proto_rawDescGZIP(), []int{7}
}

func (x *Plugin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Plugin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plugin) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Plugin) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Plugin) GetRender() string {
	if x != nil {
		return x.Render
	}
	return ""
}

var File_extender_proto protoreflect.FileDescriptor

var file_extender_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x45, 0x0a, 0x0e, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73,
	0x22, 0xa4, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x61,
	0x69, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x73, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x10, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x73, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xa6, 0x01, 0x0a,
	0x0b, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x07, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x22, 0x4c, 0x0a, 0x15, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33,
	0x0a, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x73, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x73, 0x22, 0x4b, 0x0a, 0x14, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73,
	0x22, 0xb2, 0x01, 0x0a, 0x18, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x55, 0x6e, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x74, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x22, 0x5a, 0x20, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x69, 0x6e, 0x6b,
	0x65, 0x72, 0x2f, 0x65, 0x6f, 0x73, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_extender_proto_rawDescOnce sync.Once
	file_extender_proto_rawDescData = file_extender_proto_rawDesc
)

func file_extender_proto_rawDescGZIP() []byte {
	file_extender_proto_rawDescOnce.Do(func() {
		file_extender_proto_rawDescData = protoimpl.X.CompressGZIP(file_extender_proto_rawDescData)
	})
	return file_extender_proto_rawDescData
}

var file_extender_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_extender_proto_goTypes = []interface{}{
	(*ExtendsRequest)(nil),           // 0: service.ExtendsRequest
	(*ExtendsResponse)(nil),          // 1: service.ExtendsResponse
	(*ExtendsBasicInfo)(nil),         // 2: service.ExtendsBasicInfo
	(*ExtendsInfo)(nil),              // 3: service.ExtendsInfo
	(*ExtendsInstallRequest)(nil),    // 4: service.ExtendsInstallRequest
	(*ExtendsUpdateRequest)(nil),     // 5: service.ExtendsUpdateRequest
	(*ExtendsUninstallResponse)(nil), // 6: service.ExtendsUninstallResponse
	(*Plugin)(nil),                   // 7: service.Plugin
}
var file_extender_proto_depIdxs = []int32{
	2, // 0: service.ExtendsRequest.extends:type_name -> service.ExtendsBasicInfo
	3, // 1: service.ExtendsResponse.extends:type_name -> service.ExtendsInfo
	2, // 2: service.ExtendsResponse.failExtends:type_name -> service.ExtendsBasicInfo
	7, // 3: service.ExtendsInfo.plugins:type_name -> service.Plugin
	2, // 4: service.ExtendsInstallRequest.extends:type_name -> service.ExtendsBasicInfo
	2, // 5: service.ExtendsUpdateRequest.extends:type_name -> service.ExtendsBasicInfo
	2, // 6: service.ExtendsUninstallResponse.extends:type_name -> service.ExtendsBasicInfo
	2, // 7: service.ExtendsUninstallResponse.failExtends:type_name -> service.ExtendsBasicInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_extender_proto_init() }
func file_extender_proto_init() {
	if File_extender_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_extender_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendsRequest); i {
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
		file_extender_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendsResponse); i {
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
		file_extender_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendsBasicInfo); i {
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
		file_extender_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendsInfo); i {
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
		file_extender_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendsInstallRequest); i {
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
		file_extender_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendsUpdateRequest); i {
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
		file_extender_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendsUninstallResponse); i {
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
		file_extender_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugin); i {
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
			RawDescriptor: file_extender_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_extender_proto_goTypes,
		DependencyIndexes: file_extender_proto_depIdxs,
		MessageInfos:      file_extender_proto_msgTypes,
	}.Build()
	File_extender_proto = out.File
	file_extender_proto_rawDesc = nil
	file_extender_proto_goTypes = nil
	file_extender_proto_depIdxs = nil
}
