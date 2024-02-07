package use_cases

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/droquedev/e-commerce/pkg/nats"
	"github.com/droquedev/e-commerce/users/internal/entities"
	"github.com/nats-io/stan.go"
)

var natsConn stan.Conn

func (u *UserUseCases) UserCreator(user *entities.User) error {

	var err error
	// filters := []shared_domain.Filters{
	// 	{Field: "username", Operator: "=", Value: user.Username, Logic: "OR"},
	// 	{Field: "email", Operator: "=", Value: user.Email},
	// }

	natsConn := nats.GetNatsConn().NatsConn

	if err != nil {
		log.Fatal(err)
	}

	users, _ := u.userRepository.FindUserByUsername(user.Username)

	if users != nil {
		return errors.New("username or email already exists")
	}

	err2 := u.userRepository.Persist(user)

	if err2 != nil {
		return err2
	}

	data, err := json.Marshal(user)

	if err != nil {
		return err
	}

	natsConn.Publish("user:created", data)

	return nil
}
