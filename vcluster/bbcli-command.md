# Command

- ./bin/ace-linux-amd64 auth login --username=appscode --password=password
- ./bin/ace-linux-amd64 auth logout 

## Cluster flow

- ./bin/ace-linux-amd64 cluster list --access-token=hfhejh38423ufsf
- ./bin/ace-linux-amd64 cluster connect  --name=vcluster --credential=linode-cred --access-token=hfhejh38423ufsf
- ./bin/ace-linux-amd64 cluster check --provider=linode --name=vcluster --access-token=hfhejh38423ufsf 
- ./bin/ace-linux-amd64 cluster import --provider=linode --credential=linode-cred --id=134912 --name=vcluster --display-name=vcluster --access-token=hfhejh38423ufsf
- ./bin/ace-linux-amd64 cluster remove --name=test-cluster1 --all-features --access-token=hfhejh38423ufsf