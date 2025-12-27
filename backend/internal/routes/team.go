package routes

import (
	"github.com/batuhankanra/Chat-App/internal/handlers"
	"github.com/batuhankanra/Chat-App/internal/middleware"
	"github.com/gin-gonic/gin"
)

func TeamPublicRoute(r *gin.RouterGroup) {
	auth := r.Group("/teams")
	{
		auth.GET("/", middleware.JWTAuth(), handlers.GetMyTeams)
		auth.POST("/", middleware.JWTAuth(), handlers.CreateTeam)
	}
}
