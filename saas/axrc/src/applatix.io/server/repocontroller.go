package main

import (
	"github.com/gin-gonic/gin"
	"applatix.io/axerror"
)

func PingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(axerror.REST_STATUS_OK, "pong")
		return
	}
}
