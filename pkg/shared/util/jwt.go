package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetClaims(bearer string) (jwt.StandardClaims, error) {
	tokenString := bearer[len("Bearer "):]
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(jwt.StandardClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return jwt.StandardClaims{}, err
}

func GetToken(c *gin.Context) string {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	return authHeader[len(BEARER_SCHEMA):]
}
