package middleware

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"

	"github.com/kuroko918/myapp/cmd/app/utilities"
	"github.com/kuroko918/myapp/cmd/app/domain"
	"github.com/kuroko918/myapp/cmd/app/interfaces/database"
	"github.com/kuroko918/myapp/cmd/app/usecase"
)

type AuthMiddleWare struct {
	Interactor usecase.UserInteractor
}

func NewAuthMiddleWare(sqlHandler database.SqlHandler) *AuthMiddleWare {
	return &AuthMiddleWare{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (middleware *AuthMiddleWare) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		opt := option.WithCredentialsFile("./serviceAccountKey.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			c.JSON(500, utilities.NewError(err))
			c.Abort()
			return
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			c.JSON(500, utilities.NewError(err))
			c.Abort()
			return
		}

		authToken := c.Request.Header.Get("Authorization")

		token, err := client.VerifyIDToken(context.Background(), authToken)
		if err != nil {
			c.JSON(500, utilities.NewError(err))
			c.Abort()
			return
		}

		// DBに存在しない user は保存
		u := domain.User{
			ID: token.UID,
			Name: token.Claims["name"].(string),
			Email: token.Claims["email"].(string),
			Avater: token.Claims["picture"].(string),
		}
		_, err = middleware.Interactor.GetOrAdd(u)
		if err != nil {
			c.JSON(500, utilities.NewError(err))
			c.Abort()
			return
		}

		c.Set("AuthenticatedUserId", token.UID)
	}
}
