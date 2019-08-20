package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/daymenu/shipping/user/handler"
	"github.com/daymenu/shipping/user/token"

	"github.com/daymenu/shipping/user/model"

	pb "github.com/daymenu/shipping/user/proto/user"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
)

func main() {
	db, err := model.CreateConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(&pb.User{}).ModifyColumn("password", "varchar(1000)")

	// user := &model.UserModel{DB: db}
	tokenService := &token.TokenService{}
	srv := micro.NewService(
		micro.Name("daymenu.shipping.srv.user"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWapper),
	)

	srv.Init()
	pb.RegisterUserServiceHandler(srv.Server(), &handler.User{
		DB:           db,
		TokenService: tokenService,
	},
	)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
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
		url := meta["Url"]

		whiteList := map[string]bool{"/api/user/login": true}
		fmt.Println(url, whiteList[url])
		if white := whiteList[url]; !white {
			fmt.Println("no white:", url)
			if tokenStr == "" {
				return fmt.Errorf("please input token")
			}
			tokenService := token.TokenService{}
			custom, err := tokenService.Decode(tokenStr)
			if custom.User.GetId() == 0 {
				return fmt.Errorf("login error:%s", err.Error())
			}
		}
		err := fn(ctx, req, resp)
		return err
	}
}
