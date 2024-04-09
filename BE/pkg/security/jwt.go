package security

import (
	"fmt"
	"thelastking/kingseafood/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const JWT_KEY = "asfdgrhtheerte"

func GenerateAccessToken(data *model.Users) (string, error) {
	fmt.Printf("data.UserID: %v\n", data.UserID)
	claims := &model.AccessToken{
		UserID: data.UserID,
		Role:   data.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return "", fmt.Errorf("generating JWT Access Token failed: %w", err)
	}

	return tokenString, nil
}
func GenerateRefreshToken(data *model.Users) (string, error) {
	refreshClaims := &model.RefreshToken{
		UserID: data.UserID,
		Role:   data.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 90 * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	tokenString, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return "", fmt.Errorf("generating JWT Refresh Token failed: %w", err)
	}
	return tokenString, nil
}
