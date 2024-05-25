# pod-state-mysql



#Day1 

 kubectl 多配置集群管理



```
export KUBECONFIG=~/.kube/config:~/.kube/kubconfig2
kubectl config view --flatten > ~/.kube/merged_kubeconfig
```



不太好用，还是手动修改吧。

```
#一个访问凭证时#
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: ***
    server: https://192.168.*.*:6443
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: kubernetes-admin
  name: kubernetes-admin@kubernetes
current-context: kubernetes-admin@kubernetes
kind: Config
preferences: {}
users:
- name: kubernetes-admin
  user:
    client-certificate-data: ***
    client-key-data: ***


#两个访问凭证时#
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: ***
    server: https://42.194.*.*:443/
  name: cls-ec6ymsmo
- cluster:
    certificate-authority-data: ***
    server: https://192.168.*.*:6443
  name: kubernetes
contexts:
- context:
    cluster: cls-ec6ymsmo
    user: "10000******"
  name: cls-ec6ymsmo-10000******
- context:
    cluster: kubernetes
    user: kubernetes-admin
  name: kubernetes-admin@kubernetes
current-context: kubernetes-admin@kubernetes
kind: Config
preferences: {}
users:
- name: "100006621061"
  user:
    client-certificate-data: ***
    client-key-data: ***
- name: kubernetes-admin
  user:
    client-certificate-data: ***
    client-key-data: ***
```

