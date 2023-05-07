package middlewares

import (
	"fmt"
	"strings"
	"time"

	"github.com/dimasyudhana/event-booking-app/config/common"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(id uint, email, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.MapClaims{
		"exp":      expirationTime.Unix(),
		"id":       id,
		"email":    email,
		"username": username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(common.JWTSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

type Claims struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (c *Claims) Valid() error {
	if c.ID == 0 || c.Email == "" || c.Username == "" {
		return fmt.Errorf("invalid claims")
	}
	return nil
}

func ValidateJWT(authHeader string) (*Claims, error) {
	if authHeader == "" {
		return nil, fmt.Errorf("missing Authorization header")
	}
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.JWTSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid or expired token")
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid or expired token")
	}
}
