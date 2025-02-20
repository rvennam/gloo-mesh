---
title: "Enterprise"
menuTitle: Enterprise
description: Registering a cluster with Gloo Mesh enterprise edition
weight: 30
---

{{% notice note %}} Gloo Mesh Enterprise is required for this feature. {{% /notice %}}

Relay is an alternative mode of deploying Gloo Mesh that confers several advantages discussed in [this document]({{% versioned_link_path fromRoot="/concepts/relay" %}}).
Cluster registration in relay mode simply consists of installing the relay agent.

This guide will walk you through the basics of registering clusters using the `meshctl` tool or with Helm. We will be using the two cluster contexts mentioned in the Gloo Mesh installation guide, `kind-mgmt-cluster` and `kind-remote-cluster`. Your cluster context names may differ, so please substitute the proper values.

## Register A Cluster

In order to identify a cluster as being managed by Gloo Mesh Enterprise, we have to *register* it in our installation. Registration ensures we are aware of the cluster, and we have properly configured a remote relay *agent* to talk to the local relay *server*. In this example, we will register our remote cluster with Gloo Mesh Enterprise running on the management cluster.

### Register with `meshctl`

We can use the CLI tool `meshctl` to register our remote cluster. The command we use will be `meshctl cluster register enterprise`. This is specific to Gloo Mesh **Enterprise**, and different in nature than the `meshctl cluster register community` command.

To register our remote cluster, there are a few key pieces of information we need:

1. **cluster name** - The name we would like to register the cluster with.
1. **remote-context** - The Kubernetes context with access to the remote cluster being registered.
1. **relay-server-address** - The address of the relay server running on the management cluster.

First, let's get the `relay-server-address`. Assuming you are using Kind and Istio with an ingress provisioned per the [Gloo Mesh Enterprise prerequisites guide]({{% versioned_link_path fromRoot="/setup/enterprise_prerequisites" %}}), you can retrieve the ingress address by running the following commands. If your cluster architecture is different or you exposed the `enterprise-networking` service in a different way, your steps will be different.

```shell
MGMT_CONTEXT=kind-mgmt-cluster # Update value as needed
kubectl config use-context $MGMT_CONTEXT

MGMT_INGRESS_ADDRESS=$(kubectl get node -ojson | jq -r ".items[0].status.addresses[0].address")
MGMT_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].nodePort}')
RELAY_ADDRESS=${MGMT_INGRESS_ADDRESS}:${MGMT_INGRESS_PORT}
```

Let's set variables for the remaining values:

```bash
CLUSTER_NAME=remote-cluster # Update value as needed
REMOTE_CONTEXT=kind-remote-cluster # Update value as needed
```

Now we are ready to register the remote cluster:

```shell
meshctl cluster register enterprise \
  --remote-context=$REMOTE_CONTEXT \
  --relay-server-address $RELAY_ADDRESS \
  $CLUSTER_NAME
```

You should see the following output:

```shell
Registering cluster
📃 Copying root CA relay-root-tls-secret.gloo-mesh to remote cluster from management cluster
📃 Copying bootstrap token relay-identity-token-secret.gloo-mesh to remote cluster from management cluster
💻 Installing relay agent in the remote cluster
Finished installing chart 'enterprise-agent' as release gloo-mesh:enterprise-agent
📃 Creating remote-cluster KubernetesCluster CRD in management cluster
⌚ Waiting for relay agent to have a client certificate
         Checking...
         Checking...
🗑 Removing bootstrap token
✅ Done registering cluster!
```

The `meshctl` command accomplished the following activities:

* Created the `gloo-mesh` namespace
* Copied over the root CA certificate to the remote cluster
* Copied the boostrap token to the remote cluster
* Installed the relay agent in the remote cluster
* Created the KubernetesCluster CRD in the management cluster

When registering a remote cluster using Helm, you will need to run through these tasks yourself. The next section details how to accomplish those tasks and install the relay agent with Helm.

### Register with Helm

You can also register a remote cluster using the Enterprise Agent Helm repository. The same information used for `meshctl` registration will be needed here as well. You will also need to complete the following pre-requisites before running the Helm installation.

* Creating the `gloo-mesh` namespace
* Copying over the self-signed root CA certificate from the management cluster (`relay-root-tls-secret`)
* Copy over the validation token for relay agent initialization (`relay-identity-token-secret`)

If you have not followed these steps, the relay agent deployment will fail.

#### Prerequisites

First create the namespace in the remote cluster:

```shell
CLUSTER_NAME=remote-cluster
REMOTE_CONTEXT=kind-remote-cluster # Update value as needed

kubectl create ns gloo-mesh --context $REMOTE_CONTEXT
```

Now we will get the value of the root CA certificate and create a secret in the remote cluster:

```shell
MGMT_CONTEXT=kind-mgmt-cluster # Update value as needed

kubectl get secret relay-root-tls-secret -n gloo-mesh --context $MGMT_CONTEXT -o jsonpath='{.data.ca\.crt}' | base64 -d > ca.crt

kubectl create secret generic relay-root-tls-secret -n gloo-mesh --context $REMOTE_CONTEXT --from-file ca.crt=ca.crt

rm ca.crt
```

By adding the root CA certificate to the remote cluster, the installation of the relay agent will trust the TLS certificate from the relay server. We also need to copy over the bootstrap token used for initial communication. This token is only used to validate initial communication between the agent and server. Once the gRPC connection is established, the relay server will issue a client certificate to the relay agent to establish an mutually authenticated TLS session.

```shell
kubectl get secret relay-identity-token-secret -n gloo-mesh --context $MGMT_CONTEXT -o jsonpath='{.data.token}' | base64 -d > token

kubectl create secret generic relay-identity-token-secret -n gloo-mesh --context $REMOTE_CONTEXT --from-file token=token

rm token
```

With these tasks accomplished, we are now ready to deploy the relay agent using Helm.

#### Install the Enterprise Agent

We are going to install the Enterprise Agent from the Helm repository located at `https://storage.googleapis.com/gloo-mesh-enterprise/enterprise-agent`.
Make sure to review the Helm values options before installing. Some notable values include:

* `relay.cluster` will be the name by which the cluster is referenced in all Gloo Mesh configuration.
* `relay.serverAddress` is the address by which the Gloo Mesh management plane can be accessed.
* `relay.authority` is the host header that will be passed to the server on the Gloo Mesh management plane.

Also note that the Enterprise Agent's version should match that of the `enterprise-networking` component running on the
management cluster. Run `meshctl version` on the management cluster to review the `enterprise-networking` version.

If you haven't already, you can add the repository by running the following:

```shell
helm repo add gloo-mesh-enterprise https://storage.googleapis.com/gloo-mesh-enterprise/gloo-mesh-enterprise
helm repo update
```

First we will get the ingress address for the relay server. These commands assume you have exposed the relay server through the Istio ingress gateway.

```shell
MGMT_CONTEXT=kind-mgmt-cluster # Update value as needed
kubectl config use-context $MGMT_CONTEXT

MGMT_INGRESS_ADDRESS=$(kubectl get node -ojson | jq -r ".items[0].status.addresses[0].address")
MGMT_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].nodePort}')
RELAY_ADDRESS=${MGMT_INGRESS_ADDRESS}:${MGMT_INGRESS_PORT}
```

Then we will set our variables:

```shell
CLUSTER_NAME=remote-cluster
REMOTE_CONTEXT=kind-remote-cluster # Update value as needed
ENTERPRISE_NETWORKING_VERSION=<current version> # Update based on meshctl version output
```

And now we will deploy the relay agent in the remote cluster.


```bash
helm install enterprise-agent enterprise-agent/enterprise-agent \
  --namespace gloo-mesh \
  --set relay.serverAddress=${RELAY_ADDRESS} \
  --set relay.authority=enterprise-networking.gloo-mesh \
  --set relay.cluster=${CLUSTER_NAME} \
  --kube-context=${REMOTE_CONTEXT} \
  --version ${ENTERPRISE_NETWORKING_VERSION}
```

#### Add a Kubernetes Cluster Object

We've successfully deployed the relay agent in the remote cluster. Now we need to add a `KubernetesCluster` object to the management cluster to make the relay server aware of the remote cluster. The `metadata.name` of the object must match the value passed for `relay.cluster` in the Helm chart above. The `spec.clusterDomain` must match the [local cluster domain](https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/) of the Kubernetes cluster.


```shell
kubectl apply --context $MGMT_CONTEXT -f- <<EOF
apiVersion: multicluster.solo.io/v1alpha1
kind: KubernetesCluster
metadata:
  name: ${CLUSTER_NAME}
  namespace: gloo-mesh
spec:
  clusterDomain: cluster.local
EOF
```

#### Validate the Registration

We can validate the registration process by first checking to make sure the relay agent pod and secrets have been created on the remote cluster:

```shell
kubectl get pods -n gloo-mesh --context $REMOTE_CONTEXT

NAME                                READY   STATUS    RESTARTS   AGE
enterprise-agent-64fc8cc9c5-v7b97   1/1     Running   7          25m

kubectl get secrets -n gloo-mesh --context $REMOTE_CONTEXT

NAME                                     TYPE                                  DATA   AGE
default-token-fcx9w                      kubernetes.io/service-account-token   3      18h
enterprise-agent-token-55mvq             kubernetes.io/service-account-token   3      25m
relay-client-tls-secret                  Opaque                                3      6m24s
relay-identity-token-secret              Opaque                                1      29m
relay-root-tls-secret                    Opaque                                1      18h
sh.helm.release.v1.enterprise-agent.v1   helm.sh/release.v1                    1      25m
```

The `relay-client-tls-secret` secret is the client certificate issued by the relay server. Seeing that entry, we know at the very least communication between the relay agent and server was successful. 

We can also check the logs on the `enterprise-networking` pod on the management cluster for communication from the remote cluster.

```shell
kubectl -n gloo-mesh --context $MGMT_CONTEXT logs deployment/enterprise-networking | grep $CLUSTER_NAME
```

You should see messages similar to:

```shell
{"level":"debug","ts":1616160185.5505846,"logger":"pull-resource-deltas","msg":"recieved request for delta: response_nonce:\"1\"","metadata":{":authority":["enterprise-networking.gloo-mesh.svc.cluster.local:11100"],"content-type":["application/grpc"],"user-agent":["grpc-go/1.34.0"],"x-cluster-id":["remote-cluster"]},"peer":"10.244.0.17:40074"}
```

## Next Steps

And we're done! Any meshes in that cluster will be discovered and available for configuration by Gloo Mesh Enterprise. See the guide on [installing Istio]({{% versioned_link_path fromRoot="/guides/installing_istio" %}}), to see how to easily get Istio running on that cluster.
