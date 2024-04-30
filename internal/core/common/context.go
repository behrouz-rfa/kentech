package common

type ContextKey string

const (
	PreloadersKey           ContextKey = "preloaders"
	AuthorizationContextKey ContextKey = "Authorization"
	GinContextKey           ContextKey = "gin"

	// authorizationHeaderKey is the key for authorization header in the request
	AuthorizationHeaderKey = "authorization"
	// authorizationType is the accepted authorization type
	AuthorizationType = "bearer"
	// authorizationPayloadKey is the key for authorization payload in the context
	AuthorizationPayloadKey = "authorization_payload"
)

const GinContextKeyUserClaim string = "user"
