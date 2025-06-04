package password

import (
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/forgot/password/verification/email"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func ForgotPassword(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client) {

    forgotPassRoutes := router.Group("/password")
    email.RequestOtp(forgotPassRoutes,DB,redisClient)
	email.ValidateOtp(forgotPassRoutes,DB,redisClient)
}