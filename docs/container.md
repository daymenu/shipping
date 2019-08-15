# 集装箱微服务

- 安装依赖库  

```shell
GOPROXY=https://goproxy.io go get -u github.com/micro/go-micro

GOPROXY=https://goproxy.io go get -u  github.com/jinzhu/gorm
```
- 目录规划
shipping 是项目文件
shipping>container为container微服务的文件夹  
以下代码以container目录为基础  
- 编写container.proto

```protobuf

syntax="proto3";

// 定义包名，和以后的服务名字一样的哦
package daymenu.shipping.srv.container;

// 集装箱服务
service ContainerService {
    rpc Create(Container) returns (Response){} // 创建集装箱
    rpc Update(Container) returns (Response){} // 修改集装箱
    rpc Get(Request) returns (Response){} // 获取集装箱
    rpc Use(Request) returns (Response){} // 使用集装箱
    rpc Page(Request) returns (Response){} // 集装箱列表
    rpc GiveBack(Request) returns (Response){} // 归还集装箱
}

message Container {
    int64 id = 1; // 编号
    string customer_id =  2; //集装箱所属客户编号
    string origin = 3; // 出发地
    string user_id = 4; //集装箱所属用户编号
    int64 height = 5; // 集装箱高度
    int64 width = 6; //集装箱宽度
    int64 long = 7; // 集装箱长度
    int32 status = 8; //1 可用 2 正在使用 3 已经报废
}

message Request {
    int64 height = 1; // 货物箱高度
    int64 width = 2; // 货物箱宽度
    int64 long = 3; // 货物箱长度
    int64 page = 4; // 几页
    int64 pageSize = 5; //每页几条
    int64 id = 6;//id
    string name = 7; // 集装箱名字
    repeated Container containers = 8;
}

message Response {
    int32 code = 1; // 200 成功
    repeated Container containers = 2;// 获取到的集装箱
    Container container = 3;//获取到的一个集装箱
}
```
- 使用protoc生成对应的go代码

```shell
protoc -I. --go_out=plugins=micro:.  ./proto/container/container.proto
```
- 查看生成的go代码，并关注其中的服务接口

```go
// Server API for ContainerService service

type ContainerServiceHandler interface {
	Create(context.Context, *Container, *Response) error
	Update(context.Context, *Container, *Response) error
	Get(context.Context, *Request, *Response) error
	Use(context.Context, *Request, *Response) error
	Page(context.Context, *Request, *Response) error
	GiveBack(context.Context, *Request, *Response) error
}
```

- 编写model
1. 新建model文件夹
2. 新建database.go文件，内容如下

```go
package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// gorm require
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// CreateConn 创建连接
func CreateConn() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	fmt.Println(host, name, user, password)
	return gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			name,
		),
	)
}
```

3. 新建container.go

```go
package model

import (
	"fmt"

	pb "github.com/daymenu/shipping/container/proto/container"
	"github.com/jinzhu/gorm"
)

// Container 定义Container接口
type Container interface {
	Get(*pb.Request) (*pb.Containers, error)
	Create(*pb.Container) error
	GiveBack(*pb.Containers) error
}

// ContainerModel 结构体
type ContainerModel struct {
	DB *gorm.DB
}

// Use 根据请求获取符合要求的集装箱，然后锁定
// ps: 查找体积大于请求的集装箱
func (cm *ContainerModel) Use(r *pb.Request) ([]*pb.Container, error) {
	var cs []*pb.Container
	var canBluck int64
	cIds := make([]int64, 10)
	bluk := r.Height * r.Width * r.Long
	// 下面代码可以好好优化
	for canBluck > bluk {
		var tcs []*pb.Container // 临时
		sql := "status=1"
		cm.DB.Where(sql).Limit(10).Find(&tcs)
		for _, item := range tcs {
			if item.Id == 0 {
				return nil, fmt.Errorf("not enough container")
			}
			canBluck += item.Height * item.Width * item.Long
			if canBluck > bluk {
				cs = append(cs, item)
				cIds = append(cIds, item.Id)
				break
			}
		}
	}

	if num := cm.DB.Where("id in (?)", cIds).Updates(map[string]int{"status": 2}).RowsAffected; num < 1 {
		return nil, fmt.Errorf("container: giveback faild")
	}
	return cs, nil
}

// Page 分页获取数据
func (cm *ContainerModel) Page(r *pb.Request) ([]*pb.Container, error) {
	var cs []*pb.Container
	rows, err := cm.DB.Where("name like ?", r.GetId()).Where("").Limit(r.Page).Rows()
	if err != nil {
		return cs, err
	}
	for rows.Next() {
		var c pb.Container
		rows.Scan(&c)
		cs = append(cs, &c)
	}
	return cs, nil
}

// Create 创建一个集装箱
func (cm *ContainerModel) Create(c *pb.Container) error {
	if err := cm.DB.Create(c).Error; err != nil {
		return err
	}
	return nil
}

// Update 创建一个集装箱
func (cm *ContainerModel) Update(c *pb.Container) error {
	if err := cm.DB.Create(c).Error; err != nil {
		return err
	}
	return nil
}

// Get 获取集装箱
func (cm *ContainerModel) Get(c *pb.Request) (*pb.Container, error) {
	var container pb.Container
	if err := cm.DB.Where("id=?", c.Id).First(&container).Error; err != nil {
		return nil, err
	}
	return &container, nil
}

// GiveBack 归还集装箱
func (cm *ContainerModel) GiveBack(cs []*pb.Container) error {
	cIds := make([]int64, len(cs))
	for _, item := range cs {
		cIds = append(cIds, item.Id)
	}
	if num := cm.DB.Where("id in (?)", cIds).Updates(map[string]int{"status": 1}).RowsAffected; num < 1 {
		return fmt.Errorf("container: giveback faild")
	}
	return nil
}
```

-----
- 编写 handler
1. 新建handler文件夹
2. 新建container.go文件

```go
package handler

import (
	"context"

	"github.com/daymenu/shipping/container/model"

	pb "github.com/daymenu/shipping/container/proto/container"
	"github.com/jinzhu/gorm"
)

// IContainer 定义接口
type IContainer interface {
	Create(context.Context, *pb.Container, *pb.Response) error
	Update(context.Context, *pb.Container, *pb.Response) error
	Get(context.Context, *pb.Request, *pb.Response) error
	Use(context.Context, *pb.Request, *pb.Response) error
	Page(context.Context, *pb.Request, *pb.Response) error
	GiveBack(context.Context, *pb.Containers, *pb.Response) error
}

// Container 结构体
type Container struct {
	DB *gorm.DB
}

// Create 创建一个集装箱
func (c *Container) Create(ctx context.Context, container *pb.Container, rep *pb.Response) error {
	cm := model.ContainerModel{DB: c.DB}
	if err := cm.Create(container); err != nil {
		return err
	}
	return nil
}

// Update 修改一个集装箱
func (c *Container) Update(ctx context.Context, container *pb.Container, rep *pb.Response) error {
	cm := model.ContainerModel{DB: c.DB}
	if err := cm.Create(container); err != nil {
		return err
	}
	return nil
}

// Get 获取集装箱
func (c *Container) Get(ctx context.Context, req *pb.Request, rep *pb.Response) error {
	cm := model.ContainerModel{DB: c.DB}
	container, err := cm.Get(req)
	if err != nil {
		return err
	}
	rep.Container = container
	return nil
}

// Use 使用集装箱
func (c *Container) Use(ctx context.Context, req *pb.Request, rep *pb.Response) error {
	cm := model.ContainerModel{DB: c.DB}
	containers, err := cm.Use(req)
	if err != nil {
		return err
	}
	rep.Containers = containers
	return nil
}

// Page 获取集装箱
func (c *Container) Page(ctx context.Context, req *pb.Request, rep *pb.Response) error {
	cm := model.ContainerModel{DB: c.DB}
	containers, err := cm.Page(req)
	if err != nil {
		return err
	}
	rep.Containers = containers
	return nil
}

// GiveBack 归还集装箱
func (c *Container) GiveBack(ctx context.Context, req *pb.Request, rep *pb.Response) error {
	cm := model.ContainerModel{DB: c.DB}
	if err := cm.GiveBack(req.Containers); err != nil {
		return err
	}
	return nil
}

```

- 新建 main.go

```go
package main

import (
	"github.com/micro/go-micro"

	"github.com/daymenu/shipping/container/handler"
	"github.com/daymenu/shipping/container/model"

	pb "github.com/daymenu/shipping/container/proto/container"
)

func main() {
	db, err := model.CreateConn()
	if err != nil {
	}
	db.AutoMigrate(&pb.Container{})
	c := &handler.Container{DB: db}
	srv := micro.NewService(
		micro.Name("daymenu.shipping.srv.container"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterContainerServiceHandler(srv.Server(), c)
	if err := srv.Run(); err != nil {
	}
}
```

-  编译
	GOPROXY=https://goproxy.io go build -o container .

- 编写container的Dockerfile

```dockerfile
FROM debian:latest

# MICRO 配置
ENV MICRO_ADDRESS=:50051
ENV MICRO_REGISTRY=consul
ENV DB_HOST=mariadb.host
ENV DB_NAME=shipping
ENV DB_USER=root
ENV DB_PASSWORD=123456

# app 
RUN mkdir /app  
WORKDIR /app  
ADD container /app/container

# 开启container微服务
CMD ["./container",  "--registry_address=registry:8500"]
```

- 编写docker-compose.yaml

```yaml
version: '3'
services:
# consul 注册中心
 consul:
  image: consul:latest
  hostname: registry
  container_name: registry
  ports:
  - "8300:8300"
  - "8400:8400"
  - "8500:8500"
  - "8600:53/udp"
  environment:
   GOMAXPROCS: 2
  # micro api 网关
 micro:
  command: api --handler=api
  image: microhq/micro:latest
  links:
    - consul
    - container-api
  ports:
    - "8080:8080"
  environment:
    MICRO_REGISTRY: consul
    MICRO_API_NAMESPACE: daymenu.shippping.api
    MICRO_REGISTRY_ADDRESS: registry:8500
 micro-web:
  command: web
  image: microhq/micro:latest
  links:
    - consul
  ports:
    - "8082:8082"
  environment:
    MICRO_REGISTRY: consul
    MICRO_REGISTRY_ADDRESS: registry:8500
 mariadb-service:
  image: mariadb:latest
  hostname: mariadb.host
  container_name: mariadb.host
  environment:
   MYSQL_ROOT_PASSWORD: 123456
  volumes: ["./volume/mysql:/data/msyql"]
 container-service:
  build: ./container
  ports: ["50051:50051"]
  links: ["mariadb-service", "consul"]
  depends_on: ["mariadb-service", "consul"]
 container-api:
  build: ./api
  ports: ["50052:50051"]
  links: ["consul"]
 container-cli:
  build: ./container/cli
  links: ["consul", "micro"]
  environment:
   CONSUL_HTTP_ADDR: registry:8500
  depends_on: ["mariadb-service", "consul", "container-service"]
```

- 编译docker

```shell
docker-compose build
```

- 运行

```shell
 docker-compose up -d
```

- 在本机访问localhost:8082
点击Registry就可以发现所有服务了