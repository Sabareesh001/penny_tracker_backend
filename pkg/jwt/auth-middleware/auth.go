package authmiddleware

import (
	"os"
	"strings"

	jwtAuth "github.com/Sabareesh001/penny_tracker_backend/pkg/jwt"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeJWT(ctx *gin.Context){
	SECRET_KEY := os.Getenv("JWT_SECRET")
	tokenStr := ctx.GetHeader("Authorization")
	if(tokenStr==""){
        tokenCookie,err := ctx.Cookie("auth_token")
		if(err!=nil){
			response.UnauthorizedAccess(ctx)
			return
		}
		tokenStr = tokenCookie
	}else{
		tokenStr = strings.TrimPrefix(tokenStr,"Bearer ")
	}
    token,err := jwtAuth.ValidateJwt(tokenStr,SECRET_KEY)
	if(err!=nil){
		response.UnauthorizedAccess(ctx)
		return
             	}
    
	claims := token.Claims.(jwt.MapClaims)

	ctx.Set("userId",claims["userId"].(float64))
}