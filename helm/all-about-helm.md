# Helm

- It is ```package manager for Kubernetes```, to package YAML files and distribute them in public and private repositories.
- Helm helps you to streamline installation and management of Kubernetes application.
- Helm uses a packaging format called charts.

## Helm Charts
- A chart is a collection of files that describe a related set of Kubernetes resources.
- Helm Charts are simply Kubernetes YAML manifests combined into a single package that can be advertised to your
  Kubernetes clusters.
- A single chart might be used to deploy something simple, like a Nginx pod, or something complex, like a full web
  app stack with HTTP server, databases.
- Helm Charts help you define, install, and upgrade even the most complex Kubernetes application.
- Bundle of YAML files.
- Create your own Helm Charts with Helm.
- Push them to Helm repositories.
- Download and use existing ones.


## Three Big Concept

- A ```Chart``` is a Helm package. It contains all of the resource definitions necessary to run an applicaion,
  tool, or service inside of a Kubernetes cluster.
- A ```Repository``` is the place where charts can be collected and shared.
- A ```Release ``` is an instance of a chart running in a Kubernetes cluster. One chart can often be installed
  many times into the same cluster. And each time it is installed, a new release is created. Consider a MySQL
  chart, if you want two databases running your cluster, you can install that chart twice. Each one will have
  its own release, which will in turn have its own release name.

## Helm Chart Structure

```
    mychart/
        Chart.yaml
        values.yaml
        charts/
        templates/
```
Top level ```mychart``` folder -> name of chart. <br>
Chart.yaml -> meta info about chart. Example: name, version number, dependencies etc. <br>
values.yaml -> values for the template files. Default values that you can override. <br>
chart folder -> chart dependencies. It contains other dependant charts. <br>
template folder -> the actual template files. Where you put the actual manifest you are deploying with the
chart. For example, you might be deploying an nginx deployment that need a service, configmap and secrets.<br>
When give command ```helm install <chartname>``` template files will be filled with the values from values.yaml. <br>

## Value injection into template files:

values.yaml (default)
```
imagename: myapp
port: 8080
version: 1.0.0
```
my-values.yaml (override values)
```
version: 2.0.0
```
result (.Values object)
```
imagename: myapp
port: 8080
version: 2.0.0
```
```helm install --values=my-values.yaml <chartname>```

Or on Command line: ```helm install --set version=2.0.0```


## Helm Commands

- For installing Helm charts
```helm install```

- Search for a chart in repository
```helm search```

- List all the deployed releases.
```helm list```

- Upgrade your releases with new version.
```helm upgrade```

- Delete the release and all deployed resources.
```helm delete```

