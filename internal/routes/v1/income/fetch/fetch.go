package fetch_income

import (
	"net/http"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FetchIncome(router *gin.RouterGroup, DB *gorm.DB) {

	router.GET("/",func(ctx *gin.Context) {
		var row float32
		UserId,contains := userId.GetUserId(ctx)

        if(!contains){
			response.SomethingWentWrong(ctx)
			return
		}

        fetchIncome := DB.Table("users").Select("monthly_income").Where("id = ?",UserId).Find(&row)

		if(fetchIncome.Error != nil){
			response.SomethingWentWrong(ctx)
			return
		}

		ctx.AbortWithStatusJSON(http.StatusOK,gin.H{"message":"Successfully fetched Income","data":row})

	})

}