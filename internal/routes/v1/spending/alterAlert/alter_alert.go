package alter_alert

import (
	"fmt"
	"net/http"

	"github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/email"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAlertPercentage(router *gin.RouterGroup, DB *gorm.DB) {
	router.GET("/alert-percentage", func(ctx *gin.Context) {

		UserId, _ := userId.GetUserId(ctx)

		var u user.User
		if err := DB.Select("alert_percentage").First(&u, "id = ?", UserId).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				response.DataInAdequate(ctx)
				return
			}
			response.SomethingWentWrong(ctx)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Alert percentage fetched successfully",
			"data":u.Alert_Percentage,
		})
	})
}

func PatchAlertPercentage(router *gin.RouterGroup, DB *gorm.DB) {
	router.PATCH("/alert-percentage", func(ctx *gin.Context) {

		UserId, _ := userId.GetUserId(ctx)

		var body struct {
			Alert_Percentage float32 `json:"alert_percentage"`
		}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			response.DataInAdequate(ctx)
			return
		}

		if body.Alert_Percentage < 0 || body.Alert_Percentage > 100 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Alert percentage must be between 0 and 100",
			})
			return
		}

		if err := DB.Model(&user.User{}).
			Where("id = ?", UserId).
			Update("alert_percentage", body.Alert_Percentage).Error; err != nil {

			response.SomethingWentWrong(ctx)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Alert percentage updated successfully",
			"data": gin.H{
				"alert_percentage": body.Alert_Percentage,
			},
		})
	})
}

	func SendAlertEmail(router *gin.RouterGroup, DB *gorm.DB) {
	router.POST("/alert/send", func(ctx *gin.Context) {
		UserId, _ := userId.GetUserId(ctx)

		// Fetch user's name and email
		var u user.User
		if err := DB.Select("email", "first_name").First(&u, "id = ?", UserId).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				response.DataInAdequate(ctx)
				return
			}
			response.SomethingWentWrong(ctx)
			return
		}

		// Predefined subject and body
		subject := "⚠️ Spending Threshold Alert from Penny Tracker"
		body := fmt.Sprintf(`
Hello %s,

This is an automatic alert from Penny Tracker.

Your spending balance has fallen below your configured safe threshold.
This means you’ve reached the alert limit you set for your monthly spending goal.

Take a moment to review your recent expenses and consider adjusting your spending or saving targets to stay on track.

Stay smart with your money 💰,
— The Penny Tracker Team
`, u.First_Name)

		// Send email
		msg, err := email.ComposeEmail(u.Email, subject, body)
		if err != nil {
			response.SomethingWentWrong(ctx)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": msg,
			"data": gin.H{
				"email":   u.Email,
				"subject": subject,
			},
		})
	})
}