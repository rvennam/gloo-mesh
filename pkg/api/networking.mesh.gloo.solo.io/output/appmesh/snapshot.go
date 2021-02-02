// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./snapshot.go -destination mocks/snapshot.go

// Definitions for Output Snapshots
package appmesh

import (
	"context"
	"encoding/json"
	"sort"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/pkg/multicluster"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/contrib/pkg/output"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appmesh_k8s_aws_v1beta2 "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"
	appmesh_k8s_aws_v1beta2_sets "github.com/solo-io/external-apis/pkg/api/appmesh/appmesh.k8s.aws/v1beta2/sets"
)

// this error can occur if constructing a Partitioned Snapshot from a resource
// that is missing the partition label
var MissingRequiredLabelError = func(labelKey, resourceKind string, obj ezkube.ResourceId) error {
	return eris.Errorf("expected label %v not on labels of %v %v", labelKey, resourceKind, sets.Key(obj))
}

// the snapshot of output resources produced by a translation
type Snapshot interface {

	// return the set of AppmeshK8SAws_V1Beta2_VirtualServices with a given set of labels
	AppmeshK8SAws_V1Beta2_VirtualServices() []LabeledAppmeshK8SAws_V1Beta2_VirtualServiceSet
	// return the set of AppmeshK8SAws_V1Beta2_VirtualNodes with a given set of labels
	AppmeshK8SAws_V1Beta2_VirtualNodes() []LabeledAppmeshK8SAws_V1Beta2_VirtualNodeSet
	// return the set of AppmeshK8SAws_V1Beta2_VirtualRouters with a given set of labels
	AppmeshK8SAws_V1Beta2_VirtualRouters() []LabeledAppmeshK8SAws_V1Beta2_VirtualRouterSet

	// apply the snapshot to the local cluster, garbage collecting stale resources
	ApplyLocalCluster(ctx context.Context, clusterClient client.Client, errHandler output.ErrorHandler)

	// apply resources from the snapshot across multiple clusters, garbage collecting stale resources
	ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, errHandler output.ErrorHandler)

	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)
}

type snapshot struct {
	name string

	appmeshK8SAwsV1Beta2VirtualServices []LabeledAppmeshK8SAws_V1Beta2_VirtualServiceSet
	appmeshK8SAwsV1Beta2VirtualNodes    []LabeledAppmeshK8SAws_V1Beta2_VirtualNodeSet
	appmeshK8SAwsV1Beta2VirtualRouters  []LabeledAppmeshK8SAws_V1Beta2_VirtualRouterSet
	clusters                            []string
}

func NewSnapshot(
	name string,

	appmeshK8SAwsV1Beta2VirtualServices []LabeledAppmeshK8SAws_V1Beta2_VirtualServiceSet,
	appmeshK8SAwsV1Beta2VirtualNodes []LabeledAppmeshK8SAws_V1Beta2_VirtualNodeSet,
	appmeshK8SAwsV1Beta2VirtualRouters []LabeledAppmeshK8SAws_V1Beta2_VirtualRouterSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) Snapshot {
	return &snapshot{
		name: name,

		appmeshK8SAwsV1Beta2VirtualServices: appmeshK8SAwsV1Beta2VirtualServices,
		appmeshK8SAwsV1Beta2VirtualNodes:    appmeshK8SAwsV1Beta2VirtualNodes,
		appmeshK8SAwsV1Beta2VirtualRouters:  appmeshK8SAwsV1Beta2VirtualRouters,
		clusters:                            clusters,
	}
}

// automatically partitions the input resources
// by the presence of the provided label.
func NewLabelPartitionedSnapshot(
	name,
	labelKey string, // the key by which to partition the resources

	appmeshK8SAwsV1Beta2VirtualServices appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet,
	appmeshK8SAwsV1Beta2VirtualNodes appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet,
	appmeshK8SAwsV1Beta2VirtualRouters appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) (Snapshot, error) {

	partitionedAppmeshK8SAws_V1Beta2_VirtualServices, err := partitionAppmeshK8SAws_V1Beta2_VirtualServicesByLabel(labelKey, appmeshK8SAwsV1Beta2VirtualServices)
	if err != nil {
		return nil, err
	}
	partitionedAppmeshK8SAws_V1Beta2_VirtualNodes, err := partitionAppmeshK8SAws_V1Beta2_VirtualNodesByLabel(labelKey, appmeshK8SAwsV1Beta2VirtualNodes)
	if err != nil {
		return nil, err
	}
	partitionedAppmeshK8SAws_V1Beta2_VirtualRouters, err := partitionAppmeshK8SAws_V1Beta2_VirtualRoutersByLabel(labelKey, appmeshK8SAwsV1Beta2VirtualRouters)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		partitionedAppmeshK8SAws_V1Beta2_VirtualServices,
		partitionedAppmeshK8SAws_V1Beta2_VirtualNodes,
		partitionedAppmeshK8SAws_V1Beta2_VirtualRouters,
		clusters...,
	), nil
}

// simplified constructor for a snapshot
// with a single label partition (i.e. all resources share a single set of labels).
func NewSinglePartitionedSnapshot(
	name string,
	snapshotLabels map[string]string, // a single set of labels shared by all resources

	appmeshK8SAwsV1Beta2VirtualServices appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet,
	appmeshK8SAwsV1Beta2VirtualNodes appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet,
	appmeshK8SAwsV1Beta2VirtualRouters appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet,
	clusters ...string, // the set of clusters to apply the snapshot to. only required for multicluster snapshots.
) (Snapshot, error) {

	labeledAppmeshK8SAws_V1Beta2_VirtualService, err := NewLabeledAppmeshK8SAws_V1Beta2_VirtualServiceSet(appmeshK8SAwsV1Beta2VirtualServices, snapshotLabels)
	if err != nil {
		return nil, err
	}
	labeledAppmeshK8SAws_V1Beta2_VirtualNode, err := NewLabeledAppmeshK8SAws_V1Beta2_VirtualNodeSet(appmeshK8SAwsV1Beta2VirtualNodes, snapshotLabels)
	if err != nil {
		return nil, err
	}
	labeledAppmeshK8SAws_V1Beta2_VirtualRouter, err := NewLabeledAppmeshK8SAws_V1Beta2_VirtualRouterSet(appmeshK8SAwsV1Beta2VirtualRouters, snapshotLabels)
	if err != nil {
		return nil, err
	}

	return NewSnapshot(
		name,

		[]LabeledAppmeshK8SAws_V1Beta2_VirtualServiceSet{labeledAppmeshK8SAws_V1Beta2_VirtualService},
		[]LabeledAppmeshK8SAws_V1Beta2_VirtualNodeSet{labeledAppmeshK8SAws_V1Beta2_VirtualNode},
		[]LabeledAppmeshK8SAws_V1Beta2_VirtualRouterSet{labeledAppmeshK8SAws_V1Beta2_VirtualRouter},
		clusters...,
	), nil
}

// apply the desired resources to the cluster state; remove stale resources where necessary
func (s *snapshot) ApplyLocalCluster(ctx context.Context, cli client.Client, errHandler output.ErrorHandler) {
	var genericLists []output.ResourceList

	for _, outputSet := range s.appmeshK8SAwsV1Beta2VirtualServices {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.appmeshK8SAwsV1Beta2VirtualNodes {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.appmeshK8SAwsV1Beta2VirtualRouters {
		genericLists = append(genericLists, outputSet.Generic())
	}

	output.Snapshot{
		Name:        s.name,
		ListsToSync: genericLists,
	}.SyncLocalCluster(ctx, cli, errHandler)
}

// apply the desired resources to multiple cluster states; remove stale resources where necessary
func (s *snapshot) ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, errHandler output.ErrorHandler) {
	var genericLists []output.ResourceList

	for _, outputSet := range s.appmeshK8SAwsV1Beta2VirtualServices {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.appmeshK8SAwsV1Beta2VirtualNodes {
		genericLists = append(genericLists, outputSet.Generic())
	}
	for _, outputSet := range s.appmeshK8SAwsV1Beta2VirtualRouters {
		genericLists = append(genericLists, outputSet.Generic())
	}

	output.Snapshot{
		Name:        s.name,
		Clusters:    s.clusters,
		ListsToSync: genericLists,
	}.SyncMultiCluster(ctx, multiClusterClient, errHandler)
}

func partitionAppmeshK8SAws_V1Beta2_VirtualServicesByLabel(labelKey string, set appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet) ([]LabeledAppmeshK8SAws_V1Beta2_VirtualServiceSet, error) {
	setsByLabel := map[string]appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "AppmeshK8SAws_V1Beta2_VirtualService", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "AppmeshK8SAws_V1Beta2_VirtualService", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = appmesh_k8s_aws_v1beta2_sets.NewVirtualServiceSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedAppmeshK8SAws_V1Beta2_VirtualServices []LabeledAppmeshK8SAws_V1Beta2_VirtualServiceSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledAppmeshK8SAws_V1Beta2_VirtualServiceSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedAppmeshK8SAws_V1Beta2_VirtualServices = append(partitionedAppmeshK8SAws_V1Beta2_VirtualServices, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedAppmeshK8SAws_V1Beta2_VirtualServices, func(i, j int) bool {
		leftLabelValue := partitionedAppmeshK8SAws_V1Beta2_VirtualServices[i].Labels()[labelKey]
		rightLabelValue := partitionedAppmeshK8SAws_V1Beta2_VirtualServices[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedAppmeshK8SAws_V1Beta2_VirtualServices, nil
}

func partitionAppmeshK8SAws_V1Beta2_VirtualNodesByLabel(labelKey string, set appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet) ([]LabeledAppmeshK8SAws_V1Beta2_VirtualNodeSet, error) {
	setsByLabel := map[string]appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "AppmeshK8SAws_V1Beta2_VirtualNode", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "AppmeshK8SAws_V1Beta2_VirtualNode", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = appmesh_k8s_aws_v1beta2_sets.NewVirtualNodeSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedAppmeshK8SAws_V1Beta2_VirtualNodes []LabeledAppmeshK8SAws_V1Beta2_VirtualNodeSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledAppmeshK8SAws_V1Beta2_VirtualNodeSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedAppmeshK8SAws_V1Beta2_VirtualNodes = append(partitionedAppmeshK8SAws_V1Beta2_VirtualNodes, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedAppmeshK8SAws_V1Beta2_VirtualNodes, func(i, j int) bool {
		leftLabelValue := partitionedAppmeshK8SAws_V1Beta2_VirtualNodes[i].Labels()[labelKey]
		rightLabelValue := partitionedAppmeshK8SAws_V1Beta2_VirtualNodes[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedAppmeshK8SAws_V1Beta2_VirtualNodes, nil
}

func partitionAppmeshK8SAws_V1Beta2_VirtualRoutersByLabel(labelKey string, set appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet) ([]LabeledAppmeshK8SAws_V1Beta2_VirtualRouterSet, error) {
	setsByLabel := map[string]appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet{}

	for _, obj := range set.List() {
		if obj.Labels == nil {
			return nil, MissingRequiredLabelError(labelKey, "AppmeshK8SAws_V1Beta2_VirtualRouter", obj)
		}
		labelValue := obj.Labels[labelKey]
		if labelValue == "" {
			return nil, MissingRequiredLabelError(labelKey, "AppmeshK8SAws_V1Beta2_VirtualRouter", obj)
		}

		setForValue, ok := setsByLabel[labelValue]
		if !ok {
			setForValue = appmesh_k8s_aws_v1beta2_sets.NewVirtualRouterSet()
			setsByLabel[labelValue] = setForValue
		}
		setForValue.Insert(obj)
	}

	// partition by label key
	var partitionedAppmeshK8SAws_V1Beta2_VirtualRouters []LabeledAppmeshK8SAws_V1Beta2_VirtualRouterSet

	for labelValue, setForValue := range setsByLabel {
		labels := map[string]string{labelKey: labelValue}

		partitionedSet, err := NewLabeledAppmeshK8SAws_V1Beta2_VirtualRouterSet(setForValue, labels)
		if err != nil {
			return nil, err
		}

		partitionedAppmeshK8SAws_V1Beta2_VirtualRouters = append(partitionedAppmeshK8SAws_V1Beta2_VirtualRouters, partitionedSet)
	}

	// sort for idempotency
	sort.SliceStable(partitionedAppmeshK8SAws_V1Beta2_VirtualRouters, func(i, j int) bool {
		leftLabelValue := partitionedAppmeshK8SAws_V1Beta2_VirtualRouters[i].Labels()[labelKey]
		rightLabelValue := partitionedAppmeshK8SAws_V1Beta2_VirtualRouters[j].Labels()[labelKey]
		return leftLabelValue < rightLabelValue
	})

	return partitionedAppmeshK8SAws_V1Beta2_VirtualRouters, nil
}

func (s snapshot) AppmeshK8SAws_V1Beta2_VirtualServices() []LabeledAppmeshK8SAws_V1Beta2_VirtualServiceSet {
	return s.appmeshK8SAwsV1Beta2VirtualServices
}

func (s snapshot) AppmeshK8SAws_V1Beta2_VirtualNodes() []LabeledAppmeshK8SAws_V1Beta2_VirtualNodeSet {
	return s.appmeshK8SAwsV1Beta2VirtualNodes
}

func (s snapshot) AppmeshK8SAws_V1Beta2_VirtualRouters() []LabeledAppmeshK8SAws_V1Beta2_VirtualRouterSet {
	return s.appmeshK8SAwsV1Beta2VirtualRouters
}

func (s snapshot) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	appmeshK8SAwsV1Beta2VirtualServiceSet := appmesh_k8s_aws_v1beta2_sets.NewVirtualServiceSet()
	for _, set := range s.appmeshK8SAwsV1Beta2VirtualServices {
		appmeshK8SAwsV1Beta2VirtualServiceSet = appmeshK8SAwsV1Beta2VirtualServiceSet.Union(set.Set())
	}
	snapshotMap["appmeshK8SAwsV1Beta2VirtualServices"] = appmeshK8SAwsV1Beta2VirtualServiceSet.List()
	appmeshK8SAwsV1Beta2VirtualNodeSet := appmesh_k8s_aws_v1beta2_sets.NewVirtualNodeSet()
	for _, set := range s.appmeshK8SAwsV1Beta2VirtualNodes {
		appmeshK8SAwsV1Beta2VirtualNodeSet = appmeshK8SAwsV1Beta2VirtualNodeSet.Union(set.Set())
	}
	snapshotMap["appmeshK8SAwsV1Beta2VirtualNodes"] = appmeshK8SAwsV1Beta2VirtualNodeSet.List()
	appmeshK8SAwsV1Beta2VirtualRouterSet := appmesh_k8s_aws_v1beta2_sets.NewVirtualRouterSet()
	for _, set := range s.appmeshK8SAwsV1Beta2VirtualRouters {
		appmeshK8SAwsV1Beta2VirtualRouterSet = appmeshK8SAwsV1Beta2VirtualRouterSet.Union(set.Set())
	}
	snapshotMap["appmeshK8SAwsV1Beta2VirtualRouters"] = appmeshK8SAwsV1Beta2VirtualRouterSet.List()

	snapshotMap["clusters"] = s.clusters

	return json.Marshal(snapshotMap)
}

// LabeledAppmeshK8SAws_V1Beta2_VirtualServiceSet represents a set of appmeshK8SAwsV1Beta2VirtualServices
// which share a common set of labels.
// These labels are used to find diffs between AppmeshK8SAws_V1Beta2_VirtualServiceSets.
type LabeledAppmeshK8SAws_V1Beta2_VirtualServiceSet interface {
	// returns the set of Labels shared by this AppmeshK8SAws_V1Beta2_VirtualServiceSet
	Labels() map[string]string

	// returns the set of VirtualServicees with the given labels
	Set() appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledAppmeshK8SAws_V1Beta2_VirtualServiceSet struct {
	set    appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet
	labels map[string]string
}

func NewLabeledAppmeshK8SAws_V1Beta2_VirtualServiceSet(set appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet, labels map[string]string) (LabeledAppmeshK8SAws_V1Beta2_VirtualServiceSet, error) {
	// validate that each VirtualService contains the labels, else this is not a valid LabeledAppmeshK8SAws_V1Beta2_VirtualServiceSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on AppmeshK8SAws_V1Beta2_VirtualService %v", k, v, item.Name)
			}
		}
	}

	return &labeledAppmeshK8SAws_V1Beta2_VirtualServiceSet{set: set, labels: labels}, nil
}

func (l *labeledAppmeshK8SAws_V1Beta2_VirtualServiceSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledAppmeshK8SAws_V1Beta2_VirtualServiceSet) Set() appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet {
	return l.set
}

func (l labeledAppmeshK8SAws_V1Beta2_VirtualServiceSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list appmesh_k8s_aws_v1beta2.VirtualServiceList
		if err := cli.List(ctx, &list, client.MatchingLabels(l.labels)); err != nil {
			return nil, err
		}
		var items []ezkube.Object
		for _, item := range list.Items {
			item := item // pike
			items = append(items, &item)
		}
		return items, nil
	}

	return output.ResourceList{
		Resources:    desiredResources,
		ListFunc:     listFunc,
		ResourceKind: "VirtualService",
	}
}

// LabeledAppmeshK8SAws_V1Beta2_VirtualNodeSet represents a set of appmeshK8SAwsV1Beta2VirtualNodes
// which share a common set of labels.
// These labels are used to find diffs between AppmeshK8SAws_V1Beta2_VirtualNodeSets.
type LabeledAppmeshK8SAws_V1Beta2_VirtualNodeSet interface {
	// returns the set of Labels shared by this AppmeshK8SAws_V1Beta2_VirtualNodeSet
	Labels() map[string]string

	// returns the set of VirtualNodees with the given labels
	Set() appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledAppmeshK8SAws_V1Beta2_VirtualNodeSet struct {
	set    appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet
	labels map[string]string
}

func NewLabeledAppmeshK8SAws_V1Beta2_VirtualNodeSet(set appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet, labels map[string]string) (LabeledAppmeshK8SAws_V1Beta2_VirtualNodeSet, error) {
	// validate that each VirtualNode contains the labels, else this is not a valid LabeledAppmeshK8SAws_V1Beta2_VirtualNodeSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on AppmeshK8SAws_V1Beta2_VirtualNode %v", k, v, item.Name)
			}
		}
	}

	return &labeledAppmeshK8SAws_V1Beta2_VirtualNodeSet{set: set, labels: labels}, nil
}

func (l *labeledAppmeshK8SAws_V1Beta2_VirtualNodeSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledAppmeshK8SAws_V1Beta2_VirtualNodeSet) Set() appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet {
	return l.set
}

func (l labeledAppmeshK8SAws_V1Beta2_VirtualNodeSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list appmesh_k8s_aws_v1beta2.VirtualNodeList
		if err := cli.List(ctx, &list, client.MatchingLabels(l.labels)); err != nil {
			return nil, err
		}
		var items []ezkube.Object
		for _, item := range list.Items {
			item := item // pike
			items = append(items, &item)
		}
		return items, nil
	}

	return output.ResourceList{
		Resources:    desiredResources,
		ListFunc:     listFunc,
		ResourceKind: "VirtualNode",
	}
}

// LabeledAppmeshK8SAws_V1Beta2_VirtualRouterSet represents a set of appmeshK8SAwsV1Beta2VirtualRouters
// which share a common set of labels.
// These labels are used to find diffs between AppmeshK8SAws_V1Beta2_VirtualRouterSets.
type LabeledAppmeshK8SAws_V1Beta2_VirtualRouterSet interface {
	// returns the set of Labels shared by this AppmeshK8SAws_V1Beta2_VirtualRouterSet
	Labels() map[string]string

	// returns the set of VirtualRouteres with the given labels
	Set() appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet

	// converts the set to a generic format which can be applied by the Snapshot.Apply functions
	Generic() output.ResourceList
}

type labeledAppmeshK8SAws_V1Beta2_VirtualRouterSet struct {
	set    appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet
	labels map[string]string
}

func NewLabeledAppmeshK8SAws_V1Beta2_VirtualRouterSet(set appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet, labels map[string]string) (LabeledAppmeshK8SAws_V1Beta2_VirtualRouterSet, error) {
	// validate that each VirtualRouter contains the labels, else this is not a valid LabeledAppmeshK8SAws_V1Beta2_VirtualRouterSet
	for _, item := range set.List() {
		for k, v := range labels {
			// k=v must be present in the item
			if item.Labels[k] != v {
				return nil, eris.Errorf("internal error: %v=%v missing on AppmeshK8SAws_V1Beta2_VirtualRouter %v", k, v, item.Name)
			}
		}
	}

	return &labeledAppmeshK8SAws_V1Beta2_VirtualRouterSet{set: set, labels: labels}, nil
}

func (l *labeledAppmeshK8SAws_V1Beta2_VirtualRouterSet) Labels() map[string]string {
	return l.labels
}

func (l *labeledAppmeshK8SAws_V1Beta2_VirtualRouterSet) Set() appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet {
	return l.set
}

func (l labeledAppmeshK8SAws_V1Beta2_VirtualRouterSet) Generic() output.ResourceList {
	var desiredResources []ezkube.Object
	for _, desired := range l.set.List() {
		desiredResources = append(desiredResources, desired)
	}

	// enable list func for garbage collection
	listFunc := func(ctx context.Context, cli client.Client) ([]ezkube.Object, error) {
		var list appmesh_k8s_aws_v1beta2.VirtualRouterList
		if err := cli.List(ctx, &list, client.MatchingLabels(l.labels)); err != nil {
			return nil, err
		}
		var items []ezkube.Object
		for _, item := range list.Items {
			item := item // pike
			items = append(items, &item)
		}
		return items, nil
	}

	return output.ResourceList{
		Resources:    desiredResources,
		ListFunc:     listFunc,
		ResourceKind: "VirtualRouter",
	}
}

type builder struct {
	ctx      context.Context
	name     string
	clusters []string

	appmeshK8SAwsV1Beta2VirtualServices appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet
	appmeshK8SAwsV1Beta2VirtualNodes    appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet
	appmeshK8SAwsV1Beta2VirtualRouters  appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet
}

func NewBuilder(ctx context.Context, name string) *builder {
	return &builder{
		ctx:  ctx,
		name: name,

		appmeshK8SAwsV1Beta2VirtualServices: appmesh_k8s_aws_v1beta2_sets.NewVirtualServiceSet(),
		appmeshK8SAwsV1Beta2VirtualNodes:    appmesh_k8s_aws_v1beta2_sets.NewVirtualNodeSet(),
		appmeshK8SAwsV1Beta2VirtualRouters:  appmesh_k8s_aws_v1beta2_sets.NewVirtualRouterSet(),
	}
}

// the output Builder uses a builder pattern to allow
// iteratively collecting outputs before producing a final snapshot
type Builder interface {

	// add AppmeshK8SAws_V1Beta2_VirtualServices to the collected outputs
	AddAppmeshK8SAws_V1Beta2_VirtualServices(appmeshK8SAwsV1Beta2VirtualServices ...*appmesh_k8s_aws_v1beta2.VirtualService)

	// get the collected AppmeshK8SAws_V1Beta2_VirtualServices
	GetAppmeshK8SAws_V1Beta2_VirtualServices() appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet

	// add AppmeshK8SAws_V1Beta2_VirtualNodes to the collected outputs
	AddAppmeshK8SAws_V1Beta2_VirtualNodes(appmeshK8SAwsV1Beta2VirtualNodes ...*appmesh_k8s_aws_v1beta2.VirtualNode)

	// get the collected AppmeshK8SAws_V1Beta2_VirtualNodes
	GetAppmeshK8SAws_V1Beta2_VirtualNodes() appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet

	// add AppmeshK8SAws_V1Beta2_VirtualRouters to the collected outputs
	AddAppmeshK8SAws_V1Beta2_VirtualRouters(appmeshK8SAwsV1Beta2VirtualRouters ...*appmesh_k8s_aws_v1beta2.VirtualRouter)

	// get the collected AppmeshK8SAws_V1Beta2_VirtualRouters
	GetAppmeshK8SAws_V1Beta2_VirtualRouters() appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet

	// build the collected outputs into a label-partitioned snapshot
	BuildLabelPartitionedSnapshot(labelKey string) (Snapshot, error)

	// build the collected outputs into a snapshot with a single partition
	BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (Snapshot, error)

	// add a cluster to the collected clusters.
	// this can be used to collect clusters for use with MultiCluster snapshots.
	AddCluster(cluster string)

	// returns the set of clusters currently stored in this builder
	Clusters() []string

	// merge all the resources from another Builder into this one
	Merge(other Builder)

	// create a clone of this builder (deepcopying all resources)
	Clone() Builder

	// return the difference between the snapshot in this builder's and another
	Delta(newSnap Builder) output.SnapshotDelta
}

func (b *builder) AddAppmeshK8SAws_V1Beta2_VirtualServices(appmeshK8SAwsV1Beta2VirtualServices ...*appmesh_k8s_aws_v1beta2.VirtualService) {
	for _, obj := range appmeshK8SAwsV1Beta2VirtualServices {
		if obj == nil {
			continue
		}
		contextutils.LoggerFrom(b.ctx).Debugf("added output AppmeshK8SAws_V1Beta2_VirtualService %v", sets.Key(obj))
		b.appmeshK8SAwsV1Beta2VirtualServices.Insert(obj)
	}
}
func (b *builder) AddAppmeshK8SAws_V1Beta2_VirtualNodes(appmeshK8SAwsV1Beta2VirtualNodes ...*appmesh_k8s_aws_v1beta2.VirtualNode) {
	for _, obj := range appmeshK8SAwsV1Beta2VirtualNodes {
		if obj == nil {
			continue
		}
		contextutils.LoggerFrom(b.ctx).Debugf("added output AppmeshK8SAws_V1Beta2_VirtualNode %v", sets.Key(obj))
		b.appmeshK8SAwsV1Beta2VirtualNodes.Insert(obj)
	}
}
func (b *builder) AddAppmeshK8SAws_V1Beta2_VirtualRouters(appmeshK8SAwsV1Beta2VirtualRouters ...*appmesh_k8s_aws_v1beta2.VirtualRouter) {
	for _, obj := range appmeshK8SAwsV1Beta2VirtualRouters {
		if obj == nil {
			continue
		}
		contextutils.LoggerFrom(b.ctx).Debugf("added output AppmeshK8SAws_V1Beta2_VirtualRouter %v", sets.Key(obj))
		b.appmeshK8SAwsV1Beta2VirtualRouters.Insert(obj)
	}
}

func (b *builder) GetAppmeshK8SAws_V1Beta2_VirtualServices() appmesh_k8s_aws_v1beta2_sets.VirtualServiceSet {
	return b.appmeshK8SAwsV1Beta2VirtualServices
}
func (b *builder) GetAppmeshK8SAws_V1Beta2_VirtualNodes() appmesh_k8s_aws_v1beta2_sets.VirtualNodeSet {
	return b.appmeshK8SAwsV1Beta2VirtualNodes
}
func (b *builder) GetAppmeshK8SAws_V1Beta2_VirtualRouters() appmesh_k8s_aws_v1beta2_sets.VirtualRouterSet {
	return b.appmeshK8SAwsV1Beta2VirtualRouters
}

func (b *builder) BuildLabelPartitionedSnapshot(labelKey string) (Snapshot, error) {
	return NewLabelPartitionedSnapshot(
		b.name,
		labelKey,

		b.appmeshK8SAwsV1Beta2VirtualServices,
		b.appmeshK8SAwsV1Beta2VirtualNodes,
		b.appmeshK8SAwsV1Beta2VirtualRouters,
		b.clusters...,
	)
}

func (b *builder) BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (Snapshot, error) {
	return NewSinglePartitionedSnapshot(
		b.name,
		snapshotLabels,

		b.appmeshK8SAwsV1Beta2VirtualServices,
		b.appmeshK8SAwsV1Beta2VirtualNodes,
		b.appmeshK8SAwsV1Beta2VirtualRouters,
		b.clusters...,
	)
}

func (b *builder) AddCluster(cluster string) {
	b.clusters = append(b.clusters, cluster)
}

func (b *builder) Clusters() []string {
	return b.clusters
}

func (b *builder) Merge(other Builder) {
	if other == nil {
		return
	}

	b.AddAppmeshK8SAws_V1Beta2_VirtualServices(other.GetAppmeshK8SAws_V1Beta2_VirtualServices().List()...)
	b.AddAppmeshK8SAws_V1Beta2_VirtualNodes(other.GetAppmeshK8SAws_V1Beta2_VirtualNodes().List()...)
	b.AddAppmeshK8SAws_V1Beta2_VirtualRouters(other.GetAppmeshK8SAws_V1Beta2_VirtualRouters().List()...)
	for _, cluster := range other.Clusters() {
		b.AddCluster(cluster)
	}
}

func (b *builder) Clone() Builder {
	if b == nil {
		return nil
	}
	clone := NewBuilder(b.ctx, b.name)

	for _, appmeshK8SAwsV1Beta2VirtualService := range b.GetAppmeshK8SAws_V1Beta2_VirtualServices().List() {
		clone.AddAppmeshK8SAws_V1Beta2_VirtualServices(appmeshK8SAwsV1Beta2VirtualService.DeepCopy())
	}
	for _, appmeshK8SAwsV1Beta2VirtualNode := range b.GetAppmeshK8SAws_V1Beta2_VirtualNodes().List() {
		clone.AddAppmeshK8SAws_V1Beta2_VirtualNodes(appmeshK8SAwsV1Beta2VirtualNode.DeepCopy())
	}
	for _, appmeshK8SAwsV1Beta2VirtualRouter := range b.GetAppmeshK8SAws_V1Beta2_VirtualRouters().List() {
		clone.AddAppmeshK8SAws_V1Beta2_VirtualRouters(appmeshK8SAwsV1Beta2VirtualRouter.DeepCopy())
	}
	for _, cluster := range b.Clusters() {
		clone.AddCluster(cluster)
	}
	return clone
}

func (b *builder) Delta(other Builder) output.SnapshotDelta {
	delta := output.SnapshotDelta{}
	if b == nil {
		return delta
	}

	// calculate delta between VirtualServices
	appmeshK8SAwsV1Beta2VirtualServiceDelta := b.GetAppmeshK8SAws_V1Beta2_VirtualServices().Delta(other.GetAppmeshK8SAws_V1Beta2_VirtualServices())
	appmeshK8SAwsV1Beta2VirtualServiceGvk := schema.GroupVersionKind{
		Group:   "appmesh.k8s.aws",
		Version: "v1beta2",
		Kind:    "VirtualService",
	}
	delta.AddInserted(appmeshK8SAwsV1Beta2VirtualServiceGvk, appmeshK8SAwsV1Beta2VirtualServiceDelta.Inserted)
	delta.AddRemoved(appmeshK8SAwsV1Beta2VirtualServiceGvk, appmeshK8SAwsV1Beta2VirtualServiceDelta.Removed)
	// calculate delta between VirtualNodes
	appmeshK8SAwsV1Beta2VirtualNodeDelta := b.GetAppmeshK8SAws_V1Beta2_VirtualNodes().Delta(other.GetAppmeshK8SAws_V1Beta2_VirtualNodes())
	appmeshK8SAwsV1Beta2VirtualNodeGvk := schema.GroupVersionKind{
		Group:   "appmesh.k8s.aws",
		Version: "v1beta2",
		Kind:    "VirtualNode",
	}
	delta.AddInserted(appmeshK8SAwsV1Beta2VirtualNodeGvk, appmeshK8SAwsV1Beta2VirtualNodeDelta.Inserted)
	delta.AddRemoved(appmeshK8SAwsV1Beta2VirtualNodeGvk, appmeshK8SAwsV1Beta2VirtualNodeDelta.Removed)
	// calculate delta between VirtualRouters
	appmeshK8SAwsV1Beta2VirtualRouterDelta := b.GetAppmeshK8SAws_V1Beta2_VirtualRouters().Delta(other.GetAppmeshK8SAws_V1Beta2_VirtualRouters())
	appmeshK8SAwsV1Beta2VirtualRouterGvk := schema.GroupVersionKind{
		Group:   "appmesh.k8s.aws",
		Version: "v1beta2",
		Kind:    "VirtualRouter",
	}
	delta.AddInserted(appmeshK8SAwsV1Beta2VirtualRouterGvk, appmeshK8SAwsV1Beta2VirtualRouterDelta.Inserted)
	delta.AddRemoved(appmeshK8SAwsV1Beta2VirtualRouterGvk, appmeshK8SAwsV1Beta2VirtualRouterDelta.Removed)
	return delta
}
