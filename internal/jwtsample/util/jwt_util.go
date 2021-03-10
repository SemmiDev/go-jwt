package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type Data struct {
	Name string
	Identifier string
	Email string
}

func secretKey() []byte {
	var secretKeyBytes = []byte("sammidev")
	return secretKeyBytes
}

func JwtBuildAndSignJSON(data Data) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": data.Name,
		"identifier": data.Identifier,
		"email": data.Email,
	})

	tokenString, _ := token.SignedString(secretKey())
	return tokenString
}

func JwtValidate(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey(), nil
	})

	var data Data
	if token != nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			var userNameClaim = claims["name"].(string)
			var userIdentifirrClaim = claims["identifier"].(string)
			var userEmailClaim = claims["email"].(string)

			data = Data{userNameClaim, userIdentifirrClaim, userEmailClaim}
		}
	}
	return data, err
}
