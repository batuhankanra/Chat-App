package repository

import (
	"context"
	"errors"
	"time"

	"github.com/batuhankanra/Chat-App/internal/db"
	"github.com/batuhankanra/Chat-App/internal/models"
	"github.com/batuhankanra/Chat-App/internal/utils"
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

func AddMemberToTeam(teamId, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objTeamId, err := utils.ToObjectId(teamId)
	if err != nil {
		return err
	}
	objUserId, err := utils.ToObjectId(userId)
	if err != nil {
		return err
	}

	resp, err := db.DB.Collection(models.TeamCollection).UpdateOne(ctx, bson.M{"_id": objTeamId, "members": bson.M{"$ne": objUserId}}, bson.M{"$addToSet": bson.M{"members": objUserId}})
	if err != nil {
		return err
	}
	if resp.MatchedCount == 0 {
		return errors.New("team not found or user already member")
	}
	return nil
}
func IsteamOwner(teamId, userId string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objTeamId, err := utils.ToObjectId(teamId)
	if err != nil {
		return false, err
	}
	objUserId, err := utils.ToObjectId(userId)
	if err != nil {
		return false, err
	}
	count, err := db.DB.Collection(models.TeamCollection).CountDocuments(
		ctx,
		bson.M{
			"_id":     objTeamId,
			"ownerId": objUserId,
		},
	)
	return count > 0, err
}
func UpdateTeamName(teamId, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objId, err := utils.ToObjectId(teamId)
	if err != nil {
		return err
	}
	res, err := db.DB.Collection(models.TeamCollection).UpdateOne(
		ctx,
		bson.M{"_id": objId},
		bson.M{
			"$set": bson.M{
				"name":      name,
				"updatedAt": time.Now(),
			},
		},
	)
	if res.MatchedCount == 0 {
		return errors.New("team not found")

	}
	return err
}
func DeleteTeam(teamId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objId, err := utils.ToObjectId(teamId)
	if err != nil {
		return err
	}
	_, err = db.DB.Collection(models.TeamCollection).DeleteOne(
		ctx,
		bson.M{"_id": objId},
	)
	return err
}
func RemoveMember(teamId, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objId, err := utils.ToObjectId(teamId)
	if err != nil {
		return errors.New("sda")
	}
	objUserId, err := utils.ToObjectId(userId)
	if err != nil {
		return errors.New("sa")
	}
	res, err := db.DB.Collection(models.TeamCollection).UpdateOne(ctx, bson.M{"_id": objId, "members": objUserId}, bson.M{"$pull": bson.M{"members": objUserId}})
	if res.MatchedCount == 0 {
		return errors.New("team not found")
	}
	if res.ModifiedCount == 0 {
		return errors.New("member not removed")
	}
	return err
}
