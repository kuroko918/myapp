package interceptor

import (
	"context"

	"google.golang.org/api/option"
	firebase "firebase.google.com/go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"github.com/kuroko918/myapp/cmd/grpc-app/interfaces/database"
	"github.com/kuroko918/myapp/cmd/grpc-app/usecase"
)

type AuthInterceptor struct {
	Interactor usecase.UserInteractor
}

func NewAuthInterceptor(dbHandler database.DbHandler) *AuthInterceptor {
	return &AuthInterceptor{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				DbHandler: dbHandler,
			},
		},
	}
}

func (interceptor *AuthInterceptor) Auth(ctx context.Context) (newCtx context.Context, err error) {
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return
	}

	authToken, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return
	}

	token, err := client.VerifyIDToken(ctx, authToken)
	if err != nil {
		return
	}

	newCtx = context.WithValue(ctx, "AuthenticatedUserId", token.UID)
	return
}
