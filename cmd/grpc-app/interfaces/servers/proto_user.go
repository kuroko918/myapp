package servers

import (
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"

	"github.com/kuroko918/myapp/cmd/grpc-app/domain"

	userpb "github.com/kuroko918/myapp/cmd/grpc-app/proto/user"
)

func ProtoUser(u domain.User) (user *userpb.User, err error) {
	ucat, err := ptypes.TimestampProto(u.CreatedAt)
	if err != nil {
		err = status.Errorf(codes.Internal, "created_at cannot be parsed to timestamp")
		return
	}
	uuat, err := ptypes.TimestampProto(u.UpdatedAt)
	if err != nil {
		err = status.Errorf(codes.Internal, "updated_at cannot be parsed to timestamp")
		return
	}

	user = &userpb.User{
		Id:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Avatar:    u.Avatar,
		CreatedAt: ucat,
		UpdatedAt: uuat,
	}
	return
}
