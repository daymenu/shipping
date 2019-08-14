package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"

	pb "github.com/daymenu/shipping/api/proto/container"
	api "github.com/micro/go-micro/api/proto"
)

// Container 结构体
type Container struct{}

// Get 实现方法
func (c *Container) Get(ctx context.Context, req *api.Request, rep *api.Response) error {
	log.Print("received container get request")
	log.Print(req.Get)
	return nil
}
func main() {
	srv := micro.NewService(
		micro.Name("daymenu.shippping.api.container"),
		micro.Version("latest"),
	)
	srv.Init()
	c := Container{}
	pb.RegisterContainerServiceHandler(srv.Server(), &c)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
