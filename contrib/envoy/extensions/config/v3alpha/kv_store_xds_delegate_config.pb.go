// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: contrib/envoy/extensions/config/v3alpha/kv_store_xds_delegate_config.proto

package v3alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/common/key_value/v3"
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

// [#extension: envoy.xds_delegates.kv_store]
//
// Configuration for a KeyValueStore-based XdsResourcesDelegate implementation. This implementation
// updates the underlying KV store with xDS resources received from the configured management
// servers, enabling configuration to be persisted locally and used on startup in case connectivity
// with the xDS management servers could not be established.
//
// The KV Store based delegate's handling of wildcard resources (empty resource list or "*") is
// designed for use with O(100) resources or fewer, so it's not currently advised to use this
// feature for large configurations with heavy use of wildcard resources.
type KeyValueStoreXdsDelegateConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for the KeyValueStore that holds the xDS resources.
	// [#allow-fully-qualified-name:]
	KeyValueStoreConfig *v3.KeyValueStoreConfig `protobuf:"bytes,1,opt,name=key_value_store_config,json=keyValueStoreConfig,proto3" json:"key_value_store_config,omitempty"`
}

func (x *KeyValueStoreXdsDelegateConfig) Reset() {
	*x = KeyValueStoreXdsDelegateConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValueStoreXdsDelegateConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValueStoreXdsDelegateConfig) ProtoMessage() {}

func (x *KeyValueStoreXdsDelegateConfig) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValueStoreXdsDelegateConfig.ProtoReflect.Descriptor instead.
func (*KeyValueStoreXdsDelegateConfig) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_rawDescGZIP(), []int{0}
}

func (x *KeyValueStoreXdsDelegateConfig) GetKeyValueStoreConfig() *v3.KeyValueStoreConfig {
	if x != nil {
		return x.KeyValueStoreConfig
	}
	return nil
}

var File_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto protoreflect.FileDescriptor

var file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6b, 0x76, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x78, 0x64, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x2d, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x1e,
	0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x58, 0x64, 0x73,
	0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x6a,
	0x0a, 0x16, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x13, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xa0, 0x01, 0x0a, 0x2d, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x1d, 0x4b, 0x76,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x58, 0x64, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x33,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_rawDescOnce sync.Once
	file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_rawDescData = file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_rawDesc
)

func file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_rawDescGZIP() []byte {
	file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_rawDescOnce.Do(func() {
		file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_rawDescData)
	})
	return file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_rawDescData
}

var file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_goTypes = []interface{}{
	(*KeyValueStoreXdsDelegateConfig)(nil), // 0: envoy.extensions.config.v3alpha.KeyValueStoreXdsDelegateConfig
	(*v3.KeyValueStoreConfig)(nil),         // 1: envoy.config.common.key_value.v3.KeyValueStoreConfig
}
var file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.config.v3alpha.KeyValueStoreXdsDelegateConfig.key_value_store_config:type_name -> envoy.config.common.key_value.v3.KeyValueStoreConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_init() }
func file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_init() {
	if File_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValueStoreXdsDelegateConfig); i {
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
			RawDescriptor: file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_goTypes,
		DependencyIndexes: file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_depIdxs,
		MessageInfos:      file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_msgTypes,
	}.Build()
	File_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto = out.File
	file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_rawDesc = nil
	file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_goTypes = nil
	file_contrib_envoy_extensions_config_v3alpha_kv_store_xds_delegate_config_proto_depIdxs = nil
}
