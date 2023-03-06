# Chart Template Guide

- Helm charts are structured like this:
  ```
    mychart/
        Chart.yaml
        values.yaml
        charts/
        templates/
  ```
  The ```templates/``` directory is for template files. When helm evaluates a chart, it will send all of the
  files in the ```templates/``` directory through the template rendering engine. It then collects the results
  of those templates and sends them on to Kubernetes.
- To create a chart: ```helm create <chart-name>```
- ```_helpers.tpl```: A place to put template helpers that you can re-use throughout the chart.
- ```helm get manifest <chart-name>``` command takes a release name and prints out all of the Kubernetes resources that 
  were uploaded to the server.
- ```helm uninstall <chart-name>```: To uninstall the release.
- Hard-coding the ```name:``` into a resource is usually considered to be bad practice. **Names should be
  unique to a release.**
- When you want to test the template rendering, but not actually install anything, you can use ```helm install --debug --dry-run <chart-name> <chart-path>```.
  This will render the templates. But instead of installing the chart, it will return the rendered template
  to you so you can see the output.
- Using ```--dry-run``` will make it easier to test the code, but it won't ensure the Kubernetes itself
  will accept the templates you generate. It's best not to assume that chart will install just because
  ```--dry-run``` works.

## Built-in Objects

- ```Release.Setvice```: The service that is rendering the present template. On helm, this is always ```Helm```.
- Built-in Objects:
  - Release
  - Values
  - Chart
  - Files
  - Capabilities
  - Template
- ```Values``` object provides access to values passed into the chart.
- If we need to delete a key from the default values, we may override the value of the key to be ```null```,
  in which case Helm will remove the key from the overridden values merge.

## Template Functions and Pipelines

- Drawing on a concept from UNIX, pipelines are a tool for chaining together a series of template
  commands to compactly express a series of transformations. In other words, **pipelines are an
  efficient wat of getting several things done in a sequence**.
- Example of pipeline: ```{{ .Values.favorite.food | upper | quote }}```. By this ```food``` value will
  be uppercase letter.
- One function frequently used in templates is the ```default``` function: ```default DEFAULT_VALUE GIVEN_VALUE```.
  This function allows us to specify a default value inside of the template, in case the value is omitted.
  Example: ```{{ .Values.favorite.drink | default "tea" | quote }}```. 

  
## Subcharts and Global Values

- Few important details about application subcharts:
  - A subchart is considered "stand-alone", which means a subchart can never explicitly depend on its
    parent chart.
  - For that reason, a subchart cannot access the values of its parent.
  - A parent chart can override values of subcharts.
  - Helm has a concept of ```global values``` that can be accessed by all charts.
- Global values are values that can be accessed from any chart or subchart by exactly the same name.
  Globals require explicit declaration. The Values data type has a reserved section called ```Values.global```
  where global values can be set.
- One advantage of using ```inlcude``` is that ```include``` can dynamically reference templates. Example: ```{{ include $mytemplate }}```.

## Debugging Templates

- There are few commands that can help you debug:
  - ```helm lint``` is your go-to tool for verifying that your chart follows best practices.
  - ```helm template --debug``` will test rendering chart templates locally.
  - ```helm install --dry-run --debug:```. It's a great way to have the server render your templates, 
    then return the resulting manifest file.
  - ```helm get manifest``` This is a good way to see what templates are installed on the server.
