package api

import (
	"github.com/gin-gonic/gin"
	"ginTest/api/v1"
)

func AddRouters(r *gin.Engine)  {
	v1.AddRouter(r)
}
