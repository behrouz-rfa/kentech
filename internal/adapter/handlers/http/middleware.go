package http

import (
	"strings"

	"github.com/behrouz-rfa/kentech/internal/core/common"
	"github.com/behrouz-rfa/kentech/internal/core/model"
	"github.com/behrouz-rfa/kentech/internal/core/ports"
	"github.com/gin-gonic/gin"
)

// authMiddleware is a middleware to check if the user is authenticated
func authMiddleware(token ports.Auth) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(common.AuthorizationHeaderKey)

		isEmpty := len(authorizationHeader) == 0
		if isEmpty {
			err := model.ErrEmptyAuthorizationHeader
			handleAbort(ctx, err)
			return
		}

		fields := strings.Fields(authorizationHeader)
		isValid := len(fields) == 2
		if !isValid {
			err := model.ErrInvalidAuthorizationHeader
			handleAbort(ctx, err)
			return
		}

		currentAuthorizationType := strings.ToLower(fields[0])
		if currentAuthorizationType != common.AuthorizationType {
			err := model.ErrInvalidAuthorizationType
			handleAbort(ctx, err)
			return
		}

		accessToken := fields[1]
		payload, err := token.Verify(accessToken)
		if err != nil {
			handleAbort(ctx, err)
			return
		}

		ctx.Set(common.AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}

// adminMiddleware is a middleware to check if the user is an admin
func adminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := GetAuthPayload(ctx, common.AuthorizationPayloadKey)

		if payload.UserID == "" {
			err := model.ErrForbidden
			handleAbort(ctx, err)
			return
		}

		ctx.Next()
	}
}

// 2211
