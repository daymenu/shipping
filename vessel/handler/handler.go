package handler

import (
	"context"

	"github.com/daymenu/shipping/vessel/model"
	pb "github.com/daymenu/shipping/vessel/proto/vessel"
	"github.com/jinzhu/gorm"
)

// Vessel 结构体
type Vessel struct {
	DB *gorm.DB
}

// Create 创建货轮
func (v *Vessel) Create(ctx context.Context, vessel *pb.Vessel, resp *pb.Reponse) error {
	vm := model.VesselModel{DB: v.DB}
	vv, err := vm.Create(vessel)
	if err != nil {
		return err
	}
	resp.Vessel = vv
	return nil
}

// Update 创建货轮
func (v *Vessel) Update(ctx context.Context, vessel *pb.Vessel, resp *pb.Reponse) error {
	vm := model.VesselModel{DB: v.DB}
	vv, err := vm.Update(vessel)
	if err != nil {
		return err
	}
	resp.Vessel = vv
	return nil
}

// Page 列表
func (v *Vessel) Page(ctx context.Context, req *pb.Request, resp *pb.Reponse) error {
	vm := model.VesselModel{DB: v.DB}
	vessels, count, err := vm.Page(req)
	if err != nil {
		return err
	}
	resp.Vessels = vessels
	resp.Count = count
	return nil
}

// Get 获取一个
func (v *Vessel) Get(ctx context.Context, req *pb.Request, resp *pb.Reponse) error {
	vm := model.VesselModel{DB: v.DB}
	vessel, err := vm.Get(req)
	if err != nil {
		return err
	}
	resp.Vessel = vessel
	return nil
}

// FindAvaliable 超照可用的货轮
func (v *Vessel) FindAvaliable(ctx context.Context, req *pb.Request, resp *pb.Reponse) error {
	vm := model.VesselModel{DB: v.DB}
	vessel, err := vm.FindAvaiable(req)
	if err != nil {
		return err
	}
	resp.Vessel = vessel
	return nil
}
