# Service, Load balancing, Networking

- Every Pod in a cluster gets its own unique cluster-wide IP address.
- Pods can communicate with all other Pods on any other Node without NAT(Network Address Translation).
- Agents on a Node(e.g. system daemons, kubelet) can communicate with all Pods on that Node.
- Kubernetes IP addresses exist at the **Pos scope** - containers within a Pod share their network namespaces - 
  including IP address and MAC address. This means that containers within a Pod can all reach each others ports on ```localhost```.

## Service

- Expose an application running in your cluster behind a single outward-facing endpoint, even when the workload is
  split across multiple backends.