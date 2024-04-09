package model

import "github.com/dgrijalva/jwt-go"

type AccessToken struct {
	UserID string `json:"user_id"`
	Role   string `json:"-"`
	jwt.StandardClaims
}

type RefreshToken struct {
	UserID string `json:"user_id"`
	Role   string `json:"-"`
	jwt.StandardClaims
}
