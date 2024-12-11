// middlewares/login_limiter_middleware.go
package middlewares

import (
	"log"
	"net"
	"project_auth_jwt/database"

	"github.com/gin-gonic/gin"
)

type LoginLimiterMiddlewareInterface interface {
	Limit() gin.HandlerFunc
	IPWhitelistMiddleware() gin.HandlerFunc
}

type LoginLimiterMiddleware struct {
	Cacher database.Cacher
}

func NewLoginLimiterMiddleware(cacher database.Cacher) *LoginLimiterMiddleware {
	return &LoginLimiterMiddleware{
		Cacher: cacher,
	}
}

func (m *LoginLimiterMiddleware) IPWhitelistMiddleware() gin.HandlerFunc {
	whitelistedIPs := []string{
		"http://localhost/",
	}

	return func(c *gin.Context) {
		clientIP, _, err := net.SplitHostPort(c.Request.RemoteAddr)
		if err != nil {
			c.AbortWithStatus(403)
			return
		}

		allowed := false
		for _, ip := range whitelistedIPs {
			if clientIP == ip {
				allowed = true
				break
			}
		}

		if !allowed {
			c.AbortWithStatus(403)
			return
		}

		c.Next()
	}
}

func (m *LoginLimiterMiddleware) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP, _, err := net.SplitHostPort(c.Request.RemoteAddr)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}

		log.Println(clientIP)

		key := "login_attempts:" + clientIP

		// Cek apakah IP sudah diblokir
		blockedKey := "blocked_ip:" + clientIP
		if m.Cacher.Exsist(blockedKey) != nil {
			c.AbortWithStatus(429)
			return
		}

		// Increment login attempts
		attempts, err := m.Cacher.Incr(key)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}

		// Set expiration untuk tracking attempts
		m.Cacher.Expire(key, 1)

		if attempts > 3 {
			// Blokir IP selama 10 menit
			m.Cacher.SetExpire(blockedKey, "1", 10)
			c.AbortWithStatus(429)
			return
		}

		c.Next()
	}
}
