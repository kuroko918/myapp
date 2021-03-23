package database

import (
	"context"
	"encoding/json"

	"github.com/kuroko918/myapp/cmd/grpc-app/domain"
)

type UserRepository struct {
	DbHandler
}

func (repo *UserRepository) Get(ctx context.Context, userId string) (user domain.User, err error) {
	userSnap, err := repo.Collection("users").Doc(userId).Get(ctx)
	if err != nil {
		return
	}
	userJson, err := json.Marshal(userSnap.Data())
	if err != nil {
		return
	}

	if err = json.Unmarshal(userJson, &user); err != nil {
		return
	}
	return
}

func (repo *UserRepository) UpdateAll(ctx context.Context, u domain.User) (user domain.User, err error) {
	_, err = repo.Doc("users/" + u.ID).Set(ctx, u)
	if err != nil {
		return
	}

	userSnap, err := repo.Doc("users/" + u.ID).Get(ctx)
	if err != nil {
		return
	}
	userJson, err := json.Marshal(userSnap.Data())
	if err != nil {
		return
	}
	if err = json.Unmarshal(userJson, &user); err != nil {
		return
	}
	return
}
