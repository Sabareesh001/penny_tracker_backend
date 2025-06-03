package email

import(
	"github.com/Sabareesh001/penny_tracker_backend/pkg/otp"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/email"
	userModel "github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"context"
	"strconv"
	"time"
)

func EmailVerification(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){
    emailRoutes := router.Group("/email")
	RequestOtp(emailRoutes,DB,redisClient)
	ValidateOtp(emailRoutes,DB,redisClient)
}

func RequestOtp(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){
      
		router.POST("/requestOtp",func(ctx *gin.Context) {

			
			type Body struct{
				Email string `json:"email"`
				UserId int `json:"userId"`
			}
			body  := Body{}
			err := ctx.ShouldBindJSON(&body)
			if err!=nil {
				ctx.JSON(400,gin.H{"error":"Data Inadequate"})
				return
			}
			
			var existingList []userModel.User;
			DB.Where("email = ?",body.Email).Find(&existingList)
			
			if(len(existingList)>0){
				ctx.JSON(400,gin.H{"error":"Mobile Number Already Used"})
				return
			}

			emailCtx := context.Background()
            key := body.Email+"user:"+strconv.Itoa(body.UserId)
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
		type Body struct{
				Email string `json:"email"`
				UserId int `json:"userId"`
				Otp int `json:"otp"`
			}
		body := Body{}

		ctx.ShouldBindJSON(&body)
      
        key := body.Email+"user:"+strconv.Itoa(body.UserId)
		emailCtx := context.Background() 

	    originalOtp := redisClient.Get(emailCtx,key)



		if originalOtp.Err() == redis.Nil || originalOtp.Val() != strconv.Itoa(body.Otp) {
            ctx.JSON(400,gin.H{"error":"Invalid OTP 🚫"})
			return;
		}

        ctx.JSON(200,gin.H{"success":"OTP is verified Successfully 🎉"})
	 })
}