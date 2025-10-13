package saving

import (
	alter_monthly_saving_target "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/saving/alter"
	fetch_monthly_saving_target "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/saving/fetch"
	authmiddleware "github.com/Sabareesh001/penny_tracker_backend/pkg/jwt/auth-middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SavingRoutes(router *gin.RouterGroup, DB *gorm.DB) {
    savingRoutes := router.Group("/saving",authmiddleware.AuthorizeJWT)
	fetch_monthly_saving_target.FetchMonthlySavingTarget(savingRoutes,DB)
	alter_monthly_saving_target.AlterMonthlySavingTarget(savingRoutes,DB)
}
