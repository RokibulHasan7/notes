## Commands

- Building the image
    ```
        docker build -t <image-name> .
    ```
  
- Run the docker image
    ```
        docker run -d -p <port-num>:<port-num> <image-id>
    ```

- To see the list of docker image
    ```
        docker image ls
    ```
  
- Find running containers
    ```
        docker container ls
    ```
  
- Stop the container
    ```
        docker container stop <container-id>
    ```
- Push the Image to the DockerHub.
    ```
        docker push <DockerHub-Username>/<repo-name>:<TAG-of-image>
    ```
  
  - Remove an Docker Image
       ```
            docker rmi <id>
            docker rmi <repository> : <tag>
            
            docker rmi -f <id>
       ```
    
- Display logs of a container
    ```
        docker logs <container-name>
    ```

- Scan docker image
  ```
      docker scan <image-name>
  ```
  


- Usefull Links
  https://www.digitalocean.com/community/tutorials/how-to-remove-docker-images-containers-and-volumes