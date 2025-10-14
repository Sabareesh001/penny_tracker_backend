package fetch_spending

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Sabareesh001/penny_tracker_backend/internal/database/models/payments"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetMonthlyTotal(router *gin.RouterGroup, DB *gorm.DB) {
	router.GET("/monthly-total", func(ctx *gin.Context) {
		// Read month and year from query params
		monthStr := ctx.Query("month")
		yearStr := ctx.Query("year")
        UserId,_ := userId.GetUserId(ctx)
		month, err1 := strconv.Atoi(monthStr)
		year, err2 := strconv.Atoi(yearStr)

		if err1 != nil || err2 != nil || month < 1 || month > 12 {
			response.DataInAdequate(ctx)
			return
		}

		// Calculate start and end of month
		startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		endDate := startDate.AddDate(0, 1, 0) // first day of next month

		var total float64
		if err := DB.Model(&payments.Payments{}).
			Where("datetime >= ? AND datetime < ? AND user=?", startDate, endDate,UserId).
			Select("SUM(amount)").Scan(&total).Error; err != nil {
			response.SomethingWentWrong(ctx)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Total payments fetched successfully",
			"data": gin.H{
				"month": month,
				"year":  year,
				"total": total,
			},
		})
	})
}
