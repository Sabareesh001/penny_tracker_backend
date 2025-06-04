package changepassword

import (
	"context"

	userModel "github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/hashing/bcrypt"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func ChangePassword(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client) {
	router.PUT("/change-password",func(ctx *gin.Context) {
		changePassCtx := context.Background()
		type Body struct {
			Email string `json:"email"`
			Username string `json:"username"`
			NewPassword string `json:"newPassword"`
		}

		body := Body{}

		parsingError := ctx.ShouldBindJSON(&body)

        if(parsingError!=nil || body.Email=="" || body.NewPassword==""){
			response.DataInAdequate(ctx)
			return
		}
         
        redisKey := body.Email+":forgot"
		isVerified := redisClient.Get(changePassCtx,redisKey)

        if(isVerified.Err()!=nil){
			response.SomethingWentWrong(ctx)
			return
		}

		if isVerified.Val()!="verified" {
           response.SomethingWentWrong(ctx)
		   return
		}

		model := userModel.User{}

		query := DB.Where("username=? AND email=?",body.Username,body.Email).First(&model)
        if(query.Error!=nil){
			if(query.Error.Error()=="record not found"){
				response.NoMatchingRecords(ctx)
				return;
			}
			response.SomethingWentWrong(ctx)
		}
        model.Password = bcrypt.BcryptGetHash(body.NewPassword)
        updatePassword := DB.Save(&model)
		if(updatePassword.Error != nil){
			response.SomethingWentWrong(ctx)
			return
		}
		ctx.AbortWithStatusJSON(200,gin.H{"message":"Successfully Updated Password ✅"})
		redisClient.Del(changePassCtx,redisKey)
	})
}