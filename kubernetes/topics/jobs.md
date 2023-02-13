# Jobs

- A Job creates Pods that run until successfully termination (for instance, exit with 0).
- A regular Pod will continually restart regardless of its exit code.
- Jobs are a low-level primitive and can be used directly for simple workloads.
- Jobs are useful for things you only want to do once, such as database migrations or batch jobs. If run as a regular
  Pod, your database migration task **would run in a loop**, continually repopulating the database after every exit.
- The Job object is responsible for creating and managing Pods defined in a template in the job specification. These 
  Pods generally run until successful completion. ```The Job object coordinates running a number of Pods in parallel.```
- If the Pod fails before a successful termination, the job controller will create a new Pod based on the Pod template
  in the job specification.
- There is a small chance that duplicate Pods will be created for a specific task during certain failure scenarios.
- The Job object will automatically pick a unique label and use it to identify the Pods it creates.
- In advanced scenarios users can choose to turn off this automatic behavior and manually specify labels and selectors.
- By setting ```restartPolicy: Never```, we are telling the kubectl not to restart the Pod on failure.
- It's a good practice to set ```restartPolicy: OnFailure```, so that failed Pods are return in place.
- 

## Job Patterns

- Jobs are designed to manage batch-like workloads where work items are proceed by one or more Pods.
- By default, each job runs a single Pod once until successful termination.
- The Job pattern is defined by two primary attributes of a Job: ```the number of job completions and the number of Pods to run in parallel```
- In the case of the "run once until completion" pattern, the ``completions`` and ```parallelism``` parameters are set to 1.
- Job patterns:
    - One shot - A single Pod running once until successful termination.
    - Parallel fixed completions - One or more Pods running one or more times until reaching a fixed completion count.
    - Work queue: parallel jobs - One or more Pods running once until successful termination.

  
## Work Queues

- A common use case for jobs is to process work from a work queue. In this scenario, some task creates a number of work
  items and publishes them to a work queue. A worker job can be run to process each work item until the work queue is empty.
- Once the first Pod exits with a zero exit code, the job will start winding down and will not start any new Pods. This
  means that none of the workers should exit until the work is done and they are all in the process of finishing up.
- As the queue empties, the consumer Pods will exit cleanly and the ```consumers``` job will be considered complete.

## CronJobs

- Sometimes you want to schedule a job to be run at a **certain interval**. To achieve this, you can declare a CronJob in
  Kubernetes, which is responsible for **creating a new Job object at a particular interval**.
- 

## Commands

- Create a One-shot job by using the command-line tool.
```
kubectl run -i oneshot \
--image=gcr.io/kuar-demo/kuard-amd64:blue \
--restart=OnFailure \
--command /kuard \
-- --keygen-enable \
    --keygen-exit-on-complete \
    --keygen-num-to-gen 10
```

Here -i option to kubectl indicates that this is an interactive command. <br>
```--restart=OnFailure``` is the option that tells kubectl to create a Job object. <br>
All of the options after -- are command-line arguments to the container image. <br>

- See the Jobs.
```
kubectl get jobs -a
```

- Using labels we can clean up the stuffs we created.
```
kubectl delete <label,lable,label> -l <label>
```
