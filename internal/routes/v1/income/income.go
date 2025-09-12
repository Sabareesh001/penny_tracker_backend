package income

import (
	alter_income "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/income/alter"
	fetch_income "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/income/fetch"
	authmiddleware "github.com/Sabareesh001/penny_tracker_backend/pkg/jwt/auth-middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IncomeRoutes(router *gin.RouterGroup, DB *gorm.DB) {
    incomeRoutes := router.Group("/income",authmiddleware.AuthorizeJWT)
	fetch_income.FetchIncome(incomeRoutes,DB)
	alter_income.AlterIncome(incomeRoutes,DB)
}
