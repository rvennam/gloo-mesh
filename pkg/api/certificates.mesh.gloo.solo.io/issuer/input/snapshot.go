// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./snapshot.go -destination mocks/snapshot.go

// The Input Snapshot contains the set of all:
// * IssuedCertificates
// * CertificateRequests
// read from a given cluster or set of clusters, across all namespaces.
//
// A snapshot can be constructed from either a single Manager (for a single cluster)
// or a ClusterWatcher (for multiple clusters) using the SnapshotBuilder.
//
// Resources in a MultiCluster snapshot will have their ClusterName set to the
// name of the cluster from which the resource was read.

package input

import (
	"context"
	"encoding/json"

	"github.com/solo-io/skv2/pkg/resource"
	"github.com/solo-io/skv2/pkg/verifier"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/hashicorp/go-multierror"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	certificates_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1"
	certificates_mesh_gloo_solo_io_v1_types "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1"
	certificates_mesh_gloo_solo_io_v1_sets "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1/sets"
)

// SnapshotGVKs is a list of the GVKs included in this snapshot
var SnapshotGVKs = []schema.GroupVersionKind{

	schema.GroupVersionKind{
		Group:   "certificates.mesh.gloo.solo.io",
		Version: "v1",
		Kind:    "IssuedCertificate",
	},
	schema.GroupVersionKind{
		Group:   "certificates.mesh.gloo.solo.io",
		Version: "v1",
		Kind:    "CertificateRequest",
	},
}

// the snapshot of input resources consumed by translation
type Snapshot interface {

	// return the set of input IssuedCertificates
	IssuedCertificates() certificates_mesh_gloo_solo_io_v1_sets.IssuedCertificateSet
	// return the set of input CertificateRequests
	CertificateRequests() certificates_mesh_gloo_solo_io_v1_sets.CertificateRequestSet
	// update the status of all input objects which support
	// the Status subresource (across multiple clusters)
	SyncStatusesMultiCluster(ctx context.Context, mcClient multicluster.Client, opts SyncStatusOptions) error
	// update the status of all input objects which support
	// the Status subresource (in the local cluster)
	SyncStatuses(ctx context.Context, c client.Client, opts SyncStatusOptions) error
	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)
}

// options for syncing input object statuses
type SyncStatusOptions struct {

	// sync status of IssuedCertificate objects
	IssuedCertificate bool
	// sync status of CertificateRequest objects
	CertificateRequest bool
}

type snapshot struct {
	name string

	issuedCertificates  certificates_mesh_gloo_solo_io_v1_sets.IssuedCertificateSet
	certificateRequests certificates_mesh_gloo_solo_io_v1_sets.CertificateRequestSet
}

func NewSnapshot(
	name string,

	issuedCertificates certificates_mesh_gloo_solo_io_v1_sets.IssuedCertificateSet,
	certificateRequests certificates_mesh_gloo_solo_io_v1_sets.CertificateRequestSet,

) Snapshot {
	return &snapshot{
		name: name,

		issuedCertificates:  issuedCertificates,
		certificateRequests: certificateRequests,
	}
}

func NewSnapshotFromGeneric(
	name string,
	genericSnapshot resource.ClusterSnapshot,
) Snapshot {

	issuedCertificateSet := certificates_mesh_gloo_solo_io_v1_sets.NewIssuedCertificateSet()
	certificateRequestSet := certificates_mesh_gloo_solo_io_v1_sets.NewCertificateRequestSet()

	for _, snapshot := range genericSnapshot {

		issuedCertificates := snapshot[schema.GroupVersionKind{
			Group:   "certificates.mesh.gloo.solo.io",
			Version: "v1",
			Kind:    "IssuedCertificate",
		}]

		for _, issuedCertificate := range issuedCertificates {
			issuedCertificateSet.Insert(issuedCertificate.(*certificates_mesh_gloo_solo_io_v1_types.IssuedCertificate))
		}
		certificateRequests := snapshot[schema.GroupVersionKind{
			Group:   "certificates.mesh.gloo.solo.io",
			Version: "v1",
			Kind:    "CertificateRequest",
		}]

		for _, certificateRequest := range certificateRequests {
			certificateRequestSet.Insert(certificateRequest.(*certificates_mesh_gloo_solo_io_v1_types.CertificateRequest))
		}

	}
	return NewSnapshot(
		name,
		issuedCertificateSet,
		certificateRequestSet,
	)
}

func (s snapshot) IssuedCertificates() certificates_mesh_gloo_solo_io_v1_sets.IssuedCertificateSet {
	return s.issuedCertificates
}

func (s snapshot) CertificateRequests() certificates_mesh_gloo_solo_io_v1_sets.CertificateRequestSet {
	return s.certificateRequests
}

func (s snapshot) SyncStatusesMultiCluster(ctx context.Context, mcClient multicluster.Client, opts SyncStatusOptions) error {
	var errs error

	if opts.IssuedCertificate {
		for _, obj := range s.IssuedCertificates().List() {
			clusterClient, err := mcClient.Cluster(obj.ClusterName)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
			if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	if opts.CertificateRequest {
		for _, obj := range s.CertificateRequests().List() {
			clusterClient, err := mcClient.Cluster(obj.ClusterName)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
			if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	return errs
}

func (s snapshot) SyncStatuses(ctx context.Context, c client.Client, opts SyncStatusOptions) error {
	var errs error

	if opts.IssuedCertificate {
		for _, obj := range s.IssuedCertificates().List() {
			if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	if opts.CertificateRequest {
		for _, obj := range s.CertificateRequests().List() {
			if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	return errs
}

func (s snapshot) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	snapshotMap["issuedCertificates"] = s.issuedCertificates.List()
	snapshotMap["certificateRequests"] = s.certificateRequests.List()
	return json.Marshal(snapshotMap)
}

// builds the input snapshot from API Clients.
type Builder interface {
	BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error)
}

// Options for building a snapshot
type BuildOptions struct {

	// List options for composing a snapshot from IssuedCertificates
	IssuedCertificates ResourceBuildOptions
	// List options for composing a snapshot from CertificateRequests
	CertificateRequests ResourceBuildOptions
}

// Options for reading resources of a given type
type ResourceBuildOptions struct {

	// List options for composing a snapshot from a resource type
	ListOptions []client.ListOption

	// If provided, ensure the resource has been verified before adding it to snapshots
	Verifier verifier.ServerResourceVerifier
}

// build a snapshot from resources across multiple clusters
type multiClusterBuilder struct {
	clusters multicluster.Interface
	client   multicluster.Client
}

// Produces snapshots of resources across all clusters defined in the ClusterSet
func NewMultiClusterBuilder(
	clusters multicluster.Interface,
	client multicluster.Client,
) Builder {
	return &multiClusterBuilder{
		clusters: clusters,
		client:   client,
	}
}

func (b *multiClusterBuilder) BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error) {

	issuedCertificates := certificates_mesh_gloo_solo_io_v1_sets.NewIssuedCertificateSet()
	certificateRequests := certificates_mesh_gloo_solo_io_v1_sets.NewCertificateRequestSet()

	var errs error

	for _, cluster := range b.clusters.ListClusters() {

		if err := b.insertIssuedCertificatesFromCluster(ctx, cluster, issuedCertificates, opts.IssuedCertificates); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertCertificateRequestsFromCluster(ctx, cluster, certificateRequests, opts.CertificateRequests); err != nil {
			errs = multierror.Append(errs, err)
		}

	}

	outputSnap := NewSnapshot(
		name,

		issuedCertificates,
		certificateRequests,
	)

	return outputSnap, errs
}

func (b *multiClusterBuilder) insertIssuedCertificatesFromCluster(ctx context.Context, cluster string, issuedCertificates certificates_mesh_gloo_solo_io_v1_sets.IssuedCertificateSet, opts ResourceBuildOptions) error {
	issuedCertificateClient, err := certificates_mesh_gloo_solo_io_v1.NewMulticlusterIssuedCertificateClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "certificates.mesh.gloo.solo.io",
			Version: "v1",
			Kind:    "IssuedCertificate",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	issuedCertificateList, err := issuedCertificateClient.ListIssuedCertificate(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range issuedCertificateList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		issuedCertificates.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertCertificateRequestsFromCluster(ctx context.Context, cluster string, certificateRequests certificates_mesh_gloo_solo_io_v1_sets.CertificateRequestSet, opts ResourceBuildOptions) error {
	certificateRequestClient, err := certificates_mesh_gloo_solo_io_v1.NewMulticlusterCertificateRequestClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "certificates.mesh.gloo.solo.io",
			Version: "v1",
			Kind:    "CertificateRequest",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	certificateRequestList, err := certificateRequestClient.ListCertificateRequest(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range certificateRequestList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		certificateRequests.Insert(&item)
	}

	return nil
}

// build a snapshot from resources in a single cluster
type singleClusterBuilder struct {
	mgr         manager.Manager
	clusterName string
}

// Produces snapshots of resources read from the manager for the given cluster
func NewSingleClusterBuilder(
	mgr manager.Manager,
) Builder {
	return NewSingleClusterBuilderWithClusterName(mgr, "")
}

// Produces snapshots of resources read from the manager for the given cluster.
// Snapshot resources will be marked with the given ClusterName.
func NewSingleClusterBuilderWithClusterName(
	mgr manager.Manager,
	clusterName string,
) Builder {
	return &singleClusterBuilder{
		mgr:         mgr,
		clusterName: clusterName,
	}
}

func (b *singleClusterBuilder) BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error) {

	issuedCertificates := certificates_mesh_gloo_solo_io_v1_sets.NewIssuedCertificateSet()
	certificateRequests := certificates_mesh_gloo_solo_io_v1_sets.NewCertificateRequestSet()

	var errs error

	if err := b.insertIssuedCertificates(ctx, issuedCertificates, opts.IssuedCertificates); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertCertificateRequests(ctx, certificateRequests, opts.CertificateRequests); err != nil {
		errs = multierror.Append(errs, err)
	}

	outputSnap := NewSnapshot(
		name,

		issuedCertificates,
		certificateRequests,
	)

	return outputSnap, errs
}

func (b *singleClusterBuilder) insertIssuedCertificates(ctx context.Context, issuedCertificates certificates_mesh_gloo_solo_io_v1_sets.IssuedCertificateSet, opts ResourceBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "certificates.mesh.gloo.solo.io",
			Version: "v1",
			Kind:    "IssuedCertificate",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			"", // verify in the local cluster
			b.mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	issuedCertificateList, err := certificates_mesh_gloo_solo_io_v1.NewIssuedCertificateClient(b.mgr.GetClient()).ListIssuedCertificate(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range issuedCertificateList.Items {
		item := item // pike
		item.ClusterName = b.clusterName
		issuedCertificates.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertCertificateRequests(ctx context.Context, certificateRequests certificates_mesh_gloo_solo_io_v1_sets.CertificateRequestSet, opts ResourceBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "certificates.mesh.gloo.solo.io",
			Version: "v1",
			Kind:    "CertificateRequest",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			"", // verify in the local cluster
			b.mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	certificateRequestList, err := certificates_mesh_gloo_solo_io_v1.NewCertificateRequestClient(b.mgr.GetClient()).ListCertificateRequest(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range certificateRequestList.Items {
		item := item // pike
		item.ClusterName = b.clusterName
		certificateRequests.Insert(&item)
	}

	return nil
}

// build a snapshot from resources in a single cluster
type inMemoryBuilder struct {
	getSnapshot func() (resource.ClusterSnapshot, error)
}

// Produces snapshots of resources read from the manager for the given cluster
func NewInMemoryBuilder(
	getSnapshot func() (resource.ClusterSnapshot, error),
) Builder {
	return &inMemoryBuilder{
		getSnapshot: getSnapshot,
	}
}

func (i *inMemoryBuilder) BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error) {
	genericSnap, err := i.getSnapshot()
	if err != nil {
		return nil, err
	}

	issuedCertificates := certificates_mesh_gloo_solo_io_v1_sets.NewIssuedCertificateSet()
	certificateRequests := certificates_mesh_gloo_solo_io_v1_sets.NewCertificateRequestSet()

	genericSnap.ForEachObject(func(cluster string, gvk schema.GroupVersionKind, obj resource.TypedObject) {
		switch obj := obj.(type) {
		// insert IssuedCertificates
		case *certificates_mesh_gloo_solo_io_v1_types.IssuedCertificate:
			issuedCertificates.Insert(obj)
		// insert CertificateRequests
		case *certificates_mesh_gloo_solo_io_v1_types.CertificateRequest:
			certificateRequests.Insert(obj)
		}
	})

	return NewSnapshot(
		name,

		issuedCertificates,
		certificateRequests,
	), nil
}
