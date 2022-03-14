package jwt

import (
	"RestGo/pkg/domain/repository"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Interactor struct {
	secretKey string
	issuer    string
	cache     repository.Inmemory
}

func NewJWTInteractor(cc repository.Inmemory) *Interactor {
	return &Interactor{
		secretKey: "rahasia",
		issuer:    "kusnandartoni",
		cache:     cc,
	}
}

func (i *Interactor) GenerateToken(userId string) (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Time.Add(time.Now(), 10*time.Minute).Unix(),
		Id:        userId,
		Issuer:    i.issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(i.secretKey))
	if err != nil {
		return "", err
	}
	err = i.cache.Set(t, userId)
	if err != nil {
		return "", err
	}
	return t, nil
}

func (i *Interactor) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		val, err := i.cache.Get(token.Raw)
		if err != nil {
			return nil, err
		}
		fmt.Println(val)
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token ", token.Header["alg"])
		}
		return []byte(i.secretKey), nil
	})
}
