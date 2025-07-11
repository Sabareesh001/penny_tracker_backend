package email

import (
	generateotp "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/forgot/password/verification/generateOtp"
	validateotp "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/forgot/password/verification/validateOtp"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)


func EmailRoutes(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client){
	emailRoutes := router.Group("/email");
	generateotp.RequestOtp(emailRoutes,DB,redisClient)
	validateotp.ValidateOtp(emailRoutes,DB,redisClient)
}