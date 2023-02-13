# Service Meshes

- When considering adopting a service mesh, you have to balance the complexity of adding a new component(generally
  provided by a third party) to your list of dependencies.
- There are **three general capabilities** provided by most service  mesh implementations: **network encryption and 
  authorization, traffic shaping, and observability.**
- Installing a service mesh on your kubernetes cluster automatically provides encryption to network traffic between
  every Pod in the cluster. The service mesh adds a sidecar container to every Pod, which transparently intercepts all
  network communication. In addition to securing the communication, mTLS(Mutual Transport Layer Security) adds
  identity to the encryption using client certificates so your application securely knows the identity of every network
  client.
- ```Dog-fooding``` - A new version of the software is tried internally before anywhere else. In a dog-fooding model,
  you may run version Y of your service for a day to a week (or longer) for a subset of users before you roll it out
  broadly to your full set of users.
- ```Automatic introspection``` is another important capability provided by a service mesh.
- ```The service mesh is implemented once for an entire cluster.```
- A service mesh is a distributed system that adds complexity to your application design. The service mesh is deeply
  integrated into the core communication of your microservices.
- ```When a service mesh fails, your entire application stops working.```
- For small applications, a service mesh is an unnecessary complexity.
- A part of the service mesh needs to be present within every one of your Pods.
- Any REST API request to create a Pod is first routed to admission controller. The service mesh admission
  controller modifies the Pod definition by adding the ```sidecar```. Because this admission controller is installed by the
  cluster administrator, it transparently and consistently implements a service mesh for an entire cluster.
- The best service mesh for you is likely the one that your cloud provider supplies for you.
- Service meshes contain powerful functionality that adds security and flexibility to your application. At the same
  time, a service mesh adds complexity to the operations of your cluster and is another potential source of outages
  of your application.