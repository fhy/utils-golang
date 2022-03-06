package utils

import (
	"errors"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Status(http.StatusBadRequest)
				debug.PrintStack()
				logger.Error("ErrorHandler:", err)
				ResponseErrorJson(c, errors.New(err.(string)))
			}
		}()
		c.Next()
	}
}
