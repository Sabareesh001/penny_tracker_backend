package gender_routes

import (
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/gender/fetch"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GenderRoutes(router *gin.RouterGroup, DB *gorm.DB) {
	genderRoutes := router.Group("/gender")
    fetch.GetGender(genderRoutes,DB);
}