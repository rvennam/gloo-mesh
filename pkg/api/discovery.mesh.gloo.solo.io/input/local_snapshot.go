// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./local_snapshot.go -destination mocks/local_snapshot.go

// The Input SettingsSnapshot contains the set of all:
// * Settings
// read from a given cluster or set of clusters, across all namespaces.
//
// A snapshot can be constructed from either a single Manager (for a single cluster)
// or a ClusterWatcher (for multiple clusters) using the SettingsSnapshotBuilder.
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

	settings_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1"
	settings_mesh_gloo_solo_io_v1_types "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1"
	settings_mesh_gloo_solo_io_v1_sets "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1/sets"
)

// SnapshotGVKs is a list of the GVKs included in this snapshot
var SettingsSnapshotGVKs = []schema.GroupVersionKind{

	schema.GroupVersionKind{
		Group:   "settings.mesh.gloo.solo.io",
		Version: "v1",
		Kind:    "Settings",
	},
}

// the snapshot of input resources consumed by translation
type SettingsSnapshot interface {

	// return the set of input Settings
	Settings() settings_mesh_gloo_solo_io_v1_sets.SettingsSet
	// update the status of all input objects which support
	// the Status subresource (across multiple clusters)
	SyncStatusesMultiCluster(ctx context.Context, mcClient multicluster.Client, opts SettingsSyncStatusOptions) error
	// update the status of all input objects which support
	// the Status subresource (in the local cluster)
	SyncStatuses(ctx context.Context, c client.Client, opts SettingsSyncStatusOptions) error
	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)
}

// options for syncing input object statuses
type SettingsSyncStatusOptions struct {

	// sync status of Settings objects
	Settings bool
}

type snapshotSettings struct {
	name string

	settings settings_mesh_gloo_solo_io_v1_sets.SettingsSet
}

func NewSettingsSnapshot(
	name string,

	settings settings_mesh_gloo_solo_io_v1_sets.SettingsSet,

) SettingsSnapshot {
	return &snapshotSettings{
		name: name,

		settings: settings,
	}
}

func NewSettingsSnapshotFromGeneric(
	name string,
	genericSnapshot resource.ClusterSnapshot,
) SettingsSnapshot {

	settingsSet := settings_mesh_gloo_solo_io_v1_sets.NewSettingsSet()

	for _, snapshot := range genericSnapshot {

		settings := snapshot[schema.GroupVersionKind{
			Group:   "settings.mesh.gloo.solo.io",
			Version: "v1",
			Kind:    "Settings",
		}]

		for _, settings := range settings {
			settingsSet.Insert(settings.(*settings_mesh_gloo_solo_io_v1_types.Settings))
		}

	}
	return NewSettingsSnapshot(
		name,
		settingsSet,
	)
}

func (s snapshotSettings) Settings() settings_mesh_gloo_solo_io_v1_sets.SettingsSet {
	return s.settings
}

func (s snapshotSettings) SyncStatusesMultiCluster(ctx context.Context, mcClient multicluster.Client, opts SettingsSyncStatusOptions) error {
	var errs error

	if opts.Settings {
		for _, obj := range s.Settings().List() {
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

func (s snapshotSettings) SyncStatuses(ctx context.Context, c client.Client, opts SettingsSyncStatusOptions) error {
	var errs error

	if opts.Settings {
		for _, obj := range s.Settings().List() {
			if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	return errs
}

func (s snapshotSettings) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	snapshotMap["settings"] = s.settings.List()
	return json.Marshal(snapshotMap)
}

// builds the input snapshot from API Clients.
type SettingsBuilder interface {
	BuildSnapshot(ctx context.Context, name string, opts SettingsBuildOptions) (SettingsSnapshot, error)
}

// Options for building a snapshot
type SettingsBuildOptions struct {

	// List options for composing a snapshot from Settings
	Settings ResourceSettingsBuildOptions
}

// Options for reading resources of a given type
type ResourceSettingsBuildOptions struct {

	// List options for composing a snapshot from a resource type
	ListOptions []client.ListOption

	// If provided, ensure the resource has been verified before adding it to snapshots
	Verifier verifier.ServerResourceVerifier
}

// build a snapshot from resources across multiple clusters
type multiClusterSettingsBuilder struct {
	clusters multicluster.Interface
	client   multicluster.Client
}

// Produces snapshots of resources across all clusters defined in the ClusterSet
func NewMultiClusterSettingsBuilder(
	clusters multicluster.Interface,
	client multicluster.Client,
) SettingsBuilder {
	return &multiClusterSettingsBuilder{
		clusters: clusters,
		client:   client,
	}
}

func (b *multiClusterSettingsBuilder) BuildSnapshot(ctx context.Context, name string, opts SettingsBuildOptions) (SettingsSnapshot, error) {

	settings := settings_mesh_gloo_solo_io_v1_sets.NewSettingsSet()

	var errs error

	for _, cluster := range b.clusters.ListClusters() {

		if err := b.insertSettingsFromCluster(ctx, cluster, settings, opts.Settings); err != nil {
			errs = multierror.Append(errs, err)
		}

	}

	outputSnap := NewSettingsSnapshot(
		name,

		settings,
	)

	return outputSnap, errs
}

func (b *multiClusterSettingsBuilder) insertSettingsFromCluster(ctx context.Context, cluster string, settings settings_mesh_gloo_solo_io_v1_sets.SettingsSet, opts ResourceSettingsBuildOptions) error {
	settingsClient, err := settings_mesh_gloo_solo_io_v1.NewMulticlusterSettingsClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "settings.mesh.gloo.solo.io",
			Version: "v1",
			Kind:    "Settings",
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

	settingsList, err := settingsClient.ListSettings(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range settingsList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		settings.Insert(&item)
	}

	return nil
}

// build a snapshot from resources in a single cluster
type singleClusterSettingsBuilder struct {
	mgr         manager.Manager
	clusterName string
}

// Produces snapshots of resources read from the manager for the given cluster
func NewSingleClusterSettingsBuilder(
	mgr manager.Manager,
) SettingsBuilder {
	return NewSingleClusterSettingsBuilderWithClusterName(mgr, "")
}

// Produces snapshots of resources read from the manager for the given cluster.
// Snapshot resources will be marked with the given ClusterName.
func NewSingleClusterSettingsBuilderWithClusterName(
	mgr manager.Manager,
	clusterName string,
) SettingsBuilder {
	return &singleClusterSettingsBuilder{
		mgr:         mgr,
		clusterName: clusterName,
	}
}

func (b *singleClusterSettingsBuilder) BuildSnapshot(ctx context.Context, name string, opts SettingsBuildOptions) (SettingsSnapshot, error) {

	settings := settings_mesh_gloo_solo_io_v1_sets.NewSettingsSet()

	var errs error

	if err := b.insertSettings(ctx, settings, opts.Settings); err != nil {
		errs = multierror.Append(errs, err)
	}

	outputSnap := NewSettingsSnapshot(
		name,

		settings,
	)

	return outputSnap, errs
}

func (b *singleClusterSettingsBuilder) insertSettings(ctx context.Context, settings settings_mesh_gloo_solo_io_v1_sets.SettingsSet, opts ResourceSettingsBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "settings.mesh.gloo.solo.io",
			Version: "v1",
			Kind:    "Settings",
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

	settingsList, err := settings_mesh_gloo_solo_io_v1.NewSettingsClient(b.mgr.GetClient()).ListSettings(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range settingsList.Items {
		item := item // pike
		item.ClusterName = b.clusterName
		settings.Insert(&item)
	}

	return nil
}

// build a snapshot from resources in a single cluster
type inMemorySettingsBuilder struct {
	getSnapshot func() (resource.ClusterSnapshot, error)
}

// Produces snapshots of resources read from the manager for the given cluster
func NewInMemorySettingsBuilder(
	getSnapshot func() (resource.ClusterSnapshot, error),
) SettingsBuilder {
	return &inMemorySettingsBuilder{
		getSnapshot: getSnapshot,
	}
}

func (i *inMemorySettingsBuilder) BuildSnapshot(ctx context.Context, name string, opts SettingsBuildOptions) (SettingsSnapshot, error) {
	genericSnap, err := i.getSnapshot()
	if err != nil {
		return nil, err
	}

	settings := settings_mesh_gloo_solo_io_v1_sets.NewSettingsSet()

	genericSnap.ForEachObject(func(cluster string, gvk schema.GroupVersionKind, obj resource.TypedObject) {
		switch obj := obj.(type) {
		// insert Settings
		case *settings_mesh_gloo_solo_io_v1_types.Settings:
			settings.Insert(obj)
		}
	})

	return NewSettingsSnapshot(
		name,

		settings,
	), nil
}
