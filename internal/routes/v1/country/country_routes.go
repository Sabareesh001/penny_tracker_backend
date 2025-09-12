package country_routes

import (
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/country/fetch"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CountryRoutes(router *gin.RouterGroup, DB *gorm.DB) {
	countryRoutes := router.Group("/country")
    fetch.GetCountry(countryRoutes,DB);
}