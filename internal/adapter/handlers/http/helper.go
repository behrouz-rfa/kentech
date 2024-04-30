package http

import (
	"github.com/behrouz-rfa/kentech/internal/core/model"
	"github.com/gin-gonic/gin"
)

// GetAuthPayload is a helper function to get the auth payload from the context
func GetAuthPayload(ctx *gin.Context, key string) *model.TokenPayload {
	return ctx.MustGet(key).(*model.TokenPayload)
}

// toMap is a helper function to add meta and data to a map
func toMap(m meta, data any, key string) map[string]any {
	return map[string]any{
		"meta": m,
		key:    data,
	}
}
