# Kubernetes Networking Abstractions

## StatefulSets

- StatefulSets are a workload abstraction in Kubernetes to manage Pods like you would a deployment.
- StatefulSets add the following features for applications:
    - Stable, unique network identifiers.
    - Stable, persistent storage.
    - Ordered, graceful deployment and scaling.
    - Ordered, automated rolling updates.
- The deployment resource is better suited for applications that do not have these requirements (for example, a service that
  stores data in an external database).
- StatefulSets require a service, preferably **headless**, to be responsible for the network identity of the Pods, and end users
  are responsible for creating this service.
- StatefulSets offer functionality for a specific use case in Kubernetes. They should not be used for everyday application
  deployment.

## Endpoints

- Endpoints help identify what Pods are running for the service it powers.
- Endpoints are created and managed by services.
- Each endpoint contains a list of ports (which apply to all pods) and two lists of addresses: ready and unready.
- Addresses are listed in ```.addresses``` if they are passing Pod readiness checks.
- Addresses are listed in ```.notReadyAddresses``` if they are not passed Pod readiness checks. This makes endpoints a
  ```service discovery tool```, where you can watch an ```Endpoints``` object to see the health and addresses of all Pods.
- For large deployments, the endpoint object can become very large, so much so that it can actually slow down changes in the
  cluster. To solve that issue, the Kubernetes maintainers have come up with endpoints slices.

### Endpoint Slices

- In a typical cluster, Kubernetes runs ```kube-proxy``` on every Node.
- ```kube-proxy``` is responsible for the per-node portions of making services work, by handling routing and outbound
  load balancing to all the Pods in a service. To do that, kube-proxy watches all endpoints in the cluster so it knows all
  applicable Pods that all services should route to.
- Many Kubernetes users consider endpoints watches to be the **ultimate bottleneck of cluster size**.
- Endpoint slices are an approach that allows kube-proxy's fundamental design to continue, while drastically reducing the watch
  bottleneck in large clusters where large services are used.
- Endpoint slices have similar contents of ```Endpoints``` objects but also ```include an array of endpoints```.
- With "regular" endpoints, a Kubernetes service creates one endpoint for all Pods in the service. A service creates 
  ```multiple``` endpoint slices, each containing a ```subset``` of Pods. **The union of all endpoints slices for a service
  contains all Pods in the service.** 
- This way, an IP address change (due to a new Pod, a deleted Pod, or a Pod's health changing) will result in a much smaller
  data transfer to watchers.
- Because **Kubernetes doesn't have a transactional API, the same address may appear temporarily in multiple slices.** Any
  conde consuming endpoint slices (such as kube-proxy) must be able to account for this.
- The maximum number of addresses in an endpoint slice is set using the ```--max-endpoints-per-slice kube-controller-manager```
  flag. The current default is 100 and the maximum is 1000. The endpoint slide controller attempt to fill existing slices
  before creating new ones, but does not rebalance endpoint slice.
- The endpoint slice controller mirrors endpoint to endpoint slice, to allow systems to continue writing endpoints while
  treating endpoint slice as the source of truth. There are four exceptions that will prevent mirroring:
  - There is no corresponding service.
  - The corresponding service resource selects Pods.
  - The Endpoints object has the label ```endpointslice.kubernetes.io/skip-mirror: true```.
  - The Endpoints object has the annotation ```control-plane.alpha.kubernetes.io/leader```.
- You can fetch all endpoint slices for a specific service by fetching endpoint slices filtered to the desired name in
  ```.metadata.labels."kubernetes.io/service-name"```.

## Kubernetes Services

- A service in Kubernetes is a load balancing abstraction within a cluster.
- There are four type of services: ```ClusterIP, NodePort, LoadBalancer, and ExternalName```.
- Services use a standard Pod selector to match Pod. The service includes all matching Pods.
- There is also a ```pause container``` that is created for each Pod. The pause container manages the namespaces for the Pod.
- **The pause container is the parent container for all running containers in Pod. It holds and shares all the namespaces 
  for the Pod.**

### NodePort

- A NodePort service provides a simple way for external software, such as a load balancer, to route traffic to the Pods.
- The software only needs to be aware of node IP addresses, and the service's port(s).
- A NodePort service exposes a fixed port on all Nodes, which routes to applicable Pods.
- A NodePort service uses the ```.spec.ports.[].nodePort``` field to specify the port to open on all Nodes, for the
  corresponding port on Pods.
- Using a NodePort service, external users can connect to the NodePort on any Node and be routed to a Pod on a Node
  that has a Pod backing that service.
- ```ExternalTrafficPolicy``` indicates how a service will route external traffic to either node-local or cluster-wide endpoints.
- A ```cluster``` value means that for each worker Node, the **kube-proxy iptables** rules are set up to route the traffic
  to the Pods backing the service anywhere in the cluster.
- A ```local``` value means the **kube-proxy iptables** rules are set up only on the worker Nodes with relevant Pods running
  to route the traffic local to the worker Node.
- Using Local also allows application developers to preserve the source IP of the user request.
- If you set ```externalTrafficPolicy``` to the value ```local```, kube-proxy will proxy requests only node-local endpoints
  and will not forward traffic to other Nodes.
- **A NodePort deployment will fail if it cannot allocate the requested port. Also, ports must be tracked across all 
  applications using a NodePort service.**
- Downside of using the NodePort service type is that the load balancer or client software must be aware of the Node IP addresses.

### ClusterIP

- A ClusterIP service provides an internal load balancer with a single IP address that maps to all matching (and ready) Pods.
- The IP addresses of Pods share the life cycle of the Pod and thus are not reliable for clients to use for request. Services
  help overcome this Pod networking design.
- The service's IP address must be within the CIDR set in ```service-cluster-ip-range```, in the API server.
- ```kube-proxy``` is responsible for making the ClusterIP service address route to all applicable Pods. In "normal" 
  configurations, kube-proxy performs **L4 load balancing**, which may not be sufficient.
- **A particular use case example for ClusterIP is when a workload requires a load balancer within the same cluster.**
- **The ClusterIP service is for internal traffic to the cluster, and it suffers the same issues as endpoints do. As the
  service size grows, updated to it will slow.**

### Headless

- A headless service isn't a formal type of service (i.e., there is no .spec.type:Headless). A headless service is a 
  service with ```.spec.clusterIP: "None"```.
- When ClusterIP is set to None, the service does not support any load balancing functionality.
- A headless service provides a generic way to watch endpoints, without needing to interact with the Kubernetes API.
- Headless services allow developers to deploy multiple copies of a Pod in a deployment.
- Headless has a specific use case and is not typically used for deployments.
- **If developers need to let the client decide which endpoints to use, headless is the appropriate type of service to deploy.**
- Two example of headless services are ```clusterd databases adn application that have client-side load-balancing logic built into the code```.

### ExternalName Service

- ExternalName is a special type of service that does not have selectors and uses DNS names instead.
- If developers are migrating an application into Kubernetes but its dependencies are staying external to the cluster,
  ExternalName service allows them to define a DNS record internal to the cluster no matter where the service actually runs.
- The ExternalName service allows developers to map a service to a DNS name.
- Sending traffic to a headless service via a DNS record is possible but inadvisable. **DNS is a notoriously poor way to
  load balance.**
- If you need to be able to send traffic to the service's DNS address, consider a (standard) ClusterIP or LoadBalancer service.
- **The "correct" way to use a headless service is to query the service's A/AAAA DNS record and use that data in a server-side
  or client-side load balancer.

### LoadBalancer

- LoadBalancer service exposes services external to the cluster network.
- LoadBalancer services handle L4 traffic, so they will work for any TCP and UDP service, provided the load balancer
  selected supports L4 traffic.
- Ingress handles L7 traffic.
- Once the load balancer has been provisioned, its IP address will be written to ```.status.loadBalancer.ingress.ip```.
- LoadBalancer services are useful for exposing TCP or UDP services to the outside world.
- Test provisioned load balancer: ```kubectl ger svc <loadBalancer name>```.
- It is important to remember that LoadBalancer services require specific integrations and will not work without cloud
  provider support, or manually installed software such as MetalLB.

## Ingress

- Ingress is a Kubernetes-specific L7(HTTP) load balancer, which is accessible externally, contrasting with L4 ClusterIP
  service, which is internal to the cluster. This is the typical choice for exposing an HTTP(S) workload to external users.
- An ingress can be a single entry point into an API or a microservice-based architecture.
- Ingress is a configuration spec (with multiple implementations) for routing HTTP traffic to Kubernetes services.
- To manage traffic in a cluster with ingress, there are two components required: ```the controller and rules```.
  The controller manages ingress Pods, and the rules deployed define how the traffic is routed.

### Ingress Controllers and Rules

- We call ingress implementations ingress controllers.
- In Kubernetes, a controller is software that is responsible for managing a typical resource type and making reality match
  the desired state.
- There are two general kind of controllers: ```external load balancer controllers``` and ```internal load balancer controller```.
- External load balancer controllers create a load balancer that exists "outside" the cluster, such as a cloud provider product.
- Internal load balancer controllers deploy a load balancer that runs within the cluster and do not directly solve the problem
  of routing consumers to the load balancer.
- **The primary motivation for choosing an internal load balancer is cost reduction.**
- An internal load balancer for ingress can route traffic for multiple ingress objects, whereas **an external load balancer
  controller typically needs one load balancer per ingress**.
- Ingresses have a "default backend" where requests are routed if no rule matches.
- It is possible to use multiple ingress controllers in a single cluster, using ```IngressClass``` settings. An ingress
  class represents an ingress controller, and therefore a specific ingress implementation.
- **Ingress only supports HTTP(S) requests**, which is insufficient if your service uses a different protocol (e.g., most
  databases use their own protocols).
- Some things to consider when deciding on the ingress for your cluster:
  - Protocol support: Do you need more than TCP/UDP, for example gRPC integration or WebSocket?
  - Commercial support: Do you need commercial support?
  - Advanced features: Are JWT/oAuth2 authentication or circuit breakers requirements for your applications?
  - API gateway features: Do you need some API gateway functionalities such as rate-limiting?
  - Traffic distribution: Does your application require support for specialized traffic distribution like canary A/B 
    testing or mirroring?

## Service Meshes

- A new cluster with the default options has some limitations.
- A ```service mesh``` is an API-driven infrastructure layer for handling service-to-service communication.
- From a security point of view, all traffic inside the cluster is **unencrypted between Pods**, and each application team
  that runs a service must configure monitoring separately for each service.
- Service meshes support more than the basic deployment type; **they support rolling updates and re-creations, like Canary does**.
- With service meshes, developers can add fault testing.
- There are several pieces of functionality that a service mesh enhances or provides in a default Kubernetes cluster network:
  - Service Discovery
  - Load Balancing
  - Communication Resiliency
  - Security
  - Observability
  - Routing Control
  - API
- There are several options to use when deploying a service mesh:
  - Istio
  - Consul
  - AWS App Mesh
  - Linkerd
- The best use case for a service mesh is mTLS between services. Other higher-level use cases for developers include
  circuit breaking and fault testing for APIs.
- 