# Bootstrap

This repository contains the scripting, configuration, GitOps content, and documentation necessary to bootstrap a region of a cloud service for Red Hat, based on OpenShift and the Red Hat products and supported community projects that we leverage in our reference architecture.  

# Install

## OpenShift

An initial OpenShift cluster is required to act as the bootstrap cluster and centralized/global control plane.

Obtaining or provisioning the bootstrap cluster is beyond the scope of this document, but the [example install-config.yaml](./examples/install-config.yaml)
was generated by the OpenShift installer to create an OCP cluster and is used as the basis for testing this project. 

The rest of this document assumes you are `kubeadmin` with a KUBECONFIG.


## Secrets

AWS accounts and image pull secrets are inputs into the provisioning process. They will be obtained by external
processes, stored in predictable vault paths, and delivered to our clusters.

Until there is Vault, we must create the `aws-creds` and `pull-secret` Secrets for each cluster provisioned.

Currently, this project assumes the same credentials and pull secrets for all clusters.

From ACM, retrieve the `aws-creds` and `pull-secret`, store them locally, and apply them to your cluster/namespace when needed.

```
# TODO: get the correct namespace for these secrets
oc get secret aws-creds -n hive-managed-clusters -o yaml > aws-creds.yaml
oc get secret pull-secret -n hive-managed-clusters -o yaml > pull-secret.yaml
```

# Clusters

Clusters are defined in this repository.  `cluster-10` and `cluster-20` are the examples.

The process for adding `cluster-30` is currently manual. Improvements to workflows and tooling around adding new
clusters is expected to continue.

To add `cluster-30`, do the following:

```
1. Copy/paste ./clusters/overlay/region-02 as ./clusters/overlay/region-03
2. Find/Replace 'cluster-20' with 'cluster-30' in the files in ./clusters/overlay/region-03
3. Add the new cluster overlay to ./regional-clusters/kustomization.yaml
4. Add the new cluster to the end of ./bootstrap.sh for observing status
5. Create a Pull Request and submit to the repository.
6. Run ./bootstrap.sh to wait for cluster-30 to be provisioned or watch the ACM console 
```

See https://github.com/openshift-online/bootstrap/pull/48

Argo applies the new cluster as part of the [regional clusters](./gitops-applications/regional-clusters.application.yaml) gitops application.


```mermaid
sequenceDiagram
   actor Admin
   actor Installer
   actor OCP
   actor Argo
   actor Git
   actor Tekton
   
   Admin->>Installer: Provision/Get Cluster
   Installer->>Admin: obtain KubeConfig
   
   Admin->>OCP: run ./bootstrap.sh
   OCP->>Admin: RHCP (ACM, Argo, Pipelines) is installed
      
   Argo->>Git: Get desired state
   Git->>Argo: Regional OCM YAML
   
   Argo->>OCP: Argo applies YAML
   OCP->>Argo: Deployment status
   
   Admin->>OCP: oc get pods 
   OCP->>Admin: Regional OCM is green 
    
```
