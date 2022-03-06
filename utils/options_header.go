package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OptionsHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			//c.Header("Access-Control-Allow-Origin", "*")
			//c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			//c.Header("Access-Control-Allow-Headers", "authorization,origin,content-type,accept,x-auth-token,User-Agent")
			//c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.Header("Content-Type", "application/json")
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
