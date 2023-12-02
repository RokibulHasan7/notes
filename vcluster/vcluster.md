# Virtual cluster

## Commands

- Install vcluster cli
    - curl -L -o vcluster "https://github.com/loft-sh/vcluster/releases/latest/download/vcluster-linux-amd64" && chmod +x vcluster;
    - sudo mv vcluster /usr/local/bin;

- Create vcluster
    - vcluster create vcluster-1 
    - vcluster create vcluster-1 --expose
  
- Connect to vcluster
    - vcluster connect <vcluster name>

- Disconnect to vcluster
    - vcluster disconnect

- Delete vCluster.
    - vcluster delete my-vcluster
- 

