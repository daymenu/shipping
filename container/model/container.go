package model

import (
	"fmt"

	pb "github.com/daymenu/shipping/container/proto/container"
	"github.com/jinzhu/gorm"
)

// ContainerModel 结构体
type ContainerModel struct {
	DB *gorm.DB
}

// Use 根据请求获取符合要求的集装箱，然后锁定
// ps: 查找体积大于请求的集装箱
func (cm *ContainerModel) Use(r *pb.Request) ([]*pb.Container, error) {
	var cs []*pb.Container
	var canBluck int64
	cIds := make([]int64, 10)
	bluk := r.Height * r.Width * r.Long
	// 下面代码可以好好优化
	for canBluck > bluk {
		var tcs []*pb.Container // 临时
		sql := "status=1"
		cm.DB.Where(sql).Limit(10).Find(&tcs)
		for _, item := range tcs {
			if item.Id == 0 {
				return nil, fmt.Errorf("not enough container")
			}
			canBluck += item.Height * item.Width * item.Long
			if canBluck > bluk {
				cs = append(cs, item)
				cIds = append(cIds, item.Id)
				break
			}
		}
	}

	if num := cm.DB.Where("id in (?)", cIds).Updates(map[string]int{"status": 2}).RowsAffected; num < 1 {
		return nil, fmt.Errorf("container: giveback faild")
	}
	return cs, nil
}

// Page 分页获取数据
func (cm *ContainerModel) Page(r *pb.Request) ([]*pb.Container, error) {
	var cs []*pb.Container

	if r.GetName() != "" {
		cm.DB = cm.DB.Where("name like ?", r.GetName())
	}
	cm.DB.Limit(r.GetPageSize()).Offset(((r.GetPage() - 1) * r.GetPageSize())).Find(&cs)
	return cs, nil
}

// Create 创建一个集装箱
func (cm *ContainerModel) Create(c *pb.Container) (*pb.Container, error) {
	if err := cm.DB.Create(c).Error; err != nil {
		return c, err
	}
	return c, nil
}

// Update 创建一个集装箱
func (cm *ContainerModel) Update(c *pb.Container) (*pb.Container, error) {
	if err := cm.DB.Update(c).Error; err != nil {
		return c, err
	}
	return c, nil
}

// Get 获取集装箱
func (cm *ContainerModel) Get(c *pb.Request) (*pb.Container, error) {
	var container pb.Container
	if err := cm.DB.Where("id=?", c.GetId()).First(&container).Error; err != nil {
		return nil, fmt.Errorf("%s%s%d%s%d", err, ":", c.GetId(), c.GetName(), c.GetWidth())
	}
	return &container, nil
}

// GiveBack 归还集装箱
func (cm *ContainerModel) GiveBack(cs []*pb.Container) error {
	cIds := make([]int64, len(cs))
	for _, item := range cs {
		cIds = append(cIds, item.Id)
	}
	if num := cm.DB.Where("id in (?)", cIds).Updates(map[string]int{"status": 1}).RowsAffected; num < 1 {
		return fmt.Errorf("container: giveback faild")
	}
	return nil
}
