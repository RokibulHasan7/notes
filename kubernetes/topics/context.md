# Context

- A ```context``` is a combination of three key pieces of information: Cluster, User, Namespace.
- Each context represents a specific configuration for the ```kubectl``` command-line tool to interact with a K8s cluster.
- Contexts are helpful when you work with multiple Kubernetes clusters or environments.
  By switching contexts, you can easily switch between different cluster configurations without modifying the kubectl commands themselves.
- ```kubectl config current-context```
- ```kubectl config get-contexts```
- ```kubectl config use-context <context-name>```
- ```kubectl config set-context <context-name> --cluster=<cluster-name> --user=<user-name> --namespace=<namespace>```
- ```kubectl config delete-context <context-name>```
