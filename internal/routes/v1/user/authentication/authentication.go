package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/authentication/userPass"
)

func Auth(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client){
	loginRoutes := router.Group("/auth")
	userpass.UserPassValidation(loginRoutes,DB,redisClient)
}
