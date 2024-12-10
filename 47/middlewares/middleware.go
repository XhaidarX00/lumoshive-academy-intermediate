package middlewares

import (
	"project_auth_jwt/database"
	"project_auth_jwt/services"

	"go.uber.org/zap"
)

type Middleware struct {
	Auth  AuthMiddlewareInterface
	Limit LoginLimiterMiddlewareInterface
}

func NewMiddleware(log *zap.Logger, cacher database.Cacher, jwt services.JWTServiceInterface) Middleware {
	return Middleware{
		Auth:  NewAuthMiddleware(jwt),
		Limit: NewLoginLimiterMiddleware(cacher),
	}
}
