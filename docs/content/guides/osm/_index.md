---
title: "Open Service Mesh"
menuTitle: Open Service Mesh
description: How to get started using OSM with Gloo Mesh
weight: 80
---

[Open Service Mesh (OSM)](https://openservicemesh.io/) is the brand new Open Source, Envoy proxy based, service mesh from Microsoft.

OSM introduces a new API set of usage patterns for managing a service mesh. OSM is supported by Gloo Mesh, which can translate, configure, and manage instances of OSM in your environment. 

In this guide, we will walk through the process of installing OSM and a sample application. Then we will use *Access Policies* and *Traffic Policies* from Gloo Mesh to configure the settings in OSM to allow communication between the services in the sample application. The sample application being installed is a variant of the Bookstore application. You can view the topology of the application [here](https://github.com/openservicemesh/osm/blob/main/img/book-thief-app-topology.jpg).

## Before you begin
To illustrate these concepts, we will assume that you have the following:

* Kubernetes cluster for installation or have Docker and Kind to run a local cluster
* The `meshctl` CLI tool for managing Gloo Mesh
* The [OSM installer](https://github.com/openservicemesh/osm/releases)

## Installing OSM

OSM can be installed multiple different ways. The easiest method is using `meshctl`, but it can also be installed directly using the `osm` binary, along with an existing Kubernetes cluster.

{{% notice note %}}
In this guide, we will assume that you are using OSM v0.4.0.
{{% /notice %}}

### Install using `meshctl`

To get up and running with OSM simply run:

```shell script
meshctl demo osm init
``` 

This command will create a local Kind cluster, install OSM, install Gloo Mesh, and deploy a sample application.

### Install manually

If you prefer to install the components yourself, first you will need to install OSM using the `osm` CLI tool.

```shell
osm install
kubectl rollout status deployment --timeout 300s -n osm-system osm-controller
```

After a few moments you should see a successful deployment of OSM.

```shell
Waiting for deployment "osm-controller" rollout to finish: 0 of 1 updated replicas are available...
deployment "osm-controller" successfully rolled out
```

Next we will install Gloo Mesh, as outlined in the Setup guide for Gloo Mesh:

```shell
meshctl install community --register
```

Finally, we will deploy the sample application.

```shell
kubectl create ns bookstore
kubectl create ns bookthief 
kubectl create ns bookwarehouse 
kubectl create ns bookbuyer

osm namespace add bookstore --enable-sidecar-injection
osm namespace add bookthief --enable-sidecar-injection
osm namespace add bookwarehouse --enable-sidecar-injection
osm namespace add bookbuyer --enable-sidecar-injection

kubectl apply -f https://raw.githubusercontent.com/solo-io/gloo-mesh/v0.11.3/ci/osm-demo.yaml

kubectl rollout status deployment --timeout 300s -n bookstore bookstore-v1
kubectl rollout status deployment --timeout 300s -n bookstore bookstore-v2
kubectl rollout status deployment --timeout 300s -n bookthief bookthief
kubectl rollout status deployment --timeout 300s -n bookwarehouse bookwarehouse
kubectl rollout status deployment --timeout 300s -n bookbuyer bookbuyer
```

{{% notice note %}}
For osm version v0.3.0, the flag --enable-sidecar-injection doesn't exist, so you will want to use:
```
osm namespace add bookstore
osm namespace add bookthief
osm namespace add bookwarehouse
osm namespace add bookbuyer
```
{{% /notice %}}

## OSM Basics

The OSM default operating mode is "restrictive" mode. This means that services cannot to talk to each other unless specifically allowed via the API.

The sample application installed is a variant of the Bookstore application with the following components:

* Book buyer - Sends requests to the Bookstore services
* Book thief - Sends requests to the Bookstore services
* Bookstore v1 - Sends requests to the Book warehouse service, receives requests from the Book buyer and thief
* Bookstore v2 - Sends requests to the Book warehouse service, receives requests from the Book buyer and thief
* Book warehouse - Receives requests from the Bookstore service

The restrictive mode of OSM means that **none** of the services can talk to each other yet.

To check the status of our services, we can port-forward to our the bookthief component of the application. 

```shell
kubectl port-forward -n bookthief deploy/bookthief 8080:80
```

Once the port-forward is open navigate to `localhost:8080` in any browser. The numbers should read `0` and `0`
as we have not enabled the bookthief or bookbuyer to interact with the bookstores.

![Bookthief no books]({{% versioned_link_path fromRoot="/img/bookthief-no-books.png" %}})

## Configuring OSM

In order to configure OSM to allow traffic between the various services, and properly split traffic between the two bookstores, we need to apply the following two resources.

First, set the cluster name on which OSM is installed as an environment variable. If you installed OSM using `meshctl demo`, the cluster name would be `mgmt-cluster`.
```shell
OSM_CLUSTER=osm_installation_cluster_name
```

{{< tabs >}}
{{< tab name="YAML file" codelang="yaml">}}
apiVersion: networking.mesh.gloo.solo.io/v1
kind: AccessPolicy
metadata:
  name: osm-access-policy
  namespace: gloo-mesh
spec:
  destinationSelector:
  - kubeServiceRefs:
      services:
      - clusterName: $OSM_CLUSTER
        name: bookstore-v1
        namespace: bookstore
      - clusterName: $OSM_CLUSTER
        name: bookstore-v2
        namespace: bookstore
  sourceSelector:
  - kubeServiceAccountRefs:
      serviceAccounts:
      - clusterName: $OSM_CLUSTER
        name: bookthief
        namespace: bookthief
---
apiVersion: networking.mesh.gloo.solo.io/v1
kind: TrafficPolicy
metadata:
  name: osm-traffic-policy
  namespace: gloo-mesh
spec:
  trafficShift:
    destinations:
    - kubeService:
        clusterName: $OSM_CLUSTER
        name: bookstore-v1
        namespace: bookstore
      weight: 50
    - kubeService:
        clusterName: $OSM_CLUSTER
        name: bookstore-v2
        namespace: bookstore
      weight: 50
  destinationSelector:
  - kubeServiceRefs:
      services:
      - clusterName: $OSM_CLUSTER
        name: bookstore
        namespace: bookstore
{{< /tab >}}
{{< tab name="CLI inline" codelang="shell" >}}
kubectl apply -f - <<EOF
apiVersion: networking.mesh.gloo.solo.io/v1
kind: AccessPolicy
metadata:
  name: osm-access-policy
  namespace: gloo-mesh
spec:
  destinationSelector:
  - kubeServiceRefs:
      services:
      - clusterName: $OSM_CLUSTER
        name: bookstore-v1
        namespace: bookstore
      - clusterName: $OSM_CLUSTER
        name: bookstore-v2
        namespace: bookstore
  sourceSelector:
  - kubeServiceAccountRefs:
      serviceAccounts:
      - clusterName: $OSM_CLUSTER
        name: bookthief
        namespace: bookthief
---
apiVersion: networking.mesh.gloo.solo.io/v1
kind: TrafficPolicy
metadata:
  name: osm-traffic-policy
  namespace: gloo-mesh
spec:
  trafficShift:
    destinations:
    - kubeService:
        clusterName: $OSM_CLUSTER
        name: bookstore-v1
        namespace: bookstore
      weight: 50
    - kubeService:
        clusterName: $OSM_CLUSTER
        name: bookstore-v2
        namespace: bookstore
      weight: 50
  destinationSelector:
  - kubeServiceRefs:
      services:
      - clusterName: $OSM_CLUSTER
        name: bookstore
        namespace: bookstore
EOF
{{< /tab >}}
{{< /tabs >}}

{{% notice note %}}
For osm version v0.3.0, the namespace will have to be changed to `default`.
{{% /notice %}}

The first resource above is a Gloo Mesh `AccessPolicy`. This is the resource which governs access between different services of the application. In this case we are taking advantage of the ability to specify multiple sources and destinations on a rule to allow communication from the bookthief to both bookstore instances. The configuration settings will manifest as `HTTPRouteGroup` and `TrafficTarget` custom resources for OSM.

The second resource is our `TrafficPolicy`. This resource governs how traffic is controlled, and manipulated within our application. This is how we split the traffic between the two separate bookstore instances. The configuration settings will manifest as `TrafficSplits` custom resources for OSM.

Once you have applied the two policies, the bookthief service will start "stealing" books from the bookstore. You can validate this by checking the bookthief page we set up a port-forward to earlier.

![Bookthief success]({{% versioned_link_path fromRoot="/img/bookthief-lotsa-books.png" %}})

## Next Steps

In this guide we installed the Open Service Mesh (OMS) service mesh on a Kubernetes cluster and managed it with Gloo Mesh. From this simple example, we can expand to creating a [Virtual Mesh]({{% versioned_link_path fromRoot="/guides/federate_identity/" %}}) across clusters and [enabling multicluster communications]({{% versioned_link_path fromRoot="/guides/multicluster_communication/" %}}).
