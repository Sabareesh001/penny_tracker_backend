package bcrypt

import "golang.org/x/crypto/bcrypt"

func BcryptGetHash(pass string) string {
	hashedPass,err := bcrypt.GenerateFromPassword([]byte(pass),10)
	if(err!=nil){
		panic("Could not hash Password")
	}
	return string(hashedPass)
}