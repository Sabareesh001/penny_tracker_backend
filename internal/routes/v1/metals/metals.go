package metals

import (
	alter_metal_resource "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/metals/alterMetalResource"
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/metals/fetch"
	get_price "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/metals/getPrice"

	authmiddleware "github.com/Sabareesh001/penny_tracker_backend/pkg/jwt/auth-middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MetalRoutes(router *gin.RouterGroup, DB *gorm.DB) {
		metalRoutes := router.Group("metal",authmiddleware.AuthorizeJWT)
		get_price.GetPrice(metalRoutes,DB);
		alter_metal_resource.AlterMetalResource(metalRoutes,DB)
		fetch.FetchAllMetals(metalRoutes,DB);
}