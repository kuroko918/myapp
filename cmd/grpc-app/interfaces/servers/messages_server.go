package servers

import (
	"context"
	"time"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/kuroko918/myapp/cmd/grpc-app/domain"
	"github.com/kuroko918/myapp/cmd/grpc-app/interfaces/database"
	"github.com/kuroko918/myapp/cmd/grpc-app/usecase"

	messagepb "github.com/kuroko918/myapp/cmd/grpc-app/proto/message"
)

type MessagesServer struct {
	MessageInteractor usecase.MessageInteractor
	UserInteractor usecase.UserInteractor
	messagepb.UnimplementedMessageServiceServer
}

func NewMessagesServer(dbHandler database.DbHandler) *MessagesServer {
	return &MessagesServer{
		MessageInteractor: usecase.MessageInteractor{
			MessageRepository: &database.MessageRepository{
				DbHandler: dbHandler,
			},
		},
		UserInteractor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				DbHandler: dbHandler,
			},
		},
	}
}

func (server *MessagesServer) PostMessage(ctx context.Context, req *messagepb.PostMessageParams) (message *messagepb.Message, err error) {
	userId := ctx.Value("AuthenticatedUserId").(string)

	id, _ := uuid.NewRandom()
	timeNow := time.Now()
	m := domain.Message{
		ID: id.String(),
		Content: req.GetContent(),
		UserId: userId,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
	m, err = server.MessageInteractor.Add(ctx, m)
	if err != nil {
		return
	}

	user, err := server.UserInteractor.User(ctx, userId)
	if err != nil {
		return
	}
	m.User = user
	message, err = ProtoMessage(m)
	if err != nil {
		return
	}
	return
}

func (server *MessagesServer) PatchMessage(ctx context.Context, req *messagepb.PatchMessageParams) (message *messagepb.Message, err error) {
	currentUserId := ctx.Value("AuthenticatedUserId").(string)
	messageId := req.GetId()
	m, err := server.MessageInteractor.Message(ctx, messageId)
	if err != nil {
		return
	}
	if currentUserId != m.UserId {
		err = status.Errorf(codes.PermissionDenied, "You have no authorization")
		return
	}

	messageMap := map[string]interface{}{
		"ID": messageId,
		"Content": req.GetContent(),
		"UpdatedAt": time.Now(),
	}
	m, err = server.MessageInteractor.Update(ctx, messageMap)
	if err != nil {
		return
	}

	user, err := server.UserInteractor.User(ctx, currentUserId)
	if err != nil {
		return
	}
	m.User = user
	message, err = ProtoMessage(m)
	if err != nil {
		return
	}
	return
}

func (server *MessagesServer) DeleteMessage(ctx context.Context, req *messagepb.DeleteMessageParams) (_ *emptypb.Empty, err error) {
	currentUserId := ctx.Value("AuthenticatedUserId").(string)
	messageId := req.GetId()
	m, err := server.MessageInteractor.Message(ctx, messageId)
	if err != nil {
		return
	}
	if currentUserId != m.UserId {
		err = status.Errorf(codes.PermissionDenied, "You have no authorization")
		return
	}

	message := domain.Message{
		ID: messageId,
	}
	err = server.MessageInteractor.DeleteById(ctx, message)
	if err != nil {
		return new(emptypb.Empty), err
	}
	return new(emptypb.Empty), nil
}

func (server *MessagesServer) GetMessages(ctx context.Context, _ *emptypb.Empty) (messages *messagepb.MessageList, err error) {
	var message *messagepb.Message
	var messageList []*messagepb.Message

	ms, err := server.MessageInteractor.Messages(ctx)
	if err != nil {
		return
	}

	for _, m := range ms {
		var user domain.User
		user, err = server.UserInteractor.User(ctx, m.UserId)
		if err != nil {
			return
		}
		m.User = user
		message, err = ProtoMessage(m)
		if err != nil {
			return
		}
		messageList = append(messageList, message)
	}

	messages = &messagepb.MessageList{
		Messages: messageList,
	}
	return
}
