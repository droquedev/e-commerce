package repository

import (
	"errors"

	"github.com/droquedev/e-commerce/users-service/internal/entities"
)

type UserPostgresRepository struct {
}

func NewUserPostgresRepository() entities.UserRepository {
	return &UserPostgresRepository{}
}

var users = []entities.User{
	{
		Username: "user1",
		Email:    "email@email.com",
		Password: "password",
	},
}

// FindUserByUsername implements entities.UserRepository.
func (*UserPostgresRepository) FindUserByUsername(username string) (*entities.User, error) {
	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

// Persist implements entities.UserRepository.
func (*UserPostgresRepository) Persist(user *entities.User) error {
	users = append(users, *user)
	return nil
}
