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
