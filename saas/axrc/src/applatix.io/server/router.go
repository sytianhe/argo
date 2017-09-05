package main

import (
	"applatix.io/common"
	"github.com/gin-gonic/gin"
)

func GetRouterAxrc() *gin.Engine {
	router := gin.Default()
	router.Use(common.ValidateCache)
	router.Use(common.GetGZipHandler())

	v1 := router.Group("v1")
	{
		v1.GET("ping", PingHandler())

	}

	return router
}
