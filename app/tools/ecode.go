package tools

import "fmt"

type ECode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (e *ECode) String() string {
	return fmt.Sprintf("code:%d, message:%s, data:%v", e.Code, e.Message, e.Data)
}

var (
	OK       = ECode{Code: 200, Message: "OK", Data: nil}
	NotLogin = ECode{Code: 10001, Message: "未登录", Data: nil}
	ParamErr = ECode{Code: 10002, Message: "参数错误", Data: nil}
)
