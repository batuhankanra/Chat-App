package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/batuhankanra/Chat-App/internal/config"
	"github.com/batuhankanra/Chat-App/internal/db"
	"github.com/batuhankanra/Chat-App/internal/models"
	"github.com/batuhankanra/Chat-App/internal/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "password error"})
		return
	}
	user := models.User{
		Username:     req.USername,
		Email:        req.Email,
		PasswordHash: hash,
		IsActive:     true,
		Teams:        []string{},
		BaseModel: models.BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.DB.Collection(models.UserCollection)
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "registered"})
}

func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "All areas must be filled!!"})
		return
	}
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.DB.Collection(models.UserCollection)
	err := collection.FindOne(ctx, bson.M{
		"email": req.Email,
	}).Decode(&user)
	if err != nil || !utils.CheckPassword(req.Password, user.PasswordHash) {

		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid credentials"})
		return
	}
	jwtCfg := config.LoadConfig().JWTSecret
	token, err := utils.GeneratedToken(user.ID.Hex(), jwtCfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Token error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"username": user.Username,
			"email":    user.Email,
		},
	})

}
