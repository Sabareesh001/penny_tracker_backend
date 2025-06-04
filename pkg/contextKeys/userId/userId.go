package userId

import (
	"strconv"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
)

func GetUserId(ctx *gin.Context) string{
	userId,exists := ctx.Get("userId")
	UserId :=  strconv.Itoa(int(userId.(float64)))
		if(!exists){
			response.SomethingWentWrong(ctx)
			ctx.Abort()
		}
	return UserId
}