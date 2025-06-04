package email

import (
	"time"

	userModel "github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/body/email"
	generateotp "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/verification/email/generateOtp"
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/verification/email/validate"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func EmailVerification(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){
    emailRoutes := router.Group("/email")
	RequestOtp(emailRoutes,DB,redisClient)
	ValidateOtp(emailRoutes,DB,redisClient)
}

func RequestOtp(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){
        
		router.POST("/requestOtp",func(ctx *gin.Context) {
		   UserId := userId.GetUserId(ctx)
           generateotp.GenerateOtp(ctx,DB,redisClient,UserId,"OTP for Penny Tracker Registration","user:",time.Second*40)
		})
} 

func ValidateOtp(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){
	
	 router.POST("/validateOtp",func(ctx *gin.Context) {
		
        UserId := userId.GetUserId(ctx)

        model := userModel.User{}
     
        body := email.Body{}

        err :=  validate.ValidateOtp(ctx,DB,redisClient,"user:",UserId,&body)
        if(err!=nil){
			ctx.Abort()
			return
		}
        model.IsEmailVerified = "1"

		updateEmailVerification := DB.Save(&model)

        if updateEmailVerification.Error != nil {
			response.SomethingWentWrong(ctx)
			return
		}

        ctx.JSON(200,gin.H{"success":"OTP is verified Successfully 🎉"})
	 })
}

		
