# container api

推荐micro例子仓库： [examples](https://github.com/micro/examples.git)
*说明：例子有的还是过时了，请自己甄别*

## 编写proto
```go
// 定义版本
syntax = "proto3";
//定义服务名 
package daymenu.shippping.api.container;
//导入micro的api proto文件
import "github.com/micro/go-micro/api/proto/api.proto";
service ContainerService {
    // 访问路径 /container/containerService/page
	rpc Page(go.api.Request) returns(go.api.Response) {};
	rpc Get(go.api.Request) returns(go.api.Response) {};
	rpc Use(go.api.Request) returns(go.api.Response) {};
	rpc GiveBack(go.api.Request) returns(go.api.Response) {};
	rpc Create(go.api.Request) returns(go.api.Response) {};
	rpc Update(go.api.Request) returns(go.api.Response) {};
}
```

## 执行proto

```
protoc -I. --proto_path=${GOPATH}/src --go_out=plugins=micro:. ./proto/container/container.proto
```

## 编写main.go

```go
package main

import (
	"log"
	"os"

	"github.com/daymenu/shipping/api/app"
	api "github.com/daymenu/shipping/api/proto/container"
	pb "github.com/daymenu/shipping/container/proto/container"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

func main() {
	// 注册为consul
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			os.Getenv("CONSUL_HTTP_ADDR"),
		}
	})

	// 建立consul类型的服务
	srv := micro.NewService(micro.Registry(reg))

    //建立集装箱结构体
	c := app.Container{}

	// 建立container 服务的客户端
    c.Service = pb.NewContainerServiceClient("daymenu.shipping.srv.container", srv.Client())
    
	// 定义服务
	apiSrv := micro.NewService(
		micro.Name("daymenu.shippping.api.container"),
		micro.Version("latest"),
	)
	// 初始化服务
    apiSrv.Init()
    
    //注册服务
	api.RegisterContainerServiceHandler(apiSrv.Server(), &c)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
```

## container api实现
目录位置:/api/app/

```go

// Container 结构体
type Container struct {
	Service pb.ContainerServiceClient
}

// Get 实现方法
func (container *Container) Get(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	//初始成功
	resp.StatusCode = 200
	apiReq := APIRequest{request: req}
	id, err := apiReq.GetInt64("id")
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "api:请传入正确的容器ID")
		return nil
	}

	//调用微服务
	response, err := container.Service.Get(ctx, &pb.Request{
		Id: id,
	})

	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10001, "api:"+err.Error())
		return nil
	}

	ct, err := json.Marshal(response.GetContainer())

	resp.Body = string(ct)

	return nil
}
```

```Dockerfile
FROM debian

# MICRO 配置
ENV MICRO_ADDRESS=:50051
ENV MICRO_REGISTRY=consul

# app 
RUN mkdir /app  
WORKDIR /app  
ADD container-api /app/container-api

# 开启vessel微服务
CMD ["./container-api", "--registry_address=registry:8500"]
```

