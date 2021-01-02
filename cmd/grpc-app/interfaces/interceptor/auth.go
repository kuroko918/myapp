package interceptor

import (
	"context"

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

func NewAuthInterceptor(sqlHandler database.SqlHandler) *AuthInterceptor {
	return &AuthInterceptor{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (interceptor *AuthInterceptor) Auth(ctx context.Context) (newCtx context.Context, err error) {
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return
	}

	authToken, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return
	}

	token, err := client.VerifyIDToken(context.Background(), authToken)
	if err != nil {
		return
	}

	// DBに存在しない user は保存する
	u := domain.User{
		ID: token.UID,
		Name: token.Claims["name"].(string),
		Email: token.Claims["email"].(string),
		Avatar: token.Claims["picture"].(string),
	}
	_, err = interceptor.Interactor.GetOrAdd(u)
	if err != nil {
		return
	}

	newCtx = context.WithValue(ctx, "AuthenticatedUserId", token.UID)
	return
}
