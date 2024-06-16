package auth

import (
	"context"
	"errors"
	"log"
	"scoreboard/config"
	"scoreboard/internal/models/auth"
	"scoreboard/internal/utils/database"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
)

var claims auth.CustomClaims

func GenrateAccessToken(user *auth.User) (string, error) {
	claims = auth.CustomClaims{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), //add config to it later.
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret")) // add config to it later.
}

func GenrateRefreshToken(user *auth.User) (string, error) {
	claims = auth.CustomClaims{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), // add config to it later.
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret")) // add config to it later.
}

func CreateNewToken(user *auth.User) (string, string, error) {
	accessToken, err := GenrateAccessToken(user)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := GenrateRefreshToken(user)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func ParseRefreshToken(refreshToken string) (*auth.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // add config to it later.
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return &claims, nil
}

func ValidateAccessToken(accessToken string) (*auth.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // add config to it later.
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return &claims, nil
}

func ValidateRefreshToken(refreshToken string) (*auth.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // add config to it later.
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return &claims, nil
}

func GetToken(c *gin.Context) string {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		return ""
	}
	token = ExtractToken(token)
	return token
}

func ExtractToken(authHeader string) string {
	// Check if the string starts with "Bearer "
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		// Invalid token
		log.Println("Invalid or missing bearer token")
		return ""
	}

	// Return the token part
	return parts[1]
}

func GetCurrentUser(c *gin.Context) (string, auth.Role, error) {
	token := GetToken(c)
	if token == "" {
		return "", "", nil
	}
	claims, err := ValidateAccessToken(token)
	if err != nil {
		return "", "", err
	}
	var user auth.User
	err = database.GetCollection(config.AppConfig.Collections.Users).FindOne(context.TODO(), bson.M{"email": claims.Email, "isActive": true}).Decode(&user)
	if err != nil {
		return "", "", err
	}
	if user.Email == "" {
		return "", "", errors.New("user not found")
	}
	return user.Email, user.Role, nil
}

func GenerateTokens(name, email string, role auth.Role) (string, string, error) {
	accessToken, err := GenrateAccessToken(&auth.User{Name: name, Email: email, Role: role})
	if err != nil {
		return "", "", err
	}
	refreshToken, err := GenrateRefreshToken(&auth.User{Name: name, Email: email, Role: role})
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
