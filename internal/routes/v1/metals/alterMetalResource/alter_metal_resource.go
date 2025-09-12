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
		
		
		existingEntry := metals.UserMetalTracking{}
		
		existingRecordFetch := DB.Where("user=? AND metal = ?",UserId,bodyContent.Metal_id).Find(&existingEntry)
		
		if(existingRecordFetch.Error != nil){
			newEntry := metals.UserMetalTracking{ User: userIdInInt,Metal: bodyContent.Metal_id,Weight:float64(bodyContent.Weight),Status: metals.StatusEnabled}
			newEntryQuery := DB.Save(&newEntry)
			if(newEntryQuery.Error == nil){
				 response.SuccesfullyInserted(ctx)
				 return
			}else{
			     response.SomethingWentWrong(ctx)
				 return	
			}
		}

		existingEntry.Weight = float64(bodyContent.Weight)

        updateRecord := DB.Save(&existingEntry)

        if(updateRecord.Error!=nil){
			response.SomethingWentWrong(ctx)
			return
		}

		response.SuccesfullyUpdated(ctx)


	})
}