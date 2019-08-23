package app

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/daymenu/shipping/vessel/proto/vessel"
	microapi "github.com/micro/go-micro/api/proto"
)

// Vessel 结构体
type Vessel struct {
	Service pb.VesselServiceClient
}

// Page 列表
func (vessel *Vessel) Page(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
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
	response, err := vessel.Service.Page(ctx, &pb.Request{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(12000, "permisssion:"+err.Error())
		return nil
	}
	resp.Body = APISuccess(struct {
		Count   int64        `json:"count"`
		Vessels []*pb.Vessel `json:"vessels"`
	}{
		Count:   response.GetCount(),
		Vessels: response.GetVessels(),
	})
	return nil
}

// Get 列表
func (vessel *Vessel) Get(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	return nil
}

// Create 列表
func (vessel *Vessel) Create(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	apiReq := APIRequest{request: req}
	ctx, _ = apiReq.AddAuth(ctx)
	resp.StatusCode = 200
	if req.Method != "POST" {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "请以POST方式提交数据")
		return nil
	}

	var v pb.Vessel
	err := json.Unmarshal([]byte(req.GetBody()), &v)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "json parse:"+err.Error())
		return nil
	}

	vfs := []ValidateForm{
		{Key: "capacity", Field: fmt.Sprintf("%d", v.GetCapacity()), Msg: "请填写正确的容量"},
		{Key: "max_weight", Field: fmt.Sprintf("%d", v.GetMaxWeight()), Msg: "请填写正确的容量"},
		{Key: "name", Field: v.GetName(), Msg: "请填写货轮名称"},
		{Key: "long", Field: fmt.Sprintf("%d", v.GetLong()), Msg: "请填写正确的高度"},
		{Key: "width", Field: fmt.Sprintf("%d", v.GetWidth()), Msg: "请填写正确的宽度"},
		{Key: "height", Field: fmt.Sprintf("%d", v.GetHeight()), Msg: "请填写正确的高度"},
		{Key: "owner_id", Field: v.GetOwnerId(), Msg: "请填写货轮所属人"},
	}

	if err := AutoCheck(vfs); err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10000, err.Error())
		return nil
	}

	response, err := vessel.Service.Create(ctx, &v)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10001, "api创建失败："+err.Error())
		return nil
	}

	resp.Body = APISuccess(struct {
		Vessel *pb.Vessel `json:"vessel"`
	}{
		Vessel: response.GetVessel(),
	})
	return nil
}

// Update 列表
func (vessel *Vessel) Update(ctx context.Context, req *microapi.Request, resp *microapi.Response) error {
	apiReq := APIRequest{request: req}
	ctx, _ = apiReq.AddAuth(ctx)
	resp.StatusCode = 200
	if req.Method != "POST" {
		resp.StatusCode = 500
		resp.Body = APIError(11000, "请以POST方式提交数据")
		return nil
	}

	var v pb.Vessel
	err := json.Unmarshal([]byte(req.GetBody()), &v)
	vfs := []ValidateForm{
		{Key: "capacity", Field: fmt.Sprintf("%d", v.GetCapacity()), Msg: "请填写正确的容量"},
		{Key: "max_weight", Field: fmt.Sprintf("%d", v.GetMaxWeight()), Msg: "请填写正确的容量"},
		{Key: "name", Field: v.GetName(), Msg: "请填写货轮名称"},
		{Key: "long", Field: fmt.Sprintf("%d", v.GetLong()), Msg: "请填写正确的高度"},
		{Key: "width", Field: fmt.Sprintf("%d", v.GetWidth()), Msg: "请填写正确的宽度"},
		{Key: "height", Field: fmt.Sprintf("%d", v.GetHeight()), Msg: "请填写正确的高度"},
		{Key: "owner_id", Field: v.GetOwnerId(), Msg: "请填写货轮所属人"},
	}

	if err := AutoCheck(vfs); err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10000, err.Error())
		return nil
	}
	response, err := vessel.Service.Update(ctx, &v)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = APIError(10001, "api修改失败："+err.Error())
		return nil
	}

	resp.Body = APISuccess(struct {
		Vessel *pb.Vessel `json:"vessel"`
	}{
		Vessel: response.GetVessel(),
	})
	return nil
}

// FindAvaiable 查找可用货轮
func (vessel *Vessel) FindAvaiable(ctx context.Context, req microapi.Request, resp microapi.Response) error {

	return nil
}
