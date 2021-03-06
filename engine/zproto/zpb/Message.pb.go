// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: Message.proto

package zpb

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

type COMPONENT_TYPE int32

const (
	COMPONENT_TYPE_CENTER     COMPONENT_TYPE = 0
	COMPONENT_TYPE_DISPATCHER COMPONENT_TYPE = 1
	COMPONENT_TYPE_LOGIN      COMPONENT_TYPE = 2
	COMPONENT_TYPE_GATE       COMPONENT_TYPE = 3
	COMPONENT_TYPE_GAME       COMPONENT_TYPE = 4
)

// Enum value maps for COMPONENT_TYPE.
var (
	COMPONENT_TYPE_name = map[int32]string{
		0: "CENTER",
		1: "DISPATCHER",
		2: "LOGIN",
		3: "GATE",
		4: "GAME",
	}
	COMPONENT_TYPE_value = map[string]int32{
		"CENTER":     0,
		"DISPATCHER": 1,
		"LOGIN":      2,
		"GATE":       3,
		"GAME":       4,
	}
)

func (x COMPONENT_TYPE) Enum() *COMPONENT_TYPE {
	p := new(COMPONENT_TYPE)
	*p = x
	return p
}

func (x COMPONENT_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (COMPONENT_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_Message_proto_enumTypes[0].Descriptor()
}

func (COMPONENT_TYPE) Type() protoreflect.EnumType {
	return &file_Message_proto_enumTypes[0]
}

func (x COMPONENT_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use COMPONENT_TYPE.Descriptor instead.
func (COMPONENT_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{0}
}

type Property_Type int32

const (
	Property_Type_INVALID Property_Type = 0
	Property_Type_INT32   Property_Type = 1
	Property_Type_String  Property_Type = 2
	Property_Type_INT64   Property_Type = 3
	Property_Type_BOOL    Property_Type = 4
)

// Enum value maps for Property_Type.
var (
	Property_Type_name = map[int32]string{
		0: "INVALID",
		1: "INT32",
		2: "String",
		3: "INT64",
		4: "BOOL",
	}
	Property_Type_value = map[string]int32{
		"INVALID": 0,
		"INT32":   1,
		"String":  2,
		"INT64":   3,
		"BOOL":    4,
	}
)

func (x Property_Type) Enum() *Property_Type {
	p := new(Property_Type)
	*p = x
	return p
}

func (x Property_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Property_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_Message_proto_enumTypes[1].Descriptor()
}

func (Property_Type) Type() protoreflect.EnumType {
	return &file_Message_proto_enumTypes[1]
}

func (x Property_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Property_Type.Descriptor instead.
func (Property_Type) EnumDescriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{1}
}

type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    Property_Type `protobuf:"varint,1,opt,name=type,proto3,enum=zpb.Property_Type" json:"type,omitempty"`
	PInt    int32         `protobuf:"varint,2,opt,name=p_int,json=pInt,proto3" json:"p_int,omitempty"`
	PString string        `protobuf:"bytes,3,opt,name=p_string,json=pString,proto3" json:"p_string,omitempty"`
	PFloat  uint64        `protobuf:"varint,4,opt,name=p_float,json=pFloat,proto3" json:"p_float,omitempty"`
	PBool   bool          `protobuf:"varint,5,opt,name=p_bool,json=pBool,proto3" json:"p_bool,omitempty"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{0}
}

func (x *Property) GetType() Property_Type {
	if x != nil {
		return x.Type
	}
	return Property_Type_INVALID
}

func (x *Property) GetPInt() int32 {
	if x != nil {
		return x.PInt
	}
	return 0
}

func (x *Property) GetPString() string {
	if x != nil {
		return x.PString
	}
	return ""
}

func (x *Property) GetPFloat() uint64 {
	if x != nil {
		return x.PFloat
	}
	return 0
}

func (x *Property) GetPBool() bool {
	if x != nil {
		return x.PBool
	}
	return false
}

type PropertyMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   int32     `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *Property `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PropertyMap) Reset() {
	*x = PropertyMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertyMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertyMap) ProtoMessage() {}

func (x *PropertyMap) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertyMap.ProtoReflect.Descriptor instead.
func (*PropertyMap) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{1}
}

func (x *PropertyMap) GetKey() int32 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *PropertyMap) GetValue() *Property {
	if x != nil {
		return x.Value
	}
	return nil
}

type ADD_ENGINE_COMPONENT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        COMPONENT_TYPE `protobuf:"varint,1,opt,name=type,proto3,enum=zpb.COMPONENT_TYPE" json:"type,omitempty"`
	ListenAddr  string         `protobuf:"bytes,2,opt,name=listen_addr,json=listenAddr,proto3" json:"listen_addr,omitempty"`
	ComponentId int32          `protobuf:"varint,3,opt,name=component_id,json=componentId,proto3" json:"component_id,omitempty"`
}

func (x *ADD_ENGINE_COMPONENT) Reset() {
	*x = ADD_ENGINE_COMPONENT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADD_ENGINE_COMPONENT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADD_ENGINE_COMPONENT) ProtoMessage() {}

func (x *ADD_ENGINE_COMPONENT) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADD_ENGINE_COMPONENT.ProtoReflect.Descriptor instead.
func (*ADD_ENGINE_COMPONENT) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{2}
}

func (x *ADD_ENGINE_COMPONENT) GetType() COMPONENT_TYPE {
	if x != nil {
		return x.Type
	}
	return COMPONENT_TYPE_CENTER
}

func (x *ADD_ENGINE_COMPONENT) GetListenAddr() string {
	if x != nil {
		return x.ListenAddr
	}
	return ""
}

func (x *ADD_ENGINE_COMPONENT) GetComponentId() int32 {
	if x != nil {
		return x.ComponentId
	}
	return 0
}

type ADD_ENGINE_COMPONENT_ACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComponentId   int32                   `protobuf:"varint,1,opt,name=component_id,json=componentId,proto3" json:"component_id,omitempty"`
	ComponentList []*ADD_ENGINE_COMPONENT `protobuf:"bytes,2,rep,name=component_list,json=componentList,proto3" json:"component_list,omitempty"`
}

func (x *ADD_ENGINE_COMPONENT_ACK) Reset() {
	*x = ADD_ENGINE_COMPONENT_ACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADD_ENGINE_COMPONENT_ACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADD_ENGINE_COMPONENT_ACK) ProtoMessage() {}

func (x *ADD_ENGINE_COMPONENT_ACK) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADD_ENGINE_COMPONENT_ACK.ProtoReflect.Descriptor instead.
func (*ADD_ENGINE_COMPONENT_ACK) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{3}
}

func (x *ADD_ENGINE_COMPONENT_ACK) GetComponentId() int32 {
	if x != nil {
		return x.ComponentId
	}
	return 0
}

func (x *ADD_ENGINE_COMPONENT_ACK) GetComponentList() []*ADD_ENGINE_COMPONENT {
	if x != nil {
		return x.ComponentList
	}
	return nil
}

type GET_ALL_LOGIN_COMPONENT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GET_ALL_LOGIN_COMPONENT) Reset() {
	*x = GET_ALL_LOGIN_COMPONENT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GET_ALL_LOGIN_COMPONENT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GET_ALL_LOGIN_COMPONENT) ProtoMessage() {}

func (x *GET_ALL_LOGIN_COMPONENT) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GET_ALL_LOGIN_COMPONENT.ProtoReflect.Descriptor instead.
func (*GET_ALL_LOGIN_COMPONENT) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{4}
}

type GET_ALL_LOGIN_COMPONENT_ACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginList []string `protobuf:"bytes,1,rep,name=login_list,json=loginList,proto3" json:"login_list,omitempty"`
}

func (x *GET_ALL_LOGIN_COMPONENT_ACK) Reset() {
	*x = GET_ALL_LOGIN_COMPONENT_ACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GET_ALL_LOGIN_COMPONENT_ACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GET_ALL_LOGIN_COMPONENT_ACK) ProtoMessage() {}

func (x *GET_ALL_LOGIN_COMPONENT_ACK) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GET_ALL_LOGIN_COMPONENT_ACK.ProtoReflect.Descriptor instead.
func (*GET_ALL_LOGIN_COMPONENT_ACK) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{5}
}

func (x *GET_ALL_LOGIN_COMPONENT_ACK) GetLoginList() []string {
	if x != nil {
		return x.LoginList
	}
	return nil
}

type SYNC_PROXY_PROPERTY struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyMapList []*PropertyMap `protobuf:"bytes,1,rep,name=property_map_list,json=propertyMapList,proto3" json:"property_map_list,omitempty"`
}

func (x *SYNC_PROXY_PROPERTY) Reset() {
	*x = SYNC_PROXY_PROPERTY{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SYNC_PROXY_PROPERTY) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SYNC_PROXY_PROPERTY) ProtoMessage() {}

func (x *SYNC_PROXY_PROPERTY) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SYNC_PROXY_PROPERTY.ProtoReflect.Descriptor instead.
func (*SYNC_PROXY_PROPERTY) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{6}
}

func (x *SYNC_PROXY_PROPERTY) GetPropertyMapList() []*PropertyMap {
	if x != nil {
		return x.PropertyMapList
	}
	return nil
}

type SET_REMOTE_PROPERTY struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyMapList []*PropertyMap `protobuf:"bytes,1,rep,name=property_map_list,json=propertyMapList,proto3" json:"property_map_list,omitempty"`
}

func (x *SET_REMOTE_PROPERTY) Reset() {
	*x = SET_REMOTE_PROPERTY{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SET_REMOTE_PROPERTY) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SET_REMOTE_PROPERTY) ProtoMessage() {}

func (x *SET_REMOTE_PROPERTY) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SET_REMOTE_PROPERTY.ProtoReflect.Descriptor instead.
func (*SET_REMOTE_PROPERTY) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{7}
}

func (x *SET_REMOTE_PROPERTY) GetPropertyMapList() []*PropertyMap {
	if x != nil {
		return x.PropertyMapList
	}
	return nil
}

var File_Message_proto protoreflect.FileDescriptor

var file_Message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x7a, 0x70, 0x62, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x7a, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x5f, 0x69,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x49, 0x6e, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x5f, 0x66,
	0x6c, 0x6f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x70, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x44, 0x0a, 0x0b, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x4d, 0x61, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x7a, 0x70, 0x62, 0x2e,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x83, 0x01, 0x0a, 0x14, 0x41, 0x44, 0x44, 0x5f, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x7a, 0x70, 0x62, 0x2e, 0x43, 0x4f, 0x4d,
	0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64,
	0x64, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x7f, 0x0a, 0x18, 0x41, 0x44, 0x44, 0x5f, 0x45, 0x4e, 0x47,
	0x49, 0x4e, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43,
	0x4b, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x7a,
	0x70, 0x62, 0x2e, 0x41, 0x44, 0x44, 0x5f, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x19, 0x0a, 0x17, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x4c,
	0x4c, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x4e, 0x45, 0x4e,
	0x54, 0x22, 0x3c, 0x0a, 0x1b, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x4c, 0x4f, 0x47,
	0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x4b,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x53, 0x0a, 0x13, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x5f, 0x50, 0x52,
	0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x12, 0x3c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x7a, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x4d, 0x61, 0x70, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4d, 0x61, 0x70,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x13, 0x53, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x4d, 0x4f,
	0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x12, 0x3c, 0x0a, 0x11, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x7a, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x4d, 0x61, 0x70, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x4d, 0x61, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x2a, 0x4b, 0x0a, 0x0e, 0x43, 0x4f, 0x4d,
	0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0a, 0x0a, 0x06, 0x43,
	0x45, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x53, 0x50, 0x41,
	0x54, 0x43, 0x48, 0x45, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x47, 0x49, 0x4e,
	0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04,
	0x47, 0x41, 0x4d, 0x45, 0x10, 0x04, 0x2a, 0x48, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x49,
	0x4e, 0x54, 0x36, 0x34, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x04,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x7a, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Message_proto_rawDescOnce sync.Once
	file_Message_proto_rawDescData = file_Message_proto_rawDesc
)

func file_Message_proto_rawDescGZIP() []byte {
	file_Message_proto_rawDescOnce.Do(func() {
		file_Message_proto_rawDescData = protoimpl.X.CompressGZIP(file_Message_proto_rawDescData)
	})
	return file_Message_proto_rawDescData
}

var file_Message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_Message_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_Message_proto_goTypes = []interface{}{
	(COMPONENT_TYPE)(0),                 // 0: zpb.COMPONENT_TYPE
	(Property_Type)(0),                  // 1: zpb.Property_Type
	(*Property)(nil),                    // 2: zpb.Property
	(*PropertyMap)(nil),                 // 3: zpb.PropertyMap
	(*ADD_ENGINE_COMPONENT)(nil),        // 4: zpb.ADD_ENGINE_COMPONENT
	(*ADD_ENGINE_COMPONENT_ACK)(nil),    // 5: zpb.ADD_ENGINE_COMPONENT_ACK
	(*GET_ALL_LOGIN_COMPONENT)(nil),     // 6: zpb.GET_ALL_LOGIN_COMPONENT
	(*GET_ALL_LOGIN_COMPONENT_ACK)(nil), // 7: zpb.GET_ALL_LOGIN_COMPONENT_ACK
	(*SYNC_PROXY_PROPERTY)(nil),         // 8: zpb.SYNC_PROXY_PROPERTY
	(*SET_REMOTE_PROPERTY)(nil),         // 9: zpb.SET_REMOTE_PROPERTY
}
var file_Message_proto_depIdxs = []int32{
	1, // 0: zpb.Property.type:type_name -> zpb.Property_Type
	2, // 1: zpb.PropertyMap.value:type_name -> zpb.Property
	0, // 2: zpb.ADD_ENGINE_COMPONENT.type:type_name -> zpb.COMPONENT_TYPE
	4, // 3: zpb.ADD_ENGINE_COMPONENT_ACK.component_list:type_name -> zpb.ADD_ENGINE_COMPONENT
	3, // 4: zpb.SYNC_PROXY_PROPERTY.property_map_list:type_name -> zpb.PropertyMap
	3, // 5: zpb.SET_REMOTE_PROPERTY.property_map_list:type_name -> zpb.PropertyMap
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_Message_proto_init() }
func file_Message_proto_init() {
	if File_Message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
		file_Message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyMap); i {
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
		file_Message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADD_ENGINE_COMPONENT); i {
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
		file_Message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADD_ENGINE_COMPONENT_ACK); i {
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
		file_Message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GET_ALL_LOGIN_COMPONENT); i {
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
		file_Message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GET_ALL_LOGIN_COMPONENT_ACK); i {
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
		file_Message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SYNC_PROXY_PROPERTY); i {
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
		file_Message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SET_REMOTE_PROPERTY); i {
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
			RawDescriptor: file_Message_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Message_proto_goTypes,
		DependencyIndexes: file_Message_proto_depIdxs,
		EnumInfos:         file_Message_proto_enumTypes,
		MessageInfos:      file_Message_proto_msgTypes,
	}.Build()
	File_Message_proto = out.File
	file_Message_proto_rawDesc = nil
	file_Message_proto_goTypes = nil
	file_Message_proto_depIdxs = nil
}
