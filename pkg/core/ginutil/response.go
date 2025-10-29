package ginutil

import (
	"encoding/json"
	"gin-scaffold/pkg/core"
	"net/http"
	"reflect"

	"gin-scaffold/pkg/logging"

	"gin-scaffold/pkg/core/errorsx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResData(c *gin.Context, data interface{}) {
	ResJSON(c, http.StatusOK, Response{
		Success: true,
		Code:    http.StatusOK,
		Data:    data,
	})
}
func ResOKWithMessage(c *gin.Context, message string) {
	ResJSON(c, http.StatusOK, Response{
		Success: true,
		Code:    http.StatusOK,
		Message: message,
	})
}

func ResOK(c *gin.Context) {
	ResJSON(c, http.StatusOK, Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "success",
	})
}

func ResPage(c *gin.Context, data interface{}, pagination *core.Pagination) {

	reflectValue := reflect.Indirect(reflect.ValueOf(data))
	if reflectValue.IsNil() {
		data = make([]interface{}, 0)
	}

	ResJSON(c, http.StatusOK, Response{
		Success: true,
		Code:    http.StatusOK,
		Data:    data,
		Meta:    pagination,
	})
}

func ResList(c *gin.Context, list interface{}) {

	var total int64
	reflectValue := reflect.Indirect(reflect.ValueOf(list))
	if reflectValue.IsNil() {
		list = make([]interface{}, 0)
		total = 0
	} else {
		total = int64(reflectValue.Len())
	}

	ResJSON(c, http.StatusOK, Response{
		Success: true,
		Code:    http.StatusOK,
		Data:    list,
		Meta:    core.Pagination{Total: total},
	})
}

func ResJSON(c *gin.Context, status int, data interface{}) {
	buf, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()

}

func ResError(c *gin.Context, err error) {
	ctx := c.Request.Context()
	var res *errorsx.ResponseError

	if err != nil {
		res = errorsx.UnwrapResponseError(err)
	}

	if res == nil {
		res = errorsx.NewInternal("Internal server error").WithError(err)
	}

	if status := res.StatusCode; status >= 400 && status < 500 {
		logging.WithContext(ctx).WithOptions(zap.AddCallerSkip(1)).Warn(res.Message, zap.Error(err))
	} else if status >= 500 {
		ctx = logging.NewStackContext(ctx, err)
		logging.WithContext(ctx).WithOptions(zap.AddCallerSkip(1)).Error(res.Message, zap.Error(err))
	}

	ResJSON(c, res.StatusCode, Response{
		Code:    res.Code,
		Message: res.Message,
	})
}
