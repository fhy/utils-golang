package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

var jsonContentType = []string{"application/json; charset=utf-8"}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

func ResponseSuccessJson(c *gin.Context, data interface{}, statusCode ...int) {
	httpStatus := http.StatusOK
	if len(statusCode) > 0 {
		httpStatus = statusCode[0]
	}
	val, _ := json.Marshal(addResponseExtraInfo(c, gin.H{"code": SUCCESS_CODE_CUSTOM, "msg": SUCCESS_MSG_CUSTOM, "data": data}))
	c.Status(httpStatus)
	writeContentType(c.Writer, jsonContentType)
	c.Writer.Write(val)
	c.Writer.Flush()
}

func ResponseSuccessJsonWithPagination(c *gin.Context, data interface{}, p *Pagination, statusCode ...int) {
	httpStatus := http.StatusOK
	if len(statusCode) > 0 {
		httpStatus = statusCode[0]
	}
	val, _ := json.Marshal(addResponseExtraInfo(c, gin.H{"code": SUCCESS_CODE_CUSTOM, "msg": SUCCESS_MSG_CUSTOM, "data": data, "pagination": p}))
	c.Status(httpStatus)
	writeContentType(c.Writer, jsonContentType)
	c.Writer.Write(val)
	c.Writer.Flush()
}

func ResponseFailedJson(c *gin.Context, code int, message string, data interface{}, statusCode ...int) {
	httpStatus := http.StatusOK
	if len(statusCode) > 0 {
		httpStatus = statusCode[0]
	}
	val, _ := json.Marshal(addResponseExtraInfo(c, gin.H{"code": code, "msg": message, "data": data}))
	c.Status(httpStatus)
	writeContentType(c.Writer, jsonContentType)
	c.Writer.Write(val)
	c.Writer.Flush()
	c.Abort()
}

func OutHttpJson(c *gin.Context, code int, message string, data interface{}, statusCode ...int) {
	httpStatus := http.StatusOK
	if len(statusCode) > 0 {
		httpStatus = statusCode[0]
	}
	c.JSON(httpStatus, addResponseExtraInfo(c, gin.H{"code": code, "message": message, "data": data}))
}

func ResponseErrorJson(c *gin.Context, err error) {
	if err != nil && gin.Mode() == gin.DebugMode {
		ResponseFailedJson(c, ERRCODE_CUSTOM, err.Error(), nil, http.StatusInternalServerError)
	} else {
		ResponseFailedJson(c, ERRCODE_CUSTOM, ERRMSG_CUSTOM, nil)
	}
}

func addResponseExtraInfo(c *gin.Context, content gin.H) gin.H {
	if extraInfo, exists := c.Get(EXTRA_INFO); exists {
		content[EXTRA_INFO] = extraInfo
	}
	return content
}

func AddExtraInfoToContext(c *gin.Context, key string, value interface{}) {
	if extraInfo, exists := c.Get(EXTRA_INFO); exists {
		if m, ok := extraInfo.(map[string]interface{}); ok {
			m[key] = value
			c.Set(EXTRA_INFO, m)
			return
		}
	}
	c.Set(EXTRA_INFO, gin.H{key: value})
}
