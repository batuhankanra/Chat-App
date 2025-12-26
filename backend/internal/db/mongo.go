package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/batuhankanra/Chat-App/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func MongoConnect(cfg config.Config) {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDb connection error:", err)
		return
	}
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Mongo ping failed:", err)

	}
	DB = client.Database(cfg.DBName)
	fmt.Println("mongo connected")
}
