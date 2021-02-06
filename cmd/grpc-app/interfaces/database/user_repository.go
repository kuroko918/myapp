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

// func (repo *UserRepository) GetOrStore(ctx context.Context, u domain.User) (user domain.User, err error) {
// 	if _, err = repo.Collection("users").Doc(u.ID).Set(ctx, u); err != nil {
// 		return
// 	}
// 	user = u
// 	return
// }
