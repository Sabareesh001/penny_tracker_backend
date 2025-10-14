package fetch_payment

import (
	"net/http"

	"github.com/Sabareesh001/penny_tracker_backend/internal/database/models/payments"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FetchPayments(router *gin.RouterGroup, DB *gorm.DB) {
	router.GET("/all", func(ctx *gin.Context) {
		var allPayments []payments.Payments
        UserId,_ := userId.GetUserId(ctx);
		// Fetch all payments from DB
		if err := DB.Where("user=?",UserId).Find(&allPayments).Error; err != nil {
			response.SomethingWentWrong(ctx);
			return
		}

		// Success response
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Payments fetched successfully",
			"data":    allPayments,
		})
	})
}
