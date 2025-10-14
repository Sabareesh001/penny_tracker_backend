package alter_alert

import (
	"net/http"

	"github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
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
