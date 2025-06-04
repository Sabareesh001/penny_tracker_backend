package generateotp

import (
	"context"
	"strconv"
	"time"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/email"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/otp"
	userModel "github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"
)

func GenerateOtp(ctx *gin.Context, DB *gorm.DB, redisClient *redis.Client,key string,subject string,purpose string,ttl time.Duration) {

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
            redisKey := body.Email+purpose+key
            keyAlreadyExist := redisClient.Get(emailCtx,redisKey)
			if keyAlreadyExist.Err()!=redis.Nil {
              ctx.JSON(400,gin.H{"error":"Wait for sometime to request OTP again ⌚"})
			  return
			}

			generatedOTP := otp.GetOtpWithNumbers(4)

			redisClient.Set(emailCtx,redisKey,generatedOTP,ttl)
            
            msg,err := email.ComposeEmail(body.Email,subject,"\nYour OTP is "+strconv.Itoa(generatedOTP)+"\n\n"+"Note : This Otp is only Valid for "+ttl.String())
          
           if(err!=nil){
			ctx.JSON(500,gin.H{"message":msg})
			return
		   }

            ctx.JSON(200,gin.H{"message":"OTP generation Successfull ✅"})

}