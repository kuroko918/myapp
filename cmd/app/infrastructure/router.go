package infrastructure

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/kuroko918/myapp/cmd/app/interfaces/controllers"
	"github.com/kuroko918/myapp/cmd/app/interfaces/middleware"
)

var Router *gin.Engine

func init() {
	router := gin.New()

	// Global Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// CORS
	router.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Access-Control-Allow-Headers", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowOrigins: []string{"http://localhost:3000"},
		MaxAge:       24 * time.Hour,
	}))

	// Authentification
	authMiddleware := middleware.NewAuthMiddleWare(NewSqlHandler())
	router.Use(authMiddleware.Auth())

	messagesController := controllers.NewMessagesController(NewSqlHandler())
	router.GET("/messages", func(c *gin.Context) { messagesController.Index(c) })
	router.POST("/message", func(c *gin.Context) { messagesController.Create(c) })
	router.PUT("/message/:id", func(c *gin.Context) { messagesController.Update(c) })
	router.DELETE("/message/:id", func(c *gin.Context) { messagesController.Delete(c) })

	Router = router
}
