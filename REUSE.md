# How to Reuse This Repository

This repository uses a **base + fork** pattern for managing multi-cluster infrastructure at scale.

## Repository Strategy

**Upstream (bootstrap):** https://github.com/openshift-online/bootstrap
- Shared reusable templates in `bases/`
- Hub cluster operators (`operators/{name}/global/`)
- Common patterns everyone benefits from

**Fork (bootstrap-clm):** https://github.com/openshift-online/bootstrap-clm
- Your specific cluster definitions (`regions/`, `clusters/`)
- Instance-specific configurations (`operators/{name}/{cluster}/`)
- Private deployment settings

## Directory Structure

```
bootstrap/
├── bases/                       # Reusable templates (upstream)
│   ├── clusters/ocp/           # OpenShift base templates
│   ├── clusters/eks/           # AWS EKS templates
│   ├── clusters/gcp/           # GCP GKE templates (future)
│   └── pipelines/              # Pipeline templates
├── regions/                     # Define clusters here (fork)
│   └── us-west-2/
│       └── my-cluster/
│           └── region.yaml
├── clusters/                    # Auto-generated instances
├── operators/                   # Operator deployments
│   └── {name}/
│       ├── global/             # Hub cluster instance
│       └── {cluster}/          # Managed cluster instances
├── pipelines/                   # Pipeline deployments
│   └── {name}/
│       ├── global/             # Hub pipelines
│       └── {cluster}/          # Cluster-specific pipelines
└── bin/                         # Management tools
```

## Bootstrap the Hub Cluster

```bash
# Clone and bootstrap
git clone https://github.com/openshift-online/bootstrap.git
cd bootstrap

oc login https://api.your-hub.example.com:6443
oc apply -k operators/openshift-gitops/global
oc apply -k gitops-applications/
```

**Hub installs:**
- OpenShift GitOps (ArgoCD)
- Advanced Cluster Management (ACM)
- OpenShift Pipelines (Tekton)
- Vault (secret management)

## Add Your First Cluster

```bash
./bin/cluster-create

# Prompts:
# - Cluster name: my-cluster
# - Type: ocp or eks
# - Region: us-west-2
# - Instance type: m5.2xlarge

git add .
git commit -m "Add my-cluster"
git push origin main
```

ArgoCD automatically:
- Creates cluster provisioning resources
- Generates operator installations
- Deploys pipelines
- Configures proper sync ordering

## How Bases Work

Each cluster instance references shared bases:

```yaml
# clusters/my-cluster/kustomization.yaml
resources:
  - ../../bases/clusters/ocp/    # Reuses OpenShift base

patches:
  - region: us-west-2             # Adds instance config
    instanceType: m5.2xlarge
```

## Multi-Cloud Support

**Current:**
- `bases/clusters/ocp/` - OpenShift on AWS/Azure/GCP/bare metal (via ACM)
- `bases/clusters/eks/` - AWS EKS

**Future:**
- `bases/clusters/gcp/` - GCP GKE
- `bases/clusters/oci/` - Oracle OKE
- `bases/clusters/aks/` - Azure AKS

## Scaling Pattern

To deploy operators/pipelines across N clusters:

1. Define cluster in `regions/{region}/{name}/region.yaml`
2. Run `./bin/cluster-create` (generates instances)
3. Commit and push
4. ArgoCD deploys everything

**Result:**
```
operators/openshift-pipelines/cluster-01/
operators/openshift-pipelines/cluster-02/
pipelines/hello-world/cluster-01/
pipelines/hello-world/cluster-02/
```

All reference same `bases/`, customized per cluster.

## Monitoring

```bash
# Check applications
oc get applications -n openshift-gitops

# Monitor clusters
oc get clusterdeployments -A     # OpenShift
oc get clusters -A               # EKS

# Health check
./bin/monitor-health
```

## Management Consoles

```bash
# ArgoCD
oc get route openshift-gitops-server -n openshift-gitops

# ACM
oc get route multicloud-console -n open-cluster-management

# Vault
oc get route vault -n vault
```

## Contributing Back

**Fork changes** (your deployments) stay private.

**Base improvements** (templates, operators) can be contributed upstream:
1. Improve `bases/` in your fork
2. PR to upstream bootstrap repo
3. Everyone benefits from shared patterns

## Support

- [QUICKSTART.md](./docs/getting-started/QUICKSTART.md) - Quick start guide
- [ARCHITECTURE.md](./docs/architecture/ARCHITECTURE.md) - Visual diagrams
- [BOOTSTRAP.md](./BOOTSTRAP.md) - Bootstrap details
- [NAVIGATION.md](./NAVIGATION.md) - Repository navigation
