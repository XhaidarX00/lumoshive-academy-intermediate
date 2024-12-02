package middleware

import (
	"net/http"
	"project-voucher-team3/database"
	"project-voucher-team3/utils"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	Cacher database.Cacher
}

func NewMiddleware(cacher database.Cacher) Middleware {
	return Middleware{
		Cacher: cacher,
	}
}

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		id := c.GetHeader("ID-KEY")
		val, err := m.Cacher.Get(id)
		if err != nil {
			utils.ResponseError(c, "server error", "Token expired", http.StatusUnauthorized)
			c.Abort()
			return
		}

		if val == "" || val != token {
			utils.ResponseError(c, "Unauthorized", "Token is not valid", http.StatusUnauthorized)
			c.Abort()
			return
		}

		// before request
		c.Next()

	}
}
