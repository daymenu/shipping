package model

import (
	"fmt"

	pb "github.com/daymenu/shipping/vessel/proto/vessel"
	"github.com/jinzhu/gorm"
)

// VesselModel 结构体
type VesselModel struct {
	DB *gorm.DB
}

// Page 分页
func (vm *VesselModel) Page(req *pb.Request) ([]*pb.Vessel, int64, error) {
	var vessels []*pb.Vessel
	var count int64
	if req.GetSearch() != "" {
		vm.DB = vm.DB.Where("name like ?", req.GetSearch())
	}
	vm.DB.Model(&pb.Vessel{}).Count(&count)
	vm.DB.Limit(req.GetPageSize()).Offset((req.GetPage() - 1) * req.GetPageSize()).Find(&vessels)
	return vessels, count, nil
}

//Create 创建货轮
func (vm *VesselModel) Create(vessel *pb.Vessel) (*pb.Vessel, error) {
	if err := vm.DB.Create(vessel).Error; err != nil {
		return vessel, err
	}
	return nil, nil
}

//Update 修改货轮
func (vm *VesselModel) Update(vessel *pb.Vessel) (*pb.Vessel, error) {
	if err := vm.DB.Save(vessel).Error; err != nil {
		return vessel, err
	}
	return nil, nil
}

//Get 获取货轮
func (vm *VesselModel) Get(req *pb.Request) (*pb.Vessel, error) {
	var vessel pb.Vessel
	if err := vm.DB.Where("id=?", req.GetId()).First(&vessel).Error; err != nil {
		return nil, fmt.Errorf("%s%s%d", err, ":", req.GetId())
	}
	return &vessel, nil
}

// FindAvaiable 获取可用货轮
func (vm *VesselModel) FindAvaiable(req *pb.Request) (*pb.Vessel, error) {

	return nil, nil
}
