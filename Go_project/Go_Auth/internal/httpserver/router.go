package httpserver

import (
	"go-auth/internal/app"
	"go-auth/internal/user"

	"github.com/gin-gonic/gin"
)

func NewRouter(a *app.App) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.GET("/health", health)
	userRepo := user.NewRepo(a.DB)
	userScv := user.NewService(userRepo, a.Config.JWT_SECRET)
	userHandler := user.NewHandler(userScv)
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	return r
}
