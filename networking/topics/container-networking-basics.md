# Container Networking Basics

## Containers

- **Cgroups and namespaces are Linux primitives to create containers.**
- An example of "low-level" functionality is creating cgroups and namespaces for containers.
- Example of "high-level" functionality: Build and test containers and deploy them.
- Low-level container runtime functionality:
    - Creating containers
    - Running containers
- High-level container runtime functionality:
    - Formatting container images
    - Building container images
    - Managing container images
    - Managing instance of containers
    - Sharing container images
- Open Container Initiative (OCI) promotes common, minimal, open standards and specifications for container technology.
- The three values guiding the OCI project are as follows:
  - Composable: Tools for managing containers should have clean interfaces.
  - Decentralized: The format and runtime should be well specified and developed by the community, not one organization.
  - Minimalist: The OCI spec strives to do several things well, be minimal and stable, and enable innovation and experimentation.
- ```LXC (Linux Containers)``` combines cgroups and namespaces to provide an isolated environment for running applications.
- ```runC``` is the most widely used container runtime developed initially as part of Docker and was later extracted as a 
  separate tool and library.
- runC is a command-line tool for running applications packaged according to the OCI format and is a compliant implementation
  of the OCI spec. 
- runC uses ```lincontainer```.
- ```containerd``` is a high-level runtime that was split off from Docker.
- ```Docker``` began as a monolith application, building all the previous functionality into a single binary known as the
  ```Docker engine```. The engine contained the ```Docker client``` or ```CLI``` that allows developers to build, run, and 
  push containers and images. The Docker server runs as a ```daemon``` to manage the data volumes and networks for running 
  containers. The client communicates to the server through the Docker API. It uses containerd to manage the container 
  life cycle, and it uses runC to spawn the container process.
- To run a container, the Docker engine creates the image and passes it to ```containerd```. containerd calls ```containerd-shim```,
  which uses runC to run the container.
- ```CRI``` is a plugin interface that enables Kubernetes, via **kubelet**, to communicate with any container runtime
  that satisfies the CRI interface.

### Container Primitives

- ```cgroups``` control access to resources in the kernel for our containers, and ```namespaces``` are individual slices
  of resources to manage separately from the root namespaces, i.e., the host.
- A ```cgroup``` is a Linux kernel feature that limits, accounts for, and isolates resource usage.
- ```cgroups``` allow administrators to control different CPU systems and memory for particulate processes.
- ```runC``` will create the cgroups for the container at creation time. ```A cgroup``` controls how much of a resource a container
  can use, while ```namespaces``` control what processes inside the container can see.
- ```lscgroup``` is a command-line tool that lists all the cgroups currently in the system.
- ```Namespaces``` are features of the Linux kernel that isolate and virtualize system resources of a collection of processes.

### Container Network Basics

- There are several modes for container networking:
  - None: 
    - Use this mode when the container does not need network access.
  - Bridge: 
    - The container runs in a ```private network``` internal to the host.
    - Communication with other containers in the network is open.
    - Bridge mode is the default mode of networking when the **--net** option is not specified.
  - Host:
    - The container shares the same IP address and the network namespace as that of the host.
    - This mode is useful if the container needs access to network resources on the hosts.
  - Macvlan:
    - Macvlan uses a parent interface. This interface can be a host interface such as eth0.
    - Macvlan allows a physical interface to have multiple MAC and IP addresses using Macvlan subinterfaces.
    - macvlan has four types: Private, VEPA, Bridge(which Docker default uses), and Passthrough.
    - Most cloud providers block Macvlan networking. Administrative access to networking equipment is needed.
  - IPvlan:
    - IPvlan is similar to Macvlan, with a significant difference: IPvlan does not assign MAC addresses to created subinterfaces.
    - IPvlan has two modes, L2 or L3.
  - Overlay:
    - Overlay allows for the extension of the same network across hosts in a container cluster.
    - The overlay network virtually sits on top of the underlay/physical networks.
  - Custom:
    - Custom bridge networking is the same as bridge networking but uses a bridge explicitly created for that container.

- Container-defined networking allows a container to share the address and network configuration of another container. This
  sharing enables **process isolation** between containers, where each container runs one service but where services can
  still communicate with one another on 127.0.0.1.
- Docker for MAC and Windows does not support host networking mode.
- When Docker starts, it creates a virtual bridge interface, docker0, on the host machine and assigns it a random IP address
  from the private 1918 range. The bridge passes packets between two connected devices, just like a physical bridge does.
- Each new container gets one interface automatically attached to the docker0 bridge.
- ```Bridge networks``` are for containers running on the same host. Communicating with containers running on different
  hosts can use an ```overlay network```.
- Docker uses the ```concept of local and global drivers```. Local drivers, a bridge, for example, are ```host-centric```
  and do not do cross-node coordination. Global drivers rely on libkv, a key-value store abstraction, to coordinate across
  machines.
- One technology that helps with routing between hosts for containers is a ```VXLAN```.
- Container Network Interface (CNI) is the software interface between the container runtime and the network implementation.
