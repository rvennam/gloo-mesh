// Code generated by skv2. DO NOT EDIT.

/*
	Utility for manually building input snapshots. Used primarily in tests.
*/
package input

import (
	certificates_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1alpha2"
	certificates_mesh_gloo_solo_io_v1alpha2_sets "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1alpha2/sets"
)

type InputSnapshotManualBuilder struct {
	name string

	certificatesMeshGlooSoloIoV1Alpha2IssuedCertificates  certificates_mesh_gloo_solo_io_v1alpha2_sets.IssuedCertificateSet
	certificatesMeshGlooSoloIoV1Alpha2CertificateRequests certificates_mesh_gloo_solo_io_v1alpha2_sets.CertificateRequestSet
}

func NewInputSnapshotManualBuilder(name string) *InputSnapshotManualBuilder {
	return &InputSnapshotManualBuilder{
		name: name,

		certificatesMeshGlooSoloIoV1Alpha2IssuedCertificates:  certificates_mesh_gloo_solo_io_v1alpha2_sets.NewIssuedCertificateSet(),
		certificatesMeshGlooSoloIoV1Alpha2CertificateRequests: certificates_mesh_gloo_solo_io_v1alpha2_sets.NewCertificateRequestSet(),
	}
}

func (i *InputSnapshotManualBuilder) Build() Snapshot {
	return NewSnapshot(
		i.name,

		i.certificatesMeshGlooSoloIoV1Alpha2IssuedCertificates,
		i.certificatesMeshGlooSoloIoV1Alpha2CertificateRequests,
	)
}
func (i *InputSnapshotManualBuilder) AddCertificatesMeshGlooSoloIo_V1Alpha2_IssuedCertificates(certificatesMeshGlooSoloIoV1Alpha2IssuedCertificates []*certificates_mesh_gloo_solo_io_v1alpha2.IssuedCertificate) *InputSnapshotManualBuilder {
	i.certificatesMeshGlooSoloIoV1Alpha2IssuedCertificates.Insert(certificatesMeshGlooSoloIoV1Alpha2IssuedCertificates...)
	return i
}
func (i *InputSnapshotManualBuilder) AddCertificatesMeshGlooSoloIo_V1Alpha2_CertificateRequests(certificatesMeshGlooSoloIoV1Alpha2CertificateRequests []*certificates_mesh_gloo_solo_io_v1alpha2.CertificateRequest) *InputSnapshotManualBuilder {
	i.certificatesMeshGlooSoloIoV1Alpha2CertificateRequests.Insert(certificatesMeshGlooSoloIoV1Alpha2CertificateRequests...)
	return i
}
