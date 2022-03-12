package util

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var jwtSecret = []byte("rahasia123")

type Claims struct {
	jwt.StandardClaims
}

func GenerateToken(id string) (string, error) {
	claims := Claims{}
	claims.Id = id
	claims.Issuer = "kusnandardev"
	claims.ExpiresAt = int64(60*time.Second) + time.Now().Unix()

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func Hash(text string) (string, error) {
	pwd := []byte(text)

	hashedPwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return text, err
	}
	return string(hashedPwd), nil
}

func Compare(hashed string, text string) error {
	hashedPwd := []byte(hashed)
	pwd := []byte(text)

	err := bcrypt.CompareHashAndPassword(hashedPwd, pwd)
	if err != nil {
		return err
	}
	return nil
}
