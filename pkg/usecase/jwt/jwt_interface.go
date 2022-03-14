package jwt

import "github.com/dgrijalva/jwt-go"

type InputPort interface {
	GenerateToken(userId string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
