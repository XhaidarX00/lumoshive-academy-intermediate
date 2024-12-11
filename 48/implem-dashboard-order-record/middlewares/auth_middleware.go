// middlewares/auth_middleware.go
package middlewares

import (
	"project_auth_jwt/services"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareInterface interface {
	Authenticate() gin.HandlerFunc
}

type AuthMiddleware struct {
	jwtService services.JWTServiceInterface
}

func NewAuthMiddleware(jwtService services.JWTServiceInterface) AuthMiddlewareInterface {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatus(401)
			return
		}

		claims, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
