package auth

import (
	"errors"
	"time"

	"github.com/behrouz-rfa/kentech/internal/core/model"
	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	secret string
}

func NewAuth(secret string) *Auth {
	return &Auth{
		secret: secret,
	}
}

func (a *Auth) Create(data model.TokenPayload) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		UserID:   data.UserID,
		Username: data.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.secret)
}

func (a *Auth) Verify(signedToken string) (*model.TokenPayload, error) {
	claim := &JWTClaim{}
	token, err := jwt.ParseWithClaims(
		signedToken,
		claim,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(a.secret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}

	return &model.TokenPayload{
		UserID:   claim.UserID,
		Username: claim.Username,
	}, nil
}

type JWTClaim struct {
	Username string `json:"username"`
	UserID   string `json:"userID"`
	jwt.StandardClaims
}
