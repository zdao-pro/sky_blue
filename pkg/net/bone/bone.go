package bone

import (
	"github.com/zdao-pro/sky_blue/pkg/net/bone/middleware"

	"github.com/gin-gonic/gin"
)

//Server is a gin Engine
var Server *gin.Engine

//Init http server
func Init() *gin.Engine {
	Server = gin.New()
	Server.Use(middleware.GetTracer(), middleware.GetLogger(middleware.LogConfig{}))
	return Server
}
