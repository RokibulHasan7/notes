# Labels

- Labels are key/value pairs that can be attached to Kubernetes objects such as Pods and ReplicaSets.
- Labels provide the foundation for grouping object.
- Labels are meant for querying, filtering or otherwise differentiating Pods from each other.
- Labels provide **identifying metadata for objects**.
- Labels values are strings with a maximum length of 63 characters.


## Commands

- Apply or update labels on objects after create.
```
kubectl label deployments <deployment-name> "<label-key>=<label-value>"
```

- To see a label value as a column.
```
kubectl get deployments -L <label-key>
```
- Remove a label.
```
kubectl label deployments <deployment-name> "<label-name>-"
```

## Label Selectors

- Label selectors are used to filter Kubernetes objects based on a set of labels.
- To see the list of Pods that have a specific label.
```
kubectl get pods --selector="<label-name>"
```
We can use two selectors separated by a comma.
```
kubectl get pods --selector-"<label-key=label-value>,<label-key=label-value>"
```
- To see the list with a specific label.
```
kubectl get deployments --selector="label-name"
```
You can also see the list of deployments where that specific label not used.
```
kubectl get deployments --selector="!<label-name>"
```

# Annotations

- Annotations provide a storage mechanism that resembles labels: key/value pairs designed to hold non-identifying
  information that tools and libraries can leverage.
- Annotations provide a place to store additional metadata for Kubernetes objects where the sole purpose of the 
  metadata is assisting tools and libraries.
- Annotations can be used for the tool itself or to pass configuration information between external systems.
- Annotations are used to provide extra information about where an object came from, how to use it, or policy around
  that object.
- Annotations are used in various places in Kubernetes, with the primary use case being rolling deployments.
- During rolling deployments, annotations are used to track rollout status and provide the necessary information
  required to roll back a deployment to a previous state.
- Annotations are good for small bits of data that are highly associated with a specific resource.
- 