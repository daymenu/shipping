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
func (c *Container) Create(ctx context.Context, container *pb.Container, resp *pb.Response) error {
	cm := model.ContainerModel{DB: c.DB}
	cc, err := cm.Create(container)
	if err != nil {
		return err
	}
	resp.Container = cc
	return nil
}

// Update 修改一个集装箱
func (c *Container) Update(ctx context.Context, container *pb.Container, resp *pb.Response) error {
	cm := model.ContainerModel{DB: c.DB}
	cc, err := cm.Update(container)
	if err != nil {
		return err
	}
	resp.Container = cc
	return nil
}

// Get 获取集装箱
func (c *Container) Get(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	cm := model.ContainerModel{DB: c.DB}
	container, err := cm.Get(req)
	if err != nil {
		return err
	}
	resp.Code = 200
	resp.Container = container
	return nil
}

// Use 使用集装箱
func (c *Container) Use(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	cm := model.ContainerModel{DB: c.DB}
	containers, err := cm.Use(req)
	if err != nil {
		return err
	}
	resp.Code = 200
	resp.Containers = containers
	return nil
}

// Page 获取集装箱
func (c *Container) Page(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	cm := model.ContainerModel{DB: c.DB}
	containers, err := cm.Page(req)
	if err != nil {
		return err
	}
	resp.Code = 200
	resp.Containers = containers
	return nil
}

// GiveBack 归还集装箱
func (c *Container) GiveBack(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	cm := model.ContainerModel{DB: c.DB}
	if err := cm.GiveBack(req.Containers); err != nil {
		return err
	}
	resp.Code = 200
	return nil
}
