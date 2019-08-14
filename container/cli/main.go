package main

import (
	"context"
	"fmt"
	"log"

	"os"

	"github.com/micro/go-micro"

	"github.com/micro/go-micro/registry"

	"github.com/micro/go-micro/registry/consul"

	pb "github.com/daymenu/shipping/container/proto/container"
)

func main() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			os.Getenv("CONSUL_HTTP_ADDR"),
		}
	})
	srv := micro.NewService(micro.Registry(reg))
	c := pb.NewContainerServiceClient("daymenu.shipping.srv.container", srv.Client())
	fmt.Println(c)
	rep, err := c.Create(context.Background(), &pb.Container{
		Id:         1,
		CustomerId: "54896",
		Origin:     "cz",
		UserId:     "0001",
		Height:     60,
		Width:      60,
		Long:       60,
	})
	if err != nil {
		log.Fatal(err)
	}
	if rep.Code == 1 {
		log.Print(rep)
	}
}
