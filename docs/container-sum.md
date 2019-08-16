# containre 整体运行

跟着文档走的这一步的话你的集装箱的微服务就写好了。

## 编译

```shell
docker-compose build
```

## 运行

```shell
docker-compose up -d
```

## 建立数据库
```shell
docker exec  -it 容器id
mysql -u root -p123456
create database shipping
```

## 通过micro web 查看服务

```
linux 直接访问localhost:8082
docker 启动后显示
docker is configured to use the default machine with IP **192.168.99.100**
****中间的ip是我们要使用的ip
windows 访问 192.168.99.100：8082

点击Registry可以查看注册的微服务
访问 localhost:8050/container/containerService/get
就会发现{"Code":10000,"Msg":"api:请传入正确的容器ID"}
这是因为没有传入id的参数
```

###  程序的二进制文件也已提交，可以直接看展示效果
###  如果自己写程序编译的话有可能会碰到这样或这那样的问题，请查看错误信息，然后搜索解决
### 完成