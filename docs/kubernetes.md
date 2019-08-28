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
# 一个简单的例子
## 创建 mysql rc
[mysql-rc.yaml](kubeyaml/mysql-rc.yaml)
```yaml
apiVersion: v1
# 副本控制器
kind: ReplicationController
metadata:
# RC 名称
    name: mysql
spec:
# POD 副本期待数量
    replicas: 1
# 符合目标的Pod 拥有此标签
    selector:
        app: mysql
# 根据此模板创建POD的副本
    template:
        metadata:
# Pod 副本拥有的标签，对应 RC的Selector
            labels:
                app: mysql
        spec:
# pod 容器定义部分
            containers:
# 容器名称
                -   name: mysql
# 容器对应的docker-image
                    image: mysql
# 容器暴露的端口号
                    ports:
                        -   containerPort: 3306
# 设置容器的环境变量
                    env:
                        -   name: MYSQL_ROOT_PASSWORD
                            value: "123456"
```
1. 运行 mysql rc
```shell
sudo kubectl create -f mysql-rc.yaml
```
2. 查看运行状态
```
sudo kubectl get rc
sudo kubectl get pods
docker ps |grep mysql
```
3. 删除 mysql rc
```shell
sudo kubectl delete -f mysql-rc.yaml
```
## 创建 mysql service
[mysql-svc.yaml](kubeyaml/mysql-svc.yaml)
```yaml
apiVersion: v1
kind: Service
metadata:
## 服务名
    name:   mysql
spec:
## Service 的虚拟端口
    ports:
        -   port:   3306
## 确定哪些POD副本对应到本服务
    selector:
        app:    mysql
```
1. 查看运行的服务
```shell
sudo kubectl get service
```
输出 
|NAME      |   TYPE        |CLUSTER-IP    |  EXTERNAL-IP   |PORT(S)  |  AGE|
|:---:|:---:|:---:|:---:|:---:|:---:|
|kubernetes   |ClusterIP   |10.96.0.1  |     <none>      |  443/TCP  |  16h|
|mysql   |     ClusterIP   |10.108.38.254  | <none>        |3306/TCP  | 10s|

2. 通过mysql客户端连接刚刚创建的服务
```
mysql -uroot -h10.108.38.254 -p123456
```

# 基本概念
## Master
Kubernetes里的Master指的是集群控制节点
## Node
除了Mater，Kubernetes集群中的其他机器被称为Node节点。
```shell
# 查看 nodes
sudo kubectl get nodes
```
## Pod
Pod 是Kubernetes的最重要也最基本的概念，运行用户容器的基本单位。pause容器共享 ip及volume
## Label
可以附加到各种资源对象上，可以分组管理
## RC
它定义了一个期望的场景，即声明某种Pod的副本数量在任意时刻都符合某个预期值。
## Deployment
- 创建一个Deployment对象来生成对应的Replica Set 并完成Pod副本的创建过程
- 检查Deployment的状态来看部署动作是否完成
- 更新Deployment来创建新的pod
- 回退旧的Deployment
## Service
Service 其实就是我们说的微服务架构中的微服务
##  Volume
Volume 是Pod中能够被多个容器访问的共享目录。
- emptyDir 一个emptyDir Volume 是在pod分配到Node时创建的。从它的名称就可以看出，它的初始内容为空。pod移除时也会被移除
```yaml
spec:
    volumes:
        - name: dataVol
            emptyDir: {}
    containers:
        volumeMounts:
            - mountPath: /mydata-data
                name: datavol
```
- hostPath hostPath为在pod上挂载宿主机上的文件或目录
1. 容器应用程序生成的日志文件需要永久保存
2. 需要访问宿主机上docker引擎内部数据结构的容器应用时，可以 通过 定义 hostPah为宿主机/var/lib/docker目录
```yaml
volumes:
    -   name: "persistent-storage"
        hostPath:
            path: "/data"
```
## Persistent Volume
Kubernets 可以理解为某个网络存储中对应 的一块存储，
```
apiVersion: v1
kind: PersistentVolume
metadata:
    name: pv003
spec:
    capacity:
        storage: 5Gi
    accessModes:
        - ReadWriteOnce
    nfs:
        path: /data
        server:172.17.0.2
```
accessModes:
ReadWriteOnce
ReadOnlyMany
ReadWriteMany
## pvc
```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
    name: myclaim
spec:
    accessModes:
        - ReadWriteOnce
    resources:
        requests:
            storage: 8Gi
```
pod中volume定义
```yaml
volumes:
    -   name: mypd
        persistentVolumeClaim:
            claimName: myclaim
```
## Namespace （命名空间）
```
apiVersion: v1
kind:   namespace
metadata:
    name: development
```