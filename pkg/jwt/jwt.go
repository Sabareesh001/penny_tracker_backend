package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

func AssignJWT(claims jwt.MapClaims,SECRET_KEY []byte) (string,error){

        token  := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

       signedToken,err := token.SignedString(SECRET_KEY);

	   if(err!=nil){
           return "",err
	   }
       return signedToken,nil
}


func ValidateJwt(token string,SECRET_KEY string) (*jwt.Token,error){
     parsedToken,err := jwt.Parse(token,func(t *jwt.Token) (any, error) {
		return []byte(SECRET_KEY),nil
	 })
	 if(err!=nil){
		return nil,err
	 }
    return parsedToken,nil
}