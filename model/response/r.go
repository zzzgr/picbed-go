package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// R 响应对象
type R struct {
	Code      ResCode     `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func Ok(c *gin.Context, msg string, data interface{}) {
	result(c, http.StatusOK, CodeSuccess, msg, data)
}

func Fail(c *gin.Context, httpStatusCode int, code ResCode) {
	result(c, httpStatusCode, code, code.Msg(), nil)
}

func FailMsg(c *gin.Context, httpStatusCode int, msg string) {
	result(c, httpStatusCode, CodeFail, msg, nil)
}

func result(c *gin.Context, httpStatusCode int, code ResCode, msg string, data interface{}) {
	c.JSON(httpStatusCode,
		&R{
			Code:      code,
			Msg:       msg,
			Data:      data,
			Timestamp: time.Now().Unix(),
		})
}
