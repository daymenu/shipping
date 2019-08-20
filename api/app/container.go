package app

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/daymenu/shipping/container/proto/container"
	microapi "github.com/micro/go-micro/api/proto"
)

// Container 结构体
type Container struct {
	Service pb.ContainerServiceClient
}

// Get 实现方法
func (container *Container) Get(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	apiReq := APIRequest{request: req}
	ctx, _ = apiReq.AddAuth(ctx)
	//初始成功
	resp.StatusCode = 200
	id, err := apiReq.GetInt64("id")
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "api:请传入正确的容器ID")
		return nil
	}

	//调用微服务
	response, err := container.Service.Get(ctx, &pb.Request{
		Id: id,
	})

	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10001, "api:"+err.Error())
		return nil
	}

	resp.Body = APISuccess(struct {
		Container *pb.Container `json:"container"`
	}{
		Container: response.GetContainer(),
	})

	return nil
}

// Page 容器分页
func (container *Container) Page(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	apiReq := APIRequest{request: req}
	ctx, _ = apiReq.AddAuth(ctx)
	page, err := apiReq.GetInt64("page")
	if err != nil {
		page = 1
	}
	pageSize, err := apiReq.GetInt64("pageSize")
	if err != nil {
		pageSize = 10
	}
	response, err := container.Service.Page(ctx, &pb.Request{
		Page:     page,
		PageSize: pageSize,
	})

	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10001, err.Error())
		return nil
	}
	resp.Body = APISuccess(struct {
		Count      int64           `json:"count"`
		Containers []*pb.Container `json:"containers"`
	}{
		Count:      response.GetCount(),
		Containers: response.GetContainers(),
	})
	return nil
}

// Create  创建集装箱
func (container *Container) Create(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	apiReq := APIRequest{request: req}
	ctx, _ = apiReq.AddAuth(ctx)
	resp.StatusCode = 200
	if req.Method != "POST" {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "请以POST方式提交数据")
		return nil
	}
	var c pb.Container
	err := json.Unmarshal([]byte(req.GetBody()), &c)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "json parse"+err.Error())
		return nil
	}
	vfs := []ValidateForm{
		{Key: "notempty", Field: c.GetCustomerId(), Msg: "请填写正确的客户id"},
		{Key: "notempty", Field: c.GetOrigin(), Msg: "请填写来源"},
		{Key: "notempty", Field: c.GetUserId(), Msg: "请填写用户Id"},
		{Key: "num", Field: fmt.Sprintf("%d", c.GetWidth()), Msg: "请传入正确的集装箱宽度"},
		{Key: "num", Field: fmt.Sprintf("%d", c.GetHeight()), Msg: "请传入正确的集装箱高度"},
		{Key: "num", Field: fmt.Sprintf("%d", c.GetLong()), Msg: "请传入正确的集装箱长度"},
		{Key: "num", Field: fmt.Sprintf("%d", c.GetStatus()), Msg: "请选择状态"},
	}

	if err := AutoCheck(vfs); err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10000, err.Error())
		return nil
	}
	response, err := container.Service.Create(ctx, &c)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10001, "api创建失败："+err.Error())
		return nil
	}

	resp.Body = APISuccess(struct {
		Container *pb.Container `json:"container"`
	}{
		Container: response.GetContainer(),
	})
	return nil
}

// Update  创建集装箱
func (container *Container) Update(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	apiReq := APIRequest{request: req}
	ctx, _ = apiReq.AddAuth(ctx)
	resp.StatusCode = 200
	if req.Method != "POST" {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "请以POST方式提交数据")
		return nil
	}

	var c pb.Container
	err := json.Unmarshal([]byte(req.GetBody()), &c)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "json parse"+err.Error())
		return nil
	}
	vfs := []ValidateForm{
		{Key: "notempty", Field: c.GetCustomerId(), Msg: "请填写正确的客户id"},
		{Key: "notempty", Field: c.GetOrigin(), Msg: "请填写来源"},
		{Key: "notempty", Field: c.GetUserId(), Msg: "请填写用户Id"},
		{Key: "num", Field: fmt.Sprintf("%d", c.GetWidth()), Msg: "请传入正确的集装箱宽度"},
		{Key: "num", Field: fmt.Sprintf("%d", c.GetHeight()), Msg: "请传入正确的集装箱高度"},
		{Key: "num", Field: fmt.Sprintf("%d", c.GetLong()), Msg: "请传入正确的集装箱长度"},
		{Key: "num", Field: fmt.Sprintf("%d", c.GetStatus()), Msg: "请选择状态"},
	}

	if err := AutoCheck(vfs); err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10000, err.Error())
		return nil
	}
	response, err := container.Service.Update(ctx, &c)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10001, "api修改失败："+err.Error())
		return nil
	}

	resp.Body = APISuccess(struct {
		Container *pb.Container `json:"container"`
	}{
		Container: response.GetContainer(),
	})
	return nil
}

// Use  创建集装箱
func (container *Container) Use(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	apiReq := APIRequest{request: req}
	ctx, _ = apiReq.AddAuth(ctx)
	var c pb.Request
	err := json.Unmarshal([]byte(req.GetBody()), &c)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "json parse"+err.Error())
		return nil
	}
	vfs := []ValidateForm{
		{Key: "num", Field: fmt.Sprintf("%d", c.GetWidth()), Msg: "请传入正确的集装箱宽度"},
		{Key: "num", Field: fmt.Sprintf("%d", c.GetHeight()), Msg: "请传入正确的集装箱高度"},
		{Key: "num", Field: fmt.Sprintf("%d", c.GetLong()), Msg: "请传入正确的集装箱长度"},
	}
	if err := AutoCheck(vfs); err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10000, err.Error())
		return nil
	}
	response, err := container.Service.Use(ctx, &c)

	resp.Body = APISuccess(struct {
		Containers []*pb.Container `json:"containers"`
	}{
		Containers: response.GetContainers(),
	})
	return nil
}

// GiveBack  创建集装箱
func (container *Container) GiveBack(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	apiReq := APIRequest{request: req}
	ctx, _ = apiReq.AddAuth(ctx)
	resp.StatusCode = 200
	if req.Method != "POST" {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "请以POST方式提交数据")
		return nil
	}
	var c pb.Request
	err := json.Unmarshal([]byte(req.GetBody()), &c)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "json parse"+err.Error())
		return nil
	}
	vfs := []ValidateForm{
		{Key: "num", Field: fmt.Sprintf("%d", c.GetId()), Msg: "请传入正确的集装箱ID"},
	}
	if err := AutoCheck(vfs); err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10000, err.Error())
		return nil
	}
	response, err := container.Service.GiveBack(ctx, &c)

	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10001, "api归还失败："+err.Error())
		return nil
	}

	resp.Body = APISuccess(struct {
		Container *pb.Container `json:"container"`
	}{
		Container: response.GetContainer(),
	})
	return nil
}
