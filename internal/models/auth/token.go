package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	AccessToken  string `json:"accessToken" validate:"required" bson:"accessToken"`
	RefreshToken string `json:"refreshToken" validate:"required" bson:"refreshToken"`
}

type CustomClaims struct {
	Name  string `json:"name" validate:"required" bson:"name"`
	Email string `json:"email" validate:"required" bson:"email"`
	Role  Role   `json:"role" validate:"required" bson:"role"`
	jwt.RegisteredClaims
}
