package spending

import (
	alter_alert "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/spending/alterAlert"
	fetch_spending "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/spending/fetch"
	authmiddleware "github.com/Sabareesh001/penny_tracker_backend/pkg/jwt/auth-middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func SpendingRoutes(router *gin.RouterGroup,DB *gorm.DB){
    spendingRoutes := router.Group("/spending",authmiddleware.AuthorizeJWT)
	fetch_spending.GetMonthlyTotal(spendingRoutes,DB)
	alter_alert.GetAlertPercentage(spendingRoutes,DB)
	alter_alert.PatchAlertPercentage(spendingRoutes,DB)
	alter_alert.SendAlertEmail(spendingRoutes,DB)
}

