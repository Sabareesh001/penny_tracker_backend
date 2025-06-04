package email

import (
	"context"
	"strconv"
	"time"

	userModel "github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/email"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/otp"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
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

			type Body struct{
				Email string `json:"email"`
			}
			body  := Body{}
			err := ctx.ShouldBindJSON(&body)
			if err!=nil {
				response.DataInAdequate(ctx)
				return
			}
			
			var existingList []userModel.User;
			DB.Where("email = ?",body.Email).Find(&existingList)
			
			if len(existingList)==0 {
				response.UnauthorizedAccess(ctx)
				return
			}

			emailCtx := context.Background()
            key := body.Email+"user:"+UserId
            keyAlreadyExist := redisClient.Get(emailCtx,key)
			if keyAlreadyExist.Err()!=redis.Nil {
              ctx.JSON(400,gin.H{"error":"Wait for sometime to request OTP again ⌚"})
			  return
			}

			generatedOTP := otp.GetOtpWithNumbers(4)

			redisClient.Set(emailCtx,key,generatedOTP,30*time.Second)
            
            msg,err := email.ComposeEmail(body.Email,"OTP for Penny Tracker Registration","\nYour OTP is "+strconv.Itoa(generatedOTP)+"\n\n"+"Note : This Otp is only Valid for 40 seconds")
          
           if(err!=nil){
			ctx.JSON(500,gin.H{"message":msg})
			return
		   }

            ctx.JSON(200,gin.H{"message":"OTP generation Successfull ✅"})

			})
} 

func ValidateOtp(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){
	 router.POST("/validateOtp",func(ctx *gin.Context) {
		
        UserId := userId.GetUserId(ctx)

		type Body struct{
				Email string `json:"email"`
				Otp int `json:"otp"`
			}
		body := Body{}

		parseError := ctx.ShouldBindJSON(&body)
      
		if parseError!=nil{
			response.DataInAdequate(ctx)
			return
		}
		
        model := userModel.User{}

        matchingRecord := DB.Where("id=? AND email=?",UserId,body.Email).First(&model)

        if matchingRecord.RowsAffected == 0 {
			response.NoMatchingRecords(ctx)
			return
		}

		if matchingRecord.Error != nil {
			response.SomethingWentWrong(ctx)
			return
		}

        key := body.Email+"user:"+UserId
		emailCtx := context.Background()

	    originalOtp := redisClient.Get(emailCtx,key)

		if originalOtp.Err() == redis.Nil || originalOtp.Val() != strconv.Itoa(body.Otp) {
            ctx.JSON(400,gin.H{"error":"Invalid OTP 🚫"})
			return;
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