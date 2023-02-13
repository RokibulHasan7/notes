# Deployments

- The Deployment object exists to manage the release of new versions.
- The actual mechanics of the software rollout performed by a Deployment are controlled by a Deployment controller that 
  runs in the Kubernetes cluster itself.
- Scaling the Deployments has also scaled the ReplicaSet it controls.
- if you have a current Deployment in progress, you can use ```kubectl rollout status``` to obtain the current status
  of that rollout.

### Deployment Strategies

- There are two different rollout strategies, ```Recreate``` and ```RollingUpdate```.
- Recreate - It simply updates the ReplicaSet it manages to use the new image and terminates all of the Pods associated
  with the deployment. The ReplicaSet notices that it no longer has any replicas and re-creates all Pods using the new
  image. Once the Pods are re-created, they are running the new version.
- Recreate is fast and simple, it will result in workload downtime. For this **Recreate strategy should be used only 
  for the Deployments where a service downtime is acceptable.
- RollingUpdate - The RollingUpdate strategy is the generally preferable strategy for any user-facing
  service. While it is slower than Recreate, it is also significantly more sophisticated
  and robust. Using RollingUpdate, you can roll out a new version of your service
  while it is still receiving user traffic, without any downtime.

### Configuring a rolling update

- There are two parameters you can use to tune the rolling update behavior: ```maxUnavailable``` and ```maxSurge```.
- The **maxUnavailable** parameter sets the maximum number of Pods that can be unavailable during a rolling update.
- Using a percentage is a good approach for most services.
- The **maxSurge** parameter controls how many extra resources can be created to achieve a rollout.
- Setting **minReadySeconds** to 60 indicates that the Deployment must wait for 60
  seconds after seeing a Pod become healthy before moving on to updating the next
  Pod.


## Commands

- Create Deployment from YAML file.
```
kubectl create -f <yaml-file>
```

- Download deployment into a YAML file.
```
kubectl get deploymets <deployment-name> -o yaml > <deployment-yaml-file>
kubectl replace -f <deployment-yaml-file> --save-config
```

- See the label selector by looking at the deployment object.
```
kubectl get deployments <deployment-name> \
-o jsonpath --template {.spec.selector.matchLabels}
```

- Find specific ReplicaSet.
```
kubectl get replicasets --selector=<label-key=lebel-value>
```

- Scaling deployments.
```
kubectl scale deployments <deployment-name> --replicas=X
```

- Check rollout status.
```
kubectl rollout status deployments <deployment-name>
```

- See the old and new ReplicaSet managed by the deployment.
```
kubectl get replicasets -o wide
```

- Pause a rollout.
```
kubectl rollout pause deployments <deloyment-name>
```

- Resume a rollout.
```
kubectl rollout resume deployments <deployment-name>
```

- Rollout history.
```
kubectl rollout history deployment <deployment-name>
```

- See more details about a particular revision.
```
kubectl rollout history deployment <deployment-name> --revision=<rollout-id-number>
```

- Rollback to the last update.
```
kubectl rollout undo deployments <deployment-name>
```

- You can roll back to a specific version in the history using the --to-revision flag.
```
kubectl rollout undo deployments <deployment-name> --to-revision=X
```

- Delete a Deployment.
```
kubectl delete deployments <deployment-name>
```