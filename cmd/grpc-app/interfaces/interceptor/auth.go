package interceptor

import (
	"context"
	"time"

	"google.golang.org/api/option"
	firebase "firebase.google.com/go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"github.com/kuroko918/myapp/cmd/grpc-app/domain"
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

	// DBに存在しない user は保存する
	timeNow := time.Now()
	u := domain.User{
		ID: token.UID,
		Name: token.Claims["name"].(string),
		Email: token.Claims["email"].(string),
		Avatar: token.Claims["picture"].(string),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
	_, err = interceptor.Interactor.GetOrAdd(ctx, u)
	if err != nil {
		return
	}

	newCtx = context.WithValue(ctx, "AuthenticatedUserId", token.UID)
	return
}
