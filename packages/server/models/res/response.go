package res

import (
	"net/http"
	"web/utils"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int 		`json:"code"`
	Data any 		`json:"data"`
	Msg  string `json:"msg"`
}

type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List	T			`json:"list"`
}


const(
	Success = 0
	Error = 10
)

func Result(code int, data any, msg string, c *gin.Context){
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(data any, msg string, c *gin.Context){
	Result(Success, data, msg, c)
}
func OkWith(c *gin.Context){
	Result(Success, map[string]string{}, "成功", c)
}
func OkWithData(data any, c *gin.Context){
	Result(Success, data, "成功", c)
}
func OkWithMessage(msg string, c *gin.Context){
	Result(Success, map[string]string{}, msg, c)
}
func OkWithList(list any, count int64, c *gin.Context){
	OkWithData(ListResponse[any]{
		List:  list,
		Count: count,
	}, c)
}


func Fail(data any, msg string, c *gin.Context){
	Result(Error, data, msg, c)
}
func FailWithMessage(msg string, c *gin.Context){
	Result(Error, map[string]any{}, msg, c)
}
func FailWithError(err error, obj any, c *gin.Context){
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, c)
}
func FailWithCode(code ErrorCode, c *gin.Context){
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(int(code), map[string]any{}, "未知错误", c)
}
