package handlers

import (
	"net/http"
	"time"

	"github.com/batuhankanra/Chat-App/internal/models"
	"github.com/batuhankanra/Chat-App/internal/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTeam(ctx *gin.Context) {
	userID := ctx.GetString("userId")
	var req models.CreateTeamRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "You need to choose a team"})
		return
	}
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "ss"})
		return
	}
	team := models.Team{

		Name:    req.Name,
		OwnerID: objID,
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

func AddTeamMember(c *gin.Context) {
	userId := c.GetString("userId")
	teamId := c.Param("teamId")

	var req models.AddTeamMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "you must select a user"})
		return
	}

	isOwner, err := repository.IsteamOwner(teamId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "authorization error"})
		return
	}
	if !isOwner {
		c.JSON(http.StatusForbidden, gin.H{"msg": "only owner can add members"})
		return
	}
	if err := repository.AddMemberToTeam(teamId, req.UserID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "there is no such user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "member added"})
}

func UpdateTeam(c *gin.Context) {
	userId := c.GetString("userId")
	teamId := c.Param("teamId")
	isOwner, _ := repository.IsteamOwner(teamId, userId)
	if !isOwner {
		c.JSON(http.StatusForbidden, gin.H{"msg": "only owner"})
		return
	}
	var req models.CreateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "cannot be empty"})
		return
	}
	if err := repository.UpdateTeamName(teamId, req.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Success"})
}
func DeleteTeam(c *gin.Context) {
	userId := c.GetString("userId")
	teamId := c.Param("teamId")
	isOwner, _ := repository.IsteamOwner(teamId, userId)
	if !isOwner {
		c.JSON(http.StatusForbidden, gin.H{"msg": "only is owner"})
		return
	}
	if err := repository.DeleteTeam(teamId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "team deleted"})
}

func RemoveMember(c *gin.Context) {
	userId := c.GetString("userId")
	teamId := c.Param("teamId")
	removeUserId := c.Param("removeId")
	isOwner, _ := repository.IsteamOwner(teamId, userId)
	if !isOwner || userId == removeUserId {
		c.JSON(http.StatusForbidden, gin.H{"msg": "only is owner"})
		return
	}
	if err := repository.RemoveMember(teamId, removeUserId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "member removed"})

}
