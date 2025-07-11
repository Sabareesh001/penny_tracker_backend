package verification

import (
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/verification/email"
	mobile "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/verification/mobile"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Verification(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){
	verificationRoutes := router.Group("/verify")
    mobile.MobileVerification(verificationRoutes,DB,redisClient)
	email.EmailVerification(verificationRoutes,DB,redisClient)
}
