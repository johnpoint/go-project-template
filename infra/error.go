package infra

import "PROJECT_NAME/pkg/errorhelper"

var (
	// 请求异常 401xx
	ReqParseError = &errorhelper.Err{Code: 40100, Message: "请求参数异常"}
)
