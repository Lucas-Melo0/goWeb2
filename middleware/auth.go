package middleware

import (
	"main/pkg/store/web"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("token") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, ""))
			c.Abort()
			return
		}
		c.Next()
	}
}
