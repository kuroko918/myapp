package database

import (
	"github.com/kuroko918/myapp/cmd/app/domain"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) GetOrStore(u domain.User) (user domain.User, err error) {
	if err = repo.FirstOrCreate(&u).Error; err != nil {
		return
	}
	user = u
	return
}
