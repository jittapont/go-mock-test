package main

type User struct {
	Name string
}

//go:generate mockgen -destination=./mocks/mock_repository.go -package=mocks main UserRepository
type UserRepository interface {
	InsertUser(user *User) error
	GetUsers() ([]*User, error)
}

func main() {

}
