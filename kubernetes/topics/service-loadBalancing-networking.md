# Service, Load balancing, Networking

- Every Pod in a cluster gets its own unique cluster-wide IP address.
- Pods can communicate with all other Pods on any other Node without NAT(Network Address Translation).
- Agents on a Node(e.g. system daemons, kubelet) can communicate with all Pods on that Node.
- Kubernetes IP addresses exist at the **Pos scope** - containers within a Pod share their network namespaces - 
  including IP address and MAC address. This means that containers within a Pod can all reach each others ports on ```localhost```.

## Service

- Expose an application running in your cluster behind a single outward-facing endpoint, even when the workload is
  split across multiple backends.
- Expose a Replicaset which is already running in the cluster:
```
kubectl expose rs go-demo-2 \
    --name=go-demo-2-svc \
    --target-port=28017 \
    --type=NodePort
```
Line 1: We specified that we want to expose a Replicaset(rs). <br>
Line 2: The name of the new Service should be <replicaSet-name>-svc.<br>
Line 3: The port that should be exposed is 28017. <br>
Line 4: We specified that the type of the Service should be NodePort. <br>
As a result, the target port will be exposed on every node of the cluster to the outside world, and it will be routed to one of the Pods controlled by the ReplicaSet.
- ClusterIP (the default type) exposes the port only inside the cluster.
- The LoadBalancer type is only useful when combined with cloud provider’s load balancer.
- ```ExternalName``` maps a service to an external address (e.g., kubernetes.io).
- Create Service.
```
kubectl create -f <yaml-file-name>
```
- Get service.
```
Kubectl get -f <yaml-file-name>
```
- Look at the endpoints.
```
kubectl get ep <Replicaset-name> -o yaml
```
- Requests will be sent to those Pods randomly. That randomness results in something similar to round-robin load balancing. 
  If the number of Pods does not change, each will receive an approximately equal number of requests.
- While livenessProbe is used to determine whether a Pod is alive or it should be replaced by a new one, the readinessProbe is used by the iptables.
  A Pod that does not pass the readinessProbe will be excluded and will not receive requests. 
- ```nohup kubectl port-forward service/go-demo-2-api --address 0.0.0.0 3000:8080 > /dev/null 2>&1 &```:
  - is used to run a kubectl port-forward command in the background while discarding the output.
  - This command forwards traffic from the local port 3000 to the Kubernetes service named go-demo-2-api on port 8080.
  - nohup: It prevents the process from being terminated when the terminal session is closed.
  - kubectl port-forward service/go-demo-2-api: It establishes a port forwarding connection to the specified Kubernetes service.
  - --address 0.0.0.0: It binds the port forwarding to all network interfaces, allowing external access.
  - 3000:8080: It specifies the local and remote port numbers for the port forwarding.
  - '> /dev/null': It redirects standard output to the null device, discarding the output.
  - 2>&1: It redirects standard error to standard output.
  - &: It runs the command in the background, allowing you to continue using the terminal.
  
- Sequence of events related to service discovery and components involved:
  - When the api container go-demo-2 tries to connect with the go-demo-2-db Service, it looks at the nameserver configured in /etc/resolv.conf.
    kubelet configured the nameserver with the kube-dns Service IP (10.96.0.10) during the Pod scheduling process.
  - The container queries the DNS server listening to port 53. go-demo-2-db DNS gets resolved to the service IP 10.0.0.19. 
    This DNS record was added by kube-dns during the service creation process.
  - The container uses the service IP which forwards requests through the iptables rules. They were added by kube-proxy during Service and Endpoint creation process.
  - Since we only have one replica of the go-demo-2-db Pod, iptables forwards requests to just one endpoint.
    If we had multiple replicas, iptables would act as a load balancer and forward requests randomly among Endpoints of the Service.
  