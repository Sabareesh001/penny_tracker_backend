package fetch_monthly_saving_target

import (
	"net/http"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MonthlySavingTargetData struct {
	MonthlySavingTarget float64 `json:"monthly_saving_target"`
	MonthlySavingCurrency string `json:"monthly_saving_currency"`
}

func FetchMonthlySavingTarget(router *gin.RouterGroup, DB *gorm.DB) {

	router.GET("/target", func(ctx *gin.Context) {
		var row MonthlySavingTargetData
		UserId, contains := userId.GetUserId(ctx)

		if !contains {
			response.SomethingWentWrong(ctx)
			return
		}

		fetchTarget := DB.Table("users").
			Select("monthly_saving_target", "monthly_saving_currency").
			Where("id = ?", UserId).
			First(&row)

		if fetchTarget.Error != nil {
			response.SomethingWentWrong(ctx)
			return
		}

		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "Successfully fetched monthly saving target",
			"data":    row,
		})
	})
}
