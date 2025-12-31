package handlers

import (
	"net/http"
	"time"

	"github.com/batuhankanra/Chat-App/internal/models"
	"github.com/batuhankanra/Chat-App/internal/repository"
	"github.com/batuhankanra/Chat-App/internal/utils"
	"github.com/gin-gonic/gin"
)

func CreateChannel(c *gin.Context) {
	userId := c.GetString("userId")
	var req models.CreateChannelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "cannot be empty"})
		return
	}
	ismember, _ := repository.IsteamMember(req.TeamID, userId)
	if !ismember {
		c.JSON(http.StatusForbidden, gin.H{"msg": "not team member"})
		return
	}
	if req.IsPrivate && len(req.Members) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "private channel request"})
		return
	}
	objTeamId, err := utils.ToObjectId(req.TeamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error"})
		return
	}
	objMembers, err := utils.ToObjectIdSlice(req.Members)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid member id"})
		return
	}
	channel := models.Channel{
		TeamID: objTeamId,
		BaseModel: models.BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Members:   objMembers,
		IsPrivate: req.IsPrivate,
		Name:      req.Name,
	}
	if err := repository.CreateChannel(&channel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "channel create failed"})
		return
	}
	c.JSON(http.StatusCreated, channel)
}
