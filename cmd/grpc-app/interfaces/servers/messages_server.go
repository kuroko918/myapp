package servers

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/kuroko918/myapp/cmd/grpc-app/domain"
	"github.com/kuroko918/myapp/cmd/grpc-app/interfaces/database"
	"github.com/kuroko918/myapp/cmd/grpc-app/usecase"

	messagepb "github.com/kuroko918/myapp/cmd/grpc-app/proto/message"
)

type MessagesServer struct {
	Interactor usecase.MessageInteractor
	messagepb.UnimplementedMessageServiceServer
}

func NewMessagesServer(sqlHandler database.SqlHandler) *MessagesServer {
	return &MessagesServer{
		Interactor: usecase.MessageInteractor{
			MessageRepository: &database.MessageRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (server *MessagesServer) PostMessage(ctx context.Context, req *messagepb.PostMessageParams) (message *messagepb.Message, err error) {
	m := domain.Message{
		Content: req.GetContent(),
		UserId: ctx.Value("AuthenticatedUserId").(string),
	}
	m, err = server.Interactor.Add(m)
	if err != nil {
		return
	}

	message, err = ProtoMessage(m)
	if err != nil {
		return
	}
	return
}

func (server *MessagesServer) PutMessage(ctx context.Context, req *messagepb.PutMessageParams) (message *messagepb.Message, err error) {
	m := domain.Message{
		ID: req.GetId(),
	}
	m, err = server.Interactor.Update(m, "Content", req.GetContent())
	if err != nil {
		return
	}

	message, err = ProtoMessage(m)
	if err != nil {
		return
	}
	return
}

func (server *MessagesServer) DeleteMessage(ctx context.Context, req *messagepb.DeleteMessageParams) (_ *emptypb.Empty, err error) {
	message := domain.Message{
		ID: req.GetId(),
	}
	err = server.Interactor.DeleteById(message)
	if err != nil {
		return new(emptypb.Empty), err
	}
	return new(emptypb.Empty), nil
}

func (server *MessagesServer) GetMessages(ctx context.Context, _ *emptypb.Empty) (messages *messagepb.MessageList, err error) {
	var message *messagepb.Message
	var messageList []*messagepb.Message

	ms, err := server.Interactor.Messages()
	if err != nil {
		return
	}

	for _, m := range ms {
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
