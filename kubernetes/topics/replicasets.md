# ReplicaSets

- A ReplicaSet acts as a cluster-wide Pod manager, ensuring that the right types and numbers of Pods are running
  at all times.
- Pods managed by ReplicaSets are automatically rescheduled under certain failure conditions, such as node failures and
  network partitions.
- ReplicaSet combines a cookie cutter and a desired number of cookies into a single API object. When we define a
  ReplicaSet, we define a specification for the Pods we want to create (the "cookie cutter") and a desired number
  of replicas. Additionally, we need to define a way of finding Pods that the ReplicaSet should control.
- ```Reconciliation Loop``` - The central concept behind a reconciliation loop is the notion of desire state versus
  observed or current state.
- The reconciliation loop is constantly running, observing the current state of the world and taking action to try
  to make the observed state match the desired state.
- Reconciliation loop is an inherently goal-driven, self-healing system.
- ReplicaSets and Pods is loosely coupled. Though ReplicaSets create and manage Pods, they do not own the pods they
  create.
- This notion of “coming in the front door” is another central design concept in Kubernetes. 
  In a similar decoupling, ReplicaSets that create multiple Pods and the services
  that load balance to those Pods are also totally separate, decoupled API objects.
- When the number of Pods in the current state is less than the number of Pods in the desired state, the ReplicaSet
  controller will create new Pods using a template contained in the ReplicaSet specification.
- ReplicaSet monitor cluster state using a set of Pod labels to filter Pod listing and track Pods running within a
  cluster. When initially created, a replicaSet fetches a Pod listing from the Kubernetes API and filters the results
  by labels.
- Most applications should be scalable and all must be fault tolerant. Pods do not provide those features, ReplicaSets do.
- The apiVersion, kind, and metadata fields are mandatory with all Kubernetes objects. ReplicaSet is no exception, i.e., it is also a Kubernetes object.
- The bad news is that ReplicaSets are rarely used independently. You will almost never create a ReplicaSet directly just as you’re not going to create Pods.
  Instead, we tend to create ReplicaSets through Deployments. In other words, we use ReplicaSets to create and control Pods, and Deployments to create ReplicaSets (and a few other things).



## Disadvantage of ReplicaSet

- When a server misbehaves, Pod-level health checks will automatically restart the Pod. But if your health checks
  are incomplete, a Pod can be misbehaving but still be part of the replicated set. In these situations, while it would
  work to simply kill the Pod, that would leave your developers with **only logs to debug the problem**.


## Design ReplicaSet

- ReplicaSets are designed to represent a single, scalable microservice inside your architecture.
- Their key characteristic is that every Pod the ReplicaSet controller creates is **entirely homogenous**.
- ReplicaSets are designed for ```stateless (or nearly stateless) services```.

## ReplicaSet Spec

- All ReplicaSets must have a unique name (defined using the metadata.name field).
- A ```spec``` section that describe the number of Pods(replicas) that should be running cluster-wide at any given time.
- A Pod template that describes the Pod to be created when the defined number of replicas is not met.

## Find ReplicaSet from a Pod

- Sometimes you may wonder if a Pod is being managed by a ReplicaSet, and if it
  is, which one. To enable this kind of discovery, the ReplicaSet controller adds an
  ownerReferences section to every Pod that it creates. If you run the following, look
  for the ownerReferences section:
```
kubectl get pods <pod-name> -o=jsonoath='{.metadata.ownerReferences[0].name}'
```
This will list the name of the ReplicaSet that is managing this Pod.

## Autoscaling

- Kubernetes makes a distinction between ```horizontal scaling``` which involves creating additional replicas of a 
  Pod, and ```vertical scaling``` which involves increasing the resources required for a particular Pod (such as
  increasing the CPU required for the Pod).
- Autoscaling requires the presence of the **metric-server** in your cluster. The metric-server keeps track of metrics
  and provides an API for consuming metrics that HPA (Horizontal Pod Autoscaling) uses when making scaling decisions.
- CPU based autoscaling is most useful for request-based systems that consume CPU proportionally to the number of
  requests they are receiving, while using a relatively static amount of memory.



## Commands

- Create ReplicaSet using YAML file.
```
kubectl apply -f <yaml-file>
```
- Inspect a ReplicaSet.
```
kubectl describe <replicaSet-name>
```
- Find the Pods managed by a ReplicaSet using selector.
```
kubectl get pods -l <labels-name>
```
- Scaling ReplicaSets (imperative scale command).
```
kubectl scale replicasets <replicaSet-name> --replicas=<replica-number>
```
- To scale a ReplicaSet
```
kubectl autoscale <replicaSet-name> --min=2 --max=5 --cpu-percent=80
```
This command creates an auto-scaler that scales between two and five replicas with a CPU threshold of 80%.

- To view, modify or delete resource.
```
kubectl get hpa
```
- Delete ReplicaSets.
```
kubectl delete replicaset <replicaSet-name>
```

- If you don't want to delete the Pods that the ReplicaSet is managing, you can set the ```--cascade flag``` to
  false to ensure only the ReplicaSet object is deleted and not the Pods.
```
kubectl delete replicaset <replicaSet-name> --cascade=false
```

- To reuse the pods.
```
kubectl create -f <yaml-file> --save-config
```

## Selector

- We use it to select which pods should be included in the ReplicaSet. It does not distinguish between the Pods created 
  by a ReplicaSet or some other process. In other words, ReplicaSets and Pods are decoupled. If Pods that match the selector exist, 
  ReplicaSet will do nothing. If they don’t, it will create as many Pods to match the value of the replicas field. Not only that 
  ReplicaSet creates the Pods that are missing, but it also monitors the cluster and ensures that the desired number of replicas 
  is (almost) always running. In case there are already more running Pods with the matching selector, some will be terminated to match the number set in replicas.
