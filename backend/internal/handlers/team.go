package handlers

import (
	"net/http"
	"time"

	"github.com/batuhankanra/Chat-App/internal/models"
	"github.com/batuhankanra/Chat-App/internal/repository"
	"github.com/gin-gonic/gin"
)

func CreateTeam(ctx *gin.Context) {
	userID := ctx.GetString("userId")
	var req models.CreateTeamRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "You need to choose a team"})
		return
	}
	team := models.Team{

		Name:    req.Name,
		OwnerID: userID,
		Members: []string{userID},
		BaseModel: models.BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	if err := repository.CreateTeam(&team); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Team created failed"})
		return
	}
	ctx.JSON(http.StatusCreated, team)
}

func GetMyTeams(ctx *gin.Context) {
	userID := ctx.GetString("userId")
	teams, err := repository.GetTeamsByUser(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "could not fetch teams"})
		return
	}
	ctx.JSON(http.StatusOK, teams)
}
