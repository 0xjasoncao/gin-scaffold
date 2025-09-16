package token

import (
	"context"
	"github.com/0xjasoncao/gin-scaffold/pkg/cache"
	"github.com/0xjasoncao/gin-scaffold/pkg/errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Service interface {
	IssuingToken(ctx context.Context, userId string) (string, error)
	DestroyToken(ctx context.Context, token string) error
	Parse(ctx context.Context, token string) (*Claims, error)
}

// Claims represents the custom JWT claims structure
type Claims struct {
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}

// DefaultKey is the fallback HMAC secret key used if no key is provided in settings
// This should be replaced with a secure key in production environments
const DefaultKey = "bzOojgQuL6notiVZHnYIe3MeayB5cRdtHKUlqDltb9A="

// NewTokenService creates a new instance of the JWT token service
func NewTokenService(settings *Settings, store Store) (Service, error) {
	if settings.Key == "" {
		return nil, errors.New("jwt-key cannot be empty")
	}

	return &JwtToken{Settings: settings, Store: store}, nil

}

// Settings accessToken setting
type Settings struct {
	ExpiresAtSeconds int
	Key              string
	Issuer           string
	cache.Cache
}

type JwtToken struct {
	*Settings
	Store
}

// IssuingToken generates a signed JWT token for the specified user ID
// It uses HS256 algorithm for signing and includes both custom and standard claims
func (j *JwtToken) IssuingToken(ctx context.Context, userId string) (string, error) {
	claims := Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.ExpiresAtSeconds) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.Key))

}

// DestroyToken handles token invalidation
func (j *JwtToken) DestroyToken(ctx context.Context, tokenStr string) error {
	claims, err := j.Parse(ctx, tokenStr)
	if err != nil {
		return err
	}
	expired := time.Until(time.Unix(claims.ExpiresAt.Unix(), 0))

	return j.Store.Set(ctx, tokenStr, expired)

}

// Parse verifies a JWT token string and extracts the custom claims
func (j *JwtToken) Parse(ctx context.Context, tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.Key), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.Wrap(err, "invalid token")
}
