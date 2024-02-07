package entities

type UserRepository interface {
	FindUserByUsername(username string) (*User, error)
	Persist(user *User) error
}
