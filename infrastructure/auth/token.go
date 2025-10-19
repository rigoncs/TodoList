package auth

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("very-secret-jwt")

type Claim struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type TokenService interface {
	GenerateToken(ctx context.Context, id uint, username string) (string, error)
	ParseToken(ctx context.Context, token string) (*Claim, error)
}

type JWTTokenService struct{}

func (j JWTTokenService) GenerateToken(ctx context.Context, id uint, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claim := Claim{
		Id:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todolist",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func (j JWTTokenService) ParseToken(ctx context.Context, token string) (*Claim, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claim); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
