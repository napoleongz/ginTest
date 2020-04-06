package cmd

import (
	"github.com/gin-gonic/gin"
	"ginTest/api"
	"github.com/DeanThompson/ginpprof"
)

func API()  {

	r := gin.Default()

	api.AddRouters(r)

	ginpprof.Wrapper(r)


	r.Run(":8888")


}
