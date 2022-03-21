package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateToken(UserId string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func getSecretKey() string {
	secreKey := os.Getenv("SECRET_KEY")

	if secreKey == "" {
		return "yzlhdy"
	}
	return secreKey
}

func NewJwtService() JwtService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "yzlhdy",
	}
}

// GenerateToken 生成token
func (j *jwtService) GenerateToken(userId string) string {
	claims := &jwtCustomClaim{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, error := token.SignedString([]byte(j.secretKey))
	if error != nil {
		panic(error)
	}
	return tokenString

}

// ValidateToken 验证token
func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
