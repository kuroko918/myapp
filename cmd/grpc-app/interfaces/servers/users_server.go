package servers

import (
	"context"
	"time"

	"github.com/kuroko918/myapp/cmd/grpc-app/domain"
	"github.com/kuroko918/myapp/cmd/grpc-app/interfaces/database"
	"github.com/kuroko918/myapp/cmd/grpc-app/usecase"

	userpb "github.com/kuroko918/myapp/cmd/grpc-app/proto/user"
)

type UsersServer struct {
	UserInteractor usecase.UserInteractor
	userpb.UnimplementedUserServiceServer
}

func NewUsersServer(dbHandler database.DbHandler) *UsersServer {
	return &UsersServer{
		UserInteractor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				DbHandler: dbHandler,
			},
		},
	}
}

func (server *UsersServer) PutUser(ctx context.Context, req *userpb.PutUserParams) (user *userpb.User, err error) {
	u := domain.User{
		ID: req.GetId(),
		Name: req.GetName(),
		Email: req.GetEmail(),
		Avatar: req.GetAvatar(),
		UpdatedAt: time.Now(),
	}
	m, err := server.UserInteractor.UpdateAll(ctx, u)
	if err != nil {
		return
	}

	user, err = ProtoUser(m)
	if err != nil {
		return
	}
	return
}

func (server *UsersServer) GetUser(ctx context.Context, req *userpb.GetUserParams) (user *userpb.User, err error) {
	userId := req.GetId()
	u, err := server.UserInteractor.User(ctx, userId)

	user, err = ProtoUser(u)
	if err != nil {
		return
	}
	return
}
