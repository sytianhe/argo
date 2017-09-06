package main

import (
	"applatix.io/axerror"
	"github.com/gin-gonic/gin"
)

func PingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(axerror.REST_STATUS_OK, "pong")
		return
	}
}
