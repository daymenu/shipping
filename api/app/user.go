package app

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/daymenu/shipping/user/proto/user"
	microapi "github.com/micro/go-micro/api/proto"
)

// User user 结构体
type User struct {
	Service pb.UserServiceClient
}

// Page 列表
func (user *User) Page(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	apiReq := APIRequest{request: req}
	ctx, _ = apiReq.AddAuth(ctx)
	page, err := apiReq.GetInt64("page")
	if err != nil {
		page = 1
	}

	pageSize, err := apiReq.GetInt64("pageSize")
	if err != nil {
		pageSize = 1
	}
	response, err := user.Service.Page(ctx, &pb.Request{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(12000, "permisssion:"+err.Error())
		return nil
	}
	resp.Body = APISuccess(struct {
		Count int64      `json:"count"`
		Users []*pb.User `json:"users"`
	}{
		Count: response.GetCount(),
		Users: response.GetUsers(),
	})
	return nil
}

// Get 列表
func (user *User) Get(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	return nil
}

// Create 列表
func (user *User) Create(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	apiReq := APIRequest{request: req}
	ctx, _ = apiReq.AddAuth(ctx)
	resp.StatusCode = 200
	if req.Method != "POST" {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "请以POST方式提交数据")
		return nil
	}

	var u pb.User
	err := json.Unmarshal([]byte(req.GetBody()), &u)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "json parse:"+err.Error())
		return nil
	}

	vfs := []ValidateForm{
		{Key: "email", Field: u.GetEmail(), Msg: "请填写正确的邮箱"},
		{Key: "username", Field: u.GetName(), Msg: "请填写正确的用户名，以字母开头的数字结尾的组合"},
		{Key: "notempty", Field: u.GetCompany(), Msg: "请填写正确的公司名称"},
		{Key: "mobile", Field: u.GetMobile(), Msg: "请填写正确的手机号码"},
		{Key: "password", Field: u.GetPassword(), Msg: "请填写6-20位密码"},
		{Key: "num", Field: fmt.Sprintf("%d", u.GetStatus()), Msg: "请选择状态"},
	}

	if err := AutoCheck(vfs); err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10000, err.Error())
		return nil
	}

	response, err := user.Service.Create(ctx, &u)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10001, "api创建失败："+err.Error())
		return nil
	}

	resp.Body = APISuccess(struct {
		User *pb.User `json:"user"`
	}{
		User: response.GetUser(),
	})
	return nil
}

// Update 列表
func (user *User) Update(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	apiReq := APIRequest{request: req}
	ctx, _ = apiReq.AddAuth(ctx)
	resp.StatusCode = 200
	if req.Method != "POST" {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "请以POST方式提交数据")
		return nil
	}

	var u pb.User
	err := json.Unmarshal([]byte(req.GetBody()), &u)
	vfs := []ValidateForm{
		{Key: "email", Field: u.GetEmail(), Msg: "请填写正确的邮箱"},
		{Key: "username", Field: u.GetName(), Msg: "请填写正确的用户名，以字母开头的数字结尾的组合"},
		{Key: "notempty", Field: u.GetCompany(), Msg: "请填写正确的公司名称"},
		{Key: "mobile", Field: u.GetMobile(), Msg: "请填写正确的手机号码"},
		{Key: "num", Field: fmt.Sprintf("%d", u.GetStatus()), Msg: "请选择状态"},
	}

	if err := AutoCheck(vfs); err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10000, err.Error())
		return nil
	}
	if err := Check("password", u.GetPassword()); u.GetPassword() != "" && err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请输入密码")
		return nil
	}
	response, err := user.Service.Update(ctx, &u)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10001, "api修改失败："+err.Error())
		return nil
	}

	resp.Body = APISuccess(struct {
		User *pb.User `json:"user"`
	}{
		User: response.GetUser(),
	})
	return nil
}

// Login 列表
func (user *User) Login(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	apiReq := APIRequest{request: req}
	ctx, _ = apiReq.AddAuth(ctx)
	resp.StatusCode = 200
	if req.Method != "POST" {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "请以POST方式提交数据")
		return nil
	}

	var u pb.User
	err := json.Unmarshal([]byte(req.GetBody()), &u)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "json parse:"+err.Error())
		return nil
	}

	vfs := []ValidateForm{
		{Key: "username", Field: u.GetName(), Msg: "请填写正确的用户名，以字母开头的数字结尾的组合"},
		{Key: "password", Field: u.GetPassword(), Msg: "请填写6-20位密码"},
	}

	if err := AutoCheck(vfs); err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10000, err.Error())
		return nil
	}
	token, err := user.Service.Login(ctx, &u)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10001, err.Error())
		return nil
	}
	resp.StatusCode = 200
	resp.Body = APISuccess(struct {
		Token string `json:"token"`
	}{
		Token: token.GetToken(),
	})
	return nil
}

// UserInfo 列表
func (user *User) UserInfo(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	apiReq := APIRequest{request: req}
	ctx, _ = apiReq.AddAuth(ctx)
	tokenStr, err := apiReq.HeaderString("X-Token")
	if err != nil {
		return err
	}
	fmt.Println(tokenStr)
	response, err := user.Service.UserInfo(ctx, &pb.Token{Token: tokenStr})
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(12000, err.Error())
	}
	resp.StatusCode = 200
	fmt.Println("user:", response.GetUser())
	resp.Body = APISuccess(struct {
		User *pb.User `json:"user"`
	}{
		User: response.GetUser(),
	})
	return nil
}
