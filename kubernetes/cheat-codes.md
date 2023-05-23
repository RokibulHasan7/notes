# Cheat Codes

- Delete the ```first``` namespace which is stuck in ```terminating``` state.
    ```
    NS=`kubectl get ns |grep Terminating | awk 'NR==1 {print $1}'` && kubectl get namespace "$NS" -o json   | tr -d "\n" | sed "s/\"finalizers\": \[[^]]\+\]/\"finalizers\": []/"   | kubectl replace --raw /api/v1/namespaces/$NS/finalize -f -
    ```
- Namespace deleted but Pod is still stuck "Terminating" state.
    ```
    kubectl delete pod <PODNAME> --grace-period=0 --force --namespace <NAMESPACE>
    ```
  