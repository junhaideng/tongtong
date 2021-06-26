package main

import "errors"

// Healthy 用来检查后端接口
var Healthy = "backend is fine"

// DefaultQueryNum 默认请求获取到的通知的数量
var DefaultQueryNum = 10

// SuccessCode 处理正确的code
var SuccessCode = 200

// Success 处理正确的msg
var Success = "success"

// ErrCode 处理错误返回的code
var ErrCode = -1

// ErrFailedGet 爬取信息错误的时候返回的msg
var ErrFailedGet = errors.New("获取信息失败")

// ErrBadRequest 请求参数不正确返回的msg
var ErrBadRequest = errors.New("请求参数错误")
