# DaemonSets

- A DaemonSet ensures that a copy of a Pod is running across a set of nodes in a Kubernetes cluster.
- DaemonSets are **used to deploy system daemons such as log collectors and monitoring agents**, which typically must run
  on every node.
- DaemonSets share similar functionality with ReplicaSets; both create Pods that are expected to be long-running service
  and ensure that the desired state and the observed state of the cluster match.

### ReplicaSet vs DaemonSet

- ReplicaSet should be used when your application is completely decoupled from the node and you can run multiple copies
  on a given node without special consideration.
- DaemonSets should be used when a single copy of your application must run on all or a subset of the nodes in the cluster.
- If you find yourself wanting a single Pod per node, then a DaemonSet is the correct Kubernetes resource to use.
- if you find yourself building a homogeneous replicated service to serve user traffic, then a ReplicaSet is probably
  the right Kubernetes resource to use.
- ```You should generally not use schedulling restrictions or other parameters to ensure that Pods do not colocate on the same node.```
- You can use DaemonSets to install software on nodes in a cloud-based cluster.

### DaemonSet Scheduler

- A daemonSet will create a copy of a Pod every node unless a node selector is used.
- Pods created by DaemonSets are ignored by the Kubernetes scheduler.
- Like ReplicaSets, DaemonSets are managed by a reconciliation loop that measures the desired state with the
  observed state.

### Updating a DaemonSet

- DaemonSet can be rolled out using the same **RollingUpdate** strategy that Deployments use.
- You can configure the update strategy using the ```spec.updateStrategy.type``` field, which should have the value
  ```RollingUpdate```.
- ```spec.minReadySeconds``` - Determines how long a Pod must be "ready" before the rolling update proceeds to 
  upgrade subsequent Pods.
- ```spec.updateStrategy.rollingUpdate.maxUnavilable``` - indicates how many Pods maybe simultaneously updated by
  the rolling update.
- You will likely want to set ```spec.minReadySeconds``` to a reasonably long value, for example 30-60 seconds, to
  ensure that your Pod is truly healthy before the rollout proceeds.
- The setting for ```spec.updateStrategy.rollingUpdate.maxUnavailable``` is more likely to be application dependent.
  **Setting it to 1 is a safe.**
- A good approach might be to set **maxUnavailable** to 1 and only increase it if users or administrators complain about
  DaemonSet rollout speed.
- 

## Commands

- Create a DaemonSet from YAML file.
```
kubectl apply -f <yaml-file>
```

- To see where a specific Pod is assigned.
```
kubectl get pods -l <label> -o wide
```

- See Rollout status.
```
kubectl rollout status daemonSets <daemonset-name>
```

- Delete a DaemonSet.
```
kubectl delete -f <daemonSet-name>
```

