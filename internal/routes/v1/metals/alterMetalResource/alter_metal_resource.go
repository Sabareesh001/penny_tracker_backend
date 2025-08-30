package alter_metal_resource

import (
	"encoding/json"
	"io"
	"strconv"
	"github.com/Sabareesh001/penny_tracker_backend/internal/database/models/metals"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AlterMetalResource(router *gin.RouterGroup, DB *gorm.DB){
	router.POST("/weight",func(ctx *gin.Context) {
		type Body struct {
			Metal_id int `json:"metal_id"`
			Weight float32 `json:"weight"`
		}

		UserId,contains := userId.GetUserId(ctx);

		if(!contains){
            response.SomethingWentWrong(ctx);
			return
		}

		bodyContent := Body{}

		
		body,err := io.ReadAll(ctx.Request.Body)
		
		if(err!=nil){
			response.SomethingWentWrong(ctx)
			return
		}
		
        err = json.Unmarshal(body,&bodyContent)
		
		if(err!=nil){
			response.SomethingWentWrong(ctx)
			return
		}

		userIdInInt,err := strconv.Atoi(UserId);
		
		if(err!=nil){
			response.UnauthorizedAccess(ctx);
		}
		
		newEntry := metals.UserMetalTracking{ User: userIdInInt,Metal: bodyContent.Metal_id,Weight:float64(bodyContent.Weight) }
		
		newEntryQuery := DB.Save(&newEntry)
		
		if(newEntryQuery.Error == nil){
			 response.SuccesfullyInserted(ctx)
			 return
		}

		existingRecord := metals.UserMetalTracking{};

		findExistingRecord := DB.Where("metal=? AND user=?",bodyContent.Metal_id,UserId).Find(&existingRecord)

		if(findExistingRecord.Error!=nil){
			response.SomethingWentWrong(ctx)
			return
		}

	})
}