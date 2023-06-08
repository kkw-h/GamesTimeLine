//Package response
/*
@Title: response.go
@Description
@Author: kkw 2023/1/3 16:11
*/
package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ShowTypeSilent       = 0 //无需提示
	ShowTypeWarn         = 1 //警告提示
	ShowTypeErr          = 2 //错误提示
	ShowTypeNotification = 4 //通知提示
)

type Response struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data,omitempty"`
	ErrorCode    int         `json:"error_code,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
	ShowType     int32       `json:"show_type,omitempty"` //error display type： 0 silent; 1 message.warn; 2 message.error; 4 notification; 9 page
}
type Error struct {
	Code     int    `json:"code"`
	ErrCode  int    `json:"err_code"`
	Message  string `json:"message"`
	ShowType int32  `json:"show_type"`
}

func ErrorResult(ctx *gin.Context, err *Error) {
	errMsg := &Response{
		Success:      false,
		ErrorCode:    err.ErrCode,
		ErrorMessage: err.Message,
		ShowType:     err.ShowType,
	}
	ctx.JSON(err.Code, errMsg)
}

func SuccessResult(ctx *gin.Context, response interface{}) {
	msg := &Response{
		Success: true,
		Data:    response,
	}
	ctx.JSON(http.StatusOK, msg)
}
func Success(ctx *gin.Context) {
	msg := &Response{
		Success: true,
	}
	ctx.JSON(http.StatusOK, msg)
}

/**
  1001-1100 为系统内部错误专用
*/

var DBFindError = &Error{
	Code:     http.StatusInternalServerError,
	ErrCode:  1001,
	Message:  "内部错误,请稍后再试",
	ShowType: ShowTypeErr,
}

var BindError = &Error{
	Code:     http.StatusBadRequest,
	ErrCode:  1002,
	Message:  "添加的数据错误",
	ShowType: ShowTypeErr,
}
