package app

import (
	"context"
	"log"

	pb "github.com/daymenu/shipping/container/proto/container"
	microapi "github.com/micro/go-micro/api/proto"
)

// Container 结构体
type Container struct {
	Service pb.ContainerServiceClient
}

// Get 实现方法
func (container *Container) Get(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	//初始成功
	resp.StatusCode = 200
	apiReq := APIRequest{request: req}
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
	resp.StatusCode = 200
	if req.Method != "POST" {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "请以POST方式提交数据")
		return nil
	}

	apiReq := APIRequest{request: req}
	log.Println(req.Post)
	width, err := apiReq.PostInt64("width")
	if err != nil && width <= 0 {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请传入正确的集装箱宽度")
		return nil
	}
	height, err := apiReq.PostInt64("height")
	if err != nil && height <= 0 {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请传入正确的集装箱高度度")
		return nil
	}
	customerID, err := apiReq.PostString("customer_id")
	if err != nil && customerID == "" {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请填写正确的客户id")
		return nil
	}
	origin, err := apiReq.PostString("origin")
	if err != nil && origin == "" {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请填写来源")
		return nil
	}
	userID, err := apiReq.PostString("user_id")
	if err != nil && userID == "" {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请填写正确的用户id")
		return nil
	}
	long, err := apiReq.PostInt64("long")
	if err != nil && long <= 0 {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请传入正确的集装箱长度")
		return nil
	}
	status, err := apiReq.PostInt32("status")
	if err != nil && status <= 0 {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请选择集装箱状态")
		return nil
	}
	response, err := container.Service.Create(ctx, &pb.Container{
		Width:      width,
		Height:     height,
		CustomerId: customerID,
		Origin:     origin,
		UserId:     userID,
		Long:       long,
		Status:     status,
	})
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
	resp.StatusCode = 200
	if req.Method != "POST" {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "请以POST方式提交数据")
		return nil
	}
	apiReq := APIRequest{request: req}

	id, err := apiReq.PostInt64("id")
	if err != nil && id < 0 {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请填写要修改的集装箱ID")
		return nil
	}
	long, err := apiReq.PostInt64("long")
	if err != nil && long < 0 {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请传入正确的集装箱长度")
		return nil
	}
	width, err := apiReq.PostInt64("width")
	if err != nil && width < 0 {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请传入正确的集装箱宽度")
		return nil
	}
	height, err := apiReq.PostInt64("height")
	if err != nil && height < 0 {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请传入正确的集装箱高度度")
		return nil
	}
	customerID, err := apiReq.PostString("customer_id")
	if err != nil && customerID == "" {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请填写正确的客户id")
		return nil
	}
	origin, err := apiReq.PostString("origin")
	if err != nil && origin == "" {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请填写来源")
		return nil
	}
	userID, err := apiReq.PostString("user_id")
	if err != nil && userID == "" {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请填写正确的用户id")
		return nil
	}
	status, err := apiReq.PostInt32("status")
	if err != nil && status < 0 {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请选择集装箱状态")
		return nil
	}
	response, err := container.Service.Update(ctx, &pb.Container{
		Id:         id,
		Width:      width,
		Height:     height,
		CustomerId: customerID,
		Origin:     origin,
		UserId:     userID,
		Long:       long,
		Status:     status,
	})
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

	long, err := apiReq.PostInt64("long")
	if err != nil && long < 0 {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请传入正确的集装箱长度")
		return nil
	}
	width, err := apiReq.PostInt64("width")
	if err != nil && width < 0 {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请传入正确的集装箱宽度")
		return nil
	}
	height, err := apiReq.PostInt64("height")
	if err != nil && height < 0 {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请传入正确的集装箱高度度")
		return nil
	}
	response, err := container.Service.Use(ctx, &pb.Request{
		Width:  width,
		Height: height,
		Long:   long,
	})

	resp.Body = APISuccess(struct {
		Containers []*pb.Container `json:"containers"`
	}{
		Containers: response.GetContainers(),
	})
	return nil
}

// GiveBack  创建集装箱
func (container *Container) GiveBack(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	resp.StatusCode = 200
	if req.Method != "POST" {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "请以POST方式提交数据")
		return nil
	}
	apiReq := APIRequest{request: req}

	id, err := apiReq.PostInt64("id")
	if err != nil && id < 0 {
		resp.StatusCode = 500
		resp.Body = APIError(10000, "请填写要归还的集装箱ID")
		return nil
	}

	response, err := container.Service.GiveBack(ctx, &pb.Request{
		Id: id,
	})

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
