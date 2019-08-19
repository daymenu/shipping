package model

import (
	"fmt"

	pb "github.com/daymenu/shipping/user/proto/user"
	"github.com/jinzhu/gorm"
)

// UserModel 结构体
type UserModel struct {
	DB *gorm.DB
}

// Page 分页获取数据
func (u *UserModel) Page(r *pb.Request) ([]*pb.User, int64, error) {
	var users []*pb.User
	var count int64

	if r.GetName() != "" {
		u.DB = u.DB.Where("name like ?", r.GetName())
	}

	if r.GetEmail() != "" {
		u.DB = u.DB.Where("email like ?", r.GetEmail())
	}
	u.DB.Model(&pb.User{}).Count(&count)
	u.DB.Limit(r.GetPageSize()).Offset(((r.GetPage() - 1) * r.GetPageSize())).Find(&users)
	return users, count, nil
}

// Create 创建用户
func (u *UserModel) Create(user *pb.User) (*pb.User, error) {
	if err := u.DB.Create(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Update 修改用户
func (u *UserModel) Update(user *pb.User) (*pb.User, error) {
	if err := u.DB.Save(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Get 获取用户
func (u *UserModel) Get(req *pb.Request) (*pb.User, error) {
	var user pb.User
	if err := u.DB.Where("id=?", req.GetId()).First(&user).Error; err != nil {
		return nil, fmt.Errorf("%s%s%s%s", err, ":", req.GetId(), req.GetName())
	}
	return &user, nil
}

// GetByEmail 获取用户
func (u *UserModel) GetByEmail(email string) (*pb.User, error) {
	var user pb.User
	if err := u.DB.Where("email=?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("%s%s%s", err, ":", email)
	}
	return &user, nil
}

// GetByNameOrEmail 获取用户
func (u *UserModel) GetByNameOrEmail(name string) (*pb.User, error) {
	var user pb.User
	if err := u.DB.Where("name=? or email=?", name, name).First(&user).Error; err != nil {
		return nil, fmt.Errorf("%s%s%s", err, ":", name)
	}
	return &user, nil
}
