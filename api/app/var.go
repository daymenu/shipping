package app

import microapi "github.com/micro/go-micro/api/proto"

// APIResponse 返回错误结构体
type APIResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// APIRequest  封装下获取参数的方法
type APIRequest struct {
	request *microapi.Request
}
