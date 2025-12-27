package middleware

import (
	"net/http"
	"strings"

	"github.com/batuhankanra/Chat-App/internal/config"
	"github.com/batuhankanra/Chat-App/internal/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "you need to login!"})
			return
		}
		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "You need to login!!"})
			return
		}
		jwtCfg := config.LoadConfig().JWTSecret
		claims, err := utils.ParseToken(parts[1], jwtCfg)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "You need to login!!!"})
			return
		}
		ctx.Set("userId", claims.UserID)
		ctx.Next()
	}
}
