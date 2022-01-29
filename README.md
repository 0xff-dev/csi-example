# csi-example
Examples of implementing CSI related interfaces

# Usage
```shell
go mod download
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
docker build -t your-image .

# file deploy/deploy.yaml, you shole replace all 172.22.50.227/system_containers/csi-example:beta with your image
kubecel apply -f deploy
```

# Reference
[csi spec](https://github.com/container-storage-interface/spec/blob/master/spec.md#rpc-interface)  
[csi-driver-host-path](https://github.com/kubernetes-csi/csi-driver-host-path)  
[csi blog](https://blog.dianduidian.com/post/%E5%BC%80%E5%8F%91%E8%87%AA%E5%B7%B1%E7%9A%84csi%E5%AD%98%E5%82%A8%E6%8F%92%E4%BB%B6/)  