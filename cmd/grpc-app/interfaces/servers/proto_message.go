package servers

import (
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"

	"github.com/kuroko918/myapp/cmd/grpc-app/domain"

	messagepb "github.com/kuroko918/myapp/cmd/grpc-app/proto/message"
	userpb "github.com/kuroko918/myapp/cmd/grpc-app/proto/user"
)

func ProtoMessage(m domain.Message) (message *messagepb.Message, err error) {
	ucat, err := ptypes.TimestampProto(m.User.CreatedAt)
	if err != nil {
		err = status.Errorf(codes.Internal, "created_at cannot be parsed to timestamp")
		return
	}
	uuat, err := ptypes.TimestampProto(m.User.UpdatedAt)
	if err != nil {
		err = status.Errorf(codes.Internal, "updated_at cannot be parsed to timestamp")
		return
	}

	mcat, err := ptypes.TimestampProto(m.CreatedAt)
	if err != nil {
		err = status.Errorf(codes.Internal, "created_at cannot be parsed to timestamp")
		return
	}
	muat, err := ptypes.TimestampProto(m.UpdatedAt)
	if err != nil {
		err = status.Errorf(codes.Internal, "updated_at cannot be parsed to timestamp")
		return
	}

	user := &userpb.User{
		Id:        m.User.ID,
		Name:      m.User.Name,
		Email:     m.User.Email,
		Avatar:    m.User.Avatar,
		CreatedAt: ucat,
		UpdatedAt: uuat,
	}

	message = &messagepb.Message{
		Id:        m.ID,
		Content:   m.Content,
		UserId:    m.UserId,
		CreatedAt: mcat,
		UpdatedAt: muat,
		User:      user,
	}
	return
}
