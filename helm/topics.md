# Charts

## The Chart File Structure

- A chart is organized as a collection of files inside of a directory. The directory name is the name of
  of the chart (without versioning information). Thus a chart describing WordPress would be stored in a
  ```wordpress/``` directory.
  ```
    wordpress/
        Chart.yaml
        LICENSE
        READMR.md
        values.yaml
        values.schema.json
        charts/
        crds/
        templates/
        templates/NOTES.txt
  ```
  
- The ```type``` field defines the type of chart. There are two types: ```application``` and ```library```.
- Application is the default type and it is the standard chart which can be operated on fully. The
  ```library chart``` provides utilities or functions for the chart builder. **A library chart differs
  from an application chart because it is not installable and usually doesn't contain any resource
  objects.** An application chart can be used as a library chart.
- In Helm, one chart may depend on any number of other charts. These dependencies can be dynamically
  linked using the ```dependencies``` field in ```Chart.yaml``` or brought in to the ```charts/``` directory and
  managed manually.
  ```
    dependencies:
      - name: apache
        version: 1.2.3
        repository: https://example.com/charts
      - name: mysql
        version: 3.2.1
        repository: https://another.example.com/charts
  ```
  Once you have defined dependencies, you can run ```helm dependency update``` and it will use your 
  dependency file to download all the specified charts into your ```chart/``` directory for you.
- ```helm create mychart``` - it will create a folder with basic files and folder.
