package database

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/firestore"

	"github.com/kuroko918/myapp/cmd/grpc-app/domain"
)

type MessageRepository struct {
	DbHandler
}

func (repo *MessageRepository) DeleteById(ctx context.Context, message domain.Message) (err error) {
	_, err = repo.Collection("messages").Doc(message.ID).Delete(ctx)
	if err != nil {
		return
	}
	return
}

func (repo *MessageRepository) FindAll(ctx context.Context) (messages []domain.Message, err error) {
	messageSnaps, err := repo.Collection("messages").OrderBy("CreatedAt", firestore.Asc).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	for _, message := range messageSnaps {
		messageJson, err := json.Marshal(message.Data())
    if err != nil {
			return nil, err
		}

		var message domain.Message
		if err = json.Unmarshal(messageJson, &message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return
}

func (repo *MessageRepository) Store(ctx context.Context, m domain.Message) (message domain.Message, err error) {
	if _, err = repo.Collection("messages").Doc(m.ID).Set(ctx, m); err != nil {
		return
	}
	message = m
	return
}

func (repo *MessageRepository) Update(ctx context.Context, m map[string]interface{}) (message domain.Message, err error) {
	_, err = repo.Doc("messages/" + m["ID"].(string)).Set(ctx, m, firestore.MergeAll)
	if err != nil {
		return
	}

	messageSnap, err := repo.Doc("messages/" + m["ID"].(string)).Get(ctx)
	if err != nil {
		return
	}
	messageJson, err := json.Marshal(messageSnap.Data())
	if err != nil {
		return
	}
	if err = json.Unmarshal(messageJson, &message); err != nil {
		return
	}
	return
}
