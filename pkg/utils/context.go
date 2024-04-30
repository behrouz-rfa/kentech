package utils

import (
	"context"

	"github.com/behrouz-rfa/kentech/internal/core/common"
)

func LoadersFromContext(ctx context.Context) []string {
	if loaders, ok := ctx.Value(common.PreloadersKey).([]string); ok {
		return loaders
	}

	return []string{}
}
