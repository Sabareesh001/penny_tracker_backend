package validate

import (
	"context"
	"errors"
	"strconv"

	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/body/email"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)



func ValidateOtp(ctx *gin.Context, DB *gorm.DB, redisClient *redis.Client,purpose string,key string,body *email.Body) error {

	parseError := ctx.ShouldBindJSON(body)

	if parseError != nil {
		response.DataInAdequate(ctx)
		return parseError
	}

	redisKey := body.Email + purpose + key
	emailCtx := context.Background()

	originalOtp := redisClient.Get(emailCtx, redisKey)

	if originalOtp.Err() == redis.Nil || originalOtp.Val() != strconv.Itoa(body.Otp) {
		ctx.JSON(400, gin.H{"error": "Invalid OTP 🚫"})
		return errors.New("Not valid")
	}
    
	
	redisClient.Del(emailCtx,redisKey)
	ctx.JSON(200,gin.H{"message":"OTP is verified Successfully 🎉"})
	
	return nil
}