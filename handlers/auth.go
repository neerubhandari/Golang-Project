package handlers

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-API-KEY") !=
			("eUbP9shywUygMx7u") {
			c.AbortWithStatus(401)
		}
		c.Next()
	}
}
