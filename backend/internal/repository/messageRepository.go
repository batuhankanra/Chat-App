package repository

import (
	"context"
	"time"

	"github.com/batuhankanra/Chat-App/internal/db"
	"github.com/batuhankanra/Chat-App/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMessages(channelId string, limit, offset int64) ([]models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Find().SetSort(bson.M{"createdAt": -1}).SetLimit(limit).SetSkip(offset)
	cursor, err := db.DB.Collection(models.MessageCollection).Find(ctx, bson.M{"channelId": channelId}, opts)
	if err != nil {
		return nil, err
	}
	var message []models.Message
	if err := cursor.All(ctx, &message); err != nil {
		return nil, err
	}
	return message, nil
}

func SaveMessage(msg *models.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.DB.Collection(models.MessageCollection).InsertOne(ctx, msg)
	return err
}
