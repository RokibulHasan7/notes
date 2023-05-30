# POD

- Atomic unit of Scheduling.
- A Pod is a collection of application containers and volumes running in the same execution environment.
- Each container within a Pod runs in its own **cgroup**, but they share a number of Linux namespaces.
- Application running in the same Pod share the same IP address and port space (network namespace), have the same
  hostname(UTS namespace) and can communicate using native interprocess communication channels over System V IPC or
  POSIX message queues (IPC namespace).
- When a Pod is deleted or a container restarts, any and all data in the container's filesystem is also deleted.

## POD Manifest

- The Kubernetes API server accepts and processes Pod manifest before storing them in ``` persistent storage (etcd)```.
- Scheduler places the Pods onto nodes depending on the resources and other constraints expressed in the Pod manifest.
- Scheduler can place multiple Pods on the same machine as long as there are sufficient resources.
- Scheduling multiple replicas of the same application onto the same machine is worse for reliability, since the
  machine is a single failure domain.
- The Kubernetes Scheduler tries to ensure that Pods from same application are distributed onto different machines
  for reliability in the presence of such failures.
- ```metadata``` provides information that does not influence how the Pod behaves. 

## Commands

- Create a Pod
```
kubectl create -f <yml-file>
```

- To see Pods
```
kubectl get pods
```

- In some cases, you might want to retrieve a bit more information by specifying wide output.
```
kubectl get pods -o wide
```

- If you’d like to parse the output, using json format is probably the best option.
```
kubectl get pods -o json

kubectl get pods -o yaml
```

- Run a pod after creating manifest
```
kubectl apply -f <yaml-file>

kubectl port-forward kuard <port-number>:<port-number>
```

- Delete pod
```
kubectl delete pods/<pod-name>

kubectl delete -f <yaml-file-name>  ---> another way
```

- To see pod details
```
kubectl describe pods <pod-name>
```

- See the pods running on port
```
lsof -i :<port-number>
```

- Kill the running pod
```
kill -i <pod-id>
```

- To see current logs from running instance.
```
kubectl logs <pod-name>
```

- To see what's going on into Pod.
```
kubectl exec <pod-name> -- date
```
- If we'd like to create a Pod with a Mongo database.
```
kubectl run db --image mongo
```
We can confirm that a container based on the mongo image is indeed running inside the cluster by listing all the containers 
based on the mongo image.
```
docker exec -it <node-name> ctr container ls | grep mongo
```
The above approach used to run Pods is not the best one. We used the imperative way to tell Kubernetes what to do. Even 
though there are cases when that might be useful, most of the time we want to leverage the declarative approach.

### Liveness Probe

- Liveness determines if an application is running properly. Containers that fail liveness checks are restarted.
- Liveness health checks run application-specific logic, like loading a web page, to verify that the application
  is not just still running, but is functioning properly.
- It has to be defined in Pod manifest.
- ```initialDelaySeconds: X``` - probe set will not be called until X seconds after all the containers in the Pod
  are created.
- ```periodSeconds: X``` - Kubernetes will call the probe every X seconds.
- ```failureThreshold: X``` - If more than X consecutive probes fail, the container will fail and restart.
- ```timeoutSeconds: X``` - The probe must respond within the X second timeout and the HTTP status code must be equal
  to or greater than 200 and less than 400 to be considered successful.
- While the default response to a failed liveness check is to restart the Pod, the actual behaviour is governed by
  the Pod's **restartPolicy**. There are three options for the restart policy: **Always** (the default), **OnFailure**
  (restart only on liveness failure or nonzero process exit code), or **Never**.


### Readiness Probe

- Readiness describes when a container is ready to serve user requests.
- Containers that fail readiness checks are removed from service load balancers.
- Readiness probes are configured similarly to liveness probes.
- Combining the readiness and liveness probes helps ensure only healthy container are running within the cluster.

### Resource Management

- **Resource Request** specify the minimum amount of a resource required to run the application.
- **Resource Limits** specify the maximum amount of a resource that an application can consume.
- Resources are requested ```per container, not per Pod```.
- The Kubernetes scheduler will ensure that the sum of all requests of all Pods on a **node** does not exceed the capacity
  of the node.
- A Pod is **guaranteed** to have at least the requested resources when running on the node.
- CPU requests are implemented suing the ```cpu-shares``` functionality in the Linux kernel.

### Use Voolume with Pods

- Two different containers in a Pod can mount the same volume at different mount paths.
- Different ways of using volumes with pods - Communication/ synchronization, Cache, Persistent data, Mounting the
  host filesystem.

### Sequence of Pod Scheduling

- The sequence of events that transpired with the ```kubectl create -f db.yml``` command is as follows:
  1. Kubernetes client (kubectl) sent a request to the API server requesting creation of a Pod defined in the db.yml file.
  2. Since the scheduler is watching the API server for new events, it detected that there is an unassigned Pod.
  3. The scheduler decided which node to assign the Pod to and sent that information to the API server.
  4. Kubelet is also watching the API server. It detected that the Pod was assigned to the node it is running on.
  5. Kubelet sent a request to Docker requesting the creation of the containers that form the Pod.
  6. Finally, Kubelet sent a request to the API server notifying it that the Pod was created successfully.
  
### Playing With the Running Pod

- 
