package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stranik28/HackLoaylityProgramm/internal/middleware"
	"github.com/stranik28/HackLoaylityProgramm/internal/service"
)

func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/auth", service.AuthHandler)
	middleware.UserMiddleware(r)
	middleware.AdminMiddleware(r)
	return r
}
