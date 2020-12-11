package middleware

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"

	"github.com/kuroko918/myapp/cmd/app/interfaces/controllers"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		opt := option.WithCredentialsFile("./serviceAccountKey.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			c.JSON(500, controllers.NewError(err))
			c.Abort()
			return
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			c.JSON(500, controllers.NewError(err))
			c.Abort()
			return
		}

		authToken := c.Request.Header.Get("Authorization")

		token, err := client.VerifyIDToken(context.Background(), authToken)
		if err != nil {
			c.JSON(500, controllers.NewError(err))
			c.Abort()
		}
		c.Set("AuthenticatedUserId", token.UID)
	}
}
