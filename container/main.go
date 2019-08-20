package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/server"

	"github.com/daymenu/shipping/container/handler"
	"github.com/daymenu/shipping/container/model"

	pb "github.com/daymenu/shipping/container/proto/container"
	pbUser "github.com/daymenu/shipping/user/proto/user"
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
		micro.WrapHandler(AuthWapper),
	)

	srv.Init()

	pb.RegisterContainerServiceHandler(srv.Server(), c)
	if err := srv.Run(); err != nil {
	}
}

// AuthWapper 验证登录
func AuthWapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		// Note this is now uppercase (not entirely sure why this is...)
		tokenStr := meta["Token"]

		// 注册为consul
		reg := consul.NewRegistry(func(op *registry.Options) {
			op.Addrs = []string{
				os.Getenv("CONSUL_HTTP_ADDR"),
			}
		})

		// 建立consul类型的服务
		srv := micro.NewService(micro.Registry(reg))

		userService := pbUser.NewUserServiceClient("daymenu.shipping.srv.user", srv.Client())
		validateToken, err := userService.ValidateToken(ctx, &pbUser.Token{
			Token: tokenStr,
		})
		if err != nil || !validateToken.Valid {
			return fmt.Errorf("login error:%s", err.Error())
		}
		err = fn(ctx, req, resp)
		return err
	}
}
