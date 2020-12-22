package gin

import (
	"net/http"
	"testing"

	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/zdao-pro/sky_blue/pkg/net/http/gin"
)

//Server is a gin Engine
var Server *gin.Engine

//Init http server
func TestGin(t *testing.T) {
	log.Init(nil)
	Server = gin.Default()
	Server.GET("/ping", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})
	Server.Run()
}
