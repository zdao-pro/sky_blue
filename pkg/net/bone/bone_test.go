package bone

import (
	"testing"

	"github.com/zdao-pro/sky_blue/pkg/log"

	"github.com/gin-gonic/gin"
)

func cli(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func TestApollo(t *testing.T) {
	log.Init(nil)
	e := Init()
	e.GET("/cli", cli)
	e.Run()
}
