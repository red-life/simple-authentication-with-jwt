package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type IJWTPackage interface {
	CreateToken(email string)  (string,error)
	ValidateToken(token string) (string, error)
}

func NewJWTPackage(SecretKey string) IJWTPackage{
	return &JWTService{
		SecretKey: SecretKey,
	}
}

type JWTService struct {
	SecretKey string
}

func (J JWTService) CreateToken(email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)


	signedToken, err := token.SignedString([]byte(J.SecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (J JWTService) ValidateToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(J.SecretKey), nil
	})
	if err != nil {
		return "",err
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return payload["email"].(string), nil
	}

	return "",errors.New("invalid token")
}



