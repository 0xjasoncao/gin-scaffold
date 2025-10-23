package token

import (
	"context"
	"time"

	"gin-scaffold/pkg/errorsx"
	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	IssuingToken(ctx context.Context, payload Payload) (*IssuingTokenInfo, error)
	DestroyToken(ctx context.Context, token string) error
	Parse(ctx context.Context, token string) (*Claims, error)
}

// Payload custom payload
type Payload struct {
	UserID      uint64 `json:"userID"`
	AuthorityId uint   `json:"authority_id"`
}

// Claims represents the custom JWT claims structure
type Claims struct {
	Payload
	jwt.RegisteredClaims
}

// DefaultKey is the fallback HMAC secret key used if no key is provided in settings
// This should be replaced with a secure key in production environments
const (
	DefaultKey              = "bzOojgQuL6notiVZHnYIe3MeayB5cRdtHKUlqDltb9A="
	DefaultExpiresAtSeconds = 3600
)

type jwtToken struct {
	*Settings
	Store
}

// NewTokenService creates a new instance of the JWT token service
func NewTokenService(settings *Settings, store Store) (Service, error) {
	if settings.Key == "" {
		settings.Key = DefaultKey
	}
	if settings.ExpiresAtSeconds <= 0 {
		settings.ExpiresAtSeconds = DefaultExpiresAtSeconds
	}

	return &jwtToken{Settings: settings, Store: store}, nil

}

// Settings defines the configuration for token generation.
// It includes expiration settings, signing key, issuer, and a cache backend.
type Settings struct {
	ExpiresAtSeconds int    // Token expiration time in seconds
	Key              string // Signing key for the token
	Issuer           string // Token issuer identifier
}

// IssuingTokenInfo represents detailed information about an issued token.
// It includes issue time, expiration time, and the access token itself.
type IssuingTokenInfo struct {
	ExpiresAt   int64  `json:"expires_at"`             // Expiration timestamp (Unix seconds)
	AccessToken string `json:"access_token,omitempty"` // Access token string
	IssuedAt    int64  `json:"issued_at"`              // Issue timestamp (Unix seconds)
}

// IssuingToken generates a signed JWT token
// It uses HS256 algorithm for signing and includes both custom and standard claims
func (j *jwtToken) IssuingToken(ctx context.Context, payload Payload) (*IssuingTokenInfo, error) {

	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Duration(j.ExpiresAtSeconds) * time.Second))
	issuedAt := jwt.NewNumericDate(time.Now())
	claims := Claims{
		Payload: Payload{
			UserID:      payload.UserID,
			AuthorityId: payload.AuthorityId,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString([]byte(j.Key))
	if err != nil {
		return nil, err
	}

	tokenInfo := &IssuingTokenInfo{
		ExpiresAt:   expiresAt.Time.Unix(),
		IssuedAt:    issuedAt.Time.Unix(),
		AccessToken: signedString,
	}
	return tokenInfo, nil

}

// DestroyToken handles token invalidation
func (j *jwtToken) DestroyToken(ctx context.Context, tokenStr string) error {
	claims, err := j.Parse(ctx, tokenStr)
	if err != nil {
		return err
	}
	expired := time.Until(time.Unix(claims.ExpiresAt.Unix(), 0))

	return j.Store.Set(ctx, tokenStr, expired)

}

// Parse verifies a JWT token string and extracts the custom claims
func (j *jwtToken) Parse(ctx context.Context, tokenStr string) (*Claims, error) {

	if exists, err := j.Store.Check(ctx, tokenStr); err != nil {
		return nil, err
	} else if exists {
		return nil, errorsx.New("expired token")
	}
	token, err := jwt.ParseWithClaims(
		tokenStr, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.Key), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return nil, errorsx.WithStack(err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errorsx.Wrap(err, "invalid token")
}
