// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/discovery/v1alpha2/workload.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//*
//The Workload is an abstraction for a workload/client which mesh-discovery has discovered to be part of a
//given mesh (i.e. its traffic is managed by an in-mesh sidecar).
type WorkloadSpec struct {
	// Specifies the underlying type of workload that this Workload is abstracting.
	//
	// Types that are valid to be assigned to WorkloadType:
	//	*WorkloadSpec_Kubernetes
	WorkloadType isWorkloadSpec_WorkloadType `protobuf_oneof:"workload_type"`
	// The mesh with which this workload is associated.
	Mesh *v1.ObjectRef `protobuf:"bytes,4,opt,name=mesh,proto3" json:"mesh,omitempty"`
	// instances of the proxy running within this workload (e.g Envoy processes)
	ProxyInstances []*WorkloadSpec_ProxyInstance `protobuf:"bytes,6,rep,name=proxy_instances,json=proxyInstances,proto3" json:"proxy_instances,omitempty"`
	// Appmesh specific metadata.
	AppMesh              *WorkloadSpec_AppMesh `protobuf:"bytes,5,opt,name=app_mesh,json=appMesh,proto3" json:"app_mesh,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *WorkloadSpec) Reset()         { *m = WorkloadSpec{} }
func (m *WorkloadSpec) String() string { return proto.CompactTextString(m) }
func (*WorkloadSpec) ProtoMessage()    {}
func (*WorkloadSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_15f8df912dc5fd52, []int{0}
}
func (m *WorkloadSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadSpec.Unmarshal(m, b)
}
func (m *WorkloadSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadSpec.Marshal(b, m, deterministic)
}
func (m *WorkloadSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadSpec.Merge(m, src)
}
func (m *WorkloadSpec) XXX_Size() int {
	return xxx_messageInfo_WorkloadSpec.Size(m)
}
func (m *WorkloadSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadSpec.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadSpec proto.InternalMessageInfo

type isWorkloadSpec_WorkloadType interface {
	isWorkloadSpec_WorkloadType()
	Equal(interface{}) bool
}

type WorkloadSpec_Kubernetes struct {
	Kubernetes *WorkloadSpec_KubernetesWorkload `protobuf:"bytes,1,opt,name=kubernetes,proto3,oneof" json:"kubernetes,omitempty"`
}

func (*WorkloadSpec_Kubernetes) isWorkloadSpec_WorkloadType() {}

func (m *WorkloadSpec) GetWorkloadType() isWorkloadSpec_WorkloadType {
	if m != nil {
		return m.WorkloadType
	}
	return nil
}

func (m *WorkloadSpec) GetKubernetes() *WorkloadSpec_KubernetesWorkload {
	if x, ok := m.GetWorkloadType().(*WorkloadSpec_Kubernetes); ok {
		return x.Kubernetes
	}
	return nil
}

func (m *WorkloadSpec) GetMesh() *v1.ObjectRef {
	if m != nil {
		return m.Mesh
	}
	return nil
}

func (m *WorkloadSpec) GetProxyInstances() []*WorkloadSpec_ProxyInstance {
	if m != nil {
		return m.ProxyInstances
	}
	return nil
}

func (m *WorkloadSpec) GetAppMesh() *WorkloadSpec_AppMesh {
	if m != nil {
		return m.AppMesh
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*WorkloadSpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*WorkloadSpec_Kubernetes)(nil),
	}
}

// Information describing a Kubernetes-based workload (e.g. a Deployment or DaemonSet).
type WorkloadSpec_KubernetesWorkload struct {
	//*
	//Resource ref to the underlying kubernetes controller which is managing the pods associated with the workloads.
	//It has the generic name controller as it can represent a deployment, daemonset, or statefulset.
	Controller *v1.ClusterObjectRef `protobuf:"bytes,1,opt,name=controller,proto3" json:"controller,omitempty"`
	//*
	//These are the labels directly from the pods that this controller owns.
	//NB: these labels are read directly from the pod template metadata.labels
	//defined in the workload spec.
	//We need these to determine which services are backed by this workload.
	PodLabels map[string]string `protobuf:"bytes,2,rep,name=pod_labels,json=podLabels,proto3" json:"pod_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Service account attached to the pods owned by this controller.
	ServiceAccountName   string   `protobuf:"bytes,3,opt,name=service_account_name,json=serviceAccountName,proto3" json:"service_account_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkloadSpec_KubernetesWorkload) Reset()         { *m = WorkloadSpec_KubernetesWorkload{} }
func (m *WorkloadSpec_KubernetesWorkload) String() string { return proto.CompactTextString(m) }
func (*WorkloadSpec_KubernetesWorkload) ProtoMessage()    {}
func (*WorkloadSpec_KubernetesWorkload) Descriptor() ([]byte, []int) {
	return fileDescriptor_15f8df912dc5fd52, []int{0, 0}
}
func (m *WorkloadSpec_KubernetesWorkload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadSpec_KubernetesWorkload.Unmarshal(m, b)
}
func (m *WorkloadSpec_KubernetesWorkload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadSpec_KubernetesWorkload.Marshal(b, m, deterministic)
}
func (m *WorkloadSpec_KubernetesWorkload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadSpec_KubernetesWorkload.Merge(m, src)
}
func (m *WorkloadSpec_KubernetesWorkload) XXX_Size() int {
	return xxx_messageInfo_WorkloadSpec_KubernetesWorkload.Size(m)
}
func (m *WorkloadSpec_KubernetesWorkload) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadSpec_KubernetesWorkload.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadSpec_KubernetesWorkload proto.InternalMessageInfo

func (m *WorkloadSpec_KubernetesWorkload) GetController() *v1.ClusterObjectRef {
	if m != nil {
		return m.Controller
	}
	return nil
}

func (m *WorkloadSpec_KubernetesWorkload) GetPodLabels() map[string]string {
	if m != nil {
		return m.PodLabels
	}
	return nil
}

func (m *WorkloadSpec_KubernetesWorkload) GetServiceAccountName() string {
	if m != nil {
		return m.ServiceAccountName
	}
	return ""
}

// Describes an instance of a single workload proxy (e.g. a Pod container or VM process)
type WorkloadSpec_ProxyInstance struct {
	// the node ID of the proxy (Envoy) running on this instance
	ProxyNodeId          string   `protobuf:"bytes,1,opt,name=proxy_node_id,json=proxyNodeId,proto3" json:"proxy_node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkloadSpec_ProxyInstance) Reset()         { *m = WorkloadSpec_ProxyInstance{} }
func (m *WorkloadSpec_ProxyInstance) String() string { return proto.CompactTextString(m) }
func (*WorkloadSpec_ProxyInstance) ProtoMessage()    {}
func (*WorkloadSpec_ProxyInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_15f8df912dc5fd52, []int{0, 1}
}
func (m *WorkloadSpec_ProxyInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadSpec_ProxyInstance.Unmarshal(m, b)
}
func (m *WorkloadSpec_ProxyInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadSpec_ProxyInstance.Marshal(b, m, deterministic)
}
func (m *WorkloadSpec_ProxyInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadSpec_ProxyInstance.Merge(m, src)
}
func (m *WorkloadSpec_ProxyInstance) XXX_Size() int {
	return xxx_messageInfo_WorkloadSpec_ProxyInstance.Size(m)
}
func (m *WorkloadSpec_ProxyInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadSpec_ProxyInstance.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadSpec_ProxyInstance proto.InternalMessageInfo

func (m *WorkloadSpec_ProxyInstance) GetProxyNodeId() string {
	if m != nil {
		return m.ProxyNodeId
	}
	return ""
}

// Information relevant to AppMesh-injected workloads.
type WorkloadSpec_AppMesh struct {
	// The value of the env var APPMESH_VIRTUAL_NODE_NAME on the Appmesh envoy proxy container.
	VirtualNodeName string `protobuf:"bytes,1,opt,name=virtual_node_name,json=virtualNodeName,proto3" json:"virtual_node_name,omitempty"`
	// Needed for declaring Appmesh VirtualNode listeners.
	Ports                []*WorkloadSpec_AppMesh_ContainerPort `protobuf:"bytes,2,rep,name=ports,proto3" json:"ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *WorkloadSpec_AppMesh) Reset()         { *m = WorkloadSpec_AppMesh{} }
func (m *WorkloadSpec_AppMesh) String() string { return proto.CompactTextString(m) }
func (*WorkloadSpec_AppMesh) ProtoMessage()    {}
func (*WorkloadSpec_AppMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_15f8df912dc5fd52, []int{0, 2}
}
func (m *WorkloadSpec_AppMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadSpec_AppMesh.Unmarshal(m, b)
}
func (m *WorkloadSpec_AppMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadSpec_AppMesh.Marshal(b, m, deterministic)
}
func (m *WorkloadSpec_AppMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadSpec_AppMesh.Merge(m, src)
}
func (m *WorkloadSpec_AppMesh) XXX_Size() int {
	return xxx_messageInfo_WorkloadSpec_AppMesh.Size(m)
}
func (m *WorkloadSpec_AppMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadSpec_AppMesh.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadSpec_AppMesh proto.InternalMessageInfo

func (m *WorkloadSpec_AppMesh) GetVirtualNodeName() string {
	if m != nil {
		return m.VirtualNodeName
	}
	return ""
}

func (m *WorkloadSpec_AppMesh) GetPorts() []*WorkloadSpec_AppMesh_ContainerPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

// k8s application container ports.
type WorkloadSpec_AppMesh_ContainerPort struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Protocol             string   `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkloadSpec_AppMesh_ContainerPort) Reset()         { *m = WorkloadSpec_AppMesh_ContainerPort{} }
func (m *WorkloadSpec_AppMesh_ContainerPort) String() string { return proto.CompactTextString(m) }
func (*WorkloadSpec_AppMesh_ContainerPort) ProtoMessage()    {}
func (*WorkloadSpec_AppMesh_ContainerPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_15f8df912dc5fd52, []int{0, 2, 0}
}
func (m *WorkloadSpec_AppMesh_ContainerPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadSpec_AppMesh_ContainerPort.Unmarshal(m, b)
}
func (m *WorkloadSpec_AppMesh_ContainerPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadSpec_AppMesh_ContainerPort.Marshal(b, m, deterministic)
}
func (m *WorkloadSpec_AppMesh_ContainerPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadSpec_AppMesh_ContainerPort.Merge(m, src)
}
func (m *WorkloadSpec_AppMesh_ContainerPort) XXX_Size() int {
	return xxx_messageInfo_WorkloadSpec_AppMesh_ContainerPort.Size(m)
}
func (m *WorkloadSpec_AppMesh_ContainerPort) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadSpec_AppMesh_ContainerPort.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadSpec_AppMesh_ContainerPort proto.InternalMessageInfo

func (m *WorkloadSpec_AppMesh_ContainerPort) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *WorkloadSpec_AppMesh_ContainerPort) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type WorkloadStatus struct {
	// The observed generation of the Workload.
	// When this matches the Workload's metadata.generation it indicates that mesh-networking
	// has reconciled the latest version of the Workload.
	ObservedGeneration   int64    `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkloadStatus) Reset()         { *m = WorkloadStatus{} }
func (m *WorkloadStatus) String() string { return proto.CompactTextString(m) }
func (*WorkloadStatus) ProtoMessage()    {}
func (*WorkloadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_15f8df912dc5fd52, []int{1}
}
func (m *WorkloadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadStatus.Unmarshal(m, b)
}
func (m *WorkloadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadStatus.Marshal(b, m, deterministic)
}
func (m *WorkloadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadStatus.Merge(m, src)
}
func (m *WorkloadStatus) XXX_Size() int {
	return xxx_messageInfo_WorkloadStatus.Size(m)
}
func (m *WorkloadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadStatus proto.InternalMessageInfo

func (m *WorkloadStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func init() {
	proto.RegisterType((*WorkloadSpec)(nil), "discovery.mesh.gloo.solo.io.WorkloadSpec")
	proto.RegisterType((*WorkloadSpec_KubernetesWorkload)(nil), "discovery.mesh.gloo.solo.io.WorkloadSpec.KubernetesWorkload")
	proto.RegisterMapType((map[string]string)(nil), "discovery.mesh.gloo.solo.io.WorkloadSpec.KubernetesWorkload.PodLabelsEntry")
	proto.RegisterType((*WorkloadSpec_ProxyInstance)(nil), "discovery.mesh.gloo.solo.io.WorkloadSpec.ProxyInstance")
	proto.RegisterType((*WorkloadSpec_AppMesh)(nil), "discovery.mesh.gloo.solo.io.WorkloadSpec.AppMesh")
	proto.RegisterType((*WorkloadSpec_AppMesh_ContainerPort)(nil), "discovery.mesh.gloo.solo.io.WorkloadSpec.AppMesh.ContainerPort")
	proto.RegisterType((*WorkloadStatus)(nil), "discovery.mesh.gloo.solo.io.WorkloadStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-mesh/api/discovery/v1alpha2/workload.proto", fileDescriptor_15f8df912dc5fd52)
}

var fileDescriptor_15f8df912dc5fd52 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0x92, 0xf4, 0xf7, 0xf6, 0x4b, 0x0b, 0x43, 0x17, 0x91, 0x8b, 0x50, 0x55, 0x36, 0x15,
	0x52, 0xed, 0x36, 0x5d, 0x80, 0x50, 0xa5, 0xaa, 0xad, 0x10, 0x54, 0x2d, 0x6d, 0x65, 0x84, 0x90,
	0x58, 0x60, 0x26, 0xf6, 0x25, 0x71, 0xe3, 0xcc, 0x1d, 0xcd, 0x8c, 0x0d, 0x79, 0x19, 0xd6, 0xbc,
	0x00, 0x8f, 0xc1, 0x4b, 0xf0, 0x24, 0x68, 0xc6, 0x8e, 0x49, 0x54, 0x54, 0x15, 0xb1, 0xf2, 0xfd,
	0xf1, 0x39, 0xf7, 0xcc, 0xb9, 0x9a, 0x81, 0xe3, 0x7e, 0x6a, 0x06, 0x79, 0xcf, 0x8f, 0x69, 0x14,
	0x68, 0xca, 0x68, 0x27, 0xa5, 0xa0, 0x9f, 0x11, 0xed, 0x8c, 0x50, 0x0f, 0x02, 0x2e, 0xd3, 0x20,
	0x49, 0x75, 0x4c, 0x05, 0xaa, 0x71, 0x50, 0xec, 0xf1, 0x4c, 0x0e, 0x78, 0x37, 0xf8, 0x4c, 0x6a,
	0x98, 0x11, 0x4f, 0x7c, 0xa9, 0xc8, 0x10, 0xdb, 0xa8, 0x7f, 0xf1, 0x2d, 0xca, 0xb7, 0x78, 0xdf,
	0x92, 0xf9, 0x29, 0x79, 0x1b, 0x7a, 0x58, 0x74, 0x1d, 0x51, 0x4c, 0x0a, 0x83, 0x62, 0xcf, 0x7d,
	0x4b, 0xa4, 0xb7, 0xde, 0xa7, 0x3e, 0xb9, 0x30, 0xb0, 0x51, 0x59, 0xdd, 0xfa, 0xba, 0x08, 0xff,
	0xbf, 0xab, 0x46, 0xbc, 0x91, 0x18, 0xb3, 0x0f, 0x00, 0xc3, 0xbc, 0x87, 0x4a, 0xa0, 0x41, 0xdd,
	0x69, 0x6c, 0x36, 0xb6, 0x57, 0xba, 0x07, 0xfe, 0x2d, 0x53, 0xfd, 0x69, 0xb8, 0x7f, 0x56, 0x63,
	0x27, 0xe5, 0x57, 0xff, 0x85, 0x53, 0x8c, 0x6c, 0x17, 0xe6, 0x2c, 0x45, 0x67, 0xce, 0x31, 0x3f,
	0xf4, 0x9d, 0x42, 0xab, 0xbb, 0xe6, 0xbb, 0xec, 0x5d, 0x63, 0x6c, 0x42, 0xfc, 0x14, 0xba, 0x3f,
	0xd9, 0x47, 0x58, 0x93, 0x8a, 0xbe, 0x8c, 0xa3, 0x54, 0x68, 0xc3, 0x45, 0x8c, 0xba, 0xb3, 0xb0,
	0xd9, 0xda, 0x5e, 0xe9, 0x3e, 0xbd, 0xbb, 0xac, 0x2b, 0x4b, 0x70, 0x5a, 0xe1, 0xc3, 0x55, 0x39,
	0x9d, 0x6a, 0x76, 0x0e, 0x4b, 0x5c, 0xca, 0xc8, 0xe9, 0x9a, 0x77, 0xba, 0xf6, 0xee, 0x4e, 0x7d,
	0x24, 0xe5, 0x6b, 0xd4, 0x83, 0x70, 0x91, 0x97, 0x81, 0xf7, 0xbd, 0x09, 0xec, 0xa6, 0x0d, 0xec,
	0x04, 0x20, 0x26, 0x61, 0x14, 0x65, 0x19, 0xaa, 0xca, 0xd8, 0xc7, 0x7f, 0x38, 0xfe, 0x49, 0x96,
	0x6b, 0x83, 0xea, 0xb7, 0x0b, 0x53, 0x30, 0x76, 0x0d, 0x20, 0x29, 0x89, 0x32, 0xde, 0xc3, 0x4c,
	0x77, 0x9a, 0xce, 0x86, 0xb3, 0x7f, 0xd9, 0x8e, 0x7f, 0x45, 0xc9, 0xb9, 0x63, 0x7b, 0x21, 0x8c,
	0x1a, 0x87, 0xcb, 0x72, 0x92, 0xb3, 0x5d, 0x58, 0xd7, 0xa8, 0x8a, 0x34, 0xc6, 0x88, 0xc7, 0x31,
	0xe5, 0xc2, 0x44, 0x82, 0x8f, 0xb0, 0xd3, 0xda, 0x6c, 0x6c, 0x2f, 0x87, 0xac, 0xea, 0x1d, 0x95,
	0xad, 0x0b, 0x3e, 0x42, 0xef, 0x00, 0x56, 0x67, 0xe9, 0xd8, 0x3d, 0x68, 0x0d, 0x71, 0xec, 0x4e,
	0xbb, 0x1c, 0xda, 0x90, 0xad, 0xc3, 0x7c, 0xc1, 0xb3, 0x1c, 0x3b, 0x4d, 0x57, 0x2b, 0x93, 0xe7,
	0xcd, 0x67, 0x0d, 0x6f, 0x1f, 0xda, 0x33, 0x6b, 0x62, 0x5b, 0xd0, 0x2e, 0x17, 0x2f, 0x28, 0xc1,
	0x28, 0x4d, 0x2a, 0x9a, 0x15, 0x57, 0xbc, 0xa0, 0x04, 0x4f, 0x13, 0xef, 0x47, 0x03, 0x16, 0xab,
	0x0d, 0xb0, 0x27, 0x70, 0xbf, 0x48, 0x95, 0xc9, 0x79, 0x56, 0x22, 0x9c, 0xda, 0x12, 0xb3, 0x56,
	0x35, 0x2c, 0xca, 0x4a, 0x65, 0x6f, 0x61, 0x5e, 0x92, 0x32, 0x13, 0x0f, 0x0f, 0xff, 0x7a, 0xdf,
	0xfe, 0x09, 0x09, 0xc3, 0x53, 0x81, 0xea, 0x8a, 0x94, 0x09, 0x4b, 0x36, 0xef, 0x10, 0xda, 0x33,
	0x75, 0xc6, 0x60, 0xce, 0x76, 0x9c, 0x8c, 0x76, 0xe8, 0x62, 0xe6, 0xc1, 0x92, 0xbb, 0x7c, 0x31,
	0x65, 0x95, 0x0b, 0x75, 0x7e, 0xbc, 0x06, 0xed, 0xc9, 0x8d, 0x8f, 0xcc, 0x58, 0xe2, 0xd6, 0x11,
	0xac, 0xd6, 0xe3, 0x0d, 0x37, 0xb9, 0x66, 0x01, 0x3c, 0xa0, 0x9e, 0x75, 0x1f, 0x93, 0xa8, 0x8f,
	0x02, 0x15, 0x37, 0x29, 0x09, 0x37, 0xa1, 0x15, 0xb2, 0x49, 0xeb, 0x65, 0xdd, 0x39, 0xbe, 0xfc,
	0xf6, 0xf3, 0x51, 0xe3, 0xfd, 0xe9, 0xad, 0xaf, 0x8f, 0x1c, 0xf6, 0x67, 0x5f, 0xa0, 0x9b, 0x36,
	0xd4, 0x6f, 0x52, 0x6f, 0xc1, 0xc9, 0xdd, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0x16, 0xf4, 0xe7,
	0xa6, 0xd1, 0x04, 0x00, 0x00,
}

func (this *WorkloadSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkloadSpec)
	if !ok {
		that2, ok := that.(WorkloadSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.WorkloadType == nil {
		if this.WorkloadType != nil {
			return false
		}
	} else if this.WorkloadType == nil {
		return false
	} else if !this.WorkloadType.Equal(that1.WorkloadType) {
		return false
	}
	if !this.Mesh.Equal(that1.Mesh) {
		return false
	}
	if len(this.ProxyInstances) != len(that1.ProxyInstances) {
		return false
	}
	for i := range this.ProxyInstances {
		if !this.ProxyInstances[i].Equal(that1.ProxyInstances[i]) {
			return false
		}
	}
	if !this.AppMesh.Equal(that1.AppMesh) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *WorkloadSpec_Kubernetes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkloadSpec_Kubernetes)
	if !ok {
		that2, ok := that.(WorkloadSpec_Kubernetes)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Kubernetes.Equal(that1.Kubernetes) {
		return false
	}
	return true
}
func (this *WorkloadSpec_KubernetesWorkload) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkloadSpec_KubernetesWorkload)
	if !ok {
		that2, ok := that.(WorkloadSpec_KubernetesWorkload)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Controller.Equal(that1.Controller) {
		return false
	}
	if len(this.PodLabels) != len(that1.PodLabels) {
		return false
	}
	for i := range this.PodLabels {
		if this.PodLabels[i] != that1.PodLabels[i] {
			return false
		}
	}
	if this.ServiceAccountName != that1.ServiceAccountName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *WorkloadSpec_ProxyInstance) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkloadSpec_ProxyInstance)
	if !ok {
		that2, ok := that.(WorkloadSpec_ProxyInstance)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ProxyNodeId != that1.ProxyNodeId {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *WorkloadSpec_AppMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkloadSpec_AppMesh)
	if !ok {
		that2, ok := that.(WorkloadSpec_AppMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.VirtualNodeName != that1.VirtualNodeName {
		return false
	}
	if len(this.Ports) != len(that1.Ports) {
		return false
	}
	for i := range this.Ports {
		if !this.Ports[i].Equal(that1.Ports[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *WorkloadSpec_AppMesh_ContainerPort) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkloadSpec_AppMesh_ContainerPort)
	if !ok {
		that2, ok := that.(WorkloadSpec_AppMesh_ContainerPort)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if this.Protocol != that1.Protocol {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *WorkloadStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkloadStatus)
	if !ok {
		that2, ok := that.(WorkloadStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ObservedGeneration != that1.ObservedGeneration {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
