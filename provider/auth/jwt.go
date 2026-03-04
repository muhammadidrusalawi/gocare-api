package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/muhammadidrusalawi/gocare-api/provider/cache"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID string, email string, role string) (string, error) {
	expireStr := os.Getenv("JWT_EXPIRE")

	expire, err := time.ParseDuration(expireStr)
	if err != nil {
		expire = 24 * time.Hour
	}

	now := time.Now()
	expTime := now.Add(expire)

	claims := jwt.MapClaims{
		"jti":     uuid.NewString(),
		"user_id": userID,
		"email":   email,
		"role":    role,
		"iat":     now.Unix(),
		"nbf":     now.Unix(),
		"exp":     expTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

func RevokeToken(tokenString string) error {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return err
	}

	jti, ok := claims["jti"].(string)
	if !ok {
		return errors.New("invalid jti")
	}

	expFloat, ok := claims["exp"].(float64)
	if !ok {
		return errors.New("invalid exp")
	}

	expTime := time.Unix(int64(expFloat), 0)
	ttl := time.Until(expTime)

	if ttl <= 0 {
		return nil
	}

	key := "blacklist:" + jti

	return cache.Client.Set(cache.Ctx, key, "revoked", ttl).Err()
}

func IsTokenRevoked(jti string) (bool, error) {
	key := "blacklist:" + jti

	result, err := cache.Client.Get(cache.Ctx, key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return false, nil
		}
		return false, err
	}

	return result == "revoked", nil
}
