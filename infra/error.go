package infra

import "github.com/johnpoint/go-bootstrap/berror"

var (
	// ReqParseError 请求异常 401xx
	ReqParseError = &berror.BErr{Code: 40100, Message: "请求参数异常"}
)
