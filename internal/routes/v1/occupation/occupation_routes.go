package occupation_routes

import (
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/occupation/fetch"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OccupationRoutes(router *gin.RouterGroup, DB *gorm.DB) {
	occupationRoutes := router.Group("/occupation")
    fetch.GetOccupation(occupationRoutes,DB);
}