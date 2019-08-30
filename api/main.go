package main

import (
	"log"
	"os"

	"github.com/daymenu/shipping/api/app"
	api "github.com/daymenu/shipping/api/proto/container"
	userapi "github.com/daymenu/shipping/api/proto/user"
	vesselapi "github.com/daymenu/shipping/api/proto/vessel"
	pb "github.com/daymenu/shipping/container/proto/container"
	userPb "github.com/daymenu/shipping/user/proto/user"
	vesselPb "github.com/daymenu/shipping/vessel/proto/vessel"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"

	_ "github.com/micro/go-plugins/registry/kubernetes"
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

	c := app.Container{}
	user := app.User{}
	vessel := app.Vessel{}
	// 建立container 服务的客户端
	c.Service = pb.NewContainerServiceClient("daymenu.shipping.srv.container", srv.Client())
	user.Service = userPb.NewUserServiceClient("daymenu.shipping.srv.user", srv.Client())
	vessel.Service = vesselPb.NewVesselServiceClient("daymenu.shipping.srv.vessel", srv.Client())
	// 定义服务
	apiSrv := micro.NewService(
		micro.Name("daymenu.shippping.api.api"),
		micro.Version("latest"),
	)
	// 初始化服务
	apiSrv.Init()
	api.RegisterContainerServiceHandler(apiSrv.Server(), &c)
	userapi.RegisterUserHandler(apiSrv.Server(), &user)
	vesselapi.RegisterVesselHandler(apiSrv.Server(), &vessel)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
