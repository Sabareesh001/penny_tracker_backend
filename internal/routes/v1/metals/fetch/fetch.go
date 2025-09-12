package fetch

import (
	"net/http"

	"github.com/Sabareesh001/penny_tracker_backend/internal/database/models/metals"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FetchAllMetals(router *gin.RouterGroup, DB *gorm.DB){

	
	router.GET("/",func(ctx *gin.Context) {
		
		UserID,contains := userId.GetUserId(ctx);

		if(!contains){
			response.UnauthorizedAccess(ctx)
			return
		}

        type MetalData struct {
			metals.Metals
			Status string
		}

	    var rows []MetalData
		fetchAllMetals := DB.Table("metals").
	Select("metals.*, COALESCE(umt.status, '0') as status").
	Joins("LEFT JOIN user_metal_trackings AS umt ON umt.metal = metals.id AND umt.user = "+UserID).
	Scan(&rows)
		if(fetchAllMetals.Error != nil){
			response.SomethingWentWrong(ctx)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusOK,gin.H{"message":"sucessfully fetched metals","data":rows})
	})

}