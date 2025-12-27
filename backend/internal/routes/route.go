package routes

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.SetTrustedProxies(nil)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("/api")

	AuthPublicRoute(api)
	TeamPublicRoute(api)

	return r
}
