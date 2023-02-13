# ConfigMap

- Kubernetes object that defines a small filesystem.
- It can be used when defining the environment or command line for your container.
- ConfigMap is combined with the Pod right before it is run.
- The container image and the POd definition can be reused by many workloads just by changing the ConfigMap that is used.

# Secrets

- Kubernetes Secrets API provides an application-centric mechanism for exposing
  sensitive configuration information to applications in a way that’s easy to audit and
  leverages native OS isolation primitives.
- Kubernetes Secrets are stored in plain test in the **etcd** storage for the cluster.

## Create Secret

- Secrets are created using the Kubernetes API or the kubectl command-line tool.
- Secrets hold one or more data elements as a collection of key/value pairs.
- Command to create a secret ``` kubectl create secret <secret-name> ```.

## Consuming Secrets

- **Secret volume** can be used to access Secret.
- Secret data can be exposed to Pods using the Secret volume type.
- **Secret volume** are managed by the **kubelet** and are created at Pod creation time.
- Secrets are stored on ```tmpfs volumes (aka RAM disks) ```.
- Each data element of a secret is stored in a separate file under the target mount point specified in the volume mount.

## Private Container Registries

- A special use case for secrets is to store access credentials for private container registries.

## Commands

- Create a configmap from commandline.
```
kubectl create configmap my-config \
--from-file=my-config.txt \
--from-literal=extra-param=extra-value \
--from-literal=another-param=another-value
```

Get this as a YAML file.
```
kubectl get configmaps my-config -o yaml
```

- create secret.
```
kubectl create secret generic kuard-tls \
--from-file=kuard.crt \
--from-file=kuard.key
```