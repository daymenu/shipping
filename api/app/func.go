package app

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// APIError  api错误封装
func APIError(code int, msg string) string {
	result, err := json.Marshal(Error{
		Code: code,
		Msg:  msg,
	})
	if err != nil {
		return `{"json":"parse error"}`
	}
	return string(result)
}

//GetInt 获取int参数
func (api *APIRequest) GetInt(key string) (int, error) {
	name, ok := api.request.Get[key]
	if !ok || len(name.Values) == 0 {
		return 0, fmt.Errorf("api get [%s] is none", key)
	}
	id, err := strconv.ParseInt(name.Values[0], 10, 0)
	return int(id), err
}

//GetInt32 获取int32参数
func (api *APIRequest) GetInt32(key string) (int32, error) {
	name, ok := api.request.Get[key]
	if !ok || len(name.Values) == 0 {
		return 0, fmt.Errorf("api get [%s] is none", key)
	}
	id, err := strconv.ParseInt(name.Values[0], 10, 0)
	return int32(id), err
}

//GetInt64 获取int64参数
func (api *APIRequest) GetInt64(key string) (int64, error) {
	name, ok := api.request.Get[key]
	if !ok || len(name.Values) == 0 {
		return 0, fmt.Errorf("api get [%s] is none", key)
	}
	return strconv.ParseInt(name.Values[0], 10, 0)
}

//GetString 获取String参数
func (api *APIRequest) GetString(key string) (string, error) {
	name, ok := api.request.Get[key]
	if !ok || len(name.Values) == 0 {
		return "", fmt.Errorf("api get [%s] is none", key)
	}
	return name.Values[0], nil
}

//GetStrings 获取String参数
func (api *APIRequest) GetStrings(key string) ([]string, error) {
	name, ok := api.request.Get[key]
	if !ok || len(name.Values) == 0 {
		return nil, fmt.Errorf("api get [%s] is none", key)
	}
	return name.Values, nil
}
