package currency_routes

import (
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/currency/convert"
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/currency/fetch"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CurrencyRoutes(router *gin.RouterGroup, DB *gorm.DB) {
	currencyRoutes := router.Group("/currency")
    fetch.GetCurrencies(currencyRoutes,DB)
	convert.ConvertCurrency(currencyRoutes,DB)
}