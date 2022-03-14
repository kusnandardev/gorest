package middleware

import (
	"RestGo/pkg/usecase/jwt"
	djwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "auth required"})
			c.Abort()
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := jwt.NewJWTInteractor().ValidateToken(tokenString)

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token, please re-login"})
			c.Abort()
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		c.Set("claims", token.Claims.(djwt.MapClaims))
		c.Next()
	}
}
