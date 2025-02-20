// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/certificates/v1/certificate_request.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Possible states in which a CertificateRequest can exist.
type CertificateRequestStatus_State int32

const (
	// The CertificateRequest has yet to be picked up by the issuer.
	CertificateRequestStatus_PENDING CertificateRequestStatus_State = 0
	// The issuer has replied to the request and the `signedCertificate` and `signingRootCa`
	// status fields will be populated.
	CertificateRequestStatus_FINISHED CertificateRequestStatus_State = 1
	// Processing the certificate workflow failed.
	CertificateRequestStatus_FAILED CertificateRequestStatus_State = 2
)

// Enum value maps for CertificateRequestStatus_State.
var (
	CertificateRequestStatus_State_name = map[int32]string{
		0: "PENDING",
		1: "FINISHED",
		2: "FAILED",
	}
	CertificateRequestStatus_State_value = map[string]int32{
		"PENDING":  0,
		"FINISHED": 1,
		"FAILED":   2,
	}
)

func (x CertificateRequestStatus_State) Enum() *CertificateRequestStatus_State {
	p := new(CertificateRequestStatus_State)
	*p = x
	return p
}

func (x CertificateRequestStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CertificateRequestStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_enumTypes[0].Descriptor()
}

func (CertificateRequestStatus_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_enumTypes[0]
}

func (x CertificateRequestStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CertificateRequestStatus_State.Descriptor instead.
func (CertificateRequestStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_rawDescGZIP(), []int{1, 0}
}

//
//CertificateRequests are generated by the Gloo Mesh agent installed on managed clusters.
//They are used to request a signed certificate from the certificate issuer (the Gloo Mesh server) based on a private key
//generated by the agent (which never leaves the managed cluster).
//
//When Gloo Mesh creates an IssuedCertificate on a managed cluster, the local Gloo Mesh Agent
//will generate a CertificateRequest corresponding to it.
//
//Gloo Mesh will then process the certificate signing request contained in the
//`CertificateRequestSpec` and write the signed SSL certificate back as a Kubernetes secret in the managed cluster,
//and update the `CertificateRequestStatus` to point to that secret.
type CertificateRequestSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base64-encoded data for the PKCS#10 Certificate Signing Request issued
	// by the Gloo Mesh agent deployed in the managed cluster, corresponding
	// to the IssuedRequest received by the Gloo Mesh agent.
	CertificateSigningRequest []byte `protobuf:"bytes,1,opt,name=certificate_signing_request,json=certificateSigningRequest,proto3" json:"certificate_signing_request,omitempty"`
}

func (x *CertificateRequestSpec) Reset() {
	*x = CertificateRequestSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateRequestSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateRequestSpec) ProtoMessage() {}

func (x *CertificateRequestSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateRequestSpec.ProtoReflect.Descriptor instead.
func (*CertificateRequestSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_rawDescGZIP(), []int{0}
}

func (x *CertificateRequestSpec) GetCertificateSigningRequest() []byte {
	if x != nil {
		return x.CertificateSigningRequest
	}
	return nil
}

type CertificateRequestStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The most recent generation observed in the the CertificateRequest metadata.
	// If the `observedGeneration` does not match `metadata.generation`, the issuer has not processed the most
	// recent version of this request.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// Any error observed which prevented the CertificateRequest from being processed.
	// If the error is empty, the request has been processed successfully
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// The current state of the CertificateRequest workflow reported by the issuer.
	State CertificateRequestStatus_State `protobuf:"varint,3,opt,name=state,proto3,enum=certificates.mesh.gloo.solo.io.CertificateRequestStatus_State" json:"state,omitempty"`
	// The signed intermediate certificate issued by the CA.
	SignedCertificate []byte `protobuf:"bytes,4,opt,name=signed_certificate,json=signedCertificate,proto3" json:"signed_certificate,omitempty"`
	// The root CA used by the issuer to sign the certificate.
	SigningRootCa []byte `protobuf:"bytes,5,opt,name=signing_root_ca,json=signingRootCa,proto3" json:"signing_root_ca,omitempty"`
}

func (x *CertificateRequestStatus) Reset() {
	*x = CertificateRequestStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateRequestStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateRequestStatus) ProtoMessage() {}

func (x *CertificateRequestStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateRequestStatus.ProtoReflect.Descriptor instead.
func (*CertificateRequestStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_rawDescGZIP(), []int{1}
}

func (x *CertificateRequestStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *CertificateRequestStatus) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CertificateRequestStatus) GetState() CertificateRequestStatus_State {
	if x != nil {
		return x.State
	}
	return CertificateRequestStatus_PENDING
}

func (x *CertificateRequestStatus) GetSignedCertificate() []byte {
	if x != nil {
		return x.SignedCertificate
	}
	return nil
}

func (x *CertificateRequestStatus) GetSigningRootCa() []byte {
	if x != nil {
		return x.SigningRootCa
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65,
	0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x58, 0x0a, 0x16, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x3e, 0x0a, 0x1b, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x19, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xbe, 0x02, 0x0a, 0x18,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x54, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3e,
	0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x11, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x63, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x61, 0x22, 0x2e, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x42, 0x4c, 0x5a, 0x46,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_goTypes = []interface{}{
	(CertificateRequestStatus_State)(0), // 0: certificates.mesh.gloo.solo.io.CertificateRequestStatus.State
	(*CertificateRequestSpec)(nil),      // 1: certificates.mesh.gloo.solo.io.CertificateRequestSpec
	(*CertificateRequestStatus)(nil),    // 2: certificates.mesh.gloo.solo.io.CertificateRequestStatus
}
var file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_depIdxs = []int32{
	0, // 0: certificates.mesh.gloo.solo.io.CertificateRequestStatus.state:type_name -> certificates.mesh.gloo.solo.io.CertificateRequestStatus.State
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_init() }
func file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateRequestSpec); i {
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
		file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateRequestStatus); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_certificate_request_proto_depIdxs = nil
}
