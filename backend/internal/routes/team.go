package routes

import (
	"github.com/batuhankanra/Chat-App/internal/handlers"
	"github.com/batuhankanra/Chat-App/internal/middleware"
	"github.com/gin-gonic/gin"
)

func TeamPublicRoute(r *gin.RouterGroup) {
	team := r.Group("/teams")
	{
		team.GET("/", middleware.JWTAuth(), handlers.GetMyTeams)
		team.POST("/", middleware.JWTAuth(), handlers.CreateTeam)
		team.PUT("/:teamId", middleware.JWTAuth(), handlers.UpdateTeam)
		team.DELETE("/:teamId", middleware.JWTAuth(), handlers.DeleteTeam)
		team.POST("/:teamId/members", middleware.JWTAuth(), handlers.AddTeamMember)
		team.DELETE("/:teamId/members/:removeId", middleware.JWTAuth(), handlers.RemoveMember)

	}
}
