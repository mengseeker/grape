// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: api/v1/confd/controller.proto

package confd

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

type ApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string   `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	Project     *Project `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *ApiRequest) Reset() {
	*x = ApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_confd_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiRequest) ProtoMessage() {}

func (x *ApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_confd_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiRequest.ProtoReflect.Descriptor instead.
func (*ApiRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_confd_controller_proto_rawDescGZIP(), []int{0}
}

func (x *ApiRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *ApiRequest) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

type ApiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Project *Project `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *ApiResponse) Reset() {
	*x = ApiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_confd_controller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResponse) ProtoMessage() {}

func (x *ApiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_confd_controller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResponse.ProtoReflect.Descriptor instead.
func (*ApiResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_confd_controller_proto_rawDescGZIP(), []int{1}
}

func (x *ApiResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ApiResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApiResponse) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

var File_api_v1_confd_controller_proto protoreflect.FileDescriptor

var file_api_v1_confd_controller_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x64, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x64, 0x1a, 0x18, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x6c, 0x0a, 0x0b, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x32, 0xc5, 0x01, 0x0a, 0x09, 0x41, 0x70, 0x69, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x64, 0x2e, 0x41, 0x70, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x64, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x64, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x64, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3c, 0x0a, 0x03, 0x44, 0x65, 0x6c, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x64, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x64, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x14, 0x5a, 0x12, 0x67, 0x72, 0x61, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_confd_controller_proto_rawDescOnce sync.Once
	file_api_v1_confd_controller_proto_rawDescData = file_api_v1_confd_controller_proto_rawDesc
)

func file_api_v1_confd_controller_proto_rawDescGZIP() []byte {
	file_api_v1_confd_controller_proto_rawDescOnce.Do(func() {
		file_api_v1_confd_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_confd_controller_proto_rawDescData)
	})
	return file_api_v1_confd_controller_proto_rawDescData
}

var file_api_v1_confd_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_confd_controller_proto_goTypes = []interface{}{
	(*ApiRequest)(nil),  // 0: api.v1.confd.ApiRequest
	(*ApiResponse)(nil), // 1: api.v1.confd.ApiResponse
	(*Project)(nil),     // 2: api.v1.confd.Project
}
var file_api_v1_confd_controller_proto_depIdxs = []int32{
	2, // 0: api.v1.confd.ApiRequest.project:type_name -> api.v1.confd.Project
	2, // 1: api.v1.confd.ApiResponse.project:type_name -> api.v1.confd.Project
	0, // 2: api.v1.confd.ApiServer.Set:input_type -> api.v1.confd.ApiRequest
	0, // 3: api.v1.confd.ApiServer.Get:input_type -> api.v1.confd.ApiRequest
	0, // 4: api.v1.confd.ApiServer.Del:input_type -> api.v1.confd.ApiRequest
	1, // 5: api.v1.confd.ApiServer.Set:output_type -> api.v1.confd.ApiResponse
	1, // 6: api.v1.confd.ApiServer.Get:output_type -> api.v1.confd.ApiResponse
	1, // 7: api.v1.confd.ApiServer.Del:output_type -> api.v1.confd.ApiResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_confd_controller_proto_init() }
func file_api_v1_confd_controller_proto_init() {
	if File_api_v1_confd_controller_proto != nil {
		return
	}
	file_api_v1_confd_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_confd_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiRequest); i {
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
		file_api_v1_confd_controller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiResponse); i {
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
			RawDescriptor: file_api_v1_confd_controller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_confd_controller_proto_goTypes,
		DependencyIndexes: file_api_v1_confd_controller_proto_depIdxs,
		MessageInfos:      file_api_v1_confd_controller_proto_msgTypes,
	}.Build()
	File_api_v1_confd_controller_proto = out.File
	file_api_v1_confd_controller_proto_rawDesc = nil
	file_api_v1_confd_controller_proto_goTypes = nil
	file_api_v1_confd_controller_proto_depIdxs = nil
}
