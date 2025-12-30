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

func CreateChannel(channel *models.Channel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.DB.Collection(models.MessageCollection).InsertOne(ctx, channel)
	return err

}

func GetChannelsByTeam(TeamId, userId string) ([]models.Channel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objTeamId, err := utils.ToObjectId(TeamId)
	if err != nil {
		return nil, err
	}
	objUserId, err := utils.ToObjectId(userId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"teamId": objTeamId,
		"$or": []bson.M{
			{"isPrivate": false},
			{"members": objUserId},
		},
	}
	cursor, err := db.DB.Collection(models.ChannelCollection).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var channel []models.Channel
	if err := cursor.All(ctx, &channel); err != nil {
		return nil, err
	}
	return channel, nil
}
func UpdateChannelName(channelId, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objCID, err := utils.ToObjectId(channelId)
	if err != nil {
		return err
	}
	res, err := db.DB.Collection(models.ChannelCollection).UpdateOne(
		ctx, bson.M{"_id": objCID}, bson.M{"$set": bson.M{"name": name, "updatedAt": time.Now()}},
	)
	if res.MatchedCount == 0 {
		return errors.New("channel not found")
	}
	return err

}

func DeleteChannel(channelId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objID, err := utils.ToObjectId(channelId)
	if err != nil {
		return err
	}
	_, err = db.DB.Collection(models.ChannelCollection).DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
