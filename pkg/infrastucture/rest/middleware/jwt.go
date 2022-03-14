package middleware

import (
	"RestGo/pkg/adapter/db/inmemory"
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
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "auth required"})
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := jwt.NewJWTInteractor(inmemory.DefaultCacheHandler()).ValidateToken(tokenString)

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token, please re-login"})
			return
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.Set("claims", token.Claims.(djwt.MapClaims))
		c.Next()
	}
}
