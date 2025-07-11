package userId

import (
	"strconv"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
)

func GetUserId(ctx *gin.Context) (string,bool){
	userId,exists := ctx.Get("userId")
	if(!exists){
		response.DataInAdequate(ctx)
		return "",false
	}
	UserId :=  strconv.Itoa(int(userId.(float64)))
	return UserId,true
}