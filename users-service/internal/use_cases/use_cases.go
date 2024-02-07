package use_cases

import "github.com/droquedev/e-commerce/users-service/internal/entities"

type UserUseCases struct {
	userRepository entities.UserRepository
}

func NewUserUseCases(userRepo entities.UserRepository) *UserUseCases {
	return &UserUseCases{
		userRepository: userRepo,
	}
}
