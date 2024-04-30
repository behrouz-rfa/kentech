package ports

import (
	"github.com/behrouz-rfa/kentech/internal/core/model"
)

type Auth interface {
	Create(info model.TokenPayload) (*model.JWTToken, error)
	Verify(token string) (*model.TokenPayload, error)
}
