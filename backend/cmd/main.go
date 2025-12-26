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

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Run(":" + cfg.Port)

}
