# docker
## 安装docker
### 下载公用软件

 ```shell
 sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
 ```

### 添加docker源

```shell
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo apt-key fingerprint 0EBFCD88

sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
```

###  查看版本
```shell
apt-cache madison docker-ce
```

###  安装
```shell
sudo apt-get install docker-ce docker-ce-cli containerd.io
```

### 卸载
```shell
sudo apt-get purge docker-ce // 卸载
sudo rm -rf /var/lib/docker // 清除程序docker运行时文件
```

## docker 入门
### login 登录
> 登录到docker hub，类似github 只不过里边放的镜像

```
docker login
```
### pull 拉取镜像
```shell
docker pull nginx // 拉取nginx镜像
```
### images 查看本地镜像
```shell
docker images
docker images nginx
```
- -a 查看所有镜像
- -q仅显示数字IDs
### rmi 删除镜像
```shell
docker rmi nginx
docker rmi $(docker images -q)// 删除所有镜像
```
### run 运行容器
```shell
docker run -d -p 8080:80 --rm nginx
```
- -d 容器在后台运行
- -p 端口映射 8080映射到容器的80端口
--rm 容器停止后立即删除该容器
### ps 查看容器
```shell
docker ps // 查看正在运行的容器
docker ps -a // 查看所有容器
```
### exec 进入容器
```shell
docker exec -it 0ea20ae979a8 /bin/bash // -it 后面是docker ps查看的容器ID
```
### stop 停止容器
```shell
docker stop 0ea20ae979a8
```
### rm删除容器
```shell
// 删除制定容器ID的容器
docker rm 0ea20ae979a8

//删除所有已停止的容器
docker rm $(docker ps -qa) 
```
## Dockerfile
> 目标：使用容器编译并运行一个go程序

### 简单的go程序
#### 在主机新建main.go内容如下
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("hello docker")
    fmt.Println("hello ", os.Getenv("GO-NAME"))
}
```
#### 编写dockerfile
```shell
# 设置要拉取的镜像
FROM golang:latest AS gotest

# 设置环境变量
ENV GO-NAME test

# 运行一个命令
RUN mkdir -p /app

# 切换当前目录
WORKDIR /app

# 将主机的main.go复制到容器的当前目录下
ADD main.go .
RUN go get github.com/go-sql-driver/mysql
RUN go build -o godocker main.go

# 设置部署容器镜像的实例时要运行的命令
CMD ["./godocker"]

```
#### 通过dockerfile生成镜像
```shell
docker build . -t gotest
```
> . Dockerfile所在的目录 -t 给镜像命名

#### 运行镜像
```shell
docker run --rm gotest
```
# docker-compose
## 安装
```shell
//下载
sudo curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
//给可执行权限
sudo chmod +x /usr/local/bin/docker-compose
```
## docker-compose.yaml
> 目标：使用go程序并且使用一下mariadb

1. 在目录下新建docker-compose.yaml内容如下
```yaml
# 指定版本
version: '3'
# 定义服务
services:
#服务名字
 mariadb-service:
 # 镜像名字
  image: mariadb:latest
  # 容器名字
  container_name: mariadb.host
  ports:
   - 3306:3306
  #设置环境变量
  environment:
   MYSQL_ROOT_PASSWORD: 123456
 go-service:
  build: .
  links: ["mariadb-service"]
  depends_on: ["mariadb-service"]
  environment:
   MYSQL_HOST: mariadb.host
```
2. 在当前目录下新建 main.go，内容如下
```go
// 连接mariadb 并且输出mysql数据库中的所有表
package main

import (
	"database/sql"
	"fmt"
    "log"
    "os"

	_ "github.com/go-sql-driver/mysql" // mysql 驱动
)

func main() {
    host := os.Getenv("MYSQL_HOST")
    driver := fmt.Sprintf("root:123456@tcp(%s:3306)/mysql",host)

	db, err := sql.Open("mysql", driver)
	if err != nil {
		log.Fatal(err)
	}

	if db.Ping() != nil {
		log.Fatal(err)
	}
	log.Print("connect mysql success")

	rows, err := db.Query("show tables;")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tables")
	for rows.Next() {
		var table string
		rows.Scan(&table)
		fmt.Println(table)
	}
}
```
3. 在当前目录下新建Dockfile文件内容如下
```shell
# 设置要拉取的镜像
FROM golang:latest AS godb

# 运行一个命令
RUN mkdir -p /app

# 切换当前目录
WORKDIR /app

# 将主机的main.go复制到容器的当前目录下
ADD main.go .

RUN go get github.com/go-sql-driver/mysql
RUN go build -o godb main.go

# 设置部署容器镜像的实例时要运行的命令
CMD ["./godb"]

```
4. 运行
```shell
// 编译
docker-compose build
// 运行
docker-compose up
-d 可以让容器在后台运行
// 停止服务
docker-compose stop
// 删除已经停止的容器
docker-compose rm 
```