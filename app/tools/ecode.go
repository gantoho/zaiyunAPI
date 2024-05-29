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
	NotLogin = ECode{Code: 701, Message: "Login Error", Data: nil}
	ParamErr = ECode{Code: 702, Message: "Param Error", Data: nil}
)
