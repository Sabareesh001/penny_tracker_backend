package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SomethingWentWrong(ctx *gin.Context){
	ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":"Something Went Wrong 😖"})
}

func DataInAdequate(ctx *gin.Context){
	ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"error":"Data Inadequate 🗑️"})
}

func NoMatchingRecords(ctx *gin.Context){
	ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"error":"No Matching Records 🥹"})
}

func NoSuchUserExist(ctx *gin.Context){
	ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"error":"No Such User Exist 🚫"})
}

func UnauthorizedAccess(ctx *gin.Context){
	ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error":"Unauthorized Access 🚫"})
}

func SuccesfullyInserted(ctx *gin.Context){
	ctx.AbortWithStatusJSON(http.StatusCreated,gin.H{"message":"Sucessfully Inserted ✅"})
}

func SuccesfullyUpdated(ctx *gin.Context){
	ctx.AbortWithStatusJSON(http.StatusAccepted,gin.H{"message":"Successfully Updated ✅"})
}