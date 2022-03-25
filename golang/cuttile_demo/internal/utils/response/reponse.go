package response

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code    int         `json:"code,omitempty"`  // 业务响应状态码
	Message string      `json:"msg"`  // 提示消息
	Data    interface{} `json:"data"`  // 数据
}

func (g *Gin) Response(code int, msg string, data interface{}) {
	g.Ctx.JSON(200, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
	return
}

type ResponseLis struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
	Page    int         `json:"page,omitempty"`
	Amounts int64       `json:"amounts,omitempty"`
}

func (g *Gin) ResponseLis(code int, msg string, data interface{}, page int, amounts int64) {
	g.Ctx.JSON(200, ResponseLis{
		Code:    code,
		Message: msg,
		Data:    data,
		Page:    page,
		Amounts: amounts,
	})
	return
}
