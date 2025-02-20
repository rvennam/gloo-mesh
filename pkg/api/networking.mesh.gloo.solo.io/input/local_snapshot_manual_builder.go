// Code generated by skv2. DO NOT EDIT.

/*
	Utility for manually building input snapshots. Used primarily in tests.
*/
package input

import (
	settings_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1"
	settings_mesh_gloo_solo_io_v1_sets "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1/sets"

	discovery_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	discovery_mesh_gloo_solo_io_v1_sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1/sets"

	networking_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	networking_mesh_gloo_solo_io_v1_sets "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1/sets"

	networking_enterprise_mesh_gloo_solo_io_v1beta1 "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1beta1"
	networking_enterprise_mesh_gloo_solo_io_v1beta1_sets "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1beta1/sets"

	observability_enterprise_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/observability.enterprise.mesh.gloo.solo.io/v1"
	observability_enterprise_mesh_gloo_solo_io_v1_sets "github.com/solo-io/gloo-mesh/pkg/api/observability.enterprise.mesh.gloo.solo.io/v1/sets"

	v1_sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1 "k8s.io/api/core/v1"

	multicluster_solo_io_v1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	multicluster_solo_io_v1alpha1_sets "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/sets"
)

type InputLocalSnapshotManualBuilder struct {
	name string

	settings settings_mesh_gloo_solo_io_v1_sets.SettingsSet

	destinations discovery_mesh_gloo_solo_io_v1_sets.DestinationSet
	workloads    discovery_mesh_gloo_solo_io_v1_sets.WorkloadSet
	meshes       discovery_mesh_gloo_solo_io_v1_sets.MeshSet

	trafficPolicies networking_mesh_gloo_solo_io_v1_sets.TrafficPolicySet
	accessPolicies  networking_mesh_gloo_solo_io_v1_sets.AccessPolicySet
	virtualMeshes   networking_mesh_gloo_solo_io_v1_sets.VirtualMeshSet

	wasmDeployments     networking_enterprise_mesh_gloo_solo_io_v1beta1_sets.WasmDeploymentSet
	virtualDestinations networking_enterprise_mesh_gloo_solo_io_v1beta1_sets.VirtualDestinationSet

	accessLogRecords observability_enterprise_mesh_gloo_solo_io_v1_sets.AccessLogRecordSet

	secrets v1_sets.SecretSet

	kubernetesClusters multicluster_solo_io_v1alpha1_sets.KubernetesClusterSet
}

func NewInputLocalSnapshotManualBuilder(name string) *InputLocalSnapshotManualBuilder {
	return &InputLocalSnapshotManualBuilder{
		name: name,

		settings: settings_mesh_gloo_solo_io_v1_sets.NewSettingsSet(),

		destinations: discovery_mesh_gloo_solo_io_v1_sets.NewDestinationSet(),
		workloads:    discovery_mesh_gloo_solo_io_v1_sets.NewWorkloadSet(),
		meshes:       discovery_mesh_gloo_solo_io_v1_sets.NewMeshSet(),

		trafficPolicies: networking_mesh_gloo_solo_io_v1_sets.NewTrafficPolicySet(),
		accessPolicies:  networking_mesh_gloo_solo_io_v1_sets.NewAccessPolicySet(),
		virtualMeshes:   networking_mesh_gloo_solo_io_v1_sets.NewVirtualMeshSet(),

		wasmDeployments:     networking_enterprise_mesh_gloo_solo_io_v1beta1_sets.NewWasmDeploymentSet(),
		virtualDestinations: networking_enterprise_mesh_gloo_solo_io_v1beta1_sets.NewVirtualDestinationSet(),

		accessLogRecords: observability_enterprise_mesh_gloo_solo_io_v1_sets.NewAccessLogRecordSet(),

		secrets: v1_sets.NewSecretSet(),

		kubernetesClusters: multicluster_solo_io_v1alpha1_sets.NewKubernetesClusterSet(),
	}
}

func (i *InputLocalSnapshotManualBuilder) Build() LocalSnapshot {
	return NewLocalSnapshot(
		i.name,

		i.settings,

		i.destinations,
		i.workloads,
		i.meshes,

		i.trafficPolicies,
		i.accessPolicies,
		i.virtualMeshes,

		i.wasmDeployments,
		i.virtualDestinations,

		i.accessLogRecords,

		i.secrets,

		i.kubernetesClusters,
	)
}
func (i *InputLocalSnapshotManualBuilder) AddSettings(settings []*settings_mesh_gloo_solo_io_v1.Settings) *InputLocalSnapshotManualBuilder {
	i.settings.Insert(settings...)
	return i
}
func (i *InputLocalSnapshotManualBuilder) AddDestinations(destinations []*discovery_mesh_gloo_solo_io_v1.Destination) *InputLocalSnapshotManualBuilder {
	i.destinations.Insert(destinations...)
	return i
}
func (i *InputLocalSnapshotManualBuilder) AddWorkloads(workloads []*discovery_mesh_gloo_solo_io_v1.Workload) *InputLocalSnapshotManualBuilder {
	i.workloads.Insert(workloads...)
	return i
}
func (i *InputLocalSnapshotManualBuilder) AddMeshes(meshes []*discovery_mesh_gloo_solo_io_v1.Mesh) *InputLocalSnapshotManualBuilder {
	i.meshes.Insert(meshes...)
	return i
}
func (i *InputLocalSnapshotManualBuilder) AddTrafficPolicies(trafficPolicies []*networking_mesh_gloo_solo_io_v1.TrafficPolicy) *InputLocalSnapshotManualBuilder {
	i.trafficPolicies.Insert(trafficPolicies...)
	return i
}
func (i *InputLocalSnapshotManualBuilder) AddAccessPolicies(accessPolicies []*networking_mesh_gloo_solo_io_v1.AccessPolicy) *InputLocalSnapshotManualBuilder {
	i.accessPolicies.Insert(accessPolicies...)
	return i
}
func (i *InputLocalSnapshotManualBuilder) AddVirtualMeshes(virtualMeshes []*networking_mesh_gloo_solo_io_v1.VirtualMesh) *InputLocalSnapshotManualBuilder {
	i.virtualMeshes.Insert(virtualMeshes...)
	return i
}
func (i *InputLocalSnapshotManualBuilder) AddWasmDeployments(wasmDeployments []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) *InputLocalSnapshotManualBuilder {
	i.wasmDeployments.Insert(wasmDeployments...)
	return i
}
func (i *InputLocalSnapshotManualBuilder) AddVirtualDestinations(virtualDestinations []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) *InputLocalSnapshotManualBuilder {
	i.virtualDestinations.Insert(virtualDestinations...)
	return i
}
func (i *InputLocalSnapshotManualBuilder) AddAccessLogRecords(accessLogRecords []*observability_enterprise_mesh_gloo_solo_io_v1.AccessLogRecord) *InputLocalSnapshotManualBuilder {
	i.accessLogRecords.Insert(accessLogRecords...)
	return i
}
func (i *InputLocalSnapshotManualBuilder) AddSecrets(secrets []*v1.Secret) *InputLocalSnapshotManualBuilder {
	i.secrets.Insert(secrets...)
	return i
}
func (i *InputLocalSnapshotManualBuilder) AddKubernetesClusters(kubernetesClusters []*multicluster_solo_io_v1alpha1.KubernetesCluster) *InputLocalSnapshotManualBuilder {
	i.kubernetesClusters.Insert(kubernetesClusters...)
	return i
}
