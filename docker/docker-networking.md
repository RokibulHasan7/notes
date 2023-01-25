## Docker Network

- The application doesn't work on the other system sue to the difference in computer environments.
    So the solution is, **Docker Networking**.
- **What is Docker Networking?** Docker networking enables a user to link a Docker container to as many networks as requires.
- Docker Networks are used to provide complete isolation for Docker containers.
- A user can add containers to more than one network.

### Advantage of Docker Network
- Rapid Deployment
- Portability
- Better Efficiency
- Faster Configuration
- Scalability
- Security

### Container Network Model

- It is an isolated sandbox that holds the network configuration of containers.
- Sandbox is created when a user requests to generate an endpoint on the network.
- It can have several endpoints in a network, as it represents a container's network configuration such as IP-address, MAC-address, DNS etc.
- End points establishes the connectivity for container services (within a network) with other services.
- It helps in providing connectivity among several endpoints that belong to the same network and isolate them from the rest.
- Docker Engine is the base engine installed on your host machine to build and run containers using Docker components and services.
- Network Drivers task is to manage the network with multiple drivers.
- Network Architecture - It provides the entry-point into libnetwork in order to maintain networks, whereas libnetwork supports multiple virtual drivers.

### Network Drivers
- Bridge
  - It is a private default network created on the host.
  - Container linked to this network have an internal IP address through which they communicate with each other easily.
  - The Docker server (daemon) creates virtual ethernet bridge **docker0** that performs the operation by automatically delivering packets among various network interfaces.
- Host
  - It is a public network.
  - It utilizes the host's IP address and TCP port space in order to display the services running inside the container.
  - It effectively disables network isolation between the docker host and the docker containers which means using this network driver a user will be unable to run multiple containers on the same host.
- None
  - In this network driver, the docker containers will neither have any access to external networks nor will it be able to communicate with other containers.
  - This option is used when a user wants to disable the networking access to a container.
  - In simple terms, None is called a loopback interface which means it has no external network interfaces.
- Overlay
  - This is utilized for creating an internal private network to the docker nodes in the docker swarm cluster.
  - It is an important network driver in Docker networking.
- Macvlan
  - It simplifies the communication process between containers.
  - This network assigns a MAC address to the Docker container. With this Mac address, the Docker server (daemon) routes the network traffic to a router.
  - It is suitable when a user wants to directly connect the container to the physical network rather than Docker host's.

Note: Docker Daemon is a server which interacts with the operating system and performs all kind of services.
