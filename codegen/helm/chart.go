package helm

import (
	"os"
	"strings"

	"github.com/solo-io/gloo-mesh/codegen/io"
	"github.com/solo-io/gloo-mesh/pkg/common/defaults"
	"github.com/solo-io/gloo-mesh/pkg/common/version"
	"github.com/solo-io/skv2/codegen/model"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// this file provides the source of truth for the Gloo Mesh Helm chart.
// it is imported into the root level generate.go to generate the Portal manifest
var (
	registry = func() string {
		registry := os.Getenv("IMAGE_REGISTRY")
		if registry == "" {
			registry = defaultRegistry
		}
		return registry
	}()
	// built-in skv2 templates we don't need
	filterTemplates = func(outPath string) bool {
		return outPath == "templates/namespace.yaml" ||
			outPath == "templates/chart.yaml" ||
			outPath == "templates/configmap.yaml"
	}
)

var Chart = &model.Chart{
	Operators: []model.Operator{
		discoveryOperator(),
		NetworkingOperator("networking"),
	},
	FilterTemplate: filterTemplates,
	Data: model.Data{
		ApiVersion:  "v1",
		Name:        "gloo-mesh",
		Description: "Helm chart for Gloo Mesh.",
		Version:     version.Version,
		Dependencies: []model.Dependency{{
			Name:       "gloo-mesh-crds",
			Version:    version.Version,
			Repository: GetHelmRepository() + "gloo-mesh-crds",
		}},
	},
	Values: DefaultValues(),
}

// return the Helm repository where Gloo Mesh charts are stored.
// will point to public url on releases, else install/helm/_output/charts/
func GetHelmRepository() string {
	// WARNING: hard-coupled with makefile convention
	if os.Getenv("RELEASE") == "\"true\"" {
		return "https://storage.googleapis.com/gloo-mesh/"
	}
	// file path must be relative to repo root
	return "file://../"
}

var CrdsChart = &model.Chart{
	FilterTemplate: func(outPath string) bool {
		return strings.Contains(outPath, "templates") || outPath == "values.yaml"
	},
	Data: model.Data{
		ApiVersion:  "v1",
		Name:        "gloo-mesh-crds",
		Description: "CRDs required by Gloo Mesh management controllers (i.e. discovery and networking).",
		Version:     version.Version,
	},
}

var AgentCrdsChart = &model.Chart{
	FilterTemplate: func(outPath string) bool {
		return strings.Contains(outPath, "templates") || outPath == "values.yaml"
	},
	Data: model.Data{
		ApiVersion:  "v1",
		Name:        "agent-crds",
		Description: "CRDs required by Gloo Mesh remote agents (i.e. cert-agent and enterprise-agent).",
		Version:     version.Version,
	},
}

var CertAgentChart = &model.Chart{
	Operators: []model.Operator{
		certAgentOperator(),
	},
	FilterTemplate: filterTemplates,
	Data: model.Data{
		ApiVersion:  "v1",
		Name:        "cert-agent",
		Description: "Helm chart for the Gloo Mesh Certificate Agent.",
		Version:     version.Version,
	},
	Values: ChartValues{
		Verbose: false,
	},
}

var (
	defaultRegistry = "gcr.io/gloo-mesh"
)

func discoveryOperator() model.Operator {
	var rbacPolicies []rbacv1.PolicyRule
	rbacPolicies = append(rbacPolicies, io.ClusterWatcherInputTypes.RbacPoliciesWatch()...)
	rbacPolicies = append(rbacPolicies, io.DiscoveryLocalInputTypes.RbacPoliciesWatch()...)
	rbacPolicies = append(rbacPolicies, io.DiscoveryOutputTypes.Snapshot.RbacPoliciesWrite()...)
	return model.Operator{
		Name: "discovery",
		Deployment: model.Deployment{
			Image: glooMeshImage(),
			Resources: &v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("125m"),
					v1.ResourceMemory: resource.MustParse("256Mi"),
				},
			},
		},
		Service: model.Service{
			Type: v1.ServiceTypeClusterIP,
			Ports: []model.ServicePort{
				{
					Name:        "metrics",
					DefaultPort: int32(defaults.MetricsPort),
				},
			},
		},
		Rbac: rbacPolicies,
		Args: []string{
			"discovery",
			"--metrics-port={{ $.Values.discovery.ports.metrics | default $.Values.defaultMetricsPort }}",
			"--settings-name={{ $.Values.glooMeshOperatorArgs.settingsRef.name }}",
			"--settings-namespace={{ $.Values.glooMeshOperatorArgs.settingsRef.namespace }}",
			"--verbose={{ $.Values.verbose }}",
		},
		Env: []v1.EnvVar{
			{
				Name: defaults.PodNamespaceEnv,
				ValueFrom: &v1.EnvVarSource{
					FieldRef: &v1.ObjectFieldSelector{
						FieldPath: "metadata.namespace",
					},
				},
			},
		},
	}
}

// exported for use in Enterprise chart
func NetworkingOperator(name string) model.Operator {
	var rbacPolicies []rbacv1.PolicyRule
	rbacPolicies = append(rbacPolicies, io.ClusterWatcherInputTypes.RbacPoliciesWatch()...)
	rbacPolicies = append(rbacPolicies, io.NetworkingInputTypes.RbacPoliciesWatch()...)
	rbacPolicies = append(rbacPolicies, io.NetworkingInputTypes.RbacPoliciesUpdateStatus()...)
	rbacPolicies = append(rbacPolicies, io.LocalNetworkingOutputTypes.Snapshot.RbacPoliciesWrite()...)
	rbacPolicies = append(rbacPolicies, io.IstioNetworkingOutputTypes.Snapshot.RbacPoliciesWrite()...)
	rbacPolicies = append(rbacPolicies, io.SmiNetworkingOutputTypes.Snapshot.RbacPoliciesWrite()...)
	rbacPolicies = append(rbacPolicies, io.CertificateIssuerInputTypes.RbacPoliciesWatch()...)
	rbacPolicies = append(rbacPolicies, io.CertificateIssuerInputTypes.RbacPoliciesUpdateStatus()...)
	return model.Operator{
		Name: name,
		Deployment: model.Deployment{
			Image: glooMeshImage(),
			Resources: &v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("125m"),
					v1.ResourceMemory: resource.MustParse("256Mi"),
				},
			},
		},
		Rbac: rbacPolicies,
		Args: []string{
			"networking",
			"--metrics-port={{ $.Values.defaultMetricsPort }}",
			"--settings-name={{ $.Values.glooMeshOperatorArgs.settingsRef.name }}",
			"--settings-namespace={{ $.Values.glooMeshOperatorArgs.settingsRef.namespace }}",
			"--verbose={{ $.Values.verbose }}",
			"--disallow-intersecting-config={{ $.Values.disallowIntersectingConfig }}",
			"--watch-output-types={{ $.Values.watchOutputTypes }}",
		},
		Env: []v1.EnvVar{
			{
				Name: defaults.PodNamespaceEnv,
				ValueFrom: &v1.EnvVarSource{
					FieldRef: &v1.ObjectFieldSelector{
						FieldPath: "metadata.namespace",
					},
				},
			},
		},
	}
}

func certAgentOperator() model.Operator {
	var rbacPolicies []rbacv1.PolicyRule
	rbacPolicies = append(rbacPolicies, io.CertificateAgentInputTypes.RbacPoliciesWatch()...)
	rbacPolicies = append(rbacPolicies, io.CertificateAgentInputTypes.RbacPoliciesUpdateStatus()...)
	rbacPolicies = append(rbacPolicies, io.CertificateAgentOutputTypes.Snapshot.RbacPoliciesWrite()...)
	// ability to bounce pods
	rbacPolicies = append(rbacPolicies, rbacv1.PolicyRule{
		Verbs:     []string{"*"},
		APIGroups: []string{""},
		Resources: []string{"pods"},
	})
	return model.Operator{
		Name: "cert-agent",
		Deployment: model.Deployment{
			Image: certAgentImage(),
			Resources: &v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("50m"),
					v1.ResourceMemory: resource.MustParse("128Mi"),
				},
			},
		},
		Service: model.Service{
			Type: v1.ServiceTypeClusterIP,
			Ports: []model.ServicePort{
				{
					Name:        "metrics",
					DefaultPort: int32(defaults.MetricsPort),
				},
			},
		},
		Rbac: rbacPolicies,
		Args: []string{
			"--metrics-port={{ $.Values.certAgent.ports.metrics | default $.Values.defaultMetricsPort }}",
			"--verbose={{ $.Values.verbose }}",
		},
		Env: []v1.EnvVar{
			{
				Name: defaults.PodNamespaceEnv,
				ValueFrom: &v1.EnvVarSource{
					FieldRef: &v1.ObjectFieldSelector{
						FieldPath: "metadata.namespace",
					},
				},
			},
		},
	}
}

// both glooMesh operators share same image
func glooMeshImage() model.Image {
	return model.Image{
		Registry:   registry,
		Repository: "gloo-mesh",
		Tag:        version.Version,
		PullPolicy: v1.PullIfNotPresent,
	}
}

func certAgentImage() model.Image {
	return model.Image{
		Registry:   registry,
		Repository: "cert-agent",
		Tag:        version.Version,
		PullPolicy: v1.PullIfNotPresent,
	}
}
