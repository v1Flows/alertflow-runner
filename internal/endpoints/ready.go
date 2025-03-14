package endpoints

import (
	"github.com/AlertFlow/runner/config"
	"github.com/gin-gonic/gin"
)

func ReadyEndpoint(cfg config.Config, router *gin.Engine) {
	router.GET("/ready", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}
