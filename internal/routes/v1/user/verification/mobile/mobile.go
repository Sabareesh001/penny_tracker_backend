package mobile

import (
	"context"
	"strconv"
	"time"

	userModel "github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/otp"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func MobileVerification(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){
    mobileRoutes := router.Group("/mobile")
	RequestOtp(mobileRoutes,DB,redisClient)
	ValidateOtp(mobileRoutes,DB,redisClient)
}

func RequestOtp(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){
      
		router.POST("/requestOtp",func(ctx *gin.Context) {

			
			type Body struct{
				Mobile string `json:"mobile"`
				UserId int `json:"userId"`
			}
			body  := Body{}
			err := ctx.ShouldBindJSON(&body)
			if err!=nil {
				response.DataInAdequate(ctx)
				return
			}
			
			var existingList []userModel.User;
			DB.Where("phone = ?",body.Mobile).Find(&existingList)
			
			if(len(existingList)>0){
				ctx.JSON(400,gin.H{"error":"Mobile Number Already Used"})
				return
			}

			mobileCtx := context.Background()
            key := body.Mobile+"user:"+strconv.Itoa(body.UserId)
            keyAlreadyExist := redisClient.Get(mobileCtx,key)
			if keyAlreadyExist.Err()!=redis.Nil {
              ctx.JSON(400,gin.H{"error":"Wait for sometime to request OTP again ⌚"})
			  return
			}
			redisClient.Set(mobileCtx,key,otp.GetOtpWithNumbers(4),30*time.Second)

            ctx.JSON(200,gin.H{"message":"OTP generation Successfull ✅"})

			})
} 

func ValidateOtp(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){
	 router.POST("/validateOtp",func(ctx *gin.Context) {
		type Body struct{
				Mobile string `json:"mobile"`
				UserId int `json:"userId"`
				Otp int `json:"otp"`
			}
		body := Body{}

		ctx.ShouldBindJSON(&body)
      
        key := body.Mobile+"user:"+strconv.Itoa(body.UserId)
		mobileCtx := context.Background() 

	    originalOtp := redisClient.Get(mobileCtx,key)



		if originalOtp.Err() == redis.Nil || originalOtp.Val() != strconv.Itoa(body.Otp) {
            ctx.JSON(400,gin.H{"error":"Invalid OTP 🚫"})
			return;
		}

        ctx.JSON(200,gin.H{"success":"OTP is verified Successfully 🎉"})
	 })
}