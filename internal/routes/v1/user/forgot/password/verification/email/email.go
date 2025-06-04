package email

import (
	"context"
	"time"

	generateotp "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/verification/email/generateOtp"
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/verification/email/validate"
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/body/email"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RequestOtp(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client) {
	router.POST("/requestOtp", func(ctx *gin.Context) {
		generateotp.GenerateOtp(ctx, DB, redisClient,"","Forgot Password : Reset Request","forgotPass:",time.Second*50)
	})
}



func ValidateOtp(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client){
	router.POST("/validateOtp",func(ctx *gin.Context) {

	   body := email.Body{}

        err :=  validate.ValidateOtp(ctx,DB,redisClient,"forgotPass:","",&body)
        if(err!=nil){
			ctx.Abort()
			return
		}
        
	    forgotValidationCtx := context.Background()

		redisClient.Set(forgotValidationCtx,body.Email+":forgot","verified",time.Minute*5)

        ctx.JSON(200,gin.H{"success":"OTP is verified Successfully 🎉"})
	 })
}