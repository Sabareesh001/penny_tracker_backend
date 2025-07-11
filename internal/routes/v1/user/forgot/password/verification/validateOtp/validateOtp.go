package validateotp


import (
	"context"
	"time"
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/verification/email/validate"
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/body/email"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)


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

	 })
}