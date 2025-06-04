# ACME

"ACM Everywhere" is a base platform of best practices, best-of-breed components, and common patterns using
[Advanced Cluster Management's](https://www.redhat.com/en/technologies/management/advanced-cluster-management)
declarative resource and application management technologies.

Modern service architectures will span clusters and clouds with many different applications deployed across
environments that must be built, tested, progressively delivered, and supported by live SRE staff. ACMEverywhere helps managed this complexity
with declarative CRUD (Clusters, Resources, Users, and Deployments) that is infrastructure-as-code and delivered through pipelines.

## Fully Declarative

A live service represents both platform and business logic components. In fact, you have to bootstrap a *lot* of your
base platform before you can write any business logic that provides value. For example, a service would need, at a minimum, 
a code repository and a build system, an image repository, some fleet of clusters to deploy to complete with peering and properly secure networking, 
an Identity Provider to manage access to your clusters, something to orchestrate your deployments throughout your fleet, Vault to manage secrets,
and more. You haven't gotten to any business logic yet.

### CRUD: Clusters, Resources, Users, Deployments

#### Clusters

CAPI provides the building blocks for declaring a fleet 

`Cluster` - top-level Kind that declares your cluster, its network config, and its control plane and worker nodes.

TODO: Research and Replace with OpenShift/HCP equivalents -- `AWSCluster`, `KubeadmControlPlane`, `MachineDeployment`, `AWSMachineTemplate`. See [example](generated-examples/capi-int-cluster-example.yaml).

`Peering` -- describes relationships and networking between clusters, implemented using such things as VPCs and PrivateLinks and the equivalents across cloud providers.

TODO: Research and define the schema for Peering (app-interface to start and/or community offering)

`CloudAccount` -- cloud accounts own resources, including clusters and cloud resources. 

`ClusterAuthentication` -- a cluster's configured IDP and auth solution

#### Resources

<< Use Radius (or similar plugin architecture) to declare cloud resources (e.g, Postgres in RDS/Aurora/Pod) >>

`CloudResource` -- a resource provided by a hyperscalar. Resources will follow plugins/interfaces/recipes so that a postgres
database is seamlessly provided by pods, AWS RDS instances, Azure Aurora, or any other flavor desired and implemented.

CloudResources are deployed to specific Namespaces on Clusters 

`CloudAccount` -- cloud accounts own resources, including clusters and cloud resources.

#### Users

`IdentityProvider` -- an IDP secures access to all clusters, resources, and deployments.

`User` -- Users are internal developers, engineers, and other SRE staff who managed the clusters, resources, and deployments.

`Role`, `Permission`, `RoleBinding` -- full RBAC model to authorize Users to perform specific actions across clusters and deployments.


#### Deployments

`ImageRegistry` -- an image registry to host container images, such as Quay, ACR, and ECR.

`Build` --  the build system for service development, such as Konflux.

`Repository` -- the repository for project source code, such as GitHub or GitLab.

`Component` -- a specific piece of functionality, such as an individual microservice.

`Application` -- a collection of components create an Application.

`ComponentImageRegistry` -- a image registry configured for a specific component.

`ComponentBuild` -- a build configured for a specific component.

`ComponentRepository` -- a code repository configured for a specific component.


## Entity Relationship Diagram

Draft (at best).  Deployment section TDB based on Argo Rollouts.

![ERD](the%20big%20erd.drawio.png )







### Users

<< Explain IAM strategy.  Keycloak, roles, permissions, etc. >>

### Deployments

<< Explain progressive delivery of deployments throughout the fleet with tests, soak time, and metrics >>

