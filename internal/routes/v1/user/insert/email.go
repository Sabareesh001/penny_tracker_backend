package email

import (
	"fmt"

	userModel "github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/responses" 
)

func InsertEmail(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client) {
   router.POST("/email",func(ctx *gin.Context) {
	type Body struct {
		UserId  int `form:"userId" binding:"required"`
		Email string `form:"email" binding:"required"`
	}

    body := Body{}
    model := userModel.User{}
	ctx.ShouldBind(&body)

	matchingRecord := DB.Where("id=?",body.UserId).First(&model)
    model.Email = body.Email
    fmt.Println(model)
	isUpdated:=DB.Save(&model)

    if matchingRecord.Error != nil || isUpdated.Error!=nil {
		response.SomethingWentWrong(ctx)
		return
	}

    ctx.JSON(200,gin.H{"message":"Email Updated Successfully ✅"})

   })
}