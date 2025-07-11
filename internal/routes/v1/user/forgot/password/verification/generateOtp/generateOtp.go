package generateotp

import (
	"time"

	generateotp "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/verification/email/generateOtp"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)


func RequestOtp(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client) {
	router.POST("/requestOtp", func(ctx *gin.Context) {
		generateotp.GenerateOtp(ctx, DB, redisClient,"","Forgot Password : Reset Request","forgotPass:",time.Second*50)
	})
}

