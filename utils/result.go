package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpOkResult(c *gin.Context, message string, result interface{})  {
	c.JSON(http.StatusOK, HttpResult(0, message, result))
}

func HttpResult(status int, message string, result interface{})(map[string]interface{}) {
	return gin.H{
		"status" : status,
		"message" : message,
		"result" :result,
	}
}
