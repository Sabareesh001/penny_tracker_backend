package fetch_income

import (
	"net/http"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IncomeData struct {
	MonthlyIncome         float64 `json:"monthly_income"`
	MonthlyIncomeCurrency string  `json:"monthly_income_currency"`
}

func FetchIncome(router *gin.RouterGroup, DB *gorm.DB) {

	router.GET("/",func(ctx *gin.Context) {
		var row IncomeData
		UserId,contains := userId.GetUserId(ctx)

        if(!contains){
			response.SomethingWentWrong(ctx)
			return
		}

        fetchIncome := DB.Table("users").Select("monthly_income","monthly_income_currency").Where("id = ?",UserId).First(&row)

		if(fetchIncome.Error != nil){
			response.SomethingWentWrong(ctx)
			return
		}

		ctx.AbortWithStatusJSON(http.StatusOK,gin.H{"message":"Successfully fetched Income","data":row})

	})

}