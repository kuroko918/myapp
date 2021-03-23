package servers

import (
	"context"
	"time"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	currentUserId := ctx.Value("AuthenticatedUserId").(string)
	userId := req.GetId()
	if currentUserId != userId {
		err = status.Errorf(codes.PermissionDenied, "You have no authorization")
		return
	}

	createdAt, err := ptypes.Timestamp(req.GetCreatedAt())
	if err != nil {
		err = status.Errorf(codes.Internal, "created_at cannot be parsed to timestamp")
		return
	}

	u := domain.User{
		ID: userId,
		Name: req.GetName(),
		Email: req.GetEmail(),
		Avatar: req.GetAvatar(),
		CreatedAt: createdAt,
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
