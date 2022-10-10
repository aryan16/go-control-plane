// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: envoy/api/v2/discovery.proto

package apiv2

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	any "github.com/golang/protobuf/ptypes/any"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// A DiscoveryRequest requests a set of versioned resources of the same type for
// a given Envoy node on some API.
// [#next-free-field: 7]
type DiscoveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The version_info provided in the request messages will be the version_info
	// received with the most recent successfully processed response or empty on
	// the first request. It is expected that no new request is sent after a
	// response is received until the Envoy instance is ready to ACK/NACK the new
	// configuration. ACK/NACK takes place by returning the new API config version
	// as applied or the previous API config version respectively. Each type_url
	// (see below) has an independent version associated with it.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The node making the request.
	Node *core.Node `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	// List of resources to subscribe to, e.g. list of cluster names or a route
	// configuration name. If this is empty, all resources for the API are
	// returned. LDS/CDS may have empty resource_names, which will cause all
	// resources for the Envoy instance to be returned. The LDS and CDS responses
	// will then imply a number of resources that need to be fetched via EDS/RDS,
	// which will be explicitly enumerated in resource_names.
	ResourceNames []string `protobuf:"bytes,3,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty"`
	// Type of the resource that is being requested, e.g.
	// "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment". This is implicit
	// in requests made via singleton xDS APIs such as CDS, LDS, etc. but is
	// required for ADS.
	TypeUrl string `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// nonce corresponding to DiscoveryResponse being ACK/NACKed. See above
	// discussion on version_info and the DiscoveryResponse nonce comment. This
	// may be empty only if 1) this is a non-persistent-stream xDS such as HTTP,
	// or 2) the client has not yet accepted an update in this xDS stream (unlike
	// delta, where it is populated only for new explicit ACKs).
	ResponseNonce string `protobuf:"bytes,5,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	// This is populated when the previous :ref:`DiscoveryResponse <envoy_api_msg_DiscoveryResponse>`
	// failed to update configuration. The *message* field in *error_details* provides the Envoy
	// internal exception related to the failure. It is only intended for consumption during manual
	// debugging, the string provided is not guaranteed to be stable across Envoy versions.
	ErrorDetail *status.Status `protobuf:"bytes,6,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
}

func (x *DiscoveryRequest) Reset() {
	*x = DiscoveryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_discovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryRequest) ProtoMessage() {}

func (x *DiscoveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_discovery_proto_msgTypes[0]
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
	return file_envoy_api_v2_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *DiscoveryRequest) GetVersionInfo() string {
	if x != nil {
		return x.VersionInfo
	}
	return ""
}

func (x *DiscoveryRequest) GetNode() *core.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *DiscoveryRequest) GetResourceNames() []string {
	if x != nil {
		return x.ResourceNames
	}
	return nil
}

func (x *DiscoveryRequest) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *DiscoveryRequest) GetResponseNonce() string {
	if x != nil {
		return x.ResponseNonce
	}
	return ""
}

func (x *DiscoveryRequest) GetErrorDetail() *status.Status {
	if x != nil {
		return x.ErrorDetail
	}
	return nil
}

// [#next-free-field: 7]
type DiscoveryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The version of the response data.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The response resources. These resources are typed and depend on the API being called.
	Resources []*any.Any `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	// [#not-implemented-hide:]
	// Canary is used to support two Envoy command line flags:
	//
	// * --terminate-on-canary-transition-failure. When set, Envoy is able to
	//   terminate if it detects that configuration is stuck at canary. Consider
	//   this example sequence of updates:
	//   - Management server applies a canary config successfully.
	//   - Management server rolls back to a production config.
	//   - Envoy rejects the new production config.
	//   Since there is no sensible way to continue receiving configuration
	//   updates, Envoy will then terminate and apply production config from a
	//   clean slate.
	// * --dry-run-canary. When set, a canary response will never be applied, only
	//   validated via a dry run.
	Canary bool `protobuf:"varint,3,opt,name=canary,proto3" json:"canary,omitempty"`
	// Type URL for resources. Identifies the xDS API when muxing over ADS.
	// Must be consistent with the type_url in the 'resources' repeated Any (if non-empty).
	TypeUrl string `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// For gRPC based subscriptions, the nonce provides a way to explicitly ack a
	// specific DiscoveryResponse in a following DiscoveryRequest. Additional
	// messages may have been sent by Envoy to the management server for the
	// previous version on the stream prior to this DiscoveryResponse, that were
	// unprocessed at response send time. The nonce allows the management server
	// to ignore any further DiscoveryRequests for the previous version until a
	// DiscoveryRequest bearing the nonce. The nonce is optional and is not
	// required for non-stream based xDS implementations.
	Nonce string `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// [#not-implemented-hide:]
	// The control plane instance that sent the response.
	ControlPlane *core.ControlPlane `protobuf:"bytes,6,opt,name=control_plane,json=controlPlane,proto3" json:"control_plane,omitempty"`
}

func (x *DiscoveryResponse) Reset() {
	*x = DiscoveryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_discovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoveryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryResponse) ProtoMessage() {}

func (x *DiscoveryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_discovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveryResponse.ProtoReflect.Descriptor instead.
func (*DiscoveryResponse) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *DiscoveryResponse) GetVersionInfo() string {
	if x != nil {
		return x.VersionInfo
	}
	return ""
}

func (x *DiscoveryResponse) GetResources() []*any.Any {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *DiscoveryResponse) GetCanary() bool {
	if x != nil {
		return x.Canary
	}
	return false
}

func (x *DiscoveryResponse) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *DiscoveryResponse) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *DiscoveryResponse) GetControlPlane() *core.ControlPlane {
	if x != nil {
		return x.ControlPlane
	}
	return nil
}

// DeltaDiscoveryRequest and DeltaDiscoveryResponse are used in a new gRPC
// endpoint for Delta xDS.
//
// With Delta xDS, the DeltaDiscoveryResponses do not need to include a full
// snapshot of the tracked resources. Instead, DeltaDiscoveryResponses are a
// diff to the state of a xDS client.
// In Delta XDS there are per-resource versions, which allow tracking state at
// the resource granularity.
// An xDS Delta session is always in the context of a gRPC bidirectional
// stream. This allows the xDS server to keep track of the state of xDS clients
// connected to it.
//
// In Delta xDS the nonce field is required and used to pair
// DeltaDiscoveryResponse to a DeltaDiscoveryRequest ACK or NACK.
// Optionally, a response message level system_version_info is present for
// debugging purposes only.
//
// DeltaDiscoveryRequest plays two independent roles. Any DeltaDiscoveryRequest
// can be either or both of: [1] informing the server of what resources the
// client has gained/lost interest in (using resource_names_subscribe and
// resource_names_unsubscribe), or [2] (N)ACKing an earlier resource update from
// the server (using response_nonce, with presence of error_detail making it a NACK).
// Additionally, the first message (for a given type_url) of a reconnected gRPC stream
// has a third role: informing the server of the resources (and their versions)
// that the client already possesses, using the initial_resource_versions field.
//
// As with state-of-the-world, when multiple resource types are multiplexed (ADS),
// all requests/acknowledgments/updates are logically walled off by type_url:
// a Cluster ACK exists in a completely separate world from a prior Route NACK.
// In particular, initial_resource_versions being sent at the "start" of every
// gRPC stream actually entails a message for each type_url, each with its own
// initial_resource_versions.
// [#next-free-field: 8]
type DeltaDiscoveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The node making the request.
	Node *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// Type of the resource that is being requested, e.g.
	// "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment".
	TypeUrl string `protobuf:"bytes,2,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// DeltaDiscoveryRequests allow the client to add or remove individual
	// resources to the set of tracked resources in the context of a stream.
	// All resource names in the resource_names_subscribe list are added to the
	// set of tracked resources and all resource names in the resource_names_unsubscribe
	// list are removed from the set of tracked resources.
	//
	// *Unlike* state-of-the-world xDS, an empty resource_names_subscribe or
	// resource_names_unsubscribe list simply means that no resources are to be
	// added or removed to the resource list.
	// *Like* state-of-the-world xDS, the server must send updates for all tracked
	// resources, but can also send updates for resources the client has not subscribed to.
	//
	// NOTE: the server must respond with all resources listed in resource_names_subscribe,
	// even if it believes the client has the most recent version of them. The reason:
	// the client may have dropped them, but then regained interest before it had a chance
	// to send the unsubscribe message. See DeltaSubscriptionStateTest.RemoveThenAdd.
	//
	// These two fields can be set in any DeltaDiscoveryRequest, including ACKs
	// and initial_resource_versions.
	//
	// A list of Resource names to add to the list of tracked resources.
	ResourceNamesSubscribe []string `protobuf:"bytes,3,rep,name=resource_names_subscribe,json=resourceNamesSubscribe,proto3" json:"resource_names_subscribe,omitempty"`
	// A list of Resource names to remove from the list of tracked resources.
	ResourceNamesUnsubscribe []string `protobuf:"bytes,4,rep,name=resource_names_unsubscribe,json=resourceNamesUnsubscribe,proto3" json:"resource_names_unsubscribe,omitempty"`
	// Informs the server of the versions of the resources the xDS client knows of, to enable the
	// client to continue the same logical xDS session even in the face of gRPC stream reconnection.
	// It will not be populated: [1] in the very first stream of a session, since the client will
	// not yet have any resources,  [2] in any message after the first in a stream (for a given
	// type_url), since the server will already be correctly tracking the client's state.
	// (In ADS, the first message *of each type_url* of a reconnected stream populates this map.)
	// The map's keys are names of xDS resources known to the xDS client.
	// The map's values are opaque resource versions.
	InitialResourceVersions map[string]string `protobuf:"bytes,5,rep,name=initial_resource_versions,json=initialResourceVersions,proto3" json:"initial_resource_versions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// When the DeltaDiscoveryRequest is a ACK or NACK message in response
	// to a previous DeltaDiscoveryResponse, the response_nonce must be the
	// nonce in the DeltaDiscoveryResponse.
	// Otherwise (unlike in DiscoveryRequest) response_nonce must be omitted.
	ResponseNonce string `protobuf:"bytes,6,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	// This is populated when the previous :ref:`DiscoveryResponse <envoy_api_msg_DiscoveryResponse>`
	// failed to update configuration. The *message* field in *error_details*
	// provides the Envoy internal exception related to the failure.
	ErrorDetail *status.Status `protobuf:"bytes,7,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
}

func (x *DeltaDiscoveryRequest) Reset() {
	*x = DeltaDiscoveryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_discovery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaDiscoveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaDiscoveryRequest) ProtoMessage() {}

func (x *DeltaDiscoveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_discovery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaDiscoveryRequest.ProtoReflect.Descriptor instead.
func (*DeltaDiscoveryRequest) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_discovery_proto_rawDescGZIP(), []int{2}
}

func (x *DeltaDiscoveryRequest) GetNode() *core.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *DeltaDiscoveryRequest) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *DeltaDiscoveryRequest) GetResourceNamesSubscribe() []string {
	if x != nil {
		return x.ResourceNamesSubscribe
	}
	return nil
}

func (x *DeltaDiscoveryRequest) GetResourceNamesUnsubscribe() []string {
	if x != nil {
		return x.ResourceNamesUnsubscribe
	}
	return nil
}

func (x *DeltaDiscoveryRequest) GetInitialResourceVersions() map[string]string {
	if x != nil {
		return x.InitialResourceVersions
	}
	return nil
}

func (x *DeltaDiscoveryRequest) GetResponseNonce() string {
	if x != nil {
		return x.ResponseNonce
	}
	return ""
}

func (x *DeltaDiscoveryRequest) GetErrorDetail() *status.Status {
	if x != nil {
		return x.ErrorDetail
	}
	return nil
}

// [#next-free-field: 7]
type DeltaDiscoveryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The version of the response data (used for debugging).
	SystemVersionInfo string `protobuf:"bytes,1,opt,name=system_version_info,json=systemVersionInfo,proto3" json:"system_version_info,omitempty"`
	// The response resources. These are typed resources, whose types must match
	// the type_url field.
	Resources []*Resource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	// Type URL for resources. Identifies the xDS API when muxing over ADS.
	// Must be consistent with the type_url in the Any within 'resources' if 'resources' is non-empty.
	TypeUrl string `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// Resources names of resources that have be deleted and to be removed from the xDS Client.
	// Removed resources for missing resources can be ignored.
	RemovedResources []string `protobuf:"bytes,6,rep,name=removed_resources,json=removedResources,proto3" json:"removed_resources,omitempty"`
	// The nonce provides a way for DeltaDiscoveryRequests to uniquely
	// reference a DeltaDiscoveryResponse when (N)ACKing. The nonce is required.
	Nonce string `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *DeltaDiscoveryResponse) Reset() {
	*x = DeltaDiscoveryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_discovery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaDiscoveryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaDiscoveryResponse) ProtoMessage() {}

func (x *DeltaDiscoveryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_discovery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaDiscoveryResponse.ProtoReflect.Descriptor instead.
func (*DeltaDiscoveryResponse) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_discovery_proto_rawDescGZIP(), []int{3}
}

func (x *DeltaDiscoveryResponse) GetSystemVersionInfo() string {
	if x != nil {
		return x.SystemVersionInfo
	}
	return ""
}

func (x *DeltaDiscoveryResponse) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *DeltaDiscoveryResponse) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *DeltaDiscoveryResponse) GetRemovedResources() []string {
	if x != nil {
		return x.RemovedResources
	}
	return nil
}

func (x *DeltaDiscoveryResponse) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource's name, to distinguish it from others of the same type of resource.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The aliases are a list of other names that this resource can go by.
	Aliases []string `protobuf:"bytes,4,rep,name=aliases,proto3" json:"aliases,omitempty"`
	// The resource level version. It allows xDS to track the state of individual
	// resources.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The resource being tracked.
	Resource *any.Any `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_discovery_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_discovery_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_discovery_proto_rawDescGZIP(), []int{4}
}

func (x *Resource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Resource) GetAliases() []string {
	if x != nil {
		return x.Aliases
	}
	return nil
}

func (x *Resource) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Resource) GetResource() *any.Any {
	if x != nil {
		return x.Resource
	}
	return nil
}

var File_envoy_api_v2_discovery_proto protoreflect.FileDescriptor

var file_envoy_api_v2_discovery_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x02,
	0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x22, 0xf9, 0x01, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x22, 0xff,
	0x03, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x38, 0x0a, 0x18, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x16, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x3c, 0x0a, 0x1a, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x5f, 0x75, 0x6e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x55, 0x6e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x7c, 0x0a, 0x19, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x17, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x35, 0x0a,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x1a, 0x4a, 0x0a, 0x1c, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xdc, 0x01, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x2b, 0x0a, 0x11,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22,
	0x84, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x93, 0x01, 0x0a, 0x1a, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x42, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67,
	0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x70, 0x69,
	0x76, 0x32, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x1c, 0x12, 0x1a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_api_v2_discovery_proto_rawDescOnce sync.Once
	file_envoy_api_v2_discovery_proto_rawDescData = file_envoy_api_v2_discovery_proto_rawDesc
)

func file_envoy_api_v2_discovery_proto_rawDescGZIP() []byte {
	file_envoy_api_v2_discovery_proto_rawDescOnce.Do(func() {
		file_envoy_api_v2_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_api_v2_discovery_proto_rawDescData)
	})
	return file_envoy_api_v2_discovery_proto_rawDescData
}

var file_envoy_api_v2_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_envoy_api_v2_discovery_proto_goTypes = []interface{}{
	(*DiscoveryRequest)(nil),       // 0: envoy.api.v2.DiscoveryRequest
	(*DiscoveryResponse)(nil),      // 1: envoy.api.v2.DiscoveryResponse
	(*DeltaDiscoveryRequest)(nil),  // 2: envoy.api.v2.DeltaDiscoveryRequest
	(*DeltaDiscoveryResponse)(nil), // 3: envoy.api.v2.DeltaDiscoveryResponse
	(*Resource)(nil),               // 4: envoy.api.v2.Resource
	nil,                            // 5: envoy.api.v2.DeltaDiscoveryRequest.InitialResourceVersionsEntry
	(*core.Node)(nil),              // 6: envoy.api.v2.core.Node
	(*status.Status)(nil),          // 7: google.rpc.Status
	(*any.Any)(nil),                // 8: google.protobuf.Any
	(*core.ControlPlane)(nil),      // 9: envoy.api.v2.core.ControlPlane
}
var file_envoy_api_v2_discovery_proto_depIdxs = []int32{
	6, // 0: envoy.api.v2.DiscoveryRequest.node:type_name -> envoy.api.v2.core.Node
	7, // 1: envoy.api.v2.DiscoveryRequest.error_detail:type_name -> google.rpc.Status
	8, // 2: envoy.api.v2.DiscoveryResponse.resources:type_name -> google.protobuf.Any
	9, // 3: envoy.api.v2.DiscoveryResponse.control_plane:type_name -> envoy.api.v2.core.ControlPlane
	6, // 4: envoy.api.v2.DeltaDiscoveryRequest.node:type_name -> envoy.api.v2.core.Node
	5, // 5: envoy.api.v2.DeltaDiscoveryRequest.initial_resource_versions:type_name -> envoy.api.v2.DeltaDiscoveryRequest.InitialResourceVersionsEntry
	7, // 6: envoy.api.v2.DeltaDiscoveryRequest.error_detail:type_name -> google.rpc.Status
	4, // 7: envoy.api.v2.DeltaDiscoveryResponse.resources:type_name -> envoy.api.v2.Resource
	8, // 8: envoy.api.v2.Resource.resource:type_name -> google.protobuf.Any
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_envoy_api_v2_discovery_proto_init() }
func file_envoy_api_v2_discovery_proto_init() {
	if File_envoy_api_v2_discovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_api_v2_discovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_envoy_api_v2_discovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoveryResponse); i {
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
		file_envoy_api_v2_discovery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeltaDiscoveryRequest); i {
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
		file_envoy_api_v2_discovery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeltaDiscoveryResponse); i {
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
		file_envoy_api_v2_discovery_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
			RawDescriptor: file_envoy_api_v2_discovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_api_v2_discovery_proto_goTypes,
		DependencyIndexes: file_envoy_api_v2_discovery_proto_depIdxs,
		MessageInfos:      file_envoy_api_v2_discovery_proto_msgTypes,
	}.Build()
	File_envoy_api_v2_discovery_proto = out.File
	file_envoy_api_v2_discovery_proto_rawDesc = nil
	file_envoy_api_v2_discovery_proto_goTypes = nil
	file_envoy_api_v2_discovery_proto_depIdxs = nil
}
