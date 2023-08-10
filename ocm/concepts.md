# OCM
- https://youtu.be/5S5n0-xl2yE
- OCM aims:
    - Cluster Inventory Management
    - Cluster Workload Placement
    - Cluster Workload Distribution
- Hub & Spoke Architecture
    - Pull based model
- Managed Cluster:
    - Klusterlet
        - registration-agent
        - work-agent
    - addon agent
- Hub cluster:
  - Cluster manager
    - registration-controller
    - placement-controller
  - Management addons
    - Policy addon
    - Application addon
    - Addon
- Hub clusters IP address needs to be accessible for managed clusters. Managed clusters can be in private cluster.
- open-cluster-management and open-cluster-management-hub ns in hub.
- Switch to managed cluster:
  - Run `clusteradm join --hub-token <token> --cluster-name <cluster-name>`
- In hub cluster there will be a new ns for managed cluster. Named cluster-ns.
- clustermanager: configuration object for ocm.
- ManifestWork: If you want to create a kubernetes resource in managed cluster, you have to create a custom resource with this kind.
- **ManifestWork**: ManifestWork is used to define a group of Kubernetes resources on the hub to be applied to the managed cluster. In the open-cluster-management project, a `ManifestWork` resource must be created in the cluster namespace. 
- **Status tracking**: Work agent will track all the resources defined in `ManifestWork` and update its status. There are two types of status in manifestwork. The `resourceStatus` tracks the status of each manifest in the `ManifestWork` and `conditions` reflects the overall status of the `ManifestWork`.
- Create clutersets in hub cluster:
  - app1
  - app2
- After that add c1 cluster to app1 clusterset and add c2 cluster to app2.
- kubedb operator on all the clusters:
  - created a Placement object. set clustersets: global.













## Commands

- kind get clusters
- kind create cluster --name=c1
- kind export kubeconfig --name=hub
  - kubectl get nodes
  - helm ls -A
  - clusteradm init --wait: make this cluster HUB.
  - kubectl get clustermanager
  - kuebctl get managedclusters
  - clusteradm get token
  - clusteradm accept --clusters c1
  - 
- clusteradm join --hub-token <token> --cluster-name <cluster-name>
- For kind cluster need to add one flag: clusteradm join --hub-token <token> --cluster-name <cluster-name> --force-internal-endpoint-lookup
- clusteradm get clustersets

### clustersets

- clusteradm create clusterset app1
- clusteradm clusterset set app1 --clusters c1

- clusteradm create clusterset app2
- clusteradm clusterset set app2 --clusters c2

- kubectl create ns app1
- clusteradm clusterset bind app1 --namespace app1

- kubectl create ns app2
- clusteradm clusterset bind app2 --namespace app2

- kuebctl create ns kubedb
- kubectl create ns kubeops
- clusteradm clusterset bind global --namespace kubedb
- clusteradm clusterset bind global --namespace kubeops
- clusteradm get clustersets

- kubectl apply -f kubeops/
- kubectl get placementdecisions -A
- kubectl get manifestwork -A

- To add manifestwork enable in cluster-manager yaml
  - kubectl edit ClusterManager cluster-manager
- kubectl get helmrepositories -A
- kubectl get helmrelease -A
- kubectl get licensestatues



## Examples from Doc


- ManagedClusterSet:
```
$ clusteradm create clusterset example-clusterset
$ clusteradm get clustersets
<ManagedClusterSet>
└── <default>
│   ├── <BoundNamespace>
│   ├── <Status> No ManagedCluster selected
└── <example-clusterset>
│   ├── <BoundNamespace>
│   ├── <Status> No ManagedCluster selected
└── <global>
    └── <BoundNamespace>
    └── <Status> 1 ManagedClusters selected
```

```
$ clusteradm clusterset set example-clusterset --clusters managed1
$ clusteradm get clustersets
<ManagedClusterSet>
└── <default>
│   ├── <BoundNamespace>
│   ├── <Status> No ManagedCluster selected
└── <example-clusterset>
│   ├── <BoundNamespace>
│   ├── <Status> 1 ManagedClusters selected
└── <global>
    └── <BoundNamespace>
    └── <Status> 1 ManagedClusters selected
```

- Binding the ManagedClusterSet to a workspace namespace:
```
$ clusteradm clusterset bind example-clusterset --namespace default
$ clusteradm get clustersets
<ManagedClusterSet>
└── <default>
│   ├── <BoundNamespace>
│   ├── <Status> No ManagedCluster selected
└── <example-clusterset>
│   ├── <Status> 1 ManagedClusters selected
│   ├── <BoundNamespace> default
└── <global>
    └── <BoundNamespace>
    └── <Status> 1 ManagedClusters selected
```