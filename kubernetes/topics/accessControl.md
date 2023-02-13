# Role-Based Access Control for Kubernetes

- Role-based access control provides a mechanism for restricting both access to and actions on Kubernetes APIs to
  ensure that only authorized users have access.
- It's important to remember that anyone who can run arbitrary code inside the Kubernetes cluster can effectively obtain
  root privileges on the entire cluster.
- Every request to Kubernetes is first ```authenticated```.
- Kubernetes does not have a built-in identity store, focusing instead on integrating other identity sources within itself.
- Every request to Kubernetes is associated with some identity.
- A request with no identity is associated with the ```system:unauthenticated``` group.
- Kubernetes makes a distinction between user identities and service account identities. Service accounts are 
  created and managed by Kubernetes itself and generally associated with components running inside the cluster.
  User accounts are all other accounts associated with actual users of the cluster, and often include automation
  like continuous delivery services that run outside the cluster.
- Kubernetes supports a number of authentication providers:
  - HTTP Basic Authentication (largely deprecated).
  - x509 client certificates.
  - Static token files on the host.
  - Cloud authentication providers, such as Azure Active Directory and AWS Identity and Access Management(IAM).
  - Authentication webhooks.
- ```You should always use different identities for different applications in your cluster.``` For example, you 
  should have one identity for your production frontends, a different identity for the production backends, and all
  production identities should be distinct from development identities.
- ```You should also have different identities for different clusters.```
- For example, Azure Active Directory supplies an **open source identity provider for Pods** as do other popular identity
  providers.
- When the Kubernetes API server start up, it automatically installs a number of default ClusterRoles that are defined
  in the code of the API server itself. This means that if you modify any built-in cluster role, those modifications
  are transient. Whenever the API sever is restarted your changes will be overwritten.
  - To Prevent this from happening, before you make any other modifications, you need to add the ```rbac.authorization.kubernetes.io/autoupdate```
    annotation with a value of ```false``` to the built-in ClusterRole resource. If this annotation is set to ```false```,
    the API server will not overwrite the modified ClusterRole resource.
- Kubernetes RBAC supports the usage of an ```aggregation rule``` to combine multiple roles in a new role.
- A best practice for managing ClusterRole resources is to create a number of fine-grained cluster roles and then
  aggregate them to form higher-level or broader cluster roles.

## Role and Role Bindings

- Identity is just the beginning of authorization in Kubernetes. Once Kubernetes knows the identity of the request, it
  needs to determine if the request is authorized for that user. To achieve this, **it uses roles and role bindings.**
- A ```role``` is a set of abstract capabilities. For example, the ```appdev``` role might represent the ability to
  create Pods and Services. A ```role binding``` is an assignment of a role to one or more identities. Thus, binding
  the ```appdev``` role to the user identity ```alice``` indicates that Alice has the ability to create Pods and Services.
- In Kubernetes, two pairs of related resources represent roles and role bindings. One pair is scoped to a 
  namespace (Role and RoleBinding), while the other pair is scoped to the cluster (ClusterRole and ClusterBinding).
- Role resources are namespaced and represent capabilities within that single namespace.
- Four built-in roles designed for generic end users:
  - The ```cluster-admin``` role provides complete access to the entire cluster.
  - The ```admin``` role provides complete access to a complete namespace.
  - The ```edit``` role allows an end user to modify resources in a namespace.
  - The ```view``` role allows for read-only access to a namespace.

- Most clusters already have numerous ClusterRole bindings set up, and you can view these bindings with
  ```kubectl get clusterrolebindings```.


## Commmands

- Test whether a specific user can perform a specific action.
```
kubectl auth can-i create pods
```
Test sub-resources like logs or port-forwarding with the --subresource command-line flag:
```
kubectl auth can-i get pods --subresource=logs
```

- Get default clusterRoles.
```
kubectl get clusterroles
```

- Managing RBAC in source control.
```
kubectl auth reconcile -f <yaml-file>
```
```reconcile``` command operates like ```kubectl apply``` and will reconcile a set of roles and role bindings with
the current state of the cluster.
