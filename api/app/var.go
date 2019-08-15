package app

import microapi "github.com/micro/go-micro/api/proto"

// Error 返回错误结构体
type Error struct {
	Code int
	Msg  string
}

// APIRequest  封装下获取参数的方法
type APIRequest struct {
	request *microapi.Request
}
