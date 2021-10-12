// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: grape/api/v1/confd/types.proto

package confd

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "grape/api/v1/core"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *FileConfig) Reset() {
	*x = FileConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grape_api_v1_confd_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileConfig) ProtoMessage() {}

func (x *FileConfig) ProtoReflect() protoreflect.Message {
	mi := &file_grape_api_v1_confd_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileConfig.ProtoReflect.Descriptor instead.
func (*FileConfig) Descriptor() ([]byte, []int) {
	return file_grape_api_v1_confd_types_proto_rawDescGZIP(), []int{0}
}

func (x *FileConfig) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileConfig) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
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
		mi := &file_grape_api_v1_confd_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvConfig) ProtoMessage() {}

func (x *EnvConfig) ProtoReflect() protoreflect.Message {
	mi := &file_grape_api_v1_confd_types_proto_msgTypes[1]
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
	return file_grape_api_v1_confd_types_proto_rawDescGZIP(), []int{1}
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

type Configs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// string run = 2;
	FileConfigs []*FileConfig `protobuf:"bytes,3,rep,name=fileConfigs,proto3" json:"fileConfigs,omitempty"`
	EnvConfigs  []*EnvConfig  `protobuf:"bytes,4,rep,name=envConfigs,proto3" json:"envConfigs,omitempty"`
	Service     string        `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
	Group       string        `protobuf:"bytes,6,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *Configs) Reset() {
	*x = Configs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grape_api_v1_confd_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configs) ProtoMessage() {}

func (x *Configs) ProtoReflect() protoreflect.Message {
	mi := &file_grape_api_v1_confd_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configs.ProtoReflect.Descriptor instead.
func (*Configs) Descriptor() ([]byte, []int) {
	return file_grape_api_v1_confd_types_proto_rawDescGZIP(), []int{2}
}

func (x *Configs) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Configs) GetFileConfigs() []*FileConfig {
	if x != nil {
		return x.FileConfigs
	}
	return nil
}

func (x *Configs) GetEnvConfigs() []*EnvConfig {
	if x != nil {
		return x.EnvConfigs
	}
	return nil
}

func (x *Configs) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Configs) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

type ServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version      string              `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Service      string              `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Default      *Configs            `protobuf:"bytes,3,opt,name=default,proto3" json:"default,omitempty"`
	GroupConfigs map[string]*Configs `protobuf:"bytes,4,rep,name=groupConfigs,proto3" json:"groupConfigs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ServerConfig) Reset() {
	*x = ServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grape_api_v1_confd_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig) ProtoMessage() {}

func (x *ServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_grape_api_v1_confd_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return file_grape_api_v1_confd_types_proto_rawDescGZIP(), []int{3}
}

func (x *ServerConfig) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServerConfig) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ServerConfig) GetDefault() *Configs {
	if x != nil {
		return x.Default
	}
	return nil
}

func (x *ServerConfig) GetGroupConfigs() map[string]*Configs {
	if x != nil {
		return x.GroupConfigs
	}
	return nil
}

var File_grape_api_v1_confd_types_proto protoreflect.FileDescriptor

var file_grape_api_v1_confd_types_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x64, 0x1a, 0x1e, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x33, 0x0a, 0x09, 0x45, 0x6e, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xd4, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x66,
	0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x3d, 0x0a,
	0x0a, 0x65, 0x6e, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x64, 0x2e, 0x45, 0x6e, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x0a, 0x65, 0x6e, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0xaf, 0x02, 0x0a,
	0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x35, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52,
	0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x56, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x64, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x1a, 0x5c, 0x0a, 0x11, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x14,
	0x5a, 0x12, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grape_api_v1_confd_types_proto_rawDescOnce sync.Once
	file_grape_api_v1_confd_types_proto_rawDescData = file_grape_api_v1_confd_types_proto_rawDesc
)

func file_grape_api_v1_confd_types_proto_rawDescGZIP() []byte {
	file_grape_api_v1_confd_types_proto_rawDescOnce.Do(func() {
		file_grape_api_v1_confd_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_grape_api_v1_confd_types_proto_rawDescData)
	})
	return file_grape_api_v1_confd_types_proto_rawDescData
}

var file_grape_api_v1_confd_types_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grape_api_v1_confd_types_proto_goTypes = []interface{}{
	(*FileConfig)(nil),   // 0: grape.api.v1.confd.FileConfig
	(*EnvConfig)(nil),    // 1: grape.api.v1.confd.EnvConfig
	(*Configs)(nil),      // 2: grape.api.v1.confd.Configs
	(*ServerConfig)(nil), // 3: grape.api.v1.confd.ServerConfig
	nil,                  // 4: grape.api.v1.confd.ServerConfig.GroupConfigsEntry
}
var file_grape_api_v1_confd_types_proto_depIdxs = []int32{
	0, // 0: grape.api.v1.confd.Configs.fileConfigs:type_name -> grape.api.v1.confd.FileConfig
	1, // 1: grape.api.v1.confd.Configs.envConfigs:type_name -> grape.api.v1.confd.EnvConfig
	2, // 2: grape.api.v1.confd.ServerConfig.default:type_name -> grape.api.v1.confd.Configs
	4, // 3: grape.api.v1.confd.ServerConfig.groupConfigs:type_name -> grape.api.v1.confd.ServerConfig.GroupConfigsEntry
	2, // 4: grape.api.v1.confd.ServerConfig.GroupConfigsEntry.value:type_name -> grape.api.v1.confd.Configs
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_grape_api_v1_confd_types_proto_init() }
func file_grape_api_v1_confd_types_proto_init() {
	if File_grape_api_v1_confd_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grape_api_v1_confd_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileConfig); i {
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
		file_grape_api_v1_confd_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_grape_api_v1_confd_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configs); i {
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
		file_grape_api_v1_confd_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig); i {
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
			RawDescriptor: file_grape_api_v1_confd_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grape_api_v1_confd_types_proto_goTypes,
		DependencyIndexes: file_grape_api_v1_confd_types_proto_depIdxs,
		MessageInfos:      file_grape_api_v1_confd_types_proto_msgTypes,
	}.Build()
	File_grape_api_v1_confd_types_proto = out.File
	file_grape_api_v1_confd_types_proto_rawDesc = nil
	file_grape_api_v1_confd_types_proto_goTypes = nil
	file_grape_api_v1_confd_types_proto_depIdxs = nil
}