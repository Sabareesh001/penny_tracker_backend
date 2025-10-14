package add_payment

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

func AddPayment(router *gin.RouterGroup, DB *gorm.DB) {
	router.POST("/add", func(ctx *gin.Context) {

        UserId,_ := userId.GetUserId(ctx)

		var req struct {
			Category    int     `json:"category" binding:"required"`
			Label       string  `json:"label" binding:"required"`
			Amount      float64 `json:"amount" binding:"required"`
			Currency string `json:"currency" binding:"required"`
		}

		// Bind and validate incoming JSON
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid input: " + err.Error(),
			})
			return
		}
        UserIdNum,err := strconv.Atoi(UserId);
		if(err!=nil){
			response.SomethingWentWrong(ctx);
			return;
		}
		// Create the payment record
		payment := payments.Payments{
			User: UserIdNum,
			Category:    req.Category,
			Label:       req.Label,
			Amount:      req.Amount,
			Currency: req.Currency,
			Datetime:    time.Now(),
		}

		// Insert into DB
		if err := DB.Create(&payment).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to add payment: " + err.Error(),
			})
			return
		}

		// Success response
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Payment added successfully",
			"data":    payment,
		})
	})
}
