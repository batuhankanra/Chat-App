package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	MongoURI  string
	RedisURI  string
	DBName    string
	JWTSecret string
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Error:.env file not fount")
	}
	return Config{
		Port:      getEnv("PORT", "8080"),
		MongoURI:  getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		RedisURI:  getEnv("REDIS_URI", "localhost:6379"),
		DBName:    getEnv("DB_NAME", "DailyPlanner"),
		JWTSecret: getEnv("JWT_SECRET", "varsayılan-gizli-anahtar"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
