package helper

import (
	"errors"
	dto "palm_code_be/src/app/dto/user"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("secret_key")

// TokenClaims menyimpan klaim JWT
type TokenClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken membuat token JWT
func GenerateToken(data *dto.RegisterModel) (string, error) {
	expirationTime := time.Now().Add(120 * time.Minute)
	claims := &TokenClaims{
		UserID: data.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// VerifyToken memverifikasi token JWT
func VerifyToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

const (
	Page    = int64(1)
	PerPage = int64(10)
)

var (
	ErrNotFound = errors.New("resource not found")
)
