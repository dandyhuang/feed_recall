// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.2
// source: conf/dict.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DictInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ClassName string `protobuf:"bytes,2,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	Path      string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DictInfo) Reset() {
	*x = DictInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_dict_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictInfo) ProtoMessage() {}

func (x *DictInfo) ProtoReflect() protoreflect.Message {
	mi := &file_conf_dict_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictInfo.ProtoReflect.Descriptor instead.
func (*DictInfo) Descriptor() ([]byte, []int) {
	return file_conf_dict_proto_rawDescGZIP(), []int{0}
}

func (x *DictInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DictInfo) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *DictInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Dict struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dicts []*DictInfo `protobuf:"bytes,1,rep,name=dicts,proto3" json:"dicts,omitempty"`
}

func (x *Dict) Reset() {
	*x = Dict{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_dict_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dict) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dict) ProtoMessage() {}

func (x *Dict) ProtoReflect() protoreflect.Message {
	mi := &file_conf_dict_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dict.ProtoReflect.Descriptor instead.
func (*Dict) Descriptor() ([]byte, []int) {
	return file_conf_dict_proto_rawDescGZIP(), []int{1}
}

func (x *Dict) GetDicts() []*DictInfo {
	if x != nil {
		return x.Dicts
	}
	return nil
}

var File_conf_dict_proto protoreflect.FileDescriptor

var file_conf_dict_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a,
	0x08, 0x44, 0x69, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x22, 0x32, 0x0a, 0x04, 0x44, 0x69, 0x63, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x64, 0x69, 0x63, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x64,
	0x69, 0x63, 0x74, 0x73, 0x42, 0x1f, 0x5a, 0x1d, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_dict_proto_rawDescOnce sync.Once
	file_conf_dict_proto_rawDescData = file_conf_dict_proto_rawDesc
)

func file_conf_dict_proto_rawDescGZIP() []byte {
	file_conf_dict_proto_rawDescOnce.Do(func() {
		file_conf_dict_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_dict_proto_rawDescData)
	})
	return file_conf_dict_proto_rawDescData
}

var file_conf_dict_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_conf_dict_proto_goTypes = []interface{}{
	(*DictInfo)(nil), // 0: kratos.api.DictInfo
	(*Dict)(nil),     // 1: kratos.api.Dict
}
var file_conf_dict_proto_depIdxs = []int32{
	0, // 0: kratos.api.Dict.dicts:type_name -> kratos.api.DictInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_conf_dict_proto_init() }
func file_conf_dict_proto_init() {
	if File_conf_dict_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_dict_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictInfo); i {
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
		file_conf_dict_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dict); i {
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
			RawDescriptor: file_conf_dict_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_dict_proto_goTypes,
		DependencyIndexes: file_conf_dict_proto_depIdxs,
		MessageInfos:      file_conf_dict_proto_msgTypes,
	}.Build()
	File_conf_dict_proto = out.File
	file_conf_dict_proto_rawDesc = nil
	file_conf_dict_proto_goTypes = nil
	file_conf_dict_proto_depIdxs = nil
}
