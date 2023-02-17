package jwt

import (
	"github.com/golang-jwt/jwt"
	"log"
)

// JWT signing Key
type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	Id          int64
	AuthorityId int64
	jwt.StandardClaims
}

// ParseToken parses the token.
// TODO_hewen: 完善token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString, &CustomClaims{}, 
		func(token *jwt.Token) (interface{}, error) {
			return j.SigningKey, nil})
	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return nil, err
	}
	claims, _ := token.Claims.(*CustomClaims)
	return claims, nil
}