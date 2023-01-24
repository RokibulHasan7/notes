# Docker Info

- In Docker, a DockerFile is used to build the image using **build command** and that image is stored into the registry using **push commmand**.
- When you run **pull command**, the Docker image(NGNIX) is retrieved from the registry.
- Finally, a single container(NGNIX) is built using Docker Image through the **run command**.

### Dockerfile, Docker Image
- Dockerfile is a simple text file which consists of instructions to build Docker images.
- A Docker Image is a template of instructions which is used to create Docker containers.
- Docker Image is comprised of multiple layers and by default it starts with a base layer.
- Each image layer depends on the layer below it.
- Image layers are created by executing each command in the Dockerfile.
- This image consists of only read-only format.

### Example of Dockerfile

```azure
FROM ubuntu:18.04
PULL ./file
RUN make /file 
CMD python /file/file.py
```
This code will give Image with four layers.

Layer 1 -> FROM ubuntu:18.04

Layer 4 -> CMD python /file/file.py

Here, FROM - Creates a layer from the ubuntu:18.04.

PULL - Adds files from your Docker repository.

RUN - Builds your container.

CMD - Specifies what command to run within the container.

### Docker Commands

- ENTRYPOINT -> allows to specify a command along with the parameters.
    ```
        ENTRYPOINT application "arg,arg1"  ---> Syntax
  
        ENTRYPOINT echo "Hello, $name"
    ```

- ADD -> command helps in copying data into a Docker image.
    ```
        ADD /[source]/[destination]  --> Syntax
        
        ADD /root_folder/test_folder
    ```

- ENV -> provides default values for variables which can be accessed within the container.
    ```
        ENV key value --> Syntax
  
        ENV value_1
    ```
  
- MAINTAINER -> declares the author field of the images.
    ```
        MAINTAINER [name]   ---> Syntax
  
        MAINTAINER author_name
    ```
  
### Docker Daemon
- Docker Daemon interacts with the operating system in order to create or manage Docker Containers.

### Docker Swarm
- Docker Swarm is a service which allows users to create and manage a cluster of Docker nodes and **schedule container**.
- Each node of a Docker Swarm is a Docker daemon and all Docker daemons interact using the Docker API.
- Here services can be deployed and accessed by nodes of same cluster.
- Docker Swarm can reschedule containers on Node failures.
- Swarm node has a backup folder, in case the main node fails, it can be used to restore the data onto a new swarm.

### Feature of Docker Swarm
- Decentralized access
- High security
- Auto load balancing
- High scalability
- Roll back a task

### More about Docker Swarm
- Containers are launched using services.
- A service is a group of containers of the same image.
- Service enables to scale your application.
- Before you can deploy a service in Docker Swarm, you must have at least one node deployed.
- There are two types of nodes in Docker Swarm.
  1. Manager node - Manager node maintains cluster management tasks.
  2. Worker node - Worker nodes receive and execute tasks from manager node.
  
### How does Docker Swarm work?
- Manager node knows the status of all the worker nodes in a cluster.
- Worker nodes accept tasks sent from manager node.
- Every worker node has as an agent, which reports on the state of the node's tasks to the manager.
- The worker nodes communicate with the manager node using API over HTTP.
- In Docker Swarm, services can be deployed and accessed by any node of same cluster.
- While creating a service, a user has to specify which container image to use.
- Here, a service is either global or replicated.
- A global service will run on every Swarm node.
- In a replicated service, the manager node distributes tasks to worker nodes.
- **A service is a description of a task or the state, whereas a task does the work.**
- Docker enables a user to create services, which can start tasks.
- When a task is assigned to a node, it cannot be assigned to another node.
- **It is possible to have multiple manager nodes on Swarm, but there will be only one primary manager node, which gets elected by the other manager nodes.**

### Manager node and Worker node
#### Manager node
- API - Based on the CLI command a service is created.
- Orchestration - Creates tasks for each service.
- Task Allocation - Allocate IP address (of worker node) to tasks.
- Dispatcher and Scheduler - Assigns and Instructs worker nodes to run a task.

#### Worker node
- Checks for the task - Connects to manager node and checks for new tasks.
- Execute the task - Execute the assigned tasks.

 