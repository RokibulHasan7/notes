# Linux Networking

## Network Interface

- Computer use a ```network interface``` to communicate with the outside world.
- IP addresses are assigned to network interfaces. A typical interface may have one IPv4 address and one IPv6 address,
  but multiple addresses can be assigned to the same interface.
- If you run ```ifconfig```, you will see a list of all network interfaces and their configurations (including IP addresses).
- The ```loopback interface``` in a special interface for same-host communication. 127.0.0.1 is the standard IP 
  address for the loopback interface. The loopback interface is commonly abbreviated as ```lo```.

## Bridge Interface

- Bridge interface is a type of network interface that connects two or more separate networks together and allows 
  devices on one network to communicate with devices on the other network.
- A bridge interface acts as a bridge between the two networks, forwarding data packets between them.
- In Kubernetes, Bridge allow Pods with their individual network interfaces, to interact with the broader network
  via the node's network interface.
- Bridge can be managed and created using ```ip``` and ```brctl``` command.

## Netfilter

- Netfilter, included in Linux since 2.3, is a critical component of packet handling.
- Netfilter is a framework of kernel hooks, which allow userspace programs to handle packets on behalf of the kernel.
- Netfilter was created jointly with ```iptables```, to separate kernel and userspace code.
- **Netfilter has five hooks.**
- A packet originating from a local process will always trigger NF_IP_LOCAL_OUT hooks and then NF_IP_POST_ROUTING hooks.
- If a process sends a packet destined for the same host, it triggers the NF_IP_LOCAL_OUT and then the NF_IP_POST_ROUTING
  hooks before "reentering" the system and triggering the NF_IP_PRE_ROUTING and NF_IP_LOCAL_IN hooks.

## Conntrack

- Conntrack is a component of Netfilter used to track the state of connections to (and from) the machine.
- Conntrack tracking directly associates packets with a particular connection.
- Conntrack is important on systems that handle firewalling or NAT.
- NAT relies on Conntrack to function.
- The internet uses addresses for routing and computers use port numbers for application mapping.
- A flow contains metadata about the connection and its state. Conntrack refers to these connections as flows.
- Conntrack stores flows in a hash table.
- ```conntrack -L``` shows all current flows.

## iptables

- ```iptables``` is staple of Linux sysadmins.
- ```iptables``` can be used to create firewall and audit logs, mutate and reroute packets, and even implement crude
  connection fan-out. 
- iptables uses Netfilter, which allows iptables to intercept and mutate packets.
- Most Linux distributions are replacing ```iptables``` with ```nftables```, a similar but more performant tool
  built a top Netfilter.
- There are three key concepts in iptables: ```tables, chains and rules```. They are considered hierarchical in 
  nature: ```a table contains chains, and a chain contains rules```.
- The three most commonly applicable tables are: Filter(for Firewall-related rules), NAT(for NAT-related rules) and
  Mangle(for non-NAT packet-mutating rules).
- iptables has ```five tables```. which are:
    - Filter: The filter table handle acceptance and rejection of packets.
    - NAT: The NAT table is used to modify the source or destination IP addresses.
    - Mangle: The mangle table can perform general-purpose editing of packet headers, but it is not intended for NAT.
      It can also "mark" the packet with iptables-only metadata.
    - Raw: The Raw table allows for packet mutation before connection tracking and other tables are handled. Its
      most common use is to disable connection tracking for some packets.
    - Security: SELinux uses the security table for packet handling. It is not applicable on a machine that is not
      using SELinux.
- iptables executes tables in a particular order: Raw, Mangle, NAT, Filter.
- The order of execution is chains, then tables.
- ```iptables chains``` are a list of rules. When a packet  triggers or passes through a chain, each rules is
  sequentially evaluated, until the packet matches a "terminating target" (such as DROP), or the packet reaches
  the end of the chain.
- **The built-in, "top-level" chains are PREROUTING, INPUT, NAT, OUTPUT, and POSTROUTING.** These are powered by
  Netfilter hooks. Each chain corresponds to a hook.
- iptables is, effectively, running tens or hundreds or thousands of ```if statements``` against every single 
  packet that goes in or out of your system. That has measurable impact on packet latency, CPU use, and network
  throughput.
- A well-organized set of chains reduces this overhead by eliminating effectively redundant checks or actions.
- **iptables's performance given a service with many Pods is still a problem in Kubernetes, which makes other
  solutions with less or ````no iptables```` use, such as ```IPVS or eBPF```, more appealing.**
- iptables can masquerade connections, making it appear as if the packets came from their own IP address.
- In Kubernetes, masquerading can make Pods use their Node's IP address, despite the fact that Pods have unique IP
  addresses. This is necessary to communicate outside the cluster in many setups, where Pods have internal IP
  addresses that cannot communicate directly with the internet.
- The MASQUERADE target is similar to SNAT. 
- iptables can perform connection-level load balancing or more accurately, connection fan-out.

### iptables Rules

- ```Rules``` have to parts: a match condition and an action (called a target).
- ```The match condition``` describe a **packet attribute**. If the packet matches, the action will be executed. 
  if the packet does not match, iptables will move to check the next rule.
- There are two kind of target actions: ```terminating and nonterminating```.
- A terminating target will stop iptables from checking subsequent targets in the chain, essentially acting as a 
  final decision.
- A nonterminating target will allow iptables to continue checking subsequent targets in the chain. 
- ```ACCEPT, DROP, REJECT and RETURN``` are all terminating targets.
- ```ACCEPT and RETURN``` are terminating only ```within their chain```.


## IPVS (IP Virtual Server)

- IP Virtual Server(IPVS) is a Linux connection (L4) load balancer.
- **IPVS supports multiple load balancing modes (in contrast with the iptables one)**. This allows IPVS to spread
  load more effectively than iptables, depending on IPVS configuration and traffic patterns.
- There are three aspects to look when it comes to issues with iptables as a load balancer:
  - Number of Nodes in the Cluster
  - Time
  - Latency
- **IPVS supports session affinity**. Which is exposed as an option in services (Service.spec.sessionAffinity and
  Service.spec.sessionAffinityConfig). 
- Repeated connections, within the session affinity time window, will route to the same host. This can be useful for
  **scenario such as minimizing cache misses**.
- To create a basic load balancer with two equally weighted destinations, run ```ipvsadm -A -t <address> -s <mode>```.
  -A, -E and -D are used to add, edit and delete virtual services respectively.
- List the IPVS hosts: ```ipvsadm -L```.

## eBPF (extended Berkeley Packet Filter)

- eBPF is a programming system that allows special sandboxed programs to run in the kernel without passing back
  and forth between kernel and user space, like we saw with Netfilter and iptables.
- **BPF** is a technology used in the kernel, among other things, **to analyze network traffic**.
- BPF supports filtering packets, which allows a userspace process to supply a filter that specifies which packets
  it wants to inspect.
- One of BPF's use cases is tcpdump.
- An eBPF program has direct access to syscalls. eBPF programs can directly watch and block syscalls, without the
  usual approach of adding kernel hooks to an userspace program. Because of its performance characteristics, it is
  well suited for writing networking software.
- There are many reasons to use eBPF with Kubernetes:
  - Performance (hashing table versus iptables list)
  - Tracing: Using BPF, we can gather Pod and Container-level network statistics.
  - Auditing ```kubectl exec``` with eBPF
  - Security
    - Seccomp
    - Falco
- The most common use of eBPF in Kubernetes is ```Cilium, CNI and service implementation```.
- **Cilium replaces kube-proxy, which writes iptables rules to map a services IP address to its corresponding Pods.**
- Through eBPF, Cilium can **intercept and route all packets directly in the Kernel, which is faster and allows for
  application-level (layer 7) load balancing.


## Network Troubleshooting Tools


| Case                        | Tools                                              |
|-----------------------------|----------------------------------------------------|
| Checking connectivity       | traceroute, ping, telnet, netcat                   |
| Port scanning               | nmap                                               | 
| Checking DNS records        | dig, commands mentioned in "checking connectivity" |
| Checking HTTP/1             | cURL, telnet, netcat                               |
| Checking HTTPS              | openSSL, cURL                                      |
| Checking listening programs | netstat                                            |


