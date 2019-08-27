# 安装
## kubectl 安装
```shell
sudo curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl \
&& sudo chmod +x kubectl
&& sudo mv  kubectl /usr/local/bin/
```
## minikube 安装
```shell
sudo curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
&& sudo chmod +x minikube \
&& sudo mv minikube /usr/local/bin/
```

## 启动
```shell
sudo minikube start --vm-driver=none --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers

// 不出意外肯定会失败

// 浏览器下载以下两个文件
https://storage.googleapis.com/kubernetes-release/release/v1.15.2/bin/linux/amd64/kubeadm
https://storage.googleapis.com/kubernetes-release/release/v1.15.2/bin/linux/amd64/kubelet

// 然后把这两移动到 ~/.miniku/cache/v1.15.2/文件夹下，然后再执行启动命令就成功了
```