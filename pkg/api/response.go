package api

import (
	"encoding/json"
	"net/http"
	"reflect"

	"gin-scaffold/pkg/logging"

	"gin-scaffold/pkg/errorsx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// PaginationResponse 分页元数据
type PaginationResponse struct {
	Total    int64 `json:"total"`              //总数量
	PageSize int   `json:"pageSize,omitempty"` //每页数量
}

func ResData(c *gin.Context, data interface{}) {
	ResJSON(c, http.StatusOK, Response{
		Success: true,
		Code:    http.StatusOK,
		Data:    data,
	})
}
func ResWithMessage(c *gin.Context, message string) {
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

func ResPage(c *gin.Context, data interface{}, pr *PaginationResponse) {

	reflectValue := reflect.Indirect(reflect.ValueOf(data))
	if reflectValue.IsNil() {
		data = make([]interface{}, 0)
	}

	ResJSON(c, http.StatusOK, Response{
		Success: true,
		Code:    http.StatusOK,
		Data:    data,
		Meta:    pr,
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
		Meta:    &PaginationResponse{Total: total},
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
