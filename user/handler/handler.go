package handler

import (
	"context"
	"fmt"

	"github.com/daymenu/shipping/user/model"
	pb "github.com/daymenu/shipping/user/proto/user"
	"github.com/daymenu/shipping/user/token"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//User 用户
type User struct {
	DB           *gorm.DB
	TokenService token.Authable
}

// Get 获取用户
func (u *User) Get(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	um := model.UserModel{DB: u.DB}
	uu, err := um.Get(req)
	if err != nil {
		return err
	}
	resp.Code = 200
	resp.User = uu
	return nil
}

// Create  创建
func (u *User) Create(ctx context.Context, user *pb.User, resp *pb.Response) error {
	um := model.UserModel{DB: u.DB}
	hashbytes, err := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashbytes)
	uu, err := um.Create(user)
	if err != nil {
		return err
	}
	resp.Code = 200
	resp.User = uu
	return nil
}

// Update  创建
func (u *User) Update(ctx context.Context, user *pb.User, resp *pb.Response) error {
	um := model.UserModel{DB: u.DB}
	if user.GetPassword() != "" {
		hashbytes, err := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashbytes)
	}
	uu, err := um.Update(user)
	if err != nil {
		return err
	}
	resp.Code = 200
	resp.User = uu
	return nil
}

// Page 　列表
func (u *User) Page(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	um := model.UserModel{DB: u.DB}
	users, count, err := um.Page(req)
	if err != nil {
		return err
	}
	resp.Code = 200
	resp.Users = users
	resp.Count = count
	return nil
}

// Login 登录
func (u *User) Login(ctx context.Context, user *pb.User, token *pb.Token) error {
	um := model.UserModel{DB: u.DB}
	uu, err := um.GetByNameOrEmail(user.GetName())

	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(uu.Password), []byte(user.Password)); err != nil {
		return err
	}
	tokenStr, err := u.TokenService.Encode(uu)
	if err != nil {
		return err
	}
	token.Token = tokenStr
	return nil
}

//UserInfo 用户信息
func (u *User) UserInfo(ctx context.Context, token *pb.Token, resp *pb.Response) error {
	custom, err := u.TokenService.Decode(token.Token)
	if err != nil {
		return err
	}
	resp.User = custom.User
	return nil
}

// ValidateToken 验证token
func (u *User) ValidateToken(ctx context.Context, token *pb.Token, resp *pb.Token) error {
	claims, err := u.TokenService.Decode(token.Token)
	if err != nil {
		return err
	}
	if claims.User.Id == 0 {
		return fmt.Errorf("invalid user")
	}
	resp.Valid = true
	return nil
}
