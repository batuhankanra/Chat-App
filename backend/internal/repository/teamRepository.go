package repository

import (
	"context"
	"time"

	"github.com/batuhankanra/Chat-App/internal/db"
	"github.com/batuhankanra/Chat-App/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateTeam(team *models.Team) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.DB.Collection(models.TeamCollection).InsertOne(ctx, team)
	return err
}

func GetTeamsByUser(userID string) ([]models.Team, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := db.DB.Collection(models.TeamCollection).Find(ctx, bson.M{"members": userID})
	if err != nil {
		return nil, err
	}
	var teams []models.Team
	if err = cursor.All(ctx, &teams); err != nil {
		return nil, err
	}
	return teams, err
}
