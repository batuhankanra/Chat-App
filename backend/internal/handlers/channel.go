package handlers

import (
	"net/http"
	"time"

	"github.com/batuhankanra/Chat-App/internal/models"
	"github.com/batuhankanra/Chat-App/internal/repository"
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
	channel := models.Channel{
		TeamID: req.TeamID,
		BaseModel: models.BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Members:   req.Members,
		IsPrivate: req.IsPrivate,
		Name:      req.Name,
	}
	if err := repository.CreateChannel(&channel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "channel create failed"})
		return
	}
	c.JSON(http.StatusCreated, channel)
}
