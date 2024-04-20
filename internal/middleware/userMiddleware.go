package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stranik28/HackLoaylityProgramm/internal/helper"
	"github.com/stranik28/HackLoaylityProgramm/internal/logger"
	"github.com/stranik28/HackLoaylityProgramm/internal/service"
	"net/http"
	"strings"
)

func UserMiddleware(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(userAuthenticate())
	incomingRoutes.GET("/user/check_jwt", service.CheckJWT)
}

func userAuthenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			logger.Log.Warn("No token provided")
			c.IndentedJSON(500, "No Token Provided")
			c.Abort()
			return
		}
		if !strings.HasPrefix(clientToken, "Bearer ") {
			c.String(http.StatusUnauthorized, "Invalid authorization format")
			return
		}
		// Извлекаем только сам токен, убрав префикс "Bearer "
		clientToken = clientToken[len("Bearer "):]
		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			logger.Log.Error("The token is invalid")
			c.JSON(http.StatusInternalServerError, gin.H{"There's an error Message for you": err})
			c.Abort()
			return
		}

		c.Set("id", claims.Id)
		c.Next()
	}
}
