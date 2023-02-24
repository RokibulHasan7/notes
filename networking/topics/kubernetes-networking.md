# Kubernetes Networking Introduction

- Kubernetes networking looks to solve these four networking issues:
    - Highly coupled container-to-container communications
    - Pod-to-pod communications
    - Pod-to-service communications
    - External-to-service communications
- Docker networking model uses a virtual bridge network by default, which is defined per host and is a private network where
  container attach.
- The container's IP address is allocated a private IP address.

## Kubernetes Networking Model

- Kubernetes networking model natively supports multihost cluster networking.
- Pods can communicate with each other by default, regardless of which host they are deployed on.
- A Pod has a unique IP address. If Pod did not have unique IP addresses, then two Pods on a node could contend for the
  same port (such as two web servers, both trying to listen on port 80).
- ```StatefulSets``` are a built-in workload type intended for workloads such as databases, which maintain a pod identity
  concept and give a new pod the same name and IP address as the pod it replaces.
- Every Kubernetes node runs a component called the ```kubelet```, which manages pods on the node.
- The networking functionality in the Kubelet comes from API interactions with a CNI plugin on the node. 
- The CNI plugin is what manages Pod IP addresses and individual container network provisioning.

## Node and Pod Network Layout

- Nodes and Pods must have L3 connectivity in this IP address space.
- Without L3 connectivity, TCP handshakes would not be possible, as the ```SYN-ACK``` could not be delivered.
- Pods do not have MAC addresses. So, L2 connectivity to pods is not possible.
- There are broadly three approaches, with many variations, to structuring a cluster's network:
  - Isolated Networks
  - Flat Networks
  - Island Networks

## Isolated Networks

- In an isolated cluster network, nodes are routable on the broader network (i.e., hosts that are not part of the cluster
  can reach nodes in the cluster), but pods are not.
- Pods cannot reach other Pods outside the cluster.
- Kubernetes API server will need to be routable from the broader network, if external systems or users should be able to
  access the Kubernetes API.

### Flat Networks

- In a Flat network, all Pods have an IP address that is routable from the broader network.
- A load balancer outside the cluster can load balance pods, such as gRPC client in another cluster.
- This model requires a large, contiguous IP address space for each cluster(i.e., a range of IP addresses where every IP
  address in the range in under your control).
- Kubernetes requires a single CIDR for pod IP addresses (for each IP family).
- This model is achievable with a private subnet; however it is much harder and more expensive to do with public IP
  addresses, especially IPv4 addresses.

### Island Networks

- Island cluster network is a combination of Isolated and flat networks.
- In an Island cluster setup, nodes have L3 connectivity with the broader network, but Pods do not.
- Traffic to and from Pods must pass through some form of proxy, through nodes. This is achieved by ```iptables``` source NAT
  on a pod's packets leaving the node. This setup called ```masquerading```, uses SNAT to rewrite packet sources from the 
  pod's IP address to the node's IP address.
- Sharing an IP address while also using NAT hides the individual pod IP addresses.


- ```Control plane``` refers to all the functions and processes that determine which path to use to send the packet or frame.
- ```Data plane``` refers to all the functions and processes that forward packets/frames from one interface to another
  based on control plane logic.

- ```Kube-controller-manager``` includes multiple controllers that manage the Kubernetes network stack. Administrators set
  the cluster CIDR here.


  
## Kubelet

- The Kubelet is a single binary that runs on every worker node in a cluster.
- The Kubelet is responsible for managing any Pods scheduled to the Node and providing status updates for te Node and Pods on it.
- Kubelet primarily acts as a ```coordinator``` for other software on the Node.
- Kubelet manages a container networking implementation (CNI) and a container runtime (CRI).
- Kubelet makes an ADD call to the CNI, which tells the CNI plugin to create the pod network.


## Pod Readiness and Probes

- Pod readiness is an additional indication of whether the Pod is ready to serve traffic.
- Pod readiness determines whether the Pod address shows up in the ```Endpoints``` object from an external source.
- Probes effect the ```.Status.Phase``` field of a Pod.
- Pod phases:
  - Pending: The Pod has been accepted by the cluster, but one or more of the containers has not been set up and made ready
    to run. This includes the time a pod spends waiting to be scheduled as well as the time spent downloading container
    images over the network.
  - Running: The Pod has been scheduled to a Node and all the containers have been created. At least one container is still
    running or is in the process of starting or restarting.
  - Succeeded: All containers in the Pod have terminated in success and will not be restarted.
  - Failed: All containers in the Pod have terminated, and at least one container has terminated in failure. That is, the 
    container either exited with nonzero status or was terminated by the system.
  - Unknown: For some reason the state of the Pod could not be determined. This phase typically occurs due to an error in
    communicating with the Kubelet where the pod should be running.
- The Kubelet performs several types of health checks for individual containers in a Pod:
  - Liveness Probes(livenessProbe)
  - Readiness Probes(readinessProbe)
  - Startup Probes(startupProbe)
- Kubelet must be able to connect to all containers running on that node in order to perform any HTTP health checks.
- Each probe has one of three results:
  - Success: Container passed the diagnostic.
  - Failure: Container failed the diagnostic.
  - Unknown: Diagnostic failed, so no action should be taken.
- if the probe fails more than the ```failureThreshold``` number of times, Kubernetes will consider the check to have failed.
- When a container's **readiness probe fails**, the Kubelet does not terminate it. Instead, the Kubelet writes the failure to
  the Pod's status.
- If the **liveness probes fail**, the kubelet will terminate the container. The intended use case for liveness probes is to
  let the Kubelet know when to restart a container.
- A ```startup probe``` can provide a grace period before a liveness probe can take effect. Liveness probes will not terminate a 
  container before the startup probe has succeeded.
- The ```Endpoints/EndpointsSlice``` controllers also react to failing readiness probes. If a pod's readiness probe fails,
  the Pod's IP address will not be in the endpoint object and the service will not route traffic to it.
- The ```startupProbe``` will inform the ```kubelet``` whether the application inside the container is started.
- Probe configurable options:
  - initialDelaySeconds
  - periodSeconds
  - timeoutSeconds
  - successThreshold
  - failureThreshold
- ```gRPC``` - **API to communicate from the API server to etcd**.
- **Communication between the Pods and the Kubelet is made possible by the CNI**.

## CNI Specification

- There are Four operation that a CNI plugin must support:
  - ADD: Add a container to the network.
  - DEL: Delete a container from the network.
  - CHECK: Return an error if there is a problem with the container's network.
  - VERSION: Report version information about the plugin.

### CNI Plugins

- CNI plugin has two primary responsibilities:
  - Allocate and assign unique IP addresses for Pods
  - Ensure that routes exist within Kubernetes to each Pod IP address.
- If there are too few IP addresses or it is not possible to attach sufficient to a Node, cluster admins will need to use
  a CNI plugin that supports an overlay network.
- There are **two broad categories of CNI network models**: flat networks and overlay networks.
- In a flat network, **the CNI driver uses IP addresses from the cluster's network, which typically requires many IP addresses
  to be available to the cluster.**
- In an overlay network, **the CNI driver creates a secondary network within Kubernetes, which uses the cluster's network(called
  the underlay network) to send packets.
- Overlay networks create a virtual network within the cluster.
- In an overlay network, the CNI plugin encapsulates packets.
- CNI plugin is also responsible for calling ```IPAM plugins``` for IP addressing.

### IPAM Interface

- The CNI spec has a second interface, the IP Address Management(IPAM) interface, to reduce duplication of IP allocation 
  code in CNI plugins.

- ```Cillium``` is open source software for transparently securing network connectivity between application containers.
- ```Flannel``` focuses on the network and is a simple and easy way to configure a layer 3 network fabric designed for K8s.
- ```Calico``` give Network Policy support.
- ```Weave Net``` also gives network policy support and it's network setup is Mesh overlay network.

### Kube-proxy

- Kube-proxy is another per-node daemon in Kubernetes, like Kubelet.
- Kube-proxy provides basic load balancing functionality within the cluster.
- It implements services and relies on Endpoints/EndpointSlices.
- Kube-proxy is responsible for routing requests to a service's cluster IP address to healthy Pods.
- Kube-proxy has four modes, which change its runtime mode and exact feature set: **userspace, iptables, ipvs and kernelspace**.
  - userspace Mode: userspace mode is no longer commonly used, and we suggest avoiding it unless you have a clear reason to
    use it.
  - iptables Mode: iptables mode uses iptables entirely. It is the ```deafault mode```, and the most commonly used.
    - iptables mode performs connection fan-out, instead of true load balancing. In other words, iptables mode will route
      a connection to a backend pod, and all requests made using that connection will go to the same Pod, until the
      connection is terminated.
  - ipvs Mode: ipvs mode uses IPVS, instead of iptables, for connection load balancing. ipvs mode supports ```six load balancing modes```,
    specified with ```--ipvs-scheduler```:
      - rr: Round-robin
      - lc: Least connection
      - dh: Destination hashing
      - sh: Source Hashing
      - sed: Shortest expected delay
      - nq: Never queue
    - Round-robin(rr) is the default load balancing mode.
  - kernelspace Mode: kernelspace is the newest, Windows-only mode. It provides an alternative to userspace mode for 
    Kubernetes on Windows, as iptables and ipvs are specific to Linux.

## Network Policy

- Kubernetes default behaviour is to allow traffic between any two Pods in the cluster network.
- **We strongly discourage running real clusters without ```NetworkPolicy```. Since all Pods can communicate with all other
  Pods, we strongly recommend that application owners use NetworkPolicy objects along with other application-layer security
  measures, such as authentication tokens or mutual Transport Layer Security(mTLS), for any network communication.
- NetworkPolicy is a resource type in Kubernetes that contains allow-based firewall rules.
- The NetworkPolicy resource acts as a configuration for CNI plugins, which themselves are responsible for ensuring
  connectivity between Pods.
- NetworkPolicy objects rely heavily on labels and selectors.

### Selecting Pods

- Pods are unrestricted until they are selected by a NetworkPolicy.
- A NetworkPolicy has a ```spec.policyTypes``` field containing a list of policy types.
- If we select a Pod with a NetworkPolicy that has ingress listed but not egress, then ingress will be restricted, and
  egress will not.
- The ```spec.podSelector``` field will dictate which Pods to apply the NetworkPolicy to.
- NetworkPolicy objects are ```namespaced``` objects, which means they exist in and apply to a specific namespace.
- The ```spec.podSelector``` field can select Pods only when they are in the same namespace as the NetworkPolicy.

### Rules

- If multiple NetworkPolicy objects select a Pod, all rules in each of those NetworkPolicy objects apply.
- Ingress rules and egress rules are discrete types in the NetworkPolicy API. However, they are functionally structured the
  same way. Each NetworkPolicyIngressRule/NetworkPolicyEgressRule contains a list of ports and a list of NetworkPolicyPeers.
- A ```NetworkPolicyPeer``` has four ways for rules to refer to networked entities: **ipBlock, namespaceSelector, podSelector,
  and a combination.**
- ```ipBlock``` is useful for allowing traffic to and from external systems.
- **It is common and highly recommended by security experts to keep the scope of namespaces small; typical namespace scopes are
  per an app or service group or team.**

### DNS

- The ```kube-dns``` container watches the Kubernetes API and serves DNS records based on the Kubernetes DNS specification,
  ```dnsmasq``` provides caching and stub domain support and ```sidecar``` provides metrics and health checks.
- There are several differences between CoreDNS and KubeDNS:
  - CoreDNS runs as a single container.
  - CoreDNS is a Go process that replicates and enhances the functionality of Kube-DNS.
  - CoreDNS is designed to be a general-purpose DNS server that is backward compatible with Kubernetes, and its extendable 
    plugins can do more than is provided in Kubernetes DNS specification.
- There are four options for ```dnsPolicy``` that significantly affect how DNS resolutions work inside a Pod:
  - Default: The Pod inherits the name resolution configuration from the Node that the Pods run on.
  - ClusterFirst: Any DNS query that does not match the cluster domain suffix, such as www.kubernetes.io, is sent to the
    ```upstream name server``` inherited from the Node.
  - ClusterFirstWithHostNet: For Pods running with ```hostNetwork```, admins should set the DNS policy to ClusterFirstWithHostNet.
  - None: All DNS settings use the ```dnsConfig``` field in the Pod spec.
    - If ```none```, developers will have to specify name servers in the Pod spec.
    - ```nameservers:``` is a list of IP addresses that the Pod will use as DNS server. There can be at most **three** IP
      addresses specified.
    - ```searches:``` is a list of DNS search domains for hostname lookup in the Pod. Kubernetes allows for at most **six**
      search domains.

### IPv4/IPv6 Dual Stack

- IPv4/IPv6 features enable the following features for Pod networking:
  - A single IPv4 and IPv6 address per Pod.
  - IPv4 and IPv6 services.
  - Pod cluster egress routing via IPv4 and IPv6 interfaces.
- When IPv4/IPv6 is on in a cluster, services now have an extra field in which developers can choose the ```ipFamilyPolicy```
  to deploy applications:
  - SingleStack
  - PreferDualStack
  - RequireDualStack
  - ipFamilies
