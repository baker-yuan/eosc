// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: message.proto

package eosc

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

type ProfessionConfig_ProfessionMod int32

const (
	ProfessionConfig_Worker    ProfessionConfig_ProfessionMod = 0
	ProfessionConfig_Singleton ProfessionConfig_ProfessionMod = 1
)

// Enum value maps for ProfessionConfig_ProfessionMod.
var (
	ProfessionConfig_ProfessionMod_name = map[int32]string{
		0: "Worker",
		1: "Singleton",
	}
	ProfessionConfig_ProfessionMod_value = map[string]int32{
		"Worker":    0,
		"Singleton": 1,
	}
)

func (x ProfessionConfig_ProfessionMod) Enum() *ProfessionConfig_ProfessionMod {
	p := new(ProfessionConfig_ProfessionMod)
	*p = x
	return p
}

func (x ProfessionConfig_ProfessionMod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProfessionConfig_ProfessionMod) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (ProfessionConfig_ProfessionMod) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x ProfessionConfig_ProfessionMod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProfessionConfig_ProfessionMod.Descriptor instead.
func (ProfessionConfig_ProfessionMod) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0, 0}
}

type ProfessionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string                         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`   //
	Label        string                         `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"` //
	Desc         string                         `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`   //
	Dependencies []string                       `protobuf:"bytes,4,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	AppendLabels []string                       `protobuf:"bytes,5,rep,name=appendLabels,proto3" json:"appendLabels,omitempty"`
	Drivers      []*DriverConfig                `protobuf:"bytes,6,rep,name=drivers,proto3" json:"drivers,omitempty"`
	Mod          ProfessionConfig_ProfessionMod `protobuf:"varint,7,opt,name=mod,proto3,enum=service.ProfessionConfig_ProfessionMod" json:"mod,omitempty"`
}

func (x *ProfessionConfig) Reset() {
	*x = ProfessionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfessionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfessionConfig) ProtoMessage() {}

func (x *ProfessionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfessionConfig.ProtoReflect.Descriptor instead.
func (*ProfessionConfig) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *ProfessionConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProfessionConfig) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ProfessionConfig) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ProfessionConfig) GetDependencies() []string {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

func (x *ProfessionConfig) GetAppendLabels() []string {
	if x != nil {
		return x.AppendLabels
	}
	return nil
}

func (x *ProfessionConfig) GetDrivers() []*DriverConfig {
	if x != nil {
		return x.Drivers
	}
	return nil
}

func (x *ProfessionConfig) GetMod() ProfessionConfig_ProfessionMod {
	if x != nil {
		return x.Mod
	}
	return ProfessionConfig_Worker
}

type ProfessionConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ProfessionConfig `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ProfessionConfigs) Reset() {
	*x = ProfessionConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfessionConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfessionConfigs) ProtoMessage() {}

func (x *ProfessionConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfessionConfigs.ProtoReflect.Descriptor instead.
func (*ProfessionConfigs) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *ProfessionConfigs) GetData() []*ProfessionConfig {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProfessionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`   //
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"` //
	Desc  string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`   //
}

func (x *ProfessionInfo) Reset() {
	*x = ProfessionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfessionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfessionInfo) ProtoMessage() {}

func (x *ProfessionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfessionInfo.ProtoReflect.Descriptor instead.
func (*ProfessionInfo) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *ProfessionInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProfessionInfo) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ProfessionInfo) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type ProfessionDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Label        string          `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Desc         string          `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Dependencies []string        `protobuf:"bytes,4,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	AppendLabels []string        `protobuf:"bytes,5,rep,name=appendLabels,proto3" json:"appendLabels,omitempty"`
	Drivers      []*DriverDetail `protobuf:"bytes,6,rep,name=drivers,proto3" json:"drivers,omitempty"`
}

func (x *ProfessionDetail) Reset() {
	*x = ProfessionDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfessionDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfessionDetail) ProtoMessage() {}

func (x *ProfessionDetail) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfessionDetail.ProtoReflect.Descriptor instead.
func (*ProfessionDetail) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *ProfessionDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProfessionDetail) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ProfessionDetail) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ProfessionDetail) GetDependencies() []string {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

func (x *ProfessionDetail) GetAppendLabels() []string {
	if x != nil {
		return x.AppendLabels
	}
	return nil
}

func (x *ProfessionDetail) GetDrivers() []*DriverDetail {
	if x != nil {
		return x.Drivers
	}
	return nil
}

type DriverInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Label string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Desc  string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"` //  string profession=5;
}

func (x *DriverInfo) Reset() {
	*x = DriverInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverInfo) ProtoMessage() {}

func (x *DriverInfo) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverInfo.ProtoReflect.Descriptor instead.
func (*DriverInfo) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *DriverInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DriverInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DriverInfo) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *DriverInfo) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type DriverConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Label  string            `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Desc   string            `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	Params map[string]string `protobuf:"bytes,5,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DriverConfig) Reset() {
	*x = DriverConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverConfig) ProtoMessage() {}

func (x *DriverConfig) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverConfig.ProtoReflect.Descriptor instead.
func (*DriverConfig) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *DriverConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DriverConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DriverConfig) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *DriverConfig) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *DriverConfig) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type PluginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group   string `protobuf:"bytes,6,opt,name=group,proto3" json:"group,omitempty"`
	Project string `protobuf:"bytes,7,opt,name=project,proto3" json:"project,omitempty"`
	Name    string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"` // 插件名
}

func (x *PluginInfo) Reset() {
	*x = PluginInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginInfo) ProtoMessage() {}

func (x *PluginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginInfo.ProtoReflect.Descriptor instead.
func (*PluginInfo) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (x *PluginInfo) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *PluginInfo) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *PluginInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 驱动详情
type DriverDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Driver string `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"` // 驱动名 = driverConfig.name
	Label  string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Desc   string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	//  string profession=5;
	Params map[string]string `protobuf:"bytes,5,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Plugin *PluginInfo       `protobuf:"bytes,6,opt,name=plugin,proto3" json:"plugin,omitempty"`
}

func (x *DriverDetail) Reset() {
	*x = DriverDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverDetail) ProtoMessage() {}

func (x *DriverDetail) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverDetail.ProtoReflect.Descriptor instead.
func (*DriverDetail) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{7}
}

func (x *DriverDetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DriverDetail) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *DriverDetail) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *DriverDetail) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *DriverDetail) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *DriverDetail) GetPlugin() *PluginInfo {
	if x != nil {
		return x.Plugin
	}
	return nil
}

type WorkerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Profession string `protobuf:"bytes,2,opt,name=profession,proto3" json:"profession,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Driver     string `protobuf:"bytes,4,opt,name=driver,proto3" json:"driver,omitempty"`
	Create     string `protobuf:"bytes,5,opt,name=create,proto3" json:"create,omitempty"`
	Update     string `protobuf:"bytes,6,opt,name=update,proto3" json:"update,omitempty"`
	Body       []byte `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *WorkerConfig) Reset() {
	*x = WorkerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerConfig) ProtoMessage() {}

func (x *WorkerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerConfig.ProtoReflect.Descriptor instead.
func (*WorkerConfig) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{8}
}

func (x *WorkerConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WorkerConfig) GetProfession() string {
	if x != nil {
		return x.Profession
	}
	return ""
}

func (x *WorkerConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkerConfig) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *WorkerConfig) GetCreate() string {
	if x != nil {
		return x.Create
	}
	return ""
}

func (x *WorkerConfig) GetUpdate() string {
	if x != nil {
		return x.Update
	}
	return ""
}

func (x *WorkerConfig) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type WorkerConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*WorkerConfig `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *WorkerConfigs) Reset() {
	*x = WorkerConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerConfigs) ProtoMessage() {}

func (x *WorkerConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerConfigs.ProtoReflect.Descriptor instead.
func (*WorkerConfigs) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{9}
}

func (x *WorkerConfigs) GetData() []*WorkerConfig {
	if x != nil {
		return x.Data
	}
	return nil
}

type ExtendersSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extenders map[string]string `protobuf:"bytes,1,rep,name=Extenders,proto3" json:"Extenders,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExtendersSettings) Reset() {
	*x = ExtendersSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendersSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendersSettings) ProtoMessage() {}

func (x *ExtendersSettings) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendersSettings.ProtoReflect.Descriptor instead.
func (*ExtendersSettings) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{10}
}

func (x *ExtendersSettings) GetExtenders() map[string]string {
	if x != nil {
		return x.Extenders
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xb0, 0x02, 0x0a, 0x10, 0x50, 0x72, 0x6f,
	0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x22, 0x0a, 0x0c, 0x64,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x12, 0x39, 0x0a, 0x03, 0x6d, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x52, 0x03, 0x6d, 0x6f, 0x64, 0x22,
	0x2a, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64,
	0x12, 0x0a, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x10, 0x01, 0x22, 0x42, 0x0a, 0x11, 0x50,
	0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x4e, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22,
	0xc9, 0x01, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x70,
	0x70, 0x65, 0x6e, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x07, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x22, 0x5a, 0x0a, 0x0a, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0xd2, 0x01, 0x0a, 0x0c, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x39, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x50, 0x0a, 0x0a,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x83,
	0x02, 0x0a, 0x0c, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x12, 0x39, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2b, 0x0a, 0x06,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xae, 0x01, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x3a, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x9a, 0x01, 0x0a, 0x11, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x47, 0x0a, 0x09, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73,
	0x1a, 0x3c, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x1a,
	0x5a, 0x18, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c,
	0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2f, 0x65, 0x6f, 0x73, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_message_proto_goTypes = []interface{}{
	(ProfessionConfig_ProfessionMod)(0), // 0: service.ProfessionConfig.ProfessionMod
	(*ProfessionConfig)(nil),            // 1: service.ProfessionConfig
	(*ProfessionConfigs)(nil),           // 2: service.ProfessionConfigs
	(*ProfessionInfo)(nil),              // 3: service.ProfessionInfo
	(*ProfessionDetail)(nil),            // 4: service.ProfessionDetail
	(*DriverInfo)(nil),                  // 5: service.DriverInfo
	(*DriverConfig)(nil),                // 6: service.DriverConfig
	(*PluginInfo)(nil),                  // 7: service.PluginInfo
	(*DriverDetail)(nil),                // 8: service.DriverDetail
	(*WorkerConfig)(nil),                // 9: service.WorkerConfig
	(*WorkerConfigs)(nil),               // 10: service.WorkerConfigs
	(*ExtendersSettings)(nil),           // 11: service.ExtendersSettings
	nil,                                 // 12: service.DriverConfig.ParamsEntry
	nil,                                 // 13: service.DriverDetail.ParamsEntry
	nil,                                 // 14: service.ExtendersSettings.ExtendersEntry
}
var file_message_proto_depIdxs = []int32{
	6,  // 0: service.ProfessionConfig.drivers:type_name -> service.DriverConfig
	0,  // 1: service.ProfessionConfig.mod:type_name -> service.ProfessionConfig.ProfessionMod
	1,  // 2: service.ProfessionConfigs.data:type_name -> service.ProfessionConfig
	8,  // 3: service.ProfessionDetail.drivers:type_name -> service.DriverDetail
	12, // 4: service.DriverConfig.params:type_name -> service.DriverConfig.ParamsEntry
	13, // 5: service.DriverDetail.params:type_name -> service.DriverDetail.ParamsEntry
	7,  // 6: service.DriverDetail.plugin:type_name -> service.PluginInfo
	9,  // 7: service.WorkerConfigs.data:type_name -> service.WorkerConfig
	14, // 8: service.ExtendersSettings.Extenders:type_name -> service.ExtendersSettings.ExtendersEntry
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfessionConfig); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfessionConfigs); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfessionInfo); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfessionDetail); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverInfo); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverConfig); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginInfo); i {
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
		file_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverDetail); i {
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
		file_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerConfig); i {
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
		file_message_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerConfigs); i {
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
		file_message_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendersSettings); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
