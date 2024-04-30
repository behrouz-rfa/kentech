package auth

import (
	"errors"
	"fmt"
	"github.com/behrouz-rfa/kentech/pkg/logger"

	"time"

	"github.com/behrouz-rfa/kentech/internal/core/model"
	"github.com/dgrijalva/jwt-go"
)

// Auth is an authentication service that handles token creation and verification.
type Auth struct {
	secret string
	lg     *logger.Entry
}

// NewAuth creates a new instance of the Auth service.
func NewAuth(secret string) *Auth {
	return &Auth{
		secret: secret,
		lg:     logger.General.Component("FilmService"),
	}
}

// Create generates a new JWT token with the provided user information.
func (a *Auth) Create(data model.TokenPayload) (*model.JWTToken, error) {
	expirationTime := time.Now().Add(3 * time.Hour)
	claims := &JWTClaim{
		UserID:   data.UserID,
		Username: data.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(a.secret))
	if err != nil {
		a.lg.WithError(err).Error("failed to sign JWT token")
		return nil, fmt.Errorf("failed to sign JWT token: %w", err)
	}

	return &model.JWTToken{Token: signedToken, ExpirationTime: expirationTime}, nil
}

// Verify validates the provided JWT token and returns the user information.
func (a *Auth) Verify(signedToken string) (*model.TokenPayload, error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			a.lg.Error("invalid signing method used in JWT token")
			return nil, errors.New("invalid signing method used in JWT token")
		}
		return []byte(a.secret), nil
	})
	if err != nil {
		a.lg.WithError(err).Error("failed to parse JWT token")
		return nil, fmt.Errorf("failed to parse JWT token: %w", err)
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		a.lg.Error("invalid JWT token")
		return nil, errors.New("invalid JWT token")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		a.lg.Error("JWT token has expired")
		return nil, errors.New("JWT token has expired")
	}

	return &model.TokenPayload{
		UserID:   claims.UserID,
		Username: claims.Username,
	}, nil
}

// JWTClaim represents the claims in a JWT token.
type JWTClaim struct {
	Username string `json:"username"`
	UserID   string `json:"userID"`
	jwt.StandardClaims
}
