package forgot

import (
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/forgot/password"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)
func Forgot(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client) {

    forgotRoutes := router.Group("/forgot")
	password.ForgotPassword(forgotRoutes,DB,redisClient)

}