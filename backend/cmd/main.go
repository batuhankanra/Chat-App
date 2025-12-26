package main

import (
	"github.com/batuhankanra/Chat-App/internal/config"
	"github.com/batuhankanra/Chat-App/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	cfg := config.LoadConfig()
	db.MongoConnect(cfg)
	redisCLient := db.RedisConnect(cfg.RedisURI, "", 0)
	if err := redisCLient.Ping(); err != nil {
		panic(err)
	}
	defer redisCLient.Close()

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Run(":" + cfg.Port)

}
