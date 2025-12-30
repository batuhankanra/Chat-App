package repository

import (
	"context"
	"time"

	"github.com/batuhankanra/Chat-App/internal/db"
	"github.com/batuhankanra/Chat-App/internal/models"
	"github.com/batuhankanra/Chat-App/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func IsteamMember(teamId, userId string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	objTeamId, err := utils.ToObjectId(teamId)
	if err != nil {
		return false, err
	}
	objUserId, err := utils.ToObjectId(userId)
	if err != nil {
		return false, err
	}

	count, err := db.DB.Collection(models.TeamCollection).CountDocuments(ctx, bson.M{"_id": objTeamId, "members": objUserId})
	return count > 0, err
}
func CanAccessChannel(channelId, userId string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objChannelId, err := utils.ToObjectId(channelId)
	if err != nil {
		return false, err
	}
	objUserId, err := utils.ToObjectId(userId)
	if err != nil {
		return false, err
	}
	count, err := db.DB.Collection(models.ChannelCollection).CountDocuments(ctx, bson.M{"_id": objChannelId, "$or": []bson.M{{"isPrivate": false}, {"members": objUserId}}})
	return count > 0, err
}
