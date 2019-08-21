package handler

import (
	"context"

	pb "github.com/daymenu/shipping/vessel/proto/vessel"
	"github.com/jinzhu/gorm"
)

// Vessel 结构体
type Vessel struct {
	DB *gorm.DB
}

// Create 创建货轮
func (v *Vessel) Create(ctx context.Context, vessel *pb.Vessel, resp *pb.Reponse) error {
	return nil
}

// Update 创建货轮
func (v *Vessel) Update(ctx context.Context, vessel *pb.Vessel, resp *pb.Reponse) error {
	return nil
}

// Page 列表
func (v *Vessel) Page(ctx context.Context, vessel *pb.Request, resp *pb.Reponse) error {
	return nil
}

// Get 获取一个
func (v *Vessel) Get(ctx context.Context, vessel *pb.Request, resp *pb.Reponse) error {
	return nil
}

// FindAvaliable 超照可用的货轮
func (v *Vessel) FindAvaliable(ctx context.Context, vessel *pb.Request, resp *pb.Reponse) error {
	return nil
}
