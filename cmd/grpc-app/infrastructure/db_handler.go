package infrastructure

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"github.com/kuroko918/myapp/cmd/grpc-app/interfaces/database"
)

func NewDbHandler() database.DbHandler {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}

	dbHandler := new(DbHandler)
	dbHandler.Client = client
	return dbHandler
}

type DbHandler struct {
	Client *firestore.Client
}

func (handler *DbHandler) Collection(path string) *firestore.CollectionRef {
	return handler.Client.Collection(path)
}

func (handler *DbHandler) Doc(path string) *firestore.DocumentRef {
	return handler.Client.Doc(path)
}
