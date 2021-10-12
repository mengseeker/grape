// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: grape/api/v1/view/view.proto

package view

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	core "grape/api/v1/core"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path    string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	JsPaser string `protobuf:"bytes,3,opt,name=jsPaser,proto3" json:"jsPaser,omitempty"`
	Eventer string `protobuf:"bytes,4,opt,name=eventer,proto3" json:"eventer,omitempty"`
}

func (x *FileLog) Reset() {
	*x = FileLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grape_api_v1_view_view_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileLog) ProtoMessage() {}

func (x *FileLog) ProtoReflect() protoreflect.Message {
	mi := &file_grape_api_v1_view_view_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileLog.ProtoReflect.Descriptor instead.
func (*FileLog) Descriptor() ([]byte, []int) {
	return file_grape_api_v1_view_view_proto_rawDescGZIP(), []int{0}
}

func (x *FileLog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileLog) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileLog) GetJsPaser() string {
	if x != nil {
		return x.JsPaser
	}
	return ""
}

func (x *FileLog) GetEventer() string {
	if x != nil {
		return x.Eventer
	}
	return ""
}

type ServiceLogs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string     `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Service string     `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Logs    []*FileLog `protobuf:"bytes,3,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *ServiceLogs) Reset() {
	*x = ServiceLogs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grape_api_v1_view_view_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceLogs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceLogs) ProtoMessage() {}

func (x *ServiceLogs) ProtoReflect() protoreflect.Message {
	mi := &file_grape_api_v1_view_view_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceLogs.ProtoReflect.Descriptor instead.
func (*ServiceLogs) Descriptor() ([]byte, []int) {
	return file_grape_api_v1_view_view_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceLogs) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServiceLogs) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ServiceLogs) GetLogs() []*FileLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

type AllLogs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string         `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Logs    []*ServiceLogs `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *AllLogs) Reset() {
	*x = AllLogs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grape_api_v1_view_view_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllLogs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllLogs) ProtoMessage() {}

func (x *AllLogs) ProtoReflect() protoreflect.Message {
	mi := &file_grape_api_v1_view_view_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllLogs.ProtoReflect.Descriptor instead.
func (*AllLogs) Descriptor() ([]byte, []int) {
	return file_grape_api_v1_view_view_proto_rawDescGZIP(), []int{2}
}

func (x *AllLogs) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AllLogs) GetLogs() []*ServiceLogs {
	if x != nil {
		return x.Logs
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grape_api_v1_view_view_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grape_api_v1_view_view_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_grape_api_v1_view_view_proto_rawDescGZIP(), []int{3}
}

func (x *GetRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type DelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *DelRequest) Reset() {
	*x = DelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grape_api_v1_view_view_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelRequest) ProtoMessage() {}

func (x *DelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grape_api_v1_view_view_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelRequest.ProtoReflect.Descriptor instead.
func (*DelRequest) Descriptor() ([]byte, []int) {
	return file_grape_api_v1_view_view_proto_rawDescGZIP(), []int{4}
}

func (x *DelRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type DiscoveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DiscoveryRequest) Reset() {
	*x = DiscoveryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grape_api_v1_view_view_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryRequest) ProtoMessage() {}

func (x *DiscoveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grape_api_v1_view_view_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveryRequest.ProtoReflect.Descriptor instead.
func (*DiscoveryRequest) Descriptor() ([]byte, []int) {
	return file_grape_api_v1_view_view_proto_rawDescGZIP(), []int{5}
}

var File_grape_api_v1_view_view_proto protoreflect.FileDescriptor

var file_grape_api_v1_view_view_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x76,
	0x69, 0x65, 0x77, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x67, 0x72, 0x61, 0x70, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x69, 0x65,
	0x77, 0x1a, 0x1e, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x65, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x73, 0x50, 0x61, 0x73, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x73, 0x50, 0x61, 0x73, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x71, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x6c,
	0x6f, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x61, 0x70,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x57, 0x0a, 0x07, 0x41,
	0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x32, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x69,
	0x65, 0x77, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x04,
	0x6c, 0x6f, 0x67, 0x73, 0x22, 0x26, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x26, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xd8, 0x01, 0x0a, 0x09, 0x41, 0x70, 0x69,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x1e, 0x2e,
	0x67, 0x72, 0x61, 0x70, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x69, 0x65,
	0x77, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x1a, 0x18, 0x2e,
	0x67, 0x72, 0x61, 0x70, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x1d, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x76, 0x69, 0x65, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x22,
	0x00, 0x12, 0x40, 0x0a, 0x03, 0x44, 0x65, 0x6c, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x44, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x32, 0x64, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x4c, 0x6f, 0x67, 0x73, 0x12, 0x23, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x61, 0x70,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x41, 0x6c,
	0x6c, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0x00, 0x30, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x67, 0x72, 0x61,
	0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grape_api_v1_view_view_proto_rawDescOnce sync.Once
	file_grape_api_v1_view_view_proto_rawDescData = file_grape_api_v1_view_view_proto_rawDesc
)

func file_grape_api_v1_view_view_proto_rawDescGZIP() []byte {
	file_grape_api_v1_view_view_proto_rawDescOnce.Do(func() {
		file_grape_api_v1_view_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_grape_api_v1_view_view_proto_rawDescData)
	})
	return file_grape_api_v1_view_view_proto_rawDescData
}

var file_grape_api_v1_view_view_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grape_api_v1_view_view_proto_goTypes = []interface{}{
	(*FileLog)(nil),          // 0: grape.api.v1.view.FileLog
	(*ServiceLogs)(nil),      // 1: grape.api.v1.view.ServiceLogs
	(*AllLogs)(nil),          // 2: grape.api.v1.view.AllLogs
	(*GetRequest)(nil),       // 3: grape.api.v1.view.GetRequest
	(*DelRequest)(nil),       // 4: grape.api.v1.view.DelRequest
	(*DiscoveryRequest)(nil), // 5: grape.api.v1.view.DiscoveryRequest
	(*core.Empty)(nil),       // 6: grape.api.v1.core.Empty
}
var file_grape_api_v1_view_view_proto_depIdxs = []int32{
	0, // 0: grape.api.v1.view.ServiceLogs.logs:type_name -> grape.api.v1.view.FileLog
	1, // 1: grape.api.v1.view.AllLogs.logs:type_name -> grape.api.v1.view.ServiceLogs
	1, // 2: grape.api.v1.view.ApiServer.Set:input_type -> grape.api.v1.view.ServiceLogs
	3, // 3: grape.api.v1.view.ApiServer.Get:input_type -> grape.api.v1.view.GetRequest
	4, // 4: grape.api.v1.view.ApiServer.Del:input_type -> grape.api.v1.view.DelRequest
	5, // 5: grape.api.v1.view.DiscoveryServer.StreamLogs:input_type -> grape.api.v1.view.DiscoveryRequest
	6, // 6: grape.api.v1.view.ApiServer.Set:output_type -> grape.api.v1.core.Empty
	1, // 7: grape.api.v1.view.ApiServer.Get:output_type -> grape.api.v1.view.ServiceLogs
	6, // 8: grape.api.v1.view.ApiServer.Del:output_type -> grape.api.v1.core.Empty
	2, // 9: grape.api.v1.view.DiscoveryServer.StreamLogs:output_type -> grape.api.v1.view.AllLogs
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grape_api_v1_view_view_proto_init() }
func file_grape_api_v1_view_view_proto_init() {
	if File_grape_api_v1_view_view_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grape_api_v1_view_view_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileLog); i {
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
		file_grape_api_v1_view_view_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceLogs); i {
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
		file_grape_api_v1_view_view_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllLogs); i {
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
		file_grape_api_v1_view_view_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_grape_api_v1_view_view_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelRequest); i {
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
		file_grape_api_v1_view_view_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoveryRequest); i {
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
			RawDescriptor: file_grape_api_v1_view_view_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_grape_api_v1_view_view_proto_goTypes,
		DependencyIndexes: file_grape_api_v1_view_view_proto_depIdxs,
		MessageInfos:      file_grape_api_v1_view_view_proto_msgTypes,
	}.Build()
	File_grape_api_v1_view_view_proto = out.File
	file_grape_api_v1_view_view_proto_rawDesc = nil
	file_grape_api_v1_view_view_proto_goTypes = nil
	file_grape_api_v1_view_view_proto_depIdxs = nil
}
